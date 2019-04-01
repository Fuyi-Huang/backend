package main

import (
	"github.com/fuyi-huang/backend/config"
	"github.com/fuyi-huang/backend/library"
	"github.com/fuyi-huang/backend/router"
	"github.com/gin-gonic/gin"
)

func initial() {
	config.ReadConfig()
	library.ConnectMysql()
}

func main() {
	initial()

	r := gin.Default()
	router.Location(r)
	r.Run(":8181")
}
