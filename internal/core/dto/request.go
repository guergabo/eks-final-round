package dto

import (
	"errors"
	"strconv"
	"strings"
)

type RequestSubCommand string

const (
	Book   RequestSubCommand = "BOOK"
	Cancel RequestSubCommand = "CANCEL"
)

// just raw args
type Request struct {
	Action                string
	StartingSeat          string
	NumOfConsecutiveSeats string
}

func NewRequest(action string, startingSeating string, numOfConsecutiveSeats string) *Request {
	return &Request{
		Action:                action,
		StartingSeat:          startingSeating,
		NumOfConsecutiveSeats: numOfConsecutiveSeats,
	}
}

func (r *Request) ValidBookingRequest() error {
	// check if row is valid - making row case insensitive
	row := strings.ToUpper(r.StartingSeat[:1])
	if !IsValidRow(row) {
		return errors.New("row is not valid")
	}

	// check if seat numbers are valid and available
	startingSeatNum, err := strconv.Atoi(r.StartingSeat[1:])
	if err != nil {
		return errors.New("starting seat number is not valid")
	}

	// changing number
	numOfConsecutiveSeats, err := strconv.Atoi(r.NumOfConsecutiveSeats)
	if err != nil {
		return errors.New("could not accomodate number of seats")
	}

	if !AreValidSeats(startingSeatNum, numOfConsecutiveSeats) {
		return errors.New("seats are not valid")
	}

	return nil
}

func (r *Request) ValidCancellationRequest() error {
	// check if row is valid - case insensitive
	row := strings.ToUpper(r.StartingSeat[:1])
	if !IsValidRow(row) {
		return errors.New("could not accomodate customer request")
	}

	// check if seat numbers are valid and available
	startingSeatNum, err := strconv.Atoi(r.StartingSeat[1:])
	if err != nil {
		return err
	}

	numOfConsecutiveSeats, err := strconv.Atoi(r.NumOfConsecutiveSeats)
	if err != nil {
		return errors.New("could not accomodate customer request")
	}

	if !AreValidSeats(startingSeatNum, numOfConsecutiveSeats) {
		return errors.New("could not accomodate customer request")
	}

	return nil
}

// private methods
func IsValidRow(letter string) bool {
	bs := rune(letter[0])
	ascii := int(bs)
	if ascii >= 65 && ascii <= 84 { // A=65, T=84
		return true
	}
	return false
}

func AreValidSeats(start int, num int) bool {
	if (start < 0 || num > 7) || (start+num < 1 || start+num > 8) {
		return false
	}
	return true
}
