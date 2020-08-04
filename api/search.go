package api

import (
	"fmt"

	"github.com/dhruv0711/Oauth/db"
	"github.com/gin-gonic/gin"
)

func Search(c *gin.Context) {
	email := c.Param("search")
	resp, err := db.Search(email)
	if err != nil {
		fmt.Println(err)
		return
	}
	c.JSON(200, resp)
}
