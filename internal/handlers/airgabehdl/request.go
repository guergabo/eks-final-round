package airgabehdl

type Request struct {
	Action                string
	StartingSeat          string
	NumOfConsecutiveSeats int
}

func NewRequest(action string, startingSeating string, numOfConsecutiveSeats int) *Request {
	return &Request{
		Action:                action,
		StartingSeat:          startingSeating,
		NumOfConsecutiveSeats: numOfConsecutiveSeats,
	}
}
