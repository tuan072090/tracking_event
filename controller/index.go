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

	Validate = validator.New()

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	//	add user event
	r.POST("/user-event", AddUserEvent)

	var port = ":3000"
	fmt.Println("port env là...", os.Getenv("PORT"))
	if os.Getenv("PORT") != "" {
		port = ":" + os.Getenv("PORT")
	}

	err := r.Run(port)

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
