package main

/*
// Load Wi-Fi library
#include <ESP8266WiFi.h>

const char* ssid     = "<ssid>";
const char* password = "<pass>";

void wifiConnect() {
	// Connect to Wi-Fi network with SSID and password
	WiFi.begin(ssid, password);
	while (WiFi.status() != WL_CONNECTED) {
		delay(500);
	}
	return;
}
*/
import "C"

import (
	"machine"
	"time"
)

func main() {
	C.wifiConnect()
	blinking()
}

func blinking() {
	led := machine.GPIO2
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
