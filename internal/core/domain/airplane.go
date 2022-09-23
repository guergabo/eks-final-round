package domain

type RowID string

const (
	A RowID = "A"
	B RowID = "B"
	C RowID = "C"
	D RowID = "D"
	E RowID = "E"
	F RowID = "F"
	G RowID = "G"
	H RowID = "H"
	I RowID = "I"
	J RowID = "J"
	K RowID = "K"
	L RowID = "L"
	M RowID = "M"
	N RowID = "N"
	O RowID = "O"
	P RowID = "P"
	Q RowID = "Q"
	R RowID = "R"
	S RowID = "S"
	T RowID = "T"
)

type Airplane struct {
	LastUpdate      string        `json:"last_updated"`
	Airline         string        `json:"airline"`
	FlightIdent     string        `json:"flight_ident"`
	Aircraft        string        `json:"aircraft"`
	DepartureRegion string        `json:"departure_region"`
	ArrivalRegion   string        `json:"arrival_region"`
	Status          string        `json:"status"`
	Rows            map[RowID]Row `json:"rows"`
}

func NewAirplane() *Airplane {
	return &Airplane{}
}

// need to check file for this stuff
func (a *Airplane) AreSeatsAvailable(row RowID, start int, num int) bool {
	desiredRow := a.Rows[row]
	seats := desiredRow.Seats[start:(start + num)]
	for i := range seats {
		if seats[i].Status != Available {
			return false
		}
	}
	return true
}

func (a *Airplane) AreSeatsBooked(row RowID, start int, num int) bool {
	desiredRow := a.Rows[row]
	seats := desiredRow.Seats[start:(start + num)]
	for i := range seats {
		// get error if you are canceling something that is not booked
		if seats[i].Status == Available {
			return false
		}
	}
	return true
}

func (a *Airplane) ProcessBooking(row RowID, start int, num int) {
	desiredRow := a.Rows[row]
	seats := desiredRow.Seats[start:(start + num)]
	for i := range seats {
		seats[i].Status = Booked
	}
}

func (a *Airplane) ProcessCancellation(row RowID, start int, num int) {
	desiredRow := a.Rows[row]
	seats := desiredRow.Seats[start:(start + num)]
	for i := range seats {
		seats[i].Status = Available
	}
}
