package mqtt

import (
	"github.com/Menschomat/pBox2/api/websocket"
	_ "github.com/Menschomat/pBox2/docs"
	"github.com/Menschomat/pBox2/model"
	"github.com/Menschomat/pBox2/utils"
)

// HandleFanEvent handles MQTT messages related to fans.
// It updates the fan's level if the new value is different from the current value,
// and publishes a WebSocket event to notify clients of the change.
func HandleFanEvent(cfg *model.Configuration, box *model.Box, itemID string, payload []byte) {
	fan := utils.FindFanById(itemID, box)
	if value, err := utils.GetIntValueFromPayload(payload); err == nil {
		if fan.Level != value {
			fan.Level = value
			websocket.PublishFanEvent(cfg, box, fan, value)
		}
	}
}

// HandleLightEvent handles MQTT messages related to lights.
// It updates the light's level if the new value is different from the current value,
// and publishes a WebSocket event to notify clients of the change.
func HandleLightEvent(cfg *model.Configuration, box *model.Box, itemID string, payload []byte) {
	light := utils.FindLightById(itemID, box)
	if value, err := utils.GetIntValueFromPayload(payload); err == nil {
		if light.Level != value {
			light.Level = value
			websocket.PublishLightEvent(cfg, box, light, value)
		}
	}
}

// HandleSensorEvent handles MQTT messages related to sensors.
// It updates the sensor's time series with the new value, and publishes a WebSocket event
// to notify clients of the change.
func HandleSensorEvent(cfg *model.Configuration, box *model.Box, itemID string, payload []byte) {
	sensor := utils.FindSensorById(itemID, box)
	if value, err := utils.GetFloatValueFromPayload(payload); err == nil {
		utils.StoreValueInTimeSeries(float32(value), &sensor.TimeSeries)
		websocket.PublishSensorEvent(cfg, box, sensor, value)
	}
}
