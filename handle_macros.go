package main

import (
	"os/exec"
)

func handle_macros(key string) {
	switch key {
	case "+":
		title("Quitting")
		info("calling stop playback")
		stop_playback()
		info("calling logger close")
		logger.Close()
		info("done")
	case "-":
		title("Rebooting")
		stop_playback()
		logger.Close()
		exec.Command("reboot").Run()
	case "BS":
		title("Shutting down")
		stop_playback()
		logger.Close()
		exec.Command("poweroff").Run()
	case "R_ENTER":
		info("Stopping playpack")
		stop_playback()
		update_display()
	case "INS":
		info("Display backlight off")
		display_backlight(false)
	case "DEL":
		info("Display backlight on")
		display_backlight(true)
	case "END_1":
		info("playing SWR1 BW")
		stop_playback()
		currently_playing = "SWR1 BW"
		update_display()
		audio_strean("https://liveradio.swr.de/sw282p3/swr1bw/")
	case "DOWN":
		info("playing Radio Regenbogen")
		stop_playback()
		currently_playing = "Radio Regenbogen"
		update_display()
		audio_strean("https://audiotainment-sw.streamabc.net/atsw-regenbogen1011-mp3-128-1810210")
	case "PGDN_3":
		info("playing SWR3")
		stop_playback()
		currently_playing = "SWR3"
		update_display()
		audio_strean("https://liveradio.swr.de/sw282p3/swr3/")
	}

}
