package day03

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

	if res := part1(input); res != 157 {
		t.Errorf("Expected %d to be 157", res)
	}
}

func TestPart2Example(t *testing.T) {
	input, err := os.ReadFile("example")

	if err != nil {
		log.Fatal(err)
	}

	if res := part2(input); res != 70 {
		t.Errorf("Expected %d to be 70", res)
	}
}

func TestPart1(t *testing.T) {
	input, err := os.ReadFile("input")

	if err != nil {
		log.Fatal(err)
	}

	if res := part1(input); res != 8139 {
		t.Errorf("Expected %d to be 8139", res)
	}
}

func TestPart2(t *testing.T) {
	input, err := os.ReadFile("input")

	if err != nil {
		log.Fatal(err)
	}

	if res := part2(input); res != 2668 {
		t.Errorf("Expected %d to be 2668", res)
	}
}
