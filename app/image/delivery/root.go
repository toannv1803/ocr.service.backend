package ImageDelivery

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	ImageInterface "ocr.service.backend/app/image/interface"
	ImageRepository "ocr.service.backend/app/image/repository"
	"ocr.service.backend/config"
	"ocr.service.backend/model"
	module "ocr.service.backend/module/rabbitmq"
)

var CONFIG, _ = config.NewConfig(nil)

type ImageDelivery struct {
	repository        ImageInterface.IImageRepository
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
// @Param id query string false "image id"
// @Param user_id query string false "user id"
// @Param status query string false "status"
// @Success 200 {object} []model.Image
// @Router /api/v1/images [get]
func (q *ImageDelivery) Gets(c *gin.Context) {
	var filter model.Image
	if c.BindQuery(&filter) == nil {
		arrImage, err := q.repository.Get(filter)
		if err != nil {
			c.String(500, "get data failed")
			return
		}
		c.JSON(200, arrImage)
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
// @Param id query string false "image id"
// @Param data body string false "image content"
// @Success 200 {string} string	""
// @Router /api/v1/images [post]
func (q *ImageDelivery) Update(c *gin.Context) {
	imageID := c.Query("id")
	var update model.Image
	err := c.BindJSON(&update)
	if err == nil {
		err := q.repository.Update(model.Image{Id: imageID}, model.Image{Data: update.Data})
		if err != nil {
			c.String(500, "update data failed")
			return
		}
		c.Writer.WriteHeader(200)
		return
	} else {
		c.String(500, "parser body failed, "+err.Error())
		return
	}
}

func (q *ImageDelivery) PublishTask(image model.Image) error {
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
	err = q.repository.Update(model.Image{Id: image.Id}, image)
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

func NewImageDelivery() (*ImageDelivery, error) {
	var q ImageDelivery
	var err error
	q.repository, err = ImageRepository.NewImageRepository()
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
