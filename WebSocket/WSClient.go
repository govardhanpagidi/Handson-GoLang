package main2

import (
	"flag"
	"github.com/gorilla/websocket"
	"log"
	"net/url"
	"os"
	"os/signal"
	"time"
)

var address = flag.String("addr", "52.66.206.199:8080", "http server address")

func main() {

	flag.Parse()

	cname := os.Args[1]
	log.SetFlags(0)
	s := "From WSClient-"+ cname + " : "

	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt)

	u := url.URL{Scheme: "ws", Host: *address, Path: "/ds"}
	log.Printf("Connecting to %s ", u.String())

	c, _, err := websocket.DefaultDialer.Dial(u.String(), nil)
	if err != nil {
		log.Fatal("dial", err)
	}
	defer c.Close()

	done := make(chan struct{})

	go func() {
		defer close(done)
		for {
			_, message, err := c.ReadMessage()
			if err != nil {
				log.Printf("Error :", err)
				return
			}
			log.Printf("recv :\n %s", message)

		}
	}()

	timeTicker := time.NewTicker(time.Second * 30)
	defer timeTicker.Stop()

	for {
		select {
		case <-done:
			return
		case t := <-timeTicker.C:
			err := c.WriteMessage(websocket.TextMessage, []byte( s+ t.String()))
			if err != nil {
				log.Printf("Write: ", err)
				return
			}
		case <-interrupt:
			log.Println("interrupt")
			err := c.WriteMessage(websocket.CloseMessage, websocket.FormatCloseMessage(websocket.CloseNormalClosure, "gracefully closed the connection."))
			if err != nil {
				log.Printf("write close :", err)
				return
			}
			select {
			case <-done:
			case <-time.After(time.Second):
			}
			return
		}
	}
}
