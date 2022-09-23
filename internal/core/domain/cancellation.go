package domain

type Cancellation struct {
	Row                   RowID
	StartingRowIndex      int
	NumOfConsecutiveSeats int
}

func NewCancellation(action *ActionConfig) *Cancellation {
	return &Cancellation{
		Row:                   action.Row,
		StartingRowIndex:      action.StartingRowIndex,
		NumOfConsecutiveSeats: action.NumOfConsecutiveSeats,
	}
}
