package logger

import (
	"errors"
	"io"
	"os"
	"strings"
	"sync"

	"github.com/birchwood-langham/web-service-bootstrap/config"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
)

var once sync.Once
var syncer zapcore.WriteSyncer
var core zapcore.Core
var log *zap.Logger

var CoreNotInitializedError error = errors.New("Zap core has not been initialized")

func ZapConfig() zapcore.EncoderConfig {
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	encoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder

	return encoderConfig
}

func ZapEncoder() zapcore.Encoder {
	return zapcore.NewConsoleEncoder(ZapConfig())
}

func ZapWriter(writer io.Writer) zapcore.WriteSyncer {
	if syncer == nil {
		syncer = zapcore.NewMultiWriteSyncer(zapcore.AddSync(writer), zapcore.AddSync(os.Stdout))
	}

	return syncer
}

func LumberjackLogger(fileName string, maxSize, maxBackups, maxAge int, compress bool) *lumberjack.Logger {
	return &lumberjack.Logger{
		Filename:   fileName,
		MaxSize:    maxSize,
		MaxBackups: maxBackups,
		MaxAge:     maxAge,
		Compress:   compress,
	}
}

func ZapCore() (zapcore.Core, error) {
	if core == nil {
		return nil, CoreNotInitializedError
	}

	return core, nil
}

func New(level zapcore.Level, writer io.Writer) *zap.Logger {
	once.Do(func() {
		core = zapcore.NewCore(ZapEncoder(), ZapWriter(writer), level)
		log = zap.New(core, zap.AddCaller())
	})

	defer func() {
		_ = log.Sync()
	}()

	return log
}

func ConfiguredLumberjackLogger() *lumberjack.Logger {
	return LumberjackLogger(
		viper.GetString(config.LogFilePathKey),
		viper.GetInt(config.LogFileMaxSize),
		viper.GetInt(config.LogFileMaxBackups),
		viper.GetInt(config.LogFileMaxAge),
		viper.GetBool(config.LogFileCompress),
	)
}

func ConsoleLogger() *zap.Logger {
	c := zapcore.NewCore(ZapEncoder(), zapcore.AddSync(os.Stdout), zapcore.InfoLevel)
	return zap.New(c, zap.AddCaller())
}

// ApplicationLogLevel returns the log level defined in the
// application configuration file
func ApplicationLogLevel() zapcore.Level {
	var level zapcore.Level

	switch strings.ToUpper(viper.GetString(config.LogLevelKey)) {
	case "DEBUG":
		level = zapcore.DebugLevel
	case "INFO":
		level = zapcore.InfoLevel
	case "WARN":
		level = zapcore.WarnLevel
	case "ERROR":
		level = zapcore.ErrorLevel
	case "FATAL":
		level = zapcore.FatalLevel
	case "PANIC":
		level = zapcore.PanicLevel
	default:
		level = zapcore.InfoLevel
	}

	return level
}
