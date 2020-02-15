package api_test

import (
	"fmt"
	"github.com/codetaming/skillsmapper/internal/api"
	"github.com/codetaming/skillsmapper/internal/model"
	"github.com/codetaming/skillsmapper/internal/persistence/local"
	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
	"time"
)

var a *api.API
var logger *log.Logger
var defaultSkillID = "default_skill_id"
var invalidSkillID = "invalid"

type testDef struct {
	name                   string
	in                     *http.Request
	out                    *httptest.ResponseRecorder
	expectedLocationHeader string
	expectedStatus         int
	expectedBody           string
}

func generateRequest(method string, target string, bodyFile string) *http.Request {
	body, err := os.Open(bodyFile)
	if err != nil {
		logger.Fatalf("failed to open test file: %s: %v", bodyFile, err)
	}
	request := httptest.NewRequest(method, target, body)
	request.Header.Set("Content-Type", "application/json")
	return request
}

func TestHandlers_SubmitSkill(t *testing.T) {
	tests := []testDef{
		{
			name:                   "Submit Skill",
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
			a.SubmitSkill(test.out, test.in)
			assert.Regexp(t, test.expectedLocationHeader, test.out.Header()["Location"])
			assert.Equal(t, test.expectedStatus, test.out.Code)
			assert.Regexp(t, test.expectedBody, test.out.Body)
		})
	}
}

func TestHandlers_GetSkill(t *testing.T) {
	tests := []testDef{
		{
			name:           "Get Skill",
			in:             generateRequest("GET", fmt.Sprintf("/skill/%s", defaultSkillID), "../../examples/empty.json"),
			out:            httptest.NewRecorder(),
			expectedStatus: http.StatusOK,
			expectedBody:   "",
		},
		{
			name:           "Get Invalid skill",
			in:             generateRequest("GET", fmt.Sprintf("/skill/%s", invalidSkillID), "../../examples/empty.json"),
			out:            httptest.NewRecorder(),
			expectedStatus: http.StatusNotFound,
			expectedBody:   "",
		},
	}
	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			a.GetSkill(test.out, test.in)
			assert.Equal(t, test.expectedStatus, test.out.Code)
			assert.Regexp(t, test.expectedBody, test.out.Body)
		})
	}
}

func init() {
	logger = log.New(os.Stdout, "skillsmapper-api-test ", log.LstdFlags|log.Lshortfile)
	dataStore, err := local.NewInMemoryDataStore(logger)
	if err != nil {
		logger.Fatalf("failed to create data store: %v", err)
	}
	a = api.NewAPI(logger, dataStore)
	a.SetupRoutes(mux.NewRouter())

	dataStore.PersistSkill(model.Skill{
		SkillID: defaultSkillID,
		Created: time.Time{},
		Email:   "dan@example.com",
		Tag:     "java",
		Level:   "using",
	})
}
