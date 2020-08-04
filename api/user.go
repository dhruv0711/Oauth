package api

import (
	"fmt"
	"strconv"

	"github.com/dhruv0711/Oauth/db"
	"github.com/gin-gonic/gin"
)

func User(c *gin.Context) {
	param := c.Param("id")
	id, err := strconv.Atoi(param)
	resp, err := db.User(id)
	if err != nil {
		fmt.Println(err)
		return
	}
	c.JSON(200, resp)
}
