// loggercfg.go - содержит описание структуры для параметров конфигурации
package config

import (
	"github.com/ivi81/enummethods/enumerator"
	"gitlab.cloud.gcm/i.ippolitov/go-servicelogger/config/cons"
)

// LoggerCfg -описывает параметры конфигурации логировния
// Поле Mode задает режим работы логера и должно иметь значение из следующего списка:
// - std - вывод сообшений на stdout,
// - file - запись сообщений в файл,
// - service - отправка сообщений удаленной службе логирования
type LoggerCfg struct {
	Mode             []cons.Mode          `yaml:"mode" env:"MODE"` //Режим в котором работает логер, stdin, file ,storage
	Path             string               `yaml:"path" env:"PATH"`
	LogStorage       ClientLogsStorageCfg `yaml:"logStorage" env:"STORAGE"`
	LogFormat        cons.Format          `yaml:"logFormat" env:"FORMAT"` //В каком формате пишется лог: txt, json
	LogLevel         []cons.Level         `yaml:"logLevel" env:"LEVEL"`   //Уровни логирования. Могут быть info, debug, warn, err
	DisableTimeStamp bool                 `yaml:"disableTimeStamp" env:"DISABLE_TIME_STAMP"`
}

// ClientLogsStorage описывает параметры подключения к хранилищу логов
type ClientLogsStorageCfg struct {
	Service string `yaml:"sevice" env:"SERVICE"`
	Host    string `yaml:"host" env:"HOST"`
	Port    int    `yaml:"port" env:"PORT"`
}

type Enumerator interface {
	enumerator.Enumerator
}
