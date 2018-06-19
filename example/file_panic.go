package main

import (
	"fmt"
	golog "go-log/app"
	"time"

	gowriter "github.com/joaosoft/go-writer/app"
)

func runTestFilePanic() {
	//
	// stdout fileWriter
	quit := make(chan bool)
	fileWriter := gowriter.NewFileWriter(
		gowriter.WithFileDirectory("./testing"),
		gowriter.WithFileName("dummy_"),
		gowriter.WithFileMaxMegaByteSize(1),
		gowriter.WithFileFlushTime(time.Second*5),
		gowriter.WithFileQuitChannel(quit),
	)

	//
	// log to json
	fmt.Println(":: LOG JSON")
	log := golog.NewLog(
		golog.WithLevel(golog.InfoLevel),
		golog.WithSpecialWriter(fileWriter)).
		With(
			map[string]interface{}{"level": golog.LEVEL, "timestamp": golog.TIMESTAMP, "date": golog.DATE, "time": golog.TIME},
			map[string]interface{}{"service": "log"},
			map[string]interface{}{"name": "joão"})

	// logging...
	start := time.Now()
	sum := 0
	for i := 0; i < 100000; i++ {
		log.Infof("MESSAGE %d", i+1)
		sum += 1

		if i == 50000 {
			panic("FUCKED!")
		}
	}
	elapsed := time.Since(start)
	log.Infof("ELAPSED TIME: %s", elapsed.String())

	<-time.After(time.Second * 10)
	quit <- true
}
