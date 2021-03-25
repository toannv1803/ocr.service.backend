package ImageDelivery

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/copier"
	ImageInterface "ocr.service.backend/app/image/interface"
	ImageUseCase "ocr.service.backend/app/image/usecase"
	"ocr.service.backend/config"
	"ocr.service.backend/enum"
	"ocr.service.backend/model"
	module "ocr.service.backend/module/rabbitmq"
)

var CONFIG, _ = config.NewConfig(nil)

type ImageDelivery struct {
	useCase           ImageInterface.IImageUseCase
	rabbitmq          *module.RabbitMQ
	imageTaskQueue    string
	imageSuccessQueue string
	imageErrorQueue   string
}

type RabbitMQLogin = module.RabbitMQLogin

// @tags Images
// @Summary image
// @Description get list image
// @start_time default
// @Param image_id path string true "image id"
// @Param Authorization header string true "'Bearer ' + token"
// @Success 200 {object} model.ImageResponse
// @Router /api/v1/auth/image/{image_id} [get]
func (q *ImageDelivery) GetById(c *gin.Context) {
	var agent model.Agent
	if v, ok := c.Get("agent"); ok {
		agent = v.(model.Agent)
	}
	imageId := c.Param("image_id")
	if imageId != "" {
		var arrImageResponse []model.ImageResponse
		err := q.useCase.GetsCustom(agent, model.Image{Id: imageId}, &arrImageResponse)
		if err != nil {
			switch err.Error() {
			case "not found role":
				c.String(401, "not allow")
			default:
				c.String(500, err.Error())
			}
			return
		}
		if len(arrImageResponse) == 0 {
			c.String(404, "not found")
			return
		}
		c.JSON(200, arrImageResponse[0])
		return
	} else {
		c.String(400, "require param image_id")
		return
	}
}

// @tags Images
// @Summary image
// @Description get list image
// @start_time default
// @Param Authorization header string true "'Bearer ' + token"
// @Success 200 {object} []string
// @Router /api/v1/auth/images/block-ids [get]
func (q *ImageDelivery) GetListBlockId(c *gin.Context) {
	var agent model.Agent
	if v, ok := c.Get("agent"); ok {
		agent = v.(model.Agent)
	}
	arrBlockId, err := q.useCase.GetListBlockId(agent)
	if err != nil {
		switch err.Error() {
		case "not found role":
			c.String(401, "not allow")
		default:
			c.String(500, err.Error())
		}
		return
	}
	if len(arrBlockId) == 0 {
		c.String(404, "not found")
		return
	}
	c.JSON(200, arrBlockId)
	return
}

// @tags Images
// @Summary image
// @Description get list image
// @start_time default
// @Param _ query model.ImageFilter true "_"
// @Param Authorization header string true "'Bearer ' + token"
// @Success 200 {string} string	""
// @Router /api/v1/auth/images [delete]
func (q *ImageDelivery) Delete(c *gin.Context) {
	var agent model.Agent
	if v, ok := c.Get("agent"); ok {
		agent = v.(model.Agent)
	}
	var filter model.ImageFilter
	var image model.Image
	if c.BindQuery(&filter) == nil {
		copier.Copy(&image, &filter)
		nDel, err := q.useCase.Delete(agent, image)
		if err != nil {
			switch err.Error() {
			case "not found role":
				c.String(401, "not allow")
			case "delete image require at least one query":
				c.String(400, err.Error())
			default:
				c.String(500, err.Error())
			}
			return
		}
		if nDel == 0 {
			c.String(404, "not found")
			return
		}
		c.JSON(200, "")
		return
	} else {
		c.String(400, "require param image_id")
		return
	}
}

// @tags Images
// @Summary image
// @Description get list image
// @start_time default
// @Param _ query model.ImageFilter true "_"
// @Param Authorization header string true "'Bearer ' + token"
// @Param status query string false "status"
// @Success 200 {object} []model.ImageResponse
// @Router /api/v1/auth/images [get]
func (q *ImageDelivery) Gets(c *gin.Context) {
	var agent model.Agent
	if v, ok := c.Get("agent"); ok {
		agent = v.(model.Agent)
	}
	var filter model.ImageFilter
	var image model.Image
	if c.BindQuery(&filter) == nil {
		copier.Copy(&image, &filter)
		var arrImageResponse []model.ImageResponse
		err := q.useCase.GetsCustom(agent, image, &arrImageResponse)
		if err != nil {
			switch err.Error() {
			case "not found role":
				c.String(401, "not allow")
			default:
				c.String(500, err.Error())
			}
			return
		}
		if len(arrImageResponse) == 0 {
			c.String(404, "not found")
			return
		}
		c.JSON(200, arrImageResponse)
		return
	} else {
		fmt.Println(c.BindQuery(&filter))
		c.String(500, "...")
		return
	}
}

