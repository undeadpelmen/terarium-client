package rutine

import (
	"encoding/json"
)

type OutTerMsg struct {
    Temperarure int
	Humidity float32
}

func (m *OutTerMsg) JsonFromString (msg string) error {
    return json.Unmarshal([]byte(msg), m)
}

func (m *OutTerMsg) JsonToString () (string, error) {
    json, err := json.Marshal(m)
    return string(json), err
}