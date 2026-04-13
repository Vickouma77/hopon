package domain

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// RideFareModel represents the fare details for a requested ride.
type RideFareModel struct {
	ID                primitive.ObjectID
	UserID            string
	PackageSlug       string
	TotalPriceInCents float64
}
