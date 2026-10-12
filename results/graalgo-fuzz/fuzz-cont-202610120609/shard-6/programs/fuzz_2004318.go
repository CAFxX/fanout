package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 uint64 }
func genFn0(p0 int8) uint64 {
    p0 += (p0 % (((p0 >> ((65535) & 63))) | 1))
    v6 := []int{733, -304, 32768, -593}
    p0 = 42
    _ = p0
    _ = v6
    return 18446744073709551615
}

func main() {
    var v1 uint = 311
    var v2 float64 = float64(1.5)
    v3 := make(map[string]int)
    var v4 int32 = 1
    v5 := make(map[string]uint32)
    var v7 uint8 = 255
    for fi8 := 0; fi8 < 1; fi8++ {
        v2 -= v2
        _ = fi8
    }
    v1 = (v1 - v1)
    switch -340 {
    case 2147483647:
        v1 = (v1 * v1)
    case 0:
        v9 := []int{103, -2147483648, 4294967295, -993}
        _ = v9
    default:
        fmt.Println(v7)
    }
    var v10 int64 = -783
    v11 := GenS1{F0: -1, F1: 0}
    v10 = v10
    switch 32767 {
    case 128:
        fmt.Println(v10, v10)
    case 91:
        v2 = v2
    default:
        v4 -= (v4 - 0)
    }
    var v12 uint8 = (^(v7))
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v12:", v12)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v7:", v7)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v7
}
