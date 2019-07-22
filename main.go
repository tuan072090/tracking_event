package main

import (
	"fmt"
	"tracking_event/controller"
)

func main(){
	fmt.Println("Starting gin")

	//controller.CreateTableEvent()
	controller.InitGin()
}
