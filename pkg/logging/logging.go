package logging

import (
	"fmt"
	"gin-blog/util"
	"log"
	"os"
	"path/filepath"
	"runtime"
	"time"
)

var (
	F *os.File

	DefaultPrefix      = ""
	DefaultCallerDepth = 2

	logger     *log.Logger
	logPrefix  = ""
	levelFlags = []string{"DEBUG", "INFO", "WARN", "ERROR", "FATAL"}

	LogSavePath = "runtime/logs/"
	LogSaveName = "log"
	LogFileExt  = "log"
	TimeFormat  = "20060102"
)

// 定义类型
type Level int

// 定义日志级别
const (
	DEBUG Level = iota
	INFO
	WARNING
	ERROR
	FATAL
)

func init() {
	filePath := LogSavePath + LogSaveName + time.Now().Format(TimeFormat) + "." + LogFileExt
	_, err := os.Stat(filePath)
	switch {
	case os.IsNotExist(err):
		err = util.MkDir(LogSavePath) // 目录不存在，创建日志目录
	case os.IsPermission(err):
		log.Fatalf("Permission :%v", err) // 权限不足
	}
	handle, err := os.OpenFile(filePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatalf("Fail to OpenFile :%v", err)
	}
	logger = log.New(handle, DefaultPrefix, log.LstdFlags)
}

func Debug(v ...interface{}) {
	setPrefix(DEBUG)
	logger.Println(v)
}

func Info(v ...interface{}) {
	setPrefix(INFO)
	logger.Println(v)
}

func Warn(v ...interface{}) {
	setPrefix(WARNING)
	logger.Println(v)
}

func Error(v ...interface{}) {
	setPrefix(ERROR)
	logger.Println(v)
}

func Fatal(v ...interface{}) {
	setPrefix(FATAL)
	logger.Fatalln(v)
}

func setPrefix(level Level) {
	_, file, line, ok := runtime.Caller(DefaultCallerDepth)
	if ok {
		logPrefix = fmt.Sprintf("[%s][%s:%d]", levelFlags[level], filepath.Base(file), line)
	} else {
		logPrefix = fmt.Sprintf("[%s]", levelFlags[level])
	}
	logger.SetPrefix(logPrefix)
}
