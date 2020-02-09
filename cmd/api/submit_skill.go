package api

import (
	"encoding/json"
	"github.com/codetaming/skillsmapper/internal/model"
	"github.com/google/uuid"
	"net/http"
	"os"
	"time"
)

func (api *API) SubmitSkill(w http.ResponseWriter, r *http.Request) {
	s := model.Skill{
		SkillID: uuid.Must(uuid.NewUUID()).String(),
		Created: time.Now(),
	}

	err := json.NewDecoder(r.Body).Decode(&s)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("content-type", "application/json")
	w.WriteHeader(http.StatusCreated)
	w.Header().Set("location", os.Getenv("BASE_URL")+"/skill/")
	json.NewEncoder(w).Encode(s)
}
