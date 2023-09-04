package notification_scanner

import (
	"strconv"

	"github.com/pakut2/mandarin/cmd/notification_scanner/schedule"
	firebase_admin "github.com/pakut2/mandarin/pkg/firebase"
	"github.com/pakut2/mandarin/pkg/logger"
	notification_pkg "github.com/pakut2/mandarin/pkg/notification"
	"github.com/pakut2/mandarin/pkg/utilities"
)

func scanNotifications(notificationService notification_pkg.Service, messagingService firebase_admin.MessagingService) {
	logger.Logger.Info("Scanning notifications")

	notifications, err := notificationService.GetNotifications(notification_pkg.Notification{Delivered: utilities.BoolPointer(false)})
	if err != nil {
		return
	}

	for _, notification := range *notifications {
		scheduleProvider, err := schedule.GetScheduleProvider(notification.ProviderName)
		if err != nil {
			continue
		}

		if !scheduleProvider.ShouldDeliverNotification(notification) {
			continue
		}

		notificationBody := notification.LineNumber + " departs in " + strconv.Itoa(notification.ReminderTime) + "min"
		messagingService.SendMessage(notification.DeviceToken, notification.StopName, notificationBody)

		notificationService.UpdateNotification(notification.Id, notification_pkg.Notification{Delivered: utilities.BoolPointer(true)})
	}
}
