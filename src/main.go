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
	_ "github.com/Menschomat/pBox2/docs"
	m "github.com/Menschomat/pBox2/model"
	router "github.com/Menschomat/pBox2/router"
	u "github.com/Menschomat/pBox2/utils"
	mqtt "github.com/eclipse/paho.mqtt.golang"
	"github.com/go-chi/chi/v5"
	httpSwagger "github.com/swaggo/http-swagger"
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

// GetFan godoc
//
//	@Summary		returns light
//	@Description	get light by box- and light-id
//	@Tags			light
//	@Accept			json
//	@Produce		json
//	@Param			boxId	path		string	true	"Box ID"
//	@Param			lightId	path		string	true	"Light ID"
//	@Success		200		{object}	m.Light
//	@Router			/{boxId}/lights/{lightId} [get]
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

// UpdateLight godoc
//
//	@Summary		updates light
//	@Description	get light by box- and light-id
//	@Tags			light
//	@Accept			json
//	@Produce		json
//	@Param			request	body		m.Light	true	"body"
//	@Success		200		{object}	m.Light
//	@Router			/{boxId}/lights/{lightId} [post]
func updateLight(w http.ResponseWriter, r *http.Request) {

}

// GetFan godoc
//
//	@Summary		returns fan
//	@Description	get fan by box- and fan-id
//	@Tags			fan
//	@Accept			json
//	@Produce		json
//	@Param			boxId	path		string	true	"Box ID"
//	@Param			fanId	path		string	true	"Fan ID"
//	@Success		200		{object}	m.Fan
//	@Router			/{boxId}/fans/{fanId} [get]
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

// UpdateFan godoc
//
//	@Summary		updates fan
//	@Description	get fan by box- and fan-id
//	@Tags			fan
//	@Accept			json
//	@Produce		json
//	@Param			request	body		m.Fan	true	"body"
//	@Success		200		{object}	m.Fan
//	@Router			/{boxId}/fans/{fanId} [post]
func updateFan(w http.ResponseWriter, r *http.Request) {

}

// GetEnclosure godoc
//
//	@Summary		returns whole enclosure
//	@Description	get string by ID
//	@Tags			enclosure
//	@Accept			json
//	@Produce		json
//	@Success		200	{object}	m.Enclosure
//	@Router			/enclosure [get]
func getEnclosure(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	cs.Publish([]byte("TEST"))
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(cfg.Enclosure)
}

// GetSensorData godoc
//
//	@Summary		returns sensor-data as time-series
//	@Description	get sensor-data as time-series by box- and sensor-id
//	@Tags			sensor
//	@Accept			json
//	@Produce		json
//	@Param			boxId		path		string	true	"Box ID"
//	@Param			sensorId	path		string	true	"Sensor ID"
//	@Success		200			{object}	m.TimeSeries
//	@Router			/{boxId}/sensors/{sensorId}/data [get]
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

// GetSensor godoc
//
//	@Summary		returns sensor
//	@Description	get sensor by box- and sensor-id
//	@Tags			sensor
//	@Accept			json
//	@Produce		json
//	@Param			boxId		path		string	true	"Box ID"
//	@Param			sensorId	path		string	true	"Sensor ID"
//	@Success		200			{object}	m.Sensor
//	@Router			/{boxId}/sensors/{sensorId} [get]
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

//	@title		pBox2 API-Docs
//	@version	1.0
//	@BasePath	/api/v1
func main() {
	log.Println("*_-_-_-pBox2-_-_-_*")
	rand.Seed(time.Now().UnixNano())
	if token := client.Connect(); token.Wait() && token.Error() != nil {
		panic(token.Error())
	}
	log.Println("Spawning API")
	var routes = []router.Route{
		{
			Method:      "GET",
			Path:        "/enclosure",
			HandlerFunc: getEnclosure,
		},
		{
			Method:      "GET",
			Path:        "/{boxId}/lights/{lightId}",
			HandlerFunc: getLight,
		},
		{
			Method:      "POST",
			Path:        "/{boxId}/lights/{lightId}",
			HandlerFunc: updateLight,
		},
		{
			Method:      "GET",
			Path:        "/{boxId}/fans/{fanId}",
			HandlerFunc: getLight,
		},
		{
			Method:      "POST",
			Path:        "/{boxId}/fans/{fanId}",
			HandlerFunc: updateLight,
		},
		{
			Method:      "GET",
			Path:        "/{boxId}/sensors/{sensorId}",
			HandlerFunc: getSensor,
		},
		{
			Method:      "GET",
			Path:        "/{boxId}/sensors/{sensorId}/data",
			HandlerFunc: getSensorData,
		},
	}
	r_api_v1 := router.NewRouter(routes)
	r := router.NewRouter([]router.Route{})
	r.Mount("/api/v1", r_api_v1)
	r.Mount("/swagger", httpSwagger.WrapHandler)
	r.Mount("/", cs)
	err := http.ListenAndServe(":8080", r)
	if err != nil {
		log.Fatalln("There's an error with the server", err)
	}
}
