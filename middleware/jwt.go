package middleware

import (
	"OpenMall/util/string_util"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"net/http"
)

// MyCustomClaims jwt结构体
type MyCustomClaims struct {
	Id       int    `json:"id"`
	Username string `json:"username"`
	jwt.StandardClaims
}

// 定义加密秘钥 后面放在配置文件
var mySigningKey = []byte("qwer1234")

// AuthJwt 验证jwt
func AuthJwt(signToken string) (*MyCustomClaims, error) {
	var claims MyCustomClaims
	token, err := jwt.ParseWithClaims(signToken, &claims, func(token *jwt.Token) (interface{}, error) {
		return mySigningKey, nil
	})
	if token.Valid {
		return &claims, nil
	} else {
		return nil, err
	}
}

// GenJwt 生成jwt
func GenJwt(claims MyCustomClaims) (string, error) {
	// 使用HS256加密方式
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signToken, err := token.SignedString(mySigningKey)
	if !string_util.IsNil(err) {
		return "", err
	}
	return signToken, nil
}

// JWTAuthMiddleware 基于JWT的认证中间件
func JWTAuthMiddleware(c *gin.Context) {
	// 从请求头中取出
	signToken := c.Request.Header.Get("Authorization")
	if signToken == "" {
		c.JSON(http.StatusOK, gin.H{
			"code": 1002,
			"msg":  "token为空",
		})
		c.Abort()
		return
	}
	// 校验token
	myclaims, err := AuthJwt(signToken)
	if !string_util.IsNil(err) {
		fmt.Println(err)
		c.JSON(http.StatusOK, gin.H{
			"code": 1003,
			"msg":  "token校验失败",
		})
		c.Abort()
		return
	}
	// 将用户的id放在到请求的上下文c上
	c.Set("userid", myclaims.Id)
	c.Next() // 后续的处理函数可以用过c.Get("userid")来获取当前请求的id

}
