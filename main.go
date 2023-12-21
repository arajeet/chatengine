package main

import (
	"bufio"
	"fmt"
	"net/http"
	"os"

	"github.com/gorilla/websocket"
)

type Router struct{}

var upgrader = websocket.Upgrader{}

func (ro Router) ServeHTTP(w http.ResponseWriter, r *http.Request) error {
	switch r.URL.Path {
	case "/":
		handlerx(w, r)
	case "/connect":

	default:
		handlerx(w, r)

	}
	return nil
}

func handlerx(w http.ResponseWriter, r *http.Request) {

	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		fmt.Errorf("Could Not Upgrade the http connection to  Websocket", err)
	}
	mt, msg, _ := conn.ReadMessage()
	fmt.Println(msg)
	dat, err := os.Open("/var/tmp/test.txt")
	for {

		//dat, err := os.Open("/var/tmp/test.txt")
		if err != nil {
			fmt.Errorf("Could not Open file ", err)
		}
		defer dat.Close()
		scanner := bufio.NewScanner(dat)
		for scanner.Scan() {
			conn.WriteMessage(mt, scanner.Bytes())
		}
	}
}

func getChat(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "templates/websocket.html")
}

func main() {
	http.HandleFunc("/todo", handlerx)
	http.HandleFunc("/", getChat)
	http.ListenAndServe(":8080", nil)
	fmt.Println("Hello Chat Engine")
}
