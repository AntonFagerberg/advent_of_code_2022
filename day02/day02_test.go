package day02

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

	res := parse(input)

	exp := [][]string{
		{"A", "Y"},
		{"B", "X"},
		{"C", "Z"},
	}

	if !reflect.DeepEqual(res, exp) {
		t.Fail()
	}
}

func TestPart1Example(t *testing.T) {
	input, err := os.ReadFile("example")

	if err != nil {
		log.Fatal(err)
	}

	if res := part1(parse(input)); res != 15 {
		t.Errorf("Expected %d to be 15", res)
	}
}

func TestPart1(t *testing.T) {
	input, err := os.ReadFile("input")

	if err != nil {
		log.Fatal(err)
	}

	if res := part1(parse(input)); res != 12740 {
		t.Errorf("Expected %d to be 12740", res)
	}
}

func TestPart2Example(t *testing.T) {
	input, err := os.ReadFile("example")

	if err != nil {
		log.Fatal(err)
	}

	if res := part2(parse(input)); res != 12 {
		t.Errorf("Expected %d to be 12", res)
	}
}

func TestPart2(t *testing.T) {
	input, err := os.ReadFile("input")

	if err != nil {
		log.Fatal(err)
	}

	if res := part2(parse(input)); res != 11980 {
		t.Errorf("Expected %d to be 11980", res)
	}
}
