package domain

type rowID string

const (
	A rowID = "A"
	B rowID = "B"
	C rowID = "C"
	D rowID = "D"
	E rowID = "E"
	F rowID = "F"
	G rowID = "G"
	H rowID = "H"
	I rowID = "I"
	J rowID = "J"
	K rowID = "K"
	L rowID = "L"
	M rowID = "M"
	N rowID = "N"
	O rowID = "O"
	P rowID = "P"
	Q rowID = "Q"
	R rowID = "R"
	S rowID = "S"
	T rowID = "T"
)

type Airplane struct {
	Rows map[rowID]Row
}

func NewAirplane() *Airplane {
	return &Airplane{
		Rows: make(map[rowID]Row, 20),
	}
}
