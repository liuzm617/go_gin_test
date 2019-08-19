package pageView

import (
	"github.com/gin-gonic/gin"
	"go_gin_example/utils/logging"
	"go_gin_example/utils/redis"
)

//count page view 
func PageView() gin.HandlerFunc{
	return func(c *gin.Context) {
		path := c.Request.URL.String()

		exists := redis.Exists(path)
		if exists {
			err := redis.Incr(path)
			if err != nil {
				logging.Error(err)
			}
		}else{
			err := redis.Set(path, 1, 0)
			if err!= nil{
				logging.Error(err)
			}
		}
	}
}