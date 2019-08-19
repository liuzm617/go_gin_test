package v1

import (
	"github.com/astaxie/beego/validation"
	"github.com/gin-gonic/gin"
	"go_gin_example/models"
	"go_gin_example/service"
	"go_gin_example/utils/app"
	"fmt"
	"go_gin_example/utils/common"
	"go_gin_example/utils/conf"
	"net/http"
	"strconv"
)

// @Summary create user
// @Accept json
// @Produce json
// @param user body models.Auth true "user"
// @Success 200 {object} app.Response
// @Failure 400 {object} app.Response
// @Failure 404 {object} app.Response
// @Failure 500 {object} app.Response
func AddUser(c *gin.Context) {
	var u models.Auth
	appG := app.Gin{C: c}
	err := c.Bind(&u)
	if err != nil{
		appG.Response(http.StatusBadRequest, 1, fmt.Sprint(err), nil)
		return
	}
	valid := validation.Validation{}
	ok, _ := valid.Valid(&u)
	if !ok{
		appG.Response(http.StatusBadRequest, 1, fmt.Sprint(valid.Errors[0]), nil)
		return
	}

	user, err := service.AddUser(u.Username, u.Password)
	if err != nil{
		appG.Response(http.StatusBadRequest, 1, fmt.Sprint(err), nil)
		return
	}
	appG.Response(http.StatusOK, 0, "", user)
}



// @Summary get user
// @Accept json
// @Produce json
// @param id path string true "id"
// @Success 200 {object} app.Response
// @Failure 400 {object} app.Response
// @Failure 404 {object} app.Response
// @Failure 500 {object} app.Response
func GetUser(c *gin.Context){
	appG := app.Gin{C: c}
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil{
		appG.Response(http.StatusNotFound, 1, "", nil)
	}
	user , err := service.GetUser(id)

	if err != nil{
		appG.Response(http.StatusNotFound, 1, fmt.Sprint(err), nil)
	}else{
		appG.Response(http.StatusOK, 0, "", user)
	}
}

// @Summary get users
// @Accept json
// @Produce json
// @Param page query int false "page"
// @Param limit query int false "limit"
// @Success 200 {object} app.Response
// @Failure 400 {object} app.Response
// @Failure 404 {object} app.Response
// @Failure 500 {object} app.Response
func GetUsers(c *gin.Context){
	appG := app.Gin{C: c}

	page := c.Query("page")
	limit := c.Query("limit")

	pageNum := common.QueryInt(page, 1)
	limitNum := common.QueryInt(limit, conf.AppConf.PageSize)

	users, err := service.GetUsers(pageNum, limitNum)
	if err != nil{
		appG.Response(http.StatusInternalServerError, 1, fmt.Sprint(err), nil)
		return
	}
	appG.Response(http.StatusOK, 0, "", users)

}

// @Summary del users
// @Accept json
// @Produce json
// @Param id path int True "id"
// @Success 200 {object} app.Response
// @Failure 400 {object} app.Response
// @Failure 404 {object} app.Response
// @Failure 500 {object} app.Response
func DeleteUser(c *gin.Context){
	appG := app.Gin{C: c}
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil{
		appG.Response(http.StatusNotFound, 1, "not found", nil)
		return
	}
	err = service.DeleteUser(id)
	if err != nil{
		appG.Response(http.StatusInternalServerError, 1, fmt.Sprint(err), nil)
		return
	}
	appG.Response(http.StatusOK, 0, "", nil)
}


// @Summary update users
// @Accept json
// @Produce json
// @Param id path int True "id"
// @Param body models.User True "body"
// @Success 200 {object} app.Response
// @Failure 400 {object} app.Response
// @Failure 404 {object} app.Response
// @Failure 500 {object} app.Response
func UpdateUser(c *gin.Context){
	var u models.UUser
	appG := app.Gin{C: c}

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil{
		appG.Response(http.StatusNotFound, 1, "not found", nil)
		return
	}
	err = c.Bind(&u)
	if err != nil{
		appG.Response(http.StatusBadRequest, 1, fmt.Sprint(err), nil)
		return
	}
	valid := validation.Validation{}
	ok, _ := valid.Valid(&u)
	if !ok{
		appG.Response(http.StatusBadRequest, 1, fmt.Sprint(valid.Errors[0]), nil)
		return
	}
	err = service.UpdateUser(id, u)
	if err != nil{
		appG.Response(http.StatusInternalServerError, 1, fmt.Sprint(err), nil)
		return
	}
	appG.Response(http.StatusOK, 0, "", nil)
}