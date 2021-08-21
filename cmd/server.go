package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/tada3/go-websocket/ws"
)

func main() {
	fmt.Println("Server start!")

	err := testServer(8888)
	if err != nil {
		fmt.Println("Error!", err)
	}

	fmt.Println("done")
}

func testServer(p int) error {
	http.HandleFunc("/hoge", handler)
	http.HandleFunc("/ws", wsHandler)
	addr := fmt.Sprintf("0.0.0.0:%d", p)
	http.ListenAndServe(addr, nil)

	return nil
}

func handler(w http.ResponseWriter, req *http.Request) {
	w.Write([]byte("hello world"))
}

func wsHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("wsHandler 0000")
	ws, err := ws.NewWs(w, r)
	if err != nil {
		log.Println(err)
		return
	}

	log.Println("wsHandler 100")
	err = ws.Handshake()
	if err != nil {
		log.Println(err)
		return
	}

	defer ws.Close()

	for {
		log.Println("Waiting client msg..")
		frame, err := ws.Recv()
		if err != nil {
			log.Println("Error Decoding", err)
			return
		}

		log.Printf("Received frame: %+v\n", frame)

		switch frame.Opcode {
		case 8: // Close
			log.Println("Close it!")
			return
		case 9: // Ping
			frame.Opcode = 10
			fallthrough
		case 0: // Continuation
			fallthrough
		case 1: // Text
			fallthrough
		case 2: // Binary
			if err = ws.Send(frame); err != nil {
				log.Println("Error sending", err)
				return
			}

			err = sendN(ws, frame, 10)
			if err != nil {
				log.Println("Error sending", err)
				return
			}

		}
	}

}

func sendN(ws *ws.Ws, f ws.Frame, n int) error {
	for i := 0; i < n; i++ {
		log.Println("Sending msg..", i)
		err := ws.Send(f)
		if err != nil {
			return err
		}
		log.Println("Sleep 1 second")
		time.Sleep(1 * time.Second)
	}
	return nil

}
