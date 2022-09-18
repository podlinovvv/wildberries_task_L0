package main

import (
	"fmt"
	"time"
	"wb_l0/service"
	"wb_l0/webservice"
)

func main() {
	//defer database.Close()

	//перезапуск при панике, например, если ещё не запустилась бд
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("Recovered", err)
		}
	}()

	//ждём пока в докере запустятся nats-streaming и db
	time.Sleep(10 * time.Second)
	nh := service.NewHandler()
	service.Sub(nh)
	webservice.RunServer(nh)
}
