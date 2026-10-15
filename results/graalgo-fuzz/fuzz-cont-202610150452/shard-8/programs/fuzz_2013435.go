package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 uint8 }

func main() {
    var v1 int64 = 408
    v2 := make(map[uint32]bool)
    v3 := make(map[uint32]int)
    var v4 float64 = (float64(((3.14159 / 1.0) * (((3.14159 + (0.0 / (3.14159 - (((float32((0.0 + (float64(0.0) + 1.5))) - (-2.25 + (0.0 * float32(((0.1 * float64(((1.5 / 1.0) + (((0.0 - 1e10) / (3.14159 - (3.14159 * 1.5))) - 1e10)))) + 1.5))))) + 0.1) * 3.14159)))) / -2.25) - 3.14159))) - ((0.1 - 3.14159) * 1.5))
    v5 := GenS1{F0: 4294967295, F1: 2}
    for fi6 := 0; fi6 < 5; fi6++ {
        for fi7 := 0; fi7 < 8; fi7++ {
            var v8 uint16 = 42
            _ = len(v3)
            _ = fi7
            _ = v8
        }
        _ = fi6
    }
    var v9 int64 = -1
    switch -5 {
    case 9223372036854775807:
        v10 := GenS1{F0: 42, F1: 128}
        _ = v10
    default:
        for fi11 := 0; fi11 < 17; fi11++ {
            if (-755 != -48) {
                v4 -= v4
                var v12 int16 = -208
                _ = v12
            } else {
                var v13 int = -241
                _ = v13
            }
            var v14 int8 = 42
            fmt.Println(v4)
            _ = fi11
            _ = v14
        }
    }
    v9 += (v9 & (v9 + v9))
    var v15 bool = true
    _ = len(v3)
    var v16 int = -94
    var v17 int = 42
    fmt.Println("v1:", v1)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v15
    _ = v16
    _ = v17
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v9
}
