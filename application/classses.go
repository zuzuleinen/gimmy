package application

import (
	"encoding/json"
	"errors"
	"fmt"
	"gimmy/api"
	"gimmy/infrastructure"
	"net/http"
	"time"
)

type Class struct {
	ID       string
	Name     string
	Capacity int
	From     time.Time
	To       time.Time
}

func GetClass(w http.ResponseWriter, r *http.Request) error {
	fmt.Println("classes GET")
	return infrastructure.JsonResponse(w, http.StatusOK, "")
}

func CreateClass(w http.ResponseWriter, r *http.Request) error {
	decoder := json.NewDecoder(r.Body)
	var c api.CreateClassRequest
	err := decoder.Decode(&c)

	if err != nil {
		return infrastructure.JsonResponse(w, http.StatusBadRequest, "")
	}

	err = validate(c)
	if err != nil {
		return infrastructure.JsonResponse(w, http.StatusBadRequest, err.Error())
	}

	err = createClass(c)
	return infrastructure.JsonResponse(w, http.StatusCreated, "")
}

func validate(r api.CreateClassRequest) error {
	if r.Name == "" {
		return errors.New("`name` is required")
	}
	if r.StartDate == "" {
		return errors.New("`start_date` is required")
	}
	if r.EndDate == "" {
		return errors.New("`end_date` is required")
	}
	if r.Capacity <= 0 {
		return errors.New("`capacity` should be bigger than zero")
	}
	return nil
}

func createClass(c api.CreateClassRequest) error {
	fmt.Println("Creating class", c.Name, c.StartDate, c.EndDate)
	return nil
}
