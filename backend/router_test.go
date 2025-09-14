package main

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/mbient/todo-api/routers"
	"github.com/stretchr/testify/assert"
)

func TestSomething(t *testing.T) {
	assert.True(t, true, "True is true!")
}

func TestTaskRoute(t *testing.T) {
	router := routers.TaskRouter()
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/api/v1/tasks", nil)
	router.ServeHTTP(w, req)
	assert.Equal(t, 200, w.Code)
}

func TestTaskRouteById(t *testing.T) {
	router := routers.TaskRouter()
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/api/v1/tasks/1", nil)
	router.ServeHTTP(w, req)
	assert.Equal(t, 200, w.Code)
}

func TestWrongRoute(t *testing.T) {
	router := routers.TaskRouter()
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/api/v1/nope", nil)
	router.ServeHTTP(w, req)
	assert.Equal(t, 404, w.Code)
}
