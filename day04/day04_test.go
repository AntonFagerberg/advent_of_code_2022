package day04

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

	if res := part1(input); res != 2 {
		t.Errorf("Expected %d to be 2", res)
	}
}

func TestPart1(t *testing.T) {
	input, err := os.ReadFile("input")

	if err != nil {
		log.Fatal(err)
	}

	if res := part1(input); res != 459 {
		t.Errorf("Expected %d to be 459", res)
	}
}

func TestPart2Example(t *testing.T) {
	input, err := os.ReadFile("example")

	if err != nil {
		log.Fatal(err)
	}

	if res := part2(input); res != 4 {
		t.Errorf("Expected %d to be 4", res)
	}
}

func TestPart2(t *testing.T) {
	input, err := os.ReadFile("input")

	if err != nil {
		log.Fatal(err)
	}

	if res := part2(input); res != 779 {
		t.Errorf("Expected %d to be 779", res)
	}
}
