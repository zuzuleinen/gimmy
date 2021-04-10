package api

const DateLayout = "2006-01-02"

type CreateClassRequest struct {
	Name      string `json:"name"`
	StartDate string `json:"start_date"`
	EndDate   string `json:"end_date"`
	Capacity  int    `json:"capacity"`
}

type CreateBookingRequest struct {
	Name string `json:"name"`
	Date string `json:"date"`
}
