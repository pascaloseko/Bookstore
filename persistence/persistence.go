package persistence

// DatabaseHandler interface for db persistence
type DatabaseHandler interface {
	AddEvent(Event) (error)
	FindEvent([]byte) (Event, error)
	FindEventByName(string) (Event, error)
	FindAllAvailableEvents() ([]Event, error)
}
