package main

import (
	"encoding/json"
	"errors"
	"net/http"
	"time"

	"github.com/aws/aws-lambda-go/events"
)

func extractPageHTML(url string) (string, error) {
	timeout := time.Duration(10 * time.Second)
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

	// Assumes that JSON is not malformed etc.
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
	url := "https://nookipedia.com/w/api.php?action=parse&section=2&origin=*&prop=text&page=List_of_villagers&format=json"
	HTMLString, err := extractPageHTML(url)
	if err != nil {
		return get403Response(err)
	}

	villagers, err := convertToVillagers(HTMLString)
	if err != nil {
		return get403Response(err)
	}

	// Filter
	if params["Species"] != "" {
		var filteredVillagers []Villager
		species := params["Species"]
		for _, villager := range villagers {
			if villager.Species == species {
				filteredVillagers = append(filteredVillagers, villager)
			}
		}
		villagers = filteredVillagers
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

// func main() {
// Handler("https://nookipedia.com/w/api.php?action=parse&page=List_of_villagers&format=json")
// fmt.Println(Handler(
// 	"https://nookipedia.com/w/api.php?action=parse&section=2&origin=*&prop=text&page=List_of_villagers&format=json")) // Error with client making request

// fmt.Println(Handler("https://nookipedia.com/w/api.php?action=parse&section=2&origin=*&page=List_of_villagers&format=json")) // Error with client making request
// Handler("https://en.wikipedia.org/w/api.php?action=query&origin=*&list=search&srsearch=Craig%20Noone&format=jsonfm")
// Handler("https://jsonplaceholder.typicode.com/todos/1")
// url := "https://jsonplaceholder.typicode.com/todos"
// fmt.Println("Hello World")

// timeout := time.Duration(5 * time.Second) // Response using client
// client := http.Client{
// 	Timeout: timeout,
// }

// req, err := http.NewRequest("GET", url, nil)

// if err != nil {
// 	print("Error: making request failed")
// }

// resp, err := client.Do(req)
// if err != nil {
// 	print("Error: client making request failed")
// }

// resp, err := http.Get("https://jsonplaceholder.typicode.com/todos/1") // Single JSON response of unknown type
// resp, err := http.Get(url) // Array JSON response of unknown type

// if err != nil {
// 	println("Error: Get request failed")
// }
// defer resp.Body.Close()
// if resp.StatusCode != http.StatusOK {
// 	println("Error: Status code not okay")
// }

// var result map[string]interface{} // Single JSON response of unknown type
// err = json.NewDecoder(resp.Body).Decode(&result)
// if err != nil {
// 	println("Error: Failed to decode body of response")
// }
// for key, value := range result {
// 	stringValue := fmt.Sprint(value)
// 	println("Key: ", key, "Value: ", stringValue)
// }
// println("Here is the Result ", result)

// var result []map[string]interface{} // Array JSON response of unknown type
// err = json.NewDecoder(resp.Body).Decode(&result)

// for _, value := range result {
// 	for subKey, subValue := range value {
// 		stringValue := fmt.Sprint(subValue)
// 		println(" eachresult ", subKey, stringValue)
// 	}
// }

// body, err := ioutil.ReadAll(resp.Body) // Plain text response
// println("Here is the Result ", string(body))
// }
