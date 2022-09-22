// driver adapter
// An adapter for a driver port, transforms specific technology requests
// into call on a core service
package airgabehdl

import (
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

func (hdl *CLHandler) Run(args []string) *Response {
	// requires [Action] [Starting Seat] [Number of Consecutive Seats]
	if len(args) == 0 || containsHelp(args) {
		return &Response{Status: help}
	}
	if len(args) < 3 {
		return &Response{Status: fail}
	}

	// transform command line request into a service request
	numOfConsecutiveSeats, err := strconv.Atoi(args[2])
	if err != nil {
		return &Response{Status: fail}
	}
	req := NewRequest(args[0], args[1], numOfConsecutiveSeats)

	// route service request
	var requestStatus error
	switch action := requestSubCommand(req.Action); action {
	case book:
		requestStatus = hdl.airplaneService.Book(req.StartingSeat, numOfConsecutiveSeats)
	case cancel:
		requestStatus = hdl.airplaneService.Cancel(req.StartingSeat, req.NumOfConsecutiveSeats)
	default:
		return &Response{Status: fail}
	}

	// response to customer
	if requestStatus != nil {
		return &Response{Status: fail}
	}
	return &Response{Status: success}
}

// private methods
func containsHelp(args []string) bool {
	for _, v := range args {
		if v == "-h" || v == "--help" {
			return true
		}
	}
	return false
}
