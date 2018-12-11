// package main

// import (
// 	"github.com/gorilla/websocket"
// 	"log"
// 	"net/http"
// 	"html/template"
// 	"os"
// )

// //Configuration
// var address = "localhost:8080"

// func init() {
// 	nf, err := os.Create("log.txt")
// 	if err != nil {
// 		log.Println(err)
// 	}
// 	log.SetOutput(nf)
// }

// func main() {
// 	//Attach websocket handler
// 	http.HandleFunc("/ds", handleWebsocket)

// 	// Attach home page of device server 
// 	http.HandleFunc("/", land)
// 	log.Fatal(http.ListenAndServe(address, nil))
// }

// var connectionUpgrader = websocket.Upgrader{}

// func handleWebsocket(w http.ResponseWriter, r *http.Request) {

// 	conn, err := connectionUpgrader.Upgrade(w, r, nil)

// 	if err != nil {
// 		log.Println("Upgrade : ", err)
// 	}
// 	defer conn.Close()
// 	log.Println("Received WS request")

// }

// func land(w http.ResponseWriter, r *http.Request) {
// 	homeTemplate.Execute(w, "ws://"+r.Host+"/ds")
// }

// var homeTemplate = template.Must(
// 					template.New("ServerIndex").Parse(`
// 									<!DOCTYPE html>
// 									<html>
// 										<head>
// 											<meta charset="utf-8">
// 										</head>
// 										<body>
// 											<div>
// 												<p><b> Device server </b> is up and running.</p>
// 											</div>
// 										</body>
// 									</html>
// `))