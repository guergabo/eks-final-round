package airgabehdl

type requestStatus string

const (
	success requestStatus = "SUCCESS"
	fail    requestStatus = "FAIL"
)

type Response struct {
	Status requestStatus
}

func NewResponse(ars requestStatus) *Response {
	return &Response{Status: ars}
}
