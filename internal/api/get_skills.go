package api

import (
	"net/http"
)

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
