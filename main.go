package main

import (
	"context"
	"log"
	"marketbill-messaging-service/controllers"
	"marketbill-messaging-service/models"
	"net/http"
	"os"

	"github.com/aws/aws-lambda-go/events"
)

func init() {
	profile := os.Getenv("PROFILE")
	log.Print("PROFILE : ", profile)
}

func HandleRequest(ctx context.Context, request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {

	r := models.NewLambdaResponse()
	switch request.HTTPMethod {
	case "GET":
		return controllers.HealthCheck(request)
	case "POST":
		return controllers.HandleSMS(request)
	default:
		return r.Error(http.StatusBadRequest, "Wrong http method")
	}
}

// func main() {
// 	lambda.Start(HandleRequest)
// }

// test
// func main() {
// 	setter := test.NewEnvSetter("local")
// 	setter.SetEnv()

// 	db, _ := datastore.NewPostgresql()

// 	req := models.MessagingRequest{
// 		To:       "01091751159",
// 		Template: "Default",
// 		Args:     []any{"01011010100"},
// 	}
// 	ctrl := controllers.HandleSMS(req)
// }
