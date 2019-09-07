package handler

import (
	"github.com/labstack/echo"
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"testing"
)

func TestHelloWorld(t *testing.T) {
	os.Setenv("DATASTORE_PROJECT_ID", "test-project")
	os.Setenv("DATASTORE_EMULATOR_HOST", "localhost:8059")

	e := echo.New()
	req, err := http.NewRequest("GET", "/", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()

	c := e.NewContext(req, rr)
	err = HelloWorld(c)
	if err != nil {
		t.Errorf("unexpeced handler reponse (err: %v)", err)
		return
	}

	if status := rr.Code; status != http.StatusOK {
		t.Errorf(
			"unexpected status: got (%v) want (%v)",
			status,
			http.StatusOK,
		)
	}

	expected := `{"id":12345,"content":"test"}`
	if strings.TrimSpace(rr.Body.String()) != expected {
		t.Errorf(
			"unexpected body: got (%v) want (%v)",
			rr.Body.String(),
			expected,
		)
	}
}

