package ImageRepository

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"ocr/model"
	"time"
)

type ImageRepository struct {
	collection *mongo.Collection
}

var nullImage = model.Image{}

func (q *ImageRepository) Get() (model.Image, error) {
	//q.collection.Find()
	return nullImage, nil
}

func (q *ImageRepository) InsertOne(image model.Image) (string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	res, err := q.collection.InsertOne(ctx, image)
	if err != nil {
		return "", err
	}
	return res.InsertedID.(string), nil
}

func (q *ImageRepository) Update(filter model.Image, image model.Image) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	_, err := q.collection.UpdateOne(ctx, filter, image)
	if err != nil {
		return err
	}
	return nil
}

func NewImageRepository() (*ImageRepository, error) {
	var q ImageRepository
	ctx, cancel := context.WithTimeout(context.Background(), 7*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		return nil, err
	}
	collection := client.Database("orc").Collection("images")
	q.collection = collection
	return &q, nil
}
