package main

import (
	"machine"
	"time"

	"tinygo.org/x/drivers/dht"
)

// START_PART1 OMIT
var (
	led     = machine.GPIO2
	tempPin = machine.GPIO14 //D5 of NodeMCU is GPIO14
	light   = false
)

func main() {
	led.Configure(machine.PinConfig{
		Mode: machine.PinOutput,
	})
	led.High() // turn off the led
	ch := make(chan bool)
	go measure(ch)
	for {
		trigger := <-ch
		if trigger {
			led.Low()
			light = true
		} else {
			led.High()
			light = false
		}
	}
}

// END_PART1 OMIT
// START_PART2 OMIT
func measure(trigger chan bool) {
	device := dht.New(tempPin, dht.DHT11)
	println()
	for {
		_ = device.ReadMeasurements()

		temp, _ := device.Temperature()
		prettyTempPrintln(int(temp))

		if !light && temp/10 >= 30 {
			trigger <- true
		} else if light && temp/10 < 30 {
			trigger <- false
		}
		time.Sleep(2 * time.Second)
	}
}

// END_PART2 OMIT

func prettyTempPrintln(num int) {
	h := num / 100
	d := (num - (h * 100)) / 10
	u := num - (h * 100) - (d * 10)

	print("Temp: ")
	print(h)
	print(d)
	print(",")
	print(u, "\n")
}
