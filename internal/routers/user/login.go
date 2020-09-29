package user

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

/**
* @Author: super
* @Date: 2020-09-29 10:40
* @Description:
**/

type User struct {

}

func NewUser() User{
	return User{}
}

func (u User) Login(c *gin.Context) {
	username := c.PostForm("username")
	fmt.Println(username)
	password := c.PostForm("password")
	fmt.Println(password)

	result := username + password

	c.JSON(http.StatusOK, result)
	return
}