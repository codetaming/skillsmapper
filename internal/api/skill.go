package api

import (
	"encoding/json"
	"fmt"
	"github.com/codetaming/skillsmapper/internal/model"
	"github.com/codetaming/skillsmapper/internal/persistence"
	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"net/http"
	"os"
	"time"
)

func (api *API) GetSkill(w http.ResponseWriter, r *http.Request) {
	v := mux.Vars(r)
	id := v["id"]
	s, err := api.dataStore.GetSkill(id)

	if err != nil {
		switch err.(type) {
		case *persistence.NotFoundError:
			http.Error(w, err.Error(), http.StatusNotFound)
			return
		default:
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}
	w.Header().Set("content-type", "application/json")
	json.NewEncoder(w).Encode(s)
}

func (api *API) GetSkills(w http.ResponseWriter, r *http.Request) {
	/*
		model.Skill, err = api.dataStore.GetAllSkills()
		if err != nil {
			http.Error(w, err.Error(), 500)
			return
		}

		w.Header().Set("content-type", "application/json")
		w.WriteHeader(http.StatusCreated)
		w.Header().Set("location", os.Getenv("BASE_URL")+"/skill/")
		json.NewEncoder(w).Encode(s)
	*/
}

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
	err = api.dataStore.PersistSkill(s)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	w.Header().Set("content-type", "application/json")
	w.WriteHeader(http.StatusCreated)
	w.Header().Set("location", fmt.Sprintf("%s/skill/%s", os.Getenv("BASE_URL"), s.SkillID))
	json.NewEncoder(w).Encode(s)
}
