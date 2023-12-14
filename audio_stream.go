// slightly modified example from minimp3 library

package main

import (
	"fmt"
	"io"
	"net/http"
	"sync"
	"time"

	"github.com/hajimehoshi/oto"
	"github.com/tosone/minimp3"
)

func audio_strean(mp3url string) {
	var err error

	var response *http.Response
	if response, err = http.Get(mp3url); err != nil {
		fmt.Println(err)
		return
	}

	var dec *minimp3.Decoder
	if dec, err = minimp3.NewDecoder(response.Body); err != nil {
		fmt.Println(err)
	}
	<-dec.Started()

	info("sample rate:", dec.SampleRate, "channels:", dec.Channels)

	var context *oto.Context
	if context, err = oto.NewContext(dec.SampleRate, dec.Channels, 2, 4096); err != nil {
		fmt.Println(err)
	}

	var waitForPlayOver = new(sync.WaitGroup)
	waitForPlayOver.Add(1)

	var player = context.NewPlayer()

	go func() {
		defer response.Body.Close()
	loop:
		for {
			select {
			case <-stop_playing:
				dec.Close()
				player.Close()
				context.Close()
				break loop
			default:
				var data = make([]byte, 512)
				_, err = dec.Read(data)
				if err == io.EOF {
					break loop
				}
				if err != nil {
					fmt.Println(err)
					break loop
				}
				player.Write(data)
			}
		}
		waitForPlayOver.Done()
	}()

	time.Sleep(1)
}
