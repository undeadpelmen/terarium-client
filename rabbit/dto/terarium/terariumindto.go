package terarium

import (
	"encoding/json"
	"terarium-client/rabbit/dto/animal"
)

type TerariumInDto struct {
	TerariumDto
	Animal animal.Animal
}

func (t *TerariumInDto) JsonToString () string {
	data, err := json.Marshal(t)
	if err != nil {
		return ""
	}
	
	return string(data)
}

func (t *TerariumInDto) JsonFromString (jsonString string) error {
	return json.Unmarshal([]byte(jsonString), t)
}