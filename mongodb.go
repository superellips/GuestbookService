package main

import (
	"context"
	"os"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoDb struct {
	Uri string
}

func (p *MongoDb) create(g *Guestbook) *Guestbook {
	coll := p.guestbookCollection()
	result, err := coll.InsertOne(
		context.TODO(),
		g)
	if err != nil {
		panic(err)
	}
	g.Id = result.InsertedID.(primitive.ObjectID)
	return g
}

func (p *MongoDb) read(i *primitive.ObjectID) *Guestbook {
	coll := p.guestbookCollection()
	filter := bson.D{{Key: "_id", Value: i}}
	var result Guestbook
	err := coll.FindOne(
		context.TODO(),
		filter).Decode(&result)
	if err != nil {
		panic(err)
	}
	return &result
}

func (p *MongoDb) update(g *Guestbook) *Guestbook {
	coll := p.guestbookCollection()
	filter := bson.D{{Key: "_id", Value: g.Id}}
	var result Guestbook
	err := coll.FindOneAndReplace(
		context.TODO(),
		filter,
		g).Decode(&result)
	if err != nil {
		panic(err)
	}
	return &result
}

func (p *MongoDb) delete(i *primitive.ObjectID) *Guestbook {
	coll := p.guestbookCollection()
	filter := bson.D{{Key: "_id", Value: i}}
	var deletedGuestbook Guestbook
	err := coll.FindOneAndDelete(
		context.TODO(),
		filter).Decode(&deletedGuestbook)
	if err != nil {
		panic(err)
	}
	return &deletedGuestbook
}

func (p *MongoDb) createMessage(m *Message) *Message {
	coll := p.messageCollection()
	result, err := coll.InsertOne(
		context.TODO(),
		m)
	if err != nil {
		panic(err)
	}
	m.Id = result.InsertedID.(primitive.ObjectID)
	return m
}

func (p *MongoDb) readMessagesByGuestbookId(i *primitive.ObjectID, a bool) *[]Message {
	coll := p.messageCollection()
	filter := bson.D{{Key: "guestbookId", Value: i}, {Key: "approved", Value: a}}
	var results []Message
	cur, err := coll.Find(
		context.TODO(),
		filter)
	if err != nil {
		panic(err)
	}
	cur.All(
		context.TODO(),
		&results)
	return &results
}

func (p *MongoDb) updateMessage(m *Message) *Message {
	coll := p.messageCollection()
	filter := bson.D{{Key: "_id", Value: m.Id}}
	var result Message
	err := coll.FindOneAndReplace(
		context.TODO(),
		filter,
		m).Decode(&result)
	if err != nil {
		panic(err)
	}
	return &result
}

func (p *MongoDb) deleteMessage(gbId *primitive.ObjectID, msgId *primitive.ObjectID) *Message {
	coll := p.messageCollection()
	filter := bson.D{{Key: "guestbookId", Value: gbId}, {Key: "_id", Value: msgId}}
	var deletedMessage Message
	err := coll.FindOneAndDelete(
		context.TODO(),
		filter).Decode(&deletedMessage)
	if err != nil {
		panic(err)
	}
	return &deletedMessage
}

func (p *MongoDb) guestbookCollection() *mongo.Collection {
	p.Uri = os.Getenv("GB_CONSTRING")
	client, err := mongo.Connect(context.TODO(), options.Client().
		ApplyURI(p.Uri))
	if err != nil {
		panic(err)
	}
	return client.Database("guestbook-service").Collection("guestbooks")
}

func (p *MongoDb) messageCollection() *mongo.Collection {
	p.Uri = os.Getenv("GB_CONSTRING")
	client, err := mongo.Connect(
		context.TODO(),
		options.Client().ApplyURI(p.Uri))
	if err != nil {
		panic(err)
	}
	return client.Database("guestbook-service").Collection("messages")
}
