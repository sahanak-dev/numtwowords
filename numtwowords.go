package numtwowords

import "fmt"

//MaxNum is the largest number that can be converted to words
const MaxNum = 999
const MinNum = -999

//Convert converts a number to its string representation
func Convert(num int) (string, error) {
	if num < MinNum || num > MaxNum {
		return "", fmt.Errorf("number out of range: %d of %d to %d", num, MinNum, MaxNum)
	}
	if num == 0 {
		return "Not a valid number", nil
	}
	units := [20]string{"zero", "one", "two", "three", "four", "five", "six", "seven", "eight", "nine", "ten", "eleven", "twelve", "thirteen", "fourteen", "fifteen", "sixteen", "seventeen", "eighteen", "nineteen"}
	tenns := [8]string{"twenty", "thirty", "forty", "fifty", "sixty", "seventy", "eighty", "ninety"}
	result := ""

	if num < 0 {
		result += "minus "
		num = num * -1
	}
	if num > 99 {
		hundredindex := num / 100
		result += units[hundredindex] + " hundred"
		num = num % 100
		if num == 0 {
			return result, nil
		}
		result += " and "
	}
	if num > 19 {
		tenindex := num/10 - 2
		result += tenns[tenindex]
		num = num % 10
		if num == 0 {
			return result, nil
		}
		result += " "
	}

	if num > 0 {
		result += units[num]
	}
	return result, nil
}
