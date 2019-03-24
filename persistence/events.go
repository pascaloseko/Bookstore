package persistence

import (
	"database/sql"
	"errors"
	"fmt"
	"os"

	_ "github.com/lib/pq"
)

// CRUD functions

// Db conn to database
var Db *sql.DB
var err error

func init() {
	username := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASS")
	dbName := os.Getenv("DB_NAME")

	dbURI := fmt.Sprintf("user=%s dbname=%s sslmode=disable password=%s", username, dbName, password)
	fmt.Println(dbURI)
	Db, err = sql.Open("postgres", dbURI)

	if err != nil {
		fmt.Printf("cannot open database: %v", err)
	}
}

// AddEvent adds an event
func (e *Event) AddEvent() (err error) {
	if e.LocationID >= 1 {
		err = errors.New("Location found")
	} else {
		err = errors.New("Location not found")
	}
	statement := "INSERT INTO event (name, duration, start_date, end_date, loc_id) VALUES($1, $2, $3, $4, (SELECT id FROM location WHERE id = $5)) RETURNING id;"
	stmt, err := Db.Prepare(statement)
	if err != nil {
		return err
	}
	defer stmt.Close()
	err = stmt.QueryRow(e.Name, e.Duration, e.StartDate, e.EndDate, e.LocationID).Scan(&e.ID)
	return
}

// GetEvents gets all available locations
func GetEvents() (events []Event, err error) {
	rows, err := Db.Query("SELECT * FROM event")
	if err != nil {
		return
	}

	for rows.Next() {
		event := Event{}
		if err = rows.Scan(&event.ID, &event.Name, &event.Duration, &event.StartDate, &event.EndDate, &event.LocationID); err != nil {
			return
		}

		events = append(events, event)
	}

	rows.Close()
	return
}
