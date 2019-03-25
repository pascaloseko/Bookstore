package rest

import (
	"encoding/json"
	"fmt"
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

// // FindEventHandler search for events
// func FindEventHandler(w http.ResponseWriter, r *http.Request) {
// 	http.Error(w, "This feature is coming soon", http.StatusNotImplemented)
// }

// AllEventHandler gets all available events
func AllEventHandler(w http.ResponseWriter, r *http.Request) {

	fmt.Println(r.Context().Value("user"))
	e, err := persistence.GetEvents()
	if err != nil {
		RespondWithError(w, "error occured while getting locations", http.StatusInternalServerError)
	}
	RespondWithJSON(w, http.StatusOK, e)
}

/*

 {
	"name": "Kapuka throwback",
	"duration": 9,
	"startdate": "2018-01-04",
	"enddate": "2018-01-05",
	"location": 1
}

*/

// NewEventHandler adds a new event
func NewEventHandler(w http.ResponseWriter, r *http.Request) {
	event := persistence.Event{}
	err := json.NewDecoder(r.Body).Decode(&event)
	if nil != err {
		fmt.Println(err)
		RespondWithError(w, "error occured while decoding event data", http.StatusBadRequest)
		return
	}
	if err := event.AddEvent(); err != nil {
		fmt.Println(err)
		RespondWithError(w, "error occured while persisting event", http.StatusInternalServerError)
	}
	RespondWithJSON(w, http.StatusCreated, event)
}

// UpdatEventHandler updates event instance
func UpdatEventHandler(w http.ResponseWriter, r *http.Request) {
	http.Error(w, "This feature is coming soon", http.StatusNotImplemented)
}

// DeletEventHandler deletes an event instance
func DeletEventHandler(w http.ResponseWriter, r *http.Request) {
	http.Error(w, "This feature is coming soon", http.StatusNotImplemented)
}

//RespondWithError proper error handling with status code
func RespondWithError(w http.ResponseWriter, message string, code int) {
	RespondWithJSON(w, code, map[string]string{"error": message})
}

//RespondWithJSON proper json handling showing the status code and the response
func RespondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}
