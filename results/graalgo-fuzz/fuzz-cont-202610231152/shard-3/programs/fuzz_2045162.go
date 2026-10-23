package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
type GenS2 struct { F0 int64; F1 uint }
func genFn0() int64 {
    var v5 int32 = -2147483648
    fmt.Println(v5)
    _ = v5
    return 0
}

func main() {
    var v1 int64 = 1
    var v2 uint32 = 2147483647
    v3 := make(map[uint32]int)
    var v4 int = 713
    v1 = 9223372036854775807
    v6 := GenS1{F0: (v1 | v1)}
    v1 -= (v1 ^ v1)
    var v7 float64 = (((1.5 / (3.14159 + ((((1e10 + -2.25) + float64(((0.0 + (0.1 * ((0.0 / 0.1) + ((1e10 + 1.5) / ((0.0 + (1.5 + float32(0.1))) * (-2.25 / 1.0)))))) - float32(((0.0 + float64(float32(0.1))) / 0.1))))) / 3.14159) / (((-2.25 - 1.5) + float64((3.14159 * 1e10))) * ((3.14159 / ((1.5 + 1.5) / 1.0)) + 3.14159))))) + ((1.5 * 1e10) - 3.14159)) - 0.1)
    var v8 bool = false
    var v9 uint64 = 4294967295
    fmt.Println(v1)
    if (false || ((float32(float64(0.1)) < (3.14159 * (float32(v7) / ((0.0 + 1e10) - 1.5)))) && ((v2 - (v2 << ((1) & 63))) < 964))) {
        v7 = v7
    }
    v10 := []int{}
    fmt.Println(v7, v1, v7)
    switch "a" {
    case "\n\t":
        v7 += v7
    }
    var v11 uint16 = 2
    var v12 string = "\n\t"
    for fi13 := 0; fi13 < 5; fi13++ {
        var v14 uint16 = v11
        var v15 int8 = 1
        var v16 int8 = (v15 | v15)
        _ = fi13
        _ = v14
        _ = v15
        _ = v16
    }
    var v17 float64 = float64((0.1 * ((0.1 * (1e10 / (-2.25 / 3.14159))) + float32(v7))))
    _ = len(v3)
    for fi18 := 0; fi18 < 16; fi18++ {
        for fi19 := 0; fi19 < 2; fi19++ {
            var v20 int64 = 0
            v11 = v11
            fmt.Println(v17)
            _ = fi19
            _ = v20
        }
        switch (v4 / ((2147483647) | 1)) {
        case -860:
            v21 := GenS2{F0: genFn0(), F1: 18446744073709551615}
            _ = v21
        case 42:
            v1 += genFn0()
        default:
            for fi22 := 0; fi22 < 2; fi22++ {
                var v23 int32 = 246
                _ = fi22
                _ = v23
            }
        }
        if len(v10) >= 2 {
            v10 = v10[1:len(v10)]
        }
        _ = fi18
    }
    fmt.Println(v9, v9)
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v10:", len(v10), v10)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v17
    _ = v2
    _ = v3
    _ = v4
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
