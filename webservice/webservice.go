package webservice

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"wb_l0/service"
)

var tmplt *template.Template

type Page struct {
	OID  string
	Body string
}

func RunServer(nh *service.Handler) {
	mux := http.NewServeMux()
	NewWebHandler := CreateWebHandler(nh)
	mux.HandleFunc("/", NewWebHandler.handlePage)
	mux.HandleFunc("/get", NewWebHandler.handlePage)

	//err := http.ListenAndServe("localhost:8080", mux)
	err := http.ListenAndServe(":8080", mux)
	if err != nil {
		log.Fatalln("There's an error with the server:", err)
	}
}

type WebHandler struct {
	nh *service.Handler
}

func CreateWebHandler(nh *service.Handler) *WebHandler {
	return &WebHandler{
		nh: nh,
	}
}

func (wh *WebHandler) handlePage(writer http.ResponseWriter, request *http.Request) {
	if request.Method == "GET" {
		tmplt, err := template.ParseFiles("html/index.html")
		if err != nil {
			fmt.Println(err.Error())
		}
		event := Page{}
		err = tmplt.Execute(writer, event)
		if err != nil {
			fmt.Println(err.Error())
			return
		}
	}

	if request.Method == "POST" {
		tmplt, err := template.ParseFiles("html/index.html")
		if err != nil {
			fmt.Println(err.Error())
		}

		orderId := request.FormValue("id")

		jsonToPrint, err := wh.nh.ReadOrderById(orderId)

		/*newOrder := structs.Order{}
		err = json.Unmarshal(jsonToPrint, &newOrder)
		niceJson, err := json.MarshalIndent(newOrder, "", "  ")
		if err != nil {
			fmt.Println(err)
		}

		event := Page{
			Body: string(niceJson),
		}*/

		event := Page{
			OID:  "ID: " + orderId,
			Body: string(jsonToPrint),
		}

		if err != nil {
			event = Page{
				Body: "id \"" + orderId + "\" не найден",
			}
		}

		err = tmplt.Execute(writer, event)
		if err != nil {
			fmt.Println(err.Error())
			return
		}
	}
}
