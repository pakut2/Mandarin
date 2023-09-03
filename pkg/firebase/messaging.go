package firebase_admin

import (
	"context"

	firebase "firebase.google.com/go/v4"
	"firebase.google.com/go/v4/messaging"
	"github.com/pakut2/mandarin/pkg/logger"
)

type MessagingService interface {
	SendMessage(string, string, string) error
}

type messagingService struct {
	firebaseMessaging *messaging.Client
}

func NewMessagingService(firebaseAdmin *firebase.App) (MessagingService, error) {
	messaging, err := firebaseAdmin.Messaging(context.Background())

	if err != nil {
		logger.Logger.Errorf("error getting Messaging client, err: %v", err)
		return nil, err
	}

	return &messagingService{
		firebaseMessaging: messaging,
	}, nil
}

func (s *messagingService) SendMessage(deviceToken string, title string, payload string) error {
	messageId, err := s.firebaseMessaging.Send(context.Background(), &messaging.Message{Token: deviceToken, Notification: &messaging.Notification{Title: title, Body: payload}, Android: &messaging.AndroidConfig{Notification: &messaging.AndroidNotification{Priority: messaging.PriorityHigh}}})

	if err != nil {
		logger.Logger.Errorf("error sending message, err: %v", err)
		return err
	}

	logger.Logger.Infof("message with ID: %s sent successfully", messageId)
	return nil
}
