package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"

	tag "github.com/gcottom/audiometa"
)

func saveMeta(trackUrl, title, artist, album, albumart string) ([]byte, string, error) {
	coverFileName := os.TempDir() + fmt.Sprintf("yt-dl-ui/%s+%s+cover.jpg", artist, title)
	url := albumart
	idTag, err := tag.OpenTag(trackUrl)
	if err != nil {
		return nil, "", err
	}
	if url != "" {
		response, err := http.Get(url)
		if err != nil {
			return nil, "", err
		}
		defer response.Body.Close()
		file, err := os.Create(coverFileName)
		if err != nil {
			return nil, "", err
		}
		defer file.Close()
		_, err = io.Copy(file, response.Body)
		if err != nil {
			return nil, "", err
		}
		idTag.SetAlbumArtFromFilePath(coverFileName)
	}

	idTag.SetTitle(title)
	idTag.SetAlbum(album)
	idTag.SetArtist(artist)
	idTag.Save()
	os.Remove(coverFileName)
	var singleArtist string
	if len(strings.Split(artist, ",")) > 1 {
		singleArtist = strings.Split(artist, ",")[0]
	} else {
		singleArtist = artist
	}
	newFileName := os.TempDir() + "yt-dl-ui/" + SanitizeFilename(singleArtist) + " - " + SanitizeFilename(title) + ".mp3"
	os.Rename(trackUrl, newFileName)
	outFile, err := os.ReadFile(newFileName)
	if err != nil {
		return nil, "", err
	}
	return outFile, (SanitizeFilename(singleArtist) + " - " + SanitizeFilename(title) + ".mp3"), nil

}
