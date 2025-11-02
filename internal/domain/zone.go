package domain

import "time"
import "github.com/google/uuid"

// Zone represents a physical growing area (bed, field, container)
type Zone struct {
	ID            string    `json:"id"`
	UserID        string    `json:"user_id"` // Future: for multi-user
	Name          string    `json:"name"`
	AreaSqFt      float64   `json:"area_sqft"`
	HardinessZone string    `json:"hardiness_zone"` // e.g., "10a"
	LastFrost     time.Time `json:"last_frost"`
	FirstFrost    time.Time `json:"first_frost"`
	CreatedAt     time.Time `json:"created_at"`
}

// NewZone creates a new zone with generated ID
func NewZone(name string, areaSqFt float64, zone string, lastFrost, firstFrost time.Time) *Zone {
	return &Zone{
		ID:            generateID(), // TODO: implement this
		Name:          name,
		AreaSqFt:      areaSqFt,
		HardinessZone: zone,
		LastFrost:     lastFrost,
		FirstFrost:    firstFrost,
		CreatedAt:     time.Now(),
	}
}

// Helper function to generate unique IDs (simple version)
func generateID() string {
	return uuid.New().String()
}
