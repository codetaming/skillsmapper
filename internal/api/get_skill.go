package api

import (
	"encoding/json"
	"github.com/codetaming/skillsmapper/internal/persistence"
	"github.com/gorilla/mux"
	"net/http"
	"os"
)

func (api *API) GetSkill(w http.ResponseWriter, r *http.Request) {
	v := mux.Vars(r)
	id := v["id"]
	s, err := api.dataStore.GetSkill(id)

	if err != nil {
		switch err.(type) {
		case *persistence.NotFoundError:
			http.Error(w, err.Error(), http.StatusNotFound)
		default:
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}
	w.Header().Set("content-type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Header().Set("location", os.Getenv("BASE_URL")+"/skill/")
	json.NewEncoder(w).Encode(s)
}
