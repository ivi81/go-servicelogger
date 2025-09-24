// mode.go - содержит константы описывающие набор значений опции для настройки режима логирования
// и перечислимый тип значениями которого могут быть эти константы
package cons

//go:generate stringer -type=Mode
//go:generate enummethods -type=Mode
////type level string

// Mode - тип необходимый для сопоставления набора строковых значений их индексам
// возможные значения для настройки формата логгирования
type Mode uint32

const (
	stdin = Mode(iota)
	file
	storage
)

const (
	LOGMODSTDIN   = stdin
	LOGMODFILE    = file
	LOGMODSTORAGE = storage
)
