package notification_dto

import "github.com/pakut2/mandarin/pkg/types"

type CreateNotificationDto struct {
	DeviceToken  string                     `json:"deviceToken" validate:"required" example:"fIuoGe66REq_eyZaN2V8E0"`
	ReminderTime int                        `json:"reminderTime" validate:"gte=1,lte=60" example:"10"`
	ProviderName types.ScheduleProviderName `json:"providerName" validate:"supportedProvider" example:"ztm"`
	StopId       string                     `json:"stopId" validate:"required,numeric" example:"1461"`
	StopName     string                     `json:"stopName" validate:"required" example:"Przymorze Wielkie"`
	LineNumber   string                     `json:"lineNumber" validate:"required" example:"199"`
}
