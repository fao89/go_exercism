package raindrops

import "strconv"

// Convert number to Raindrops
func Convert(number int) string {
	mod3 := number%3 == 0
	mod5 := number%5 == 0
	mod7 := number%7 == 0

	if !mod3 && !mod5 && !mod7 {
		return strconv.Itoa(number)
	}

	raindrops := ""
	if mod3 {
		raindrops += "Pling"
	}

	if mod5 {
		raindrops += "Plang"
	}

	if mod7 {
		raindrops += "Plong"
	}

	return raindrops
}
