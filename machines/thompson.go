package machines

import . "github.com/timtadh/regex-machines/inst"
import "github.com/timtadh/regex-machines/queue"

func Thompson(program InstSlice, text []byte) bool {
    var cqueue, nqueue *queue.Queue = queue.New(), queue.New()
    cqueue.Push(0)
    for tc := 0; tc <= len(text); tc++ {
        for !cqueue.Empty() {
            pc := cqueue.Pop()
            inst := program[pc]
            switch inst.Op {
            case CHAR:
                if int(tc) < len(text) && text[tc] == byte(inst.X) {
										nqueue.Push(pc + 1)
                }
            case MATCH:
                if tc == len(text) {
                    return true
                }
            case JMP:
                cqueue.Push(inst.X)
            case SPLIT:
                cqueue.Push(inst.X)
                cqueue.Push(inst.Y)
            }
        }
        cqueue, nqueue = nqueue, cqueue
    }
    return false
}
