package persistence

import (
	"database/sql"
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
