package handlers

import (
	"bytes"
	"io"
	"io/ioutil"
	"log"

	"github.com/deadlock-delegate/payfee-api/config"
	"github.com/gin-gonic/gin"
)

func WebhookEvents(c *gin.Context) {
	auth := c.GetHeader("authorization")
	token := auth + config.Webhook.VerificationKey
	if token != config.Webhook.Token {
		c.JSON(400, gin.H{"message": "Error"})
		return
	}

	var buf bytes.Buffer
	tee := io.TeeReader(c.Request.Body, &buf)
	body, _ := ioutil.ReadAll(tee)
	log.Println(string(body))

	c.JSON(200, gin.H{"message": "OK"})
}
