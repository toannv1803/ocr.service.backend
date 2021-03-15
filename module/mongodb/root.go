package mongodb

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"ocr.service.backend/config"
	"time"
)

func NewClient() (*mongo.Client, error) {
	CONFIG, err := config.NewConfig(nil)
	if err != nil {
		return nil, err
	}
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	//return  mongo.Connect(ctx, options.Client().ApplyURI("mongodb://localhost:27017"))
	var client *mongo.Client
	if CONFIG.GetString("MONGODB_USERNAME") != "" {
		client, err = mongo.Connect(ctx, options.Client().ApplyURI("mongodb://"+CONFIG.GetString("MONGODB_USERNAME")+":"+CONFIG.GetString("MONGODB_PASSWORD")+"@"+CONFIG.GetString("MONGODB_HOST")+":"+CONFIG.GetString("MONGODB_PORT")+"/?authSource="+CONFIG.GetString("MONGODB_DB")))
	} else {
		client, err = mongo.Connect(ctx, options.Client().ApplyURI("mongodb://"+CONFIG.GetString("MONGODB_HOST")+":"+CONFIG.GetString("MONGODB_PORT")+"/?authSource="+CONFIG.GetString("MONGODB_DB")))
	}
	if err != nil {
		return nil, err
	}
	err = client.Ping(ctx, readpref.Primary())
	if err != nil {
		return nil, err
	}
	return client, nil
}
