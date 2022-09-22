package ports

type AirplaneRepository interface {
	Book(startingSeat string, numOfConsecutiveSeats int) error
	Cancel(startingSeat string, numOfConsecutiveSeats int) error
}
