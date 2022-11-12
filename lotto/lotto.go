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
	MaxEven int
	MinEven int
	MaxOdd  int
	MinOdd  int
	MaxSeq  int
	MinSeq  int
}
