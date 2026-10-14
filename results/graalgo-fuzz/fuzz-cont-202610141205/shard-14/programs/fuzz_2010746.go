package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
type GenS2 struct { F0 bool; F1 int64; F2 int8 }

func main() {
    v1 := make(map[int]int)
    v2 := []int{2147483647, 0, 289, -2147483648}
    v3 := make(map[string]uint32)
    var v4 uint64 = 2
    var v5 string = ""
    fmt.Println(v5, v4)
    v5 = "hello"
    fmt.Println(v4, v5)
    switch 4294967295 {
    case 127:
        var v6 uint = 179
        _ = v6
    case -1:
        fmt.Println(v4, v5)
    }
    if ((42 < 0) && ((1e10 + 0.1) > 0.0)) {
        _ = len(v2)
        for fi7 := 0; fi7 < 14; fi7++ {
            v5 = "\n\t"
            _ = fi7
        }
    } else {
        fmt.Println(v5)
    }
    var v8 float32 = (1e10 * 0.0)
    v9 := GenS1{F0: 42}
    if true {
        var v10 uint64 = v4
        v8 -= (1.5 * v8)
        _ = v10
    }
    v11 := GenS2{F0: (false || true), F1: -144, F2: 0}
    v4 = v4
    v8 += v8
    v5 = v5
    if len(v2) > 0 {
        v2[0] = 4294967295
    }
    v12 := GenS2{F0: (0 >= 1), F1: 0, F2: 42}
    fmt.Println(v4, v4, v8)
    for fi13 := 0; fi13 < 9; fi13++ {
        v4 -= v4
        _ = fi13
    }
    var v14 uint16 = 1
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v14:", v14)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v11
    _ = v12
    _ = v14
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v8
    _ = v9
}
