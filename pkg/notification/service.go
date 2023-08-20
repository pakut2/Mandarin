package notification

import (
	"context"
	"time"

	notification_dto "github.com/pakut2/mandarin/cmd/notification_api/dto"
	"github.com/pakut2/mandarin/pkg/database"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type Service interface {
	CreateNotification(notification *notification_dto.CreateNotificationDto) (*Notification, error)
}

type service struct {
	Collection *mongo.Collection
}

func NewService() Service {
	return &service{
		Collection: database.GetCollection("notification"),
	}
}

func (s *service) CreateNotification(createNotificationDto *notification_dto.CreateNotificationDto) (*Notification, error) {
	notification := Notification{
		Id:                primitive.NewObjectID(),
		Delivered:         false,
		CreatedAt:         time.Now(),
		DeviceToken:       createNotificationDto.DeviceToken,
		PrecedenceMinutes: createNotificationDto.PrecedenceMinutes,
		ProviderName:      createNotificationDto.ProviderName,
		StopId:            createNotificationDto.StopId,
		LineNumber:        createNotificationDto.LineNumber,
	}

	_, err := s.Collection.InsertOne(context.Background(), notification)
	if err != nil {
		return nil, err
	}

	return &notification, nil
}
