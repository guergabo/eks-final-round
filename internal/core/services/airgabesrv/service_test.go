package airgabesrv

import (
	"errors"
	"testing"
)

// use a mock request
// mock repo
type MockRepo struct {
	MockError error
}

func (mr *MockRepo) Book(startingSeat string, numOfConsecutiveSeats int) error {
	return mr.MockError
}

func (mr *MockRepo) Cancel(startingSeat string, numOfConsecutiveSeats int) error {
	return mr.MockError
}

func TestHandler(t *testing.T) {
	t.Run("Successful Requests", func(t *testing.T) {
		mockSvc := MockRepo{MockError: nil}

		h := New(&mockSvc)
		if err := h.Book("mockStartinSeat", 3); err != nil {
			t.Fatalf("unexpected error, %s", err)
		}

		if err := h.Cancel("mockStartinSeat", 3); err != nil {
			t.Fatalf("unexpected error, %s", err)
		}
	})

	t.Run("Failed Requests", func(t *testing.T) {
		mockSvc := MockRepo{MockError: errors.New("something went wrong")}

		h := New(&mockSvc)
		if err := h.Book("mockStartinSeat", 3); err == nil {
			t.Fatal("expected to handle an error")
		}

		if err := h.Cancel("mockStartinSeat", 3); err == nil {
			t.Fatal("expected to handle an error")
		}
	})
}
