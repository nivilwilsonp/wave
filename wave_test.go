package wave

import (
	"math"
	"testing"
)

func TestWave(t *testing.T) {

	wav := GetWaveFile()
	generateTestData(wav)
	// if  {
	// 	t.Error("generateTestData-Failed expected >0, and got "))
	// }

	wav.CloseFile()
}

func generateTestData(wav *WaveFile) {

	var maxAmplitude float64 = 32760
	var frequency float64 = 250
	var duration float64 = 2

	var i float64
	for i = 0; i < float64(wav.Format.SampleRate)*duration; i++ {
		amplitude := float64(i / float64(wav.Format.SampleRate) * maxAmplitude)
		value := math.Sin(2 * 3.14 * i * frequency / float64(wav.Format.SampleRate))
		channel1 := amplitude * value
		channel2 := (maxAmplitude - amplitude) * value

		wav.WriteData(i16tob(uint16(channel1)))
		wav.WriteData(i16tob(uint16(channel2)))
	}

}
