package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }

func main() {
    var v1 int64 = 42
    v2 := make(map[int]bool)
    var v3 uint32 = 861
    var v4 uint64 = 498
    var v5 int32 = 1
    v6 := []int{9223372036854775807}
    for fi7 := 0; fi7 < 20; fi7++ {
        v8 := GenS1{F0: 2147483647}
        _ = fi7
        _ = v8
    }
    v5 = (v5 % ((v5) | 1))
    fmt.Println(v4)
    v5 = v5
    var v9 int8 = 127
    if true {
        switch 530 {
        case -305:
            fmt.Println(v1, v1, v3)
        case -142:
            var v10 uint64 = (v4 + 1)
            _ = v10
        default:
            fmt.Println(v5)
        }
        for fi11 := 0; fi11 < 1; fi11++ {
            var v12 int = (fi11 / ((1) | 1))
            _ = len(v6)
            _ = fi11
            _ = v12
        }
    }
    var v13 int8 = v9
    v13 -= (v13 | (v9 % ((1) | 1)))
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v13
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v9
}
