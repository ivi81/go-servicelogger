package config

import (
	mcfg "gitlab.cloud.gcm/i.ippolitov/go-microconfig/microconfig"
)

//LoggerCfg - структура данных описывающая информацию о конфигурации модуля логирования микросервисов
/*type LoggerCfg struct {
	microconfig.LoggerCfg
}
*/

//SetValuesFromEnv загружает значение перменных среды которые имеют префикс заданный в envPrefix
//в структуру LoggerCfg
/*func (cfg *LoggerCfg) SetValuesFromEnv(envPrefix string) {
	cfg.LoggerCfg.SetValuesFromEnv(envPrefix)
}
*/
//LoggerCfg -описывает параметры конфигурации логировния
//Поле Mode задает режим работы логера и должно иметь значение из следующего списка:
//- std - вывод сообшений на stdout,
//- file - запись сообщений в файл,
//- service - отправка сообщений удаленной службе логирования
type LoggerCfg struct {
	Mode             []string             `yaml:"mode"`
	Path             string               `yaml:"path"`
	LogService       ClientLogsStorageCfg `yaml:"logService"`
	LogFormat        string               `yaml:"logFormat"`
	LogLevel         []string             `yaml:"logLevel"`
	DisableTimeStamp bool                 `yaml:"disableTimeStamp"`
}

//SetValuesFromEnv загружает в параметры значения перменных окружения среды
func (cfg *LoggerCfg) SetValuesFromEnv(envPrefix string) {

	envPrefix = mcfg.JoinStr(envPrefix, LogEnvPrefix)

	if mode, ok := mcfg.LookupEnvAsSlice(mcfg.JoinStr(envPrefix, "MODE"), mcfg.StrSplitter); ok {
		cfg.Mode = mode
	}
	if path, ok := mcfg.LookupEnv(mcfg.JoinStr(envPrefix, "PATH")); ok {
		cfg.Path = path
	}

	if format, ok := mcfg.LookupEnv(mcfg.JoinStr(envPrefix, "FORMAT")); ok {
		cfg.Path = format
	}

	if level, ok := mcfg.LookupEnvAsSlice(mcfg.JoinStr(envPrefix, "LEVEL"), mcfg.StrSplitter); ok {
		cfg.LogLevel = level
	}

	if disableTimeStamp, ok := mcfg.LookupEnvAsBool(mcfg.JoinStr(envPrefix, "DISABLE_TIME_STAMP")); ok {
		cfg.DisableTimeStamp = disableTimeStamp
	}

	cfg.LogService.SetValuesFromEnv(envPrefix)
}

//ClientLogsStorage описывает параметры подключения к хранилищу логов
type ClientLogsStorageCfg struct {
	mcfg.BasicClientCfg `yaml:",inline"`
}

//SetValuesFromEnv загружает в параметры значения перменных окружения среды
func (cfg *ClientLogsStorageCfg) SetValuesFromEnv(envPrefix string) {

	envPref := mcfg.JoinStr(envPrefix, ClientLogsStoragePrefix)
	cfg.BasicClientCfg.SetValuesFromEnv(envPref)
}
