package service

import (
	"log"
	"net/http"
	"net/rpc"
	"os"
)

// RunDefaultServer help you run default rpc server for the service
func RunDefaultServer() {
	// rpc
	rpc.Register(&Locate{})
	rpc.HandleHTTP()

	// server
	http.ListenAndServe(":52201", nil)
}

func init() {
	logger, err := os.OpenFile("run.log", os.O_WRONLY|os.O_APPEND|os.O_CREATE, os.ModeAppend|os.ModePerm)
	if err != nil {
		log.Fatal(err)
	}

	// redirect log output to log file
	log.SetOutput(logger)
}
