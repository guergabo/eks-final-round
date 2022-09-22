package airgabehdl

type airplaneRequestStatus string

const (
	success airplaneRequestStatus = "SUCCESS"
	fail    airplaneRequestStatus = "FAIL"
)

type Response struct {
	Status airplaneRequestStatus
}

func NewResponse(ars airplaneRequestStatus) *Response {
	return &Response{
		Status: ars,
	}
}
