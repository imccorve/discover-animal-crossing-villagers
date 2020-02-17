package main

import "github.com/aws/aws-lambda-go/lambda"

func main() {
	lambda.Start(Handler)
	// Handler(events.APIGatewayProxyRequest{})
	// TESTING
	// htmlReadFile, err := ioutil.ReadFile("html_test.html")
	// htmlReadFile, err := ioutil.ReadFile("./parse_text_text.txt")

}
