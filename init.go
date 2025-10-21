package main

import (
	"terarium-client/rabbit"
	"time"
)

func Init(mac string, prod *rabbit.Producer) error {
	err := prod.NewQueue("terarium.init")
	if err != nil {
		return err
	}
	
	mes := rabbit.Message{
		Mac: mac,
		Message: "Terarium online",
		Time: time.Now().Format("02-01-2006 15:04"),
	}
	
	err = prod.Publish("", "terarium.init", mes.JsonToString())
	if err != nil {
		return err
	}
	
	return nil
}