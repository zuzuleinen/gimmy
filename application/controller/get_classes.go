package controller

import (
	"gimmy/application/model"
	"gimmy/infrastructure"
	"net/http"
)

type GetClassesResponse struct {
	Classes []Class `json:"classes"`
}

func GetClasses(w http.ResponseWriter, r *http.Request) error {
	var resp GetClassesResponse
	resp.Classes = make([]Class, 0)

	c := model.NewClassRepository()

	for _, v := range c.All() {
		resp.Classes = append(resp.Classes, Class{
			ID:        v.ID,
			Name:      v.Name,
			StartDate: v.StartDate,
			EndDate:   v.EndDate,
			Capacity:  v.Capacity,
		})
	}

	return infrastructure.JSON(w, http.StatusOK, resp)
}
