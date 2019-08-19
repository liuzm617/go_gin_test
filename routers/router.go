package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
	_ "go_gin_example/docs"
	"go_gin_example/middleware/pageView"
	"go_gin_example/routers/api"
	"go_gin_example/routers/api/v1"
	"go_gin_example/utils/export"
	"go_gin_example/utils/logging"
	"go_gin_example/utils/qrcode"
	"go_gin_example/utils/upload"
	"net/http"
)

// init router
func Router() *gin.Engine{
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	r.StaticFS("/export", http.Dir(export.GetExcelFullPath()))
	r.StaticFS("/upload/images", http.Dir(upload.GetImageFullPath()))
	r.StaticFS("/qrcode", http.Dir(qrcode.GetQrCodeFullPath()))

	r.POST("/auth", api.GetToken)
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	//r.POST("/upload", api.UploadImage)
	V1 := r.Group("/api/v1")
	//V1.Use(jwt.JWT())
	V1.Use(pageView.PageView())
	{
		// create user
		V1.POST("/users", v1.AddUser)
		V1.GET("/users", v1.GetUsers)
		V1.GET("/users/:id", v1.GetUser)
		V1.DELETE("/users/:id", v1.DeleteUser)
		V1.PATCH("/users/:id", v1.UpdateUser)

		V1.GET("/test_log", func(c *gin.Context) {
			logging.Info("test log")
			path := c.Request.URL.String()
			c.String(200, path)
		})
	}
	return r

}