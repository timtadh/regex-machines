package machines

import . "github.com/timtadh/regex-machines/inst"
import "github.com/timtadh/regex-machines/queue"

func Lexer(program InstSlice, text []byte) (chan bool, chan int) {
		matches := make(chan int)
		success := make(chan bool)
		go func() {
				var cqueue, nqueue *queue.Queue = queue.New(), queue.New()
				last_match := -1
				last_match_tc := -1
				emitted := false
				cqueue.Push(0)
				for tc := 0; tc <= len(text); tc++ {
						if emitted {
								tc -= 1
								emitted = false
						}
						for !cqueue.Empty() {
								pc := cqueue.Pop()
								inst := program[pc]
								switch inst.Op {
								case CHAR:
										if int(tc) >= len(text) || text[tc] != byte(inst.X) {
												break
										}
										nqueue.Push(pc + 1)
								case MATCH:
										if last_match_tc < tc {
												last_match = int(pc)
												last_match_tc = tc
										} else if last_match > int(pc) {
												last_match = int(pc)
												last_match_tc = tc
										}
								case JMP:
										cqueue.Push(inst.X)
								case SPLIT:
										cqueue.Push(inst.X)
										cqueue.Push(inst.Y)
								}
						}
						cqueue, nqueue = nqueue, cqueue
						if cqueue.Empty() && last_match != -1 {
								matches <- last_match
								cqueue.Push(0)
								emitted = true
						}
				}
				if last_match_tc == len(text) {
						success <- true
				} else {
						success <- false
				}
		}()
		return success, matches
}

