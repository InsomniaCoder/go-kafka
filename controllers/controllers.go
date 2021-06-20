package controllers

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/insomniacoder/go-kafka/models"
	"github.com/insomniacoder/go-kafka/producers"
)

func PublishComment(c *gin.Context) {

	var publishRequest models.Message
	c.BindJSON(&publishRequest)

	log.Printf("Publishing Message %s \n", publishRequest)

	err := producers.PushCommentToQueue("message", []byte(publishRequest.Text))

	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
	}

	defer c.JSON(http.StatusOK, "message published")
}
