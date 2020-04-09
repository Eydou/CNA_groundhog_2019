//
// EPITECH PROJECT, 2020
// 202unsold_2019
// File description:
// function
//

package functions

import (
	"errors"
)

//ErrorArgs check error
func ErrorArgs(args []string) (int, error) {
	if len(args) == 1 || len(args) > 3 {
		return 84, errors.New("invalid arguments")
	}
	for i := 1; i != len(args); i++ {
		for j := 0; j != len(args[i]); j++ {
			if args[i][j] < '0' || args[i][j] > '9' {
				return 84, errors.New("invalid number")
			}
		}
	}
	return 0, nil
}
