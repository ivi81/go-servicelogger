//config.go расширяет функционал go-microconfig в части настроек логирования
package config

import "gitlab.cloud.gcm/i.ippolitov/go-microconfig/microconfig"

//ServiceLoggerCfg - структура данных описывающая информацию о конфигурации модуля логирования микросервисов
type ServiceLoggerCfg struct {
	microconfig.LoggerCfg
}

//SetValuesFromEnv загружает значение перменных среды которые имеют префикс заданный в envPrefix
//в структуру ServiceLoggerCfg
func (cfg *ServiceLoggerCfg) SetValuesFromEnv(envPrefix string) {
	cfg.LoggerCfg.SetValuesFromEnv(envPrefix)
}
