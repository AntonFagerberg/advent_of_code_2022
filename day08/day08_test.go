package day08

import (
	"log"
	"os"
	"testing"
)

func TestPart1Example(t *testing.T) {
	input, err := os.ReadFile("example")

	if err != nil {
		log.Fatal(err)
	}

	if res := part1(input); res != 21 {
		t.Errorf("Expected %v to be 21", res)
	}
}

func TestPart1(t *testing.T) {
	input, err := os.ReadFile("input")

	if err != nil {
		log.Fatal(err)
	}

	if res := part1(input); res != 1825 {
		t.Errorf("Expected %v to be 1825", res)
	}
}

func TestPart2Example(t *testing.T) {
	input, err := os.ReadFile("example")

	if err != nil {
		log.Fatal(err)
	}

	if res := part2(input); res != 8 {
		t.Errorf("Expected %v to be 8", res)
	}
}

func TestPart2(t *testing.T) {
	input, err := os.ReadFile("input")

	if err != nil {
		log.Fatal(err)
	}

	if res := part2(input); res != 235200 {
		t.Errorf("Expected %v to be 235200", res)
	}
}
