package part_1

import (
	"testing"
)

func TestCalibrate(t *testing.T) {
	sum, err := Calibrate("../inputs/test-input-part-1.txt")

	if err != nil {
		t.Fatal(err)
	}

	if sum != 142 {
		t.Fatalf("Expected %d, recieved %d", 142, sum)
	}
}
