package machines

import "testing"
import "github.com/timtadh/regex-machines/inst"

// var text []byte = []byte{'a', 'b', 'a'}

func TestLexerMatch(t *testing.T) {
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
		t.Log(len(text))
    t.Log(program)
		success, matches := LexerEngine(program, text)
		go func() {
				for match := range matches {
						t.Log(match)
				}
		}()

		if ok := <-success; !ok {
        t.Error("program should have matched text but did not")
		}
}

func TestLexerNoMatch(t *testing.T) {
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
    program[14] = inst.New(inst.MATCH, 0, 0)

		t.Log("(a|b)*cba?(c|b)")
    t.Log(string(text))
    t.Log(program)
		success, matches := LexerEngine(program, text)
		go func() {
				for match := range matches {
						t.Log(match)
				}
		}()
		if ok := <-success; ok {
        t.Error("program should not have matched text but did")
		}
}

func TestLexerThreeStrings(t *testing.T) {
    //. (a|b)*cba?(c|b)bb
		var text []byte = []byte{'s', 't', 'r', 'u', 'c', 't', ' ', '*'}
    program := make(inst.InstSlice, 30)

    program[0] = inst.New(inst.SPLIT, 2, 1) // go to 1 or 2/3
    program[1] = inst.New(inst.SPLIT, 9, 13)  // go to 2 or 3
    program[2] = inst.New(inst.CHAR, 's', 0)
    program[3] = inst.New(inst.CHAR, 't', 0)
    program[4] = inst.New(inst.CHAR, 'r', 0)
    program[5] = inst.New(inst.CHAR, 'u', 0)
    program[6] = inst.New(inst.CHAR, 'c', 0)
    program[7] = inst.New(inst.CHAR, 't', 0)
    program[8] = inst.New(inst.MATCH, 0, 0)
    program[9] = inst.New(inst.SPLIT, 10, 12)
    program[10] = inst.New(inst.CHAR, ' ', 0)
    program[11] = inst.New(inst.JMP, 9, 0)
    program[12] = inst.New(inst.MATCH, 0, 0)
    program[13] = inst.New(inst.CHAR, '*', 0)
    program[14] = inst.New(inst.MATCH, 0, 0)

    t.Log(string(text))
		t.Log(len(text))
    t.Log(program)
		success, matches := LexerEngine(program, text)
		go func() {
				for match := range matches {
						t.Log(match)
				}
		}()

		if ok := <-success; !ok {
        t.Error("program should have matched text but did not")
		}
}
