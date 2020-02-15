package api

import (
	"github.com/codetaming/skillsmapper/internal/persistence"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"time"
)

type API struct {
	logger    *log.Logger
	dataStore persistence.DataStore
}

func (api *API) SetupRoutes(r *mux.Router) {
	r.HandleFunc("/skill", api.Logger(api.SubmitSkill)).Methods("POST")
	r.HandleFunc("/skill/{id}", api.GetSkill).Methods("GET")
}

func (api *API) Logger(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		startTime := time.Now()
		defer api.logger.Printf("request processed in %s\n", time.Now().Sub(startTime))
		next(w, r)
	}
}

func NewAPI(logger *log.Logger, dataStore persistence.DataStore) *API {
	return &API{
		logger:    logger,
		dataStore: dataStore,
	}
}
