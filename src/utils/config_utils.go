package utils

import (
	"encoding/json"
	"io/ioutil"

	m "github.com/Menschomat/pBox2/model"
)

func ParesConfig(path string) m.Configuration {
	file, _ := ioutil.ReadFile(path)
	data := m.Configuration{}
	_ = json.Unmarshal([]byte(file), &data)
	return data
}

func FindBoxById(id string, enclosure *m.Enclosure) *m.Box {
	for idx := range enclosure.Boxes {
		if enclosure.Boxes[idx].ID == id {
			return &enclosure.Boxes[idx]
		}
	}
	return &m.Box{}
}
func FindSensorById(id string, box *m.Box) *m.Sensor {
	for idx := range box.Sensors {
		if box.Sensors[idx].ID == id {
			return &box.Sensors[idx]
		}
	}
	return &m.Sensor{}
}
