package api

import (
	"log"
	"net/http"
	"os"
	"time"

	"github.com/arekbor/mistrzownia-radio-stats-api/store"
	"github.com/go-chi/cors"
	"github.com/gorilla/mux"
)

func (a *Api) prepareHandlers(r *mux.Router) {
	sub := r.PathPrefix("/api/v1").Subrouter()

	sub.HandleFunc("/health", a.handleApiHealth).Methods(http.MethodGet)
	sub.HandleFunc("/stats", a.handlePaginatedStats).Methods(http.MethodGet)
}

type Api struct {
	store        *store.Store
	addr         string
	writeTimeout time.Duration
	readTimeout  time.Duration
}

func New(
	addr string,
	store *store.Store,
	writeTimeout time.Duration,
	readTimeout time.Duration,
) *Api {
	return &Api{
		store:        store,
		addr:         addr,
		writeTimeout: writeTimeout,
		readTimeout:  readTimeout,
	}
}

func (a *Api) Init() error {
	r := mux.NewRouter()

	r.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{os.Getenv("CORS_URL")},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Content-Type", "Depth", "User-Agent", "X-File-Size", "X-Requested-With", "If-Modified-Since", "X-File-Name", "Cache-Control"},
		AllowCredentials: true,
	}))

	a.prepareHandlers(r)

	srv := &http.Server{
		Handler:      r,
		Addr:         a.addr,
		WriteTimeout: a.writeTimeout,
		ReadTimeout:  a.readTimeout,
	}

	log.Println("Server is running on addr: ", a.addr)

	return srv.ListenAndServe()
}
