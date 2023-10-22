package mqtt

import (
	"log"

	"github.com/Menschomat/pBox2/api/websocket"
	"github.com/Menschomat/pBox2/model"
	"github.com/Menschomat/pBox2/utils"
)

// handleFanEvent updates the fan's level and notifies clients if there's a change.
func (c *MQTTClient) handleFanEvent(box *model.Box, itemID string, payload []byte) {
	fan := utils.FindFanById(itemID, box)
	value, err := utils.GetIntValueFromPayload(payload)
	if err != nil {
		log.Printf("Error parsing payload: %v", err)
		return
	}

	if fan.Level != value {
		fan.Level = value
		websocket.PublishFanEvent(c.cfg, box, fan)
	}
}

// handleLightEvent updates the light's level and notifies clients if there's a change.
func (c *MQTTClient) handleLightEvent(box *model.Box, itemID string, payload []byte) {
	light := utils.FindLightById(itemID, box)
	value, err := utils.GetIntValueFromPayload(payload)
	if err != nil {
		log.Printf("Error parsing payload: %v", err)
		return
	}

	if light.Level != value {
		light.Level = value
		websocket.PublishLightEvent(c.cfg, box, light)
	}
}

// handleSensorEvent updates the sensor's data and notifies clients.
func (c *MQTTClient) handleSensorEvent(box *model.Box, itemID string, payload []byte) {
	sensor := utils.FindSensorById(itemID, box)
	value, err := utils.GetFloatValueFromPayload(payload)
	if err != nil {
		log.Printf("Error parsing payload: %v", err)
		return
	}

	utils.StoreValueInTimeSeries(float32(value), &sensor.TimeSeries)
	websocket.PublishSensorEvent(c.cfg, box, sensor, value)
}
