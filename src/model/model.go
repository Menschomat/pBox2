package model

type Configuration struct {
	Mqtt      Mqtt      `json:"mqtt"`
	Enclosure Enclosure `json:"enclosure"`
}

type Mqtt struct {
	Broker   string `json:"broker"`
	Port     int    `json:"port"`
	ClientID string `json:"client_id"`
	Username string `json:"username"`
	Password string `json:"password"`
	Topic    string `json:"topic"`
}

type Enclosure struct {
	ID       string `json:"id"`
	Name     string `json:"name"`
	Location string `json:"location"`
	Boxes    []Box  `json:"boxes"`
}
type Light struct {
	ID    string    `json:"id"`
	Name  string    `json:"name"`
	Pins  []int     `json:"pins"`
	Type  LightType `json:"type"`
	State bool      `json:"state"`
}
type Fan struct {
	ID         string     `json:"id"`
	Name       string     `json:"name"`
	Pin        int        `json:"pin"`
	Level      int        `json:"level"`
	TimeSeries TimeSeries `json:"-"`
}
type Sensor struct {
	ID         string     `json:"id"`
	Name       string     `json:"name"`
	Pin        int        `json:"pin"`
	Type       SensorType `json:"type"`
	Unit       string     `json:"unit"`
	Target     int        `json:"target"`
	TimeSeries TimeSeries `json:"-"`
}
type Box struct {
	ID       string   `json:"id"`
	Name     string   `json:"name"`
	Location string   `json:"location"`
	Lights   []Light  `json:"lights"`
	Fans     []Fan    `json:"fans"`
	Sensors  []Sensor `json:"sensors"`
}
type TimeSeries struct {
	Times  []string  `json:"times"`
	Values []float32 `json:"values"`
}

type SensorType string

const (
	TEMP SensorType = "TEMP"
)

type LightType string

const (
	MONO LightType = "MONO"
	//RGB    LightType = "RGB"
	//WS2811 LightType = "WS2811"
)
