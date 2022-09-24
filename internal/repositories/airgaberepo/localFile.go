// driven adapter
// An adapter for a driven port, transforms a technology agnostic request
// from the core into a specific technology request on the actor
package airgaberepo

import (
	"encoding/json"
	"errors"
	"os"

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

func (repo *localFile) Book(book *domain.Booking) error {
	airplane, err := loadState()
	if err != nil {
		return errors.New("repo error: " + err.Error())
	}

	if !airplane.AreSeatsAvailable(book.Row, book.StartingRowIndex, book.NumOfConsecutiveSeats) {
		return errors.New("repo error: seats are not available")
	}

	airplane.ProcessBooking(book.Row, book.StartingRowIndex, book.NumOfConsecutiveSeats)

	if err := storeState(current, airplane); err != nil {
		return errors.New("repo error: " + err.Error())
	}
	return nil
}

func (repo *localFile) Cancel(cancel *domain.Cancellation) error {
	airplane, err := loadState()
	if err != nil {
		return errors.New("repo error: " + err.Error())
	}

	if !airplane.AreSeatsBooked(cancel.Row, cancel.StartingRowIndex, cancel.NumOfConsecutiveSeats) {
		return errors.New("repo error: seats were already available")
	}

	airplane.ProcessCancellation(cancel.Row, cancel.StartingRowIndex, cancel.NumOfConsecutiveSeats)

	if err := storeState(current, airplane); err != nil {
		return errors.New("repo error: " + err.Error())
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
		return nil, errors.New("repo error: " + err.Error())
	}

	return airplane, nil
}

func storeState(filename string, state *domain.Airplane) error {
	if err := utils.WriteJSONFile(filename, state); err != nil {
		return errors.New("repo error: " + err.Error())
	}
	return nil
}
