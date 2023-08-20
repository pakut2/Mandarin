package notification

import (
	"time"

	"github.com/pakut2/mandarin/pkg/types"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Notification struct {
	Id                primitive.ObjectID     `json:"id" bson:"_id,omitempty"`
	Delivered         bool                   `json:"delivered" bson:"delivered"`
	CreatedAt         time.Time              `json:"createdAt" bson:"createdAt"`
	DeviceToken       string                 `json:"deviceToken" bson:"deviceToken"`
	PrecedenceMinutes int                    `json:"precedenceMinutes" bson:"precedenceMinutes"`
	ProviderName      types.ScheduleProvider `json:"providerName" bson:"providerName"`
	StopId            string                 `json:"stopId" bson:"stopId"`
	LineNumber        string                 `json:"lineNumber" bson:"lineNumber"`
}
