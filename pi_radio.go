package main

import (
	"fmt"
	"github.com/MarinX/keylogger"
	"os/exec"
)

const input_device = "/dev/input/by-id/usb-Usb_KeyBoard_Usb_KeyBoard-event-kbd"

var logger *keylogger.KeyLogger
var state string

func main() {
	title("Pi-Radio starting")

	state = "init_display"
	init_display()

	state = "init_logger"
	update_display()

	var error error
	logger, error = keylogger.New(input_device)
	if error != nil {
		info("Error starting keylogger:")
		fmt.Println(error)
	}
	events := logger.Read()

	state = "wait_net"
	update_display()
	waitForNetwork()

	state = "ready"
	update_display()

	// range of events
	for e := range events {
		if e.Type == keylogger.EvKey && e.KeyPress() {
			// EvKey is used to describe state changes of keyboards, buttons, or other key-like devices.
			// check the input_event.go for more events

			key := e.KeyString()
			info("Pressed ", key)

			// turn off NUM_LOCK-LED
			if key == "NUM_LOCK" {
				exec.Command("sh", "-c", "setleds -num < /dev/console").Run()
			}

			handle_macros(key)

		}
	}
}
