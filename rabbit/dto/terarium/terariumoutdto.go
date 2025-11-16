package terarium

import "encoding/json"

type TerariumOutDto struct {
    TerariumDto
	Temperature int
	Humidity float32
}

func (t *TerariumOutDto) JsonToString () string{
	data, err := json.Marshal(t)
	if err != nil {
		return ""
	}
	
	return string(data)
}

func (t *TerariumOutDto) JsonFromString (jsonString string) error{
	return json.Unmarshal([]byte(jsonString), t)
}