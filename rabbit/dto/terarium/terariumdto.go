package terarium

import (
	"terarium-client/rabbit"
	"terarium-client/rabbit/dto/animal"
)

type TerariumDto struct {
	rabbit.Message
	TerariumId int
	Animal animal.Animal
}