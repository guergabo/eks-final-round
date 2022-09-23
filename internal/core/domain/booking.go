package domain

type Booking struct {
	Row                   RowID
	StartingRowIndex      int
	NumOfConsecutiveSeats int
}

func NewBooking(action *ActionConfig) *Booking {
	return &Booking{
		Row:                   action.Row,
		StartingRowIndex:      action.StartingRowIndex,
		NumOfConsecutiveSeats: action.NumOfConsecutiveSeats,
	}
}
