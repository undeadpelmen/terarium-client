package main

import (
	"terarium-client/rabbit"
	"time"
)

func Init(mac string, prod *rabbit.Producer) error {
	err := prod.NewQueue("init")
	if err != nil {
		return err
	}
	
	mes := rabbit.Mesage{
		Mac: mac,
		Message: "Terarium online",
		Time: time.Now(),
	}
	
	err = prod.Publish("", "init", mes.JSON())
	if err != nil {
		return err
	}
	
	return nil
}