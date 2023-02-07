# wave
go package to create wav files .

#To Include this module run go get -u github.com/nivilwilsonp/wave

#steps to do after including process
create a file named main.go and follow the steps

package main

import (
	"github.com/nivilwilsonp/wave"
)

func main() {
	// step1:initializing wave
	w := wave.GetWaveFile()

	// step2: collect audio data([]byte) from audio input or wave decoder and call the function WriteData by passing the collected byte data
	b := make([]uint8, 0)
	w.WriteData(b)

	// call CloseFile() to close the opened file for writting
	w.CloseFile()

}
