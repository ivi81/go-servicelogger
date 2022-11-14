//const.go содержит константы необходимые для конфигурирования логера
package logger

type level string

const (
	JSONFORMAT = "json"
	TEXTFORMAT = "text"
)

const (
	MODSTDIN = "stdin"
	MODFILE  = "file"
	MODSERV  = "service"
)
