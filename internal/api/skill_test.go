package api_test

import (
	"fmt"
	"github.com/codetaming/skillsmapper/internal/api"
	"github.com/codetaming/skillsmapper/internal/model"
	"github.com/codetaming/skillsmapper/internal/persistence/local"
	"github.com/stretchr/testify/assert"
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
	"time"
)

type testDef struct {
	name                   string
	in                     *http.Request
	out                    *httptest.ResponseRecorder
	expectedLocationHeader string
	expectedStatus         int
	expectedBody           string
}

type APITests struct {
	api            *api.API
	defaultSkillID string
	invalidSkillID string
}

func TestSkills(t *testing.T) {
	log := log.New(os.Stderr, "TEST : ", log.LstdFlags|log.Lmicroseconds|log.Lshortfile)

	dataStore, err := local.NewInMemoryDataStore(log)
	if err != nil {
		log.Fatalf("failed to create data store: %v", err)
	}
	a := api.NewAPI(log, dataStore)

	apiTests := APITests{
		api:            a,
		defaultSkillID: "default_skill_id",
		invalidSkillID: "invalid_skill_id",
	}

	dataStore.PersistSkill(model.Skill{
		SkillID: apiTests.defaultSkillID,
		Created: time.Time{},
		Email:   "dan@example.com",
		Tag:     "java",
		Level:   "using",
	})

	t.Run("SubmitSkill", apiTests.SubmitSkill)
	t.Run("RoundTripSkill", apiTests.RoundTripSkill)
	t.Run("GetSkill", apiTests.GetSkill)
	t.Run("GetSkills", apiTests.GetSkills)
}

func generateRequest(method string, target string, bodyFile string) *http.Request {
	body, err := os.Open(bodyFile)
	if err != nil {
		log.Fatalf("failed to open test file: %s: %v", bodyFile, err)
	}
	request := httptest.NewRequest(method, target, body)
	request.Header.Set("Content-Type", "application/json")
	return request
}

func (a *APITests) SubmitSkill(t *testing.T) {
	tests := []testDef{
		{
			name:                   "submit skill",
			in:                     generateRequest("POST", "/skill", "../../examples/skill.json"),
			out:                    httptest.NewRecorder(),
			expectedLocationHeader: "/skill/.+",
			expectedStatus:         http.StatusCreated,
			expectedBody:           "{\"skill_id\":.+,\"created\":.+,\"email\":\"dan@example.com\",\"tag\":\"java\",\"level\":\"learning\"}",
		},
	}
	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			a.api.Router.ServeHTTP(test.out, test.in)
			assert.Regexp(t, test.expectedLocationHeader, test.out.Header()["Location"])
			assert.Equal(t, test.expectedStatus, test.out.Code)
			assert.Regexp(t, test.expectedBody, test.out.Body)
		})
	}
}

func (a *APITests) RoundTripSkill(t *testing.T) {
	name := "round trip skill"
	t.Run(name, func(t *testing.T) {
		postOut := httptest.NewRecorder()
		a.api.Router.ServeHTTP(postOut, generateRequest("POST", "/skill", "../../examples/skill.json"))
		location := postOut.Header()["Location"]
		getOut := httptest.NewRecorder()
		a.api.Router.ServeHTTP(getOut, generateRequest("GET", location[0], "../../examples/empty.json"))
		assert.Equal(t, http.StatusOK, getOut.Code)
		expectedBody := "{\"skill_id\":.+,\"created\":.+,\"email\":\"dan@example.com\",\"tag\":\"java\",\"level\":\"learning\"}"
		assert.Regexp(t, expectedBody, getOut.Body)
	})
}

func (a *APITests) GetSkills(t *testing.T) {
	name := "get skills"
	t.Run(name, func(t *testing.T) {
		getOut := httptest.NewRecorder()
		a.api.Router.ServeHTTP(getOut, generateRequest("GET", "/skill", "../../examples/empty.json"))
		assert.Equal(t, http.StatusOK, getOut.Code)
		expectedBody := "{\"skill_id\":.+,\"created\":.+,\"email\":\"dan@example.com\",\"tag\":\"java\",\"level\":\"learning\"}"
		assert.Regexp(t, expectedBody, getOut.Body)
	})
}

func (a *APITests) GetSkill(t *testing.T) {
	tests := []testDef{
		{
			name:           "get default skill",
			in:             generateRequest("GET", fmt.Sprintf("/skill/%s", a.defaultSkillID), "../../examples/empty.json"),
			out:            httptest.NewRecorder(),
			expectedStatus: http.StatusOK,
			expectedBody:   "",
		},
		{
			name:           "get invalid skill",
			in:             generateRequest("GET", fmt.Sprintf("/skill/%s", a.invalidSkillID), "../../examples/empty.json"),
			out:            httptest.NewRecorder(),
			expectedStatus: http.StatusNotFound,
			expectedBody:   "",
		},
	}
	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			a.api.Router.ServeHTTP(test.out, test.in)
			assert.Equal(t, test.expectedStatus, test.out.Code)
			assert.Regexp(t, test.expectedBody, test.out.Body)
		})
	}
}
