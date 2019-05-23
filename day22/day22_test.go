package main

import "testing"

func TestErosionLevelTable(t *testing.T) {
	table := erosionLevelTable(10, 10, 510)

	var tests = []struct {
		x, y, geologicIndex int
	}{
		{0, 0, 510},
		{1, 0, 17317},
		{0, 1, 8415},
		{1, 1, 1805},
		{10, 10, 510},
	}
	for _, test := range tests {
		if table[test.x][test.y] != test.geologicIndex {
			t.Errorf("geologicIndex[%d][%d] = %d", test.x, test.y, test.geologicIndex)
		}
	}
}

func TestRiskLevel(t *testing.T) {
	if riskLevel(10, 10, 510) != 114 {
		t.Errorf("riskLevel with target 10,10 and depth 510 = 114")
	}
}
