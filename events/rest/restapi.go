package rest

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/pascaloseko/Events/persistence"
)

// ServeAPI routes for the event service
func ServeAPI(endpoint string, databasehandler persistence.DatabaseHandler) error {
	// handler := New(databasehandler)
	r := mux.NewRouter()
	eventsrouter := r.PathPrefix("/events").Subrouter()
	// eventsrouter.Methods("GET").Path("/{SearchCriteria}/{search}").HandlerFunc(FindEventHandler)
	eventsrouter.Methods("GET").Path("").HandlerFunc(AllEventHandler)
	eventsrouter.Methods("POST").Path("").HandlerFunc(NewEventHandler)
	eventsrouter.Methods("PUT").Path("/{id}").HandlerFunc(UpdatEventHandler)
	eventsrouter.Methods("DELETE").Path("/{id}").HandlerFunc(DeletEventHandler)

	eventsrouter.Methods("POST").Path("/location").HandlerFunc(NewLocationHandler)
	eventsrouter.Methods("GET").Path("/locations").HandlerFunc(AllLocationsHandler)
	eventsrouter.Methods("GET").Path("/location/{id}").HandlerFunc(LocationHandler)

	eventsrouter.Use(JwtAuth)

	err := http.ListenAndServe(endpoint, r)
	if err != nil {
		log.Fatalf("cannot load server: %v", err)
	}
	return err
}
