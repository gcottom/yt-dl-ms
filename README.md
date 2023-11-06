# yt-dl-ms
Youtube downloader go microservice

Allows for the download and conversion of audio from youtube. This is a backend service built to work with my frontend yt-dl-ui-web. This service is built using go-kit, youtube/v2, and my audiometa package. Has a dependency on ffmpeg (used to convert yt videos to mp3 audio).