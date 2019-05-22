package main

import "fmt"

func main() {
	vals := make(map[int]bool)

	r := make([]int, 6)
	var last int
	//I0:
	//fmt.Printf("I0 %v\n", r)
	r[4] = 123
I1:
	//fmt.Printf("I1 %v\n", r)
	r[4] = r[4] & 456
	// I2:
	//fmt.Printf("I2 %v\n", r)
	if r[4] == 72 {
		r[4] = 1
	} else {
		r[4] = 0
	}
	// I3:
	//fmt.Printf("I3 %v\n", r)
	//r[1] = r[4] + r[1]
	if r[4] == 0 {
		goto I4
	} else {
		goto I5
	}
I4:
	//fmt.Printf("I4 %v\n", r)
	//r[1] = 0
	goto I1
I5:
	//fmt.Printf("I5 %v\n", r)
	r[4] = 0
I6:
	//fmt.Printf("I6 %v\n", r)
	r[3] = r[4] | 65536
	// I7:
	//fmt.Printf("I7 %v\n", r)
	r[4] = 16098955
I8:
	//fmt.Printf("I8 %v\n", r)
	r[5] = r[3] & 255
	// I9:
	//fmt.Printf("I9 %v\n", r)
	r[4] = r[4] + r[5]
	// I10:
	//fmt.Printf("I10 %v\n", r)
	r[4] = r[4] & 16777215
	// I11:
	//fmt.Printf("I11 %v\n", r)
	r[4] = r[4] * 65899
	// I12:
	//fmt.Printf("I12 %v\n", r)
	r[4] = r[4] & 16777215
	// I13:
	//fmt.Printf("I13 %v\n", r)
	if 256 > r[3] {
		r[5] = 1
	} else {
		r[5] = 0
	}
	// I14:
	//fmt.Printf("I14 %v\n", r)
	//r[1] = r[5] + r[1]
	if r[5] == 0 {
		goto I15
	} else {
		goto I16
	}
I15:
	//fmt.Printf("I15 %v\n", r)
	//r[1] = r[1] + 1
	goto I17
I16:
	//fmt.Printf("I16 %v\n", r)
	//r[1] = 27
	goto I28
I17:
	//fmt.Printf("I17 %v\n", r)
	r[5] = 0
I18:
	//fmt.Printf("I18 %v\n", r)
	r[2] = r[5] + 1
	// I19:
	//fmt.Printf("I19 %v\n", r)
	r[2] = r[2] * 256
	// I20:
	//fmt.Printf("I20 %v\n", r)
	if r[2] > r[3] {
		r[2] = 1
	} else {
		r[2] = 0
	}
	// I21:
	//fmt.Printf("I21 %v\n", r)
	//r[1] = r[2] + r[1]
	if r[2] == 0 {
		goto I22
	} else {
		goto I23
	}
I22:
	//fmt.Printf("I22 %v\n", r)
	//r[1] = r[1] + 1
	goto I24
I23:
	//fmt.Printf("I23 %v\n", r)
	//r[1] = 25
	goto I26
I24:
	//fmt.Printf("I24 %v\n", r)
	r[5] = r[5] + 1
	// I25:
	//fmt.Printf("I25 %v\n", r)
	//r[1] = 17
	goto I18
I26:
	//fmt.Printf("I26 %v\n", r)
	r[3] = r[5]
	// I27:
	//fmt.Printf("I27 %v\n", r)
	//r[1] = 7
	goto I8
I28:
	//fmt.Printf("I28 %v\n", r)
	val := r[4]
	if last == 0 {
		fmt.Printf("Part 1 = %d\n", val)
	}
	if vals[val] {
		// Found duplicate
		fmt.Printf("Part 2 = %d\n", last)
		return
	}
	vals[val] = true
	last = val
	if r[4] == r[0] {
		r[5] = 1
	} else {
		r[5] = 0
	}
	// I29:
	//fmt.Printf("I29 %v\n", r)
	//r[1] = r[5] + r[1]
	if r[5] == 0 {
		goto I30
	} else {
		//fmt.Printf("DONE final=%v", r)
		return
	}
I30:
	//fmt.Printf("I30 %v\n", r)
	//r[1] = 5
	goto I6
}
