package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"strconv"
	"strings"
	"time"

	api "github.com/Menschomat/pBox2/api"
	m "github.com/Menschomat/pBox2/model"
	u "github.com/Menschomat/pBox2/utils"
	mqtt "github.com/eclipse/paho.mqtt.golang"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

const CFG_PATH = "config.json"

var cs = api.NewChatServer()
var cfg = u.ParesConfig(CFG_PATH)
var opts = u.GetBrokerOpts(cfg, messagePubHandler, connectHandler, connectLostHandler)
var client = mqtt.NewClient(opts)
var lightStates map[string]bool
var fanStates map[string]int

var messagePubHandler mqtt.MessageHandler = func(client mqtt.Client, msg mqtt.Message) {
	if strings.HasPrefix(msg.Topic(), cfg.Mqtt.Topic) {
		log.Printf("Received message: %s from topic: %s\n", msg.Payload(), msg.Topic())
		boxId, msgType, itemId, err := u.ParseTopic(msg.Topic())
		if err != nil {
			log.Println(err.Error())
			return
		}
		log.Println("BOX_ID:", boxId, "ITEM_ID:", itemId, "MSG_TYPE", msgType, "PAYLOAD:", string(msg.Payload()))
		box := u.FindBoxById(boxId, &cfg.Enclosure)
		switch msgType {
		case "sensors":
			sensor := u.FindSensorById(itemId, box)
			log.Println("HANDLING - SENSOR-EVENT:", sensor.ID)
			if f64, err := strconv.ParseFloat(string(msg.Payload()), 32); err == nil {
				go storeValueInTimeSeries(float32(f64), &sensor.TimeSeries)
				sensorEvent, err := json.Marshal(
					m.NewSensorEvent(
						cfg.Enclosure.ID+"/"+box.ID,
						m.SensorEventBody{
							ID:   sensor.ID,
							Unit: sensor.Unit, Type: sensor.Type,
							Value: f64,
							Time:  time.Now().Format(time.RFC3339),
						},
					),
				)
				if err != nil {
					fmt.Println(err)
					return
				}
				go cs.Publish(sensorEvent)
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

func storeValueInTimeSeries(value float32, timeSeries *m.TimeSeries) {
	timeSeries.Times = append(timeSeries.Times, time.Now().Format(time.RFC3339))
	timeSeries.Values = append(timeSeries.Values, value)
	if len(timeSeries.Times) > 10 {
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

func getLight(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	box := *u.FindBoxById(chi.URLParam(r, "boxId"), &cfg.Enclosure)
	light := *u.FindLightById(chi.URLParam(r, "lightId"), &box)
	if len(light.ID) <= 0 {
		BadRequestError(w, r)
	}
	json.NewEncoder(w).Encode(light)
}

func updateLight(w http.ResponseWriter, r *http.Request) {

}

func getFan(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	box := *u.FindBoxById(chi.URLParam(r, "boxId"), &cfg.Enclosure)
	fan := *u.FindFanById(chi.URLParam(r, "fanId"), &box)
	if len(fan.ID) <= 0 {
		BadRequestError(w, r)
	}
	json.NewEncoder(w).Encode(fan)
}

func updateFan(w http.ResponseWriter, r *http.Request) {

}

func getEnclosure(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	cs.Publish([]byte("TEST"))
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(cfg.Enclosure)
}

func getSensorData(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	box := *u.FindBoxById(chi.URLParam(r, "boxId"), &cfg.Enclosure)
	sensor := *u.FindSensorById(chi.URLParam(r, "sensorId"), &box)
	if len(sensor.ID) <= 0 {
		BadRequestError(w, r)
	}
	json.NewEncoder(w).Encode(sensor.TimeSeries)
}
func getSensor(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	box := *u.FindBoxById(chi.URLParam(r, "boxId"), &cfg.Enclosure)
	sensor := *u.FindSensorById(chi.URLParam(r, "sensorId"), &box)
	if len(sensor.ID) <= 0 {
		BadRequestError(w, r)
	}
	json.NewEncoder(w).Encode(sensor)
}

func main() {
	log.Println("*_-_-_-pBox2-_-_-_*")
	rand.Seed(time.Now().UnixNano())
	if token := client.Connect(); token.Wait() && token.Error() != nil {
		panic(token.Error())
	}
	log.Println("Spawning API")
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Get("/api/v1/enclosure", getEnclosure)
	//---------------
	r.Get("/api/v1/{boxId}/lights/{lightId}", getLight)
	r.Post("/api/v1/{boxId}/lights/{lightId}", updateLight)
	//---------------
	r.Get("/api/v1/{boxId}/fans/{fanId}", getFan)
	r.Post("/api/v1/{boxId}/fans/{fanId}", updateFan)
	//---------------
	r.Get("/api/v1/{boxId}/sensors/{sensorId}", getSensor)
	r.Get("/api/v1/{boxId}/sensors/{sensorId}/data", getSensorData)
	r.Mount("/", cs)
	err := http.ListenAndServe(":8080", r)
	if err != nil {
		log.Fatalln("There's an error with the server", err)
	}
}
