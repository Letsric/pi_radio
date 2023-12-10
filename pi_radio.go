package main

import (
	"fmt"
	"github.com/MarinX/keylogger"
	"os/exec"
)

var logger *keylogger.KeyLogger

func main() {
	input_device := "/dev/input/by-id/usb-Usb_KeyBoard_Usb_KeyBoard-event-kbd"

	title("Pi-Radio starting")
	var error error
	logger, error = keylogger.New(input_device)
	if error != nil {
		info("Error starting keylogger:")
		fmt.Println(error)
	}

	// "global" vars
	events := logger.Read()

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
