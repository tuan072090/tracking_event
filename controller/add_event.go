package controller

import (
	"github.com/gin-gonic/gin"
	"tracking_event/model"
	"tracking_event/service"
)

// AddUserEvent gets the payload from user client and save direct to bigquery
func AddUserEvent (c *gin.Context){
	var event model.Event

	//	binding
	errBinding := c.ShouldBindJSON(&event)
	if errBinding != nil {
		c.JSON(400, gin.H{
			"message": "Dữ liệu không hợp lệ",
		})
		return
	}

	//	validate
	errValid := Validate.Struct(event)
	if errValid != nil {
		c.JSON(400, gin.H{
			"message": errValid.Error(),
		})
		return
	}

	//	save
	err := service.InsertEvent(event)
	if err != nil {
		c.JSON(422, gin.H{
			"message": err.Error(),
		})
		return
	}

	c.JSON(200, gin.H{
		"message": "Thành công",
	})
}
