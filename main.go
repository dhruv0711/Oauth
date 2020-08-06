package main

import (
	"github.com/dhruv0711/Oauth/api"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.StaticFile("/images/github.png", "./html/images/github.png")
	r.StaticFile("/images/linkedin.png", "./html/images/linkedin.png")
	r.LoadHTMLGlob("html/login.html")
	r.GET("/login", api.Login)
	r.GET("/githubhome", api.Githubhome)
	r.GET("/linkedinhome", api.Linkedinhome)
	r.GET("/user/:id", api.User)
	r.GET("/users/all", api.Allusers)
	r.GET("/users/search", api.Search)

	r.Run(":8084")
}
