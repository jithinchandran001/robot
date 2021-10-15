package main

import (
	"flag"
	"fmt"
	"robot/bin"
)

func main() {
	fmt.Println("App is booting...")
	bin.Run(parseArgs())
}

func parseArgs() (string, string) {
	var mode, env string

	flag.StringVar(&env, "env", ".env", "Specify environment variable config file")
	flag.StringVar(&mode, "mode", "http", "Application running mode in http or grpc")

	flag.Parse()
	return mode, env
}
