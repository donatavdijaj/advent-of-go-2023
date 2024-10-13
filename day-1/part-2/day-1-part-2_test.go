package part_1

import (
	"testing"
)

func TestCalibrate(t *testing.T) {
	sum, err := Calibrate("../inputs/test-input-part-2.txt")

	if err != nil {
		t.Fatal(err)
	}

	if sum != 281 {
		t.Fatalf("Expected %d, recieved %d", 281, sum)
	}

}
