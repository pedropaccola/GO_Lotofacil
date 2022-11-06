package lotto

import (
	"fmt"
	"math/rand"
	"sort"
	"strconv"
	"strings"
	"time"
)

type Game struct {
	GameSettings

	min     int
	max     int
	bet     int
	bets    []int
	numbers []int
}

func (g *Game) String() string {
	sa := make([]string, len(g.numbers))
	for i, v := range g.numbers {
		sa[i] = strconv.Itoa(v)
	}
	sl := strings.Join(sa, ", ")

	return fmt.Sprintf("Aposta de %d dezenas: %s\n", g.bet, sl)
}

func NewGame(s GameSettings) *Game {
	g := &Game{}

	switch g.Lt {
	case 0: //lotofacil
		g.min = 1
		g.max = 25
		g.bets = []int{15, 16, 17, 18, 19, 20}
	case 1: //lotomania
		g.min = 1
		g.max = 100
		g.bets = []int{50}
	case 2: //megasena
		g.min = 1
		g.max = 60
		g.bets = []int{6, 7, 8, 9, 10, 11, 12, 13, 14, 15}
	case 3: //quina
		g.min = 1
		g.max = 80
		g.bets = []int{5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15}
	default: //erro
	}

	g.numbers = make([]int, 0, g.max)

	return g
}

func (g *Game) CheckBet(bet int) (bool, error) {
	for _, v := range g.bets {
		if bet == v {
			g.bet = bet
			return true, nil
		}
	}
	return false, fmt.Errorf("impossivel apostar (%d) dezenas", bet)
}

func (g *Game) Game() error {

	for start := time.Now(); time.Since(start) < (time.Second * 5); {

		g.generateNumbers()

		rand.Shuffle(len(g.numbers), func(i, j int) {
			g.numbers[i], g.numbers[j] = g.numbers[j], g.numbers[i]
		})

		g.numbers = g.numbers[:g.bet]

		sort.Ints(g.numbers)

		if b, _ := g.checkOddsEvensSequences(); b {
			return nil
		}

	}
	return fmt.Errorf("nao foi possivel gerar o jogo com as restricoes colocadas")
}

func (g *Game) generateNumbers() {
	for i := g.min; i <= g.max; i++ {
		g.numbers = append(g.numbers, i)
	}
}

func (g *Game) checkOddsEvensSequences() (bool, error) {
	e := 0
	o := 0
	s := 0
	seq := 0

	for i, v := range g.numbers {
		if v%2 == 0 {
			e++
		} else {
			o++
		}

		if i != 0 && g.numbers[i] == g.numbers[i-1]+1 {
			s++
		} else {
			s = 0
		}

		if s > seq {
			seq = s
		}
	}
	seq++

	if !(e <= g.MaxEven && e >= g.MinEven && o <= g.MaxOdd && o >= g.MinOdd) {
		return false, fmt.Errorf("(%d) dezenas pares e (%d) dezenas impares", e, o)
	} else if !(seq <= g.MaxSeq && seq >= g.MinSeq) {
		return false, fmt.Errorf("(%d) dezenas em sequencia", seq)
	}
	return true, nil
}
