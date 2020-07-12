package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestExtractPageHTML(t *testing.T) {
	url := "https://nookipedia.com/w/api.php?action=parse&section=3&origin=*&prop=text&page=List_of_villagers&format=json"
	HTMLString, err := extractPageHTML(url)
	assert.NotEqual(t, 0, len(HTMLString))
	assert.Equal(t, nil, err)
}
func TestFormResponseWithParam(t *testing.T) {
	var m map[string]string
	m = make(map[string]string)
	m["Species"] = "Bird"
	response := formResponse(m)
	assert.NotEqual(t, 0, len(response.Body))
	assert.Equal(t, 200, response.StatusCode)
}

func TestFormResponse(t *testing.T) {
	var m map[string]string
	m = make(map[string]string)
	response := formResponse(m)
	assert.NotEqual(t, 0, len(response.Body))
	assert.Equal(t, 200, response.StatusCode)
}
