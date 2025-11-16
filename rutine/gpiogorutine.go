package rutine

import (
	"math/rand"
	"terarium-client/rabbit/dto/terarium"
	"time"
)

func GpioOut(outch chan terarium.TerariumOutDto, errch chan error) {
	
	
	for {
		time.Sleep(30 * time.Second)
		
		temperature := 15 + rand.Int()%5
		humidity := float32(65+rand.Int()%5) + rand.Float32()
		
		msg := terarium.TerariumOutDto{
			Temperature: temperature,
			Humidity:    humidity,
		}
		
		outch <- msg
	}
}

func GpioIn(inch chan string, logch chan string){
	for {
		time.Sleep(30 * time.Second)
		
		msg := <- inch
		
		logch <- msg
	}
}
