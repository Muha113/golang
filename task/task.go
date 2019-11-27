package task

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

// Max values depending on system bit depth
var maxUint = ^uint(0)
var minUint = 0
var maxInt = int(maxUint >> 1)
var minInt = -maxInt - 1

// MyStrToInt1 : conversion string to int 1-t way
func MyStrToInt1(str string) (int, error) {
	str = strings.TrimSpace(str)
	resultInt, err := strconv.ParseFloat(str, 64)
	if err != nil {
		return 0, fmt.Errorf("Error occured: %s", err.Error())
	}
	offset := resultInt - float64(int(resultInt))
	if resultInt > float64(maxInt) || resultInt < float64(minInt) || offset != 0.0 {
		return 0, fmt.Errorf("Error occured: %s", "str do not match int template")
	}
	return int(resultInt), nil
}

// MyStrToInt2 : conversion string to int 2-d way
func MyStrToInt2(str string) (int, error) {
	regexMatch := "^\\s*[+-]{0,1}[0-9]+(.[0]+|)\\s*$"
	regexSearch := regexp.MustCompile("[1-9]+[0-9]*")
	matched, _ := regexp.MatchString(regexMatch, str)
	if !matched {
		return 0, fmt.Errorf("Error occured: %s", "str do not match int template")
	}
	myStr := regexSearch.FindString(str)
	var coeff int
	if str[0] == '-' {
		coeff = -1
	} else {
		coeff = 1
	}
	resultInt, err := strconv.Atoi(myStr)
	if err != nil {
		return 0, fmt.Errorf("Error occured: %s", err.Error())
	}
	resultInt *= coeff
	return resultInt, nil
}
