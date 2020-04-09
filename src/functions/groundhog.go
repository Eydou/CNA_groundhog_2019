//
// EPITECH PROJECT, 2020
// CNA_groundhog_2019
// File description:
// groundhog
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
	weirdVal   []float64
	nb         int
	lastsign  bool
	switchTime int
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
	res = res / st.tabEvo[0] * 100
	fmt.Printf("\tr=%.0f%%", res)
	st.CalcDevation(index)
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
	if st.nb > index {
		st.tab = append(st.tab[:0], st.tab[1:]...)
	}
	if st.nb > index+1 {
		st.tabEvo = append(st.tabEvo[:0], st.tabEvo[1:]...)
	}
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
	st := algo{nb: 1, lastsign: true, switchTime: 0}
	reader := bufio.NewReader(os.Stdin)

	for {
		var verif bool
		str, err := reader.ReadString('\n')
		if err != nil {
			os.Exit(84)
		}
		str = strings.Replace(str, "\n", "", -1)
		if strings.Compare("STOP", str) == 0 {
			fmt.Printf("Global tendency switched %d times\n", st.switchTime)
			fmt.Printf("%d weirdest values are ", st.switchTime)
			break
		}
		number, err := strconv.ParseFloat(str, 64)
		if err != nil {
			os.Exit(84)
		}
		st.Calcul(index, number, verif)
	}
}
