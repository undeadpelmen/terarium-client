package terarium

import (
	"terarium-client/rabbit"
)

type TerariumDto struct {
	rabbit.Message
	TerariumId int
	Name string
}