package main

import (
	"log"
	"math/rand"
	"net/http"
	"time"

	api_mqtt "github.com/Menschomat/pBox2/api/mqtt"
	api_rest "github.com/Menschomat/pBox2/api/rest"
	api_websocket "github.com/Menschomat/pBox2/api/websocket"
	_ "github.com/Menschomat/pBox2/docs"
	"github.com/Menschomat/pBox2/utils"
	httpSwagger "github.com/swaggo/http-swagger"
)

var websocket = api_websocket.NewWebSocketServer()
var cfg = utils.GetConfig()
var client = api_mqtt.NewMQTTClient()

// @title		pBox2 API-Docs
// @version	1.0
// @BasePath	/api/v1
func main() {
	log.Println("*_-_-_-pBox2-_-_-_*")
	rand.Seed(time.Now().UnixNano())
	if token := client.GetClient().Connect(); token.Wait() && token.Error() != nil {
		panic(token.Error())
	}
	log.Println("Spawning API")
	appRouter := api_rest.NewBasicRouter()
	apiV1Router := api_rest.NewApiRouter(&cfg, client.GetClient())
	appRouter.Mount("/api/v1", apiV1Router)
	appRouter.Mount("/swagger", httpSwagger.WrapHandler)
	appRouter.Mount("/ws", websocket)
	fs := http.FileServer(http.Dir("./static/dashboard"))
	appRouter.Handle("/*", http.StripPrefix("", fs))
	err := http.ListenAndServe(":8080", appRouter)
	if err != nil {
		log.Fatalln("There's an error with the server", err)
	}
}
