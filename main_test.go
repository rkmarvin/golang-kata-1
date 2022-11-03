package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestWelcomeMessage(t *testing.T) {
	// given
	expected := "Hello world!"

	// than
	actual := welcomeMessage()

	// that
	assert.Equal(t, expected, actual, "message is not equal")
}
