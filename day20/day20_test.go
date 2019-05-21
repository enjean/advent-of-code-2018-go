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
		result := FurthestRoom(test.input)
		if result != test.expected {
			t.Errorf("Expected FurthestRoom(%s) = %d; Got %v", test.input, test.expected, result)
		}
	}
}
