package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 uint16 }
type GenS2 struct { F0 string; F1 int16; F2 uint8 }
func genFn0() uint8 {
    var v7 float64 = float64((float32(((1.5 + (float64(((0.0 - 0.0) / 1e10)) - 0.1)) * (((1.5 - ((0.0 * float64((1.5 * 0.1))) * 1e10)) - (0.0 + 0.0)) / -2.25))) + float32(1.5)))
    var v8 int64 = 0
    for fi9 := 0; fi9 < 1; fi9++ {
        fmt.Println(v7, fi9, v7)
        var v10 float32 = (0.0 + 1e10)
        _ = fi9
        _ = v10
    }
    v7 = 3.14159
    _ = v7
    _ = v8
    return 1
}

func main() {
    var v1 float32 = -2.25
    var v2 int64 = 568
    v3 := make(map[uint32]bool)
    var v4 uint64 = 374
    var v5 int64 = -31
    v6 := make(map[int]int)
    v11 := []int{int(v5)}
    var v12 int32 = 1
    if len(v11) > 0 {
        v11[0] = 32767
    }
    _ = len(v6)
    var v13 float32 = 1e10
    for fi14 := 0; fi14 < 8; fi14++ {
        for fi15 := 0; fi15 < 5; fi15++ {
            if true {
                var v16 uint = 255
                v17 := v6[-9223372036854775808]
                v18 := GenS1{F0: 4294967295, F1: 42}
                _ = v16
                _ = v17
                _ = v18
            }
            _ = fi15
        }
        switch "" {
        case "":
            if true {
                v19 := GenS2{F0: "", F1: 1, F2: 127}
                v1 = (v13 * v13)
                _ = v19
            }
        }
        if len(v11) > 0 {
            v11[0] = fi14
        }
        _ = fi14
    }
    v5 = (v2 & (v5 / ((0) | 1)))
    v20 := GenS1{F0: 621, F1: 383}
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v11:", len(v11), v11)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v11
    _ = v12
    _ = v13
    _ = v2
    _ = v20
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
