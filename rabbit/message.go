package rabbit

import (
	"encoding/json"
)

type Message struct {
	Mac string
	Message string
	Time string
}

func (m *Message) JsonToString() string {
	data, err := json.Marshal(m)
	if err != nil {
		return ""
	}
	
	return string(data)
}

func (m *Message) JsonFromString(jsonString string) error {
	return json.Unmarshal([]byte(jsonString), m)
}