// driver adapter
// An adapter for a driver port, transforms specific technology requests
// into call on a core service
package airgabehdl

import (
	"fmt"

	"github.com/guergabo/eks-final-round/internal/core/dto"
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

func (hdl *CLHandler) Run(args []string) *dto.Response {
	if len(args) < 3 {
		if containsHelp(args) {
			return &dto.Response{Status: dto.Help}
		}
		return &dto.Response{Status: dto.Help + dto.RequestStatus(fmt.Sprintf("\n\nERROR: requires at least 3 arg(s), only received %d", len(args)))}
	}

	req := dto.NewRequest(args[0], args[1], args[2])

	// routing
	var requestStatus error
	switch action := dto.RequestSubCommand(req.Action); action {
	case dto.Book:
		requestStatus = hdl.airplaneService.Book(req)
	case dto.Cancel:
		requestStatus = hdl.airplaneService.Cancel(req)
	default:
		return &dto.Response{Status: dto.Fail}
	}

	if requestStatus != nil {
		return &dto.Response{Status: dto.Fail}
	}
	return &dto.Response{Status: dto.Success}
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
