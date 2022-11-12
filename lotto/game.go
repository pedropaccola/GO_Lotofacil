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
	bets    []int
	unique  map[int]bool
	numbers []int
}

func (g *Game) String() string {
	strArr := make([]string, len(g.numbers))
	for i, v := range g.numbers {
		strArr[i] = strconv.Itoa(v)
	}
	str := strings.Join(strArr, ", ")

	return fmt.Sprintf("Aposta de %d dezenas: %s\n", g.Bet, str)
}

func NewGame(s GameSettings) (*Game, error) {
	g := &Game{
		GameSettings: s,
	}

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
	if err := g.bet(); err != nil {
		return nil, err
	}
	if err := g.constraints(); err != nil {
		return nil, err
	}
	return g, nil
}

func (g *Game) bet() error {
	for _, bet := range g.bets {
		if g.Bet == bet {
			return nil
		}
	}
	return fmt.Errorf("impossivel apostar (%d) dezenas", g.Bet)
}

func (g *Game) constraints() error {
	oddCount := 0
	evenCount := 0
	for i := g.min; i <= g.max; i++ {
		if i%2 == 0 {
			evenCount++
		} else {
			oddCount++
		}
	}
	if g.MaxEven+g.MaxOdd < g.Bet {
		return fmt.Errorf(
			"impossival aposta de (%d) pares e (%d) impares com (%d) dezenas",
			g.MaxEven, g.MaxOdd, g.Bet)
	}
	if g.MaxEven > evenCount || g.MaxOdd > oddCount {
		return fmt.Errorf(
			"impossival aposta de (%d) pares e (%d) impares na %s",
			g.MaxEven, g.MaxOdd, g.Lt)
	}
	if g.max-g.min < g.MaxSeq {
		return fmt.Errorf(
			"impossival sequencia de (%d) numeros na %s",
			g.MaxSeq, g.Lt)
	}
	return nil
}

func (g *Game) Generate() error {
	g.numbers = make([]int, g.Bet)
	g.unique = make(map[int]bool)

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
