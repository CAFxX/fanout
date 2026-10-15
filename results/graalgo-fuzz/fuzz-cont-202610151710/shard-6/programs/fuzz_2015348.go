package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 int64 }
func genFn0(p0 int, p1 int64) int8 {
    v4 := []int{}
    var v5 int = p0
    _ = p0
    _ = p1
    _ = v4
    _ = v5
    return -128
}

func main() {
    var v1 int16 = 1000
    v2 := []int{}
    var v3 int16 = v1
    for fi6 := 0; fi6 < 20; fi6++ {
        fmt.Println(v1, v1, fi6)
        _ = fi6
    }
    if len(v2) > 0 {
        v2[0] = 0
    }
    v1 -= (v3 + (v3 % ((v3) | 1)))
    v7 := GenS1{F0: 2147483648, F1: 676}
    var v8 uint = 42
    var v9 string = "\n\t"
    for fi10 := 0; fi10 < 3; fi10++ {
        v11 := GenS1{F0: 688, F1: 9223372036854775807}
        _ = fi10
        _ = v11
    }
    v1 = 729
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v2
    _ = v3
    _ = v7
    _ = v8
    _ = v9
}
