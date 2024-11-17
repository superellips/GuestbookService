package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Message struct {
	Id          primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	GuestbookId primitive.ObjectID `json:"guestbookId" bson:"guestbookId"`
	SenderName  string             `json:"senderName"`
	SenderEmail string             `json:"senderEmail"`
	Text        string             `json:"text"`
	Approved    bool               `json:"approved"`
}

func PostMessage(c *gin.Context) {
	var newMessage Message
	if err := c.BindJSON(&newMessage); err != nil {
		return
	}
	gbId, err := primitive.ObjectIDFromHex(c.Param("gbId"))
	if err != nil {
		return
	}
	newMessage.GuestbookId = gbId
	// Persist logic
	var db MongoDb
	newMessage.Approved = !db.read((&newMessage.GuestbookId)).RequireApproval
	db.createMessage(&newMessage)
	c.IndentedJSON(http.StatusCreated, newMessage)
}

func GetMessagesByGuestbookId(c *gin.Context) {
	id, err := primitive.ObjectIDFromHex(c.Param("gbId"))
	if err != nil {
		return
	}
	// Persist logic
	var db MongoDb
	messages := db.readMessagesByGuestbookId(&id, true)
	c.IndentedJSON(http.StatusOK, messages)
}

func PutMessage(c *gin.Context) {
	var updatedMessage Message
	if err := c.BindJSON(&updatedMessage); err != nil {
		return
	}
	// Persist logic
	var db MongoDb
	updatedMessage = *db.updateMessage(&updatedMessage)
	c.IndentedJSON(http.StatusAccepted, updatedMessage)
}

func DeleteMessage(c *gin.Context) {
	gbId, err := primitive.ObjectIDFromHex(c.Param("gbId"))
	if err != nil {
		return
	}
	msgId, err := primitive.ObjectIDFromHex(c.Param("msgId"))
	if err != nil {
		return
	}
	// Persist logic
	var db MongoDb
	deletedMessage := *db.deleteMessage(&gbId, &msgId)
	c.IndentedJSON(http.StatusAccepted, deletedMessage)
}
