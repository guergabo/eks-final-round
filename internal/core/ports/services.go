package ports

// take care of error handling
type AirplaneService interface {
	Book(startingSeat string, numOfConsecutiveSeats int) error
	Cancel(startingSeat string, numOfConsecutiveSeats int) error
}
