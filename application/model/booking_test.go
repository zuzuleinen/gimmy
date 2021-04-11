package model

import (
	"testing"
)

func TestBookWithNoClasses(t *testing.T) {
	bc := NewBookingCreator(NoClassesRepo{}, BookingRepoStub{})
	_, err := bc.Book("John", "2021-01-25", "someId")

	expectedErrorMsg := "no classes exists"
	if err != nil && err.Error() != expectedErrorMsg {
		t.Errorf("invalid error. received `%s`, expected `%s`", err, expectedErrorMsg)
	}
	if err == nil {
		t.Errorf("should return an error if no classes exists")
	}
}

func TestBookWithInvalidClassID(t *testing.T) {
	bc := NewBookingCreator(NewRepoWithClasses(), BookingRepoStub{})
	_, err := bc.Book("John", "2021-01-25", "someId")

	expectedErrorMsg := "class not found for id `someId`"
	if err != nil && err.Error() != expectedErrorMsg {
		t.Errorf("invalid error. received `%s`, expected `%s`", err, expectedErrorMsg)
	}
	if err == nil {
		t.Errorf("should return an error for class not found")
	}
}

func TestBookWithInvalidDate(t *testing.T) {
	bc := NewBookingCreator(NewRepoWithClasses(), BookingRepoStub{})
	_, err := bc.Book("John", "2021-01-25", "111")

	expectedErrorMsg := "cannot book `Yoga` for date `2021-01-25`. can only book between 2021-04-01 - 2021-04-30"
	if err != nil && err.Error() != expectedErrorMsg {
		t.Errorf("invalid error. received `%s`, expected `%s`", err, expectedErrorMsg)
	}
	if err == nil {
		t.Errorf("should return an error for invalid date")
	}
}

func TestSuccessfulBooking(t *testing.T) {
	bc := NewBookingCreator(NewRepoWithClasses(), BookingRepoStub{})
	bookingID, err := bc.Book("John", "2021-04-02", "111")

	if err != nil {
		t.Errorf("eror should be nil on successful booking")
	}

	if bookingID == "" {
		t.Errorf("booking ID should not be empty on successful booking")
	}
}

type NoClassesRepo struct {
}

func (r NoClassesRepo) All() []Class {
	return []Class{}
}
func (r NoClassesRepo) Find(id string) Class {
	return Class{}
}
func (r NoClassesRepo) Save(c Class) {
}

type RepoWithClasses struct {
	classes []Class
}

func NewRepoWithClasses() *RepoWithClasses {
	r := RepoWithClasses{
		classes: []Class{{
			ID:        "111",
			Name:      "Yoga",
			Capacity:  10,
			StartDate: "2021-04-01",
			EndDate:   "2021-04-30",
		}},
	}
	return &r
}

func (r RepoWithClasses) All() []Class {
	return r.classes
}
func (r RepoWithClasses) Find(id string) Class {
	for _, v := range r.classes {
		if v.ID == id {
			return v
		}
	}
	return Class{}
}
func (r RepoWithClasses) Save(c Class) {
}

type BookingRepoStub struct {
}

func (r BookingRepoStub) Save(b Booking) string {
	return b.ID
}
