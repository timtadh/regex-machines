package machines

import "fmt"
import . "github.com/timtadh/regex-machines/inst"
import "github.com/timtadh/regex-machines/queue"

type Match struct {
		Bytes []byte
		PC int
}

func (self Match) String() string {
		return fmt.Sprintf("<Match %v '%v'>", self.PC, string(self.Bytes))
}

func LexerEngine(program InstSlice, text []byte) (chan bool, chan Match) {
		matches := make(chan Match)
		success := make(chan bool)
		go func() {
				var cqueue, nqueue *queue.Queue = queue.New(), queue.New()
				match := -1
				match_tc := -1
				start_tc := 0
				cqueue.Push(0)
				for tc := 0; tc <= len(text); tc++ {
						for !cqueue.Empty() {
								pc := cqueue.Pop()
								inst := program[pc]
								// fmt.Printf("%v tc=%v\n", inst, tc)
								switch inst.Op {
								case CHAR:
										if int(tc) >= len(text) || text[tc] != byte(inst.X) {
												break
										}
										nqueue.Push(pc + 1)
								case MATCH:
										if match_tc < tc {
												match = int(pc)
												match_tc = tc
										} else if match > int(pc) {
												match = int(pc)
												match_tc = tc
										}
								case JMP:
										cqueue.Push(inst.X)
								case SPLIT:
										cqueue.Push(inst.X)
										cqueue.Push(inst.Y)
								}
						}
						cqueue, nqueue = nqueue, cqueue
						if cqueue.Empty() && match != -1 {
								matches <- Match{text[start_tc:match_tc], match}
								match = -1
								cqueue.Push(0)
								start_tc = tc
								tc -= 1
						}
				}
				close(matches)
				if match_tc == len(text) {
						success <- true
				} else {
						success <- false
				}
				close(success)
		}()
		return success, matches
}

