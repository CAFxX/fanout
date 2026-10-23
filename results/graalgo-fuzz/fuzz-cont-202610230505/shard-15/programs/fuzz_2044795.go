package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 bool; F2 int32 }
type GenS2 struct { F0 uint16 }

func main() {
    var v1 int64 = -1
    var v2 uint64 = 0
    var v3 uint32 = 426
    v4 := make(map[int]int)
    fmt.Println(v1, v3)
    v3 -= v3
    v2 -= v2
    v5 := GenS1{F0: "a", F1: (false || true), F2: 1}
    for fi6 := 0; fi6 < 14; fi6++ {
        v7 := GenS2{F0: 2}
        _ = fi6
        _ = v7
    }
    var v8 uint64 = 392
    var v9 int8 = -1
    var v10 int = -251
    var v11 int64 = v1
    v11 -= (-(v11))
    v12 := GenS2{F0: 32767}
    if false {
        var v13 uint64 = (v2 * (v8 & 775))
        _ = v13
    } else {
        switch (v10 / ((v10) | 1)) {
        case -1:
            var v14 float64 = 3.14159
            _ = v14
        case 775:
            v1 = v1
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v8
    _ = v9
}
