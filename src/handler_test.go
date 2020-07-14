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

func TestPaginateResultsNoParams(t *testing.T) {
	var villagers = []Villager{
		Villager{
			ImageURL:     "",
			Name:         "test1",
			JapaneseName: "test1_japanese",
			Species:      "Dog",
			Gender:       "Male",
			Personality:  "Lazy",
		},
		Villager{
			ImageURL:     "",
			Name:         "test2",
			JapaneseName: "test2_japanese",
			Species:      "Dog",
			Gender:       "Female",
			Personality:  "Snooty",
		},
		Villager{
			ImageURL:     "",
			Name:         "test3",
			JapaneseName: "test3_japanese",
			Species:      "Cat",
			Gender:       "Female",
			Personality:  "Lazy",
		},
	}
	var params map[string]string
	params = make(map[string]string)
	actual, err := paginateResults(params, villagers)
	expected := len(villagers)
	assert.Equal(t, expected, len(actual))
	assert.Equal(t, nil, err)
}

func TestPaginateResultsNoOffset(t *testing.T) {
	var villagers = []Villager{
		Villager{
			ImageURL:     "",
			Name:         "test1",
			JapaneseName: "test1_japanese",
			Species:      "Dog",
			Gender:       "Male",
			Personality:  "Lazy",
		},
		Villager{
			ImageURL:     "",
			Name:         "test2",
			JapaneseName: "test2_japanese",
			Species:      "Dog",
			Gender:       "Female",
			Personality:  "Snooty",
		},
		Villager{
			ImageURL:     "",
			Name:         "test3",
			JapaneseName: "test3_japanese",
			Species:      "Cat",
			Gender:       "Female",
			Personality:  "Lazy",
		},
	}
	var params map[string]string
	params = make(map[string]string)
	params["limit"] = "2"
	actual, err := paginateResults(params, villagers)
	expected := 2
	assert.Equal(t, expected, len(actual))
	assert.Equal(t, nil, err)
}

func TestPaginateResultsWithOffset(t *testing.T) {
	var villagers = []Villager{
		Villager{
			ImageURL:     "",
			Name:         "test1",
			JapaneseName: "test1_japanese",
			Species:      "Dog",
			Gender:       "Male",
			Personality:  "Lazy",
		},
		Villager{
			ImageURL:     "",
			Name:         "test2",
			JapaneseName: "test2_japanese",
			Species:      "Dog",
			Gender:       "Female",
			Personality:  "Snooty",
		},
		Villager{
			ImageURL:     "",
			Name:         "test3",
			JapaneseName: "test3_japanese",
			Species:      "Cat",
			Gender:       "Female",
			Personality:  "Lazy",
		},
	}

	var params map[string]string
	params = make(map[string]string)
	params["limit"] = "1"
	params["offset"] = "1"
	actual, err := paginateResults(params, villagers)
	expected := 1
	assert.Equal(t, expected, len(actual))
	assert.Equal(t, nil, err)
}

func TestPaginateResultsWithLargeLimit(t *testing.T) {
	var villagers = []Villager{
		Villager{
			ImageURL:     "",
			Name:         "test1",
			JapaneseName: "test1_japanese",
			Species:      "Dog",
			Gender:       "Male",
			Personality:  "Lazy",
		},
		Villager{
			ImageURL:     "",
			Name:         "test2",
			JapaneseName: "test2_japanese",
			Species:      "Dog",
			Gender:       "Female",
			Personality:  "Snooty",
		},
		Villager{
			ImageURL:     "",
			Name:         "test3",
			JapaneseName: "test3_japanese",
			Species:      "Cat",
			Gender:       "Female",
			Personality:  "Lazy",
		},
	}

	var params map[string]string
	params = make(map[string]string)
	params["limit"] = "10000"
	actual, err := paginateResults(params, villagers)
	expected := len(villagers)
	assert.Equal(t, expected, len(actual))
	assert.Equal(t, nil, err)
}

func TestPaginateResultsWithInvalidValues(t *testing.T) {
	var villagers = []Villager{
		Villager{
			ImageURL:     "",
			Name:         "test1",
			JapaneseName: "test1_japanese",
			Species:      "Dog",
			Gender:       "Male",
			Personality:  "Lazy",
		},
		Villager{
			ImageURL:     "",
			Name:         "test2",
			JapaneseName: "test2_japanese",
			Species:      "Dog",
			Gender:       "Female",
			Personality:  "Snooty",
		},
		Villager{
			ImageURL:     "",
			Name:         "test3",
			JapaneseName: "test3_japanese",
			Species:      "Cat",
			Gender:       "Female",
			Personality:  "Lazy",
		},
	}

	var params map[string]string
	params = make(map[string]string)
	params["limit"] = "30"
	params["offset"] = "30"
	actual, err := paginateResults(params, villagers)
	expected := 0
	assert.Equal(t, expected, len(actual))
	assert.Equal(t, nil, err)
}

func TestPaginateResultsWithLimitGreaterThanLength(t *testing.T) {
	var villagers = []Villager{
		Villager{
			ImageURL:     "",
			Name:         "test1",
			JapaneseName: "test1_japanese",
			Species:      "Dog",
			Gender:       "Male",
			Personality:  "Lazy",
		},
		Villager{
			ImageURL:     "",
			Name:         "test2",
			JapaneseName: "test2_japanese",
			Species:      "Dog",
			Gender:       "Female",
			Personality:  "Snooty",
		},
		Villager{
			ImageURL:     "",
			Name:         "test3",
			JapaneseName: "test3_japanese",
			Species:      "Cat",
			Gender:       "Female",
			Personality:  "Lazy",
		},
	}

	var params map[string]string
	params = make(map[string]string)
	params["limit"] = "2"
	params["offset"] = "1"
	actual, err := paginateResults(params, villagers)
	expected := 1
	assert.Equal(t, expected, len(actual))
	assert.Equal(t, nil, err)
}

func TestPaginateResults(t *testing.T) {
	var villagers = []Villager{
		Villager{
			ImageURL:     "",
			Name:         "test1",
			JapaneseName: "test1_japanese",
			Species:      "Dog",
			Gender:       "Male",
			Personality:  "Lazy",
		},
		Villager{
			ImageURL:     "",
			Name:         "test2",
			JapaneseName: "test2_japanese",
			Species:      "Dog",
			Gender:       "Female",
			Personality:  "Snooty",
		},
		Villager{
			ImageURL:     "",
			Name:         "test3",
			JapaneseName: "test3_japanese",
			Species:      "Cat",
			Gender:       "Female",
			Personality:  "Lazy",
		},
	}
	var params map[string]string
	params = make(map[string]string)
	params["limit"] = "1"
	params["offset"] = "0"
	actual, err := paginateResults(params, villagers)
	expected := 1
	assert.Equal(t, expected, len(actual))
	assert.Equal(t, nil, err)
}

// TODO: Improve testing so that runtime doesn't play a factor in
// PASSING or FAILING.
// func TestFormResponseWithParam(t *testing.T) {
// 	var m map[string]string
// 	m = make(map[string]string)
// 	m["Species"] = "Bird"
// 	response := formResponse(m)
// 	assert.NotEqual(t, 0, len(response.Body))
// 	assert.Equal(t, 200, response.StatusCode)
// }

// func TestFormResponse(t *testing.T) {
// 	var m map[string]string
// 	m = make(map[string]string)
// 	response := formResponse(m)
// 	assert.NotEqual(t, 0, len(response.Body))
// 	assert.Equal(t, 200, response.StatusCode)
// }
