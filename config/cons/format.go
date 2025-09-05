// format.go - содержит константы описывающие набор значений опции для настройки формата логирования
// и перечислимый тип значениями которого могут быть эти константы
package cons

//go:generate stringer -type=Format
//go:generate enummethods -type=Format
////type level string

// Format - тип необходимый для сопоставления набора строковых значений их индексам
// возможные значения для настройки формата логгирования
type Format int

const (
	text = Format(iota)
	json
)

const (
	TEXTFORMAT = text
	JSONFORMAT = json
)
