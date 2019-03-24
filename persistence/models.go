package persistence

import (
	"fmt"
)

// All data models will be here

// User struct
type User struct {
	ID        int    `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Age       int    `json:"age"`
	Bookings  []Booking
}

func (u *User) String() string {
	return fmt.Sprintf("id: %d, first_name: %s, last_name: %s, Age: %d, Bookings: %v", u.ID, u.FirstName, u.LastName, u.Age, u.Bookings)
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
