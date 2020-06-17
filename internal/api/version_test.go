package api_test

import (
	"fmt"
	"github.com/codetaming/skillsmapper/internal/api"
	"github.com/stretchr/testify/assert"
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
)

type APIVersionTests struct {
	api *api.API
}

func TestVersion(t *testing.T) {
	log := log.New(os.Stderr, "TEST : ", log.LstdFlags|log.Lmicroseconds|log.Lshortfile)
	a := api.NewAPI(log, nil)

	apiVersionTests := APIVersionTests{
		api: a,
	}

	t.Run("GetVersion", apiVersionTests.GetVersion)
}

func (a *APIVersionTests) GetVersion(t *testing.T) {
	tests := []testDef{
		{
			name:           "get default skill",
			in:             generateRequest("GET", fmt.Sprintf("/version"), "../../examples/empty.json"),
			out:            httptest.NewRecorder(),
			expectedStatus: http.StatusOK,
			expectedBody:   "",
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			a.api.Router.ServeHTTP(test.out, test.in)
			assert.Equal(t, test.expectedStatus, test.out.Code)
			assert.Regexp(t, test.expectedBody, test.out.Body)
		})
	}
}
