package main

import (
	"fmt"
	"go-log/app"
	"os"
	"time"

	gowriter "github.com/joaosoft/go-writer/app"
)

func runTestDefault() {
	//
	// log to text
	fmt.Println(":: LOG TEXT")
	log := golog.NewLog(
		golog.WithLevel(golog.InfoLevel),
		golog.WithFormatHandler(gowriter.TextFormatHandler),
		golog.WithWriter(os.Stdout)).
		With(
			map[string]interface{}{"level": golog.LEVEL, "timestamp": golog.TIMESTAMP, "date": golog.DATE, "time": golog.TIME},
			map[string]interface{}{"service": "log"},
			map[string]interface{}{"name": "joão"})

	// logging...
	log.Error("isto é uma mensagem de error")
	log.Info("isto é uma mensagem de info")
	log.Debug("isto é uma mensagem de debug")
	log.Error("")

	fmt.Println("--------------")
	<-time.After(time.Second)

	//
	// log to json
	fmt.Println(":: LOG JSON")
	log = golog.NewLog(
		golog.WithLevel(golog.InfoLevel),
		golog.WithFormatHandler(gowriter.JsonFormatHandler),
		golog.WithWriter(os.Stdout)).
		With(
			map[string]interface{}{"level": golog.LEVEL, "timestamp": golog.TIMESTAMP, "date": golog.DATE, "time": golog.TIME},
			map[string]interface{}{"service": "log"},
			map[string]interface{}{"name": "joão"})

	// logging...
	log.Errorf("isto é uma mensagem de error %s", "hello")
	log.Infof("isto é uma  mensagem de info %s ", "hi")
	log.Debugf("isto é uma mensagem de debug %s", "ehh")
}
