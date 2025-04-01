package usecase

import (
	"time"

	"github.com/pkg/errors"
)

func validateTimestamp(fromStr string, toStr string) (time.Time, time.Time, error) {
	layout := "2006-01-02T15:04:05"

	fromTime, err := time.ParseInLocation(layout, fromStr, time.UTC)
	if err != nil {
		return time.Time{}, time.Time{}, errors.Wrap(err, "Parse fromTime")
	}

	toTime, err := time.ParseInLocation(layout, toStr, time.UTC)
	if err != nil {
		return time.Time{}, time.Time{}, errors.Wrap(err, "Parse toTime")
	}

	if fromTime.After(toTime) {
		return time.Time{}, time.Time{}, errors.New("wrong time")
	}

	return fromTime, toTime, nil
}
