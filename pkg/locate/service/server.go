package service

import (
	"log"
	"net/http"
	"net/rpc"
	"os"
	"strings"

	"github.com/1000Delta/wifi-locate/pkg/locate/model"
)

// Config let you costom your service
type Config struct {
	LogPath  string
	TargetAP []string
}

// RunDefaultServer help you run default rpc server for the service
func RunDefaultServer(cfg Config) {
	if logPath := strings.TrimSpace(cfg.LogPath); logPath != "" {
		initLogger(logPath)
	} else {
		initLogger("")
	}

	// init module
	model.InitDB(log.Default())
	InitAPConvertor(cfg.TargetAP)

	// init rpc
	rpc.Register(&LocateService{})
	rpc.HandleHTTP()

	// init server
	http.ListenAndServe(":52201", nil)
}

func initLogger(path string) {
	// 默认使用 stdout
	if path == "" {
		log.SetOutput(os.Stdout)
		return
	}

	logOutput, err := os.OpenFile(path+"run.log", os.O_WRONLY|os.O_APPEND|os.O_CREATE, os.ModeAppend|os.ModePerm)
	if err != nil {
		log.Fatal(err)
	}

	// redirect log output to log file
	log.SetOutput(logOutput)
}
