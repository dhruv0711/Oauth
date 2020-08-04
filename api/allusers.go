package api

import (
	"fmt"

	"github.com/dhruv0711/Oauth/db"
	"github.com/gin-gonic/gin"
)

func Allusers(c *gin.Context) {
	resp, err := db.Allusers()
	if err != nil {
		fmt.Println(err)
		return
	}
	c.JSON(200, resp)
}
