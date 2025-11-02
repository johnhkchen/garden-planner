package domain

import "time"

// EventType represents different milestones
type EventType string

const (
	EventStartSeeds  EventType = "start_seeds"
	EventTransplant  EventType = "transplant"
	EventDirectSow   EventType = "direct_sow"
	EventFirstHarvest EventType = "first_harvest"
	EventRemove      EventType = "remove"
)

// EventStatus tracks completion state
type EventStatus string

const (
	EventUpcoming EventStatus = "upcoming"
	EventDueSoon  EventStatus = "due_soon"
	EventOverdue  EventStatus = "overdue"
	EventCompleted EventStatus = "completed"
	EventSkipped  EventStatus = "skipped"
)

// TimelineEvent represents a milestone in the growing timeline
type TimelineEvent struct {
	ID            string      `json:"id"`
	PlantIntentID string      `json:"plant_intent_id"`
	EventType     EventType   `json:"event_type"`
	ScheduledDate time.Time   `json:"scheduled_date"`
	Status        EventStatus `json:"status"`
	CompletedAt   *time.Time  `json:"completed_at,omitempty"`
	Notes         string      `json:"notes"`
	CreatedAt     time.Time   `json:"created_at"`
}
