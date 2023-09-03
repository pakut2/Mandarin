package schedule_provider_api

import "go.mongodb.org/mongo-driver/bson/primitive"

type ZtmStop struct {
	Id          primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	StopId      string             `json:"stopId" bson:"stopId"`
	LineNumbers []string           `json:"lineNumbers" bson:"lineNumbers"`
}
