package model

import (
	"encoding/json"
	"errors"
	"fmt"
	"gimmy/infrastructure"
)

type BookingRepo interface {
	Save(b Booking) string
}

type Booking struct {
	ID     string `json:"id"`
	Member string `json:"member"`
	Date   string `json:"date"`
}

func (b Booking) Key() string {
	return b.ID
}

func (b Booking) Data() []byte {
	jsn, _ := json.Marshal(&b)
	return jsn
}

type BookingCreator struct {
	classes  ClassesRepo
	bookings BookingRepo
}

func NewBookingCreator(cr ClassesRepo, br BookingRepo) *BookingCreator {
	bc := BookingCreator{classes: cr, bookings: br}
	return &bc
}

func (bc *BookingCreator) Book(member, date, classID string) (string, error) {
	if len(bc.classes.All()) == 0 {
		return "", errors.New("no classes exists")
	}

	c := bc.classes.Find(classID)
	if c.ID == "" {
		return "", fmt.Errorf("class not found for id `%s`", classID)
	}

	for _, v := range c.DailyClasses() {
		if v.Equals(date) {
			bookingID := bc.bookings.Save(Booking{
				ID:     infrastructure.GenerateID(),
				Member: member,
				Date:   date,
			})
			return bookingID, nil
		}
	}

	return "", fmt.Errorf(
		"cannot book `%s` for date `%s`. can only book between %s - %s",
		c.Name,
		date,
		c.StartDate,
		c.EndDate,
	)
}

type BookingRepository struct {
	table string
}

func NewBookingRepository() *BookingRepository {
	return &BookingRepository{
		"bookings",
	}
}

func (r BookingRepository) Save(b Booking) string {
	db := infrastructure.DB()
	id := infrastructure.GenerateID()

	db.Save(r.table, b)

	return id
}
