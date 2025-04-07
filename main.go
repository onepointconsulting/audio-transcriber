package main

import (
	"brahmakumaris/audiotranscriber/cmd"
	"log"
)

func main() {
	if err := cmd.Execute(); err != nil {
		log.Fatal(err)
	}
	transcribeAudio(cmd.AudioDir, cmd.Recursive)
}
