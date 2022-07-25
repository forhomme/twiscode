package app

import "math"

func BinToDec(in int) (out int, err error) {
	decimal := 0
	counter := float64(0)
	remain := 0

	if in == 0 {
		return 0, nil
	}

	for in != 0 {
		remain = in % 10
		decimal += remain * int(math.Pow(2.0, counter))
		in = in / 10
		counter++
	}
	return decimal, nil
}

func DecToBin(in int) (out int, err error) {
	binary := 0
	counter := 1
	remain := 0

	if in == 0 {
		return 0, nil
	}

	for in != 0 {
		remain = in % 2
		in = in / 2
		binary += remain * counter
		counter *= 10
	}
	return binary, nil
}
