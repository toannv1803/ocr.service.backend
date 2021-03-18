package ImageUseCase

import (
	"errors"
	ImageInterface "ocr.service.backend/app/image/interface"
	ImageRepository "ocr.service.backend/app/image/repository"
	"ocr.service.backend/enum"
	"ocr.service.backend/model"
)

type ImageUseCase struct {
	repository ImageInterface.IImageRepository
}

func (q *ImageUseCase) Gets(agent model.Agent, filter model.Image) ([]model.Image, error) {
	switch agent.Role {
	case enum.RoleAdmin:
		return q.repository.Get(filter)
	case enum.RoleUser:
		if agent.UserId == "" {
			return nil, errors.New("agent not found user_id")
		}
		filter.UserId = agent.UserId
		return q.repository.Get(filter)
	default: //enum.RoleAnonymous
		filter.UserId = enum.RoleAnonymous
		return q.repository.Get(filter)
	}
}
func (q *ImageUseCase) InsertOne(agent model.Agent, image model.Image) (string, error) {
	switch agent.Role {
	case enum.RoleAdmin, enum.RoleUser:
		return q.repository.InsertOne(image)
	default: //enum.RoleAnonymous:
		image.UserId = enum.RoleAnonymous
		return q.repository.InsertOne(image)
	}
}
func (q *ImageUseCase) Update(agent model.Agent, filter model.Image, image model.Image) (int64, error) {
	switch agent.Role {
	case enum.RoleAdmin:
		return q.repository.Update(filter, image)
	case enum.RoleUser:
		if agent.UserId == "" {
			return 0, errors.New("not found user_id")
		}
		filter.UserId = agent.UserId
		return q.repository.Update(filter, image)
	default: //enum.RoleAnonymous
		filter.UserId = agent.UserId
		return q.repository.Update(filter, image)
	}
}

func NewImageUseCase() (ImageInterface.IImageUseCase, error) {
	var q ImageUseCase
	var err error
	q.repository, err = ImageRepository.NewImageRepository()
	return &q, err
}
