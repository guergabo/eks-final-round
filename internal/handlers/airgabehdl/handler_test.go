package airgabehdl

import (
	"errors"
	"fmt"
	"testing"

	"github.com/guergabo/eks-final-round/internal/core/dto"
)

// mock service
type MockService struct {
	MockError error
}

func (ms *MockService) Book(req *dto.Request) error {
	return ms.MockError
}

func (ms *MockService) Cancel(req *dto.Request) error {
	return ms.MockError
}

func TestHandler(t *testing.T) {
	t.Run("Help Requests or too few arguments", func(t *testing.T) {
		mockArgs := []string{}
		mockSvc := MockService{MockError: nil}

		h := NewCLHandler(&mockSvc)

		// no arguments at all
		resp := h.Run(mockArgs)
		if resp.Status != (dto.Help + dto.RequestStatus(fmt.Sprintf("\n\nERROR: requires at least 3 arg(s), only received %d", len(mockArgs)))) {
			t.Fatalf("expected Help response instead got %s", resp.Status)
		}

		// help flag found
		mockArgs = append(mockArgs, "--help")
		resp = h.Run(mockArgs)
		if resp.Status != dto.Help {
			t.Fatalf("expected Help response instead got %s", resp.Status)
		}
	})

	t.Run("Successful requests", func(t *testing.T) {
		mockArgs := []string{"BOOK", "A1", "1"}
		mockSvc := MockService{MockError: nil}

		h := NewCLHandler(&mockSvc)

		// successful book
		resp := h.Run(mockArgs)
		if resp.Status != dto.Success {
			t.Fatalf("expected SUCCESS response instead got %s", resp.Status)
		}

		// successful cancel
		mockArgs[0] = "CANCEL"
		resp = h.Run(mockArgs)
		if resp.Status != dto.Success {
			t.Fatalf("expected SUCCESS response instead got %s", resp.Status)
		}
	})

	t.Run("Failed requests", func(t *testing.T) {
		mockArgs := []string{"BOOK", "U1", "1"}
		mockSvc := MockService{MockError: errors.New("something went wrong")}

		h := NewCLHandler(&mockSvc)

		// unsuccessful book
		resp := h.Run(mockArgs)
		if resp.Status != dto.Fail {
			t.Fatalf("expected FAIL response instead got %s", resp.Status)
		}

		// unsuccessful cancel
		mockArgs[0] = "CANCEL"
		resp = h.Run(mockArgs)
		if resp.Status != dto.Fail {
			t.Fatalf("expected FAIL response instead got %s", resp.Status)
		}
	})
}
