package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math"
	"math/rand"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/Menschomat/pBox2/api"
	_ "github.com/Menschomat/pBox2/docs"
	"github.com/Menschomat/pBox2/model"
	"github.com/Menschomat/pBox2/utils"
	mqtt "github.com/eclipse/paho.mqtt.golang"
	httpSwagger "github.com/swaggo/http-swagger"
)

const CFG_PATH = "config.json"

var websocket = api.NewWebSocketServer()
var cfg = utils.ParesConfig(CFG_PATH)
var opts = utils.GetBrokerOpts(cfg, messagePubHandler, connectHandler, connectLostHandler)
var client = mqtt.NewClient(opts)
var lightStates map[string]bool
var fanStates map[string]int

var messagePubHandler mqtt.MessageHandler = func(client mqtt.Client, msg mqtt.Message) {
	if strings.HasPrefix(msg.Topic(), cfg.Mqtt.Topic) {
		log.Printf("Received message: %s from topic: %s\n", msg.Payload(), msg.Topic())
		boxId, msgType, itemId, err := utils.ParseTopic(msg.Topic())
		if err != nil {
			log.Println(err.Error())
			return
		}
		log.Println("BOX_ID:", boxId, "ITEM_ID:", itemId, "MSG_TYPE", msgType, "PAYLOAD:", string(msg.Payload()))
		box := utils.FindBoxById(boxId, &cfg.Enclosure)
		switch msgType {
		case "sensors":
			sensor := utils.FindSensorById(itemId, box)
			log.Println("HANDLING - SENSOR-EVENT:", sensor.ID)
			if f32, err := strconv.ParseFloat(string(msg.Payload()), 32); err == nil {
				go storeValueInTimeSeries(float32(f32), &sensor.TimeSeries)
				sensorEvent, err := json.Marshal(
					model.NewSensorEvent(
						cfg.Enclosure.ID+"/"+box.ID,
						model.SensorEventBody{
							ID:   sensor.ID,
							Unit: sensor.Unit, Type: sensor.Type,
							Value: math.Round(f32*100) / 100,
							Time:  time.Now().Format(time.RFC3339),
						},
					),
				)
				if err != nil {
					fmt.Println(err)
					return
				}
				go websocket.Publish(sensorEvent)
			}
		case "light":
			log.Println("LIGHT")
		case "fans":
			log.Println("FANS")
		default:
			log.Println("UNKNOWN")
			return
		}

	}
}

func storeValueInTimeSeries(value float32, timeSeries *model.TimeSeries) {
	timeSeries.Times = append(timeSeries.Times, time.Now().Format(time.RFC3339))
	timeSeries.Values = append(timeSeries.Values, value)
	if len(timeSeries.Times) > 200 {
		timeSeries.Times = timeSeries.Times[1:]
		timeSeries.Values = timeSeries.Values[1:]
	}
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
	appRouter := api.NewBasicRouter()
	apiV1Router := api.NewApiRouter(&cfg, client)
	appRouter.Mount("/api/v1", apiV1Router)
	appRouter.Mount("/swagger", httpSwagger.WrapHandler)
	appRouter.Mount("/", websocket)
	err := http.ListenAndServe(":8080", appRouter)
	if err != nil {
		log.Fatalln("There's an error with the server", err)
	}
}
