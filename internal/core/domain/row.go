package domain

type Row struct {
	Seats []Seat `json:"seats"`
}

func NewRow() *Row {
	return &Row{
		Seats: make([]Seat, 8, 8),
	}
}
