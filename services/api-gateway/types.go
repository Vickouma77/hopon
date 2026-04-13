package main

import "hopon.io/shared/types"

// PreviewTripRequest represents the payload for requesting a trip preview/estimate.
type PreviewTripRequest struct {
	UserID      string           `json:"userID"`
	Pickup      types.Coordinate `json:"pickup"`
	Destination types.Coordinate `json:"destination"`
}
