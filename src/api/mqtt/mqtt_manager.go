package mqtt

import (
	"log"
	"strings"

	"github.com/Menschomat/pBox2/model"
	"github.com/Menschomat/pBox2/utils"
	mqtt "github.com/eclipse/paho.mqtt.golang"
)

var cfg = utils.GetConfig()
var opts = utils.GetBrokerOpts(cfg, messagePubHandler, connectHandler, connectLostHandler)
var client = mqtt.NewClient(opts)
var messagePubHandler mqtt.MessageHandler = func(client mqtt.Client, msg mqtt.Message) {
	HandleMessage(&cfg, msg)
}

var connectHandler mqtt.OnConnectHandler = func(client mqtt.Client) {
	log.Println("Connected")
	sub(client)
}

var connectLostHandler mqtt.ConnectionLostHandler = func(client mqtt.Client, err error) {
	log.Printf("Connect lost: %v", err)
}

func sub(client mqtt.Client) {
	if token := client.Subscribe(cfg.Mqtt.Topic+"/#", 0, nil); token.Wait() && token.Error() != nil {
		log.Println(token.Error())
	}
}
func isValidTopic(cfg *model.Configuration, topic string) bool {
	return strings.HasPrefix(topic, cfg.Mqtt.Topic)
}

func logMessageReceived(msg mqtt.Message) {
	log.Printf("Received message: %s from topic: %s\n", msg.Payload(), msg.Topic())
}

func GetClient() mqtt.Client {
	return client
}

func HandleMessage(cfg *model.Configuration, msg mqtt.Message) {
	if !isValidTopic(cfg, msg.Topic()) {
		return
	}

	logMessageReceived(msg)

	boxID, itemType, itemID, err := utils.ParseTopic(msg.Topic())
	if err != nil {
		log.Println(err)
		return
	}

	box := utils.FindBoxById(boxID, &cfg.Enclosure)
	switch itemType {
	case "sensors":
		handleSensorEvent(cfg, box, itemID, msg.Payload())
	case "lights":
		handleLightEvent(cfg, box, itemID, msg.Payload())
	case "fans":
		handleFanEvent(cfg, box, itemID, msg.Payload())
	default:
		log.Println("UNKNOWN")
	}
}
