package main

import (
	"machine"
	"time"

	"tinygo.org/x/drivers/dht"
)

var (
	tempPin = machine.GPIO9 //S3 of NodeMCU is GPIO9
	ledPin  = machine.GPIO2
	light   = false

	rPin = machine.GPIO14 //D5 of NodeMCU is GPIO14
	gPin = machine.GPIO12 //D6 of NodeMCU is GPIO12
	bPin = machine.GPIO13 //D7 of NodeMCU is GPIO13
)

func main() {
	ledPin.Configure(machine.PinConfig{
		Mode: machine.PinOutput,
	})
	ledPin.High()

	rPin.Configure(machine.PinConfig{Mode: machine.PinOutput})
	gPin.Configure(machine.PinConfig{Mode: machine.PinOutput})
	bPin.Configure(machine.PinConfig{Mode: machine.PinOutput})

	ch := make(chan bool)
	go measure(ch)
	for {
		trigger := <-ch
		if trigger {
			ledStripOn()
			ledPin.Low()
			light = true
		} else {
			ledStripOff()
			ledPin.High()
			light = false
		}
	}
}

func ledStripOn() {
	rPin.High()
	gPin.High()
	bPin.High()
}
func ledStripOff() {
	rPin.Low()
	gPin.Low()
	bPin.Low()
}

func measure(trigger chan bool) {
	device := dht.New(tempPin, dht.DHT11)
	for {
		_ = device.ReadMeasurements()

		temp, _ := device.Temperature()
		prettyTempPrintln(int(temp))

		if !light && temp/10 > 27 {
			trigger <- true
		} else if light && temp/10 <= 27 {
			trigger <- false
		}
		time.Sleep(2 * time.Second)
	}
}

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
