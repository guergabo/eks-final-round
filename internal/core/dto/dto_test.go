package dto

import (
	"testing"

	"github.com/guergabo/eks-final-round/internal/core/domain"
)

func TestRequestValidation(t *testing.T) {
	t.Run("Valid booking request", func(t *testing.T) {

		// success
		mockRequest := NewRequest("BOOK", "A1", "1")

		actionConfig, err := mockRequest.ValidBookingRequest()
		if err != nil {
			t.Fatalf("unexpected error: %s", err)
		}

		if actionConfig.Row != domain.A {
			t.Fatalf("expected row A instead got row %s", actionConfig.Row)
		}

		if actionConfig.StartingRowIndex != 1 {
			t.Fatalf("expected starting row index 1 instead got %d", actionConfig.StartingRowIndex)
		}

		if actionConfig.NumOfConsecutiveSeats != 1 {
			t.Fatalf("expected number of consecutive seats with value 1 instead got %d", actionConfig.NumOfConsecutiveSeats)
		}

		// fail
		mockRequest = NewRequest("BOOK", "U1", "1")

		_, err = mockRequest.ValidBookingRequest()
		if err == nil {
			t.Fatalf("expected to handle an error")
		}
	})
	t.Run("Valid cancellation request", func(t *testing.T) {
		// success
		mockRequest := NewRequest("CANCEL", "A1", "1")

		actionConfig, err := mockRequest.ValidCancellationRequest()
		if err != nil {
			t.Fatalf("unexpected error: %s", err)
		}

		if actionConfig.Row != domain.A {
			t.Fatalf("expected row A instead got row %s", actionConfig.Row)
		}

		if actionConfig.StartingRowIndex != 1 {
			t.Fatalf("expected starting row index 1 instead got %d", actionConfig.StartingRowIndex)
		}

		if actionConfig.NumOfConsecutiveSeats != 1 {
			t.Fatalf("expected number of consecutive seats with value 1 instead got %d", actionConfig.NumOfConsecutiveSeats)
		}

		// fail
		mockRequest = NewRequest("CANCEL", "U1", "1")

		_, err = mockRequest.ValidCancellationRequest()
		if err == nil {
			t.Fatalf("expected to handle an error")
		}
	})
	t.Run("Is valid row", func(t *testing.T) {
		if !IsValidRow("A") {
			t.Fatal("expected true")
		}

		if IsValidRow("U") {
			t.Fatal("expected false")
		}
	})
	t.Run("Are valid seats", func(t *testing.T) {
		if !AreValidSeats(0, 3) {
			t.Fatal("expected true")
		}

		if AreValidSeats(3, 6) {
			t.Fatal("expected false")
		}
	})
}
