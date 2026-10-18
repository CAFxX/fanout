package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
type GenS2 struct { F0 uint8; F1 uint; F2 uint }
type GenS3 struct { F0 uint64; F1 int; F2 int8 }
type GenS4 struct { F0 uint8; F1 uint32; F2 int }
type GenS5 struct { F0 string; F1 int64; F2 uint8 }
func genFn0(p0 uint32, p1 int64) int {
    p1 = p1
    var v4 uint16 = 255
    _ = p0
    _ = p1
    _ = v4
    return -9223372036854775808
}
func genFn1(p0 uint64, p1 int32, p2 int32) uint {
    v5 := GenS1{F0: 2}
    var v6 int32 = p1
    _ = p0
    _ = p1
    _ = p2
    _ = v5
    _ = v6
    return 726
}

func main() {
    var v1 uint32 = 42
    v2 := make(map[uint32]string)
    var v3 int32 = 2147483647
    v7 := []int{genFn0(2147483648, 0), -9223372036854775808, genFn0(0, -356), 399}
    for fi8 := 0; fi8 < 8; fi8++ {
        var v9 int64 = 1
        _ = fi8
        _ = v9
    }
    var v10 uint = 517
    var v11 uint32 = v1
    v12 := GenS2{F0: 168, F1: (^(v10)), F2: 42}
    v13 := GenS1{F0: (^(v1))}
    switch "\n\t" {
    case "a":
        v14 := GenS3{F0: 655, F1: 32767, F2: -128}
        _ = v14
    case "":
        v3 = v3
    case "hello":
        if (((v10 + 2) < (v10 & 42)) && (!(42 >= (v1 - 2147483648)))) {
            if len(v7) >= 2 {
                v7 = v7[1:len(v7)]
            }
            switch 127 {
            case -2147483648:
                v15 := GenS4{F0: 137, F1: (v1 << (((v10 - genFn1(4294967295, (^(v3)), 278))) & 63)), F2: genFn0(0, 42)}
                _ = v15
            case 32768:
                fmt.Println(v11, v11)
            case 1:
                var v16 float32 = ((3.14159 / float32(-2.25)) + (float32(0.1) - (-2.25 * (0.0 + 1.5))))
                _ = v16
            }
            v11 -= v11
        }
    default:
        v10 = v10
    }
    if false {
        if (v11 == v11) {
            if false {
                v17 := GenS3{F0: 214, F1: 1, F2: 1}
                _ = v17
            } else {
                var v18 uint16 = 32767
                var v19 int64 = 9223372036854775807
                _ = v18
                _ = v19
            }
            v2[v1] = "a"
            fmt.Println(v3)
        } else {
            switch 0 {
            case 931:
                var v20 int = 751
                _ = v20
            }
            var v21 uint32 = v11
            _ = v21
        }
        var v22 int8 = 0
        _ = v22
    } else {
        fmt.Println(v1, v3)
    }
    var v23 uint8 = uint8(v10)
    v1 = 2147483648
    if true {
        v24 := GenS2{F0: 2, F1: (v10 - (v10 >> (((v10 % ((genFn1(241, (v3 + (v3 - (v3 << ((v10) & 63)))), (v3 / ((128) | 1)))) | 1))) & 63))), F2: (v10 + (v10 ^ 255))}
        var v25 uint64 = 53
        v7 = append(v7, 128)
        _ = v24
        _ = v25
    }
    v3 -= (^(v3))
    v1 = (v1 ^ (v1 % ((v11) | 1)))
    switch "hello" {
    case "foo bar":
        fmt.Println(v11, v11, v11)
    case "a":
        fmt.Println(v11, v23, v23)
    default:
        fmt.Println(v10, v1)
    }
    v3 -= (^(v3))
    v26 := GenS5{F0: "\n\t", F1: 9223372036854775807, F2: v23}
    fmt.Println(v23, v3)
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v23:", v23)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v7:", len(v7), v7)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v13
    _ = v2
    _ = v23
    _ = v26
    _ = v3
    _ = v7
}
