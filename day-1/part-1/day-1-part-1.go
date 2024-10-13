package part_1

import (
	"bufio"
	"os"
)

func Calibrate(filepath string) (sum int, err error) {
	f, err := os.Open(filepath)
	if err != nil {
		return 0, err
	}

	defer func() {
		if err := f.Close(); err != nil {
			panic(err)
		}
	}()

	b := bufio.NewScanner(f)

	for b.Scan() {
		var f, l int32 = -1, -1

		for _, v := range b.Text() {
			if v >= '0' && v <= '9' {
				if f == -1 {
					f = v - 48
				} else {
					l = v - 48
				}
			}
		}

		if f == -1 {
			continue
		}

		if l == -1 {
			l = f
		}

		sum += int(f*10 + l)
	}

	return sum, err
}
