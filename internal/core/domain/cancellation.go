package domain

type Cancellation struct {
	Row                   RowID
	StartingRowIndex      int
	NumOfConsecutiveSeats int
}

func NewCancellation(row RowID, start int, num int) *Cancellation {
	return &Cancellation{
		Row:                   row,
		StartingRowIndex:      start,
		NumOfConsecutiveSeats: num,
	}
}
