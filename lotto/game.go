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
	if g.MaxSeq < 1 || g.MaxEven < 0 || g.MaxOdd < 0 {
		return fmt.Errorf(
			"impossival sequencia de 0 digitos ou valores negativos")
	}
	if g.MaxSeq > g.Bet {
		return fmt.Errorf(
			"impossival sequencia de (%d) apostando (%d) dezenas",
			g.MaxSeq, g.Bet)
	}
	return nil
}

func (g *Game) Generate() error {
	g.Numbers = make([]int, 0, g.Bet)
	found := make(map[int]bool)

	for len(g.Numbers) <= g.Bet {
		i := rand.Intn(g.max) + 1
		if found[i] {
			if len(found) < g.max {
				continue
			} else {
				return fmt.Errorf("nao foi possivel gerar o jogo")
			}
		}
		found[i] = true
		if !g.generateConstraints(i) {
			continue
		}
		g.Numbers = append(g.Numbers, i)
	}
	sort.Slice(g.Numbers, func(i, j int) bool {
		return g.Numbers[i] < g.Numbers[j]
	})
	return nil
}

func (g *Game) generateConstraints(check int) bool {
	e := 0
	o := 0
	s := 1
	seq := 0
	seqArr := make([]int, 0, len(g.Numbers))

	if len(g.Numbers) == 0 {
		if check%2 == 0 {
			if g.MaxEven > 0 {
				return true
			}
		} else {
			if g.MaxOdd > 0 {
				return true
			}
		}
		return false
	}

	// check odds/evens
	for _, v := range g.Numbers {
		seqArr = append(seqArr, v)
		if v%2 == 0 {
			e++
		} else {
			o++
		}
	}
	if check%2 == 0 && e >= g.MaxEven {
		fmt.Printf("falha 1, e=%d, %v\n", check, seqArr)
		return false
	}
	if check%2 != 0 && o >= g.MaxOdd {
		fmt.Printf("falha 2, e=%d, %v\n", check, seqArr)
		return false
	}

	// check sequences
	seqArr = append(seqArr, check)
	sort.Slice(seqArr, func(i, j int) bool {
		return seqArr[i] < seqArr[j]
	})
	for i := range seqArr {
		if i != 0 && seqArr[i] == seqArr[i-1]+1 {
			s++
		} else {
			s = 1
		}

		if s > seq {
			seq = s
		}
	}
	if seq > g.MaxSeq {
		fmt.Printf("falha 3, e=%d, %v\n", check, seqArr)
		return false
	}
	return true
}
