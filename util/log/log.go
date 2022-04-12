package log

import (
	"github.com/RyaoChengfeng/wzj-checkin/config"
	rotatelogs "github.com/lestrrat-go/file-rotatelogs"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"os"
	"path"
	"time"
)

// Logger is a global variable
//  import . "github.com/casbin/casnode/util"
//  Logger.Info("msg")
//  Logger.Debug("msg")
//  Logger.Warn("msg")
//  Logger.Error("msg")
//  Logger.Fatal("msg")
var Logger *zap.SugaredLogger

func init() {
	Logger = getLogger()
	if Logger == nil {
		panic("Logger initialization failed")
	}
	Logger.Info("Logger initialization succeeded!")
}

func getLogger() *zap.SugaredLogger {
	// print log to console
	consoleEncoder := zapcore.NewConsoleEncoder(zap.NewDevelopmentEncoderConfig())
	consoleWriter := zapcore.Lock(os.Stdout)
	// write log to file
	encoder := getEncoder()
	writeSyncer := getLogWriter()
	// zapcore.DebugLevel set default log level to DEBUG
	var allCore []zapcore.Core

	if config.C.Debug {
		allCore = append(allCore, zapcore.NewCore(consoleEncoder, consoleWriter, zapcore.DebugLevel))
		allCore = append(allCore, zapcore.NewCore(encoder, writeSyncer, zapcore.DebugLevel))
	} else {
		allCore = append(allCore, zapcore.NewCore(encoder, writeSyncer, zapcore.InfoLevel))
	}
	core := zapcore.NewTee(allCore...)

	// zap.AddCaller() add the feature to log calling function information to the log.
	logger := zap.New(core, zap.AddCaller())

	return logger.Sugar()
}

func getEncoder() zapcore.Encoder {
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder // use ISO8601TimeEncoder

	// Use uppercase letters to record log levels in log files
	encoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder
	// NewJsonEncoder return a zapcore.Encoder that writes key-value pairs in a JSON format.
	return zapcore.NewJSONEncoder(encoderConfig)
}

func getLogWriter() zapcore.WriteSyncer {
	var logPath, logFile string
	if config.C.LogConf.LogPath == "" {
		logPath = "../logs"
	} else {
		logPath = config.C.LogConf.LogPath
	}
	if config.C.LogConf.LogFile == "" {
		logFile = "qqbot.log"
	} else {
		logFile = config.C.LogConf.LogFile
	}

	baseLogPath := path.Join(logPath, logFile)

	// create log file
	_, err := os.Stat(baseLogPath)
	if err != nil {
		if os.IsNotExist(err) {
			err := os.MkdirAll(logPath, os.ModePerm)
			if err != nil {
				panic(err)
			}
		}
	}
	writer, err := rotatelogs.New(
		baseLogPath+".%Y-%m-%d-%H-%M",
		rotatelogs.WithLinkName(baseLogPath),      // generate a soft link pointing to the latest log file
		rotatelogs.WithMaxAge(30*24*time.Hour),    // maximum file save time, 30 days
		rotatelogs.WithRotationTime(24*time.Hour), // log rotating interval
	)
	if err != nil {
		panic(err)
	}

	return zapcore.AddSync(writer)
}
