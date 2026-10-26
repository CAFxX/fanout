package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
type GenS2 struct { F0 uint8; F1 uint }
type GenS3 struct { F0 int; F1 string; F2 uint64 }
func genFn0() uint {
    var v5 float32 = 0.0
    var v6 float64 = float64(1e10)
    fmt.Println(v5, v6)
    _ = v5
    _ = v6
    return 42
}
func genFn1() uint32 {
    var v7 uint = 693
    v7 = v7
    var v8 int = -197
    if false {
        var v9 uint16 = 1
        switch (v8 + 127) {
        case -2147483648:
            v9 = 1
        }
        _ = v9
    } else {
        v7 = 42
        fmt.Println(v7, v8)
    }
    _ = v7
    _ = v8
    return 338
}

func main() {
    var v1 int16 = 468
    v2 := make(map[uint32]bool)
    v3 := []int{}
    var v4 uint16 = 32767
    var v10 bool = true
    v11 := GenS1{F0: -1}
    switch -1 {
    case -979:
        v10 = v10
    default:
        v4 = (v4 * (v4 << ((genFn0()) & 63)))
    }
    var v12 int32 = 2147483647
    fmt.Println(v10, v10, v4)
    var v13 int8 = -20
    for fi14 := 0; fi14 < 6; fi14++ {
        v15 := GenS2{F0: 127, F1: 1}
        v16 := GenS3{F0: 0, F1: "a", F2: 728}
        _ = fi14
        _ = v15
        _ = v16
    }
    _ = len(v2)
    for fi17 := 0; fi17 < 10; fi17++ {
        if v10 {
            fmt.Println(v12, fi17, v4)
            v3 = append(v3, (fi17 % (((fi17 ^ fi17)) | 1)))
        }
        _ = fi17
    }
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v13
    _ = v2
    _ = v3
    _ = v4
}
