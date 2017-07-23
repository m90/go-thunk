package thunk

import (
	"errors"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"
)

type recorder struct {
	result error
}

func (r *recorder) record(v error) {
	r.result = v
}

func TestHandleSafelyWith(t *testing.T) {
	tests := []struct {
		name                   string
		handler                http.Handler
		expectedStatusCode     int
		expectedCallbackResult error
	}{
		{
			"panic",
			http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				panic("zalgo")
			}),
			500,
			errors.New("zalgo"),
		},
		{
			"ok",
			http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				w.Write([]byte("OK!"))
			}),
			200,
			nil,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			rcd := recorder{}
			ts := httptest.NewServer(HandleSafelyWith(rcd.record)(test.handler))
			res, err := http.Get(ts.URL)
			if err != nil {
				t.Errorf("Unexpected error %v", err)
			}
			if res.StatusCode != test.expectedStatusCode {
				t.Errorf("Expected staus code of %v, got %v", test.expectedStatusCode, res.StatusCode)
			}
			if !reflect.DeepEqual(rcd.result, test.expectedCallbackResult) {
				t.Errorf("Expected callback result of %v, got %v", test.expectedCallbackResult, rcd.result)
			}
		})
	}
}

func TestHandleSafely(t *testing.T) {
	tests := []struct {
		name                   string
		handler                http.Handler
		expectedStatusCode     int
		expectedCallbackResult error
	}{
		{
			"panic",
			http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				panic("zalgo")
			}),
			500,
			errors.New("zalgo"),
		},
		{
			"ok",
			http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				w.Write([]byte("OK!"))
			}),
			200,
			nil,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			ts := httptest.NewServer(HandleSafely()(test.handler))
			res, err := http.Get(ts.URL)
			if err != nil {
				t.Errorf("Unexpected error %v", err)
			}
			if res.StatusCode != test.expectedStatusCode {
				t.Errorf("Expected staus code of %v, got %v", test.expectedStatusCode, res.StatusCode)
			}
		})
	}
}
