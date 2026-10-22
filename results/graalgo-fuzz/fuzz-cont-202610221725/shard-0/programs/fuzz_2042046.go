package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 bool }

func main() {
    v1 := []int{2147483647, 0, -9223372036854775808}
    v2 := []uint32{}
    var v3 bool = true
    v4 := []int{-9223372036854775808, 2147483647, 42}
    v5 := make(map[int]string)
    var v6 float64 = 1e10
    for fi7 := 0; fi7 < 5; fi7++ {
        v3 = (1 < -128)
        for fi8 := 0; fi8 < 8; fi8++ {
            var v9 int = fi8
            _ = len(v5)
            var v10 int64 = 1
            _ = fi8
            _ = v10
            _ = v9
        }
        _ = fi7
    }
    v6 += (0.1 / (float64(((0.0 * (-2.25 - 0.0)) - 1.5)) - -2.25))
    v11 := GenS1{F0: 2147483647, F1: (((2147483647 <= -2147483648) && v3) || (1 > 536))}
    if v3 {
        v6 -= (v6 - 3.14159)
    } else {
        var v12 uint8 = 127
        v13 := GenS1{F0: 1, F1: v3}
        v6 += v6
        _ = v12
        _ = v13
    }
    var v14 uint = 42
    fmt.Println(v6)
    _ = len(v5)
    for fi15 := 0; fi15 < 15; fi15++ {
        fmt.Println(v3)
        v14 += v14
        _ = fi15
    }
    v6 = (float64(((0.0 / ((-2.25 / float32(v6)) / 0.1)) + (0.1 - (0.1 + 0.0)))) - float64(0.1))
    if (false || (v3 && true)) {
        v3 = (32768 < 994)
        v16 := v5[-347]
        _ = v16
    } else {
        fmt.Println(v6)
        var v17 bool = true
        _ = v17
    }
    var v18 int16 = -32768
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v14:", v14)
    fmt.Println("v18:", v18)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v11
    _ = v14
    _ = v18
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
