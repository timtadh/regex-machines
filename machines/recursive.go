package machines

import "inst"

func recursive(program inst.InstSlice, text []byte, pc, tc uint32) bool {
    if int(pc) >= len(program) || int(tc) > len(text) {
        return false
    }
    if program[pc].Op == inst.MATCH && tc == uint32(len(text)) {
        return true
    } else if program[pc].Op == inst.MATCH && tc != uint32(len(text)) {
        return false
    } else if program[pc].Op == inst.CHAR && tc == uint32(len(text)) {
        return false
    }
    switch program[pc].Op {
        case inst.CHAR:
            if text[tc] != byte(program[pc].X) {
                return false
            }
            return recursive(program, text, pc+1, tc+1)
        case inst.JMP:
            return recursive(program, text, program[pc].X, tc)
        case inst.SPLIT:
            if recursive(program, text, program[pc].X, tc) {
                return true
            }
            return recursive(program, text, program[pc].Y, tc)
    }
    return false
}

func Recursive(program inst.InstSlice, text []byte) bool {
    return recursive(program, text, 0, 0)
}
