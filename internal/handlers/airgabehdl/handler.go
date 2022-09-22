// driver adapter
// An adapter for a driver port, transforms specific technology requests
// into call on a core service
package airgabehdl

import "github.com/guergabo/eks-final-round/internal/core/ports"

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
	return &Response{
		Status: success,
	}
}
