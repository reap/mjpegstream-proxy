package main

import (
	"log"
	"net/http"
	"time"

	"github.com/putsi/paparazzogo"
)

func main() {

	// Local server settings
	imgPath := "/img.jpg"
	addr := ":8080"

	// MJPEG-stream settings
	user := ""
	pass := ""
	// If there is zero GET-requests for 30 seconds, mjpeg-stream will be closed.
	// Streaming will be reopened after next request.
	timeout := 30 * time.Second

	mjpegStream := "http://192.168.64.30:4408/webcam/?action=stream"

	mjpegHandler := paparazzogo.NewMjpegproxy()
	mjpegHandler.OpenStream(mjpegStream, user, pass, timeout)

	http.Handle(imgPath, mjpegHandler)

	s := &http.Server{
		Addr:    addr,
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
