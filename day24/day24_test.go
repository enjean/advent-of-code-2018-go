package main

import (
	"reflect"
	"testing"
)

func TestTargetSelection(t *testing.T) {
	immune1 := &Group{
		numUnits: 17, hitPoints: 5390,
		attackDamage: 4507, attackType: "fire", initiative: 2,
		immuneTo: nil, weakTo: []string{"radiation", "bludgeoning"},
	}
	immune2 := &Group{
		numUnits: 989, hitPoints: 1274,
		attackDamage: 25, attackType: "slashing", initiative: 3,
		immuneTo: []string{"fire"}, weakTo: []string{"bludgeoning", "slashing"},
	}
	immuneSystem := Army{immune1, immune2}

	immuneTargetSelectionOrder := immuneSystem.targetSelectionOrder()
	if immuneTargetSelectionOrder[0] != immune1 && immuneTargetSelectionOrder[1] != immune2 {
		t.Errorf("Target selection order for immune incorrect, got %v", immuneTargetSelectionOrder)
	}

	infection1 := &Group{
		numUnits: 801, hitPoints: 4706,
		attackDamage: 116, attackType: "bludgeoning", initiative: 1,
		immuneTo: nil, weakTo: []string{"radiation"},
	}
	infection2 := &Group{
		numUnits: 4485, hitPoints: 2961,
		attackDamage: 12, attackType: "slashing", initiative: 4,
		immuneTo: []string{"radiation"}, weakTo: []string{"fire"},
	}
	infection := Army{infection1, infection2}

	infectionTargetSelectionOrder := infection.targetSelectionOrder()
	if infectionTargetSelectionOrder[0] != infection1 && infectionTargetSelectionOrder[1] != infection2 {
		t.Errorf("Target selection order for immune incorrect, got %v", infectionTargetSelectionOrder)
	}

	infectionTargetSelection := infection.selectTargets(immuneSystem)

	if infectionTargetSelection[infection1] != immune1 || infectionTargetSelection[infection2] != immune2 {
		t.Errorf("Target selection for infection was %v", infectionTargetSelection)
	}

	immuneTargetSelection := immuneSystem.selectTargets(infection)

	if immuneTargetSelection[immune1] != infection2 || immuneTargetSelection[immune2] != infection1 {
		t.Errorf("Target selection for immune was %v", infectionTargetSelection)
	}
}

func TestAttack(t *testing.T) {
	immune1 := &Group{
		numUnits: 17, hitPoints: 5390,
		attackDamage: 4507, attackType: "fire", initiative: 2,
		immuneTo: nil, weakTo: []string{"radiation", "bludgeoning"},
	}
	immune2 := &Group{
		numUnits: 989, hitPoints: 1274,
		attackDamage: 25, attackType: "slashing", initiative: 3,
		immuneTo: []string{"fire"}, weakTo: []string{"bludgeoning", "slashing"},
	}

	infection1 := &Group{
		numUnits: 801, hitPoints: 4706,
		attackDamage: 116, attackType: "bludgeoning", initiative: 1,
		immuneTo: nil, weakTo: []string{"radiation"},
	}
	infection2 := &Group{
		numUnits: 4485, hitPoints: 2961,
		attackDamage: 12, attackType: "slashing", initiative: 4,
		immuneTo: []string{"radiation"}, weakTo: []string{"fire"},
	}

	infection2.attack(immune2)
	if immune2.numUnits != 989 - 84 {
		t.Errorf("Immune2 did not take expected damage %v", immune2)
	}

	immune2.attack(infection1)
	if infection1.numUnits != 801 - 4 {
		t.Errorf("infection1 did not take expected damage %v", infection1)
	}

	immune1.attack(infection2)
	if infection2.numUnits != 4485 - 51 {
		t.Errorf("infection2 did not take expected damage %v", infection2)
	}

	infection1.attack(immune1)
	if immune1.numUnits != 0 {
		t.Errorf("immune1 did not take expected damage %v", immune1)
	}
}

