package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
type GenS2 struct { F0 int32; F1 uint16; F2 uint }

func main() {
    v1 := make(map[int]string)
    var v2 int32 = 42
    var v3 uint64 = 18446744073709551615
    v4 := GenS1{F0: 255}
    v3 -= 1
    fmt.Println(v3)
    fmt.Println(v3, v3)
    delete(v1, -235)
    v5 := v1[-2147483648]
    fmt.Println(v5, v5)
    if true {
        fmt.Println(v2, v3)
        v5 = v5
    } else {
        v6 := GenS2{F0: (v2 ^ v2), F1: 138, F2: 1}
        var v7 uint32 = 476
        for fi8 := 0; fi8 < 3; fi8++ {
            var v9 int = fi8
            _ = fi8
            _ = v9
        }
        _ = v6
        _ = v7
    }
    fmt.Println(v2, v3, v5)
    var v10 uint16 = 277
    switch 4294967295 {
    case 42:
        var v11 int32 = (v2 - v2)
        _ = v11
    case -2147483648:
        v2 = -2147483648
    case 603:
        v2 += (v2 >> ((2) & 63))
    default:
        var v12 bool = true
        _ = v12
    }
    var v13 int32 = -1
    switch 0 {
    case -93:
        fmt.Println(v5, v3, v10)
    case -85:
        v13 -= v13
    case 128:
        fmt.Println(v2, v2, v3)
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v10:", v10)
    fmt.Println("v13:", v13)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v10
    _ = v13
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
