package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 int32; F2 int8 }
func genFn0(p0 uint32, p1 uint64, p2 uint8) uint8 {
    var v4 uint64 = uint64(p0)
    p0 = 0
    v5 := GenS1{F0: ("hello" + (("" + ("" + "hello")) + "\n\t")), F1: -1, F2: 1}
    _ = p0
    _ = p1
    _ = p2
    _ = v4
    _ = v5
    return (p2 & (p2 / ((128) | 1)))
}

func main() {
    var v1 uint16 = 952
    var v2 int64 = 0
    var v3 uint32 = 2147483647
    v2 -= v2
    v6 := []int{42}
    v2 = (v2 + (v2 / ((v2) | 1)))
    v7 := GenS1{F0: "\n\t", F1: 1, F2: -1}
    var v8 float32 = 0.1
    var v9 uint8 = 41
    v3 = (v3 % ((v3) | 1))
    v2 = v2
    v8 = -2.25
    for fi10 := 0; fi10 < 9; fi10++ {
        if len(v6) > 0 {
            v6[0] = (fi10 | -343)
        }
        _ = fi10
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v6:", len(v6), v6)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v2
    _ = v3
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
