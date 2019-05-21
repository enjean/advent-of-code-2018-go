package main

import "testing"

func TestFurthestRoom(t *testing.T) {
	var tests = []struct{
		input string
		expected int
	}{
		{"^WNE$", 3},
		{"^ENWWW(NEEE|SSE(EE|N))$", 10},
		{"^ENNWSWW(NEWS|)SSSEEN(WNSE|)EE(SWEN|)NNN$", 18},
		{"^ESSWWN(E|NNENN(EESS(WNSE|)SSS|WWWSSSSE(SW|NNNE)))$", 23},
		{"^WSSEESWWWNW(S|NENNEEEENN(ESSSSW(NWSW|SSEN)|WSWWN(E|WWS(E|SS))))$", 31},
	}
	for _, test := range tests {
		result := distancesToRooms(test.input).max()
		if result != test.expected {
			t.Errorf("Expected Distances(%s).Max() = %d; Got %v", test.input, test.expected, result)
		}
	}
}

func TestAtLeastNAway(t *testing.T) {
	var tests = []struct{
		input string
		n int
		expected int
	}{
		{"^WNE$", 2, 2},
		{"^ENWWW(NEEE|SSE(EE|N))$", 5, 11},
		{"^ENNWSWW(NEWS|)SSSEEN(WNSE|)EE(SWEN|)NNN$", 10, 13},
	}
	for _, test := range tests {
		result := distancesToRooms(test.input).atLeastNAway(test.n)
		if result != test.expected {
			t.Errorf("Expected Distances(%s).AtLeastNAway(%d) = %d; Got %v", test.input, test.n, test.expected, result)
		}
	}
}
