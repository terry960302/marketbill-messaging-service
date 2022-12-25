package main

import (
	"context"
	"fmt"
	"log"
	"marketbill-messaging-service/constants"
	"marketbill-messaging-service/controllers"
	"marketbill-messaging-service/datastore"
	"marketbill-messaging-service/models"
	"marketbill-messaging-service/services"
	"marketbill-messaging-service/test"
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
func main() {
	setter := test.NewEnvSetter("local")
	setter.SetEnv()

	db, _ := datastore.NewPostgresql()
	s := services.NewSmsService(db)
	res, err := s.SendDefaultSMS("01091751159", "this is test", constants.SMS)
	if err != nil {
		panic(err)
	}
	fmt.Println(res)
}
