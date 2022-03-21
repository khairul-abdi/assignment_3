package main

import (
	"assignment_3/controllers"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/", controllers.Get)
	r.Run()
}
