package service

import (
	"encoding/json"
	"go_gin_example/models"
	"go_gin_example/service/cache_service"
	"go_gin_example/utils/redis"
	"log"
)

func AddUser(username, password string) (user *models.User, err error) {
	user = &models.User{
		Username:username,
		Password:password,
	}

	err = user.AddUser()
	if err != nil{
		return
	}else{
		return user, nil
	}
}


func UpdateUser(id int, u models.UUser)error{
	err := models.UpdateUser(id, u)
	return err
}

func DeleteUser(id int)error{
	key := cache_service.GetUserKey(id)
	exists := redis.Exists(key)
	if exists{
		redis.Delete(key)
	}
	err := models.DeleteUser(id)
	if err != nil{
		return err
	}
	return nil
}

func GetUser(id int) (*models.User, error){
	// 先查询redis
	var user models.User
	var err error
	key := cache_service.GetUserKey(id)
	exists := redis.Exists(key)
	if exists{
		data, err := redis.Get(key)
		if err != nil{
			log.Println("get redis key err:", err)
			return nil, err
		}else{
			err = json.Unmarshal(data, &user)
			if err != nil{
				return nil, err
			}

		}
	}else{
		// 查询数据库
		user , err := models.GetUser(id)
		if err != nil{
			log.Println("select user err:", err)
			return nil, err
		}else{
			// set cache key
			redis.Set(key, user,redis.CACHE_USER_EXPIRE )
			return user, nil
		}
	}
	// return &user, err
	return  &user, err
}


func GetUsers(page, limit int) ([]models.User, error){
	var users []models.User
	var err error

	key := cache_service.GetUsersKey(page, limit)
	exists := redis.Exists(key)
	if exists{
		// 查询redis
		data, err := redis.Get(key)
		if err != nil{
			log.Println("get redis key err:", err)
			return nil, err
		}else{
			err = json.Unmarshal(data, &users)
			if err != nil{
				return nil, err
			}

		}
	}else{
		//查询数据库
		users, err = models.GetUsers(page, limit)
		if err != nil{
			return nil, err
		}
		if len(users) > 0{
			// 存入Redis
			redis.Set(key, users, redis.CACHE_USERS_EXPIRE)
		}

	}
	return users, err
}