package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 bool; F2 int16 }

func main() {
    var v1 int32 = 1
    var v2 uint8 = 1
    var v3 uint64 = 509
    var v4 uint64 = 42
    var v5 bool = false
    v6 := GenS1{F0: v4, F1: v5, F2: 1}
    if true {
        v5 = v5
    }
    for fi7 := 0; fi7 < 10; fi7++ {
        fmt.Println(v5)
        _ = fi7
    }
    v8 := []int{127, -169, -2147483648}
    v4 += 42
    var v9 int8 = int8(v2)
    fmt.Println(v4)
    v10 := GenS1{F0: (v3 - 407), F1: true, F2: -32768}
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v8:", len(v8), v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v8
    _ = v9
}
