package day06

import (
	"log"
	"os"
	"testing"
)

func TestPart1(t *testing.T) {
	input, err := os.ReadFile("input")

	if err != nil {
		log.Fatal(err)
	}

	if res := solve(string(input), 4); res != 1760 {
		t.Errorf("Expected %v to be 1760", res)
	}
}

func TestPart2(t *testing.T) {
	input, err := os.ReadFile("input")

	if err != nil {
		log.Fatal(err)
	}

	if res := solve(string(input), 14); res != 2974 {
		t.Errorf("Expected %v to be 2974", res)
	}
}
