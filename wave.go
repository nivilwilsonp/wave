package wave

import (
	"os"
	"strings"
	"time"
)

type WaveFileFormat struct {

	//RIFF chunk
	ChunkID   string
	ChunkSize string
	Format    string

	//fmt sub chunk
	Subchunk1ID   string
	Subchunk1Size int
	AudioFormat   int
	NumChannels   int
	SampleRate    int
	ByteRate      int
	BlockAlign    int
	BitsPerSample int

	//data sub chunk
	Subchunk2ID   string
	Subchunk2Size string
}

func (format *WaveFileFormat) SetDefaultFormat() {
	format.ChunkID = "RIFF" //Marks the file as a riff file. Characters are each 1 byte long

	format.ChunkSize = "----" /*This is the size of the
	entire file in bytes minus 8 bytes for the
	two fields not included in this count:
	ChunkID(4bytes) and ChunkSize(4 bytes). Typically, you’d fill this in after creation.*/

	format.Format = "WAVE"      //File Type Header. For our purposes, it always equals “WAVE”.
	format.Subchunk1ID = "fmt " //Format chunk marker(4bytes). Includes trailing null(a blank space)
	format.Subchunk1Size = 16   /*16 for PCM.  This is the size of the
	rest of the Subchunk which follows this number.*/

	format.AudioFormat = 1 /*PCM = 1 (i.e. Linear quantization)
	Values other than 1 indicate some
	form of compression. */
	format.NumChannels = 1                                                                // Mono = 1, Stereo = 2, etc.
	format.SampleRate = 44100                                                             // 8000, 44100,48000 etc.
	format.ByteRate = format.SampleRate * format.NumChannels * (format.Subchunk1Size / 8) //== SampleRate * NumChannels * BitsPerSample/8
	format.BlockAlign = format.NumChannels * (format.Subchunk1Size / 8)                   //== NumChannels * BitsPerSample/8
	format.BitsPerSample = 16                                                             // 8 bits = 8, 16 bits = 16, etc.

	//data sub chunk

	format.Subchunk2ID = "data"   //data” chunk header. Marks the beginning of the data section.
	format.Subchunk2Size = "----" //Size of the data section.
}

func (wavefile *WaveFile) SetFileNameAutomatically() {

	now := time.Now()

	fileName := now.Format(time.RFC3339)
	if !strings.HasSuffix(fileName, ".wav") {
		fileName += ".wav"
	}
	wavefile.FileName = fileName
}

type WaveFile struct {
	Format   WaveFileFormat
	FileName string
}

func (wavefile *WaveFile) createFile() {
	var err error
	_file, err = os.Create(wavefile.FileName)
	if err != nil {
		panic(err)
	}
}

func (wavefile *WaveFile) writeHeader() {

}

// global object of WaveFile
var _wavfile *WaveFile
var _file *os.File

func GetVersion() string {
	return "v1.0.0"
}

func GetWaveFile() *WaveFile {
	_wavfile = new(WaveFile)
	_wavfile.Format.SetDefaultFormat()
	_wavfile.SetFileNameAutomatically()
	_wavfile.createFile()

	return _wavfile
}
