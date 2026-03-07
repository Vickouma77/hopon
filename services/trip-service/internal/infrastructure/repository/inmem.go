package repository

import (
	"context"

	"hopon.io/services/trip-service/internal/domain"
)

type InmemRepository struct {
	trips    map[string]*domain.TripModel
	rideFare map[string]*domain.RideFareModel
}

func NewInmemRepository() *InmemRepository {
	return &InmemRepository{
		trips:    make(map[string]*domain.TripModel),
		rideFare: make(map[string]*domain.RideFareModel),
	}
}

func (r *InmemRepository) CreateTrip(ctx context.Context, trip *domain.TripModel) (*domain.TripModel, error) {
	r.trips[trip.ID.Hex()] = trip

	return trip, nil
}
