//
// EPITECH PROJECT, 2020
// CNA_groundhog_2019
// File description:
// groundhog by edouard touch and alexis lopes
//

package functions

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

type algo struct {
	tab        []float64
	tabEvo     []float64
	tabWeird   []float64
	nb         int
	increment  int
	lastsign   bool
	switchTime int
}

func (st *algo) SortDescending(weirdVal []float64) []float64 {
	for i := len(weirdVal); i > 0; i-- {
		for j := 1; j < i; j++ {
			if weirdVal[j-1] < weirdVal[j] {
				intermediate := st.tabWeird[j]
				intermediate2 := weirdVal[j]
				st.tabWeird[j] = st.tabWeird[j-1]
				weirdVal[j] = weirdVal[j-1]
				st.tabWeird[j-1] = intermediate
				weirdVal[j-1] = intermediate2
			}
		}
	}
	return weirdVal
}

func (st *algo) PrintWeirdestValue(index int) {
	res := 0

	if len(st.tabWeird)-2 > 4 {
		res = 5
	} else if len(st.tabWeird)-2 > 0 {
		res = len(st.tabWeird) - 2
	} else {
		res = 0
	}
	if len(st.tabWeird) > 2 {
		fmt.Printf("%d weirdest values are ", res)
		st.GetWeirdestValue(index)
	} else {
		fmt.Printf("no weirdest value\n")
	}
}

func (st *algo) GetWeirdestValue(index int) {
	weirdVal := make([]float64, 1)
	first := 0

	for third := 2; third != st.increment; third++ {
		predict := (st.tabWeird[first] + st.tabWeird[third]) / 2.
		weirdVal = append(weirdVal, math.Abs(predict-st.tabWeird[third-1.]))
		first++
	}
	st.tabWeird = append(st.tabWeird[:st.increment-1], st.tabWeird[st.increment-1+1:]...)
	st.tabWeird = append(st.tabWeird[:0], st.tabWeird[0+1:]...)
	weirdVal = append(weirdVal[:0], weirdVal[0+1:]...)
	weirdVal = st.SortDescending(weirdVal)
	fmt.Printf("[")
	if len(st.tabWeird) > 0 {
		fmt.Printf("%.1f", st.tabWeird[0])
		for i := 1; i != len(st.tabWeird) && i != 5; i++ {
			fmt.Printf(", %.1f", st.tabWeird[i])
		}
	}
	fmt.Printf("]\n")
}

func (st *algo) CalcDevation(index int) {
	resTab := float64(0)
	resX := float64(0)
	result := float64(0)
	x := make([]float64, 0)

	for i := 0; i < len(st.tab); i++ {
		resTab += st.tab[i]
	}
	resTab = resTab / float64(index)
	for i := 0; i < len(st.tab); i++ {
		x = append(x, math.Pow((st.tab[i]-resTab), 2))
	}
	for i := 0; i < len(x); i++ {
		resX += x[i]
	}
	result = math.Sqrt((1 / float64(index)) * resX)
	fmt.Printf("\ts=%.2f", result)
}

func (st *algo) CalcEvolution(index int) {
	res := float64(0)
	currentsign := bool(true)

	res = st.tabEvo[index] - st.tabEvo[0]
	res = res / math.Abs(st.tabEvo[0]) * 100
	fmt.Printf("\tr=%.0f%%", res)
	st.CalcDevation(index)
	if math.Inf(+1) != res && math.Inf(-1) != res {
		if res > 0 {
			currentsign = true
		}
		if res < 0 {
			currentsign = false
		}
		if math.Inf(+1) != res && math.Inf(-1) != res && st.lastsign != currentsign {
			fmt.Printf("\ta switch occurs")
			st.switchTime++
		}
		st.lastsign = currentsign
	}
}

func (st *algo) CalcTempInc(index int) {
	result := float64(0)

	for i := len(st.tabEvo) - 1; i != 0; i-- {
		if st.tabEvo[i] > st.tabEvo[i-1] {
			result += (st.tabEvo[i] - st.tabEvo[i-1])
		}
	}
	if result < 0 {
		result = 0
	}
	fmt.Printf("g=%.2f", result/float64(index))
}

func (st *algo) CreateTab(number float64, index int) {
	st.tab = append(st.tab, number)
	st.tabEvo = append(st.tabEvo, number)
	st.tabWeird = append(st.tabWeird, number)

	if st.nb > index {
		st.tab = append(st.tab[:0], st.tab[1:]...)
	}
	if st.nb > index+1 {
		st.tabEvo = append(st.tabEvo[:0], st.tabEvo[1:]...)
	}
	st.increment++
}

func (st *algo) Calcul(index int, number float64, verif bool) {
	if verif != true {
		st.CreateTab(number, index)
		if len(st.tab) < index {
			fmt.Printf("g=nan\tr=nan%%\ts=nan")
		} else if st.nb == index {
			fmt.Printf("g=nan\tr=nan%%")
			st.CalcDevation(index)
		} else {
			st.CalcTempInc(index)
			st.CalcEvolution(index)
		}
		fmt.Printf("\n")
		st.nb++
	}
}

//GroundHog read input
func GroundHog(index int) {
	st := algo{nb: 1, lastsign: true, switchTime: 0, increment: 0}
	reader := bufio.NewReader(os.Stdin)

	for {
		var verif bool
		str, err := reader.ReadString('\n')
		if err != nil {
			os.Exit(84)
		}
		str = strings.Replace(str, "\n", "", -1)
		if strings.Compare("STOP", str) == 0 {
			if st.nb <= index {
				os.Exit(84)
			}
			fmt.Printf("Global tendency switched %d times\n", st.switchTime)
			st.PrintWeirdestValue(index)
			break
		}
		number, err := strconv.ParseFloat(str, 64)
		if err != nil {
			os.Exit(84)
		}
		st.Calcul(index, number, verif)
	}
}
