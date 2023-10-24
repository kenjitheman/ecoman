package api

import (
	"github.com/kenjitheman/ecoman/vars"
	"net/http"
	"testing"
	"time"
)

func TestApiResponseIsOk(t *testing.T) {
	resp, err := http.Get(vars.DataUrl)
	if err != nil {
		t.Errorf("Expected no error when retrieving test data from API, got: %v", err)
	}
	res := resp.StatusCode
	if res != 200 {
		t.Errorf("Expected response status code 200, got: %v", res)
	}
	defer resp.Body.Close()
}

func TestApiResponseBodyIsNotEmpty(t *testing.T) {
	resp, err := http.Get(vars.DataUrl)
	if err != nil {
		t.Errorf("Expected no error when retrieving test data from API, got: %v", err)
	}
	body := resp.Body
	if body == nil {
		t.Errorf("Expected non-empty response body, got nil")
	}
	defer resp.Body.Close()
}

func TestApiResponseBodyIsJson(t *testing.T) {
	resp, err := http.Get(vars.DataUrl)
	if err != nil {
		t.Errorf("Expected no error when retrieving test data from API, got: %v", err)
	}
	body := resp.Body
	if body == nil {
		t.Errorf("Expected non-empty response body, got nil")
	}
	defer resp.Body.Close()
}

func TestResponseTime(t *testing.T) {
	startTime := time.Now()
	_, err := http.Get(vars.DataUrl)
	if err != nil {
		t.Errorf("Expected no error, got: %v", err)
	}
	elapsed := time.Since(startTime)
	if elapsed > 5*time.Second {
		t.Errorf("Expected response time to be less than 5 seconds, got: %v", elapsed)
	}
}
