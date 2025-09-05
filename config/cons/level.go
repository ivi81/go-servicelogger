// level.go - содержит константы описывающие набор значений опции для настройки формата логирования
// и перечислимый тип значениями которого могут быть эти константы
package cons

//go:generate stringer -type=Level
//go:generate enummethods -type=Level
////type level string

// Level - тип необходимый для сопоставления набора строковых значений их индексам
// возможные значения для настройки формата логгирования
type Level int

const (
	info = Level(iota)
	debug
	warn
	err
)

const (
	LOGLEVELINFO  = info
	LOGLEVELDEBUG = debug
	LOGLEVELWARN  = warn
	LOGLEVELERR   = err
)
