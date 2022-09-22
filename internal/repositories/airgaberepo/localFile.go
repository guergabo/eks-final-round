// driven adapter
// An adapter for a driven port, transforms a technology agnostic request
// from the core into a specific technology request on the actor
package airgaberepo

import (
	"encoding/json"
	"os"

	"github.com/guergabo/eks-final-round/internal/core/domain"
	"github.com/guergabo/eks-final-round/pkg/utils"
)

const (
	initial string = "init-state.json"
	current string = "current-state.json"
)

type localFile struct {
	fileName string
}

func NewLocalFile() *localFile {
	return &localFile{
		fileName: "init-state.json",
	}
}

// actual file manipulation
func (repo *localFile) Book(startingSeat string, numOfConsecutiveSeats int) error {
	return nil
}

func (repo *localFile) Cancel(startingSeat string, numOfConsecutiveSeats int) error {
	return nil
}

// add logic
// if current-state.json exists load that
// if not load init-state.json it is first run
func loadFile() *domain.Airplane {
	jsonFilePath := initial
	if _, err := os.Stat("current-state.json"); err == nil {
		// file exists exists
		jsonFilePath = current
	}
	byteValue := utils.MustRead(utils.ReadJSONFile(jsonFilePath))
	airplane := domain.NewAirplane()
	json.Unmarshal(byteValue, airplane)
	return airplane
}
