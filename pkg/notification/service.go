package notification

import (
	"context"
	"time"

	notification_dto "github.com/pakut2/mandarin/cmd/notification_api/dto"
	"github.com/pakut2/mandarin/pkg/database"
	"github.com/pakut2/mandarin/pkg/logger"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type Service interface {
	CreateNotification(*notification_dto.CreateNotificationDto) (*Notification, error)
	GetNotifications(GetNotificationsFilter) (*[]Notification, error)
	UpdateNotification(primitive.ObjectID, UpdateNotificationData) error
}

type service struct {
	collection *mongo.Collection
}

func NewService() Service {
	return &service{
		collection: database.GetCollection("notification"),
	}
}

func (s *service) CreateNotification(createNotificationDto *notification_dto.CreateNotificationDto) (*Notification, error) {
	notification := Notification{
		Id:           primitive.NewObjectID(),
		Delivered:    false,
		CreatedAt:    time.Now(),
		DeviceToken:  createNotificationDto.DeviceToken,
		ReminderTime: createNotificationDto.ReminderTime,
		ProviderName: createNotificationDto.ProviderName,
		StopId:       createNotificationDto.StopId,
		StopName:     createNotificationDto.StopName,
		LineNumber:   createNotificationDto.LineNumber,
	}

	if _, err := s.collection.InsertOne(context.Background(), notification); err != nil {
		logger.Logger.Errorf("error creating notification, err: %v", err)
		return nil, err
	}

	return &notification, nil
}

type GetNotificationsFilter struct {
	Delivered bool `bson:"delivered"`
}

func (s *service) GetNotifications(getNotificationsFilter GetNotificationsFilter) (*[]Notification, error) {
	getNotificationsFilterDoc, err := database.ToDoc(getNotificationsFilter)

	if err != nil {
		logger.Logger.Errorf("error parsing notifications filter data: %v, err: %v", getNotificationsFilter, err)
		return nil, err
	}

	cursor, err := s.collection.Find(context.Background(), getNotificationsFilterDoc)

	if err != nil {
		logger.Logger.Errorf("error fetching notifications, err: %v", err)
		return nil, err
	}

	notifications := make([]Notification, 0)

	if err = cursor.All(context.TODO(), &notifications); err != nil {
		logger.Logger.Errorf("error parsing notifications, err: %v", err)
		return nil, err
	}

	return &notifications, nil
}

type UpdateNotificationData struct {
	Delivered bool `bson:"delivered"`
}

func (s *service) UpdateNotification(notificationId primitive.ObjectID, updateNotificationData UpdateNotificationData) error {
	updateNotificationDataDoc, err := database.ToDoc(updateNotificationData)

	if err != nil {
		logger.Logger.Errorf("error parsing notification update data: %v, for notification with ID: %v, err: %v", updateNotificationData, notificationId, err)
		return err
	}

	if _, err = s.collection.UpdateOne(
		context.Background(),
		bson.M{"_id": notificationId},
		bson.M{"$set": updateNotificationDataDoc},
	); err != nil {
		logger.Logger.Errorf("error updating notification with ID: %v, err: %v", notificationId, err)
		return err
	}

	return nil
}
