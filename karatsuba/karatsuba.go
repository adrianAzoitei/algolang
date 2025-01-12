package karatsuba

import (
	"algorithms/utils"
	"strconv"
)

func Karatsuba(x, y int) (xy int, err error) {
	// TODO: check if x,y are the same length
	x_str := strconv.Itoa(x)
	y_str := strconv.Itoa(y)
	n := len(x_str)

	if n == 1 {
		return x * y, nil
	}

	a, err := strconv.Atoi(x_str[:(n / 2)])
	if err != nil {
		return 0, err
	}
	b, err := strconv.Atoi(x_str[(n / 2):])
	if err != nil {
		return 0, err
	}
	c, err := strconv.Atoi(y_str[:(n / 2)])
	if err != nil {
		return 0, err
	}
	d, err := strconv.Atoi(y_str[(n / 2):])
	if err != nil {
		return 0, err
	}
	ac, err := Karatsuba(a, c)
	ad, err := Karatsuba(a, d)
	bc, err := Karatsuba(b, c)
	bd, err := Karatsuba(b, d)
	xy = utils.IntPow(10, n)*ac + utils.IntPow(10, n/2)*(ad+bc) + bd
	return
}
