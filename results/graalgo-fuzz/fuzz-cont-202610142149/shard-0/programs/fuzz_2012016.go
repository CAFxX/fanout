package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
type GenS2 struct { F0 int8; F1 uint32; F2 uint64 }
type GenS3 struct { F0 uint16 }

func main() {
    var v1 uint64 = 1
    var v2 uint8 = 1
    v3 := make(map[int]uint32)
    var v4 int64 = 42
    if (true || false) {
        v5 := GenS1{F0: 65535}
        if true {
            var v6 int32 = -1
            _ = v6
        } else {
            var v7 float64 = ((3.14159 / (1.5 / 1.0)) - 1.5)
            _ = v7
        }
        v8 := []int{-266, -9223372036854775808}
        _ = v5
        _ = v8
    }
    v2 = (v2 * (v2 / ((42) | 1)))
    v9 := []int{0}
    if (-1 >= 1) {
        v4 += v4
        fmt.Println(v1)
    } else {
        v1 = v1
        switch 611 {
        case 773:
            for fi10 := 0; fi10 < 4; fi10++ {
                v4 += 795
                _ = fi10
            }
        default:
            for fi11 := 0; fi11 < 17; fi11++ {
                var v12 bool = (true && (!true))
                v1 -= (v1 + v1)
                v12 = true
                _ = fi11
                _ = v12
            }
        }
    }
    var v13 int = 4294967295
    v14 := GenS2{F0: 112, F1: 2147483648, F2: (v1 / ((2) | 1))}
    v15 := v3[v13]
    v16 := GenS3{F0: 0}
    delete(v3, 423)
    var v17 uint8 = (^(v2))
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("v15:", v15)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v9:", len(v9), v9)
    _ = v1
    _ = v13
    _ = v14
    _ = v15
    _ = v16
    _ = v17
    _ = v2
    _ = v3
    _ = v4
    _ = v9
}
