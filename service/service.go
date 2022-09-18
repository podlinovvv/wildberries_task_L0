package service

import (
	"encoding/json"
	"fmt"
	"github.com/nats-io/stan.go"
	"wb_l0/repos/Cache"
	"wb_l0/structs"
)

var id = "7"

func Sub(h *Handler) {
	handlerF := h.ToStore

	//sc, _ := stan.Connect("test-cluster", id, stan.NatsURL("0.0.0.0:4222"))
	sc, _ := stan.Connect("test-cluster", id, stan.NatsURL("natscont"+":4222"))
	_, err := sc.Subscribe("main", handlerF, stan.DeliverAllAvailable(), stan.DurableName("client-007"))
	if err != nil {
		fmt.Println("can't subscribe to the channel")
		panic(err)
	}
}

/*type Handler struct {
	db    *sql.DB
	cache *map[string][]byte
}*/

/*func newHandler() *Handler {
	c := make(map[string][]byte)

	return &Handler{
		db:    database.ConnectToDb(),
		cache: &c,
	}
}*/

type Handler struct {
	*Cache.Storage
}

func NewHandler() *Handler {
	return &Handler{Cache.NewStorage()}
}

func (h *Handler) ToStore(m *stan.Msg) {
	//fmt.Printf("Received a message: %s\n", string(m.Data))
	newOrder := structs.Order{}

	err := json.Unmarshal(m.Data, &newOrder)
	if err != nil {
		fmt.Println("Wrong JSON, can't unmarshal data")
		return
	}

	//fmt.Println(newOrder.OrderUID)
	_, ok := h.Storage.C[newOrder.OrderUID]
	if ok {
		fmt.Println("already in cache")
		return
	}

	err = h.WriteOrder(newOrder.OrderUID, m.Data)
	if err != nil {
		fmt.Println("Can't validate json\n" + err.Error())
		return
	} else {
		fmt.Println("order stored")
	}
}
