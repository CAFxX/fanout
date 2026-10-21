package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 uint64 }
func genFn0() string {
    var v5 int8 = 100
    var v6 uint16 = 32767
    var v7 uint64 = 890
    _ = v5
    _ = v6
    _ = v7
    return "a"
}

func main() {
    v1 := make(map[uint32]bool)
    v2 := []uint32{}
    var v3 float64 = float64((1.5 * 3.14159))
    var v4 uint64 = 42
    var v8 int64 = -166
    if (!false) {
        v9 := GenS1{F0: -98, F1: 765}
        _ = v9
    }
    v2 = append(v2, 4294967295)
    var v10 bool = false
    for fi11 := 0; fi11 < 17; fi11++ {
        v12 := v1[42]
        _ = fi11
        _ = v12
    }
    v4 = (v4 | v4)
    v4 += v4
    var v13 int64 = (v8 ^ (v8 + (-(v8))))
    if (float64(float32(((0.0 + (3.14159 + 3.14159)) / float64((-2.25 - 0.0))))) >= v3) {
        var v14 uint32 = 857
        _ = v14
    }
    for fi15 := 0; fi15 < 13; fi15++ {
        var v16 bool = ((2147483647 != 722) || false)
        var v17 int8 = 127
        var v18 int64 = (v13 ^ v8)
        _ = fi15
        _ = v16
        _ = v17
        _ = v18
    }
    v10 = (127 > 255)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v10:", v10)
    fmt.Println("v13:", v13)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v10
    _ = v13
    _ = v2
    _ = v3
    _ = v4
    _ = v8
}
