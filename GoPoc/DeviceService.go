package main

import (
	"github.com/gorilla/websocket"
	"log"
	"net/http"
	"html/template"
	"os"
	"github.com/nu7hatch/gouuid"
	"fmt"

)

//Configuration
var address = "localhost:8080"

func init() {
	nf, err := os.Create("log.txt")
	if err != nil {
		log.Println(err)
	}
	log.SetOutput(nf)
}

func main() {
	//Attach websocket handler
	http.HandleFunc("/ds", handleWebsocket)

	// Attach home page of device server 
	http.HandleFunc("/", land)
	log.Fatal(http.ListenAndServe(address, nil))
}

var connectionUpgrader = websocket.Upgrader{}

func handleWebsocket(w http.ResponseWriter, r *http.Request) {

	fmt.Println("Received WS request")

	conn, err := connectionUpgrader.Upgrade(w, r, nil)

	if err != nil {
		log.Println("Upgrade : ", err)
	}
	defer conn.Close()
	log.Println("Received WS request")

	for {
		messageType, message, error := conn.ReadMessage()
		if error != nil {
			log.Println("read : ", error)
		}


		val := string (message)
		val = "Server Acknowledgment for ---> " + val
		err := conn.WriteMessage(messageType, []byte (val))

		if error != nil {
			log.Println("Write : ", err)
		}

		//eventId := string (uuid.NewRandom())

		uuid, err := uuid.NewV4()
		if err != nil {
			log.Println("UUID generation failed : ", err)
		}
		log.Println("Event id, Val ",uuid.String(),val)

		go SaveEvent(uuid.String(),val)
		

		fmt.Println("Event saved")
	}

}

func land(w http.ResponseWriter, r *http.Request) {
	homeTemplate.Execute(w, "ws://"+r.Host+"/ds")
}

var homeTemplate = template.Must(
					template.New("ServerIndex").Parse(`
									<!DOCTYPE html>
									<html>
										<head>
											<meta charset="utf-8">
										</head>
										<body>
											<div>
												<p><b> Device server </b> is up and running.</p>
											</div>
										</body>
									</html>
`))