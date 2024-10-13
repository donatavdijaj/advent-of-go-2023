package day_1

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
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
		f, l := -1, -1

		for _, v := range b.Text() {
			if v >= '0' && v <= '9' {
				if f == -1 {
					f = int(v) - 48
				} else {
					l = int(v) - 48
				}
			}
		}

		if f == -1 {
			continue
		}

		if l == -1 {
			l = f
		}

		n, err := strconv.Atoi(fmt.Sprintf("%d%d", f, l))
		if err != nil {
			return 0, err
		}

		sum += n
		fmt.Println(b.Text(), n, sum)
	}

	return sum, err
}
