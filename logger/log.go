// log.go содержит функционал создания и настройки объекта логера
// в качестве ядра логирования используется библиотека logrus
package logger

import (
	"fmt"
	"io"

	"os"
	"path"
	"runtime"

	"github.com/sirupsen/logrus"
	"gitlab.cloud.gcm/i.ippolitov/go-servicelogger/config"
	"gitlab.cloud.gcm/i.ippolitov/go-servicelogger/config/cons"
)

// NewLogger создает объект логгер
func NewLogger(cfg *config.LoggerCfg) Logger {

	l := logrus.New()

	//Задаем формат логирования
	switch cfg.LogFormat {
	case cons.JSONFORMAT:
		l.Formatter = newJsonFormatter(cfg)
	case cons.TEXTFORMAT:
		l.Formatter = newTextFormatter(cfg)
	}

	l.SetOutput(io.Discard)
	hook := regHook(cfg)
	l.AddHook(hook)
	l.SetLevel(logrus.TraceLevel)

	return Logger{logrus.NewEntry(l)}
}

func newTextFormatter(cfg *config.LoggerCfg) *logrus.TextFormatter {
	formatter := &logrus.TextFormatter{}

	return formatter
}

func newJsonFormatter(cfg *config.LoggerCfg) *logrus.JSONFormatter {
	formatter := &logrus.JSONFormatter{
		CallerPrettyfier: func(frame *runtime.Frame) (function string, file string) {
			filename := path.Base(frame.File)
			return fmt.Sprintf("%s()", frame.Function), fmt.Sprintf("%s:%d", filename, frame.Line)
		},
	}

	return formatter
}

// regHook регистрирует получателей логов (места куда будут отправляться логи сервиса, файл, stdout и т.д.)
func regHook(cfg *config.LoggerCfg) *writerHook {

	hook := writerHook{
		LogLevel: setLogLevel(cfg.LogLevel),
	}

	for _, mode := range cfg.Mode {
		switch mode {
		case cons.LOGMODFILE:
			if cfg.Path != "" {
				file, err := os.OpenFile(cfg.Path, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0640)
				if err != nil {
					panic(err)
				}
				hook.Writer = append(hook.Writer, file)
			}
		case cons.LOGMODSTORAGE:
		case cons.LOGMODSTDIN:
			hook.Writer = append(hook.Writer, os.Stdout)

		}
	}

	return &hook
}

// setLogLevel переводит значения типа описывающего уровни логирования во внутренние значения уровней в logrus
// если список strL пуст то возвращает список всех уровней logrus.AllLevels
func setLogLevel(strL []cons.Level) []logrus.Level {

	if len(strL) == 0 {
		return logrus.AllLevels
	}

	lvl := make([]logrus.Level, len(strL), len(strL))

	for i, strLevel := range strL {
		if level, err := logrus.ParseLevel(strLevel.String()); err == nil {
			lvl[i] = level
		}
	}

	return lvl
}
