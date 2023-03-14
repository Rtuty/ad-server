package mongo

import (
	"github.com/valyala/fasthttp"
	"go.mongodb.org/mongo-driver/mongo"
)

type CustomerMetric struct {
	c *mongo.Collection
}

func NewAdDAO(ctx *fasthttp.RequestCtx, client *mongo.Client) (*CustomerMetric, error) {
	return &CustomerMetric{
		c: client.Database("core").Collection("metrics"),
	}, nil
}
