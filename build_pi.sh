export pi_addr="$1"
go mod tidy
ssh $pi_addr mkdir pi_radio
scp \
	pi_radio.go \
	audio_stream.go \
	logging.go \
	playback.go \
	handle_macros.go \
	display.go \
	network.go \
	go.mod \
	go.sum \
	$1:pi_radio/
ssh $pi_addr
