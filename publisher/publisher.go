package main

import (
	"encoding/json"
	"fmt"
	"github.com/nats-io/stan.go"
	"math/rand"
	"strconv"
	"time"
	"wb_l0/structs"
)

func main() {
	connection := connect()
	pub(connection)
}

func connect() stan.Conn {
	sc, err := stan.Connect("test-cluster", "1", stan.NatsURL("0.0.0.0:4222"))
	if err != nil {
		fmt.Println(err.Error())
		return nil
	}
	return sc
}

func pub(sc stan.Conn) {
	var uid int
	for {
		uid++
		data := generateRandomJson(uid)
		if err := sc.Publish("main", data); err != nil {
			fmt.Println(err.Error())
			return
		}
		fmt.Printf("sent good json with id %d\n", uid)
		if uid%4 == 0 {
			if err := sc.Publish("main", generateInvalidJson()); err != nil {
				fmt.Println(err.Error())
				return
			}
			fmt.Println("sent bad json")
		}

		time.Sleep(2 * time.Second)
	}
}

func generateRandomJson(uid int) []byte {
	var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ 0123456789")
	newOrder := structs.Order{}

	newOrder.OrderUID = strconv.Itoa(uid)

	//strings
	newOrder.TrackNumber = randSeqString(15, letters)
	newOrder.Entry = randSeqString(5, letters)
	newOrder.Locale = randSeqString(2, letters)
	newOrder.InternalSignature = randSeqString(5, letters)
	newOrder.CustomerID = randSeqString(5, letters)
	newOrder.DeliveryService = randSeqString(5, letters)
	newOrder.Shardkey = randSeqString(5, letters)
	newOrder.OofShard = randSeqString(5, letters)

	newOrder.Delivery.Name = randSeqString(10, letters)
	newOrder.Delivery.Phone = randSeqString(10, letters)
	newOrder.Delivery.Zip = randSeqString(7, letters)
	newOrder.Delivery.City = randSeqString(5, letters)
	newOrder.Delivery.Address = randSeqString(15, letters)
	newOrder.Delivery.Region = randSeqString(5, letters)
	newOrder.Delivery.Email = randSeqString(5, letters) + "@gmail.com"

	newOrder.Payment.Transaction = randSeqString(19, letters)
	newOrder.Payment.RequestID = randSeqString(10, letters)
	newOrder.Payment.Currency = randSeqString(3, letters)
	newOrder.Payment.Provider = randSeqString(5, letters)
	newOrder.Payment.Bank = randSeqString(5, letters)

	newItem := structs.Item{}
	newItem.TrackNumber = randSeqString(5, letters)
	newItem.Rid = randSeqString(5, letters)
	newItem.Name = randSeqString(5, letters)
	newItem.Size = randSeqString(5, letters)
	newItem.Brand = randSeqString(5, letters)

	//ints
	newOrder.SmID = rand.Int63n(10)

	newOrder.Payment.Amount = rand.Int63n(5)
	newOrder.Payment.PaymentDt = rand.Int63n(5)
	newOrder.Payment.DeliveryCost = rand.Int63n(5)
	newOrder.Payment.GoodsTotal = rand.Int63n(5)
	newOrder.Payment.CustomFee = rand.Int63n(5)

	newItem.ChrtID = rand.Int63n(5)
	newItem.Price = rand.Int63n(5)
	newItem.Sale = rand.Int63n(5)
	newItem.TotalPrice = rand.Int63n(5)
	newItem.NmID = rand.Int63n(5)
	newItem.Status = rand.Int63n(5)

	newOrder.Items = []structs.Item{newItem}

	b, err := json.Marshal(newOrder)
	if err != nil {
		fmt.Println("error:", err)
	}
	return b
}

func randSeqString(n int, letters []rune) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}

func generateInvalidJson() []byte {
	notValidJson := struct {
		SmID string `json:"sm_id"`
	}{}
	notValidJson.SmID = "wrond id"
	b, err := json.Marshal(notValidJson)
	if err != nil {
		fmt.Println("error:", err)
	}
	return b
}
