package main

import "fmt"
import "os"
import "math"
import "inst"
import . "machines"

func test_case(n int) (inst.InstSlice, []byte) {
    program := make(inst.InstSlice, n*3+1)
    text := make([]byte, n)
    i := uint32(0)
    for j := 0; j < n; j++ {
        program[i] = inst.New(inst.SPLIT, i+1, i+2)
        program[i+1] = inst.New(inst.CHAR, 'a', 0)
        i += 2
    }
    for j := 0; j < n; j++ {
        text[j] = 'a'
        program[i] = inst.New(inst.CHAR, 'a', 0)
        i++
    }
    program[i] = inst.New(inst.MATCH, 0, 0)
    return program, text
}

func time() float64 {
    sec, nsec, _ := os.Time()
    return float64(sec) + float64(nsec)*math.Pow(10.0, -9)
}

func main() {
    //     fmt.Println("test, recursive, backtracking, thompson")
    for i := 1; i <= 50; i++ {
        program, text := test_case(i)
        var t1, t2, t3 float64
        if i <= 20 {
            {
                s := time()
                Recursive(program, text)
                e := time()
                t1 = e - s
            }
            {
                s := time()
                Backtracking(program, text)
                e := time()
                t2 = e - s
            }
        } else {
            t1 = 0.0
            t2 = 0.0
        }
        {
            s := time()
            Thompson(program, text)
            e := time()
            t3 = e - s
        }
        fmt.Printf("%v, %f, %f, %f\n", i, t1, t2, t3)
    }
}
