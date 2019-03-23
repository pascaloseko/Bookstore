package persistence

import (
	"database/sql"
	"errors"
	"fmt"

	"github.com/pascaloseko/Events/lib/config"

	_ "github.com/lib/pq"
)

// CRUD functions

// Db conn to database
var Db *sql.DB
var err error

func init() {
	dbAddr := "dbname=" + config.Config.DbName + " user=" + config.Config.DbUser + " port=" + config.Config.DbPort +
		" password=" + config.Config.DbPassword
	Db, err = sql.Open("postgres", dbAddr)

	if err != nil {
		fmt.Printf("cannot open database: %v", err)
	}
}

// AddEvent adds an event
func (e Event) AddEvent() (err error) {
	if e.Location == nil {
		err = errors.New("Location not found")
	}
	statement := "INSERT INTO events (name, duration, start_date, end_date, location) VALUES($1, $2, $3, $4, $5) RETURNING id;"
	stmt, err := Db.Prepare(statement)
	if err != nil {
		return
	}

	defer stmt.Close()
	err = stmt.QueryRow(e.Name, e.Duration, e.StartDate, e.EndDate, e.Location.ID).Scan(&e.ID)
	return
}