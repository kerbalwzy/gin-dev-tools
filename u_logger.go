package kerbalwzygo

import (
	"log"
	"os"
	"sync"
)

type Level int

var (
	Debug Level = 0
	Info  Level = 1
	Warn  Level = 2
	Error Level = 3
)

type XLogger struct {
	level Level
	log.Logger
}

func (obj *XLogger) SetLevel(level Level) {
	obj.level = level
}

func (obj *XLogger) Level() Level {
	if obj.level == 0 {
		obj.level = Debug
	}
	return obj.level
}

func (obj *XLogger) Debug(msg ...interface{}) {
	if obj.level > Debug {
		return
	}
	obj.Printf("[DEBUG] %s", msg)
}

func (obj *XLogger) Info(msg ...interface{}) {
	if obj.level > Info {
		return
	}
	obj.Printf("[INFO] %s", msg)
}

func (obj *XLogger) Warn(msg ...interface{}) {
	if obj.level > Warn {
		return
	}
	obj.Printf("[WARN] %s", msg)
}

func (obj *XLogger) Error(msg ...interface{}) {
	if obj.level > Error {
		return
	}
	obj.Printf("[ERROR] %s", msg)
}

var logger *XLogger
var once sync.Once

func GetLogger() *XLogger {
	once.Do(func() {
		logger = &XLogger{}
		logger.SetOutput(os.Stdout)
		logger.SetFlags(log.Ldate | log.Lmicroseconds | log.Lshortfile)
	})
	return logger
}
