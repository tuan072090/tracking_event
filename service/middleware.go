package service

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

// EventAuthen function authenticates user header
func EventAuthen(c *gin.Context){

	referral := c.Request.Header.Get("referer")
	userAgent := c.Request.Header.Get("user-agent")

	test := c.GetHeader("Referer")
	test2 := c.GetHeader("user-agent")

	fmt.Println("referral....", referral)
	fmt.Println("userAgent...", userAgent)

	fmt.Println("test...", test)
	fmt.Println("test2...", test2)
	c.Next()
}
