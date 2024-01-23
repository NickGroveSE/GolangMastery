package main

import (
	"bytes"
)

func validate(ccn string) bool {

	weights := bytes.Repeat([]byte{2, 1}, 8)
	summation := 0

	for pos, num := range ccn {

		digitCalculation := int(num-'0') * int(weights[pos])

		if digitCalculation >= 10 {
			digitCalculation = digitCalculation - 9
		}

		summation = summation + digitCalculation

	}

	return summation%10 == 0

}
