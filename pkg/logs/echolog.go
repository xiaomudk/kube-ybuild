package logs

import (
	"os"

	"github.com/labstack/gommon/log"
)

const (
	// WriterConsole console输出
	WriterConsole = "console"
	// WriterFile 文件输出
	WriterFile = "file"
)

// For mapping config logger to app logger levels
var echologLevelMap = map[string]log.Lvl{
	"debug": log.DEBUG,
	"info":  log.INFO,
	"warn":  log.WARN,
	"error": log.ERROR,
	"panic": 5,
	"fatal": 6,
}

// NewEchoLogger new echo logger
func NewEchoLogger(cfg *Config) *log.Logger {
	logger := log.New("")
	logger.SetLevel(getLoggerLevel(cfg))
	if cfg.Writers == WriterFile {
		f, err := os.OpenFile(cfg.LoggerFile, os.O_CREATE|os.O_RDWR|os.O_APPEND, 0666)
		if err != nil {
			log.Errorf("cannot open '%s', (%s)", cfg.LoggerFile, err.Error())
			os.Exit(-1)
		}
		logger.SetOutput(f)
	}
	Log = logger
	return logger
}

func getLoggerLevel(cfg *Config) log.Lvl {
	level, exist := echologLevelMap[cfg.Level]
	if !exist {
		return echologLevelMap["debug"]
	}

	return level
}
