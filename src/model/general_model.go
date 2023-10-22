package model

type Configuration struct {
	Mqtt      Mqtt      `json:"mqtt"`
	Enclosure Enclosure `json:"enclosure"`
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
	Type  LightType `json:"type"`
	Level int       `json:"level"`
	State bool      `json:"state"`
}
type Switch struct {
	ID    string     `json:"id"`
	Name  string     `json:"name"`
	Type  SwitchType `json:"type"`
	State bool       `json:"state"`
}
type Fan struct {
	ID         string     `json:"id"`
	Name       string     `json:"name"`
	Level      int        `json:"level"`
	TimeSeries TimeSeries `json:"-"`
}

type Box struct {
	ID       string   `json:"id"`
	Name     string   `json:"name"`
	Location string   `json:"location"`
	Lights   []Light  `json:"lights"`
	Switches []Switch `json:"switches"`
	Fans     []Fan    `json:"fans"`
	Sensors  []Sensor `json:"sensors"`
}
type TimeSeries struct {
	Times  []string  `json:"times"`
	Values []float32 `json:"values"`
}
