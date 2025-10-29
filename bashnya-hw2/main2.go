package main

import (
	"fmt"
	"strconv"
	"strings"
)

func ones_ones(digit rune, alf []rune) string {
	var builder string
	if digit == alf[1] {
		builder = "один "
	} else if digit == alf[2] {
		builder = "два "
	} else if digit == alf[3] {
		builder = "три "
	} else if digit == alf[4] {
		builder = "четыре "
	} else if digit == alf[5] {
		builder = "пять "
	} else if digit == alf[6] {
		builder = "шесть "
	} else if digit == alf[7] {
		builder = "семь "
	} else if digit == alf[8] {
		builder = "восемь "
	} else if digit == alf[9] {
		builder = "девять "
	}
	return builder
}

func ones_thous(digit rune, alf []rune) string {
	var builder string
	if digit == alf[1] {
		builder = "одна "
	} else if digit == alf[2] {
		builder = "две "
	}
	return builder
}

func doz(digit rune, alf []rune) string {
	var builder string
	if digit == alf[2] {
		builder = "двадцать "
	} else if digit == alf[3] {
		builder = "тридцать "
	} else if digit == alf[4] {
		builder = "сорок "
	} else if digit == alf[5] {
		builder = "пятьдесят "
	} else if digit == alf[6] {
		builder = "шестьдесят "
	} else if digit == alf[7] {
		builder = "семьдесят "
	} else if digit == alf[8] {
		builder = "восемьдесят "
	} else if digit == alf[9] {
		builder = "девяносто "
	}
	return builder
}

func tens(digit rune, alf []rune) string {
	var builder string
	if digit == alf[1] {
		builder = "одиннадцать "
	} else if digit == alf[2] {
		builder = "двенадцать "
	} else if digit == alf[3] {
		builder = "тринадцать "
	} else if digit == alf[4] {
		builder = "четырнадцать "
	} else if digit == alf[5] {
		builder = "пятнадцать "
	} else if digit == alf[6] {
		builder = "шестнадцать "
	} else if digit == alf[7] {
		builder = "семнадцать "
	} else if digit == alf[8] {
		builder = "восемнадцать "
	} else if digit == alf[9] {
		builder = "девятнадцать "
	} else {
		builder = "десять "
	}
	return builder
}

func hundreds(digit rune, alf []rune) string {
	var builder string
	if digit == alf[1] {
		builder = "сто "
	} else if digit == alf[2] {
		builder = "двести "
	} else if digit == alf[3] {
		builder = "триста "
	} else if digit == alf[4] {
		builder = "четыреста "
	} else if digit == alf[5] {
		builder = "пятьсот "
	} else if digit == alf[6] {
		builder = "шестьсот "
	} else if digit == alf[7] {
		builder = "семьсот "
	} else if digit == alf[8] {
		builder = "восемьсот "
	} else if digit == alf[9] {
		builder = "девятьсот "
	}
	return builder
}

func thousands(digit1, digit2 rune, alf []rune) string {
	var builder string
	if digit1 == alf[1] {
		builder = "тысяч "
	} else if digit2 == alf[1] {
		builder = "тысяча "
	} else if digit2 == alf[2] || digit2 == alf[3] || digit2 == alf[4] {
		builder = "тысячи "
	} else {
		builder = "тысяч "
	}

	return builder
}

func transcryption(num int) string {
	alf := []rune("0123456789")
	numstring := []rune(strconv.Itoa(num))
	ans := strings.Builder{}
	for i := 0; i < 2; i++ {
		if (i == 0 && len(numstring) == 6) || i == 1 {
			ans.WriteString(hundreds(numstring[len(numstring)-(6-3*i)], alf))
		}

		if numstring[len(numstring)-(5-3*i)] == alf[1] {
			ans.WriteString(tens(numstring[len(numstring)-(4-3*i)], alf))
		} else {
			ans.WriteString(doz(numstring[len(numstring)-(5-3*i)], alf))
			if i == 0 && (numstring[len(numstring)-4] == alf[1] || numstring[len(numstring)-4] == alf[2]) {
				ans.WriteString(ones_thous(numstring[len(numstring)-4], alf))
			} else {
				ans.WriteString(ones_ones(numstring[len(numstring)-(4-3*i)], alf))
			}
		}

		if i == 0 {
			ans.WriteString(thousands(numstring[len(numstring)-5], numstring[len(numstring)-4], alf))
		}
	}

	return ans.String()
}

func main() {
	var num int
	fmt.Println("Введите число: ")
	fmt.Scanf("%v", &num)
	if num >= 12307 || num <= -12307 {
		fmt.Println("Число слишком большое")
		return
	}

	for num < 12307 {
		if num < 0 {
			num *= -1
		} else if num%7 == 0 {
			num *= 39
		} else if num%9 == 0 {
			num = num*13 + 1
		} else {
			num = (num + 2) * 3
		}

		if num%9 == 0 && num%13 == 0 {
			fmt.Println("service error")
			return
		} else {
			num += 1
		}
	}

	fmt.Println("В результате получилось число: ", transcryption(num))
}
