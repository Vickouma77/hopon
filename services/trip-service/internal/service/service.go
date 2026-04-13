package service

import (
	"context"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"hopon.io/services/trip-service/internal/domain"
)

// service is an implementation of domain.TripService.
type service struct {
	repo domain.TripRepository
}

// NewService creates a new instance of TripService.
func NewService(repo domain.TripRepository) *service {
	return &service{
		repo: repo,
	}
}

// CreateTrip creates a new trip based on the provided ride fare.
func (s *service) CreateTrip(ctx context.Context, fare *domain.RideFareModel) (*domain.TripModel, error) {
	t := &domain.TripModel{
		ID:       primitive.NewObjectID(),
		USerID:   fare.UserID,
		Status:   "pending",
		RideFare: fare,
	}

	return s.repo.CreateTrip(ctx, t)
}
