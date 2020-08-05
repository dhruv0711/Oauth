package main

import (
	"github.com/dhruv0711/Oauth/api"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.LoadHTMLGlob("templates/*")
	r.GET("/login", api.Login)
	r.GET("/githubhome", api.Githubhome)
	r.GET("/linkedinhome", api.Linkedinhome)
	r.GET("/user/:id", api.User)
	r.GET("/users/all", api.Allusers)
	r.GET("/users/search", api.Search)

	r.Run(":8084")
}
