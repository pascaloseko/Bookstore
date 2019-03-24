package rest

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/pascaloseko/Events/persistence"
)


/*

{
  "name": "Nakuru",
  "address": "254-Nakuru",
  "country": "Kenya",
  "opentime": 7,
  "closetime": 20
}

*/

//NewLocationHandler Adds a new location
func NewLocationHandler(w http.ResponseWriter, r *http.Request) {
	loc := persistence.Location{}
	err := json.NewDecoder(r.Body).Decode(&loc)
	if err != nil {
		RespondWithError(w, "error occured while decoding location data", http.StatusBadRequest)
		return
	}
	if err := loc.AddLocation(); err != nil {
		fmt.Println(err)
		RespondWithError(w, "error occured while persisting location", http.StatusInternalServerError)
	}
	RespondWithJSON(w, http.StatusCreated, loc)
}
