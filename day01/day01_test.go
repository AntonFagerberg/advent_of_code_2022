package day01

import (
	"log"
	"os"
	"reflect"
	"testing"
)

func TestParse(t *testing.T) {
	input, err := os.ReadFile("example")

	if err != nil {
		log.Fatal(err)
	}

	res := Parse(input)
	exp := [][]int{
		{1000, 2000, 3000},
		{4000},
		{5000, 6000},
		{7000, 8000, 9000},
		{10000},
	}

	if !reflect.DeepEqual(res, exp) {
		t.Errorf("Expected %d to be %d", exp, res)
	}
}

func TestSums(t *testing.T) {
	input, err := os.ReadFile("example")

	if err != nil {
		log.Fatal(err)
	}

	res := Sums(Parse(input))
	exp := []int{4000, 6000, 10000, 11000, 24000}

	if !reflect.DeepEqual(res, exp) {
		t.Errorf("Expected %d to be %d", exp, res)

	}
}

func TestPart1Example(t *testing.T) {
	input, err := os.ReadFile("example")

	if err != nil {
		log.Fatal(err)
	}

	if res := Part1(input); res != 24000 {
		t.Errorf("Expected %d to be 24000", res)
	}
}

func TestPart1(t *testing.T) {
	input, err := os.ReadFile("input")

	if err != nil {
		log.Fatal(err)
	}

	if res := Part1(input); res != 70369 {
		t.Errorf("Expected %d to be 70369", res)
	}
}

func TestPart2Example(t *testing.T) {
	input, err := os.ReadFile("example")

	if err != nil {
		log.Fatal(err)
	}

	if res := Part2(input); res != 45000 {
		t.Errorf("Expected %d to be 45000", res)
	}
}

func TestPart2(t *testing.T) {
	input, err := os.ReadFile("input")

	if err != nil {
		log.Fatal(err)
	}

	if res := Part2(input); res != 203002 {
		t.Errorf("Expected %d to be 203002", res)
	}
}
