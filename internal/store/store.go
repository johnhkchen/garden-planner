package store

import (
	"github.com/johnhkchen/garden-planner/internal/domain"
)

// Store defines the interface for data persistence
type Store interface {
	// Zones
	CreateZone(zone *domain.Zone) error
	GetZone(id string) (*domain.Zone, error)
	ListZones(userID string) ([]*domain.Zone, error)
	UpdateZone(zone *domain.Zone) error
	DeleteZone(id string) error

	// Plant Intents
	CreatePlantIntent(intent *domain.PlantIntent) error
	GetPlantIntent(id string) (*domain.PlantIntent, error)
	ListPlantIntents(zoneID string) ([]*domain.PlantIntent, error)
	UpdatePlantIntent(intent *domain.PlantIntent) error
	DeletePlantIntent(id string) error

	// Timeline Events
	CreateTimelineEvent(event *domain.TimelineEvent) error
	GetTimelineEvent(id string) (*domain.TimelineEvent, error)
	ListTimelineEvents(plantIntentID string) ([]*domain.TimelineEvent, error)
	UpdateTimelineEvent(event *domain.TimelineEvent) error
	DeleteTimelineEvent(id string) error

	// Helper: Get all events for a zone
	GetZoneTimeline(zoneID string) ([]*domain.TimelineEvent, error)
}
