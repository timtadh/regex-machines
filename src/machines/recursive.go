package machines

import "inst"

func recursive(program inst.InstSlice, text []byte, pc, tc uint32) bool {
    if int(pc) >= len(program) || int(tc) > len(text) {
        return false
    }
    switch program[pc].Op {
        case inst.MATCH:
            if tc == uint32(len(text)) {
                return true
            }
            return false
        case inst.CHAR:
            if tc == uint32(len(text)) || text[tc] != byte(program[pc].X) {
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
