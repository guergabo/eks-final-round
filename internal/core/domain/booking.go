package domain

type Booking struct {
	Row                   RowID
	StartingRowIndex      int
	NumOfConsecutiveSeats int
}

func NewBooking(row RowID, start int, num int) *Booking {
	return &Booking{
		Row:                   row,
		StartingRowIndex:      start,
		NumOfConsecutiveSeats: num,
	}
}
