package engine

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/99designs/gqlgen/handler"

	"github.com/stretchr/testify/assert"
)

func TestExecutionEngine(t *testing.T) {

}

func TestEngineWithHandler(t *testing.T) {
	h := handler.GraphQL(NewEngine(`type Schema { query: Query } type Query { name: String }`))

	t.Run("success", func(t *testing.T) {
		resp := doRequest(h, "POST", "/graphql", `{"query":"{ name }"}`)
		assert.Equal(t, http.StatusOK, resp.Code)
		assert.Equal(t, `{"data":{}}`, resp.Body.String())
	})
}

func doRequest(handler http.Handler, method string, target string, body string) *httptest.ResponseRecorder {
	r := httptest.NewRequest(method, target, strings.NewReader(body))
	w := httptest.NewRecorder()

	handler.ServeHTTP(w, r)
	return w
}
