package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }

func main() {
    v1 := []uint32{}
    var v2 int = 127
    var v3 float64 = 0.1
    var v4 uint64 = 0
    v3 = -2.25
    fmt.Println(v3, v2)
    v3 = float64(0.1)
    v5 := GenS1{F0: v4}
    v6 := GenS1{F0: (v4 | (v4 << ((0) & 63)))}
    v7 := GenS1{F0: (v4 | v4)}
    v8 := GenS1{F0: (v4 % ((v4) | 1))}
    for fi9 := 0; fi9 < 12; fi9++ {
        switch 127 {
        case -9223372036854775808:
            fmt.Println(v2, v2, v2)
        case 32767:
            var v10 string = "hello"
            _ = v10
        case -84:
            var v11 uint = 0
            _ = v11
        default:
            v3 = v3
        }
        for fi12 := 0; fi12 < 16; fi12++ {
            var v13 uint = 393
            v13 += (^(v13))
            var v14 int16 = -57
            _ = fi12
            _ = v13
            _ = v14
        }
        fmt.Println(fi9, v3)
        _ = fi9
    }
    v3 = (v3 - ((-2.25 - (0.0 * ((v3 * 0.0) / (3.14159 * (float64(1.5) * v3))))) * 3.14159))
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
