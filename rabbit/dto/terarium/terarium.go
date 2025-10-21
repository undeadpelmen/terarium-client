package terarium

import (
	"encoding/json"
	animal "terarium-client/rabbit/dto/animal"
)

type Tererarium struct {
	Id int
	Name string
	Mac string
	Animal animal.Animal
	AftorId int
}

func (t *Tererarium) JsonToString() string {
	data, err := json.Marshal(t)
	if err != nil {
		return ""
	}
	
	return string(data)
}

func (t *Tererarium) JsonFromString(jsonString string) error {
	return json.Unmarshal([]byte(jsonString), t)
}