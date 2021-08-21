package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"nhooyr.io/websocket"
	"nhooyr.io/websocket/wsjson"
)

// Calling readN() fails because nhooyr.io/websocket does not read frame by frame.
// It tries to read all available frames. 
// If you use nhooyr.io/websocket, you need to return response for each server message.

func main() {
	fmt.Println("Client start!")

	err := testClient(8888)
	if err != nil {
		fmt.Println("Error!", err)
	}

	fmt.Println("done")
}

func testClient(p int) error {

	ctx, cancel := context.WithTimeout(context.Background(), time.Minute)
	defer cancel()

	addr := fmt.Sprintf("ws://localhost:%d/ws", p)
	log.Println("Dial to ", addr)
	c, _, err := websocket.Dial(ctx, addr, nil)
	if err != nil {
		return err
	}
	defer c.Close(websocket.StatusInternalError, "the sky is falling")

	log.Println("Write hi")
	err = wsjson.Write(ctx, c, "hi")
	if err != nil {
		return err
	}

	//var v interface{}
	log.Println("Reading")
	//err = wsjson.Read(ctx, c, &v)
	mt, b, err := c.Read(ctx)
	if err != nil {
		return err
	}
	//fmt.Printf("read: %+v\n", v)
	log.Printf("Read: %+v, %s\n", mt, string(b))

	readN(ctx, c, 10)

	c.Close(websocket.StatusNormalClosure, "")

	return nil
}

func readN(ctx context.Context, c *websocket.Conn, n int) error {
	for i := 0; i < n; i++ {
		log.Println("readN", i)
		//var v interface{}
		//err := wsjson.Read(ctx, c, &v)
		mt, b, err := c.Read(ctx)
		if err != nil {
			return err
		}
		//fmt.Printf("read: %+v\n", v)
		log.Printf("Read: %+v, %s\n", mt, string(b))
	}
	return nil
}
