-- Zones table
CREATE TABLE IF NOT EXISTS zones (
    id TEXT PRIMARY KEY,
    user_id TEXT NOT NULL DEFAULT 'default_user',
    name TEXT NOT NULL,
    area_sqft REAL,
    hardiness_zone TEXT,
    last_frost DATE,
    first_frost DATE,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- Plant intents table
CREATE TABLE IF NOT EXISTS plant_intents (
    id TEXT PRIMARY KEY,
    zone_id TEXT NOT NULL,
    crop_type TEXT NOT NULL,
    quantity INTEGER NOT NULL,
    preferences TEXT, -- JSON array stored as text
    status TEXT DEFAULT 'planned',
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (zone_id) REFERENCES zones(id) ON DELETE CASCADE
);

-- Timeline events table
CREATE TABLE IF NOT EXISTS timeline_events (
    id TEXT PRIMARY KEY,
    plant_intent_id TEXT NOT NULL,
    event_type TEXT NOT NULL,
    scheduled_date DATE NOT NULL,
    status TEXT DEFAULT 'upcoming',
    completed_at TIMESTAMP,
    notes TEXT,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (plant_intent_id) REFERENCES plant_intents(id) ON DELETE CASCADE
);

-- Crop templates table (hardcoded data)
CREATE TABLE IF NOT EXISTS crop_templates (
    crop_type TEXT PRIMARY KEY,
    days_to_transplant INTEGER,
    days_to_harvest INTEGER,
    space_per_plant REAL,
    weeks_before_frost INTEGER,
    can_direct_sow BOOLEAN DEFAULT 0
);

-- Insert initial crop data
INSERT INTO crop_templates (crop_type, days_to_transplant, days_to_harvest, space_per_plant, weeks_before_frost, can_direct_sow) VALUES
('tomato', 56, 80, 4.0, 8, 0),
('pepper', 70, 90, 4.0, 10, 0),
('lettuce', 28, 45, 1.0, 4, 0),
('basil', 0, 60, 1.0, 0, 0),  -- buy starts, no seed start
('beans', 0, 55, 0.5, -2, 1);  -- direct sow after frost
