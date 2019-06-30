package main

import (
    "os"
    "time"

    "github.com/faiface/beep"
    "github.com/faiface/beep/mp3"
    "github.com/faiface/beep/speaker"
)

func main() {
    soundFile := "music.mp3"
    f, err := os.Open(soundFile)
    if err != nil {
        panic(err)
    }

    streamer, format, err := mp3.Decode(f)
    if err != nil {
        panic(err)
    }
    defer streamer.Close()

    if err := speaker.Init(format.SampleRate, format.SampleRate.N(time.Second/10)); err != nil {
        panic(err)
    }

    done := make(chan bool)
    speaker.Play(beep.Seq(
        streamer,
        beep.Callback(
            func() {
                done <- true
            },
        ),
    ))
    <-done
}
