package model

type Mqtt struct {
	Broker   string      `json:"broker"`
	Port     int         `json:"port"`
	ClientID string      `json:"client_id"`
	Username string      `json:"username"`
	Password string      `json:"password"`
	Topic    string      `json:"topic"`
	Tasmota  MqttTasmota `json:"tasmota"`
}

type MqttMessage struct {
	Payload  interface{} `json:"payload"`
	ClientId string      `json:"clientId"`
}
type MqttTasmota struct {
	Topic string `json:"topic"`
}
