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

func TestParsePath(t *testing.T) {
	var tests = []struct {
		input string
		expected *Path
	}{
		{"^WNE$", &Path{"WNE", nil}},
		{"^ENWWW(NEEE|SSE(EE|N))$", &Path{"ENWWW", []*Path{
			{"NEEE", nil},
			{"SSE", []*Path{
				{"EE", nil},
				{"N", nil},
			}},
			}}},
	}

	for _, test := range tests {
		result := ParsePath(test.input)
		if !result.equalTo(test.expected) {
			t.Errorf("Expected parse(%v) = %v; Got %v", test.input, test.expected.print(), result.print())
		}
	}

}

func (p *Path) equalTo(other *Path) bool {
	if p.route != other.route {
		return false
	}
	if len(p.branches) != len(other.branches) {
		return false
	}
	for i, branch := range p.branches {
		if !branch.equalTo(other.branches[i]) {
			return false
		}
	}
	return true
}
