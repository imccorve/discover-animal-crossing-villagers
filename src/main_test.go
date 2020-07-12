package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// func TestHandler(t *testing.T) {
// 	villagers, err := Handler(Request{})
// 	println(villagers.Body)
// 	assert.IsType(t, nil, err)
// 	assert.NotEqual(t, 0, len(villagers.Body))
// }

// TODO: Complete and Update unit tests
// func TestFilter(t *testing.T) {
// 	villagers, err := Handler(events.Request{"Bird"})
// 	println(villagers.Body)
// 	assert.IsType(t, nil, err)
// 	assert.NotEqual(t, 0, len(villagers.Body))
// }

func TestFormResponse(t *testing.T) {
	var m map[string]string
	m = make(map[string]string)
	m["Species"] = "Bird"
	response := formResponse(m)
	assert.NotEqual(t, 0, len(response.Body))
}
