package ports

import "github.com/guergabo/eks-final-round/internal/core/dto"

// take care of error handling
type AirplaneService interface {
	Book(req *dto.Request) error
	Cancel(req *dto.Request) error
}
