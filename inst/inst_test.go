package inst

import "testing"
import "fmt"

func TestPrint(t *testing.T) {
    i := New(CHAR, uint32('a'), 0)
    j := New(MATCH, 0, 0)
    k := New(JMP, 14, 0)
    l := New(SPLIT, 15, 17)
    fmt.Println(i)
    fmt.Println(j)
    fmt.Println(k)
    fmt.Println(l)
    s := make(InstSlice, 4)
    s[0] = i
    s[1] = j
    s[2] = k
    s[3] = l
    fmt.Println(s)
}
