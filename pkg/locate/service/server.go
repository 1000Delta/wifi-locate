package service

import (
	"log"
	"net/http"
	"net/rpc"
	"os"
	"strings"
)

// Config let you costom your service
type Config struct {
	LogPath string
}

// RunDefaultServer help you run default rpc server for the service
func RunDefaultServer(cfg Config) {
	if logPath := strings.TrimSpace(cfg.LogPath); logPath != "" {
		initLogger(logPath)
	} else {
		initLogger("./")
	}
	// rpc
	rpc.Register(&Locate{})
	rpc.HandleHTTP()

	// server
	http.ListenAndServe(":52201", nil)
}

func initLogger(path string) {
	logger, err := os.OpenFile(path+"run.log", os.O_WRONLY|os.O_APPEND|os.O_CREATE, os.ModeAppend|os.ModePerm)
	if err != nil {
		log.Fatal(err)
	}

	// redirect log output to log file
	log.SetOutput(logger)
}
