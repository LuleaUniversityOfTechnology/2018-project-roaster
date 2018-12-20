// Package route implements the roast API endpoint.
package route

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/LuleaUniversityOfTechnology/2018-project-roaster/analyze"
	"github.com/LuleaUniversityOfTechnology/2018-project-roaster/middleware"
	"github.com/LuleaUniversityOfTechnology/2018-project-roaster/model"
	"github.com/LuleaUniversityOfTechnology/2018-project-roaster/session"
	"github.com/gorilla/mux"
	"github.com/willeponken/causerr"
)

type snippet struct {
	Language string `json:"language"`
	Code     string `json:"code"`
}

func analyzeCode(w http.ResponseWriter, r *http.Request) (int, error) {
	var in snippet

	s, err := session.Get(r, "roaster_auth")
	if err != nil {
		return http.StatusInternalServerError, causerr.New(err, "")
	}

	// Empty user is handled as an anonymous user (not logged in).
	username, _ := s.Values["username"].(string)

	err = json.NewDecoder(r.Body).Decode(&in)
	if err != nil {
		return http.StatusBadRequest,
			causerr.New(err, "Invalid Roast data sent")
	}

	var roast *model.RoastResult

	switch in.Language {
	case "python3":
		roast, err = analyze.WithFlake8(username, in.Code)
		if err != nil {
			return http.StatusInternalServerError, causerr.New(err, "")
		}
	default:
		return http.StatusBadRequest,
			causerr.New(err,
				fmt.Sprintf("Language: '%s' is not supported", in.Language))
	}

	if username != "" {
		err = model.PutRoast(roast)
		if err != nil {
			return http.StatusInternalServerError, causerr.New(
				fmt.Errorf("failed to insert Roast for username: '%s': %v", username, err), "")
		}
	}

	err = json.NewEncoder(w).Encode(roast)
	if err != nil {
		return http.StatusInternalServerError, causerr.New(err, "")
	}

	return http.StatusOK, nil
}

// Roast adds the handlers for the Roast [/roast] endpoint.
func Roast(r *mux.Router) {
	// All handlers are required to use application/json as their
	// Content-Type.
	r.Use(middleware.EnforceContentType("application/json"))

	// Analyze code snippet [POST].
	r.Handle("", handler(analyzeCode)).Methods(http.MethodPost)
}
