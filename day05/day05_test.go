package day05

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

	if res := solve(input, 1); res != "CMZ" {
		t.Errorf("Expected %v to be CMZ", res)
	}
}

func TestPart1(t *testing.T) {
	input, err := os.ReadFile("input")

	if err != nil {
		log.Fatal(err)
	}

	if res := solve(input, 1); res != "RTGWZTHLD" {
		t.Errorf("Expected %v to be RTGWZTHLD", res)
	}
}

func TestPart2Example(t *testing.T) {
	input, err := os.ReadFile("example")

	if err != nil {
		log.Fatal(err)
	}

	if res := solve(input, 2); res != "MCD" {
		t.Errorf("Expected %v to be MCD", res)
	}
}

func TestPart2(t *testing.T) {
	input, err := os.ReadFile("input")

	if err != nil {
		log.Fatal(err)
	}

	if res := solve(input, 2); res != "STHGRZZFR" {
		t.Errorf("Expected %v to be STHGRZZFR", res)
	}
}
