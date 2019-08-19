package api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/astaxie/beego/validation"
	"go_gin_example/middleware/jwt"
	"go_gin_example/models"
	"go_gin_example/utils/app"
	"net/http"
)

//type Auth struct{
//	Username string `json:"username" valid:"Required; MaxSize(50)"`
//	Password string `json:"password" valid:"Required; MaxSize(50)"`
//}

// @Summary Get Auth token
// @Accept json
// @Produce json
// @param auth body models.Auth true "auth"
// @Success 200 {object} app.Response
// @Failure 400 {object} app.Response
// @Failure 401 {object} app.Response
// @Failure 404 {object} app.Response
// @Failure 500 {object} app.Response
// @Router /auth [post]
func GetToken(c *gin.Context){
	var u models.Auth
	appG := app.Gin{C: c}
	err := c.Bind(&u)
	if err != nil{
		appG.Response(http.StatusBadRequest,1, fmt.Sprint(err),nil  )
		return
	}
	valid := validation.Validation{}
	ok, _:= valid.Valid(&u)
	if !ok{
		appG.Response(http.StatusBadRequest, 1,fmt.Sprint(valid.Errors[0]), nil)
		return
	}
	isExist, err := models.CheckAuth(u.Username, u.Password)


	if err!= nil {
		appG.Response(http.StatusInternalServerError, 1, fmt.Sprint(err), nil)
		return
	}

	if !isExist{
		appG.Response(http.StatusNotFound, 1, "the username and password did not match any user", nil)
		return
	}

	token, err := jwt.GenerateToken(u.Username, u.Password)
	if err != nil{
		appG.Response(http.StatusInternalServerError, 1, fmt.Sprint(err), nil)
		return
	}

	appG.Response(http.StatusOK, 0, "", map[string]string{
		"token": token,
	})

}