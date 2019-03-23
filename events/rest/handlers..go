package rest

import (
	"net/http"
)

// FindEventHandler search for events
func FindEventHandler(w http.ResponseWriter, r *http.Request){
	http.Error(w, "This feature is coming soon", http.StatusNotImplemented)
}

// AllEventHandler gets all available events
func AllEventHandler(w http.ResponseWriter, r *http.Request){
	http.Error(w, "This feature is coming soon", http.StatusNotImplemented)
}

// NewEventHandler adds a new event
func NewEventHandler(w http.ResponseWriter, r *http.Request){
	http.Error(w, "This feature is coming soon", http.StatusNotImplemented)
}

// UpdatEventHandler updates event instance
func UpdatEventHandler(w http.ResponseWriter, r *http.Request){
	http.Error(w, "This feature is coming soon", http.StatusNotImplemented)
}

// DeletEventHandler deletes an event instance
func DeletEventHandler(w http.ResponseWriter, r *http.Request){
	http.Error(w, "This feature is coming soon", http.StatusNotImplemented)
}