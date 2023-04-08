package api

import (
	"errors"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/arekbor/mistrzownia-radio-stats-api/types"
	"github.com/arekbor/mistrzownia-radio-stats-api/utils"
)

func (a *Api) handleApiHealth(w http.ResponseWriter, r *http.Request) {
	err := utils.WriteJSON(w, "It works!")
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		log.Println(err)
		return
	}
}

func (a *Api) handlePaginatedStats(w http.ResponseWriter, r *http.Request) {
	var (
		invalidPageError = errors.New("Invalid page or limit").Error()
	)

	if r.URL.Query().Get("page") == "" || r.URL.Query().Get("limit") == "" {
		http.Error(w, invalidPageError, http.StatusBadRequest)
		log.Println(invalidPageError)
		return
	}

	page, err := strconv.Atoi(r.URL.Query().Get("page"))
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		log.Println(err)
		return
	}

	limit, err := strconv.Atoi(r.URL.Query().Get("limit"))
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		log.Println(err)
		return
	}

	maxLimit, err := strconv.Atoi(os.Getenv("MAX_PAGINATION_LIMIT"))
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		log.Println(err)
		return
	}
	if page < 0 {
		page = 0
	}

	page++
	if limit > maxLimit || limit <= 0 {
		http.Error(w, invalidPageError, http.StatusBadRequest)
		log.Println(invalidPageError)
		return
	}

	stats, err := a.store.GetPaginatedStats(page, limit)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		log.Println(err)
		return
	}

	count, err := a.store.GetCountOfStats()
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		log.Println(err)
		return
	}

	responseStats := types.ResponseStats{
		Count: count,
		Data:  stats,
	}

	utils.WriteJSON(w, responseStats)
}
