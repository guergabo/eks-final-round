package domain

type seatStatus string

const (
	Available seatStatus = "Available"
	Booked    seatStatus = "Booked"
)

type Seat struct {
	Status seatStatus `json:"status"`
}

func NewSeat() *Seat {
	return &Seat{
		Status: Available,
	}
}