// @tags Images
// @Summary image
// @Description update image
// @start_time default
// @Param image_id path string true "image id"
// @Param Authorization header string true "'Bearer ' + token"
// @Param body body model.ImageUpdate true "image content"
// @Success 200 {string} string	""
// @Router /api/v1/auth/image/{image_id} [post]
func (q *ImageDelivery) UpdateById(c *gin.Context) {
	imageId := c.Param("image_id")
	var agent model.Agent
	var update model.ImageUpdate
	if v, ok := c.Get("agent"); ok {
		agent = v.(model.Agent)
	}
	err := c.BindJSON(&update)
	if err == nil {
		nModify, err := q.useCase.Update(agent, model.Image{Id: imageId}, model.Image{Data: update.Data})
		if err != nil {
			switch err.Error() {
			case "not found role":
				c.String(401, "not allow")
			default:
				c.String(500, err.Error())
			}
			return
		}
		if nModify == 0 {
			c.String(404, "not found")
			return
		}
		c.Writer.WriteHeader(200)
		return
	} else {
		c.String(500, "parser body failed, "+err.Error())
		return
	}
}

func (q *ImageDelivery) PublishTask(image model.ImageTask) error {
	data, err := json.Marshal(image)
	if err != nil {
		return err
	}
	err = q.rabbitmq.SendMessage(q.imageTaskQueue, data, 0)
	if err != nil {
		return err
	}
	return nil
}
func (q *ImageDelivery) HandleTaskSuccess(message []byte, messageAction *module.MessageAction) {
	var image model.Image
	err := json.Unmarshal(message, &image)
	if err != nil {
		fmt.Println(err)
		messageAction.Ack()
		return
	}
	if image.Id == "" {
		fmt.Println("[ERROR] empty id")
		messageAction.Ack()
		return
	}
	_, err = q.useCase.Update(model.Agent{UserId: image.UserId, Role: enum.RoleUser}, model.Image{Id: image.Id}, image)
	if err != nil {
		fmt.Println(err)
		messageAction.Reject()
		return
	}
	messageAction.Ack()
}
func (q *ImageDelivery) HandleTaskError(message []byte, messageAction *module.MessageAction) {
	q.HandleTaskSuccess(message, messageAction) //success and error same handle
}

func NewImageDelivery() (ImageInterface.IImageDelivery, error) {
	var q ImageDelivery
	var err error
	q.useCase, err = ImageUseCase.NewImageUseCase()
	if err != nil {
		return nil, err
	}
	rabbitmqLogin := module.RabbitMQLogin{
		Host:     CONFIG.GetString("RABBITMQ_HOST"),
		Port:     CONFIG.GetString("RABBITMQ_PORT"),
		Username: CONFIG.GetString("RABBITMQ_USERNAME"),
		Password: CONFIG.GetString("RABBITMQ_PASSWORD"),
		VHOST:    CONFIG.GetString("RABBITMQ_VHOST"),
	}
	q.rabbitmq, err = module.NewRabbitMQ(rabbitmqLogin)
	if err != nil {
		return nil, err
	}
	q.imageTaskQueue = CONFIG.GetString("IMAGE_TASK_QUEUE")
	q.imageSuccessQueue = CONFIG.GetString("IMAGE_SUCCESS_QUEUE")
	q.imageErrorQueue = CONFIG.GetString("IMAGE_ERROR_QUEUE")
	err = q.rabbitmq.CreateQueue(q.imageTaskQueue, 10)
	if err != nil {
		return nil, err
	}
	err = q.rabbitmq.CreateQueue(q.imageSuccessQueue, 0)
	if err != nil {
		return nil, err
	}
	err = q.rabbitmq.CreateQueue(q.imageErrorQueue, 0)
	if err != nil {
		return nil, err
	}
	var consumeSuccessQueue = module.Consume{
		q.imageSuccessQueue,
		"",
		false,
		false,
		false,
		false,
		5,
	}
	var consumeErrorQueue = module.Consume{
		q.imageErrorQueue,
		"",
		false,
		false,
		false,
		false,
		5,
	}
	q.rabbitmq.Consume(consumeSuccessQueue, q.HandleTaskSuccess)
	q.rabbitmq.Consume(consumeErrorQueue, q.HandleTaskError)
	return &q, nil
}
