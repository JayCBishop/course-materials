package logger

import "log"

type Level int64

const (
	Info  Level = 1
	Debug Level = 2
)

const LOG_LEVEL = Debug

func Log(level Level, format string, v ...interface{}) {
	if LOG_LEVEL >= level {
		log.Printf(format, v...)
	}
}

func Logln(level Level, v ...interface{}) {
	if LOG_LEVEL >= level {
		log.Println(v...)
	}
}

func Panicln(level Level, v ...interface{}) {
	if LOG_LEVEL >= level {
		log.Panicln(v...)
	}
}
