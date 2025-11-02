package store

import (
	"database/sql"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
	"github.com/johnhkchen/garden-planner/internal/domain"
)

// SQLiteStore implements Store using SQLite
type SQLiteStore struct {
	db *sql.DB
}

// NewSQLiteStore creates a new SQLite store
func NewSQLiteStore(dbPath string) (*SQLiteStore, error) {
	db, err := sql.Open("sqlite3", dbPath)
	if err != nil {
		return nil, fmt.Errorf("failed to open database: %w", err)
	}

	// Enable foreign keys (SQLite doesn't enable by default)
	if _, err := db.Exec("PRAGMA foreign_keys = ON"); err != nil {
		return nil, fmt.Errorf("failed to enable foreign keys: %w", err)
	}

	// Test connection
	if err := db.Ping(); err != nil {
		return nil, fmt.Errorf("failed to ping database: %w", err)
	}

	return &SQLiteStore{db: db}, nil
}

// Close closes the database connection
func (s *SQLiteStore) Close() error {
	return s.db.Close()
}

// CreateZone inserts a new zone
func (s *SQLiteStore) CreateZone(zone *domain.Zone) error {
	query := `
		INSERT INTO zones (id, user_id, name, area_sqft, hardiness_zone, last_frost, first_frost, created_at)
		VALUES (?, ?, ?, ?, ?, ?, ?, ?)
	`
	_, err := s.db.Exec(query,
		zone.ID,
		zone.UserID,
		zone.Name,
		zone.AreaSqFt,
		zone.HardinessZone,
		zone.LastFrost,
		zone.FirstFrost,
		zone.CreatedAt,
	)
	return err
}

// GetZone retrieves a zone by ID
func (s *SQLiteStore) GetZone(id string) (*domain.Zone, error) {
	query := `
		SELECT id, user_id, name, area_sqft, hardiness_zone, last_frost, first_frost, created_at
		FROM zones
		WHERE id = ?
	`
	zone := &domain.Zone{}
	err := s.db.QueryRow(query, id).Scan(
		&zone.ID,
		&zone.UserID,
		&zone.Name,
		&zone.AreaSqFt,
		&zone.HardinessZone,
		&zone.LastFrost,
		&zone.FirstFrost,
		&zone.CreatedAt,
	)
	if err == sql.ErrNoRows {
		return nil, fmt.Errorf("zone not found: %s", id)
	}
	if err != nil {
		return nil, err
	}
	return zone, nil
}

// ListZones retrieves all zones for a user
func (s *SQLiteStore) ListZones(userID string) ([]*domain.Zone, error) {
	query := `
		SELECT id, user_id, name, area_sqft, hardiness_zone, last_frost, first_frost, created_at
		FROM zones
		WHERE user_id = ?
		ORDER BY created_at DESC
	`
	rows, err := s.db.Query(query, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var zones []*domain.Zone
	for rows.Next() {
		zone := &domain.Zone{}
		if err := rows.Scan(
			&zone.ID,
			&zone.UserID,
			&zone.Name,
			&zone.AreaSqFt,
			&zone.HardinessZone,
			&zone.LastFrost,
			&zone.FirstFrost,
			&zone.CreatedAt,
		); err != nil {
			return nil, err
		}
		zones = append(zones, zone)
	}
	return zones, rows.Err()
}

// TODO: Implement remaining methods (UpdateZone, DeleteZone, PlantIntent methods, etc.)
