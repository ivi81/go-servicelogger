package logger_test

import (
	"log"
	"os"
	"testing"

	"github.com/joho/godotenv"
	"gitlab.cloud.gcm/i.ippolitov/go-microconfig/v2/env"
	"gitlab.cloud.gcm/i.ippolitov/go-servicelogger/config"
	"gitlab.cloud.gcm/i.ippolitov/go-servicelogger/config/cons"
	"gitlab.cloud.gcm/i.ippolitov/go-servicelogger/logger"
)

var Cfg config.LoggerCfg

func TestMain(m *testing.M) {
	if err := godotenv.Load("./test_data/.test.env"); err != nil {
		log.Println(" No .env file found")
	}
	//Создаем перменную окружения хранящую путь к файлам конфигурации
	EnvKeyConfigPath := env.JoinStr(cons.LogEnvPrefix, "CONFIG_PATH")
	os.Setenv(EnvKeyConfigPath, "./test_data/config")

	//Задаем название среды развертывания
	os.Setenv("STAGE", "test")

	Cfg = config.LoggerCfg{
		Mode:             []cons.Mode{cons.LOGMODSTDIN, cons.LOGMODFILE, cons.LOGMODSTORAGE},
		Path:             "test_data/test_output.log",
		LogLevel:         []cons.Level{cons.LOGLEVELINFO, cons.LOGLEVELWARN, cons.LOGLEVELERR},
		LogFormat:        cons.JSONFORMAT,
		DisableTimeStamp: true,
		LogStorage: config.ClientLogsStorageCfg{
			Host: "zabbix.cloud",
			Port: 3333,
		},
	}

	os.Exit(m.Run())
}

func TestLoggerCfg(t *testing.T) {

	t.Run("TEST0", func(t *testing.T) {

		logger := logger.NewLogger(&Cfg)
		logger.Info("JKKKKK")

	})

}
