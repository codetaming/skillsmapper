package api

import (
	"encoding/json"
	"github.com/codetaming/skillsmapper/internal/model"
	"github.com/codetaming/skillsmapper/internal/version"
	"net/http"
)

func (api *API) GetVersion(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")
	v := model.Version{
		Version: version.Version,
	}
	json.NewEncoder(w).Encode(v)
}
