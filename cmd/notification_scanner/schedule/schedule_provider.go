package schedule

import (
	"errors"

	"github.com/pakut2/mandarin/pkg/constants"
	"github.com/pakut2/mandarin/pkg/logger"
	"github.com/pakut2/mandarin/pkg/notification"
	"github.com/pakut2/mandarin/pkg/types"
)

type ScheduleProvider struct {
	ShouldDeliverNotification func(notification notification.Notification) bool
}

func GetScheduleProvider(providerName types.ScheduleProviderName) (ScheduleProvider, error) {
	switch providerName {
	case constants.ZTM:
		return ScheduleProvider{ShouldDeliverNotification: shouldDeliverZtmNotification}, nil
	default:
		logger.Logger.Errorf("unsuported provider: %s", providerName)
		return ScheduleProvider{}, errors.New("unsuported provider")
	}
}
