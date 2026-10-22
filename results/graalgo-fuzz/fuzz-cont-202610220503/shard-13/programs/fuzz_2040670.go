package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 int16; F2 int8 }
type GenS2 struct { F0 int16; F1 uint16; F2 int32 }

func main() {
    var v1 uint64 = 778
    var v2 string = "\n\t"
    var v3 float64 = 0.0
    v4 := make(map[uint32]int)
    var v5 bool = (false || false)
    var v6 int64 = 1
    switch v2 {
    case "a":
        v7 := GenS1{F0: (v1 / ((0) | 1)), F1: 1000, F2: -1}
        _ = v7
    }
    var v8 int64 = 42
    switch 0 {
    case -9223372036854775808:
        if (float32(v3) > 3.14159) {
            v5 = ((1 > 2147483647) && (float64(1e10) < v3))
        }
    case -133:
        v3 -= v3
    case 2147483647:
        var v9 float64 = float64((0.0 - 1e10))
        _ = v9
    default:
        fmt.Println(v1, v6)
    }
    v10 := v4[2147483647]
    switch "" {
    case "a":
        v8 += v8
    default:
        var v11 float64 = (((v3 * v3) + v3) / (1.5 + float64(0.1)))
        _ = v11
    }
    v8 = v8
    v3 -= 3.14159
    fmt.Println(v3, v10)
    var v12 int8 = 100
    fmt.Println(v6)
    var v13 uint32 = 65535
    for fi14 := 0; fi14 < 9; fi14++ {
        for fi15 := 0; fi15 < 4; fi15++ {
            v16 := GenS1{F0: (v1 % ((585) | 1)), F1: -741, F2: (-(v12))}
            _ = fi15
            _ = v16
        }
        v17 := GenS2{F0: 32767, F1: 2, F2: -1}
        _ = fi14
        _ = v17
    }
    var v18 uint8 = 2
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v10
    _ = v12
    _ = v13
    _ = v18
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v8
}
