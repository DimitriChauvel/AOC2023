package main

type MinBag struct {
	red   int
	green int
	blue  int
}

func NewMinBag() MinBag {
	return MinBag{red: 1, green: 1, blue: 1}
}

func (b *MinBag) Power() int {
	return b.red * b.green * b.blue
}

func (b *MinBag) SetRed(number int) {
	if number > b.red {
		b.red = number
	}
}

func (b *MinBag) SetGreen(number int) {
	if number > b.green {
		b.green = number
	}
}

func (b *MinBag) SetBlue(number int) {
	if number > b.blue {
		b.blue = number
	}
}

func (b *MinBag) SetMinBag(color string, number int) {
	switch color {
	case "red":
		b.SetRed(number)
	case "green":
		b.SetGreen(number)
	case "blue":
		b.SetBlue(number)
	}
}
