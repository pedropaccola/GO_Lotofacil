package lotto

type Lotto int

const (
	Lotofacil Lotto = iota
	Lotomania
	Megasena
	Quina
)

func (l Lotto) String() string {
	return [...]string{
		"Lotofacil",
		"Lotomania",
		"Megasena",
		"Quina"}[l]
}

type GameSettings struct {
	Lt      Lotto
	Bet     int
	MaxEven int
	MaxOdd  int
	MaxSeq  int
}
