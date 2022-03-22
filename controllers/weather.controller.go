package controllers

import (
	"assignment_3/models"
	"fmt"
	"net/http"
	"net/url"
	"time"

	"github.com/gin-gonic/gin"
)

var result models.Weather

func Get(c *gin.Context) {
	result = models.GetStatus()

	go Looper(c)
	fmt.Println("First => ", result)
	c.JSON(200, gin.H{
		"message": result,
	})
}

func Looper(c *gin.Context) {

	ch := time.Tick(15 * time.Second)
	for next := range ch {
		result = models.GetStatus()

		fmt.Println("NEXT -> ", next)
		fmt.Println("HASIL NEXT => ", result)
		location := url.URL{Path: "/"}
		c.Redirect(http.StatusFound, location.RequestURI())
	}
}
