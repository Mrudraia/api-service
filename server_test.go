package main

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHelloWorldShouldSucceed(t *testing.T) {
	testServer := httptest.NewServer(http.HandlerFunc(handleHelloWorld))
	defer testServer.Close()

	testClient := testServer.Client()
	fmt.Println(testServer.URL)
	response, err := testClient.Get(testServer.URL)
	if err != nil {
		t.Error(err)
	}
	body, err := io.ReadAll(response.Body)
	if err != nil {
		t.Error(err)
	}

	assert.Equal(t, 200, response.StatusCode)
	assert.Equal(t, "Hello World!", string(body))
}

func TestHelloWorldShouldFail(t *testing.T) {
	testServer := httptest.NewServer(http.HandlerFunc(handleHelloWorld))
	defer testServer.Close()

	testClient := testServer.Client()

	body := strings.NewReader("some body")
	response, err := testClient.Post(testServer.URL, "application/json", body)
	if err != nil {
		t.Error(err)
	}

	assert.Equal(t, 405, response.StatusCode)
}

func TestHealthShouldSucceed(t *testing.T) {
	testServer := httptest.NewServer(http.HandlerFunc(handleHealth))
	defer testServer.Close()

	testClient := testServer.Client()
	fmt.Println(testServer.URL)
	response, err := testClient.Get(testServer.URL)
	if err != nil {
		t.Error(err)
	}
	body, err := io.ReadAll(response.Body)
	if err != nil {
		t.Error(err)
	}

	assert.Equal(t, 200, response.StatusCode)
	assert.Equal(t, "ok", string(body))
}

func TestHealthShouldFail(t *testing.T) {
	testServer := httptest.NewServer(http.HandlerFunc(handleHealth))
	defer testServer.Close()

	testClient := testServer.Client()

	body := strings.NewReader("some body")
	response, err := testClient.Post(testServer.URL, "application/json", body)
	if err != nil {
		t.Error(err)
	}

	assert.Equal(t, 405, response.StatusCode)
}
