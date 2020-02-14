package api

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
	"os"
)

func (api *API) GetSkill(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	skillID := vars["skillID"]
	s, err := api.dataStore.GetSkill(skillID)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	w.Header().Set("content-type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Header().Set("location", os.Getenv("BASE_URL")+"/skill/")
	json.NewEncoder(w).Encode(s)
}
