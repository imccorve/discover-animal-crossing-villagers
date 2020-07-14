package main

import (
	"encoding/json"
	"errors"
	"net/http"
	"strconv"
	"time"

	"github.com/aws/aws-lambda-go/events"
)

func min(vars ...int) int {
	min := vars[0]
	for _, i := range vars {
		if min > i {
			min = i
		}
	}
	return min
}

func max(vars ...int) int {
	max := vars[0]
	for _, i := range vars {
		if max < i {
			max = i
		}
	}
	return max
}
func extractPageHTML(url string) (string, error) {
	timeout := time.Duration(15 * time.Second)
	client := http.Client{
		Timeout: timeout,
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return "", err
	}

	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}

	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return "", errors.New("Something went wrong" + string(resp.StatusCode))
	}

	var result map[string]interface{} // Single JSON response of unknown type
	err = json.NewDecoder(resp.Body).Decode(&result)
	if err != nil {
		return "", err
	}

	// Assumes that Nookipedia JSON is not malformed etc.
	HTMLString := result["parse"].(map[string]interface{})["text"].(map[string]interface{})["*"].(string)

	return HTMLString, nil
}

func get403Response(err error) events.APIGatewayProxyResponse {
	response := events.APIGatewayProxyResponse{}
	response.StatusCode = 403
	response.Body = err.Error()
	return response
}
func paginateResults(params map[string]string, villagers []Villager) ([]Villager, error) {
	lenVillagers := len(villagers)
	limit := 20
	offset := 0
	var err error = nil

	paramLimit, hasLimit := params["limit"]
	paramOffset, hasOffset := params["offset"]

	if hasLimit {
		limit, err = strconv.Atoi(paramLimit)
		if err != nil {
			return nil, err
		}
	}
	if hasOffset {
		offset, err = strconv.Atoi(paramOffset)
		if err != nil {
			return nil, err
		}
	}

	limit = max(1, limit)

	start := min(lenVillagers, max(limit*offset, 0))
	end := min(start+min(limit, lenVillagers), lenVillagers)
	villagers = villagers[start:end]

	return villagers, nil
}

func formResponse(params map[string]string) events.APIGatewayProxyResponse {
	response := events.APIGatewayProxyResponse{}

	// Call Nookipedia API
	url := "https://nookipedia.com/w/api.php?action=parse&section=3&origin=*&prop=text&page=List_of_villagers&format=json"
	HTMLString, err := extractPageHTML(url)
	if err != nil {
		return get403Response(err)
	}

	villagers, err := convertToVillagers(HTMLString)
	if err != nil {
		return get403Response(err)
	}
	if len(params) > 0 {
		villagers = filterVillagers(villagers, params)
	}

	villagers, err = paginateResults(params, villagers)
	if err != nil {
		return get403Response(err)
	}

	villagersBytes, err := json.Marshal(villagers)
	if err != nil {
		return get403Response(err)
	}

	response.StatusCode = 200
	response.Headers = map[string]string{
		"Access-Control-Allow-Origin":  "*",
		"Access-Control-Allow-Headers": "*",
	}
	response.Body = string(villagersBytes)
	return response
}

// Handler handles the request and sends a response
func Handler(params events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	response := formResponse(params.QueryStringParameters)
	return response, nil
}
