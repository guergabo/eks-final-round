// services
// entry point to the core and implements port
package airgabesrv

import (
	"log"

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
	log.Println("calling booking repo service")
	if err := srv.airplaneRepository.Book(startingSeat, numOfConsecutiveSeats); err != nil {
		log.Println(err)
		return err
	}
	return nil
}

func (srv *service) Cancel(startingSeat string, numOfConsecutiveSeats int) error {
	log.Println("calling canceling repo service")
	if err := srv.airplaneRepository.Cancel(startingSeat, numOfConsecutiveSeats); err != nil {
		return err
	}
	return nil
}
