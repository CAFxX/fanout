package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 int32 }
type GenS2 struct { F0 uint16; F1 bool }
type GenS3 struct { F0 int32 }
func genFn0(p0 string, p1 int16) int {
    fmt.Println(p1, p0)
    v5 := []int{}
    if (!(1.5 >= 0.1)) {
        v5 = append(v5, -546)
    } else {
        v6 := GenS1{F0: 1, F1: 2147483647}
        _ = v6
    }
    switch 405 {
    case 128:
        v7 := GenS2{F0: 86, F1: true}
        _ = v7
    case -2147483648:
        var v8 int = -1
        _ = v8
    case 32767:
        var v9 uint64 = 42
        _ = v9
    }
    _ = p0
    _ = p1
    _ = v5
    return 72
}
func genFn1(p0 int32) string {
    p0 = 1
    switch "a" {
    case "foo bar":
        v10 := GenS2{F0: 130, F1: (false || true)}
        _ = v10
    case "":
        v11 := GenS2{F0: 0, F1: ((4294967295 < 18446744073709551615) || ((true || false) || (!false)))}
        _ = v11
    case "a":
        v12 := GenS3{F0: (p0 * (p0 >> ((0) & 63)))}
        _ = v12
    }
    v13 := GenS1{F0: (p0 ^ 0), F1: 122}
    p0 = (p0 - (p0 | p0))
    _ = p0
    _ = v13
    return "a"
}

func main() {
    v1 := make(map[int]uint32)
    var v2 bool = true
    var v3 float64 = 3.14159
    var v4 int64 = 42
    var v14 uint = 42
    var v15 int8 = 42
    fmt.Println(v2, v3)
    var v16 bool = (!((v15 ^ (v15 >> (((v14 % ((v14) | 1))) & 63))) > v15))
    var v17 float32 = (1.5 + 1.5)
    v4 += v4
    v18 := GenS1{F0: 2147483647, F1: 1}
    v19 := v1[2147483647]
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v14
    _ = v15
    _ = v16
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v3
    _ = v4
}
