package v1_test

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"go_gin_example/models"
	"go_gin_example/routers"
	"go_gin_example/utils/common"
	"go_gin_example/utils/conf"
	"go_gin_example/utils/redis"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

var router *gin.Engine

func init(){
	conf.Init()

	// 初始化mysql连接
	models.Init()
	//defer models.Close()

	//初始化redis
	redis.Init()

	router = routers.Router()
}

func TestAddUser(t *testing.T) {

	data := map[string]string{
		"username": common.GetRandomString(5),
		"password": "123",
	}

	url := "/api/v1/users"
	w, err := POST(url, data)
	assert.Equal(t, err, nil)
	assert.Equal(t, w.Code, 200)
	fmt.Println("response body", w.Body.String())

}

func TestGetUsers(t *testing.T) {
	url := "/api/v1/users"
	w, err := GET(url)
	assert.Equal(t, err, nil)
	assert.Equal(t, w.Code, 200)
	fmt.Println("get users response body:", w.Body.String())
}

func POST(url string, data interface{}) (*httptest.ResponseRecorder ,error){
	w := httptest.NewRecorder()
	jsonData, err := common.MapJson(data)
	fmt.Printf("post json:%s\n", jsonData)
	if err != nil{
		return nil, err
	}
	req, err := http.NewRequest("POST", url, strings.NewReader(jsonData))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json")
	router.ServeHTTP(w, req)
	return w, nil
}

func GET(url string) (*httptest.ResponseRecorder ,error){
	w := httptest.NewRecorder()
	req, err := http.NewRequest("GET",  url, nil)
	if err != nil{
		return nil, err
	}
	router.ServeHTTP(w, req)
	return w, nil
}