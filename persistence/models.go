package persistence

import (
	"fmt"
)

// All data models will be here

// User struct
type User struct {
	UserID        int    `json:"userId"`
	FirstName string `json:"userName"`
}

func (u *User) String() string {
	return fmt.Sprintf("id: %s, first_name: %s", u.UserID, u.FirstName)
}

// Booking struct
type Booking struct {
	ID      int    `json:"id"`
	Date    int64  `json:"date"`
	EventID []byte `json:"event"`
	Seats   int    `json:"seats"`
}

// Event struct
type Event struct {
	ID         int    `json:"id"`
	Name       string `json:"name"`
	Duration   int    `json:"duration"`
	StartDate  string `json:"startdate"`
	EndDate    string `json:"enddate"`
	LocationID int    `json:"location"`
}

// Location struct
type Location struct {
	ID        int    `json:"id"`
	Name      string `json:"name"`
	Address   string `json:"address"`
	Country   string `json:"country"`
	OpenTime  int    `json:"opentime"`
	CloseTime int    `json:"closetime"`
}
