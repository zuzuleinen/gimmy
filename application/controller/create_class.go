package controller

import (
	"encoding/json"
	"errors"
	"gimmy/application/model"
	"gimmy/infrastructure"
	"net/http"
)

type CreateClassRequest struct {
	Name      string `json:"name"`
	StartDate string `json:"start_date"`
	EndDate   string `json:"end_date"`
	Capacity  int    `json:"capacity"`
}

type CreateClassResponse struct {
	ID string `json:"id"`
}

type Class struct {
	ID        string `json:"id"`
	Name      string `json:"name"`
	StartDate string `json:"start_date"`
	EndDate   string `json:"end_date"`
	Capacity  int    `json:"capacity"`
}

func CreateClass(w http.ResponseWriter, r *http.Request) error {
	decoder := json.NewDecoder(r.Body)
	var req CreateClassRequest
	err := decoder.Decode(&req)

	if err != nil {
		return infrastructure.JsonResponse(w, http.StatusBadRequest, "")
	}

	err = req.Validate()
	if err != nil {
		return infrastructure.JsonResponse(w, http.StatusBadRequest, err.Error())
	}

	id := createClass(req)

	return infrastructure.JSON(w, http.StatusCreated, CreateClassResponse{
		ID: id,
	})
}

func (r CreateClassRequest) Validate() error {
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

func createClass(r CreateClassRequest) string {
	id := infrastructure.GenerateID()
	c := model.NewClassRepository()

	c.Save(model.Class{
		ID:        id,
		Name:      r.Name,
		Capacity:  r.Capacity,
		StartDate: r.StartDate,
		EndDate:   r.EndDate,
	})

	return id
}
