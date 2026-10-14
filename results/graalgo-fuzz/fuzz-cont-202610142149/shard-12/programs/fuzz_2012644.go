package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 int }
type GenS2 struct { F0 uint64; F1 uint32; F2 int64 }
type GenS3 struct { F0 uint32; F1 string }
func genFn0(p0 uint) uint16 {
    p0 -= p0
    v7 := []int{}
    p0 += p0
    for fi8 := 0; fi8 < 20; fi8++ {
        v9 := GenS1{F0: -356, F1: (fi8 ^ 109)}
        v10 := GenS1{F0: 2147483647, F1: (^(fi8))}
        switch fi8 {
        case -9223372036854775808:
            var v11 uint8 = 127
            _ = v11
        case 572:
            var v12 uint8 = 0
            _ = v12
        }
        _ = fi8
        _ = v10
        _ = v9
    }
    _ = p0
    _ = v7
    return 255
}
func genFn1(p0 int64, p1 int16, p2 uint16) uint32 {
    p2 += genFn0(952)
    var v13 uint64 = 152
    _ = p0
    _ = p1
    _ = p2
    _ = v13
    return 757
}

func main() {
    v1 := []int{32768, 32768}
    var v2 uint32 = 2147483648
    v3 := []int{-57, -9223372036854775808, 127, 128}
    v4 := make(map[uint32]string)
    v5 := make(map[int]string)
    var v6 uint64 = 142
    fmt.Println(v2, v2)
    fmt.Println(v2, v6)
    v2 = (v2 & 42)
    fmt.Println(v2, v2)
    fmt.Println(v6)
    v14 := GenS2{F0: v6, F1: v2, F2: 407}
    switch -9223372036854775808 {
    case 32767:
        var v15 float32 = 0.1
        _ = v15
    case 128:
        var v16 uint = 2
        _ = v16
    default:
        v2 = (v2 * v2)
    }
    v2 -= v2
    v17 := GenS3{F0: v2, F1: "\n\t"}
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    if ((^(v6)) == (v6 | (^(v6)))) {
        v18 := GenS3{F0: (^(v2)), F1: "hello"}
        _ = v18
    }
    switch "" {
    case "\n\t":
        var v19 uint = 692
        _ = v19
    case "a":
        fmt.Println(v6)
    case "":
        v2 = 0
    default:
        fmt.Println(v6, v6)
    }
    fmt.Println(v6, v6)
    switch -2147483648 {
    case 290:
        var v20 bool = (!(42 <= 255))
        _ = v20
    case 2147483647:
        fmt.Println(v2)
    }
    var v21 int64 = -629
    var v22 int8 = -109
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("v22:", v22)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v14
    _ = v17
    _ = v2
    _ = v21
    _ = v22
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
