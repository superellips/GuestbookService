package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Guestbook struct {
	Id              primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	OwnerId         primitive.ObjectID `json:"ownerId"`
	Domain          string             `json:"domain"`
	RequireApproval bool               `json:"requireApproval"`
}

func PostGuestbook(c *gin.Context) {
	var newGuestbook Guestbook
	if err := c.BindJSON(&newGuestbook); err != nil {
		return
	}
	// Persist logic
	var db MongoDb
	db.create(&newGuestbook)
	c.IndentedJSON(http.StatusCreated, newGuestbook)
}

func GetGuestbook(c *gin.Context) {
	id, err := primitive.ObjectIDFromHex(c.Param("gbId"))
	if err != nil {
		return
	}
	// Persist logic
	var db MongoDb
	guestbook := db.read(&id)
	c.IndentedJSON(http.StatusOK, guestbook)
}

func PutGuestbook(c *gin.Context) {
	var updatedGuestbook Guestbook
	if err := c.BindJSON(&updatedGuestbook); err != nil {
		return
	}
	// Persist logic
	var db MongoDb
	updatedGuestbook = *db.update(&updatedGuestbook)
	c.IndentedJSON(http.StatusAccepted, updatedGuestbook)
}

func DeleteGuestbook(c *gin.Context) {
	id, err := primitive.ObjectIDFromHex(c.Param("gbId"))
	if err != nil {
		return
	}
	// Persist logic
	var db MongoDb
	deletedGuestbook := *db.delete(&id)
	c.IndentedJSON(http.StatusAccepted, deletedGuestbook)
}
