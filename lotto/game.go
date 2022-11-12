package lotto

import (
	"fmt"
	"math/rand"
	"sort"
	"strconv"
	"strings"
)

type Game struct {
	BetSettings

	min          int
	max          int
	possibleBets []int
	Numbers      []int
}

func (g *Game) String() string {
	strArr := make([]string, len(g.Numbers))
	for i, v := range g.Numbers {
		strArr[i] = strconv.Itoa(v)
	}
	str := strings.Join(strArr, ", ")

	return fmt.Sprintf("Aposta de %d dezenas: %s\n", g.Bet, str)
}

func NewGame(s BetSettings) (*Game, error) {
	g := &Game{
		BetSettings: s,
	}

	switch g.Lt {
	case 0: //lotofacil
		g.min = 1
		g.max = 25
		g.possibleBets = []int{15, 16, 17, 18, 19, 20}
	case 1: //lotomania
		g.min = 1
		g.max = 100
		g.possibleBets = []int{50}
	case 2: //megasena
		g.min = 1
		g.max = 60
		g.possibleBets = []int{6, 7, 8, 9, 10, 11, 12, 13, 14, 15}
	case 3: //quina
		g.min = 1
		g.max = 80
		g.possibleBets = []int{5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15}
	default: //erro
		return nil, fmt.Errorf("opcao (%s) nao disponivel", g.Lt)
	}
	if err := g.bet(); err != nil {
		return nil, err
	}
	if err := g.betConstraints(); err != nil {
		return nil, err
	}
	return g, nil
}

func (g *Game) bet() error {
	for _, bet := range g.possibleBets {
		if g.Bet == bet {
			fmt.Println("possible bet")
			return nil
		}
	}
	return fmt.Errorf("impossivel apostar (%d) dezenas", g.Bet)
}

func (g *Game) betConstraints() error {
	if g.MaxEven+g.MaxOdd < g.Bet {
		return fmt.Errorf(
			"impossival aposta de (%d) pares e (%d) impares com (%d) dezenas",
			g.MaxEven, g.MaxOdd, g.Bet)
	}
	if g.max-g.min < g.MaxSeq {
		return fmt.Errorf(
			"impossival sequencia de (%d) numeros na %s",
			g.MaxSeq, g.Lt)
	}
	fmt.Println("constraints ok")
	return nil
}

func (g *Game) Generate() error {
	g.Numbers = make([]int, 0, g.Bet)
	found := make(map[int]bool)

	for len(g.Numbers) <= g.Bet {
		i := rand.Intn(g.max) + 1
		fmt.Printf("number generated %d\n", i)
		if found[i] {
			continue
		}
		found[i] = true
		fmt.Printf("number unique %d\n", i)
		if !g.generateConstraints(i) {
			continue
		}
		g.Numbers = append(g.Numbers, i)
		fmt.Printf("number appended %d\n", i)
	}
	sort.Slice(g.Numbers, func(i, j int) bool { return g.Numbers[i] < g.Numbers[j] })
	return nil
}

func (g *Game) generateConstraints(check int) bool {
	e := 0
	o := 0
	s := 0
	seq := 0

	if len(g.Numbers) == 0 {
		return true
	}

	for i, v := range g.Numbers {
		if v%2 == 0 {
			e++
		} else {
			o++
		}

		if i != 0 && g.Numbers[i] == g.Numbers[i-1]+1 {
			s++
		} else {
			s = 0
		}

		if s > seq {
			seq = s
		}
	}
	seq++

	if e <= g.MaxEven && o <= g.MaxOdd && seq <= g.MaxSeq {
		return true
	}
	return false
}

// func (g *Game) generateNumbers() {
// 	for i := g.min; i <= g.max; i++ {
// 		g.numbers = append(g.numbers, i)
// 	}
// }
