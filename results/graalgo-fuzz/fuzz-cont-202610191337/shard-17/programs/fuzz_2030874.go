package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
type GenS2 struct { F0 uint64; F1 uint32; F2 uint16 }
type GenS3 struct { F0 int; F1 uint32 }
type GenS4 struct { F0 uint32 }
type GenS5 struct { F0 uint; F1 int32 }
type GenS6 struct { F0 int32; F1 int64 }
func genFn0() bool {
    var v7 float64 = float64(-2.25)
    if (!false) {
        fmt.Println(v7)
    } else {
        v8 := []int{-9223372036854775808}
        v7 += 3.14159
        v8 = append(v8, -1)
        _ = v8
    }
    var v9 int32 = 456
    var v10 int = 127
    v7 -= float64(3.14159)
    _ = v10
    _ = v7
    _ = v9
    return true
}

func main() {
    v1 := make(map[int]string)
    v2 := []uint32{0, 365, 2147483647, 720}
    var v3 uint8 = 255
    v4 := make(map[uint32]string)
    var v5 uint = 18446744073709551615
    v6 := make(map[int]string)
    v3 -= v3
    if false {
        v5 += (^(v5))
        v11 := GenS1{F0: 1}
        _ = v11
    } else {
        if false {
            for fi12 := 0; fi12 < 8; fi12++ {
                var v13 float64 = 0.1
                fmt.Println(v5)
                _ = fi12
                _ = v13
            }
            if (1.5 != 0.0) {
                var v14 int = 32768
                var v15 int = (v14 * v14)
                v14 = 4294967295
                _ = v14
                _ = v15
            }
        }
    }
    for fi16 := 0; fi16 < 18; fi16++ {
        switch "a" {
        case "hello":
            v5 -= 1
        case "a":
            fmt.Println(v3)
        case "foo bar":
            v2 = append(v2, 2)
        default:
            fmt.Println(v5, fi16)
        }
        _ = fi16
    }
    v3 = (v3 ^ (v3 >> ((v5) & 63)))
    var v17 uint32 = 2147483647
    v18 := GenS2{F0: 0, F1: 94, F2: uint16(v17)}
    if false {
        v19 := GenS2{F0: 1, F1: v17, F2: 42}
        _ = v19
    }
    v20 := GenS3{F0: -2147483648, F1: (v17 | (v17 + v17))}
    v21 := GenS4{F0: 1}
    switch "a" {
    case "":
        v22 := GenS5{F0: uint(v3), F1: 63}
        _ = v22
    default:
        if (710 <= 4294967295) {
            fmt.Println(v3, v3)
            var v23 int8 = -128
            _ = v23
        } else {
            v5 += (v5 + 0)
            v3 += (v3 - (v3 ^ (^(v3))))
            v24 := GenS6{F0: 0, F1: 42}
            _ = v24
        }
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v17:", v17)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v17
    _ = v18
    _ = v2
    _ = v20
    _ = v21
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
