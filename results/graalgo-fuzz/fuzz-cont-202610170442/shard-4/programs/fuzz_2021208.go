package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 int }
type GenS2 struct { F0 int32 }
type GenS3 struct { F0 int8; F1 int16; F2 uint }
func genFn0(p0 int64, p1 uint32, p2 uint) int32 {
    var v5 uint = (p2 >> ((p2) & 63))
    v6 := GenS1{F0: -9223372036854775808, F1: 650}
    v7 := []int{-2147483648}
    v7 = append(v7, 4294967295)
    _ = p0
    _ = p1
    _ = p2
    _ = v5
    _ = v6
    _ = v7
    return 42
}

func main() {
    v1 := make(map[uint32]bool)
    var v2 uint = 18446744073709551615
    var v3 float64 = 1.5
    v4 := make(map[string]int)
    var v8 int16 = -1
    var v9 string = "a"
    for fi10 := 0; fi10 < 2; fi10++ {
        v3 -= float64((3.14159 / 3.14159))
        v11 := GenS2{F0: 1}
        var v12 uint32 = 173
        _ = fi10
        _ = v11
        _ = v12
    }
    v13 := []int{127, 486, 128, 810}
    var v14 float32 = (((0.0 + -2.25) + -2.25) + 1e10)
    var v15 uint64 = 598
    _ = len(v13)
    for fi16 := 0; fi16 < 8; fi16++ {
        var v17 int = fi16
        _ = fi16
        _ = v17
    }
    if (-9223372036854775808 != -1) {
        fmt.Println(v15, v3, v8)
        v18 := GenS1{F0: 106, F1: 128}
        _ = v18
    }
    var v19 float64 = 0.0
    if true {
        v20 := GenS3{F0: 127, F1: v8, F2: (v2 / (((^(v2))) | 1))}
        _ = v20
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v13:", len(v13), v13)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v13
    _ = v14
    _ = v15
    _ = v19
    _ = v2
    _ = v3
    _ = v4
    _ = v8
    _ = v9
}
