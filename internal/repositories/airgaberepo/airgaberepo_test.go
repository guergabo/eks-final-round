package airgaberepo

import (
	"os"
	"testing"

	"github.com/guergabo/eks-final-round/internal/core/domain"
)

func TestBookSeatsScenarioSimple(t *testing.T) {
	t.Run("", func(t *testing.T) {
		////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
		// "If a customer wants to reserve multiple seats together in the same row, we should be able to accommodate that or tell the customer it's not possible" //
		////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
		mockReq := &domain.Booking{
			Row:                   domain.A,
			StartingRowIndex:      0,
			NumOfConsecutiveSeats: 3,
		}

		mockLf := NewLocalFile()
		if err := mockLf.Book(mockReq); err != nil {
			t.Fatalf("unexpected error: %s", err)
		}

		// check state file updates
		airplane, err := loadState()
		if err != nil {
			t.Fatalf("unexpected error: %s", err)
		}

		seats := airplane.Rows[domain.RowID("A")].Seats[0:3]
		for i := range seats {
			if seats[i].Status != domain.Booked {
				t.Fatalf("error seat to be book but instead got: %s", seats[i].Status)
			}
		}

		///////////////////////////////////////////////////////////////
		// "A given seat cannot be reserved by more than one person" //
		///////////////////////////////////////////////////////////////
		mockReq = &domain.Booking{
			Row:                   domain.A,
			StartingRowIndex:      2,
			NumOfConsecutiveSeats: 3,
		}

		mockLf = NewLocalFile()
		if err := mockLf.Book(mockReq); err == nil {
			t.Fatal("expected to handle error")
		}

		// check if state file was modified
		airplane, err = loadState()
		if err != nil {
			t.Fatalf("unexpected error: %s", err)
		}

		attemptedSeats := airplane.Rows[domain.RowID("A")].Seats[2:5]
		if attemptedSeats[0].Status != domain.Booked {
			t.Fatalf("error expected seat A2 to be book but instead got: %s", attemptedSeats[0].Status)
		}
		availableSeats := attemptedSeats[1:]
		for i := range availableSeats {
			if availableSeats[i].Status != domain.Available {
				t.Fatalf("error expected seat to be available but instead got: %s", availableSeats[i].Status)
			}
		}

		//////////////////////////////////////////////////////////////////////////////////////////
		// "If a customer cancels their reservation, the seat is available for reserving again" //
		//////////////////////////////////////////////////////////////////////////////////////////
		mockCancelReq := &domain.Cancellation{
			Row:                   domain.A,
			StartingRowIndex:      1,
			NumOfConsecutiveSeats: 2,
		}

		mockLf = NewLocalFile()
		if err := mockLf.Cancel(mockCancelReq); err != nil {
			t.Fatalf("unexpected error %s", err)
		}

		// check if state file was modified
		airplane, err = loadState()
		if err != nil {
			t.Fatalf("unexpected error: %s", err)
		}

		seats = airplane.Rows[domain.RowID("A")].Seats
		if seats[0].Status != domain.Booked {
			t.Fatalf("error expected seat to be booked but instead got: %s", seats[0].Status)
		}

		for _, v := range seats[1:] {
			if v.Status != domain.Available {
				t.Fatalf("error expected seat to be available but instead got: %s", v.Status)
			}
		}
	})

	os.Remove("current-state.json")
}

func TestBookSeatsScenarioExample(t *testing.T) {
}
