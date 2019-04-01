package router

import (
	"github.com/gin-gonic/gin"
)

func Location(router *gin.Engine) {
	// hello world
	router.GET("/get_order", GetOrder)
	//
	router.POST("/world", world)
}
