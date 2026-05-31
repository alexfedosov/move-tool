package internal

import (
	"fmt"
	"github.com/alexfedosov/move-tool/internal/ablmodels"
	"github.com/go-audio/audio"
	"github.com/go-audio/wav"
	"io"
	"os"
	"path"
)

func writeAudioFileSlices(filePath string, outputDir string, parts int, filenamePrefix string) (*[]ablmodels.AudioFile, error) {
	source, err := readAudioSource(filePath)
	if err != nil {
		return nil, err
	}
	defer source.Close()

	samplesPerPart := len(source.Buffer.Data) / parts
	sampleDurationMilliseconds := source.SampleDurationMilliseconds(samplesPerPart)
	result := make([]ablmodels.AudioFile, parts)

	for i := 0; i < parts; i++ {
		slice := audioSlice(source.Buffer, i, parts, samplesPerPart)
		audioFile, err := writeAudioFileSlice(slice, outputDir, filenamePrefix, source.BitDepth, sampleDurationMilliseconds)
		if err != nil {
			return nil, err
		}
		result[i] = audioFile
	}

	return &result, nil
}

type audioSource struct {
	File     io.Closer
	Decoder  *wav.Decoder
	Buffer   *audio.IntBuffer
	BitDepth int
}

func readAudioSource(filePath string) (audioSource, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return audioSource{}, fmt.Errorf("could not open source file: %v", err)
	}

	decoder := wav.NewDecoder(file)
	decoder.ReadInfo()
	buf, err := decoder.FullPCMBuffer()
	if err != nil {
		file.Close()
		return audioSource{}, fmt.Errorf("could not read wave data: %v", err)
	}

	return audioSource{
		File:     file,
		Decoder:  decoder,
		Buffer:   buf,
		BitDepth: int(decoder.BitDepth),
	}, nil
}

func (s audioSource) Close() error {
	return s.File.Close()
}

func (s audioSource) SampleDurationMilliseconds(samplesPerPart int) float64 {
	return float64(s.Decoder.SampleRate) * float64(s.BitDepth) / 8 * float64(samplesPerPart) / 1000
}

type audioSliceData struct {
	Buffer *audio.IntBuffer
	Index  int
}

func audioSlice(buf *audio.IntBuffer, index int, parts int, samplesPerPart int) audioSliceData {
	start := index * samplesPerPart
	end := sliceEnd(index, parts, start, samplesPerPart, len(buf.Data))
	return audioSliceData{
		Buffer: &audio.IntBuffer{
			Format: buf.Format,
			Data:   buf.Data[start:end],
		},
		Index: index,
	}
}

func sliceEnd(index int, parts int, start int, samplesPerPart int, totalSamples int) int {
	if index == parts-1 {
		return totalSamples
	}
	return start + samplesPerPart
}

func writeAudioFileSlice(slice audioSliceData, outputDir string, filenamePrefix string, bitDepth int, duration float64) (ablmodels.AudioFile, error) {
	partFileName := fmt.Sprintf("%s_part_%d.wav", filenamePrefix, slice.Index+1)
	outputFilePath := path.Join(outputDir, partFileName)
	partFile, err := os.Create(outputFilePath)
	if err != nil {
		return ablmodels.AudioFile{}, fmt.Errorf("could not create part file: %v", err)
	}
	defer partFile.Close()

	encoder := wav.NewEncoder(partFile, slice.Buffer.Format.SampleRate, bitDepth, slice.Buffer.Format.NumChannels, 1)
	if err := encoder.Write(slice.Buffer); err != nil {
		return ablmodels.AudioFile{}, fmt.Errorf("could not write part buffer: %v", err)
	}
	if err := encoder.Close(); err != nil {
		return ablmodels.AudioFile{}, fmt.Errorf("could not close encoder: %v", err)
	}

	fmt.Printf("Slice %d saved as %s\n", slice.Index+1, outputFilePath)
	sampleFilePath := fmt.Sprintf("Samples/%s", partFileName)
	return ablmodels.AudioFile{
		FilePath: &sampleFilePath,
		Duration: duration,
	}, nil
}
