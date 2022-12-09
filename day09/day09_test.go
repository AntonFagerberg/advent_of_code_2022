package day09

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

	if res := solve(input, 2); res != 13 {
		t.Errorf("Expected %v to be 13", res)
	}
}

func TestPart1(t *testing.T) {
	input, err := os.ReadFile("input")

	if err != nil {
		log.Fatal(err)
	}

	if res := solve(input, 2); res != 6175 {
		t.Errorf("Expected %v to be 6175", res)
	}
}

func TestExample2(t *testing.T) {
	input, err := os.ReadFile("example2")

	if err != nil {
		log.Fatal(err)
	}

	if res := solve(input, 10); res != 36 {
		t.Errorf("Expected %v to be 36", res)
	}
}

func TestPart2(t *testing.T) {
	input, err := os.ReadFile("input")

	if err != nil {
		log.Fatal(err)
	}

	if res := solve(input, 10); res != 2578 {
		t.Errorf("Expected %v to be 2578", res)
	}
}
