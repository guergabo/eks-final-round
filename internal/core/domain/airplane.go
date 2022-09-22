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
	LastUpdate      string        `json:"last_updated"`
	Airline         string        `json:"airline"`
	FlightIdent     string        `json:"flight_ident"`
	Aircraft        string        `json:"aircraft"`
	DepartureRegion string        `json:"departure_region"`
	ArrivalRegion   string        `json:"arrival_region"`
	Status          string        `json:"status"`
	Rows            map[rowID]Row `json:"rows"`
}

func NewAirplane() *Airplane {
	return &Airplane{}
}
