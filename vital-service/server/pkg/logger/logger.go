package logger

import (
	"os"
	"strings"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var Log *zap.SugaredLogger

func Init(env string) {
	var zapLogger *zap.Logger

	logFile := "logs/app.log"
	if _, err := os.Stat("logs"); os.IsNotExist(err) {
		_ = os.Mkdir("logs", 0755)
	}

	fileEncoder := zapcore.NewJSONEncoder(zap.NewProductionEncoderConfig())

	fileWriter := zapcore.AddSync(&lumberjackLogger{
		Filename: logFile,
		MaxSize:  10, // megabytes
		MaxAge:   7,  // days
		Compress: true,
	})

	core := zapcore.NewCore(fileEncoder, fileWriter, zapcore.InfoLevel)

	if strings.ToLower(env) != "production" {
		consoleEncoder := zapcore.NewConsoleEncoder(zap.NewDevelopmentEncoderConfig())
		consoleWriter := zapcore.AddSync(os.Stdout)

		core = zapcore.NewTee(
			zapcore.NewCore(consoleEncoder, consoleWriter, zapcore.DebugLevel),
			zapcore.NewCore(fileEncoder, fileWriter, zapcore.InfoLevel),
		)
	}

	zapLogger = zap.New(core, zap.AddCaller(), zap.AddCallerSkip(1))

	Log = zapLogger.Sugar()
	Log.Infow("logger initialized", "environment", env)
}

func Sync() {
	if Log != nil {
		_ = Log.Sync()
	}
}

// lumberjackLogger wraps a rotating logger
type lumberjackLogger struct {
	Filename string
	MaxSize  int
	MaxAge   int
	Compress bool
}

func (l *lumberjackLogger) Write(p []byte) (n int, err error) {
	f, err := os.OpenFile(l.Filename, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return 0, err
	}
	defer f.Close()

	return f.Write(p)
}
