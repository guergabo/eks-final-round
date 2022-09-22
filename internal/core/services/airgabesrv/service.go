// services
// entry point to the core and implements port
package airgabesrv

import (
	"github.com/guergabo/eks-final-round/internal/core/ports"
)

type service struct {
	airplaneRepository ports.AirplaneRepository
}

func New(airplaneRepository ports.AirplaneRepository) *service {
	return &service{
		airplaneRepository: airplaneRepository,
	}
}

func (srv *service) Book(startingSeat string, numOfConsecutiveSeats int) error {
	if err := srv.airplaneRepository.Book(startingSeat, numOfConsecutiveSeats); err != nil {
		return err
	}
	return nil
}

func (srv *service) Cancel(startingSeat string, numOfConsecutiveSeats int) error {
	if err := srv.airplaneRepository.Cancel(startingSeat, numOfConsecutiveSeats); err != nil {
		return err
	}
	return nil
}
