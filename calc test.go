package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	var number1, operator, number2 string
	var numType int
	var calcResult int

	calc, _ := bufio.NewReader(os.Stdin).ReadString('\n')
	parts := strings.Fields(calc)

	if len(parts) != 3 {
		fmt.Println("Error: unsupported operation. Example: '4 + 4' or 'IV - IV'" +
			"\nUse spaces between objects")
		return
	}

	number1 = parts[0]
	operator = parts[1]
	number2 = parts[2]

	if isRoman(number1) && isRoman(number2) {
		numType = 1 // Оба числа римские
		num1 := romanToArabic(number1)
		num2 := romanToArabic(number2)
		calcResult = calculate(num1, operator, num2)
	} else {
		numType = 0
		num1, err1 := strconv.Atoi(number1)
		num2, err2 := strconv.Atoi(number2)
		if err1 != nil || err2 != nil || num1 > 10 || num2 > 10 || num1 <= 0 || num2 <= 0 {
			fmt.Println("Error: one of the two numbers is not supported" +
				"\nSupported symbols: 0 - 10, I - X (upper case only)" +
				"\nOnly Roman or only Arabic numerals" +
				"\nOnly integers")
			return
		}
		calcResult = calculate(num1, operator, num2)
	}

	if numType == 1 {
		RomeResult := arabicToRoman(calcResult)
		fmt.Println("Result:", RomeResult)
	} else {
		ArabResult := calcResult
		fmt.Println("Result:", ArabResult)
	}
}

func isRoman(s string) bool {
	romanNumerals := []string{"I", "II", "III", "IV", "V", "VI", "VII", "VIII", "IX", "X"}
	for _, numeral := range romanNumerals {
		if s == numeral {
			return true
		}
	}
	return false
}

func romanToArabic(roman string) int {
	romanNumerals := map[string]int{
		"I":    1,
		"II":   2,
		"III":  3,
		"IV":   4,
		"V":    5,
		"VI":   6,
		"VII":  7,
		"VIII": 8,
		"IX":   9,
		"X":    10,
	}
	return romanNumerals[roman]
}

func arabicToRoman(arabic int) string {
	if arabic <= 0 {
		fmt.Println("Error: negative Roman numeral")
		os.Exit(0)
	}
	arabicNumerals := map[int]string{
		1:   "I",
		2:   "II",
		3:   "III",
		4:   "IV",
		5:   "V",
		6:   "VI",
		7:   "VII",
		8:   "VIII",
		9:   "IX",
		10:  "X",
		11:  "XI",
		12:  "XII",
		13:  "XIII",
		14:  "XIV",
		15:  "XV",
		16:  "XVI",
		17:  "XVII",
		18:  "XVIII",
		19:  "XIX",
		20:  "XX",
		21:  "XXI",
		22:  "XXII",
		23:  "XXIII",
		24:  "XXIV",
		25:  "XXV",
		26:  "XXVI",
		27:  "XXVII",
		28:  "XXVIII",
		29:  "XXIX",
		30:  "XXX",
		31:  "XXXI",
		32:  "XXXII",
		33:  "XXXIII",
		34:  "XXXIV",
		35:  "XXXV",
		36:  "XXXVI",
		37:  "XXXVII",
		38:  "XXXVIII",
		39:  "XXXIX",
		40:  "XL",
		41:  "XLI",
		42:  "XLII",
		43:  "XLIII",
		44:  "XLIV",
		45:  "XLV",
		46:  "XLVI",
		47:  "XLVII",
		48:  "XLVIII",
		49:  "XLIX",
		50:  "L",
		51:  "LI",
		52:  "LII",
		53:  "LIII",
		54:  "LIV",
		55:  "LV",
		56:  "LVI",
		57:  "LVII",
		58:  "LVIII",
		59:  "LIX",
		60:  "LX",
		61:  "LXI",
		62:  "LXII",
		63:  "LXIII",
		64:  "LXIV",
		65:  "LXV",
		66:  "LXVI",
		67:  "LXVII",
		68:  "LXVIII",
		69:  "LXIX",
		70:  "LXX",
		71:  "LXXI",
		72:  "LXXII",
		73:  "LXXIII",
		74:  "LXXIV",
		75:  "LXXV",
		76:  "LXXVI",
		77:  "LXXVII",
		78:  "LXXVIII",
		79:  "LXXIX",
		80:  "LXXX",
		81:  "LXXXI",
		82:  "LXXXII",
		83:  "LXXXIII",
		84:  "LXXXIV",
		85:  "LXXXV",
		86:  "LXXXVI",
		87:  "LXXXVII",
		88:  "LXXXVIII",
		89:  "LXXXIX",
		90:  "XC",
		91:  "XCI",
		92:  "XCII",
		93:  "XCIII",
		94:  "XCIV",
		95:  "XCV",
		96:  "XCVI",
		97:  "XCVII",
		98:  "XCVIII",
		99:  "XCIX",
		100: "C",
	}
	return arabicNumerals[arabic]
}

func calculate(num1 int, operator string, num2 int) int {
	switch operator {
	case "+":
		return num1 + num2
	case "-":
		return num1 - num2
	case "*":
		return num1 * num2
	case "/":
		return num1 / num2
	default:
		fmt.Println("Error: not definite operator. Supported operators: ( + - / * )")
		os.Exit(1)
	}
	return 0
}
