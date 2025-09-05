// const.go содержит константы необходимые для конфигурирования логера
package logger

//go:generate stringer -type=LogFormat
//go:generate enummethods -type=LogFormat
//type level string

// LogFormat - тип необходимый для сопоставления набора строковых значений их индексам
//type LogFormat int

//const (
//	json = LogFormat(iota)
//	text
//)

//const (
//	JSONFORMAT = json
//	TEXTFORMAT = text
//)

const (
	JSONFORMAT = "json"
	TEXTFORMAT = "text"
)

const (
	MODSTDIN   = "stdin"
	MODFILE    = "file"
	MODSTORAGE = "storage"
)
