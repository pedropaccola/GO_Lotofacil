package lotto

type Lotto int

const (
	Lotofacil Lotto = iota
	Lotomania
	Megasena
	Quina
)

func (l Lotto) String() string {
	switch l {
	case 0:
		return "Lotofacil"
	case 1:
		return "Lotomania"
	case 2:
		return "Mega-Sena"
	case 3:
		return "Quina"
	default:
		return "Other"
	}
}

type GameSettings struct {
	Lt      Lotto
	MaxEven int
	MinEven int
	MaxOdd  int
	MinOdd  int
	MaxSeq  int
	MinSeq  int
}
