package main

func IdentifyGameIsPossible(b *bool, color string, number int) {
	switch color {
	case "red":
		if number > 12 {
			*b = false
		}
	case "green":
		if number > 13 {
			*b = false
		}
	case "blue":
		if number > 14 {
			*b = false
		}
	}
}
