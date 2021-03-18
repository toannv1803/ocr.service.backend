package ImageRepository

import (
	"errors"
	"fmt"
	ImageInterface "ocr.service.backend/app/image/interface"
	"ocr.service.backend/config"
	"ocr.service.backend/model"
	"ocr.service.backend/module/db"
)

type ImageRepository struct {
	db db.IDB
}

func (q *ImageRepository) Get(filter model.Image) ([]model.Image, error) {
	var arrUser []model.Image
	err := q.db.Get(filter, &arrUser)
	if err != nil {
		fmt.Println(err)
		return arrUser, errors.New("get from db failed")
	}
	return arrUser, err
}
func (q *ImageRepository) InsertOne(user model.Image) (string, error) {
	id, err := q.db.InsertOne(user)
	if err != nil {
		fmt.Println(err)
		return "", errors.New("insert to db failed")
	}
	return id, err
}
func (q *ImageRepository) Update(filter model.Image, user model.Image) (int64, error) {
	nModify, err := q.db.Update(filter, user)
	if err != nil {
		fmt.Println(err)
		return 0, errors.New("update db failed")
	}
	return nModify, err
}

func NewImageRepository() (ImageInterface.IImageRepository, error) {
	var q ImageRepository
	var err error
	CONFIG, err := config.NewConfig(nil)
	if err != nil {
		return nil, err
	}
	q.db, err = db.NewMongoRepository(CONFIG.GetString("MONGODB_DB"), "images")
	return &q, err
}
