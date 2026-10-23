package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 uint }
type GenS2 struct { F0 int16 }
type GenS3 struct { F0 uint; F1 uint8; F2 uint }
func genFn0(p0 uint, p1 uint32) int {
    var v4 int64 = -698
    v5 := []int{-9223372036854775808}
    _ = p0
    _ = p1
    _ = v4
    _ = v5
    return 0
}
func genFn1(p0 uint16) uint32 {
    var v6 int = 2147483647
    v7 := GenS1{F0: 8, F1: 29}
    _ = p0
    _ = v6
    _ = v7
    return 65535
}

func main() {
    var v1 bool = true
    var v2 int = 2147483647
    v3 := make(map[uint32]bool)
    v8 := GenS2{F0: -32768}
    v1 = true
    v9 := []int{}
    if (!v1) {
        v10 := v3[1]
        var v11 int32 = 42
        switch "\n\t" {
        case "\n\t":
            _ = len(v3)
        }
        _ = v10
        _ = v11
    }
    for fi12 := 0; fi12 < 5; fi12++ {
        var v13 uint32 = 4294967295
        if len(v9) > 0 {
            v9[0] = fi12
        }
        _ = fi12
        _ = v13
    }
    var v14 uint8 = 44
    v15 := GenS3{F0: 4294967295, F1: 1, F2: 245}
    switch (v2 + (v2 | v2)) {
    case 4294967295:
        var v16 int64 = -16
        _ = v16
    case 127:
        fmt.Println(v2, v1, v14)
    default:
        var v17 int32 = 109
        _ = v17
    }
    var v18 string = ""
    fmt.Println(v14, v2)
    fmt.Println("v1:", v1)
    fmt.Println("v14:", v14)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v9:", len(v9), v9)
    _ = v1
    _ = v14
    _ = v15
    _ = v18
    _ = v2
    _ = v3
    _ = v8
    _ = v9
}
