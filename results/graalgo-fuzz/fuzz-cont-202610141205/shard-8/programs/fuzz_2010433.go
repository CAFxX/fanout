package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 uint64; F2 int64 }
type GenS2 struct { F0 uint64; F1 int64 }
type GenS3 struct { F0 uint32 }
type GenS4 struct { F0 uint }
func genFn0() uint32 {
    v4 := GenS1{F0: 2147483647, F1: 42, F2: -9223372036854775808}
    var v5 int64 = -593
    var v6 uint16 = 255
    fmt.Println(v5, v6)
    _ = v4
    _ = v5
    _ = v6
    return 2
}
func genFn1(p0 int16, p1 int32, p2 int32) int8 {
    var v7 int32 = (^(p2))
    p2 = (p1 % ((p2) | 1))
    if (276 <= 0) {
        var v8 int64 = -1
        _ = v8
    }
    v9 := []int{32767, -1}
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    _ = v9
    return 94
}

func main() {
    v1 := make(map[int]string)
    var v2 uint64 = 42
    var v3 bool = false
    for fi10 := 0; fi10 < 11; fi10++ {
        if v3 {
            v1[9223372036854775807] = ("" + "a")
        }
        _ = fi10
    }
    fmt.Println(v3)
    v2 -= (v2 << ((4294967295) & 63))
    var v11 float32 = -2.25
    var v12 uint16 = 32767
    switch "a" {
    case "hello":
        fmt.Println(v3)
    case "":
        var v13 uint = 1
        _ = v13
    case "a":
        if (1 == 430) {
            v14 := GenS1{F0: 65535, F1: (v2 ^ (v2 - 0)), F2: 0}
            _ = v14
        } else {
            switch -1 {
            case 1:
                var v15 uint16 = v12
                _ = v15
            }
        }
    default:
        switch "" {
        case "\n\t":
            for fi16 := 0; fi16 < 8; fi16++ {
                var v17 float64 = float64((float32((((3.14159 / 1.0) - (3.14159 * -2.25)) - 3.14159)) / 1.0))
                if v3 {
                    v18 := GenS2{F0: v2, F1: 42}
                    v19 := GenS3{F0: 2}
                    v20 := GenS1{F0: genFn0(), F1: (v2 * v2), F2: 42}
                    _ = v18
                    _ = v19
                    _ = v20
                } else {
                    v2 = (^(v2))
                }
                _ = fi16
                _ = v17
            }
        case "":
            v3 = (!v3)
        case "hello":
            if v3 {
                v2 += v2
            }
        default:
            v21 := GenS1{F0: 0, F1: v2, F2: 0}
            _ = v21
        }
    }
    v22 := GenS1{F0: 524, F1: (v2 >> ((1) & 63)), F2: 42}
    var v23 uint16 = v12
    v2 += (v2 << ((2) & 63))
    v24 := GenS4{F0: 4294967295}
    v11 = (v11 - v11)
    v25 := GenS4{F0: 65535}
    fmt.Println(v3, v12, v23)
    v11 += (float32(((0.1 + 0.0) * 1.5)) / (v11 * v11))
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v2:", v2)
    fmt.Println("v23:", v23)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v11
    _ = v12
    _ = v2
    _ = v22
    _ = v23
    _ = v24
    _ = v25
    _ = v3
}
