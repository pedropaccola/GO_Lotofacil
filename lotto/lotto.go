package lotto

type LottoName int

const (
	Lotofacil LottoName = iota
	Lotomania
	Megasena
	Quina
)

func (l LottoName) String() string {
	return [...]string{
		"Lotofacil",
		"Lotomania",
		"Megasena",
		"Quina"}[l]
}

type BetSettings struct {
	Lt      LottoName
	Bet     int
	MaxEven int
	MaxOdd  int
	MaxSeq  int
}
