package ImageRepository

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"ocr.service.backend/config"
	"ocr.service.backend/model"
	"ocr.service.backend/module/mongodb"
	"os"
	"time"
)

type ImageRepository struct {
	collection *mongo.Collection
}

func (q *ImageRepository) Get(filter model.Image) ([]model.Image, error) {
	//var _filter interface{} = filter
	//if filter.Id != "" {
	//	objectId, err := primitive.ObjectIDFromHex(filter.Id)
	//	if err != nil {
	//		return nil, err
	//	}
	//	_filter = bson.M{"_id": objectId}
	//}
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	cur, err := q.collection.Find(ctx, filter)
	if err != nil {
		return nil, err
	}
	var arrOriginModel []model.Image
	defer cur.Close(ctx)
	for cur.Next(ctx) {
		var result model.Image
		err := cur.Decode(&result)
		if err == nil {
			arrOriginModel = append(arrOriginModel, result)
		} else {
			fmt.Println("GetOrigin", err)
		}
	}
	return arrOriginModel, nil
}
func (q *ImageRepository) InsertOne(image model.Image) (string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	res, err := q.collection.InsertOne(ctx, image)
	if err == nil {
		return res.InsertedID.(primitive.ObjectID).Hex(), nil
	}
	return "", err
}
func (q *ImageRepository) Delete(image model.Image) (int64, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	result, err := q.collection.DeleteMany(ctx, image)
	if err == nil {
		return result.DeletedCount, err
	}
	return 0, err
}
func (q *ImageRepository) Update(filter model.Image, image model.Image) error {
	//var _filter interface{} = filter
	//if filter.Id != "" {
	//	objectId, err := primitive.ObjectIDFromHex(filter.Id)
	//	if err != nil {
	//		return err
	//	}
	//	_filter = bson.M{"_id": objectId}
	//}
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	data := bson.D{{"$set", image}}
	_, err := q.collection.UpdateOne(ctx, filter, data)
	if err == nil {
		return err
	}
	return err
}

func NewImageRepository() (*ImageRepository, error) {
	CONFIG, err := config.NewConfig(nil)
	if err != nil {
		return nil, err
	}
	var p ImageRepository
	client, err := mongodb.NewClient()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	collection := client.Database(CONFIG.GetString("MONGODB_DB")).Collection("images")
	p.collection = collection
	return &p, nil
}
