package websocket

import (
	"encoding/json"
	"log"
	"math"
	"time"

	_ "github.com/Menschomat/pBox2/docs"
	"github.com/Menschomat/pBox2/model"
)

// PublishLightEvent publishes a WebSocket event containing the current level
// of the given light to all connected clients.
func PublishSwitchEvent(cfg *model.Configuration, box *model.Box, switc *model.Switch) {
	event, err := json.Marshal(
		model.NewSwitchEvent(
			cfg.Enclosure.ID+"/"+box.ID,
			model.SwitchEventBody{
				ID:    switc.ID,
				State: switc.State,
			},
		),
	)
	if err != nil {
		log.Println(err)
		return
	}
	wsServer.Publish(event)
}

// PublishLightEvent publishes a WebSocket event containing the current level
// of the given light to all connected clients.
func PublishLightEvent(cfg *model.Configuration, box *model.Box, light *model.Light) {
	event, err := json.Marshal(
		model.NewLightEvent(
			cfg.Enclosure.ID+"/"+box.ID,
			model.LightEventBody{
				ID:    light.ID,
				Level: light.Level,
			},
		),
	)
	if err != nil {
		log.Println(err)
		return
	}
	wsServer.Publish(event)
}

// PublishFanEvent publishes a WebSocket event containing the current level
// of the given fan to all connected clients.
func PublishFanEvent(cfg *model.Configuration, box *model.Box, fan *model.Fan) {
	event, err := json.Marshal(
		model.NewFanEvent(
			cfg.Enclosure.ID+"/"+box.ID,
			model.FanEventBody{
				ID:    fan.ID,
				Level: fan.Level,
			},
		),
	)
	if err != nil {
		log.Println(err)
		return
	}
	wsServer.Publish(event)
}

// PublishSensorEvent publishes a WebSocket event containing the current value
// of the given sensor to all connected clients.
func PublishSensorEvent(cfg *model.Configuration, box *model.Box, sensor *model.Sensor, value float64) {
	event, err := json.Marshal(
		model.NewSensorEvent(
			cfg.Enclosure.ID+"/"+box.ID,
			model.SensorEventBody{
				ID:    sensor.ID,
				Unit:  sensor.Unit,
				Type:  sensor.Type,
				Value: math.Round(value*100) / 100,
				Time:  time.Now().Format(time.RFC3339),
			},
		),
	)
	if err != nil {
		log.Println(err)
		return
	}
	wsServer.Publish(event)
}
