package main

import (
	"context"
	"fmt"
	"time"

	"github.com/go-kit/kit/metrics"
	"github.com/go-kit/log"
)

type loggingMiddleware struct {
	logger log.Logger
	next YTDLService
}
type instrumentingMiddleware struct {
	requestCount   metrics.Counter
	requestLatency metrics.Histogram
	countResult    metrics.Histogram
	next           YTDLService
}

func (mw loggingMiddleware) GetTrack(ctx context.Context, s string) (out1 string, out2 string, err error){
	defer func(begin time.Time) {
		mw.logger.Log(
			"method", "gettrack",
			"input", s,
			"outputpath", out1,
			"outputtitle", out2, 
			"err", err,
			"took", time.Since(begin),
		)
	}(time.Now())
	out1, out2, err = mw.next.GetTrack(ctx, s)
	return
}
func (mw instrumentingMiddleware) GetTrack(ctx context.Context, s string) (out1 string, out2 string, err error) {
	defer func(begin time.Time) {
		lvs := []string{"method", "gettrack", "error", fmt.Sprint(err != nil)}
		mw.requestCount.With(lvs...).Add(1)
		mw.requestLatency.With(lvs...).Observe(time.Since(begin).Seconds())
	}(time.Now())

	out1, out2, err = mw.next.GetTrack(ctx, s)
	return
}
func (mw loggingMiddleware) ConvertTrack(ctx context.Context, track, title string) (out1 string, err error){
	defer func(begin time.Time) {
		mw.logger.Log(
			"method", "converttrack",
			"inputpath", track,
			"inputtitle", title,
			"outputpath", out1, 
			"err", err,
			"took", time.Since(begin),
		)
	}(time.Now())
	out1, err = mw.next.ConvertTrack(ctx, track, title)
	return
}
func (mw instrumentingMiddleware) ConvertTrack(ctx context.Context, track, title string) (out1 string, err error){
	defer func(begin time.Time) {
		lvs := []string{"method", "converttrack", "error", fmt.Sprint(err != nil)}
		mw.requestCount.With(lvs...).Add(1)
		mw.requestLatency.With(lvs...).Observe(time.Since(begin).Seconds())
	}(time.Now())

	out1, err = mw.next.ConvertTrack(ctx, track, title)
	return
}
func (mw loggingMiddleware) SetTrackMeta(ctx context.Context, track, title, artist, album, albumart string) (fileout []byte, out1 string, err error){
	defer func(begin time.Time) {
		mw.logger.Log(
			"method", "settrackmeta",
			"inputpath", track,
			"inputtitle", title,
			"inputartist", artist,
			"inputalbum", album,
			"inputalbumartlink", albumart,
			"outputfilename", out1, 
			"err", err,
			"took", time.Since(begin),
		)
	}(time.Now())
	fileout, out1, err = mw.next.SetTrackMeta(ctx, track, title, artist, album, albumart)
	return
}
func (mw instrumentingMiddleware) SetTrackMeta(ctx context.Context, track, title, artist, album, albumart string) (fileout []byte, out1 string, err error){
	defer func(begin time.Time) {
		lvs := []string{"method", "converttrack", "error", fmt.Sprint(err != nil)}
		mw.requestCount.With(lvs...).Add(1)
		mw.requestLatency.With(lvs...).Observe(time.Since(begin).Seconds())
	}(time.Now())

	fileout, out1, err = mw.next.SetTrackMeta(ctx, track, title, artist, album, albumart)
	return
}