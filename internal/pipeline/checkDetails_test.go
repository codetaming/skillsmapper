package pipeline_test

import (
	"fmt"
	pipeline "github.com/codetaming/skillsmapper/internal/pipeline"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

const succeed = "\u2713"
const failed = "\u2717"

var exampleResponse = ``

func TestCheck(t *testing.T) {
	server := mockServer()
	defer server.Close()
	t.Log("Given the need to check pipeline.")
	{
		t.Logf("\tTest 0:\tWhen given checkDetails example.")
		raw, err := ioutil.ReadFile("../../examples/checkDetails.json")
		if err != nil {
			t.Fatalf("\t%s\tFailed to get example data: %v.", failed, err)
		}
		pipeline.Check(raw, server.URL)
		t.Logf("\t%s\tShould parse email.", succeed)
	}
}

func mockServer() *httptest.Server {
	f := func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(200)
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprintln(w, exampleResponse)
	}
	return httptest.NewServer(http.HandlerFunc(f))
}
