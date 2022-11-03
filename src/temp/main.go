// START_PART1 OMIT
package main

import (
	"machine"
	"time"

	"tinygo.org/x/drivers/dht"
)

var (
	tempPin = machine.GPIO14 //D5 of NodeMCU is GPIO14
)

func main() {
	device := dht.New(tempPin, dht.DHT11)
	println()
	for {
		_ = device.ReadMeasurements()
		temp, _ := device.Temperature()
		prettyTempPrintln(int(temp))
		time.Sleep(2 * time.Second)
	}
}

// END_PART1 OMIT
// START_PART2 OMIT
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

// END_PART2 OMIT
