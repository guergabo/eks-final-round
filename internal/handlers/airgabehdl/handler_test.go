package airgabehdl

import (
	"errors"
	"testing"
)

// mock service
type MockService struct {
	MockError error
}

func (ms *MockService) Book(startingSeat string, numOfConsecutiveSeats int) error {
	return ms.MockError
}

func (ms *MockService) Cancel(startingSeat string, numOfConsecutiveSeats int) error {
	return ms.MockError
}

func TestHandler(t *testing.T) {
	t.Run("Argument soft validation issues", func(t *testing.T) {
		mockArgs := []string{"BOOK", "A1"}
		mockSvc := MockService{MockError: nil}

		h := NewCLHandler(&mockSvc)

		// too few arguments
		resp := h.Run(mockArgs)
		if resp.Status == success {
			t.Fatalf("expected FAIL response instead got %s", resp.Status)
		}

		// consecutive seats argument is not a valid number
		mockArgs = append(mockArgs, "1t5")
		resp = h.Run(mockArgs)
		if resp.Status != fail {
			t.Fatalf("expected FAIL response instead got %s", resp.Status)
		}

	})

	t.Run("Successful requests", func(t *testing.T) {
		mockArgs := []string{"BOOK", "A1", "1"}
		mockSvc := MockService{MockError: nil}

		h := NewCLHandler(&mockSvc)

		// successful book
		resp := h.Run(mockArgs)
		if resp.Status != success {
			t.Fatalf("expected SUCCESS response instead got %s", resp.Status)
		}

		// successful cancel
		mockArgs[0] = "CANCEL"
		resp = h.Run(mockArgs)
		if resp.Status != success {
			t.Fatalf("expected SUCCESS response instead got %s", resp.Status)
		}
	})

	t.Run("Failed requests", func(t *testing.T) {
		mockArgs := []string{"BOOK", "A1", "1"}
		mockSvc := MockService{MockError: errors.New("something went wrong")}

		h := NewCLHandler(&mockSvc)

		// unsuccessful book
		resp := h.Run(mockArgs)
		if resp.Status != fail {
			t.Fatalf("expected FAIL response instead got %s", resp.Status)
		}

		// unsuccessful cancel
		mockArgs[0] = "CANCEL"
		resp = h.Run(mockArgs)
		if resp.Status != fail {
			t.Fatalf("expected FAIL response instead got %s", resp.Status)
		}
	})
}
