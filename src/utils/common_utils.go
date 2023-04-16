package utils

import (
	"encoding/binary"
	"errors"
	"fmt"
	"math"
	"strconv"
	"strings"

	m "github.com/Menschomat/pBox2/model"
	mqtt "github.com/eclipse/paho.mqtt.golang"
)

func GetBrokerOpts(
	conf m.Configuration,
	msgHandler mqtt.MessageHandler,
	conHandler mqtt.OnConnectHandler,
	lostHandler mqtt.ConnectionLostHandler,
) *mqtt.ClientOptions {
	var broker = conf.Mqtt.Broker
	var port = conf.Mqtt.Port
	opts := mqtt.NewClientOptions()
	opts.AddBroker(fmt.Sprintf("tcp://%s:%d", broker, port))
	opts.SetClientID(conf.Mqtt.ClientID)
	opts.SetUsername(conf.Mqtt.Username)
	opts.SetPassword(conf.Mqtt.Password)
	opts.SetDefaultPublishHandler(msgHandler)
	opts.OnConnect = conHandler
	opts.OnConnectionLost = lostHandler
	return opts
}

func ParseTopic(topic string) (string, string, string, error) {
	s := strings.Split(topic, "/")
	if len(s) != 4 {
		return "", "", "", errors.New("Unable to parse topic!")
	}
	return s[1], s[2], s[3], nil
}

func Float64frombytes(bytes []byte) float64 {
	bits := binary.LittleEndian.Uint64(bytes)
	float := math.Float64frombits(bits)
	return float
}
func GetFloatValueFromPayload(payload []byte) (float64, error) {
	return strconv.ParseFloat(string(payload), 32)
}

func GetIntValueFromPayload(payload []byte) (int, error) {
	value, err := strconv.ParseInt(string(payload), 0, 32)
	return int(value), err
}
