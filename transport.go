package main

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/go-kit/kit/endpoint"
)



type getTrackRequest struct {
	URL string `json:"url"`
}
type getTrackResponse struct {
	TrackUrl string `json:"url,omitempty"`
	Title string `json:"title,omitempty"`
}
type convertTrackRequest struct {
	TrackUrl string `json:"url"`
	Title string `json:"title"`
}
type convertTrackResponse struct {
	TrackUrl string `json:"url,omitempty"`
}
type setTrackMetaRequest struct {
	TrackUrl string `json:"url"`
	Title string `json:"title"`
	Artist string `json:"artist"`
	Album string `json:"album"`
	AlbumArt string `json:"albumart"`
}
type setTrackMetaResponse struct {
	TrackData []byte `json:"trackdata,omitempty"`
	FileName string `json:"filename,omitempty"`
}
type errorResponse struct {
	Error string `json:"err"`
}

func makeGetTrackEndpoint(svc YTDLService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error){
		req := request.(getTrackRequest)
		url, title, err := svc.GetTrack(ctx, req.URL)
		if err != nil {
			return errorResponse{err.Error()}, nil
		}
		return getTrackResponse{url, title}, nil
	}
}

func makeConvertTrackEndpoint(svc YTDLService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error){
		req := request.(convertTrackRequest)
		track, err := svc.ConvertTrack(ctx, req.TrackUrl, req.Title)
		if err != nil {
			return errorResponse{err.Error()}, nil
		}
		return convertTrackResponse{track}, nil
	}
}

func makeSetTrackMetaEndpoint(svc YTDLService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error){
		req := request.(setTrackMetaRequest)
		trackdata, filename, err := svc.SetTrackMeta(ctx, req.TrackUrl, req.Title, req.Artist, req.Album, req.AlbumArt)
		if err != nil {
			return errorResponse{err.Error()}, nil
		}
		return setTrackMetaResponse{trackdata, filename}, nil
	}
}
func encodeResponse(_ context.Context, w http.ResponseWriter, response interface{}) error {
	switch response.(type){
	case errorResponse: w.WriteHeader(http.StatusBadRequest)
	default: w.WriteHeader(http.StatusOK)
	}
	return json.NewEncoder(w).Encode(response)
}

func decodeGetTrackRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var request getTrackRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		return nil, err
	}
	return request, nil
}
func decodeConvertTrackRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var request convertTrackRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		return nil, err
	}
	return request, nil
}
func decodeSetTrackMetaRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var request setTrackMetaRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		return nil, err
	}
	return request, nil
}