package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Guestbook struct {
	Id              primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	OwnerId         primitive.ObjectID `json:"ownerId" bson:"ownerid"`
	Domain          string             `json:"domain" bson:"domain"`
	RequireApproval bool               `json:"requireApproval" bson:"requireapproval"`
}

func PostGuestbook(c *gin.Context) {
	createRequest, err := io.ReadAll(c.Request.Body)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to read request data"})
		return
	}
	var d map[string]interface{}
	err = json.Unmarshal(createRequest, &d)
	if err != nil {
		fmt.Println("It got all hecked up.")
		return
	}
	ownerId, err := primitive.ObjectIDFromHex(d["ownerId"].(string))
	if err != nil {
		fmt.Println("It got all hecked up.")
		return
	}
	newGuestbook := Guestbook{
		OwnerId:         ownerId,
		Domain:          d["domain"].(string),
		RequireApproval: d["requireApproval"].(bool),
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
