package router

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func world(c *gin.Context) {
	c.Request.ParseForm()
	for k, v := range c.Request.PostForm {
		fmt.Println("key: ", k, ", values: ", v)
	}
	c.JSON(http.StatusOK, gin.H{
		"status":     1,
		"error_code": 0,
		"data":       "world!",
	})
}
