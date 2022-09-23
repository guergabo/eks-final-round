package ports

import "github.com/guergabo/eks-final-round/internal/core/domain"

type AirplaneRepository interface {
	Book(book *domain.Booking) error
	Cancel(cancel *domain.Cancellation) error
}
