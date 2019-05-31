package main

import "testing"

func TestParse(t *testing.T) {
	tests := []struct{
		input string
		expected Nanobot
	}{
		{"pos=<1,1,2>, r=1", Nanobot{Position{1,1,2},1}},
		{"pos=<99663890,15679983,37262396>, r=53694281", Nanobot{Position{99663890,15679983,37262396},53694281}},
		{"pos=<79593290,-23107000,60990308>, r=96138745", Nanobot{Position{79593290,-23107000,60990308},96138745}},
	}
	for _, test := range tests {
		result := Parse(test.input)
		if result != test.expected {
			t.Errorf("Parse(%s) expected %v, got %v", test.input, test.expected, result)
		}
	}
}

func TestPositionDistanceTo(t *testing.T) {
	source := Position{0, 0,0}
	tests := []struct {
		other Position
		distance int64
	}{
		{Position{0,0,0}, 0},
		{Position{1,0,0}, 1},
		{Position{4,0,0}, 4},
		{Position{0,2,0}, 2},
		{Position{0,5,0}, 5},
		{Position{0,0,3}, 3},
		{Position{1,1,1}, 3},
		{Position{1,1,2}, 4},
		{Position{1,3,1}, 5},
	}
	for _, test := range tests {
		result := source.distanceTo(test.other)
		if result != test.distance {
			t.Errorf("expected distance(%v,%v)=%d, was %d", source, test.other, test.distance, result)
		}
	}
}

func TestInRangeOfStrongest(t *testing.T) {
	result := inRangeOfStrongest([]Nanobot{
		{Position{0, 0, 0}, 4},
		{Position{1, 0, 0}, 1},
		{Position{4, 0, 0}, 3},
		{Position{0, 2, 0}, 1},
		{Position{0, 5, 0}, 3},
		{Position{0, 0, 3}, 1},
		{Position{1, 1, 1}, 1},
		{Position{1, 1, 2}, 1},
		{Position{1, 3, 1}, 1},
	})
	if result != 7 {
		t.Errorf("inRangeOfStrongest expected 7, was %d", result)
	}
}
