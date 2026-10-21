package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 uint8 }

func main() {
    var v1 uint32 = 2147483648
    v2 := make(map[int]uint32)
    var v3 int8 = -1
    var v4 bool = false
    var v5 int16 = 1000
    for fi6 := 0; fi6 < 12; fi6++ {
        var v7 int32 = -2147483648
        for fi8 := 0; fi8 < 3; fi8++ {
            v3 = (v3 - (^(v3)))
            var v9 uint8 = 0
            v10 := GenS1{F0: 378, F1: (v9 << ((4294967295) & 63))}
            v11 := v2[fi8]
            _ = fi8
            _ = v10
            _ = v11
            _ = v9
        }
        _ = fi6
        _ = v7
    }
    v12 := GenS1{F0: 1, F1: 4}
    v5 = v5
    var v13 float64 = ((0.0 * 0.0) / (float64((1.5 - (1e10 + (float32(1.5) / (3.14159 / 1.5))))) + (((0.0 + 0.0) * 0.1) - 3.14159)))
    var v14 uint64 = 42
    var v15 float64 = (1.5 * (((v13 + float64((3.14159 * ((1.5 + 1e10) + 1e10)))) + (v13 / v13)) + v13))
    v4 = v4
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v12
    _ = v13
    _ = v14
    _ = v15
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
