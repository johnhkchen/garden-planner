package domain

import "time"

// PlantingStatus represents the lifecycle state
type PlantingStatus string

const (
	StatusPlanned    PlantingStatus = "planned"
	StatusStarted    PlantingStatus = "started"
	StatusGrowing    PlantingStatus = "growing"
	StatusHarvesting PlantingStatus = "harvesting"
	StatusFinished   PlantingStatus = "finished"
)

// PlantIntent represents a user's intention to grow something
type PlantIntent struct {
	ID          string         `json:"id"`
	ZoneID      string         `json:"zone_id"`
	CropType    string         `json:"crop_type"` // "tomato", "pepper", etc.
	Quantity    int            `json:"quantity"`
	Preferences []string       `json:"preferences"` // ["determinate", "heirloom"]
	Status      PlantingStatus `json:"status"`
	CreatedAt   time.Time      `json:"created_at"`
}

// NewPlantIntent creates a new planting intention
func NewPlantIntent(zoneID, cropType string, quantity int) *PlantIntent {
	return &PlantIntent{
		ID:        generateID(),
		ZoneID:    zoneID,
		CropType:  cropType,
		Quantity:  quantity,
		Status:    StatusPlanned,
		CreatedAt: time.Now(),
	}
}
