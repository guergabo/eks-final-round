package domain

type ActionConfig struct {
	Row                   RowID
	StartingRowIndex      int
	NumOfConsecutiveSeats int
}

func NewActionConfig(row RowID, start int, num int) *ActionConfig {
	return &ActionConfig{
		Row:                   row,
		StartingRowIndex:      start,
		NumOfConsecutiveSeats: num,
	}
}
