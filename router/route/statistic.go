// Package route implements the route API endpoint.
package route

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/LuleaUniversityOfTechnology/2018-project-roaster/middleware"
	"github.com/LuleaUniversityOfTechnology/2018-project-roaster/model"
	"github.com/gorilla/mux"
	"github.com/willeponken/causerr"
)

func retrieveRoastCountTimeseries(w http.ResponseWriter, r *http.Request) (int, error) {
	start, err := time.Parse(time.RFC3339, r.URL.Query().Get("start"))
	if err != nil {
		return http.StatusBadRequest, causerr.New(
			errors.New("invalid ?start query parameter"),
			"Missing or invalid '?start=' query parameter with timestamp formatted according to RFC3339")
	}

	end, err := time.Parse(time.RFC3339, r.URL.Query().Get("end"))
	if err != nil {
		return http.StatusBadRequest, causerr.New(
			errors.New("invalid ?end query parameter"),
			"Missing or invalid '?end=' query parameter with timestamp formatted according to RFC3339")
	}

	interval, err := time.ParseDuration(fmt.Sprintf("%s",
		r.URL.Query().Get("interval")))
	if err != nil {
		return http.StatusBadRequest, causerr.New(
			errors.New("invalid ?interval query parameters"),
			"Missing '?end=' query parameter with a time duration formatted such as '300s', '1.5h' or '2h45m'. Valid time units are 'ns', 'us' (or 'µs'), 'ms', 's', 'm', 'h'")
	}

	vars := mux.Vars(r)
	username := vars["username"]

	var timeseries model.RoastCountTimeseries
	if username == "" {
		timeseries, err = model.GetGlobalRoastCountTimeseries(start, end, interval)
	} else {
		timeseries, err = model.GetUserRoastCountTimeseries(start, end, interval, username)
	}
	if err != nil {
		return http.StatusInternalServerError, causerr.New(err, "")
	}

	err = json.NewEncoder(w).Encode(timeseries)
	if err != nil {
		return http.StatusInternalServerError, causerr.New(err, "")
	}

	return http.StatusOK, nil
}

// Statistic adds the handlers for the Statistic [/statistic] endpoint.
func Statistic(r *mux.Router) {
	// All handlers are required to use application/json as their
	// Content-Type.
	r.Use(middleware.EnforceContentType("application/json"))

	// Global Specific Roast Count Timeseries [GET].
	r.Handle("/roast", handler(retrieveRoastCountTimeseries)).
		Queries("start", "",
			"end", "",
			"interval", "").
		Methods(http.MethodGet)

	// User Specific Roast Count Timeseries [GET].
	r.Handle("/{username}/roast", handler(retrieveRoastCountTimeseries)).
		Queries("start", "",
			"end", "",
			"interval", "").
		Methods(http.MethodGet)
}
