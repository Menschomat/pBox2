package model

import (
	"time"
)

// General-------------------------------------------------------------
type SocketEvent struct {
	Time  string    `json:"time"`
	Type  EventType `json:"type"`
	Topic string    `json:"topic"`
}

type EventType string

const (
	SENSOR EventType = "SENSOR"
	LIGHT  EventType = "LIGHT"
	SWITCH EventType = "SWITCH"
	FAN    EventType = "FAN"
)

// Sensor-Event-------------------------------------------------------------
type SensorEvent struct {
	SocketEvent
	Body SensorEventBody `json:"body"`
}
type SensorEventBody struct {
	ID    string     `json:"id"`
	Type  SensorType `json:"type"`
	Unit  string     `json:"unit"`
	Value float64    `json:"value"`
	Time  string     `json:"time"`
}

func NewSensorEvent(topic string, body SensorEventBody) *SensorEvent {
	return &SensorEvent{
		SocketEvent{Topic: topic, Time: time.Now().Format(time.RFC3339), Type: SENSOR},
		body,
	}
}

// SwitchEvent-------------------------------------------------------------
type SwitchEvent struct {
	SocketEvent
	Body SwitchEventBody `json:"body"`
}
type SwitchEventBody struct {
	ID    string `json:"id"`
	State bool   `json:"state"`
}

func NewSwitchEvent(topic string, body SwitchEventBody) *SwitchEvent {
	return &SwitchEvent{
		SocketEvent{Topic: topic, Time: time.Now().Format(time.RFC3339), Type: SWITCH},
		body,
	}
}

// LightEvent-------------------------------------------------------------
type LightEvent struct {
	SocketEvent
	Body LightEventBody `json:"body"`
}
type LightEventBody struct {
	ID    string `json:"id"`
	Level int    `json:"value"`
}

func NewLightEvent(topic string, body LightEventBody) *LightEvent {
	return &LightEvent{
		SocketEvent{Topic: topic, Time: time.Now().Format(time.RFC3339), Type: LIGHT},
		body,
	}
}

// LightEvent-------------------------------------------------------------
type FanEvent struct {
	SocketEvent
	Body FanEventBody `json:"body"`
}
type FanEventBody struct {
	ID    string `json:"id"`
	Level int    `json:"value"`
}

func NewFanEvent(topic string, body FanEventBody) *FanEvent {
	return &FanEvent{
		SocketEvent{Topic: topic, Time: time.Now().Format(time.RFC3339), Type: FAN},
		body,
	}
}
