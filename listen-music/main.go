package main

import (
	"github.com/skosovsky/go-head-first/gadget"
)

type Player interface {
	Play(string)
	Stop()
}

func main() {
	mixtape := []string{"Jessie's Girl", "Whip It", "9 to 5"}
	var player Player = gadget.TapePlayer{}
	playList(player, mixtape)

	player = gadget.TapeRecorder{}
	playList(player, mixtape)

	TryOut(gadget.TapeRecorder{})
}

func playList(device Player, songs []string) {
	for _, song := range songs {
		device.Play(song)
	}
	device.Stop()
}

func TryOut(device Player) {
	device.Play("Test Track")
	device.Stop()

	recorder, ok := device.(gadget.TapeRecorder)
	if ok {
		recorder.Record()
	}
}
