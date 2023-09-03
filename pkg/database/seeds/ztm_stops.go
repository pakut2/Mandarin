package seeds

import (
	"context"
	"fmt"
	"log"
	"sort"
	"strconv"
	"time"

	mapset "github.com/deckarep/golang-set/v2"
	stop_entities "github.com/pakut2/mandarin/cmd/schedule_provider_api/entity"
	"github.com/pakut2/mandarin/pkg/database"
	"github.com/pakut2/mandarin/pkg/http_client"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func SeedZtmStopsWithLineNumbers() error {
	stops, err := http_client.Get("https://ckan.multimediagdansk.pl/dataset/c24aa637-3619-4dc2-a171-a23eec8f2172/resource/d3e96eb6-25ad-4d6c-8651-b1eb39155945/download/stopsingdansk.json")
	if err != nil {
		log.Printf("[GetZtmStopLineNumbers] error fething stops: %s", err)
		return err
	}

	trips, err := http_client.Get("https://ckan.multimediagdansk.pl/dataset/c24aa637-3619-4dc2-a171-a23eec8f2172/resource/3115d29d-b763-4af5-93f6-763b835967d6/download/stopsintrip.json")
	if err != nil {
		log.Printf("[GetZtmStopLineNumbers] error fething trips: %s", err)
		return err
	}

	currentStops := stops["stops"].([]interface{})
	currentDate := time.Now().UTC().Format("2006-01-02")
	currentTrips := trips[currentDate].(map[string]interface{})["stopsInTrip"].([]interface{})

	collection := database.GetCollection("ztmStop")

	for _, stop := range currentStops {
		stopLineNumbers := mapset.NewSet[int]()

		for _, trip := range currentTrips {
			if trip.(map[string]interface{})["stopId"] == stop.(map[string]interface{})["stopId"] {
				stopLineNumbers.Add(int(trip.(map[string]interface{})["routeId"].(float64)))
			}
		}

		lineNumbersSetSlice := stopLineNumbers.ToSlice()
		sort.Ints(lineNumbersSetSlice)

		stringLineNumbers := make([]string, len(lineNumbersSetSlice))

		for i, intLineNumber := range lineNumbersSetSlice {
			stringLineNumbers[i] = parseZtmNightLineNumber(strconv.Itoa(intLineNumber))
		}

		collection.InsertOne(context.Background(), stop_entities.ZtmStop{Id: primitive.NewObjectID(), StopId: fmt.Sprintf("%v", stop.(map[string]interface{})["stopId"]), LineNumbers: stringLineNumbers})
	}

	log.Println("done...")

	return nil
}

func parseZtmNightLineNumber(lineNumber string) string {
	if lineNumber[0] == '4' && len(lineNumber) > 1 {
		nightLineNumber := []rune(lineNumber)
		nightLineNumber[0] = rune('N')

		return string(nightLineNumber)
	}

	return lineNumber
}
