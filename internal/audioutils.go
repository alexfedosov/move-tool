package internal

import (
	"fmt"
	"github.com/alexfedosov/move-tool/internal/ablmodels"
	"github.com/go-audio/audio"
	"github.com/go-audio/wav"
	"os"
	"path"
)

func writeAudioFileSlices(filePath string, outputDir string, parts int, filenamePrefix string) (*[]ablmodels.AudioFile, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, fmt.Errorf("could not open source file: %v", err)
	}

	decoder := wav.NewDecoder(file)
	decoder.ReadInfo()

	// Read the entire wave data
	buf, err := decoder.FullPCMBuffer()
	if err != nil {
		return nil, fmt.Errorf("could not read wave data: %v", err)
	}

	// Calculate the number of samples per part
	samplesPerPart := len(buf.Data) / parts

	sampleDurationMilliseconds := float64(decoder.SampleRate) * float64(decoder.BitDepth) / 8 * float64(samplesPerPart) / 1000

	result := make([]ablmodels.AudioFile, parts)

	// Loop through each part and save it as a new file
	for i := 0; i < parts; i++ {
		start := i * samplesPerPart
		end := start + samplesPerPart
		if i == parts-1 {
			end = len(buf.Data) // Make sure the last part gets all remaining samples
		}

		partBuffer := &audio.IntBuffer{
			Format: buf.Format,
			Data:   buf.Data[start:end],
		}

		// Generate output file path
		partFileName := fmt.Sprintf("%s_part_%d.wav", filenamePrefix, i+1)
		outputFilePath := path.Join(outputDir, partFileName)
		partFile, err := os.Create(outputFilePath)
		if err != nil {
			return nil, fmt.Errorf("could not create part file: %v", err)
		}

		// Create a new encoder
		encoder := wav.NewEncoder(partFile, buf.Format.SampleRate, int(decoder.BitDepth), buf.Format.NumChannels, 1)

		if err := encoder.Write(partBuffer); err != nil {
			return nil, fmt.Errorf("could not write part buffer: %v", err)
		}

		if err := encoder.Close(); err != nil {
			return nil, fmt.Errorf("could not close encoder: %v", err)
		}

		partFile.Close()
		fmt.Printf("Slice %d saved as %s\n", i+1, outputFilePath)
		sampleFilePath := fmt.Sprintf("Samples/%s", partFileName)
		result[i] = ablmodels.AudioFile{
			FilePath: &sampleFilePath,
			Duration: sampleDurationMilliseconds,
		}
	}

	return &result, nil
}
