package routers

import(
	"github.com/gin-gonic/gin"
	"login-demo/internal/routers/user"
)

/**
* @Author: super
* @Date: 2020-09-29 10:32
* @Description:
**/

func NewRouter() *gin.Engine {
	g := gin.New()
	g.Static("/", "/Users/super/develop/login-demo/assets")

	u := user.NewUser()
	g.POST("/login", u.Login)

	return g
}