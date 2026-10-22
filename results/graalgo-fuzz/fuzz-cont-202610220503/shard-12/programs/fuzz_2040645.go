package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 int; F2 bool }
type GenS2 struct { F0 bool; F1 uint64 }
func genFn0(p0 uint32, p1 uint16) int8 {
    v7 := GenS1{F0: "\n\t", F1: 179, F2: false}
    p0 -= (p0 << ((143) & 63))
    _ = p0
    _ = p1
    _ = v7
    return 100
}

func main() {
    v1 := []int{-1, -130, 597}
    var v2 int16 = 0
    var v3 float64 = (1e10 / ((3.14159 / 3.14159) + (0.0 / 0.1)))
    var v4 int64 = 361
    var v5 float64 = (v3 / -2.25)
    v6 := make(map[int]int)
    v8 := GenS2{F0: true, F1: 0}
    v9 := GenS2{F0: true, F1: 42}
    v2 = (v2 | (v2 + -1))
    for fi10 := 0; fi10 < 13; fi10++ {
        for fi11 := 0; fi11 < 7; fi11++ {
            v5 += (v5 + v3)
            _ = fi11
        }
        var v12 int = fi10
        _ = fi10
        _ = v12
    }
    v4 -= v4
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    delete(v6, -9223372036854775808)
    for fi13 := 0; fi13 < 20; fi13++ {
        v2 -= (v2 * -729)
        if false {
            v2 += v2
        } else {
            for fi14 := 0; fi14 < 8; fi14++ {
                var v15 int32 = 0
                v3 += v5
                _ = fi14
                _ = v15
            }
        }
        fmt.Println(v3, v3, v3)
        _ = fi13
    }
    _ = len(v1)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v8
    _ = v9
}
