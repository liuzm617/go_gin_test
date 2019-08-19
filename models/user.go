package models

import (
	"github.com/jinzhu/gorm"
	"github.com/pkg/errors"
	"fmt"
	"log"
)

type User struct {
	Model
	Username string `gorm:"type:varchar(50);unique_index" json:"username"`
	Password string `json:"password"`

}

var DeletedOn  = "deleted_on=?"

// 获取token,post struct
// 创建user, post struct
type Auth struct {
	Username string `json:"username" valid:"Required; MaxSize(50)"`
	Password string `json:"password" valid:"Required; MaxSize(50)"`

}

//update user post struct
type UUser struct {
	Username  string `json:"username"`
	Password string  `json:"password"`
}

// check user exists
func CheckAuth(username, password string)(bool, error){
	var user User
	err := db.Select("id").Where(&Auth{Username:username, Password:password}).First(&user).Error
	if err != nil && err != gorm.ErrRecordNotFound{
		return false, err
	}
	if user.ID > 0{
		return true, nil
	}
	return false, nil
}


// 验证用户名是否存在
func checkUser(username string)(bool, error){
	var user User
	err := db.Select("id").Where(&Auth{Username:username}).First(&user).Error
	if err != nil && err != gorm.ErrRecordNotFound{
		return false, err
	}
	if user.ID > 0{
		return true, nil
	}
	return false, nil
}


func AddUser(username, password string) (user *User, err error){
	user = &User{
		Username:username,
		Password:password,
	}
	exists, _ := checkUser(username)
	if exists{

		err = errors.New(fmt.Sprintf("user:%s exists", username))
		return
	}

	if err = db.Create(user).Error; err != nil{
		return
	}
	return
}

func (user *User) AddUser() (err error) {

	exists, _ := checkUser(user.Username)
	if exists{
		err = errors.New(fmt.Sprintf("user:%s exists", user.Username))
		return err
	}

	if err = db.Create(user).Error; err != nil{
		return err
	}
	return nil
}

func GetUser(id int)(*User, error){
	var user User
	err := db.Where("id = ? ", id).Where(DeletedOn, "").First(&user).Error
	if err != nil && err != gorm.ErrRecordNotFound{
		log.Println("get user err:", err)
		return nil, err
	}else if err == gorm.ErrRecordNotFound{
		return nil, err
	}
	return &user, nil

}

func GetUsers(page, limit int) ([]User, error) {
	var users []User
	err := db.Where(DeletedOn, "").Offset((page-1) * limit).Limit(limit).Find(&users).Error
	return users, err
}

func DeleteUser(id int)error{
	user := User{}
	err := db.Where("id=?", id).Delete(&user).Error
	if err != nil{
		return err
	}
	return nil
}

func UpdateUser(id, data interface{}) error{
	var user User
	err := db.Model(&user).Where("id=?", id).Where(DeletedOn, "").Update(data).Error
	if err != nil{
		return err
	}
	return nil
}