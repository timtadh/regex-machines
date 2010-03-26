package inst

import "fmt"

const (
    CHAR    = iota
    SPLIT
    JMP
    MATCH
)

type Inst struct {
    Op  uint8
    X   uint32
    Y   uint32
}

type InstSlice []*Inst

func New(op uint8, x, y uint32) *Inst {
    self := new(Inst)
    self.Op = op
    self.X = x
    self.Y = y
    return self
}

func (self Inst) String() (s string) {
    switch self.Op {
        case CHAR:
            s = fmt.Sprintf("CHAR %c", byte(self.X))
        case SPLIT:
            s = fmt.Sprintf("SPLIT %v, %v", self.X, self.Y)
        case JMP:
            s = fmt.Sprintf("JMP %v", self.X)
        case MATCH:
            s = "MATCH"
    }
    return
}

func (self InstSlice) String() (s string) {
    s = "{\n"
    for i, inst := range self {
        if inst == nil { continue }
        s += fmt.Sprintf("    %v %v\n", i, inst)
    }
    s += "}"
    return
}
