package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 int }
type GenS2 struct { F0 uint8; F1 uint32; F2 int }
func genFn0(p0 uint64) int32 {
    fmt.Println(p0)
    if (684 <= 1) {
        var v5 float64 = 1e10
        v6 := GenS1{F0: -650, F1: -2147483648}
        _ = v5
        _ = v6
    }
    v7 := GenS2{F0: 29, F1: 380, F2: 994}
    for fi8 := 0; fi8 < 14; fi8++ {
        for fi9 := 0; fi9 < 9; fi9++ {
            var v10 int64 = -842
            _ = fi9
            _ = v10
        }
        fmt.Println(fi8, fi8)
        _ = fi8
    }
    _ = p0
    _ = v7
    return 342
}

func main() {
    var v1 int8 = 1
    v2 := []string{"\n\t"}
    v3 := make(map[uint32]bool)
    var v4 float64 = float64((-2.25 + -2.25))
    _ = len(v3)
    var v11 float64 = v4
    _ = len(v3)
    v12 := GenS1{F0: -9223372036854775808, F1: 127}
    var v13 uint32 = 2
    v13 = v13
    v14 := GenS1{F0: -33, F1: 1}
    if (127 >= 4294967295) {
        switch 0 {
        case -383:
            v1 = v1
        }
        v3[(v13 - 2147483647)] = (1 == 65535)
    }
    for fi15 := 0; fi15 < 11; fi15++ {
        v16 := GenS1{F0: -1, F1: fi15}
        _ = fi15
        _ = v16
    }
    v4 = (((v4 / v4) * v11) + 0.1)
    v17 := GenS2{F0: 170, F1: v13, F2: 9223372036854775807}
    v13 -= v13
    v13 = v13
    var v18 uint16 = 612
    v4 += 0.1
    for fi19 := 0; fi19 < 16; fi19++ {
        var v20 float64 = v11
        _ = fi19
        _ = v20
    }
    var v21 float64 = float64(3.14159)
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v13:", v13)
    fmt.Println("v18:", v18)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v21:", v21)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    _ = v1
    _ = v11
    _ = v12
    _ = v13
    _ = v14
    _ = v17
    _ = v18
    _ = v2
    _ = v21
    _ = v3
    _ = v4
}
