package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"sort"
	"strconv"
	"strings"
)

type Group struct {
	numUnits         int
	hitPoints        int
	attackDamage     int
	attackType       string
	initiative       int
	immuneTo, weakTo []string
}

func (g Group) effectivePower() int {
	return g.numUnits * g.attackDamage
}

func (g Group) modifyDamage(damage int, attackType string) int {
	for _, it := range g.immuneTo {
		if it == attackType {
			return 0
		}
	}
	for _, wt := range g.weakTo {
		if wt == attackType {
			return damage * 2
		}
	}
	return damage
}

func (g Group) damageTo(target Group) int {
	return target.modifyDamage(g.numUnits*g.attackDamage, g.attackType)
}

func (g *Group) attack(target *Group) {
	damage := g.damageTo(*target)
	numKilled := damage / target.hitPoints
	target.numUnits -= numKilled
	//fmt.Printf("Did %d damage, killing %d\n", damage, numKilled)
	if target.numUnits < 0 {
		target.numUnits = 0
	}
}

type Army []*Group

func (a Army) targetSelectionOrder() []*Group {
	var groups []*Group

	for _, group := range a {
		if group.numUnits > 0 {
			groups = append(groups, group)
		}
	}

	sort.Slice(groups, func(i, j int) bool {
		iPower := groups[i].effectivePower()
		jPower := groups[j].effectivePower()

		if iPower != jPower {
			return iPower > jPower
		}
		return groups[i].initiative > groups[j].initiative
	})

	return groups
}

func (a Army) selectTargets(opponent Army) map[*Group]*Group {
	selections := make(map[*Group]*Group)
	alreadyTargeted := make(map[*Group]bool)

	targetSelectionOrder := a.targetSelectionOrder()

	type targetStatsEntry struct {
		target         *Group
		damage         int
		effectivePower int
		initiative     int
	}

	for i := 0; i < len(targetSelectionOrder) && i < len(opponent); i++ {
		attacker := targetSelectionOrder[i]

		var targetStats []targetStatsEntry

		for _, oppGroup := range opponent {
			if oppGroup.numUnits == 0 {
				continue
			}
			if alreadyTargeted[oppGroup] {
				continue
			}
			wouldDamage := attacker.damageTo(*oppGroup)
			targetStats = append(targetStats, targetStatsEntry{
				target:         oppGroup,
				damage:         wouldDamage,
				effectivePower: oppGroup.effectivePower(),
				initiative:     oppGroup.initiative,
			})
		}

		sort.Slice(targetStats, func(i, j int) bool {
			if targetStats[i].damage != targetStats[j].damage {
				return targetStats[i].damage > targetStats[j].damage
			}
			if targetStats[i].effectivePower != targetStats[j].effectivePower {
				return targetStats[i].effectivePower > targetStats[j].effectivePower
			}
			return targetStats[i].initiative > targetStats[j].initiative
		})

		if len(targetStats) > 0 && targetStats[0].damage > 0 {
			target := targetStats[0].target
			selections[attacker] = target
			alreadyTargeted[target] = true
		}
	}

	return selections
}

func (a Army) totalUnits() int {
	units := 0
	for _, group := range a {
		units += group.numUnits
	}
	return units
}

type attackOrderEntry struct {
	group  *Group
	target *Group
}

func (a attackOrderEntry) String() string {
	return fmt.Sprintf("%-v attacking %-v", *a.group, *a.target)
}

func addToAttackOrder(allGroups *[]attackOrderEntry, attacking, defending Army) {
	targetSelections := attacking.selectTargets(defending)
	for _, a := range attacking {
		target := targetSelections[a]
		if target != nil {
			*allGroups = append(*allGroups, attackOrderEntry{group: a, target: target})
		}
	}
}

func performRound(army1, army2 Army) {
	var allGroups []attackOrderEntry
	addToAttackOrder(&allGroups, army1, army2)
	addToAttackOrder(&allGroups, army2, army1)

	//fmt.Printf("Attack order %-v \n", allGroups)
	sort.Slice(allGroups, func(i, j int) bool {
		return allGroups[i].group.initiative > allGroups[j].group.initiative
	})

	for _, attack := range allGroups {
		attack.group.attack(attack.target)
	}
}

func battle(army1, army2 Army) (Army, Army) {
	round := 0
	var l1, l2 int
	for army1.totalUnits() > 0 && army2.totalUnits() > 0 {
		performRound(army1, army2)
		fmt.Printf("After round %d, army 1 %d army 2 %d\n", round, army1.totalUnits(), army2.totalUnits())
		if army1.totalUnits() == l1 && army2.totalUnits() == l2 {
			// deadlock
			break
		}
		l1 = army1.totalUnits()
		l2 = army2.totalUnits()
		round++
	}

	return army1,army2
}

var groupExp = regexp.MustCompile(`(\d+) units each with (\d+) hit points (?:\((.*)\) )?with an attack that does (\d+) ([a-z]+) damage at initiative (\d+)`)
var attrExp = regexp.MustCompile(`(weak|immune) to (.*)`)

func ParseGroup(input string) Group {
	matches := groupExp.FindStringSubmatch(input)
	numUnits, _ := strconv.Atoi(matches[1])
	hitPoints, _ := strconv.Atoi(matches[2])
	attackDamage, _ := strconv.Atoi(matches[4])
	initiative, _ := strconv.Atoi(matches[6])

	var weakTo, immuneTo []string
	if matches[3] != "" {
		for _, part := range strings.Split(matches[3], ";") {
			attrMatches := attrExp.FindStringSubmatch(part)
			if attrMatches[1] == "weak" {
				weakTo = append(weakTo, strings.Split(attrMatches[2], ", ")...)
			} else if attrMatches[1] == "immune" {
				immuneTo = append(immuneTo, strings.Split(attrMatches[2], ", ")...)
			}
		}
	}

	return Group{
		numUnits:     numUnits,
		hitPoints:    hitPoints,
		attackDamage: attackDamage,
		attackType:   matches[5],
		initiative:   initiative,
		weakTo:       weakTo,
		immuneTo:     immuneTo,
	}
}

func buildArmy(spec []Group, boost int) Army {
	var groups []*Group
	for _, specGroup := range spec {
		group := specGroup
		group.attackDamage += boost
		groups = append(groups, &group)
	}
	return Army(groups)
}

func findMinBoost(immuneSystemInit, infectionInit []Group) (int, Army) {
	boost := 1
	for {
		immuneResult, infectionResult := battle(buildArmy(immuneSystemInit, boost), buildArmy(infectionInit, 0))
		if immuneResult.totalUnits() > 0 && infectionResult.totalUnits() == 0 {
			return boost, immuneResult
		}
		fmt.Printf("With boost %d, infection finished with %d\n", boost, infectionResult.totalUnits())
		boost++
	}
}

func main() {

	file, err := os.Open("day24/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var immuneSystemInit, infectionInit []Group
	var currentList *[]Group
	for scanner.Scan() {
		line := scanner.Text()
		if line == "Immune System:" {
			currentList = &immuneSystemInit
		} else if line == "Infection:" {
			currentList = &infectionInit
		} else if line != "" {
			group := ParseGroup(line)
			*currentList = append(*currentList, group)
		}
	}

	_, infectionResult := battle(buildArmy(immuneSystemInit, 0), buildArmy(infectionInit, 0))
	fmt.Printf("Part 1: %d", infectionResult.totalUnits())

	boost, immune := findMinBoost(immuneSystemInit, infectionInit)
	fmt.Printf("Part 2: With boost of %d immune has %d left\n", boost, immune.totalUnits())
}
