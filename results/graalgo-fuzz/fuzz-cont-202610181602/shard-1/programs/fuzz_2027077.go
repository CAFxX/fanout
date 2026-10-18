package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }

func main() {
    var v1 bool = false
    var v2 int64 = 1
    v3 := make(map[int]string)
    var v4 uint8 = 0
    var v5 uint16 = 255
    v6 := make(map[string]uint32)
    v2 -= (v2 << ((2) & 63))
    fmt.Println(v5)
    if (18446744073709551615 >= 325) {
        switch 128 {
        case 127:
            switch 32768 {
            case 42:
                v2 = v2
            case -1:
                v1 = v1
            default:
                var v7 uint16 = (v5 << ((0) & 63))
                _ = v7
            }
        case 32767:
            v8 := GenS1{F0: v4}
            _ = v8
        case -736:
            var v9 float64 = 0.1
            _ = v9
        }
    } else {
        var v10 int16 = -32768
        var v11 int16 = -993
        v4 -= (v4 ^ v4)
        _ = v10
        _ = v11
    }
    for fi12 := 0; fi12 < 17; fi12++ {
        var v13 int64 = (v2 % ((-9223372036854775808) | 1))
        v14 := []int{}
        _ = fi12
        _ = v13
        _ = v14
    }
    if (v2 <= v2) {
        v15 := GenS1{F0: v4}
        var v16 float64 = 1.5
        _ = v15
        _ = v16
    } else {
        var v17 bool = false
        var v18 uint = 1
        var v19 int = 2147483647
        _ = v17
        _ = v18
        _ = v19
    }
    v20 := GenS1{F0: v4}
    var v21 uint = 18446744073709551615
    var v22 uint32 = 0
    v1 = false
    v23 := []int{-371, -9223372036854775808}
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("v22:", v22)
    fmt.Println("slicelen v23:", len(v23), v23)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v23
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
