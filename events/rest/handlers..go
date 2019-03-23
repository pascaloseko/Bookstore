package rest

import (
	"encoding/json"
	"net/http"

	"github.com/pascaloseko/Events/persistence"
)

type eventServiceHandler struct {
	dbhandler persistence.DatabaseHandler
}

// New db handler
func New(databasehandler persistence.DatabaseHandler) *eventServiceHandler {
	return &eventServiceHandler{
		dbhandler: databasehandler,
	}
}

// FindEventHandler search for events
func FindEventHandler(w http.ResponseWriter, r *http.Request) {
	http.Error(w, "This feature is coming soon", http.StatusNotImplemented)
}

// AllEventHandler gets all available events
func AllEventHandler(w http.ResponseWriter, r *http.Request) {
	http.Error(w, "This feature is coming soon", http.StatusNotImplemented)
}

// NewEventHandler adds a new event
func (eh *eventServiceHandler) NewEventHandler(w http.ResponseWriter, r *http.Request) {
	event := persistence.Event{}
	err := json.NewDecoder(r.Body).Decode(&event)
	if nil != err {
		respondWithError(w, "error occured while decoding event data", http.StatusBadRequest)
		return
	}

	err = eh.dbhandler.AddEvent(event)
	if err != nil{
		respondWithError(w, "error occured while persisting event", http.StatusInternalServerError)
		return
	}
	respondWithJSON(w, http.StatusCreated, event)
}

// UpdatEventHandler updates event instance
func UpdatEventHandler(w http.ResponseWriter, r *http.Request) {
	http.Error(w, "This feature is coming soon", http.StatusNotImplemented)
}

// DeletEventHandler deletes an event instance
func DeletEventHandler(w http.ResponseWriter, r *http.Request) {
	http.Error(w, "This feature is coming soon", http.StatusNotImplemented)
}

func respondWithError(w http.ResponseWriter, message string, code int) {
	respondWithJSON(w, code, map[string]string{"error": message})
}

func respondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}
