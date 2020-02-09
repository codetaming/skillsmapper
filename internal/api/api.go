package api

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"time"
)

type API struct {
	logger *log.Logger
}

func (api *API) SetupRoutes(r *mux.Router) {
	r.HandleFunc("/skill", api.Logger(api.SubmitSkill)).Methods("POST")
}

func (api *API) Logger(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		startTime := time.Now()
		defer api.logger.Printf("request processed in %s\n", time.Now().Sub(startTime))
		next(w, r)
	}
}

func NewAPI(logger *log.Logger) *API {
	return &API{
		logger: logger,
	}
}
