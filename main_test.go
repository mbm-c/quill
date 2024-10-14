package main

import (
	"errors"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewApp(t *testing.T) {
	// Mock function that simulates app.Listen without blocking
	mockListenFunc := func(port string) error {
		t.Log("Listening on port", port)
		return nil
	}

	// Test when PORT is not set
	t.Run("Port not set", func(t *testing.T) {
		err := NewApp("", mockListenFunc)
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "$PORT must be set")
	})

	// Test when PORT is set and Listen works
	t.Run("Port is set and listen successful", func(t *testing.T) {
		// Mock the port environment variable
		t.Setenv("PORT", "8080")
		defer os.Unsetenv("PORT")

		port := os.Getenv("PORT")
		err := NewApp(port, mockListenFunc)
		assert.NoError(t, err)
	})

	// Test when Listen returns an error
	t.Run("Listen fails", func(t *testing.T) {
		// Mock function that simulates app.Listen returning an error
		mockErrorListenFunc := func(port string) error {
			t.Log("Listening on port", port)
			return errors.New("failed to start server")
		}

		t.Setenv("PORT", "8080")
		defer os.Unsetenv("PORT")

		port := os.Getenv("PORT")
		err := NewApp(port, mockErrorListenFunc)
		assert.Error(t, err)
		assert.Equal(t, "failed to start server", err.Error())
	})
}
