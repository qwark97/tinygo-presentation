package main

import (
	"machine"
	"time"

	"tinygo.org/x/drivers/espat"
)

func main() {
	led := machine.GPIO2
	led.Configure(machine.PinConfig{Mode: machine.PinOutput})
	uart := machine.DefaultUART
	uart.Configure(machine.UARTConfig{
		TX: machine.UART_RX_PIN,
		RX: machine.UART_TX_PIN,
	})
	adaptor := espat.New(uart)
	adaptor.Configure()
	if isESPConnected(adaptor) {
		connect(adaptor)
	}
	ping(led)
	select {}
}

// START OMIT
func isESPConnected(adaptor *espat.Device) bool {
	for {
		if adaptor.Connected() {
			return true
		}
		println("failed...")
		time.Sleep(1 * time.Second)
	}
}

func connect(adaptor *espat.Device) {
	adaptor.ConnectToAccessPoint("<ssid>", "<pass>", 10)

	time.Sleep(5 * time.Second)
	println(adaptor.GetClientIP())
}

func ping(led machine.Pin) {
	led.Low()
	time.Sleep(500 * time.Millisecond)
	led.High()
	time.Sleep(500 * time.Millisecond)
}

// END OMIT
