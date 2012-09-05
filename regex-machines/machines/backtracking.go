package machines

import . "regex-machines/inst"
import . "regex-machines/thread"

func Backtracking(program InstSlice, text []byte) bool {
    var stack *Stack = NewStack()
    var thread *Thread
    stack.Push(NewThread(0, 0))

    for !stack.Empty() {
        thread = stack.Pop()
    inner:
        for {
            if int(thread.Pc) >= len(program) || int(thread.Tc) > len(text) {
                return false
            }
            inst := program[thread.Pc]
            switch inst.Op {
            case CHAR:
                if int(thread.Tc) >= len(text) || text[thread.Tc] != byte(inst.X) {
                    break inner
                }
                thread.Pc++
                thread.Tc++
            case MATCH:
                if thread.Tc == uint32(len(text)) {
                    return true
                }
                break inner
            case JMP:
                thread.Pc = inst.X
            case SPLIT:
                stack.Push(NewThread(inst.Y, thread.Tc))
                thread.Pc = inst.X
            }
        }
    }
    return false
}
