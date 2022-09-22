// driven adapter
// An adapter for a driven port, transforms a technology agnostic request
// from the core into a specific technology request on the actor
package airgaberepo

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
