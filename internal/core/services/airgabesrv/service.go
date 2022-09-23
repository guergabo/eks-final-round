// services
// entry point to the core and implements port
package airgabesrv

import (
	"errors"
	"strconv"
	"strings"

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
	// simplify to validate
	if err := req.ValidBookingRequest(); err != nil {
		return errors.New("could not accomodate customer request: " + err.Error())
	}

	// transform fields to make domain object
	row := domain.RowID(strings.ToUpper(req.StartingSeat[:1]))
	startingSeatNum, err := strconv.Atoi(req.StartingSeat[1:])
	if err != nil {
		return err
	}
	numOfConsecutiveSeats, err := strconv.Atoi(req.NumOfConsecutiveSeats)
	if err != nil {
		return err
	}

	// transformation of dto to domain object if all is good
	b := domain.NewBooking(row, startingSeatNum, numOfConsecutiveSeats)

	if err := srv.airplaneRepository.Book(b); err != nil {
		return err
	}

	return nil
}

func (srv *service) Cancel(req *dto.Request) error {
	// simplify validaton
	if err := req.ValidCancellationRequest(); err != nil {
		return errors.New("could not accomodate customer request")
	}

	// transform fields to make domain object
	row := domain.RowID(strings.ToUpper(req.StartingSeat[:1]))
	startingSeatNum, err := strconv.Atoi(req.StartingSeat[1:])
	if err != nil {
		return err
	}
	numOfConsecutiveSeats, err := strconv.Atoi(req.NumOfConsecutiveSeats)
	if err != nil {
		return err
	}

	c := domain.NewCancellation(row, startingSeatNum, numOfConsecutiveSeats)

	if err := srv.airplaneRepository.Cancel(c); err != nil {
		return err
	}
	return nil
}
