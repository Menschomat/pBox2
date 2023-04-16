package main

import (
	"log"
	"math/rand"
	"net/http"
	"strings"
	"time"

	api_mqtt "github.com/Menschomat/pBox2/api/mqtt"
	api_rest "github.com/Menschomat/pBox2/api/rest"
	api_websocket "github.com/Menschomat/pBox2/api/websocket"
	_ "github.com/Menschomat/pBox2/docs"
	"github.com/Menschomat/pBox2/model"
	"github.com/Menschomat/pBox2/utils"
	mqtt "github.com/eclipse/paho.mqtt.golang"
	httpSwagger "github.com/swaggo/http-swagger"
)

const CFG_PATH = "config.json"

var websocket = api_websocket.NewWebSocketServer()
var cfg = utils.ParesConfig(CFG_PATH)
var opts = utils.GetBrokerOpts(cfg, messagePubHandler, connectHandler, connectLostHandler)
var client = mqtt.NewClient(opts)
var lightStates map[string]bool
var fanStates map[string]int

var messagePubHandler mqtt.MessageHandler = func(client mqtt.Client, msg mqtt.Message) {
	handleMessage(&cfg, msg)
}

func handleMessage(cfg *model.Configuration, msg mqtt.Message) {
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
		api_mqtt.HandleSensorEvent(cfg, box, itemID, msg.Payload())
	case "lights":
		api_mqtt.HandleLightEvent(cfg, box, itemID, msg.Payload())
	case "fans":
		api_mqtt.HandleFanEvent(cfg, box, itemID, msg.Payload())
	default:
		log.Println("UNKNOWN")
	}
}

func isValidTopic(cfg *model.Configuration, topic string) bool {
	return strings.HasPrefix(topic, cfg.Mqtt.Topic)
}

func logMessageReceived(msg mqtt.Message) {
	log.Printf("Received message: %s from topic: %s\n", msg.Payload(), msg.Topic())
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

// @title		pBox2 API-Docs
// @version	1.0
// @BasePath	/api/v1
func main() {
	log.Println("*_-_-_-pBox2-_-_-_*")
	rand.Seed(time.Now().UnixNano())
	if token := client.Connect(); token.Wait() && token.Error() != nil {
		panic(token.Error())
	}
	log.Println("Spawning API")
	appRouter := api_rest.NewBasicRouter()
	apiV1Router := api_rest.NewApiRouter(&cfg, client)
	appRouter.Mount("/api/v1", apiV1Router)
	appRouter.Mount("/swagger", httpSwagger.WrapHandler)
	appRouter.Mount("/", websocket)
	err := http.ListenAndServe(":8080", appRouter)
	if err != nil {
		log.Fatalln("There's an error with the server", err)
	}
}
