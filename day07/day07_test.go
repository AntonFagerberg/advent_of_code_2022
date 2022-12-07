package day07

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

	if res := part1(input); res != 95437 {
		t.Errorf("Expected %v to be 95437", res)
	}
}

func TestPart1(t *testing.T) {
	input, err := os.ReadFile("input")

	if err != nil {
		log.Fatal(err)
	}

	if res := part1(input); res != 2104783 {
		t.Errorf("Expected %v to be 2104783", res)
	}
}

func TestPart2Example(t *testing.T) {
	input, err := os.ReadFile("example")

	if err != nil {
		log.Fatal(err)
	}

	if res := part2(input); res != 24933642 {
		t.Errorf("Expected %v to be 24933642", res)
	}
}

func TestPart2(t *testing.T) {
	input, err := os.ReadFile("input")

	if err != nil {
		log.Fatal(err)
	}

	if res := part2(input); res != 5883165 {
		t.Errorf("Expected %v to be 5883165", res)
	}
}