func TestPerformRound(t *testing.T) {
	immune1 := &Group{
		numUnits: 17, hitPoints: 5390,
		attackDamage: 4507, attackType: "fire", initiative: 2,
		immuneTo: nil, weakTo: []string{"radiation", "bludgeoning"},
	}
	immune2 := &Group{
		numUnits: 989, hitPoints: 1274,
		attackDamage: 25, attackType: "slashing", initiative: 3,
		immuneTo: []string{"fire"}, weakTo: []string{"bludgeoning", "slashing"},
	}
	immuneSystem := Army{immune1, immune2}

	infection1 := &Group{
		numUnits: 801, hitPoints: 4706,
		attackDamage: 116, attackType: "bludgeoning", initiative: 1,
		immuneTo: nil, weakTo: []string{"radiation"},
	}
	infection2 := &Group{
		numUnits: 4485, hitPoints: 2961,
		attackDamage: 12, attackType: "slashing", initiative: 4,
		immuneTo: []string{"radiation"}, weakTo: []string{"fire"},
	}
	infection := Army{infection1, infection2}

	performRound(immuneSystem, infection)

	if immune1.numUnits != 0 {
		t.Errorf("Expected immune1 to have 0, was %v", immune1)
	}
	if immune2.numUnits != 905 {
		t.Errorf("Expected immune2 to have 905, was %v", immune2)
	}

	if infection1.numUnits != 797 {
		t.Errorf("Expected infection1 to have 797, was %v", infection1)
	}
	if infection2.numUnits != 4434 {
		t.Errorf("Expected infection2 to have 4434, was %v", infection2)
	}
}

func TestBattle(t *testing.T) {
	immune1 := &Group{
		numUnits: 17, hitPoints: 5390,
		attackDamage: 4507, attackType: "fire", initiative: 2,
		immuneTo: nil, weakTo: []string{"radiation", "bludgeoning"},
	}
	immune2 := &Group{
		numUnits: 989, hitPoints: 1274,
		attackDamage: 25, attackType: "slashing", initiative: 3,
		immuneTo: []string{"fire"}, weakTo: []string{"bludgeoning", "slashing"},
	}
	immuneSystem := Army{immune1, immune2}

	infection1 := &Group{
		numUnits: 801, hitPoints: 4706,
		attackDamage: 116, attackType: "bludgeoning", initiative: 1,
		immuneTo: nil, weakTo: []string{"radiation"},
	}
	infection2 := &Group{
		numUnits: 4485, hitPoints: 2961,
		attackDamage: 12, attackType: "slashing", initiative: 4,
		immuneTo: []string{"radiation"}, weakTo: []string{"fire"},
	}
	infection := Army{infection1, infection2}

	winner := battle(immuneSystem, infection)
	if winner.totalUnits() != 5216 {
		t.Errorf("Did not get winner in expected state, was %v", winner)
	}
}

func TestParseGroup(t *testing.T) {
	tests := []struct {
		input string
		expected Group
	}{
		{
			"17 units each with 5390 hit points (weak to radiation, bludgeoning) with an attack that does 4507 fire damage at initiative 2",
			Group{
				numUnits: 17, hitPoints: 5390,
				attackDamage: 4507, attackType: "fire", initiative: 2,
				immuneTo: nil, weakTo: []string{"radiation", "bludgeoning"},
			},
		},
		{
			"989 units each with 1274 hit points (immune to fire; weak to bludgeoning, slashing) with an attack that does 25 slashing damage at initiative 3",
			Group{
				numUnits: 989, hitPoints: 1274,
				attackDamage: 25, attackType: "slashing", initiative: 3,
				immuneTo: []string{"fire"}, weakTo: []string{"bludgeoning", "slashing"},
			},
		},
		{
			"137 units each with 3682 hit points with an attack that does 264 radiation damage at initiative 13",
			Group{
				numUnits: 137, hitPoints: 3682,
				attackDamage: 264, attackType: "radiation", initiative: 13,
				immuneTo: nil, weakTo: nil,
			},
		},
		{
			"1265 units each with 3299 hit points (weak to cold; immune to radiation) with an attack that does 25 fire damage at initiative 3",
			Group{
				numUnits: 1265, hitPoints: 3299,
				attackDamage: 25, attackType: "fire", initiative: 3,
				immuneTo: []string{"radiation"}, weakTo: []string{"cold"},
			},
		},
	}

	for _, test := range tests {
		result := ParseGroup(test.input)
		if !reflect.DeepEqual(result, test.expected) {
			t.Errorf("ParseGroup(%s) expected %v got %v", test.input, test.expected, result)
		}
	}
}
