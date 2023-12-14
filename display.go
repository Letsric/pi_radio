package main

import (
	"github.com/Letsric/tinypi_i2c"
	"tinygo.org/x/drivers"
	"tinygo.org/x/drivers/hd44780i2c"
)

var bus drivers.I2C
var screen hd44780i2c.Device

func init_display() {
	bus = tinypii2c.New(1)
	screen = hd44780i2c.New(bus, 0x27)

	config := &hd44780i2c.Config{
		Width:  20,
		Height: 4,
	}

	screen.Configure(*config)
	update_display()
}

func update_display() {
	screen.ClearDisplay()

	screen.Print([]byte("Pi-Radio\n========"))
	screen.SetCursor(0, 2)

	switch state {
	case "init_display":
		screen.Print([]byte("initializing Display"))
	case "init_logger":
		screen.Print([]byte("initializing Keypad"))
	case "wait_net":
		screen.Print([]byte("waiting for network connection"))
	case "ready":
		screen.SetCursor(0, 3)
		if currently_playing == "" {
			screen.Print([]byte("ready"))
		} else {
			screen.Print([]byte(currently_playing))
		}
	}
}

func display_backlight(status bool) {
	screen.BacklightOn(status)
}
