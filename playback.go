package main

var currently_playing = ""
var stop_playing = make(chan bool)

func stop_playback() {
	if currently_playing != "" {
		stop_playing <- true
		currently_playing = ""
	}
	update_display()
}
