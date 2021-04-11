package model

import (
	"encoding/json"
	"gimmy/application"
	"gimmy/infrastructure"
	"time"
)

type Class struct {
	ID        string `json:"id"`
	Name      string `json:"name"`
	Capacity  int    `json:"capacity"`
	StartDate string `json:"start_date"`
	EndDate   string `json:"end_date"`
}

func (c Class) FromTime() time.Time {
	date, _ := time.ParseInLocation(application.DateLayout, c.StartDate, time.UTC)
	return date
}

func (c Class) ToTime() time.Time {
	date, _ := time.ParseInLocation(application.DateLayout, c.EndDate, time.UTC)
	return date
}

func (c Class) DailyClasses() []DailyClass {
	var dcs []DailyClass

	for d := c.FromTime(); d.After(c.ToTime()) == false; d = d.AddDate(0, 0, 1) {
		dcs = append(dcs, DailyClass{
			Name:     c.Name,
			Capacity: c.Capacity,
			Date:     d,
		})
	}

	return dcs
}

type DailyClass struct {
	Name     string    `json:"name"`
	Capacity int       `json:"capacity"`
	Date     time.Time `json:"date"`
}

func (d DailyClass) Equals(date string) bool {
	return d.Date.Format(application.DateLayout) == date
}

func (c Class) Key() string {
	return c.ID
}

func (c Class) Data() []byte {
	jsn, _ := json.Marshal(&c)
	return jsn
}

type ClassesRepo interface {
	Save(c Class)
	All() []Class
	Find(id string) Class
}

type ClassRepository struct {
	db    *infrastructure.DBConn
	table string
}

func NewClassRepository() *ClassRepository {
	db := infrastructure.DB()
	return &ClassRepository{
		db,
		"classes",
	}
}

func (r ClassRepository) Save(c Class) {
	r.db.Save(r.table, c)
}

func (r ClassRepository) All() []Class {
	var cs []Class
	for _, v := range r.db.Tables[r.table] {
		var c Class
		err := json.Unmarshal(v.Data(), &c)
		if err != nil {
			panic(err)
		}
		cs = append(cs, c)
	}
	return cs
}

func (r ClassRepository) Find(id string) Class {
	for _, v := range r.All() {
		if v.ID == id {
			return v
		}
	}
	return Class{}
}
