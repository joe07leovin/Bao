package httpapi

import (
	"context"
	"errors"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"
)

type readinessFunc func(context.Context) error

func (f readinessFunc) Execute(ctx context.Context) error { return f(ctx) }
func TestHealthResponses(t *testing.T) {
	for _, tc := range []struct {
		path       string
		dependency error
		want       int
	}{
		{"/health/live", errors.New("secret database details"), http.StatusNoContent},
		{"/health/ready", nil, http.StatusNoContent},
		{"/health/ready", errors.New("secret database details"), http.StatusServiceUnavailable},
	} {
		t.Run(tc.path+http.StatusText(tc.want), func(t *testing.T) {
			called := false
			handler := Handler(readinessFunc(func(ctx context.Context) error {
				called = true
				if _, ok := ctx.Deadline(); !ok {
					t.Fatal("missing deadline")
				}
				return tc.dependency
			}), time.Second)
			response := httptest.NewRecorder()
			handler.ServeHTTP(response, httptest.NewRequest(http.MethodGet, tc.path, nil))
			if response.Code != tc.want {
				t.Fatalf("got %d want %d", response.Code, tc.want)
			}
			if strings.Contains(response.Body.String(), "secret") {
				t.Fatal("dependency details leaked")
			}
			if tc.path == "/health/live" && called {
				t.Fatal("liveness depends on database")
			}
		})
	}
}
func TestReadinessTimeout(t *testing.T) {
	handler := Handler(readinessFunc(func(ctx context.Context) error { <-ctx.Done(); return ctx.Err() }), time.Millisecond)
	response := httptest.NewRecorder()
	handler.ServeHTTP(response, httptest.NewRequest(http.MethodGet, "/health/ready", nil))
	if response.Code != http.StatusServiceUnavailable {
		t.Fatalf("got %d", response.Code)
	}
}
