package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 int32; F2 int8 }
func genFn0(p0 int8) uint32 {
    p0 -= (p0 >> ((42) & 63))
    var v4 uint8 = 0
    v5 := []int{}
    fmt.Println(p0, p0)
    _ = p0
    _ = v4
    _ = v5
    return 2
}

func main() {
    var v1 float64 = -2.25
    v2 := make(map[int]string)
    v3 := []int{}
    fmt.Println(v1)
    var v6 uint64 = 142
    v1 = (((0.0 - ((-2.25 + (-2.25 - v1)) * v1)) - v1) + ((1e10 / 1.0) - (((v1 / (3.14159 - (float64(-2.25) / 0.1))) + -2.25) + ((0.1 / 0.1) - v1))))
    v7 := GenS1{F0: -32768, F1: -1, F2: int8(v6)}
    var v8 uint = 1
    for fi9 := 0; fi9 < 5; fi9++ {
        var v10 uint = (v8 - (^(v8)))
        for fi11 := 0; fi11 < 10; fi11++ {
            v1 = 0.1
            v8 -= uint(v6)
            _ = fi11
        }
        if (!(-62 < 19)) {
            v12 := GenS1{F0: 32767, F1: 928, F2: -128}
            _ = v12
        }
        _ = fi9
        _ = v10
    }
    var v13 float64 = v1
    v14 := GenS1{F0: 0, F1: 0, F2: 127}
    v15 := GenS1{F0: int16(v6), F1: 42, F2: 52}
    v16 := v2[-858]
    fmt.Println(v6, v8, v8)
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("v16:", v16)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v6:", v6)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v13
    _ = v14
    _ = v15
    _ = v16
    _ = v2
    _ = v3
    _ = v6
    _ = v7
    _ = v8
}
