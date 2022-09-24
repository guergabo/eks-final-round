// services
// entry point to the core and implements port
package airgabesrv

import (
	"errors"

	"github.com/guergabo/eks-final-round/internal/core/domain"
	"github.com/guergabo/eks-final-round/internal/core/dto"
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

// server side validation before doing any request and having to waste time loading up file
func (srv *service) Book(req *dto.Request) error {
	actionConfig, err := req.ValidBookingRequest()
	if err != nil {
		return errors.New("service error: " + err.Error())
	}

	// transformation of dto to domain object if all is good
	b := domain.NewBooking(actionConfig)

	if err := srv.airplaneRepository.Book(b); err != nil {
		return err
	}

	return nil
}

func (srv *service) Cancel(req *dto.Request) error {
	actionConfig, err := req.ValidCancellationRequest()
	if err != nil {
		return errors.New("service error: " + err.Error())
	}

	// transformation of dto to domain object if all is good
	c := domain.NewCancellation(actionConfig)

	if err := srv.airplaneRepository.Cancel(c); err != nil {
		return err
	}
	return nil
}
