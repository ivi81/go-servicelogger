package config_test

import (
	"log"
	"os"
	"testing"

	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"

	"gitlab.cloud.gcm/i.ippolitov/go-microconfig/v2"
	"gitlab.cloud.gcm/i.ippolitov/go-microconfig/v2/env"
	"gitlab.cloud.gcm/i.ippolitov/go-servicelogger/config"
	"gitlab.cloud.gcm/i.ippolitov/go-servicelogger/config/cons"
)

func TestMain(m *testing.M) {
	if err := godotenv.Load("./test_data/.test.env"); err != nil {
		log.Println(" No .env file found")
	}
	//Создаем перменную окружения хранящую путь к файлам конфигурации
	EnvKeyConfigPath := env.JoinStr(cons.LogEnvPrefix, "CONFIG_PATH")
	os.Setenv(EnvKeyConfigPath, "./test_data/config")

	//Задаем название среды развертывания
	os.Setenv("STAGE", "test")

	os.Exit(m.Run())
}

func TestLoggerCfg(t *testing.T) {

	t.Run("TEST0 : load Cfg from env", func(t *testing.T) {

		testCfg := config.LoggerCfg{}

		expectedCfg := config.LoggerCfg{
			Mode:             []cons.Mode{cons.LOGMODSTDIN, cons.LOGMODFILE, cons.LOGMODSTORAGE},
			Path:             "another/path/to/file",
			LogLevel:         []cons.Level{cons.LOGLEVELINFO, cons.LOGLEVELWARN, cons.LOGLEVELERR},
			LogFormat:        cons.JSONFORMAT,
			DisableTimeStamp: true,
			LogStorage: config.ClientLogsStorageCfg{
				Host: "zabbix.cloud",
				Port: 3333,
			},
		}

		err := microconfig.CfgLoad(&testCfg, cons.LogEnvPrefix, false)

		if assert.NoError(t, err) {
			assert.EqualValues(t, expectedCfg, testCfg)
		} else {
			assert.ErrorContains(t, err, "no such file or directory")
		}
	})

}
