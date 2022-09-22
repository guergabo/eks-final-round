// driven adapter
// An adapter for a driven port, transforms a technology agnostic request
// from the core into a specific technology request on the actor
package airgaberepo

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"os"
	"strconv"

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
	log.Println("processing booking request")
	airplane, err := loadState()
	if err != nil {
		return err
	}

	// NEEDS ERROR HANDLING - this should happen in handler !!!!
	ri := domain.RowID(startingSeat[:1])
	desiredRow := airplane.Rows[ri]

	start, err := strconv.Atoi(startingSeat[1:])
	if err != nil {
		return err
	}

	// check for availability
	seats := desiredRow.Seats[start:(start + numOfConsecutiveSeats)]
	fmt.Println(seats)
	for i := range seats {
		fmt.Println(seats[i].Status)
		if seats[i].Status != domain.Available {
			return errors.New("could not accomodate customer requests")
		}
	}

	// update availability
	for i := range seats {
		seats[i].Status = domain.Booked
	}

	// store new state
	if err := storeState(current, airplane); err != nil {
		return err
	}
	return nil
}

// CAN YOU CANCEL IF YOU NEVER BOOKED IT?
func (repo *localFile) Cancel(startingSeat string, numOfConsecutiveSeats int) error {
	airplane, err := loadState()
	if err != nil {
		return errors.New("could not accomodate customer's request - unable to load state file")
	}

	// NEEDS ERROR HANDLING - this should happen in handler !!!!
	ri := domain.RowID(startingSeat[:1])
	desiredRow := airplane.Rows[ri]

	start, err := strconv.Atoi(startingSeat[1:])
	if err != nil {
		return errors.New("could not accomodate customer's request - unable to get seat value")
	}

	// check for availability
	seats := desiredRow.Seats[start:(start + numOfConsecutiveSeats)]
	for i := range seats {
		// get error if you are canceling something that is not booked
		if seats[i].Status == domain.Available {
			return errors.New("could not accomodate customer's request - seats not available")
		}
	}

	// update availability
	for i := range seats {
		seats[i].Status = domain.Available
	}

	// store new state
	if err := storeState(current, airplane); err != nil {
		return errors.New("could not accomodate customer's request - unable to store state")
	}

	return nil
}

// helpers
func loadState() (*domain.Airplane, error) {
	jsonFilePath := initial
	if _, err := os.Stat(current); err == nil {
		jsonFilePath = current
	}

	byteValue, err := utils.ReadJSONFile(jsonFilePath)
	if err != nil {
		return nil, err
		// return nil, errors.New("internal server error - reading json")
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
