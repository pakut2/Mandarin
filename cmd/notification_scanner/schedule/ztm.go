package schedule

import (
	"strconv"
	"strings"
	"time"

	"github.com/pakut2/mandarin/pkg/http_client"
	"github.com/pakut2/mandarin/pkg/logger"
	"github.com/pakut2/mandarin/pkg/notification"
)

const ZTM_TIMEZONE = "Europe/Warsaw"
const NIGHT_LINE_PREFIX = 'N'
const ZTM_NIGHT_LINE_PREFIX = '4'

func shouldDeliverZtmNotification(notification notification.Notification) bool {
	getZtmSchedulesEndpoint := getZtmSchedulesEndpoint(notification.LineNumber)
	schedules, err := http_client.Get(getZtmSchedulesEndpoint)

	if err != nil {
		logger.Logger.Errorf("error fething schedules from: %s, err: %v", getZtmSchedulesEndpoint, err)
		return false
	}

	stopIdNumber, err := strconv.ParseFloat(notification.StopId, 64)

	if err != nil {
		logger.Logger.Errorf("error parsing notification stopId: %s, err: %v", notification.StopId, err)
		return false
	}

	lineNumber, err := strconv.ParseFloat(parseZtmNightLineNumber(notification.LineNumber), 64)

	if err != nil {
		logger.Logger.Errorf("error parsing notification lineNumber: %s, err: %v", notification.LineNumber, err)
		return false
	}

	for _, schedule := range schedules["stopTimes"].([]interface{}) {
		if schedule.(map[string]interface{})["stopId"] == stopIdNumber && schedule.(map[string]interface{})["routeId"] == lineNumber {
			minutesToDeparture, err := getMinuteDifference(schedule.(map[string]interface{})["date"].(string), schedule.(map[string]interface{})["departureTime"].(string))

			if err != nil {
				continue
			}

			if minutesToDeparture == notification.ReminderTime {
				return true
			}
		}
	}

	return false
}

func getZtmSchedulesEndpoint(lineNumber string) string {
	currentDate := time.Now().UTC().Format("2006-01-02")

	return "https://ckan2.multimediagdansk.pl/stopTimes?date=" + currentDate + "&routeId=" + lineNumber
}

func parseZtmNightLineNumber(lineNumber string) string {
	if lineNumber[0] == ZTM_NIGHT_LINE_PREFIX {
		nightLineNumber := []rune(lineNumber)
		nightLineNumber[0] = rune(NIGHT_LINE_PREFIX)

		return string(nightLineNumber)
	}

	return lineNumber
}

func getMinuteDifference(ztmDate string, ztmIso string) (int, error) {
	currentDate := time.Now().UTC()
	ztmIsoParts := strings.Split(ztmIso, "T")
	ztmTime := ztmIsoParts[len(ztmIsoParts)-1]

	ztmIsoWithCorrectDate, err := time.Parse("2006-01-02T15:04:05", ztmDate+"T"+ztmTime)

	if err != nil {
		logger.Logger.Errorf("error converting date: %s, err: %v", ztmDate+"T"+ztmTime, err)
		return 0, err
	}

	ztmTimezone, err := time.LoadLocation(ZTM_TIMEZONE)

	if err != nil {
		logger.Logger.Errorf("error creating timezone, err: %v", err)
		return 0, err
	}

	ztmIsoReconciliated := time.Date(ztmIsoWithCorrectDate.Year(), ztmIsoWithCorrectDate.Month(), ztmIsoWithCorrectDate.Day(), ztmIsoWithCorrectDate.Hour(), ztmIsoWithCorrectDate.Minute(), 0, 0, ztmTimezone)

	return int(ztmIsoReconciliated.Sub(currentDate).Minutes()), nil
}
