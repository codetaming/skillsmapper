package api

import (
	"github.com/codetaming/skillsmapper/internal/persistence"
	"github.com/gorilla/mux"
	"log"
)

type API struct {
	logger    *log.Logger
	dataStore persistence.DataStore
	Router    *mux.Router
}

func (api *API) setupRoutes() {
	api.Router.HandleFunc("/skill", api.SubmitSkill).Methods("POST")
	api.Router.HandleFunc("/skill/{id}", api.GetSkill).Methods("GET")
	api.Router.HandleFunc("/skill", api.GetSkills).Methods("GET")
	api.Router.HandleFunc("/version", api.GetVersion).Methods("GET")
}

func NewAPI(logger *log.Logger, dataStore persistence.DataStore) *API {
	api := &API{
		logger:    logger,
		dataStore: dataStore,
		Router:    mux.NewRouter(),
	}
	api.setupRoutes()
	return api
}
