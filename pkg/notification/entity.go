package notification

import (
	"time"

	"github.com/pakut2/mandarin/pkg/types"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Notification struct {
	Id           primitive.ObjectID         `json:"id" bson:"_id,omitempty"`
	Delivered    *bool                      `json:"delivered" bson:"delivered,omitempty"`
	CreatedAt    time.Time                  `json:"createdAt" bson:"createdAt,omitempty"`
	DeviceToken  string                     `json:"deviceToken" bson:"deviceToken,omitempty"`
	ReminderTime int                        `json:"reminderTime" bson:"reminderTime,omitempty"`
	ProviderName types.ScheduleProviderName `json:"providerName" bson:"providerName,omitempty"`
	StopId       string                     `json:"stopId" bson:"stopId,omitempty"`
	StopName     string                     `json:"stopName" bson:"stopName,omitempty"`
	LineNumber   string                     `json:"lineNumber" bson:"lineNumber,omitempty"`
}
