package finishi

import (
	"fmt"
	"os"
	"time"

	"github.com/faiface/beep"
	"github.com/faiface/beep/mp3"
	"github.com/faiface/beep/speaker"
	_ "github.com/freetbash/finishi/src"
)

func Finishi() {
	f, err := os.Open("./tip.mp3")
	if err != nil {
		fmt.Println("finishi module " + err.Error())
	}

	streamer, format, err := mp3.Decode(f)
	defer streamer.Close()

	speaker.Init(format.SampleRate, format.SampleRate.N(time.Second/10))
	done := make(chan bool)
	speaker.Play(beep.Seq(streamer, beep.Callback(func() {
		done <- true
	})))

	<-done
}
