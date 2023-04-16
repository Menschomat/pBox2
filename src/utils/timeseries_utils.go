package utils

import (
	"time"

	"github.com/Menschomat/pBox2/model"
)

func StoreValueInTimeSeries(value float32, timeSeries *model.TimeSeries) {
	timeSeries.Times = append(timeSeries.Times, time.Now().Format(time.RFC3339))
	timeSeries.Values = append(timeSeries.Values, value)
	if len(timeSeries.Times) > 200 {
		timeSeries.Times = timeSeries.Times[1:]
		timeSeries.Values = timeSeries.Values[1:]
	}
}
