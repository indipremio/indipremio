package main

import (
	"errors"
	"fmt"
	"strconv"
)

var romanPack = []struct {
	arab int
	rom  string
}{
	{1000, "M"}, {900, "CM"}, {500, "D"}, {400, "CD"},
	{100, "C"}, {90, "XC"}, {50, "L"}, {40, "XL"},
	{10, "X"}, {9, "IX"}, {5, "V"}, {4, "IV"}, {1, "I"},
}

func convertArabicToRoman(num int) string {
	if num == 0 {
		return ""
	}
	for _, pair := range romanPack {
		if num >= pair.arab {
			return pair.rom + convertArabicToRoman(num-pair.arab)
		}
	}
	return ""
}
func main() {
	var firstNumber, operator, secondNumber string

	for {

		_, err := fmt.Scanln(&firstNumber, &operator, &secondNumber)
		if err != nil {
			panic(err)
		}

		arabic1, err := strconv.Atoi(firstNumber)
		if err != nil {

		}
		arabic2, err := strconv.Atoi(secondNumber)
		if err != nil {
			roma1, err := RomanToInt(firstNumber)
			if err != nil {

			}

			for {
				roma2, err := RomanToInt(secondNumber)
				if err != nil {

				}

				if roma1 < 1 || roma1 > 10 {
					panic(err)
					return
				}

				if roma2 < 1 || roma2 > 10 {
					panic(err)
					return
				}
				switch operator {

				case "+":
					num := roma1 + roma2

					roman := convertArabicToRoman(num)

					fmt.Printf("%s\n", roman)

				case "-":
					num := roma1 - roma2
					if num <= 0 {
						panic(err)
					}
					roman := convertArabicToRoman(num)
					fmt.Printf("%s\n", roman)
				case "*":
					num := roma1 * roma2
					roman := convertArabicToRoman(num)
					fmt.Printf("%s\n", roman)
				case "/":
					num := roma1 / roma2
					if num <= 0 {
						panic(err)
					}
					roman := convertArabicToRoman(num)
					fmt.Printf("%s\n ", roman)

				}
				break
			}
			return
		}

		if arabic1 < 1 || arabic1 > 10 {

			panic(err)

		}

		if arabic2 < 1 || arabic2 > 10 {
			panic(err)

		}

		switch operator {

		case "+":
			fmt.Printf("%d\n", arabic1+arabic2)
		case "-":
			fmt.Printf("%d\n", arabic1-arabic2)
		case "*":

			fmt.Printf("%d\n", arabic1*arabic2)
		case "/":
			fmt.Printf("%d\n", arabic1/arabic2)

		default:
			panic(err)
		}

	}

}
func RomanToInt(string_ string) (int, error) {
	var sum int
	for i := len(string_) - 1; i >= 0; i-- {
		switch string_[i] {
		case 73:
			sum += 1

		case 86:
			if i > 0 {
				if string_[i-1] == 73 {
					sum += 4
					i--
				} else {
					sum += 5
				}
			} else {
				sum += 5
			}
		case 88:
			if i > 0 {
				if string_[i-1] == 73 {
					sum += 9
					i--
				} else {
					sum += 10
				}
			} else {
				sum += 10
			}

		case 76:
			if i > 0 {
				if string_[i-1] == 88 {
					sum += 40
					i--
				} else {
					sum += 50
				}
			} else {
				sum += 50
			}
		case 67:
			if i > 0 {
				if string_[i-1] == 88 {
					sum += 90
					i--
				} else {
					sum += 100
				}
			} else {
				sum += 100
			}
		case 68:
			if i > 0 {
				if string_[i-1] == 67 {
					sum += 400
					i--
				} else {
					sum += 500
				}
			} else {
				sum += 500
			}
		case 77:
			if i > 0 {
				if string_[i-1] == 67 {
					sum += 900
					i--
				} else {
					sum += 1000
				}
			} else {
				sum += 1000
			}
		default:
			return 0, errors.New("invalid input")
		}
	}
	return sum, nil
}
