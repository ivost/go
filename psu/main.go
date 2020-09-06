package main

import (
	"fmt"
	mqtt "github.com/eclipse/paho.mqtt.golang"
	"github.com/shirou/gopsutil/mem"
	"log"
	"time"
)

const TOPIC = "top"
const MB = 1024 * 1024

func main() {

	opts := mqtt.NewClientOptions().AddBroker("tcp://localhost:1883")
	client := mqtt.NewClient(opts)
	if token := client.Connect(); token.Wait() && token.Error() != nil {
		log.Fatal(token.Error())
	}

	for {
		v, _ := mem.VirtualMemory()
		msg := fmt.Sprintf("Memory - Total: %v MB, Available: %v MB, Used: %d%%\n",
			int(v.Total/MB), int(v.Available/MB), int(v.UsedPercent))
		// fmt.Println(v)
		log.Print(msg)
		_ = client.Publish(TOPIC, 0, false, msg)
		//if token := client.Publish(TOPIC, 0, false, msg); token.Wait() && token.Error() != nil {
		//	log.Fatal(token.Error())
		//}

		time.Sleep(5*time.Second)
	}


}


