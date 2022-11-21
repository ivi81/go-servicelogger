package config

import "gitlab.cloud.gcm/i.ippolitov/go-microconfig/microconfig"

//LoggerCfg - структура данных описывающая информацию о конфигурации модуля логирования микросервисов
type LoggerCfg struct {
	microconfig.LoggerCfg
}

//SetValuesFromEnv загружает значение перменных среды которые имеют префикс заданный в envPrefix
//в структуру ServiceLoggerCfg
func (cfg *LoggerCfg) SetValuesFromEnv(envPrefix string) {
	cfg.LoggerCfg.SetValuesFromEnv(envPrefix)
}
