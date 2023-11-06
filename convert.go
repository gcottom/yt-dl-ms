package main

import (
	"os"
	"os/exec"
	"path/filepath"
)

func convertToMp3(inputFilePath string, title string) (string, error) {

	outFile := os.TempDir() + "yt-dl-ui/"
	outFile += title + ".mp3"
	// Specify the output MP3 file path
	outputFilePath := outFile

	// Create a new FFmpeg instance

	var args = []string{"-i", inputFilePath, "-acodec:a", "libmp3lame", "-b:a", "256k", outFile}
	ffmpeg, err := filepath.Abs("/var")
	if err != nil {
		return "", err
	}
	ffmpeg += "/ffmpeg"
	cmd := exec.Command(ffmpeg, args...)
	err = cmd.Run()
	if err != nil {
		return "", err
	}
	err = os.Remove(inputFilePath)
	if err != nil {
		return "", err
	}
	err = os.Chmod(outFile, 0666)
	if err != nil {
		return "", err
	}
	return outputFilePath, nil
}
