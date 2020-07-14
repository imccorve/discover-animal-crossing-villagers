package main

import (
	"strings"

	"github.com/PuerkitoBio/goquery"
)

// Villager struct that holds the information of each villager.
type Villager struct {
	ImageURL     string `json:"imageUrl"`
	Name         string `json:"name"`
	JapaneseName string `json:"japaneseName"`
	Species      string `json:"species"`
	Gender       string `json:"gender"`
	Personality  string `json:"personality"`
}

func filterVillagers(villagers []Villager, params map[string]string) []Villager {
	var filteredVillagers []Villager
	for _, villager := range villagers {
		keep := true
		if val, found := params["Species"]; found && val != villager.Species {
			keep = false
		}
		if val, found := params["Gender"]; found && val != villager.Gender {
			keep = false
		}
		if val, found := params["Personality"]; found && val != villager.Personality {
			keep = false
		}
		if keep {
			filteredVillagers = append(filteredVillagers, villager)
		}
	}
	return filteredVillagers
}

func convertHTMLStringToArray(HTMLString string) ([][]string, error) {
	htmlReader := strings.NewReader(HTMLString)

	var row []string
	var rows [][]string

	doc, err := goquery.NewDocumentFromReader(htmlReader)
	if err != nil {
		return nil, err
	}
	doc.Find("tbody").Each(func(index int, tbodyhtml *goquery.Selection) {
		tbodyhtml.Find("tr").Each(func(indextr int, rowhtml *goquery.Selection) {
			rowhtml.Find("th").Each(func(indexth int, tableheading *goquery.Selection) {
				hasImage := false
				tableheading.Find("img").Each(func(indexth int, tablecell *goquery.Selection) {
					class, _ := tablecell.Attr("src")
					row = append(row, class)
					hasImage = true
				})
				if hasImage == false {
					tableheadingTrim := strings.TrimSpace(tableheading.Text())
					row = append(row, tableheadingTrim)
				}
			})
			rowhtml.Find("td").Each(func(indexth int, tablecell *goquery.Selection) {
				_, attrExists := tablecell.Attr("class")
				if attrExists == false {
					tablecellTrim := strings.TrimSpace(tablecell.Text())
					row = append(row, tablecellTrim)
				}
			})
			rows = append(rows, row)
			row = nil
		})
	})
	return rows, nil
}

func convertToVillagers(HTMLString string) ([]Villager, error) {
	const UNECCESSARY_ROWS = 2
	const MAX_INFO = 5

	rows, err := convertHTMLStringToArray(HTMLString)
	if err != nil {
		return nil, err
	}

	villagers := []Villager{}
	for i := UNECCESSARY_ROWS; i < len(rows); i++ {
		if len(rows[i]) > MAX_INFO {
			newVillager := Villager{
				string(rows[i][0]),
				string(rows[i][1]),
				string(rows[i][2]),
				string(rows[i][3]),
				string(rows[i][4]),
				string(rows[i][5]),
			}
			villagers = append(villagers, newVillager)
		}
	}
	return villagers, nil
}
