package rutine

import (
	"math/rand"
	"time"
)

func GpioOut(outch chan OutTerMsg, errch chan error) {
	
	
	for {
		temperature := 15 + rand.Int()%5
		humidity := float32(65+rand.Int()%5) + rand.Float32()

		msg := OutTerMsg{
			Temperarure: temperature,
			Humidity:    humidity,
		}

		outch <- msg

		time.Sleep(30 * time.Second)
	}
}
