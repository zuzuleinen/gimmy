package controller

import (
	"encoding/json"
	"errors"
	"gimmy/application/model"
	"gimmy/infrastructure"
	"net/http"
)

type CreateBookingRequest struct {
	Name    string `json:"name"`
	Date    string `json:"date"`
	ClassID string `json:"class_id"`
}

type CreateBookingResponse struct {
	ID string `json:"id"`
}

func CreateBooking(w http.ResponseWriter, r *http.Request) error {
	decoder := json.NewDecoder(r.Body)
	var req CreateBookingRequest
	err := decoder.Decode(&req)

	if err != nil {
		return infrastructure.JsonResponse(w, http.StatusBadRequest, "")
	}

	err = req.Validate()
	if err != nil {
		return infrastructure.JsonResponse(w, http.StatusBadRequest, err.Error())
	}

	bc := model.NewBookingCreator(model.NewClassRepository(), model.NewBookingRepository())

	bookingID, err := bc.Book(req.Name, req.Date, req.ClassID)

	if err != nil {
		return infrastructure.JsonResponse(w, http.StatusBadRequest, err.Error())
	}

	return infrastructure.JSON(w, http.StatusCreated, CreateBookingResponse{
		ID: bookingID,
	})
}

func (r CreateBookingRequest) Validate() error {
	if r.Name == "" {
		return errors.New("`name` is required")
	}
	if r.Date == "" {
		return errors.New("`date` is required")
	}
	if r.ClassID == "" {
		return errors.New("`class_id` is required")
	}
	return nil
}
