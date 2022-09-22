package airgabehdl

type requestSubCommand string

const (
	book   requestSubCommand = "BOOK"
	cancel requestSubCommand = "CANCEL"
)

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
