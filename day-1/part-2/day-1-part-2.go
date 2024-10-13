package part_1

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var n = [18]string{"1", "one", "2", "two", "3", "three", "4", "four",
	"5", "five", "6", "six", "7", "seven", "8", "eight", "9", "nine"}

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
		t := b.Text()

		for i, v := range n {
			j := strings.Index(t, v)

			if j != -1 {

				x := -1

				if i%2 == 0 {
					x, err = strconv.Atoi(v)
				} else {
					x, err = strconv.Atoi(n[i-1])
				}

				if err != nil {
					return 0, err
				}

				if f == -1 {
					f = x
				} else {
					l = x
				}
			}
		}

		if f == -1 {
			continue
		}

		if l == -1 {
			l = f
		}

		fmt.Println(f, l)
	}

	return sum, err
}
