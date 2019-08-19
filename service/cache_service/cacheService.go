package cache_service

import (
	"go_gin_example/utils/redis"
	"fmt"
	"strconv"
)

func GetUserKey(id int) string{
	sid := strconv.Itoa(id)
	return redis.CACHE_USER + "_" + sid
}

func GetUsersKey(page, limit int) string {
	str := redis.CACHE_USERS + "_LIST_PAGE_%d_LIMIT_%d"
	return fmt.Sprintf(str, page, limit)
}