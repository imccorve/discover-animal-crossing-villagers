package main

import (
	"encoding/json"
	"errors"
	"net/http"
	"time"

	"github.com/aws/aws-lambda-go/events"
)

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

func formResponse(params map[string]string) events.APIGatewayProxyResponse {
	response := events.APIGatewayProxyResponse{}
	// response := APIGatewayProxyResponse{Headers: make(map[string]string)}

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
