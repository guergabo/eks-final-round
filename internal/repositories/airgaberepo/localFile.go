// driven adapter
// An adapter for a driven port, transforms a technology agnostic request
// from the core into a specific technology request on the actor
package airgaberepo

import (
	"encoding/json"
	"errors"
	"os"
	"strconv"
	"strings"

	"github.com/guergabo/eks-final-round/internal/core/domain"
	"github.com/guergabo/eks-final-round/pkg/utils"
)

const (
	initial string = "init-state.json"
	current string = "current-state.json"
)

type localFile struct {
	initFileName    string
	currentFileName string
}

func NewLocalFile() *localFile {
	return &localFile{
		initFileName:    initial,
		currentFileName: current,
	}
}

// actual file manipulation
func (repo *localFile) Book(startingSeat string, numOfConsecutiveSeats int) error {
	airplane, err := loadState()
	if err != nil {
		return err
	}

	// check if row is valid - making row case insensitive
	row := strings.ToUpper(startingSeat[:1])
	if !airplane.IsValidRow(row) {
		return errors.New("could not accomodate customer request")
	}

	// check if seat numbers are valid and available
	desiredSeat, err := strconv.Atoi(startingSeat[1:])
	if err != nil {
		return err
	}

	if !airplane.AreValidSeats(desiredSeat, numOfConsecutiveSeats) {
		return errors.New("could not accomodate customer request")
	}

	if !airplane.AreSeatsAvailable(domain.RowID(row), desiredSeat, numOfConsecutiveSeats) {
		return errors.New("could not accomodate customer request")
	}

	airplane.ProcessBooking(domain.RowID(row), desiredSeat, numOfConsecutiveSeats)

	if err := storeState(current, airplane); err != nil {
		return err
	}
	return nil
}

// CAN YOU CANCEL IF YOU NEVER BOOKED IT?
func (repo *localFile) Cancel(startingSeat string, numOfConsecutiveSeats int) error {
	airplane, err := loadState()
	if err != nil {
		return err
	}

	// check if row is valid - case insensitive
	row := strings.ToUpper(startingSeat[:1])
	if !airplane.IsValidRow(row) {
		return errors.New("could not accomodate customer request")
	}

	// check if seat numbers are valid and available
	desiredSeat, err := strconv.Atoi(startingSeat[1:])
	if err != nil {
		return err
	}

	if !airplane.AreValidSeats(desiredSeat, numOfConsecutiveSeats) {
		return errors.New("could not accomodate customer request")
	}

	// check if seats had actually been booked
	if !airplane.AreSeatsBooked(domain.RowID(row), desiredSeat, numOfConsecutiveSeats) {
		return errors.New("could not accomodate customer request")
	}

	airplane.ProcessCancellation(domain.RowID(row), desiredSeat, numOfConsecutiveSeats)

	if err := storeState(current, airplane); err != nil {
		return errors.New("could not accomodate customer's request - unable to store state")
	}

	return nil
}

// private functions
func loadState() (*domain.Airplane, error) {
	jsonFilePath := initial
	if _, err := os.Stat(current); err == nil {
		jsonFilePath = current
	}

	byteValue, err := utils.ReadJSONFile(jsonFilePath)
	if err != nil {
		return nil, err
	}

	airplane := domain.NewAirplane()
	if err := json.Unmarshal(byteValue, airplane); err != nil {
		return nil, errors.New("internal server error - unmarshalling json")
	}

	return airplane, nil
}

func storeState(filename string, state *domain.Airplane) error {
	if err := utils.WriteJSONFile(filename, state); err != nil {
		return errors.New("internal server error - store json")
	}
	return nil
}
