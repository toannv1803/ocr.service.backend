package mongodb

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
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
	if CONFIG.GetString("MONGODB_USERNAME") != "" {
		return mongo.Connect(ctx, options.Client().ApplyURI("mongodb://"+CONFIG.GetString("MONGODB_USERNAME")+":"+CONFIG.GetString("MONGODB_PASSWORD")+"@"+CONFIG.GetString("MONGODB_HOST")+":"+CONFIG.GetString("MONGODB_PORT")+"/?authSource="+CONFIG.GetString("MONGODB_DB")))
	} else {
		return mongo.Connect(ctx, options.Client().ApplyURI("mongodb://"+CONFIG.GetString("MONGODB_HOST")+":"+CONFIG.GetString("MONGODB_PORT")+"/?authSource="+CONFIG.GetString("MONGODB_DB")))
	}
}
