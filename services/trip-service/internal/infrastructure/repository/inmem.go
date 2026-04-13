package repository

import (
	"context"

	"hopon.io/services/trip-service/internal/domain"
)

// InmemRepository provides an in-memory implementation of the domain.TripRepository interface.
type InmemRepository struct {
	trips    map[string]*domain.TripModel
	rideFare map[string]*domain.RideFareModel
}

// NewInmemRepository creates a new instance of InmemRepository.
func NewInmemRepository() *InmemRepository {
	return &InmemRepository{
		trips:    make(map[string]*domain.TripModel),
		rideFare: make(map[string]*domain.RideFareModel),
	}
}

// CreateTrip saves a trip to the in-memory store.
func (r *InmemRepository) CreateTrip(ctx context.Context, trip *domain.TripModel) (*domain.TripModel, error) {
	r.trips[trip.ID.Hex()] = trip

	return trip, nil
}
