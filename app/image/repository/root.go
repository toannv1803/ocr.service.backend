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

func (q *ImageRepository) Find(filter model.Image, option db.FindOption) ([]model.Image, int64, error) {
	var arrUser []model.Image
	total, err := q.db.Find(filter, &arrUser, option)
	if err != nil {
		fmt.Println(err)
		return arrUser, 0, errors.New("get from db failed")
	}
	return arrUser, total, err
}

func (q *ImageRepository) FindCustom(filter model.Image, res interface{}, option db.FindOption) (int64, error) {
	total, err := q.db.Find(filter, res, option)
	if err != nil {
		fmt.Println(err)
		return 0, errors.New("get from db failed")
	}
	return total, nil
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
func (q *ImageRepository) Delete(filter model.Image) (int64, error) {
	nDelete, err := q.db.Delete(filter)
	if err != nil {
		fmt.Println(err)
		return 0, errors.New("delete failed")
	}
	return nDelete, nil
}

func (q *ImageRepository) Distinct(field string, filter interface{}) ([]interface{}, error) {
	arrI, err := q.db.Distinct(field, filter)
	if err != nil {
		fmt.Println(err)
		return nil, errors.New("get distinct failed")
	}
	return arrI, nil
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
