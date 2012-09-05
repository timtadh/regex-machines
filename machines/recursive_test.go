package machines

import "testing"
import "github.com/timtadh/regex-machines/inst"

func TestRecursiveMatch(t *testing.T) {
    //. (a|b)*cba?(c|b)bb
    program := make(inst.InstSlice, 20)

    program[0] = inst.New(inst.SPLIT, 1, 6)
    program[1] = inst.New(inst.SPLIT, 2, 4)
    program[2] = inst.New(inst.CHAR, 'a', 0)
    program[3] = inst.New(inst.JMP, 5, 0)
    program[4] = inst.New(inst.CHAR, 'b', 0)
    program[5] = inst.New(inst.JMP, 0, 0)
    program[6] = inst.New(inst.CHAR, 'c', 0)
    program[7] = inst.New(inst.CHAR, 'b', 0)
    program[8] = inst.New(inst.SPLIT, 9, 10)
    program[9] = inst.New(inst.CHAR, 'a', 0)
    program[10] = inst.New(inst.SPLIT, 11, 13)
    program[11] = inst.New(inst.CHAR, 'c', 0)
    program[12] = inst.New(inst.JMP, 14, 0)
    program[13] = inst.New(inst.CHAR, 'b', 0)
    program[14] = inst.New(inst.CHAR, 'b', 0)
    program[15] = inst.New(inst.CHAR, 'b', 0)
    program[16] = inst.New(inst.MATCH, 0, 0)

    t.Log(string(text))
    t.Log(program)
    if !Recursive(program, text) {
        t.Error("program should have matched text but did not")
    }
}

func TestRecursiveNoMatch(t *testing.T) {
    //. (a|b)*cba?(c|b)bb
    program := make(inst.InstSlice, 20)

    program[0] = inst.New(inst.SPLIT, 1, 6)
    program[1] = inst.New(inst.SPLIT, 2, 4)
    program[2] = inst.New(inst.CHAR, 'a', 0)
    program[3] = inst.New(inst.JMP, 5, 0)
    program[4] = inst.New(inst.CHAR, 'b', 0)
    program[5] = inst.New(inst.JMP, 0, 0)
    program[6] = inst.New(inst.CHAR, 'c', 0)
    program[7] = inst.New(inst.CHAR, 'b', 0)
    program[8] = inst.New(inst.SPLIT, 9, 10)
    program[9] = inst.New(inst.CHAR, 'a', 0)
    program[10] = inst.New(inst.SPLIT, 11, 13)
    program[11] = inst.New(inst.CHAR, 'c', 0)
    program[12] = inst.New(inst.JMP, 14, 0)
    program[13] = inst.New(inst.CHAR, 'b', 0)
    program[14] = inst.New(inst.CHAR, 'b', 0)
    program[15] = inst.New(inst.MATCH, 0, 0)

    t.Log(string(text))
    t.Log(program)
    if Recursive(program, text) {
        t.Error("program should not have matched text but did")
    }
}
