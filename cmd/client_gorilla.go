package main

import (
	"fmt"
	"log"
	"net/url"

	"github.com/gorilla/websocket"
)

func main() {
	fmt.Println("Client start!")

	err := testGClient(8888)
	if err != nil {
		fmt.Println("Error!", err)
	}

	fmt.Println("done")
}

func testGClient(p int) error {

	addr := fmt.Sprintf("localhost:%d", p)
	u := url.URL{Scheme: "ws", Host: addr, Path: "/ws"}
	log.Printf("connecting to %s", u.String())

	c, _, err := websocket.DefaultDialer.Dial(u.String(), nil)
	if err != nil {
		return err
	}
	defer c.Close()

	log.Println("Writing msg..")
	err = c.WriteMessage(websocket.TextMessage, []byte("Unko!"))
	if err != nil {
		return err
	}

	err = readN(c, 11)
	if err != nil {
		return err
	}

	err = c.WriteMessage(websocket.CloseMessage, websocket.FormatCloseMessage(websocket.CloseNormalClosure, ""))
	if err != nil {
		return err
	}

	return nil
}

func readN(c *websocket.Conn, n int) error {
	for i := 0; i < n; i++ {
		log.Println("Reading msg..", i)
		mt, message, err := c.ReadMessage()
		if err != nil {
			return err
		}
		log.Printf("recv: %+v, %s", mt, message)
	}
	return nil
}
