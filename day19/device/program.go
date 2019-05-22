package device

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type Instruction struct {
	operation string
	a, b, c   int
}

type Program struct {
	IP           int
	Instructions []Instruction
}

func Parse(filename string) *Program {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var program Program

	scanner.Scan()
	ipLine := scanner.Text()
	program.IP, _ = strconv.Atoi(ipLine[4:])

	for scanner.Scan() {
		instructionLine := scanner.Text()
		parts := strings.Split(instructionLine, " ")
		a, _ := strconv.Atoi(parts[1])
		b, _ := strconv.Atoi(parts[2])
		c, _ := strconv.Atoi(parts[3])
		program.Instructions = append(program.Instructions, Instruction{parts[0], a, b, c})
	}
	return &program
}

func (p Program) ToGo() string {
	var sb strings.Builder
	sb.WriteString("r := make([]int, 6)\n")
	for i, instruction := range p.Instructions {
		sb.WriteString(fmt.Sprintf("// I%d:\n", i))
		sb.WriteString(instruction.toGo())
		sb.WriteByte('\n')
	}
	return sb.String()
}

func (i Instruction) toGo() string {
	switch i.operation {
	case "addr":
		return fmt.Sprintf("r[%d] = r[%d] + r[%d]", i.c, i.a, i.b)
	case "addi":
		return fmt.Sprintf("r[%d] = r[%d] + %d", i.c, i.a, i.b)
	case "mulr":
		return fmt.Sprintf("r[%d] = r[%d] * r[%d]", i.c, i.a, i.b)
	case "muli":
		return fmt.Sprintf("r[%d] = r[%d] * %d", i.c, i.a, i.b)
	case "banr":
		return fmt.Sprintf("r[%d] = r[%d] & r[%d]", i.c, i.a, i.b)
	case "bani":
		return fmt.Sprintf("r[%d] = r[%d] & %d", i.c, i.a, i.b)
	case "borr":
		return fmt.Sprintf("r[%d] = r[%d] | r[%d]", i.c, i.a, i.b)
	case "bori":
		return fmt.Sprintf("r[%d] = r[%d] | %d", i.c, i.a, i.b)
	case "setr":
		return fmt.Sprintf("r[%d] = r[%d]", i.c, i.a)
	case "seti":
		return fmt.Sprintf("r[%d] = %d", i.c, i.a)
	case "gtir":
		return fmt.Sprintf("if %d > r[%d] {\n\tr[%d]=1\n} else {\n\tr[%[3]d]=0\n}", i.a, i.b, i.c)
	case "gtri":
		return fmt.Sprintf("if r[%d] > %d {\n\tr[%d]=1\n} else {\n\tr[%[3]d]=0\n}", i.a, i.b, i.c)
	case "gtrr":
		return fmt.Sprintf("if r[%d] > r[%d] {\n\tr[%d]=1\n} else {\n\tr[%[3]d]=0\n}", i.a, i.b, i.c)
	case "eqir":
		return fmt.Sprintf("if %d == r[%d] {\n\tr[%d]=1\n} else {\n\tr[%[3]d]=0\n}", i.a, i.b, i.c)
	case "eqri":
		return fmt.Sprintf("if r[%d] == %d {\n\tr[%d]=1\n} else {\n\tr[%[3]d]=0\n}", i.a, i.b, i.c)
	case "eqrr":
		return fmt.Sprintf("if r[%d] == r[%d] {\n\tr[%d]=1\n} else {\n\tr[%[3]d]=0\n}", i.a, i.b, i.c)
	}
	panic(fmt.Errorf("unknown operation %v", i.operation))
}
