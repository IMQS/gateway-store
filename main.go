package main

import (
	"flag"
	"fmt"
	"github.com/IMQS/gateway-store/config"
	"github.com/IMQS/gateway-store/globals"
	"os"
)

var configfn = flag.String("c", "", "Config file to be used if not using config service")

func main() {
	os.Exit(realMain())
}

func realMain() int {
	flag.Parse()
	config, err := config.NewConfig(*configfn)
	if err != nil {
		fmt.Printf("Failed to get the service config: %v", err)
		return 1
	}

	globals, err := globals.NewGlobals(config)
	if err != nil {
		globals.Log.Errorf("Service config error: %v", err)
		return 1
	}
	defer globals.DB.Close()

	return 0
}
