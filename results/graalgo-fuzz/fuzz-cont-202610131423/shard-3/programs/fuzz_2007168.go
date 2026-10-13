package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
type GenS2 struct { F0 uint; F1 int32; F2 int64 }
type GenS3 struct { F0 uint64; F1 int32; F2 int8 }
type GenS4 struct { F0 uint8; F1 uint64; F2 uint16 }

func main() {
    var v1 string = "\n\t"
    v2 := make(map[int]uint32)
    v3 := []uint32{0, 285, 2, 848}
    var v4 float64 = 1e10
    var v5 int16 = 1
    v1 = v1
    v6 := GenS1{F0: -246}
    v4 += float64(0.1)
    v4 = v4
    if false {
        var v7 int64 = 42
        _ = v7
    } else {
        var v8 string = ((((v1 + v1) + v1) + "") + ("hello" + "hello"))
        v9 := GenS2{F0: 1, F1: -2147483648, F2: 0}
        _ = v8
        _ = v9
    }
    if (90 < 19) {
        v10 := GenS2{F0: 2, F1: 2147483647, F2: 42}
        var v11 int16 = v5
        v12 := GenS3{F0: 18446744073709551615, F1: 643, F2: 42}
        _ = v10
        _ = v11
        _ = v12
    } else {
        var v13 uint32 = 111
        var v14 uint = 42
        _ = v13
        _ = v14
    }
    var v15 bool = true
    v16 := GenS4{F0: 128, F1: 42, F2: 490}
    fmt.Println("v1:", v1)
    fmt.Println("v15:", v15)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v15
    _ = v16
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
