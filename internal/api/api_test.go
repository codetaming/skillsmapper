package api_test

import (
	"github.com/codetaming/skillsmapper/internal/api"
	"github.com/stretchr/testify/assert"
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
)

var a *api.API
var logger *log.Logger

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
			expectedBody:           "{\"SkillID\":.+,\"Created\":.+,\"Email\":\"dan@example.com\",\"Tag\":\"java\",\"Level\":\"learning\"}",
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
