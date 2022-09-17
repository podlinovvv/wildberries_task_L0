package main

import (
	"encoding/json"
	"fmt"
	"time"
	"wb_l0/publisher"
	"wb_l0/service"
	"wb_l0/structs"
	"wb_l0/webservice"
)

func main() {
	//defer database.Close()

	defer func() {
		if err := recover(); err != nil {
			fmt.Println("Recovered", err)
		}
	}()

	//pub
	sc := publisher.Connect()
	publisher.Pub(sc)
	time.Sleep(1 * time.Second)

	time.Sleep(1 * time.Second)
	//sub
	nh := service.NewHandler()
	service.Sub(nh)

	webservice.RunServer(nh)

	time.Sleep(1 * time.Second)

	by, _ := nh.Storage.ReadOrderById("1")
	newOrder2 := structs.Order{}
	err := json.Unmarshal(by, &newOrder2)
	if err != nil {
	}
	fmt.Println(newOrder2)

	time.Sleep(1 * time.Second)

}
