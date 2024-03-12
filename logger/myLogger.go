package logger

import (
	"fmt"
	"strings"
	"sync"
	"time"

	"github.com/pkg/errors"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/pkgerrors"
	"gopkg.in/natefinch/lumberjack.v2"
)

type ILogger interface {
	Info(message string)
	Debug(message string)
	Warn(message string)
	Error(err error, message string)
	Panic(err error, message string)
}

type loggerImpl struct {
	log zerolog.Logger
}

func (impl *loggerImpl) Info(message string) {
	impl.log.Info().Msg(message)
}

func (impl *loggerImpl) Warn(message string) {
	impl.log.Warn().Msg(message)
}

func (impl *loggerImpl) Debug(message string) {
	impl.log.Debug().Msg(message)
}

func (impl *loggerImpl) Error(err error, message string) {
	err = errors.Wrap(errors.New(err.Error()), message) //To get call stack
	impl.log.Error().Stack().Err(err).Msg("")
}

func (impl *loggerImpl) Panic(err error, message string) {
	err = errors.Wrap(errors.New(err.Error()), message)
	impl.log.Error().Stack().Err(err).Msg("") //To get call stack
}

const (
	logFileName      = "d:\\NewsFeed\\Logs\\logfile.log"
	logFileMaxSize   = 10 //in MB
	logFileMaxBackup = 5  //in MB to retain old logs
	logFileMaxAge    = 14 // in days
	defaultLogLevel  = zerolog.DebugLevel
)

var once sync.Once
var instance *loggerImpl

func Get() ILogger {
	once.Do(func() {
		zerolog.ErrorStackMarshaler = pkgerrors.MarshalStack
		lLog := &lumberjack.Logger{
			Filename:   logFileName,
			MaxSize:    logFileMaxSize,
			MaxBackups: logFileMaxBackup,
			MaxAge:     logFileMaxAge,
		}

		defer lLog.Close()

		fileLogger := zerolog.NewConsoleWriter()
		fileLogger.Out = lLog
		fileLogger.NoColor = true
		fileLogger.TimeFormat = time.RFC3339 //Format of prnting Timesamp
		fileLogger.FormatLevel = func(i interface{}) string {
			return strings.ToUpper(fmt.Sprintf("[%s]", i))
		}

		multi := zerolog.MultiLevelWriter(fileLogger)
		log := zerolog.New(multi).Level(defaultLogLevel).With().Timestamp().Logger()

		instance = &loggerImpl{log: log}

	})

	return instance
}
