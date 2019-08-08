package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
	_ "go_gin_example/docs"
	"go_gin_example/middleware/jwt"
	"go_gin_example/routers/api"
	"go_gin_example/utils/export"
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
	V1.Use(jwt.JWT())
	{
		// create user
		V1.POST("/users", func(c *gin.Context){
			c.JSON(200, gin.H{
				"code":1,
			})
		})
	}
	return r

}