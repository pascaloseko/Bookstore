package rest

import (
	"bytes"
	"context"
	"encoding/json"
	"log"
	"net/http"
	"strings"

	"github.com/pascaloseko/Events/persistence"
)

// JwtAuth controls the authentication of this app to SSO
var JwtAuth = func(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		ssoURL := "http://0.0.0.0:8080/sso/verify"

		rejectURL := []string{"/sso/register", "/sso/login", "/sso/verify"}
		currentPath := r.URL.Path
		w.Header().Add("Content-Type", "application/json")
		for _, value := range rejectURL {
			if value == currentPath {
				next.ServeHTTP(w, r)
				return
			}
		}

		response := make(map[string]interface{})
		tokenHeader := r.Header.Get("Authorization")

		if tokenHeader == "" {
			response = map[string]interface{}{"error": "Missing Token"}
			RespondWithJSON(w, http.StatusBadRequest, response)
			return
		}

		splitted := strings.Split(tokenHeader, " ")

		if len(splitted) != 2 {
			response = map[string]interface{}{"error": "Invalid/Malformed Token"}
			RespondWithJSON(w, http.StatusBadRequest, response)
			return
		}

		theToken := splitted[1]
		payload := "{\"token\":\"" + theToken + "\"}"
		jsonPayload := []byte(payload)
		responseSSO, err := http.Post(ssoURL, "application/json", bytes.NewBuffer(jsonPayload))

		if err != nil {
			log.Fatalln(err)
		}
		if responseSSO.Status != "200 OK" {
			response = map[string]interface{}{"error": "Auth Error"}
			RespondWithJSON(w, http.StatusBadRequest, response)
			return
		}
		user := &persistence.User{}
		json.NewDecoder(responseSSO.Body).Decode(user)
		ctx := context.WithValue(r.Context(), "user", user.UserID)
		r = r.WithContext(ctx)
		next.ServeHTTP(w, r)
	})
}
