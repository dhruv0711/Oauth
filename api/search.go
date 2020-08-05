package api

import (
	"fmt"

	"github.com/dhruv0711/Oauth/db"
	"github.com/gin-gonic/gin"
)

func Search(c *gin.Context) {
	search := c.Request.URL.Query().Get("value")
	if search == "" {
		return
	}
	resp, err := db.Search(search)
	if err != nil {
		fmt.Println(err)
		return
	}
	c.JSON(200, resp)
}
