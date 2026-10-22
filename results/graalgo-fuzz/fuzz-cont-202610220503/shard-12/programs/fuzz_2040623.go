package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 uint64 }
type GenS2 struct { F0 uint; F1 int32 }
type GenS3 struct { F0 uint16; F1 uint }
type GenS4 struct { F0 bool }
func genFn0(p0 uint64, p1 int8, p2 uint32) uint {
    v6 := GenS1{F0: -2147483648, F1: p0}
    v7 := GenS1{F0: -1, F1: p0}
    var v8 uint16 = 1
    v9 := GenS1{F0: 1, F1: 4294967295}
    fmt.Println(p1, p1, p2)
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    _ = v7
    _ = v8
    _ = v9
    return 32
}

func main() {
    var v1 int8 = 1
    var v2 uint16 = 255
    var v3 bool = (65535 >= 534)
    var v4 int = 1
    var v5 uint = 255
    v4 = v4
    switch "foo bar" {
    case "hello":
        v4 = (v4 >> ((v5) & 63))
    default:
        v10 := GenS2{F0: (v5 % (((v5 << (((v5 + (v5 * (^(v5))))) & 63))) | 1)), F1: -109}
        _ = v10
    }
    var v11 float64 = 1.5
    var v12 float32 = 1.5
    if (65535 < uint32(v5)) {
        fmt.Println(v3)
    }
    var v13 int64 = -1
    if v3 {
        if (!false) {
            for fi14 := 0; fi14 < 11; fi14++ {
                v15 := GenS3{F0: v2, F1: (v5 & v5)}
                _ = fi14
                _ = v15
            }
        }
        var v16 bool = (v2 > v2)
        _ = v16
    }
    v11 += 1.5
    fmt.Println(v4, v12, v12)
    var v17 uint64 = 0
    fmt.Println(v13, v17)
    v1 = int8(v2)
    fmt.Println(v2)
    var v18 int = (v4 ^ v4)
    if (42 > int8(v17)) {
        v19 := GenS3{F0: v2, F1: (v5 % (((^(v5))) | 1))}
        if (!false) {
            var v20 uint16 = v2
            v21 := GenS4{F0: v3}
            _ = v20
            _ = v21
        }
        _ = v19
    }
    fmt.Println(v5)
    v22 := []int{}
    var v23 float64 = 1.5
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v22:", len(v22), v22)
    fmt.Println("v23:", v23)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v11
    _ = v12
    _ = v13
    _ = v17
    _ = v18
    _ = v2
    _ = v22
    _ = v23
    _ = v3
    _ = v4
    _ = v5
}
