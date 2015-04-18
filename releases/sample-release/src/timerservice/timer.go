package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"time"
)

func main() {
	fmt.Printf("starting...\n")
	logfile_path := "/var/vcap/sys/log/timerservice/timer.log"
	// logfile_path := "timer.log"
	logfile, err := os.OpenFile(logfile_path, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		fmt.Printf("Error happens: %v, I quit\n", err)
		os.Exit(1)
	}

	defer logfile.Close()
	var interval int
	flag.IntVar(&interval, "interval", 1, "beep interval in second")
	flag.Parse()
	logger := log.New(logfile, "timersrv:", log.LstdFlags)
	logger.Printf("Interval as %d\n", interval)
	for {
		time.Sleep(time.Duration(interval) * time.Second)
	}
}
