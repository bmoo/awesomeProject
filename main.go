package awesomeProject

import "errors"

var intToSingle = []string{
	"zero", "one", "two", "three", "four", "five", "six", "seven", "eight", "nine", "ten",
	"eleven", "twelve", "thirteen", "fourteen", "fifteen", "sixteen", "seventeen", "eighteen", "nineteen",
}

func format(input int) (string, error) {
	if input >= 1000 {
		hundreds, err := formatHundreds(input % 1000)
		if err != nil {
			return "", err
		}
		return intToSingle[input / 1000] + " thousand " + hundreds, nil
	}

	return formatHundreds(input)
}

func formatHundreds(input int) (string, error) {

	if input < 0 {
		return "", errors.New("got a number that I don't recognize")
	}

	if input < 20 {
		return intToSingle[input], nil
	}

	intToTens := []string{
		"twenty", "thirty", "forty",
	}

	remainder := input % 10
	if remainder == 0 {
		return intToTens[(input/10)-2], nil
	}

	if input/100 > 0 {
		hundredsRemainder := input % 100

		return intToSingle[input/100] + " hundred " + intToTens[(hundredsRemainder/10)-2] + " " + intToSingle[remainder], nil
	}

	return intToTens[(input/10)-2] + " " + intToSingle[remainder], nil
}
