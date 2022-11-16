package lotto

type LottoName int

const (
	Lotofacil LottoName = iota + 1
	Lotomania
	Megasena
	Quina
	LastLotto
)

func (l LottoName) String() string {
	switch l {
	case 1:
		return "Lotofacil"
	case 2:
		return "Lotomania"
	case 3:
		return "Megasena"
	case 4:
		return "Quina"
	default:
		return "Loteria Invalida"
	}
}
