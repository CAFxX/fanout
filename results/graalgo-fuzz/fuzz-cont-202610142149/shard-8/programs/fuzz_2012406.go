package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 int64 }
func genFn0(p0 uint8, p1 uint) uint64 {
    v5 := []int{-1, 32768, 0, 127}
    if len(v5) >= 2 {
        v5 = v5[1:len(v5)]
    }
    p1 = p1
    switch -9223372036854775808 {
    case -598:
        v6 := GenS1{F0: 2147483648, F1: 42}
        _ = v6
    case 597:
        _ = len(v5)
    case 2147483647:
        p0 = (p0 % ((2) | 1))
    default:
        var v7 uint16 = 2
        _ = v7
    }
    _ = p0
    _ = p1
    _ = v5
    return 0
}

func main() {
    var v1 float32 = 3.14159
    v2 := make(map[string]uint32)
    v3 := []int{42}
    v4 := make(map[int]string)
    var v8 bool = (0 <= -1)
    v8 = (!true)
    fmt.Println(v8)
    if false {
        v1 = float32((-2.25 / ((0.1 / 1.5) / float64(((v1 - ((v1 - v1) / 3.14159)) - v1)))))
    } else {
        var v9 uint8 = 255
        fmt.Println(v1, v9, v8)
        var v10 int32 = 0
        _ = v10
        _ = v9
    }
    var v11 float64 = 0.1
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    _ = len(v3)
    v12 := GenS1{F0: 578, F1: 42}
    switch -1 {
    case 32768:
        var v13 int8 = -70
        _ = v13
    }
    var v14 int32 = 1
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v14:", v14)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v8:", v8)
    _ = v1
    _ = v11
    _ = v12
    _ = v14
    _ = v2
    _ = v3
    _ = v4
    _ = v8
}
