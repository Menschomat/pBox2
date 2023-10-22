package mqtt

import (
	"log"
	"strings"

	"github.com/Menschomat/pBox2/model"
	"github.com/Menschomat/pBox2/utils"
	mqtt "github.com/eclipse/paho.mqtt.golang"
)

// MQTTClient manages MQTT client operations.
type MQTTClient struct {
	client mqtt.Client
	cfg    *model.Configuration
}

// NewMQTTClient creates and configures a new MQTT client.
func NewMQTTClient() *MQTTClient {
	cfg := utils.GetConfig()
	mqttClient := &MQTTClient{
		cfg: &cfg,
	}

	opts := utils.GetBrokerOpts(cfg, mqttClient.onMessageReceived, mqttClient.onConnected, mqttClient.onConnectionLost)
	mqttClient.client = mqtt.NewClient(opts)

	return mqttClient
}

// onMessageReceived handles incoming MQTT messages.
func (c *MQTTClient) onMessageReceived(client mqtt.Client, msg mqtt.Message) {
	if c.isTopicValid(msg.Topic()) {
		c.processMessage(msg)
	} else if c.isTopicTasmota(msg.Topic()) {
		c.processTasmotaMessage(msg) // Assuming this method is implemented elsewhere
	} else {
		log.Printf("Received message from unknown topic: %s", msg.Topic())
	}
}

// onConnected handles actions to be taken when the client connects.
func (c *MQTTClient) onConnected(client mqtt.Client) {
	log.Println("MQTT client connected")
	c.subscribeToTopics()
}

// onConnectionLost handles actions to be taken when the client loses connection.
func (c *MQTTClient) onConnectionLost(client mqtt.Client, err error) {
	log.Printf("MQTT connection lost: %v", err)
	// Consider reconnection logic here
}

// subscribeToTopics subscribes the client to relevant topics.
func (c *MQTTClient) subscribeToTopics() {
	topics := c.getDefaultTopics()
	token := c.client.SubscribeMultiple(topics, nil)
	if token.Wait() && token.Error() != nil {
		log.Printf("Error subscribing to topics: %v", token.Error())
		return
	}
	log.Println("Subscribed to topics successfully")
}

// getDefaultTopics prepares the default topics for subscription.
func (c *MQTTClient) getDefaultTopics() map[string]byte {
	topics := map[string]byte{
		c.cfg.Mqtt.Topic + "/#": 0,
	}
	if tasmotaTopic := c.cfg.Mqtt.Tasmota.Topic; tasmotaTopic != "" {
		topics[tasmotaTopic+"/#"] = 0
	}
	return topics
}

// isTopicValid checks if the topic is within the expected structure.
func (c *MQTTClient) isTopicValid(topic string) bool {
	return strings.HasPrefix(topic, c.cfg.Mqtt.Topic)
}

// processMessage processes and routes the payload of a valid message.
func (c *MQTTClient) processMessage(msg mqtt.Message) {
	log.Printf("Processing message from topic: %s", msg.Topic())
	parsedTopic, err := utils.ParseTopic(msg.Topic())
	if err != nil {
		log.Printf("Error parsing topic: %v", err)
		return
	}

	if len(parsedTopic) < 3 {
		log.Println("Invalid topic structure")
		return
	}

	boxID, itemType, itemID := parsedTopic[0], parsedTopic[1], parsedTopic[2]
	box := utils.FindBoxById(boxID, &c.cfg.Enclosure)
	if box == nil {
		log.Printf("Box not found for ID: %s", boxID)
		return
	}

	c.routeMessageByItemType(itemType, box, itemID, msg.Payload())
}

// routeMessageByItemType routes the message to the appropriate handler based on the item type.
func (c *MQTTClient) routeMessageByItemType(itemType string, box *model.Box, itemID string, payload []byte) {
	switch itemType {
	case "sensors":
		c.handleSensorEvent(box, itemID, payload)
	case "lights":
		c.handleLightEvent(box, itemID, payload)
	case "fans":
		c.handleFanEvent(box, itemID, payload)
	default:
		log.Printf("Unhandled item type: %s", itemType)
	}
}

// GetClient exposes the underlying MQTT client.
func (c *MQTTClient) GetClient() mqtt.Client {
	return c.client
}
