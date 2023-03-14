package mongo

import "go.mongodb.org/mongo-driver/mongo"

type UrlDAO struct {
	c *mongo.Collection
}
