package ImageUseCase

import (
	"errors"
	ImageInterface "ocr.service.backend/app/image/interface"
	ImageRepository "ocr.service.backend/app/image/repository"
	"ocr.service.backend/enum"
	"ocr.service.backend/model"
	"ocr.service.backend/module/db"
)

type ImageUseCase struct {
	repository ImageInterface.IImageRepository
}

func (q *ImageUseCase) Gets(agent model.Agent, filter model.Image, option ImageInterface.GetOption) ([]model.Image, int64, error) {
	switch agent.Role {
	case enum.RoleAdmin:
		return q.repository.Find(filter, db.FindOption{Skip: option.Skip, Limit: option.Limit})
	case enum.RoleUser:
		if agent.UserId == "" {
			return nil, 0, errors.New("agent not found user_id")
		}
		filter.UserId = agent.UserId
		return q.repository.Find(filter, db.FindOption{Skip: option.Skip, Limit: option.Limit})
	default: //enum.RoleAnonymous
		filter.UserId = enum.RoleAnonymous
		return q.repository.Find(filter, db.FindOption{Skip: option.Skip, Limit: option.Limit})
	}
}

func (q *ImageUseCase) GetsCustom(agent model.Agent, filter model.Image, res interface{}, option ImageInterface.GetOption) (int64, error) {
	switch agent.Role {
	case enum.RoleAdmin:
		return q.repository.FindCustom(filter, res, db.FindOption{Skip: option.Skip, Limit: option.Limit})
	case enum.RoleUser:
		if agent.UserId == "" {
			return 0, errors.New("agent not found user_id")
		}
		filter.UserId = agent.UserId
		return q.repository.FindCustom(filter, res, db.FindOption{Skip: option.Skip, Limit: option.Limit})
	default: //enum.RoleAnonymous
		filter.UserId = enum.RoleAnonymous
		return q.repository.FindCustom(filter, res, db.FindOption{Skip: option.Skip, Limit: option.Limit})
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
		filter.UserId = enum.RoleAnonymous
		return q.repository.Update(filter, image)
	}
}
func (q *ImageUseCase) Delete(agent model.Agent, filter model.Image) (int64, error) {
	if filter == (model.Image{}) {
		return 0, errors.New("delete image require at least one query")
	}
	switch agent.Role {
	case enum.RoleAdmin:
		return q.repository.Delete(filter)
	case enum.RoleUser:
		if agent.UserId == "" {
			return 0, errors.New("not found user_id")
		}
		filter.UserId = agent.UserId
		return q.repository.Delete(filter)
	default: //enum.RoleAnonymous
		filter.UserId = enum.RoleAnonymous
		return q.repository.Delete(filter)
	}
}

func (q *ImageUseCase) GetListBlockId(agent model.Agent) ([]string, error) {
	var arrString []string
	switch agent.Role {
	case enum.RoleAdmin:
		arrI, err := q.repository.Distinct("block_id", model.Image{UserId: agent.UserId})
		if err != nil {
			return nil, err
		}
		for i := range arrI {
			arrString = append(arrString, arrI[i].(string))
		}
	case enum.RoleUser:
		if agent.UserId == "" {
			return nil, errors.New("not found user_id")
		}
		arrI, err := q.repository.Distinct("block_id", model.Image{UserId: agent.UserId})
		if err != nil {
			return nil, err
		}
		for i := range arrI {
			arrString = append(arrString, arrI[i].(string))
		}
	default: //enum.RoleAnonymous
		arrI, err := q.repository.Distinct("block_id", model.Image{UserId: enum.RoleAnonymous})
		if err != nil {
			return nil, err
		}
		for i := range arrI {
			arrString = append(arrString, arrI[i].(string))
		}
	}
	return arrString, nil
}

func NewImageUseCase() (ImageInterface.IImageUseCase, error) {
	var q ImageUseCase
	var err error
	q.repository, err = ImageRepository.NewImageRepository()
	return &q, err
}
