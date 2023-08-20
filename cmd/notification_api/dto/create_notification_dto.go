package notification_dto

import "github.com/pakut2/mandarin/pkg/types"

type CreateNotificationDto struct {
	DeviceToken       string                 `json:"deviceToken" validate:"required" example:"fIuoGe66REq_eyZaN2V8E0"`
	PrecedenceMinutes int                    `json:"precedenceMinutes" validate:"gte=1,lte=60" example:"10"`
	ProviderName      types.ScheduleProvider `json:"providerName" validate:"supportedProvider" example:"ztm"`
	StopId            string                 `json:"stopId" validate:"required" example:"1461"`
	LineNumber        string                 `json:"lineNumber" validate:"required" example:"199"`
}
