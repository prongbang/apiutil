// package main

// import (
// 	"github.com/aws/aws-lambda-go/events"
// 	"github.com/aws/aws-lambda-go/lambda"
// )

// func handler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
// 	return events.APIGatewayProxyResponse{
// 		StatusCode: 200,
// 		Body:       "Hello AWS Lambda and Netlify",
// 	}, nil
// }

// func main() {
// 	// Make the handler available for Remote Procedure Call by AWS Lambda
// 	lambda.Start(handler)
// }

package main

import (
	"fmt"
	"net/http"
	"os"
	"time"
)

func greet(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World! %s", time.Now())
}

func main() {
	port := os.Getenv("_LAMBDA_SERVER_PORT")
	http.HandleFunc("/api/v1/greet", greet)
	http.ListenAndServe(":"+port, nil)
}
