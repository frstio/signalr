package main

import (
	"log"

	"github.com/frstio/signalr"
)

func main() {
	// Prepare a SignalR client.
	c := signalr.New(
		"www.cryptopia.co.nz",
		"1.5",
		"/signalr",
		`[{"name":"notificationhub"}]`,
		nil,
	)

	// Define message and error handlers.
	msgHandler := func(msg signalr.Message) { log.Println(msg) }
	panicIfErr := func(err error) {
		if err != nil {
			log.Panic(err)
		}
	}

	// Start the connection.
	err := c.Run(msgHandler, panicIfErr)
	panicIfErr(err)

	// Wait indefinitely.
	select {}
}
