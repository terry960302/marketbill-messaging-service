package main

import (
	"context"
	"log"
	"marketbill-messaging-service/handlers"
	"marketbill-messaging-service/models"
	"net/http"
	"os"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

func init() {
	profile := os.Getenv("PROFILE")
	log.Print("PROFILE : ", profile)
}

func HandleRequest(ctx context.Context, request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {

	r := models.NewLambdaResponse()
	switch request.HTTPMethod {
	case "GET":
		return handlers.HealthCheck(request)
	case "POST":
		return handlers.HandleSMS(request)
	default:
		return r.Error(http.StatusBadRequest, "Wrong http method")
	}
}

func main() {
	lambda.Start(HandleRequest)
	// os.Setenv("PROFILE", "dev")
	// os.Setenv("SENS_HOST", "https://sens.apigw.ntruss.com")
	// os.Setenv("SENS_SERVICE_ID", "ncp:sms:kr:290881020329:marketbill-project")
	// os.Setenv("SENS_ACCESS_KEY_ID", "2aJkrtHdUtk5NP4oG8yh")
	// os.Setenv("SENS_SECRET_KEY", "x2A3OXOz0P1qmaLDnTiqo2dQ7if6BzOElQEPNg6b")
	// os.Setenv("DB_USER", "marketbill")
	// os.Setenv("DB_PW", "marketbill1234!")
	// os.Setenv("DB_NET", "tcp")
	// os.Setenv("DB_HOST", "marketbill-db.ciegftzvpg1l.ap-northeast-2.rds.amazonaws.com")
	// os.Setenv("DB_PORT", "5432")
	// os.Setenv("DB_NAME", "dev-db")

	// db, _ := datastore.NewPostgresql()
	// s := services.NewSmsService(db)
	// s.SendDefaultSMS("01091751159","asdasdadads")
}
