//type.go содержит описание объекта логера микросервисов
package logger

import (
	"io"

	"github.com/sirupsen/logrus"
)

var e *logrus.Entry

//Logger - структура описывает объект логгер
type Logger struct {
	*logrus.Entry
}

//GetLoggerWithField - возвращает объект логгер вывод которого будет содержать уникальную информацию переданную в k,v
//удобно использовать для персонифицированного логирования конкретного модуля в сервисе
func (l *Logger) GetLoggerWithField(k string, v interface{}) Logger {
	return Logger{l.WithField(k, v)}

}

//GetLoggerWithFields - возвращает объект логер вывод которого будет содержать уникальную информацию перданную ввиде набора полей в fields
func (l *Logger) GetLoggerWithFields(fields map[string]interface{}) Logger {
	return Logger{l.WithFields(fields)}

}

//writerHook структура описывающая объект приемник сообщений логирования
type writerHook struct {
	Writer   []io.Writer
	LogLevel []logrus.Level
}

//Fire
func (hook *writerHook) Fire(entry *logrus.Entry) error {

	line, err := entry.String()

	if err != nil {
		return err
	}

	for _, w := range hook.Writer {
		w.Write([]byte(line))
	}

	return err
}

//Levels возвращает список заданных в конфигурации уровней логирования
func (hook *writerHook) Levels() []logrus.Level {
	return hook.LogLevel
}
