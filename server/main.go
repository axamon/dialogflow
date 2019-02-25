package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/golang/protobuf/jsonpb"
	"github.com/sirupsen/logrus"
	dialogflow "google.golang.org/genproto/googleapis/cloud/dialogflow/v2"
)

func handleWebhook(c *gin.Context) {
	var err error

	wr := dialogflow.WebhookRequest{}
	if err = jsonpb.Unmarshal(c.Request.Body, &wr); err != nil {
		logrus.WithError(err).Error("Couldn't Unmarshal request to jsonpb")
		c.Status(http.StatusBadRequest)
		return
	}

	rr := dialogflow.WebhookResponse{}

	cercorisposta := dialogflow.EventInput{
		Name: "input.Risposta",
	}

	if rr.FollowupEventInput == &cercorisposta {
		rr.FulfillmentText = "Non lo so"
	}

	q := wr.GetQueryResult()
	f := q.GetFulfillmentMessages()

	fmt.Println("la richiesta è ", f)

	fmt.Println(wr.GetQueryResult().GetOutputContexts())
	fmt.Println("")
	fmt.Println(wr.GetQueryResult().GetParameters())
	fmt.Println("")
	fmt.Println(wr.GetQueryResult().GetParameters().GetFields())
	fmt.Println("")
	fmt.Println(wr.GetQueryResult().GetQueryText())
	fmt.Println("")
	fields := wr.GetQueryResult().GetParameters().GetFields()

	fmt.Println("")

	fmt.Println(fields)

	nome := fields["nome"]
	servizio := fields["servizi"]
	mailcliente := fields["email"]

	fmt.Println(nome.GetStringValue())
	fmt.Println(servizio.GetStringValue())
	fmt.Println(mailcliente.GetStringValue())
}

/*
integrazione con mantis
w920KUdcZNzbNTveqCPW0-QLz7pVzUlS

curl --location --request POST "https://mantisomd.westeurope.cloudapp.azure.com/api/rest/issues/" \
  --header "Authorization: w920KUdcZNzbNTveqCPW0-QLz7pVzUlS" \
  --header "Content-Type: application/json" \
  --data "{
  \"summary\": \"This is a test issue\",
  \"description\": \"This is a test description\",
  \"category\": {
    \"name\": \"General\"
  },
  \"project\": {
  	\"id\": 500,
    \"name\": \"mantisbt2\"
  }
}"

*/

func main() {
	var err error

	r := gin.Default()
	r.POST("/webhook", handleWebhook)

	if err = r.Run("0.0.0.0:8080"); err != nil {
		logrus.WithError(err).Fatal("Couldn't start server")
	}
}
