package main

import "testing"

func TestErosionLevelTable(t *testing.T) {
	table := erosionLevelTable(10, 10, 10, 10, 510)

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

func TestRiskLevelTable(t *testing.T) {
	table := riskLevelTable(10, 10, 10, 10, 510)

	var tests = []struct {
		x, y int
		regionType RegionType
	}{
		{0, 0, Rocky},
		{1, 0, Wet},
		{0, 1, Rocky},
		{1, 1, Narrow},
		{10, 10, Rocky},
	}
	for _, test := range tests {
		if RegionType(table[test.x][test.y]) != test.regionType {
			t.Errorf("riskLevel[%d][%d] = %d", test.x, test.y, test.regionType)
		}
	}
}

func TestRiskLevel(t *testing.T) {
	if riskLevel(10, 10, 510) != 114 {
		t.Errorf("riskLevel with target 10,10 and depth 510 = 114")
	}
}

func TestShortestTimeTo(t *testing.T) {
	if shortestTimeTo(10, 10, 510) != 45 {
		t.Errorf("shortest trip to target 10,10 and depth 510 = 45")
	}
}
