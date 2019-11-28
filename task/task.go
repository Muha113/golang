package task

import (
	"fmt"
	"math"
	"regexp"
	"strconv"
)

// MyStrToInt1 : conversion string to int 1-t way
func MyStrToInt1(str string) (int, error) {
	resultInt, err := strconv.ParseFloat(str, 64)
	if err != nil {
		return 0, fmt.Errorf("Error occured: %s", err.Error())
	}
	offset := resultInt - float64(int(resultInt))
	if resultInt > float64(math.MaxInt64) || resultInt < float64(math.MinInt64) || offset != 0.0 {
		return 0, fmt.Errorf("Error occured: %s", "str do not match int template")
	}
	return int(resultInt), nil
}

// MyStrToInt2 : conversion string to int 2-d way
func MyStrToInt2(str string) (int, error) {
	regexMatch := regexp.MustCompile("^[+-]?(\\d+)(.[0]+|)$")
	matched := regexMatch.MatchString(str)
	if !matched {
		return 0, fmt.Errorf("Error occured: %s", "str do not match int template")
	}
	myStr := regexMatch.FindStringSubmatch(str)
	var coeff int
	if str[0] == '-' {
		coeff = -1
	} else {
		coeff = 1
	}
	resultInt, err := strconv.Atoi(myStr[1])
	if err != nil {
		return 0, fmt.Errorf("Error occured: %s", err.Error())
	}
	resultInt *= coeff
	return resultInt, nil
}
