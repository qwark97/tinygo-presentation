package main

import (
	"machine"
	"time"

	"tinygo.org/x/drivers/dht"
)

// START_PART1 OMIT
var (
	tempPin = machine.GPIO9 //S3 of NodeMCU is GPIO9
	light   = false
	led     = machine.GPIO2

	rPin = machine.GPIO14 //D5 of NodeMCU is GPIO14
	gPin = machine.GPIO12 //D6 of NodeMCU is GPIO12
	bPin = machine.GPIO13 //D7 of NodeMCU is GPIO13
)

func main() {
	led.Configure(machine.PinConfig{
		Mode: machine.PinOutput,
	})
	led.Low()
	time.Sleep(time.Second)
	led.High()
	time.Sleep(time.Second)

	rPin.Configure(machine.PinConfig{Mode: machine.PinOutput})
	gPin.Configure(machine.PinConfig{Mode: machine.PinOutput})
	bPin.Configure(machine.PinConfig{Mode: machine.PinOutput})

	ch := make(chan bool)
	go measure(ch)
	for {
		println("loop")
		trigger := <-ch
		if trigger {
			turnOnLedStrip()
			led.Low()
			light = true
		} else {
			turnOffLedStrip()
			led.High()
			light = false
		}
	}
}

func turnOnLedStrip() {
	println("turn on")
	rPin.High()
	gPin.High()
	bPin.High()
}
func turnOffLedStrip() {
	println("turn off")
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

		if !light && temp/10 >= 26 {
			trigger <- true
		} else if light && temp/10 < 26 {
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
