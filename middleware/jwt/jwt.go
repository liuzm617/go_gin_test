package jwt

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"go_gin_example/utils/common"
	"net/http"
	"time"
)

// iwt middleware
func JWT() gin.HandlerFunc{
	return func(c *gin.Context){
		var (
			code int
			msg string
			httpCode int
			data interface{}
		)

		token := c.Request.Header.Get("token")
		if token == ""{
			code = 1
			msg = "Unauthorized"
			httpCode = http.StatusUnauthorized
		}else{
			_, err := ParseToken(token)
			if err != nil{
				switch err.(*jwt.ValidationError).Errors {
				case jwt.ValidationErrorExpired:
					code = 1
					msg = "token exired"
					httpCode = http.StatusUnauthorized
				default:
					code = 1
					msg = "Invalidate token"
					httpCode = http.StatusUnauthorized
				}
			}
		}
		if code != 0{
			c.JSON(httpCode, gin.H{
				"code": code,
				"msg": msg,
				"data":data,
			})
			c.Abort()
			return
		}
		c.Next()

	}
}

var jwtSecret []byte

type Jwt struct {
	Username string `json:"username"`
	Password string `json:"password"`
	jwt.StandardClaims
}

// generate token
func GenerateToken(username, password string)(string, error){
	now := time.Now()
	expireTime := now.Add(3*time.Hour)

	claims := Jwt{
		Username:common.EncodeMD5(username),
		Password:common.EncodeMD5(password),
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expireTime.Unix(),
			Issuer: "gin",
		},
	}

	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err := tokenClaims.SignedString(jwtSecret)
	return token, err
}

// parse token
func ParseToken(token string)(*Jwt, error){
	tokenClaims, err := jwt.ParseWithClaims(token, &Jwt{}, func(token *jwt.Token) (i interface{}, e error) {
		return jwtSecret, nil
	})

	if tokenClaims != nil{
		if claims, ok := tokenClaims.Claims.(*Jwt); ok && tokenClaims.Valid {
			return claims, nil
		}
	}
	return nil, err
}