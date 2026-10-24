package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 int16 }
type GenS2 struct { F0 int16; F1 int32 }
type GenS3 struct { F0 int64 }
type GenS4 struct { F0 uint16; F1 uint32 }
func genFn0() uint32 {
    var v5 float64 = ((3.14159 * 3.14159) + (0.0 / 0.1))
    if (127 > 11) {
        var v6 int32 = -319
        var v7 int64 = 9223372036854775807
        _ = v6
        _ = v7
    } else {
        v8 := GenS1{F0: "\n\t", F1: 32767}
        v5 -= (0.1 * v5)
        fmt.Println(v5)
        _ = v8
    }
    _ = v5
    return 0
}

func main() {
    v1 := make(map[int]int)
    v2 := make(map[int]bool)
    var v3 int64 = 42
    var v4 bool = false
    _ = len(v2)
    var v9 uint64 = 2
    switch 1 {
    case 32768:
        var v10 uint64 = (v9 - v9)
        _ = v10
    default:
        if (255 <= 127) {
            v11 := GenS2{F0: 1000, F1: 994}
            _ = v11
        }
    }
    v12 := GenS3{F0: (^(v3))}
    _ = len(v1)
    var v13 int32 = 1
    fmt.Println(v3)
    var v14 int16 = -25
    v3 = v3
    fmt.Println(v4, v13)
    v15 := GenS4{F0: 32767, F1: 42}
    v14 = v14
    v4 = v4
    v16 := GenS3{F0: (-(v3))}
    v13 += (v13 + v13)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v12
    _ = v13
    _ = v14
    _ = v15
    _ = v16
    _ = v2
    _ = v3
    _ = v4
    _ = v9
}
