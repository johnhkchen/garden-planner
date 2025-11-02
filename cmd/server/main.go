package main

import (
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/johnhkchen/garden-planner/internal/domain"
	"github.com/johnhkchen/garden-planner/internal/store"
)

func main() {
	// Initialize database
	db, err := store.NewSQLiteStore("./data/garden.db")
	if err != nil {
		log.Fatal("Failed to open database:", err)
	}
	defer db.Close()

	// Create a test handler
	http.HandleFunc("/", homeHandler)
	http.HandleFunc("/api/zones", zonesHandler(db))
	http.HandleFunc("/api/test", testHandler(db))

	// Static files
	fs := http.FileServer(http.Dir("./web/static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	// Start server
	addr := ":8080"
	log.Printf("Server running on http://localhost%s", addr)
	log.Fatal(http.ListenAndServe(addr, nil))
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	html := `
		<!DOCTYPE html>
		<html>
		<head>
			<title>Garden Planner</title>
			<link rel="stylesheet" href="/static/css/style.css">
		</head>
		<body>
			<h1>Garden Planner MVP</h1>
			<p>Server is running!</p>
			<ul>
				<li><a href="/api/zones">View Zones (JSON)</a></li>
				<li><a href="/api/test">Run Test (Create sample zone)</a></li>
			</ul>
		</body>
		</html>
	`
	w.Header().Set("Content-Type", "text/html")
	w.Write([]byte(html))
}

func zonesHandler(db *store.SQLiteStore) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		zones, err := db.ListZones("default_user")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(zones)
	}
}

func testHandler(db *store.SQLiteStore) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Create a test zone
		zone := domain.NewZone(
			"Test Bed A",
			32.0, // 4x8 ft
			"10a",
			time.Date(2025, 2, 15, 0, 0, 0, 0, time.UTC),
			time.Date(2025, 11, 30, 0, 0, 0, 0, time.UTC),
		)
		zone.UserID = "default_user"

		if err := db.CreateZone(zone); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]interface{}{
			"message": "Zone created successfully",
			"zone":    zone,
		})
	}
}
