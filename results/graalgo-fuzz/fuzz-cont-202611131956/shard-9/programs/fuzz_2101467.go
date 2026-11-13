package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 string; F2 int32 }
type GenS2 struct { F0 uint8 }
func genFn0(p0 bool, p1 int, p2 uint64) int64 {
    p0 = ((p2 | (p2 / ((2) | 1))) >= p2)
    for fi7 := 0; fi7 < 14; fi7++ {
        var v8 int8 = -105
        _ = fi7
        _ = v8
    }
    _ = p0
    _ = p1
    _ = p2
    return 0
}

func main() {
    v1 := make(map[string]bool)
    var v2 uint16 = 255
    v3 := []uint32{}
    var v4 int32 = -2147483648
    var v5 int16 = 1
    v6 := make(map[int]int)
    v9 := GenS1{F0: v5, F1: "", F2: v4}
    fmt.Println(v2, v4)
    v2 = (v2 ^ 0)
    fmt.Println(v2, v2, v5)
    v3 = append(v3, 2147483648)
    v5 -= (v5 | v5)
    var v10 int16 = (v5 * 1000)
    fmt.Println(v5, v2)
    v11 := GenS2{F0: 0}
    v12 := GenS2{F0: 128}
    v5 += 1000
    v13 := GenS2{F0: 1}
    fmt.Println(v2, v4)
    for fi14 := 0; fi14 < 7; fi14++ {
        v15 := GenS1{F0: 354, F1: "a", F2: v4}
        for fi16 := 0; fi16 < 9; fi16++ {
            v5 = (v10 & 1000)
            _ = fi16
        }
        v3 = append(v3, 2147483648)
        _ = fi14
        _ = v15
    }
    fmt.Println(v4, v2, v5)
    v17 := v1[("a" + "")]
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v10:", v10)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v13
    _ = v17
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v9
}
