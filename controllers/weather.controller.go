package controllers

import (
	"assignment_3/models"
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
)

var result models.Weather

func Get(c *gin.Context) {
	result = models.GetStatus()

	go Looper()

	c.JSON(200, gin.H{
		"message": result,
	})
}

func Looper() {

	ch := time.Tick(15 * time.Second)
	for next := range ch {
		fmt.Println(next)
		result = models.GetStatus()

		fmt.Println(result)
	}
}
