package schedule_provider_api

import (
	"context"

	entities "github.com/pakut2/mandarin/cmd/schedule_provider_api/entity"
	"github.com/pakut2/mandarin/pkg/database"
	"github.com/pakut2/mandarin/pkg/logger"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type ZtmService interface {
	GetStopById(string) (*entities.ZtmStop, error)
}

type ztmService struct {
	collection *mongo.Collection
}

func NewZtmService() ZtmService {
	return &ztmService{
		collection: database.GetCollection("ztmStop"),
	}
}

func (s *ztmService) GetStopById(stopId string) (*entities.ZtmStop, error) {
	var stop entities.ZtmStop
	if err := s.collection.FindOne(context.Background(), bson.M{"stopId": stopId}).Decode(&stop); err != nil {
		logger.Logger.Errorf("error fetching stop by ID: %s, err: %v", stopId, err)
		return nil, err
	}

	return &stop, nil
}
