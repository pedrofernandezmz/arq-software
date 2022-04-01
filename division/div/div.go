package div

import (
	"errors"
)

func Division(a float64, b float64) (float64, error) {
	if b == 0 {
		return -1, errors.New("invalid zero divisor")
	}
	return a / b, nil
}
