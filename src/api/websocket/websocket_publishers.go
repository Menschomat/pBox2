package websocket

import (
	"encoding/json"
	"log"
	"math"
	"time"

	_ "github.com/Menschomat/pBox2/docs"
	"github.com/Menschomat/pBox2/model"
)

func PublishLightEvent(cfg *model.Configuration, box *model.Box, light *model.Light, level int) {
	event, err := json.Marshal(
		model.NewLightEvent(
			cfg.Enclosure.ID+"/"+box.ID,
			model.LightEventBody{
				ID:    light.ID,
				Level: level,
			},
		),
	)
	if err != nil {
		log.Println(err)
		return
	}
	wsServer.Publish(event)
}

func PublishFanEvent(cfg *model.Configuration, box *model.Box, fan *model.Fan, level int) {
	event, err := json.Marshal(
		model.NewFanEvent(
			cfg.Enclosure.ID+"/"+box.ID,
			model.FanEventBody{
				ID:    fan.ID,
				Level: level,
			},
		),
	)
	if err != nil {
		log.Println(err)
		return
	}
	wsServer.Publish(event)
}
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
