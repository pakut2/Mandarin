package notification_scanner

import (
	"time"

	firebase "firebase.google.com/go/v4"
	"github.com/go-co-op/gocron"
	firebase_admin "github.com/pakut2/mandarin/pkg/firebase"
	"github.com/pakut2/mandarin/pkg/notification"
)

func InitScanner(firebaseAdmin *firebase.App) {
	notificationService := notification.NewService()
	messagingService, err := firebase_admin.NewMessagingService(firebaseAdmin)

	if err != nil {
		panic(err)
	}

	scheduler := gocron.NewScheduler(time.UTC)

	scheduler.Cron("* * * * *").Do(func() {
		scanNotifications(notificationService, messagingService)
	})

	scheduler.StartAsync()
}
