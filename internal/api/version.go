package api

import (
	"encoding/json"
	"net/http"
)

func (api *API) GetVersion(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")
	json.NewEncoder(w).Encode("")
}
