package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 string; F2 uint16 }
type GenS2 struct { F0 int64; F1 int }

func main() {
    var v1 uint64 = 18446744073709551615
    var v2 int8 = 0
    var v3 int = 859
    var v4 int64 = 712
    var v5 string = "a"
    fmt.Println(v4, v4)
    switch v5 {
    case "\n\t":
        if false {
            v6 := GenS1{F0: 2147483648, F1: (v5 + v5), F2: 2}
            _ = v6
        }
    default:
        fmt.Println(v2)
    }
    v1 -= v1
    v7 := GenS1{F0: 1, F1: "a", F2: 977}
    v2 += 1
    var v8 uint8 = 42
    v9 := GenS2{F0: -951, F1: v3}
    var v10 uint = 42
    v3 = -1
    fmt.Println(v5)
    for fi11 := 0; fi11 < 3; fi11++ {
        fmt.Println(v2, v10, fi11)
        v12 := GenS1{F0: 313, F1: v5, F2: 2}
        _ = fi11
        _ = v12
    }
    var v13 int16 = 282
    v14 := []int{}
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v13:", v13)
    fmt.Println("slicelen v14:", len(v14), v14)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v10
    _ = v13
    _ = v14
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v7
    _ = v8
    _ = v9
}
