package main

import "context"

type YTDLService interface {
	GetTrack(context.Context, string) (string, string, error)
	ConvertTrack(context.Context, string, string) (string, error)
	SetTrackMeta(context.Context, string, string, string, string, string)([]byte, string, error)
}

type ytdlService struct {}

func (ytdlService) GetTrack(_ context.Context, url string) (string, string, error) {
	return download(url)
}
func (ytdlService) ConvertTrack(_ context.Context, track, title string) (string, error) {
	return convertToMp3(track, title)
}
func (ytdlService) SetTrackMeta(_ context.Context, track, title, artist, album, albumart string)([]byte, string, error){
	return saveMeta(track, title, artist, album, albumart)
}