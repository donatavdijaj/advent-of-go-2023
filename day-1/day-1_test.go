package day_1

import "testing"

func TestCalibrateTrebuchet(t *testing.T) {
	sum, err := Calibrate("test-input.txt")

	if err != nil {
		t.Fatal(err)
	}

	if sum != 142 {
		t.Fatalf("Expected %d, recieved %d", 142, sum)
	}

}
