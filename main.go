package main

import (
	"gimmy/application/controller"
	"gimmy/infrastructure"
)

func main() {
	a := infrastructure.NewApplication(":8282")

	a.POST("/classes", controller.CreateClass)
	a.GET("/classes", controller.GetClasses)

	a.POST("/bookings", controller.CreateBooking)

	a.Run()
}
