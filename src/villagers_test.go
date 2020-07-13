package main

import (
	"io/ioutil"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFilterVillagersPersonality(t *testing.T) {
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
	lazy_villagers := [2]Villager{villagers[0], villagers[2]}
	var params map[string]string
	params = make(map[string]string)
	params["Personality"] = "Lazy"
	assert.ElementsMatch(t, lazy_villagers, filterVillagers(villagers, params))
}

func TestFilterVillagersSpecies(t *testing.T) {
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
	cat_villagers := [1]Villager{villagers[2]}
	var params map[string]string
	params = make(map[string]string)
	params["Species"] = "Cat"
	assert.ElementsMatch(t, cat_villagers, filterVillagers(villagers, params))
}

func TestFilterVillagersGender(t *testing.T) {
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
	male_villagers := [1]Villager{villagers[0]}
	var params map[string]string
	params = make(map[string]string)
	params["Gender"] = "Male"
	assert.ElementsMatch(t, male_villagers, filterVillagers(villagers, params))
}

func TestConvertToVillagers(t *testing.T) {
	actual, err := convertToVillagers("")
	assert.Equal(t, []Villager{}, actual)
	assert.Equal(t, nil, err)
}

func TestConvertHTMLStringToArray(t *testing.T) {
	content, err := ioutil.ReadFile("./test_input/html_test.html")
	actual, err := convertHTMLStringToArray(string(content))
	assert.NotEqual(t, 0, len(actual))
	assert.Equal(t, nil, err)
}
