// driver adapter
// An adapter for a driver port, transforms specific technology requests
// into call on a core service
package airgabehdl

import (
	"log"
	"strconv"

	"github.com/guergabo/eks-final-round/internal/core/ports"
)

type CLHandler struct {
	airplaneService ports.AirplaneService
}

func NewCLHandler(airplaneService ports.AirplaneService) *CLHandler {
	return &CLHandler{
		airplaneService: airplaneService,
	}
}

// transform command line request into a service call
func (hdl *CLHandler) Run(args []string) *Response {

	// requires [Action] [Starting Seat] [Number of Consecutive Seats]
	if len(args) < 3 {
		return &Response{
			Status: fail,
		}
	}

	numOfConsecutiveSeats, err := strconv.Atoi(args[2])
	if err != nil {
		return &Response{
			Status: fail,
		}
	}

	req := Request{
		Action:                args[0],
		StartingSeat:          args[1],
		NumOfConsecutiveSeats: numOfConsecutiveSeats,
	}

	log.Printf("Action: %s, Starting Seat: %s, Number of Consecutive Seats: %d\n", req.Action, req.StartingSeat, req.NumOfConsecutiveSeats)

	var requestStatus error
	switch action := req.Action; action {
	case "BOOK":
		log.Println("calling booking service")
		requestStatus = hdl.airplaneService.Book(req.StartingSeat, numOfConsecutiveSeats)
	case "CANCEL":
		log.Println("calling canceling service")
		requestStatus = hdl.airplaneService.Cancel(req.StartingSeat, req.NumOfConsecutiveSeats)
	default:
		return &Response{
			Status: fail,
		}
	}

	if requestStatus != nil {
		return &Response{
			Status: fail,
		}
	}
	return &Response{
		Status: success,
	}
}
