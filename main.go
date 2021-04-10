package main

import (
	"gimmy/application"
	"gimmy/infrastructure"
	"net/http"
)

func main() {
	a := infrastructure.NewApplication(":8282")

	a.RegisterRoute(http.MethodPost, "/classes", application.CreateClass)

	a.RegisterRoute(http.MethodGet, "/classes", application.GetClass)
	//a.RegisterRoute(http.MethodPost, "/bookings", application.CreateBooking)

	a.Run()
}
