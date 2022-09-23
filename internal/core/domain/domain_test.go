package domain

import (
	"testing"
)

func TestAirplaneValidation(t *testing.T) {
	t.Run("Are seats available", func(t *testing.T) {
		a := NewAirplane()
		a.Rows = map[RowID]Row{A: {Seats: []Seat{{Available}, {Available}, {Available}}}}

		mockRowID, mockStart, mockNum := A, 1, 1

		// success
		resp := a.AreSeatsAvailable(mockRowID, mockStart, mockNum)
		if !resp {
			t.Fatalf("expected to get true instead got %v", resp)
		}

		// fail
		a.Rows = map[RowID]Row{A: {Seats: []Seat{{Booked}, {Booked}, {Booked}}}}
		resp = a.AreSeatsAvailable(mockRowID, mockStart, mockNum)
		if resp {
			t.Fatalf("expected to get false instead got %v", resp)
		}
	})
	t.Run("Are seats booked", func(t *testing.T) {
		a := NewAirplane()
		a.Rows = map[RowID]Row{A: {Seats: []Seat{{Booked}, {Booked}, {Booked}}}}

		mockRowID, mockStart, mockNum := A, 1, 1

		// success
		resp := a.AreSeatsBooked(mockRowID, mockStart, mockNum)
		if !resp {
			t.Fatalf("expected to get true instead got %v", resp)
		}

		// fail
		a.Rows = map[RowID]Row{A: {Seats: []Seat{{Available}, {Available}, {Available}}}}
		resp = a.AreSeatsBooked(mockRowID, mockStart, mockNum)
		if resp {
			t.Fatalf("expected to get false instead got %v", resp)
		}

	})
	t.Run("Process booking", func(t *testing.T) {
		a := NewAirplane()
		a.Rows = map[RowID]Row{A: {Seats: []Seat{{Available}, {Available}, {Available}}}}
		mockRowID, mockStart, mockNum := A, 0, 3

		// success
		a.ProcessBooking(mockRowID, mockStart, mockNum)

		for _, v := range a.Rows[A].Seats {
			if v.Status == Available {
				t.Fatalf("expected to get Booked instead got %s", v)
			}
		}
	})
	t.Run("Process cancellation", func(t *testing.T) {
		a := NewAirplane()
		a.Rows = map[RowID]Row{A: {Seats: []Seat{{Booked}, {Booked}, {Booked}}}}
		mockRowID, mockStart, mockNum := A, 0, 3

		// success
		a.ProcessCancellation(mockRowID, mockStart, mockNum)

		for _, v := range a.Rows[A].Seats {
			if v.Status == Booked {
				t.Fatalf("expected to get Available instead got %s", v)
			}
		}
	})
}
