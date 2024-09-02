package main

import (
	"bytes"
	"fmt"
	"os"
	"os/exec"
	"strings"
)

// Function to compute the MD5 hash of a specific stream
func computeStreamMD5(filePath string, streamSpecifier string) (string, error) {
	cmd := exec.Command("ffmpeg", "-loglevel", "error", "-i", filePath, "-map", streamSpecifier, "-f", "md5", "-")
	var out bytes.Buffer
	cmd.Stdout = &out

	err := cmd.Run()
	if err != nil {
		return "", err
	}

	// The output format is "MD5=xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx"
	output := strings.TrimSpace(out.String())
	if strings.HasPrefix(output, "MD5=") {
		return strings.TrimPrefix(output, "MD5="), nil
	}

	return "", fmt.Errorf("unexpected output: %s", output)
}

// Function to count the number of streams of a particular type in a file
func countStreams(filePath string, streamType string) (int, error) {
	cmd := exec.Command("ffprobe", "-loglevel", "error", "-select_streams", streamType, "-show_entries", "stream=index", "-of", "csv=p=0", filePath)
	output, err := cmd.Output()
	if err != nil {
		return 0, err
	}
	// Count the number of lines in the output, each representing a stream
	return len(strings.Split(strings.TrimSpace(string(output)), "\n")), nil
}

// Function to check if two files are identical in terms of their streams
func areFilesIdentical(file1, file2 string) (bool, error) {
	// Count the number of video streams
	numVideoStreams1, err := countStreams(file1, "v")
	if err != nil {
		return false, err
	}
	numVideoStreams2, err := countStreams(file2, "v")
	if err != nil {
		return false, err
	}

	// Count the number of audio streams
	numAudioStreams1, err := countStreams(file1, "a")
	if err != nil {
		return false, err
	}
	numAudioStreams2, err := countStreams(file2, "a")
	if err != nil {
		return false, err
	}

	// Compare the number of streams
	if numVideoStreams1 != numVideoStreams2 || numAudioStreams1 != numAudioStreams2 {
		return false, nil
	}

	// Compare the video streams
	for i := 0; i < numVideoStreams1; i++ {
		streamSpecifier := fmt.Sprintf("0:v:%d", i)
		hash1, err := computeStreamMD5(file1, streamSpecifier)
		if err != nil {
			return false, err
		}
		hash2, err := computeStreamMD5(file2, streamSpecifier)
		if err != nil {
			return false, err
		}
		if hash1 != hash2 {
			return false, nil
		}
	}

	// Compare the audio streams
	for i := 0; i < numAudioStreams1; i++ {
		streamSpecifier := fmt.Sprintf("0:a:%d", i)
		hash1, err := computeStreamMD5(file1, streamSpecifier)
		if err != nil {
			return false, err
		}
		hash2, err := computeStreamMD5(file2, streamSpecifier)
		if err != nil {
			return false, err
		}
		if hash1 != hash2 {
			return false, nil
		}
	}

	return true, nil
}

func main() {
	if len(os.Args) != 3 {
		fmt.Println("Usage: mediastreamcmp <file1> <file2>")
		return
	}

	file1 := os.Args[1]
	file2 := os.Args[2]

	identical, err := areFilesIdentical(file1, file2)
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
		return
	}

	if identical {
		fmt.Println("The files are identical in terms of audio and video streams.")
		os.Exit(0)
	} else {
		fmt.Println("The files are not identical.")
		os.Exit(2)
	}
}
