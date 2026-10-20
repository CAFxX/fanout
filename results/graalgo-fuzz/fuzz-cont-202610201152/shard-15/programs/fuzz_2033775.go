package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 int64; F2 int32 }
type GenS2 struct { F0 uint8 }
type GenS3 struct { F0 uint16; F1 uint32; F2 int16 }
func genFn0(p0 uint) bool {
    v6 := []int{-1, 1, -1, -110}
    p0 = p0
    if len(v6) >= 2 {
        v6 = v6[1:len(v6)]
    }
    v7 := GenS1{F0: (p0 & p0), F1: -9223372036854775808, F2: 2147483647}
    var v8 uint = 42
    _ = p0
    _ = v6
    _ = v7
    _ = v8
    return (!true)
}
func genFn1(p0 int, p1 uint32) uint16 {
    p0 += -944
    var v9 int32 = 0
    var v10 int64 = -78
    switch p0 {
    case 2147483647:
        switch (p0 >> ((144) & 63)) {
        case 32768:
            var v11 int16 = -454
            _ = v11
        case -1:
            fmt.Println(v9, p1)
        }
    default:
        var v12 uint = uint(p1)
        _ = v12
    }
    var v13 int32 = -633
    _ = p0
    _ = p1
    _ = v10
    _ = v13
    _ = v9
    return 42
}

func main() {
    v1 := []uint32{}
    var v2 float32 = (0.0 / 1e10)
    v3 := make(map[int]bool)
    var v4 int64 = 0
    var v5 int = 42
    var v14 uint16 = 0
    v15 := v3[(v5 & (^(v5)))]
    fmt.Println(v15, v5)
    v16 := v3[128]
    v14 = 65535
    v17 := GenS2{F0: uint8(v14)}
    v18 := GenS3{F0: v14, F1: 576, F2: 1000}
    if (42 <= 65535) {
        if false {
            v15 = v16
        } else {
            var v19 string = "hello"
            _ = v19
        }
    }
    var v20 int32 = -329
    fmt.Println(v5, v14)
    v20 = v20
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v14
    _ = v15
    _ = v16
    _ = v17
    _ = v18
    _ = v2
    _ = v20
    _ = v3
    _ = v4
    _ = v5
}
