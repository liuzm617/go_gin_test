package logging

import (
	"fmt"
	"go_gin_example/utils/common"
	"go_gin_example/utils/conf"
	"log"
	"os"
	"path"
	"path/filepath"
	"runtime"
	"strings"
	"time"
)

var (
	F *os.File
	DefaultPrefix = ""
	DefaultCallerDepth = 2

	logger *log.Logger
	logPrefix = ""
	levels = []string{"DEBUG", "INFO", "WARN", "ERROR", "FATAL"}
)

const (
	DEBUG = 0
	INFO = 1
	WARNING = 2
	ERROR = 3
	FATAL = 4
)

func Init(){
	var err error
	logName := getLogFileName()
	logDir := getLogPath()
	F, err = openLog(logDir, logName)
	if err != nil{
		log.Fatal("init log err:%v", err)
	}

	logger = log.New(F, DefaultPrefix, log.LstdFlags)
}

func getLogPath() string{
	return conf.AppConf.RuntimeRootPath + conf.AppConf.LogSavePath
}

func getLogFileName() string{
	return fmt.Sprintf("%s%s.%s",
		conf.AppConf.LogSaveName,
		time.Now().Format(conf.AppConf.TimeFormat),
		conf.AppConf.LogFileExt,
		)
}
func openLog(filePath, fileName string)(*os.File, error){
	var fullPath string
	var logDir string
	// fullPath = path.Join(filePath, fileName)
	if !(strings.HasPrefix(filePath, "/") || strings.Index(filePath, ":") == 1){
		// not abs path ,linux or windows
		pwd, err := os.Getwd()
		if err != nil{
			return nil, err
		}
		logDir = path.Join(pwd, filePath)

	}else{
		logDir = filePath
	}
	if common.NotPermission(logDir){
		return nil, fmt.Errorf("Permission denied for path:%s", logDir)
	}

	// path dose not exists
	if !common.IsExists(logDir){
		err := common.Mkdir(logDir)
		if err != nil{
			return nil, fmt.Errorf("Mkdir path: %s, err: %v", logDir, err)
		}
	}
	fullPath = path.Join(logDir, fileName)
	f, err := os.OpenFile(fullPath, os.O_APPEND|os.O_CREATE|os.O_RDWR,os.ModePerm)
	if err != nil{
		return nil, fmt.Errorf("Fail to open file:%s", fullPath)
	}
	return f, nil
}


// Debug output logs at debug level
func Debug(v ...interface{}) {
	setPrefix(DEBUG)
	logger.Println(v)
}

// Info output logs at info level
func Info(v ...interface{}) {
	setPrefix(INFO)
	logger.Println(v)
}

// Warn output logs at warn level
func Warn(v ...interface{}) {
	setPrefix(WARNING)
	logger.Println(v)
}

// Error output logs at error level
func Error(v ...interface{}) {
	setPrefix(ERROR)
	logger.Println(v)
}

// Fatal output logs at fatal level
func Fatal(v ...interface{}) {
	setPrefix(FATAL)
	logger.Fatalln(v)
}

// setPrefix set the prefix of the log output
func setPrefix(level int) {
	_, file, line, ok := runtime.Caller(DefaultCallerDepth)
	if ok {
		logPrefix = fmt.Sprintf("[%s][%s:%d]", levels[level], filepath.Base(file), line)
	} else {
		logPrefix = fmt.Sprintf("[%s]", levels[level])
	}

	logger.SetPrefix(logPrefix)
}
