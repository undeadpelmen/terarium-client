package main

import (
	"fmt"
	"log"
	"os"
	"terarium-client/rabbit"
	"terarium-client/sys"
	"time"

	"github.com/joho/godotenv"
)

const (
	Reset  = "\033[0m"
	Red    = "\033[31m"
	Green  = "\033[32m"
	Yellow = "\033[33m"
	Blue   = "\033[34m"
	Purple = "\033[35m"
	Cyan   = "\033[36m"
	Gray   = "\033[90m"
	LightRed = "\033[91m"
	LightGreen = "\033[92m"
	LightYellow = "\033[93m"
	LightBlue = "\033[94m"
	LightPurple = "\033[95m"
	LightGray = "\033[97m"
	LightCyan = "\033[96m"
)

func main() {
	//get env
	godotenv.Load()
	rabbit_user := os.Getenv("RABBIT_USER")
	rabbit_pass := os.Getenv("RABBIT_PASS")
	rabbit_host := os.Getenv("RABBIT_HOST")
	rabbit_port := os.Getenv("RABBIT_PORT")
	
	
	//Set color prefix to console output
	log.SetPrefix(fmt.Sprintf("%s[%s%s%s]%s ",LightBlue , LightPurple, time.Now().Format("02-01-2006 15:04"), LightBlue, LightGreen))
	log.SetFlags(0)
	log.Printf("Starting Terarium Client\n")
	
	//Get MAC
	mac, err := sys.GetMAC()
	if err != nil {
		log.Fatalf(LightRed + "Error: %v" + Reset, err)
	}
	
	log.Printf("Your MAC addres: %s", mac)
	log.Printf("Connecting to RabbitMQ\n")
	
	//Init rabbit mq
	//Create consumer
	cons, err := rabbit.NewConsumer(rabbit_user, rabbit_pass, rabbit_host, rabbit_port, mac)
	if err != nil {
		log.Fatalf(LightRed + "Error: %v" + Reset, err)
	}
	defer cons.Close()
	
	//Create producer
	prod, err := rabbit.NewProducer(rabbit_user, rabbit_pass, rabbit_host, rabbit_port, mac)
	if err != nil {
		log.Fatalf(LightRed + "Error: %v" + Reset, err)
	}
	defer prod.Close()
	
	//Say back-end what we are online
	err = Init(mac, prod)
	if err != nil {
		log.Fatalf(LightRed + "Error: %v" + Reset, err)
	}
	
	log.Printf("Sucseesful conect to Rabbit MQ\n")
	
	//Create queues
	prod.NewQueue("out/" + mac)
	cons.NewQueue("in/" + mac)
	
	//Create chanals
	consumed := make(chan string)
	errchan  := make(chan error)
	
	//Start consume gorutine
	go Consume("in", mac, cons, consumed, errchan)
	
	//Start produce gorutine
	go Produce("out", mac, prod, consumed, errchan)
	
	log.Printf("Waiting for messages\n")
	
	//Error handler
	for {
		err := <- errchan
		log.Fatalf(LightRed + "Error: %v" + Reset, err)
	}
}