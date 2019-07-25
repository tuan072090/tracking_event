package controller

import (
	"fmt"
	"os"
	"tracking_event/service"

	"github.com/gin-gonic/gin"
	"gopkg.in/go-playground/validator.v9"
)

var Validate *validator.Validate

// InitGin inits the routers of gin
func InitGin() {
	r := gin.Default()

	gin.SetMode("release")

	Validate = validator.New()

	r.GET("/ping", ping)
	//	add user event
	r.POST("/user-event", service.EventAuthen, AddUserEvent)
	r.GET("events", QueryEvent)

	fmt.Println("Listen on port " + service.GetPort())
	err := r.Run(":" + service.GetPort())
	if err != nil {
		os.Exit(1)
	}
}

// CreateTableEvent creates the table of event
func CreateTableEvent() {
	err := service.CreateNewTable()

	if err != nil {
		fmt.Println("Lỗi tạo table cmnr")
		os.Exit(1)
		return
	}
	fmt.Println("ok")
}

func ping(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "pong",
	})
}
