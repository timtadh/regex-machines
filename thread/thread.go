package thread

import "container/list"

type Thread struct {
    Pc uint32
    Tc uint32
}

func NewThread(pc, tc uint32) *Thread {
    self := new(Thread)
    self.Pc = pc
    self.Tc = tc
    return self
}

type Stack struct {
    list *list.List
}

func NewStack() *Stack {
    self := new(Stack)
    self.list = list.New()
    return self
}

func (self *Stack) Empty() bool { return self.list.Len() <= 0 }

func (self *Stack) Push(t *Thread) {
    self.list.PushFront(t)
}

func (self *Stack) Pop() *Thread {
    e := self.list.Front()
    t, _ := e.Value.(*Thread)
    self.list.Remove(e)
    return t
}
