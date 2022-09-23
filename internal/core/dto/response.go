package dto

type RequestStatus string

const (
	Success RequestStatus = "SUCCESS"
	Fail    RequestStatus = "FAIL"
	Help    RequestStatus = `
Usage:
	airgabe [action] [starting seats] [number of consecutive seats]

Examples:

:: book seat A1
airgabe BOOK A1 1

:: cancel seat A1 
airgabe CANCEL A1 1

Flags:
	-h, --help, help for airgabe`
)

type Response struct {
	Status RequestStatus
}

func NewResponse(ars RequestStatus) *Response {
	return &Response{Status: ars}
}
