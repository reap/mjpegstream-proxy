package main

import (
	"net/http"
	"time"

	log "github.com/sirupsen/logrus"

	"github.com/putsi/paparazzogo"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
)

var serverPort string
var streamUrl string
var streamUser string
var streamPassword string

func main() {

	readConfiguration()

	// Local server settings
	imgPath := "/img.jpg"

	// If there is zero GET-requests for 30 seconds, mjpeg-stream will be closed.
	// Streaming will be reopened after next request.
	timeout := 30 * time.Second

	mjpegHandler := paparazzogo.NewMjpegproxy()
	mjpegHandler.OpenStream(streamUrl, streamUser, streamPassword, timeout)

	http.Handle(imgPath, mjpegHandler)

	s := &http.Server{
		Addr:    ":" + serverPort,
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

func readConfiguration() {
	// allow setting values also with commandline flags
	pflag.String("http_port", "8080", "Server port")
	viper.BindPFlag("http_port", pflag.Lookup("http_port"))

	pflag.String("stream_url", "http://192.168.64.30:4408/webcam/?action=stream", "Url of the mjpegstream")
	viper.BindPFlag("stream_url", pflag.Lookup("stream_url"))

	pflag.String("stream_user", "", "Username")
	viper.BindPFlag("stream_user", pflag.Lookup("stream_user"))

	pflag.String("stream_password", "", "Password")
	viper.BindPFlag("stream_password", pflag.Lookup("stream_password"))

	// parse values from environment variables
	viper.AutomaticEnv()

	serverPort = viper.GetString("http_port")
	streamUrl = viper.GetString("stream_url")
	streamUser = viper.GetString("stream_user")
	streamPassword = viper.GetString("stream_password")

	log.Infof("Using server port %s", serverPort)
	log.Infof("Using stream %s (%s / %s)", streamUrl, streamUser, streamPassword)
}
