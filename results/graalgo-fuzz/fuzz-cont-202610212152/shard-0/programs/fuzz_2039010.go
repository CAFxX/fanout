package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
type GenS2 struct { F0 uint64; F1 int8 }
type GenS3 struct { F0 int32; F1 int; F2 uint32 }
func genFn0(p0 int32) uint16 {
    p0 += 2147483647
    var v5 int64 = 0
    v6 := GenS1{F0: -1}
    fmt.Println(v5, v5)
    _ = p0
    _ = v5
    _ = v6
    return 192
}
func genFn1(p0 int, p1 bool, p2 int) int16 {
    p1 = (!p1)
    switch p2 {
    case 672:
        var v7 uint64 = 4294967295
        _ = v7
    default:
        v8 := []int{9223372036854775807, (p0 | p2), (^(p2))}
        _ = v8
    }
    for fi9 := 0; fi9 < 17; fi9++ {
        p1 = (!p1)
        _ = fi9
    }
    v10 := GenS1{F0: -966}
    _ = p0
    _ = p1
    _ = p2
    _ = v10
    return -32768
}

func main() {
    var v1 uint64 = 270
    v2 := make(map[uint32]int)
    var v3 uint8 = 206
    v4 := []string{"a", ("hello" + "\n\t"), ("" + "\n\t"), ("hello" + "\n\t")}
    fmt.Println(v1, v1)
    _ = len(v4)
    v11 := GenS2{F0: 1, F1: -128}
    var v12 int8 = -128
    v13 := GenS3{F0: -1, F1: 237, F2: 2}
    fmt.Println(v12)
    fmt.Println(v3, v3, v12)
    _ = len(v2)
    var v14 uint = 780
    for fi15 := 0; fi15 < 8; fi15++ {
        v1 = 42
        v1 = v1
        var v16 uint = (^(v14))
        _ = fi15
        _ = v16
    }
    v1 -= v1
    if false {
        v14 = (v14 % (((v14 - v14)) | 1))
    } else {
        var v17 string = ""
        _ = v17
    }
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("v14:", v14)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    _ = v1
    _ = v11
    _ = v12
    _ = v13
    _ = v14
    _ = v2
    _ = v3
    _ = v4
}
