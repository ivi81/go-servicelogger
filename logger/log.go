//log.go содержит функционал создания и настройки объекта логера
//в качестве ядра логирования используется библиотека logrus
package logger

import (
	"fmt"
	"io"

	"os"
	"path"
	"runtime"

	"github.com/sirupsen/logrus"
	"gitlab.cloud.gcm/i.ippolitov/go-servicelogger/config"
)

//NewLogger создает объект логгер
func NewLogger(cfg *config.ServiceLoggerCfg) Logger {

	l := logrus.New()
	switch cfg.LogFormat {
	case JSONFORMAT:
		l.Formatter = newJsonFormatter(cfg)
	case TEXTFORMAT:
		l.Formatter = newTextFormatter(cfg)
	}

	l.SetOutput(io.Discard)
	hook := regHook(cfg)
	l.AddHook(hook)
	l.SetLevel(logrus.TraceLevel)

	return Logger{logrus.NewEntry(l)}
}

func newTextFormatter(cfg *config.ServiceLoggerCfg) *logrus.TextFormatter {
	formatter := &logrus.TextFormatter{}

	return formatter
}

func newJsonFormatter(cfg *config.ServiceLoggerCfg) *logrus.JSONFormatter {
	formatter := &logrus.JSONFormatter{
		CallerPrettyfier: func(frame *runtime.Frame) (function string, file string) {
			filename := path.Base(frame.File)
			return fmt.Sprintf("%s()", frame.Function), fmt.Sprintf("%s:%s", filename, frame.Line)
		},
	}

	return formatter
}

//regHook производит получателей логов (места куда будут отправляться логи сервиса, файл, stdout и т.д.)
func regHook(cfg *config.ServiceLoggerCfg) *writerHook {

	hook := writerHook{
		LogLevel: setLogLevel(cfg.LogLevel),
	}

	if ok := find(cfg.Mode, MODFILE); ok {
		if cfg.Path != "" {
			file, err := os.OpenFile(cfg.Path, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0640)
			if err != nil {
				panic(err)
			}
			hook.Writer = append(hook.Writer, file)
		}
	}

	if ok := find(cfg.Mode, MODSTDIN); ok {
		hook.Writer = append(hook.Writer, os.Stdout)
	}

	return &hook
}

//setLogLevel переводит значения строк описывающих уровни логирования во внутренние значения данных уровней в logrus
//если список strL пуст то возвращает список всех уровней logrus.AllLevels
func setLogLevel(strL []string) []logrus.Level {

	var l []logrus.Level
	for _, strLevel := range strL {
		if level, err := logrus.ParseLevel(strLevel); err == nil {
			l = append(l, level)
		}
	}
	if len(l) == 0 {
		return logrus.AllLevels
	}
	return l
}

//find поиск строки в срезе строк (вспомогательная функция)
func find(s []string, x string) bool {
	for _, str := range s {
		if str == x {
			return true
		}
	}
	return false
}
