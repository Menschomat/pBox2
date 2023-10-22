package model

type Sensor struct {
	ID         string     `json:"id"`
	Name       string     `json:"name"`
	Pin        int        `json:"pin"`
	Type       SensorType `json:"type"`
	Unit       string     `json:"unit"`
	Target     int        `json:"target"`
	TimeSeries TimeSeries `json:"-"`
}

type SensorType string

const (
	TEMP SensorType = "TEMP"
)

type LightType string

const (
	MONO LightType = "MONO"
)

type SwitchType string

const (
	TASMOTA_MQTT SwitchType = "TASMOTA_MQTT"
)
