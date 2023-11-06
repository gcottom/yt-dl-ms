package main

import (
	"net/http"
	"os"

	kitprometheus "github.com/go-kit/kit/metrics/prometheus"
	httptransport "github.com/go-kit/kit/transport/http"
	"github.com/go-kit/log"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"

	"github.com/rs/cors"
)

func main() {
	logger := log.NewLogfmtLogger(os.Stderr)
	
	fieldKeys := []string{"method", "error"}
	requestCount := kitprometheus.NewCounterFrom(prometheus.CounterOpts{
		Namespace: "ytdl_group",
		Subsystem: "ytdl_service",
		Name:      "request_count",
		Help:      "Number of requests received.",
	}, fieldKeys)
	requestLatency := kitprometheus.NewSummaryFrom(prometheus.SummaryOpts{
		Namespace: "ytdl_group",
		Subsystem: "ytdl_service",
		Name:      "request_latency_microseconds",
		Help:      "Total duration of requests in microseconds.",
	}, fieldKeys)
	countResult := kitprometheus.NewSummaryFrom(prometheus.SummaryOpts{
		Namespace: "ytdl_group",
		Subsystem: "ytdl_service",
		Name:      "count_result",
		Help:      "The result of each count method.",
	}, []string{})
	c := cors.New(cors.Options{
        AllowedOrigins: []string{"*", "null"},
        AllowedMethods: []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
        AllowedHeaders: []string{"Content-Type"},
    })

	var svc YTDLService
	svc = ytdlService{}
	svc = loggingMiddleware{logger, svc}
	svc = instrumentingMiddleware{requestCount, requestLatency, countResult, svc}

	

	getTrackHandler := httptransport.NewServer(
		makeGetTrackEndpoint(svc),
		decodeGetTrackRequest,
		encodeResponse,
	)
	convertTrackHandler := httptransport.NewServer(
		makeConvertTrackEndpoint(svc),
		decodeConvertTrackRequest,
		encodeResponse,
	)
	setTrackMetaHandler := httptransport.NewServer(
		makeSetTrackMetaEndpoint(svc),
		decodeSetTrackMetaRequest,
		encodeResponse,
	)
	http.Handle("/track", c.Handler(getTrackHandler))
	http.Handle("/convert", c.Handler(convertTrackHandler))
	http.Handle("/meta", c.Handler(setTrackMetaHandler))
	http.Handle("/metrics", promhttp.Handler())
	http.ListenAndServe(":8080", nil)
}

