package main

import (
	"testing"

	"github.com/aws/aws-lambda-go/events"

	"github.com/stretchr/testify/assert"
)

// func TestHandler(t *testing.T) {
// 	villagers, err := Handler(Request{})
// 	println(villagers.Body)
// 	assert.IsType(t, nil, err)
// 	assert.NotEqual(t, 0, len(villagers.Body))
// }

// TODO: Complete and Update unit tests
func TestFilter(t *testing.T) {
	villagers, err := Handler(events.Request{"Bird"})
	println(villagers.Body)
	assert.IsType(t, nil, err)
	assert.NotEqual(t, 0, len(villagers.Body))
}
