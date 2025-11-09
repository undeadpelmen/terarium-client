package terarium

import (
	"terarium-client/rabbit"
	"terarium-client/rabbit/dto/animal"
)

type TerariumInDto struct {
	rabbit.Message
	TerariumId int
	Animal animal.Animal
}