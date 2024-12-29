package main

import (
	"net/http"
	"time"

	"github.com/dogmatiq/ferrite"
	log "github.com/sirupsen/logrus"

	"github.com/putsi/paparazzogo"
)

func main() {
	var serverPort = ferrite.String("SERVER_PORT", "Server port").WithDefault("8080").Required()
	var streamUrl = ferrite.URL("STREAM_URL", "Stream url").Required()
	var streamUser = ferrite.String("STREAM_USER", "Stream authentication user").WithDefault("").Required()
	var streamPassword = ferrite.String("STREAM_PASSWORD", "Stream authentication password").WithDefault("").Required()

	log.Infof("Using server port %s", serverPort.Value())
	log.Infof("Using stream %s (%s / %s)", streamUrl.Value(), streamUser.Value(), streamPassword.Value())
	// Local server settings
	imgPath := "/img.jpg"

	// If there is zero GET-requests for 30 seconds, mjpeg-stream will be closed.
	// Streaming will be reopened after next request.
	timeout := 30 * time.Second

	mjpegHandler := paparazzogo.NewMjpegproxy()
	mjpegHandler.OpenStream(streamUrl.Value().String(), streamUser.Value(), streamPassword.Value(), timeout)

	http.Handle(imgPath, mjpegHandler)

	s := &http.Server{
		Addr:    ":" + serverPort.Value(),
		Handler: mjpegHandler,
		// Read- & Write-timeout prevent server from getting overwhelmed in idle connections
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	log.Fatal(s.ListenAndServe())

	block := make(chan bool)
	// time.Sleep(time.Second * 30)
	// mp.CloseStream()
	// mp.OpenStream(newMjpegstream, newUser, newPass, newTimeout)
	<-block

}
