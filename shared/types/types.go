package types

// Route represents the route calculated between pickup and destination.
type Route struct {
	Distance float64     `json:"distance"`
	Duration float64     `json:"duration"`
	Geometry []*Geometry `json:"geometry"`
}

// Geometry contains the coordinate sequence for rendering the route.
type Geometry struct {
	Coordinates []*Coordinate `json:"coordinates"`
}

// Coordinate represents a geographical point with latitude and longitude.
type Coordinate struct {
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
}
