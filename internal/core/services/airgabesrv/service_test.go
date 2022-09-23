package airgabesrv

import (
	"errors"
	"testing"

	"github.com/guergabo/eks-final-round/internal/core/domain"
	"github.com/guergabo/eks-final-round/internal/core/dto"
)

// use a mock request
// mock repo
type MockRepo struct {
	MockError error
}

func (mr *MockRepo) Book(book *domain.Booking) error {
	return mr.MockError
}

func (mr *MockRepo) Cancel(cancel *domain.Cancellation) error {
	return mr.MockError
}

func TestHandler(t *testing.T) {
	t.Run("Successful Requests", func(t *testing.T) {
		mockSvc := MockRepo{MockError: nil}
		mockRequest := dto.NewRequest("BOOK", "A1", "1")

		h := New(&mockSvc)
		if err := h.Book(mockRequest); err != nil {
			t.Fatalf("unexpected error, %s", err)
		}

		mockRequest = dto.NewRequest("CANCEL", "A1", "1")
		if err := h.Cancel(mockRequest); err != nil {
			t.Fatalf("unexpected error, %s", err)
		}
	})

	t.Run("Failed Requests", func(t *testing.T) {
		mockSvc := MockRepo{MockError: errors.New("something went wrong")}

		// row is invalid
		mockRequest := dto.NewRequest("BOOK", "U1", "7")
		h := New(&mockSvc)
		if err := h.Book(mockRequest); err == nil {
			t.Fatal("expected to handle an error")
		}

		// starting seat is invalid
		mockRequest = dto.NewRequest("BOOK", "B25", "1")
		if err := h.Book(mockRequest); err == nil {
			t.Fatal("expected to handle an error")
		}

		// number of consecutive seats is invalid
		mockRequest = dto.NewRequest("CANCEL", "B1", "21")
		if err := h.Cancel(mockRequest); err == nil {
			t.Fatal("expected to handle an error")
		}
	})
}
