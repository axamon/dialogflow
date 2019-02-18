package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/golang/protobuf/jsonpb"
	"github.com/sirupsen/logrus"
	"google.golang.org/genproto/googleapis/cloud/dialogflow/v2"
)

func handleWebhook(c *gin.Context) {
	var err error

	wr := dialogflow.WebhookRequest{}
	if err = jsonpb.Unmarshal(c.Request.Body, &wr); err != nil {
		logrus.WithError(err).Error("Couldn't Unmarshal request to jsonpb")
		c.Status(http.StatusBadRequest)
		return
	}
	fmt.Println(wr.GetQueryResult().GetOutputContexts())
	fmt.Println(wr.GetQueryResult().GetParameters())
	fmt.Println(wr.GetQueryResult().GetQueryText())
	fields := wr.GetQueryResult().GetParameters().GetFields()

	fmt.Println("")

	fmt.Println(fields)

	nome := fields["given-name"]
	servizio := fields["servizi"]
	mailcliente := fields["email"]

	fmt.Println(nome.GetStringValue())
	fmt.Println(servizio.GetStringValue())
	fmt.Println(mailcliente.GetStringValue())
}

func main() {
	var err error

	r := gin.Default()
	r.POST("/webhook", handleWebhook)

	if err = r.Run("0.0.0.0:8080"); err != nil {
		logrus.WithError(err).Fatal("Couldn't start server")
	}
}
