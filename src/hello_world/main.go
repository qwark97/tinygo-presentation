package main

import (
	"machine" // <--- TinyGo specific
	"time"
)

var led = machine.GPIO2

func main() {
	led.Configure(machine.PinConfig{
		Mode: machine.PinOutput,
	})

	for {
		led.Low()
		time.Sleep(time.Second)
		led.High()
		time.Sleep(time.Second)
	}
}
