package utils

import (
	"encoding/json"
	"io/ioutil"

	m "github.com/Menschomat/pBox2/model"
)

const CFG_PATH = "config.json"

var cfg = paresConfig(CFG_PATH)

func paresConfig(path string) m.Configuration {
	file, _ := ioutil.ReadFile(path)
	data := m.Configuration{}
	println([]byte(file))
	_ = json.Unmarshal([]byte(file), &data)
	return data
}

func GetConfig() m.Configuration {
	return cfg
}

func FindBoxById(id string, enclosure *m.Enclosure) *m.Box {
	for idx := range enclosure.Boxes {
		if enclosure.Boxes[idx].ID == id {
			return &enclosure.Boxes[idx]
		}
	}
	return &m.Box{}
}
func FindSwitchByIdInEnc(id string, enclosure *m.Enclosure) (*m.Switch, *m.Box) {
	for idx := range enclosure.Boxes {
		for s_idx := range enclosure.Boxes[idx].Switches {
			if enclosure.Boxes[idx].Switches[s_idx].ID == id {
				return &enclosure.Boxes[idx].Switches[s_idx], &enclosure.Boxes[idx]
			}
		}
	}
	return &m.Switch{}, &m.Box{}
}

func FindSwitchById(id string, box *m.Box) *m.Switch {
	println(id)
	println(box.Switches[0].ID)
	for s_idx := range box.Switches {
		if box.Switches[s_idx].ID == id {
			return &box.Switches[s_idx]
		}

	}
	return &m.Switch{}
}
func FindSensorById(id string, box *m.Box) *m.Sensor {
	for idx := range box.Sensors {
		if box.Sensors[idx].ID == id {
			return &box.Sensors[idx]
		}
	}
	return &m.Sensor{}
}
func FindLightById(id string, box *m.Box) *m.Light {
	for idx := range box.Lights {
		if box.Lights[idx].ID == id {
			return &box.Lights[idx]
		}
	}
	return &m.Light{}
}
func FindFanById(id string, box *m.Box) *m.Fan {
	for idx := range box.Fans {
		if box.Fans[idx].ID == id {
			return &box.Fans[idx]
		}
	}
	return &m.Fan{}
}
