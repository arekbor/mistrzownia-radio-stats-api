package api

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/arekbor/mistrzownia-radio-stats-api/utils"
)

func (a *Api) handleApiHealth(w http.ResponseWriter, r *http.Request) {
	utils.WriteJSON(w, "It works!")
}

func (a *Api) handlePaginatedStats(w http.ResponseWriter, r *http.Request) {
	if r.URL.Query().Get("page") == "" {
		http.Error(w, errors.New("Invalid page number").Error(), http.StatusBadRequest)
		return
	}

	page, err := strconv.Atoi(r.URL.Query().Get("page"))
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	if page < 0 {
		http.Error(w, errors.New("Invalid page number").Error(), http.StatusBadRequest)
		return
	}

	stats, err := a.store.GetPaginatedStats(page, 5)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	utils.WriteJSON(w, stats)
}
