package main

import "fmt"

func main() {
	r := []int{1, 0, 0, 0, 0, 0}
	// I0:
	//r[3] = r[3] + 16
	goto I17
	I1:
		fmt.Printf("Finding sum of factors of %d\n", r[5])
	r[4] = 1
	 I2:
	r[2] = 1
	 I3:
	r[1] = r[4] * r[2]
	// I4:
	if r[1] == r[5] {
		r[1]=1
	} else {
		r[1]=0
	}
	// I5:
	//r[3] = r[1] + r[3]
	if r[1] == 1 { goto I7 }
	// I6:
	//r[3] = r[3] + 1
	goto I8
	 I7:
	r[0] = r[4] + r[0]
	I8:
	r[2] = r[2] + 1
	// I9:
	if r[2] > r[5] {
		r[1]=1
	} else {
		r[1]=0
	}
	// I10:
	//r[3] = r[3] + r[1]
	if r[1] == 1 { goto I12 }
	// I11:
	//r[3] = 2
	goto I3
	 I12:
	r[4] = r[4] + 1
	// I13:
	if r[4] > r[5] {
		r[1]=1
	} else {
		r[1]=0
	}
	// I14:
	//r[3] = r[1] + r[3]
	if r[1] == 1 {goto I16}
	// I15:
	//r[3] = 1
	goto I2
	 I16:
	//r[3] = r[3] * r[3]
	fmt.Println("Solved")
	return
	 I17:
	r[5] = r[5] + 2
	// I18:
	r[5] = r[5] * r[5]
	// I19:
	//r[5] = r[3] * r[5]
	r[5] *= 19
	// I20:
	r[5] = r[5] * 11
	// I21:
	r[1] = r[1] + 6
	// I22:
	//r[1] = r[1] * r[3]
	r[1] *= 22
	// I23:
	r[1] = r[1] + 13
	// I24:
	r[5] = r[5] + r[1]
	// I25:
	//r[3] = r[3] + r[0]
	if r[0] == 1 { goto I27 }
	// I26:
	//r[3] = 0
	goto I1
	 I27:
	//r[1] = r[3]
	r[1] = 27
	// I28:
	//r[1] = r[1] * r[3]
	r[1] *= 28
	// I29:
	//r[1] = r[3] + r[1]
	r[1] += 29
	// I30:
	//r[1] = r[3] * r[1]
	r[1] *= 30
	// I31:
	//r[1] = r[1] * 14
	r[1] *= 14
	// I32:
	//r[1] = r[1] * r[3]
	r[1] *= 32
	// I33:
	r[5] = r[5] + r[1]
	// I34:
	r[0] = 0
	// I35:
	//r[3] = 0
	goto I1
}
