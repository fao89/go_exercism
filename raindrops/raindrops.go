package raindrops

import "strconv"

// Convert number to Raindrops
func Convert(number int) string {
	raindrops := ""

	if number%3 == 0 {
		raindrops += "Pling"
	}

	if number%5 == 0 {
		raindrops += "Plang"
	}

	if number%7 == 0 {
		raindrops += "Plong"
	}

	if len(raindrops) == 0 {
		return strconv.Itoa(number)
	}

	return raindrops
}
