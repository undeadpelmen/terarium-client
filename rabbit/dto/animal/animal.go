package animal

import "encoding/json"

type Animal struct {
	Id int
	Name string
	Lat_name string
	Day_max_t int
	Day_min_t int
	Night_max_t int
	Night_min_t int
	Uv_time float32
	Uv_spec  float32
	Uv_power int
	Humidity_max int
	Humidity_min int
	Day_len float32
	Feed_rate float32
	Kide_feed_rate float32
	Food string
	Vitamins string	
}

func (a *Animal) JsonToString() string {
	data, err := json.Marshal(a)
	if err != nil {
		return ""
	}
	
	return string(data)
}

func (a *Animal) JsonFromString(jsonString string) error {
	return json.Unmarshal([]byte(jsonString), a)
}
