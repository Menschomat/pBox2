package mqtt

import (
	"log"
	"strings"

	"github.com/Menschomat/pBox2/api/websocket"
	"github.com/Menschomat/pBox2/model"
	"github.com/Menschomat/pBox2/utils"
	mqtt "github.com/eclipse/paho.mqtt.golang"
)

// isTopicTasmota checks if the topic is related to Tasmota devices.
func (c *MQTTClient) isTopicTasmota(topic string) bool {
	return strings.HasPrefix(topic, c.cfg.Mqtt.Tasmota.Topic)
}

// processTasmotaMessage processes messages from Tasmota topics.
func (c *MQTTClient) processTasmotaMessage(msg mqtt.Message) {
	if !c.isTopicTasmota(msg.Topic()) {
		return
	}

	//log.Printf("Received message from Tasmota topic: %s: %s\n", msg.Topic(), msg.Payload())

	topicParts, err := utils.ParseTopic(msg.Topic())
	if err != nil {
		log.Printf("Error parsing topic: %v", err)
		return
	}

	if len(topicParts) < 3 {
		log.Println("Invalid topic structure")
		return
	}

	deviceId, method := topicParts[1], topicParts[2]
	switc, box := utils.FindSwitchByIdInEnc(deviceId, &c.cfg.Enclosure)
	if switc == nil || box == nil {
		log.Printf("Device or box not found for ID: %s", deviceId)
		return
	}

	switch method {
	case "POWER":
		c.handleSwitchEvent(switc, box, msg.Payload())
	default:
		//log.Printf("Unhandled method: %s", method)
	}
}

// handleSwitchEvent handles switch events from Tasmota devices.
func (c *MQTTClient) handleSwitchEvent(sw *model.Switch, box *model.Box, payload []byte) {
	stateStr := utils.GetStringFromPayload(payload)
	if stateStr == "" {
		log.Println("Empty payload")
		return
	}

	sw.State = StrToBool(stateStr)
	websocket.PublishSwitchEvent(c.cfg, box, sw)
}

// strToBool converts "on" and "off" strings to their boolean equivalents.
func StrToBool(s string) bool {
	stateMap := map[string]bool{"on": true, "off": false}
	return stateMap[strings.ToLower(s)]
}

// strToBool converts "on" and "off" strings to their boolean equivalents.
func BoolToStr(s bool) string {
	stateMap := map[bool]string{true: "on", false: "off"}
	return stateMap[s]
}
