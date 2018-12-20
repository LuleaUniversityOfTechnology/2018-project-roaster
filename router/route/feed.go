// Package route implements the feed API endpoint.
package route

import (
	"encoding/json"
	"errors"
	"net/http"
	"strconv"
	"strings"

	"github.com/LuleaUniversityOfTechnology/2018-project-roaster/middleware"
	"github.com/LuleaUniversityOfTechnology/2018-project-roaster/model"
	"github.com/gorilla/mux"
	"github.com/willeponken/causerr"
)

func getBooleanFromString(s string) bool {
	if strings.ToLower(s) == "true" {
		return true
	}

	return false
}

func retrieveFeed(w http.ResponseWriter, r *http.Request) (int, error) {
	pageStr := r.URL.Query().Get("page")
	if pageStr == "" {
		return http.StatusBadRequest, causerr.New(
			errors.New("invalid query parameters"),
			"Must provide '?page=' query parameter with a single number value")
	}

	page, err := strconv.ParseUint(pageStr, 10, 64)
	if err != nil {
		// Internal server error because the router should have
		// made sure the `page` query is a number.
		return http.StatusInternalServerError, causerr.New(err, "")
	}

	username := r.URL.Query().Get("user")
	friends := getBooleanFromString(r.URL.Query().Get("friends"))

	if username != "" && friends {
		return http.StatusBadRequest, causerr.New(
			errors.New("request for friends is missing user query parameter"),
			"Friends query parameter also requires the user query parameter")
	}

	var feed model.Feed
	if username == "" {
		feed, err = model.GetGlobalFeed(page)
	} else {
		feed, err = model.GetUserFeed(username, friends, page)
	}
	if err != nil {
		return http.StatusInternalServerError, causerr.New(err, "")
	}

	err = json.NewEncoder(w).Encode(feed)
	if err != nil {
		return http.StatusInternalServerError, causerr.New(err, "")
	}

	return http.StatusOK, nil
}

// Feed adds the handlers for the Feed [/feed] endpoint.
func Feed(r *mux.Router) {
	// All handlers are required to use application/json as their
	// Content-Type.
	r.Use(middleware.EnforceContentType("application/json"))

	// Global Feed [GET].
	r.Handle("", handler(retrieveFeed)).
		Queries("page", "{[0-9]+}").
		Methods(http.MethodGet)
}
