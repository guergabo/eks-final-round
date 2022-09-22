package airgabehdl

type requestStatus string

const (
	success requestStatus = "SUCCESS"
	fail    requestStatus = "FAIL"
	help    requestStatus = `
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
	Status requestStatus
}

func NewResponse(ars requestStatus) *Response {
	return &Response{Status: ars}
}
