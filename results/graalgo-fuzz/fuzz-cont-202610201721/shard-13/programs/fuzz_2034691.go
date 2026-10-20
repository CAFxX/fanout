package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 int64; F1 uint64; F2 int32 }
type GenS3 struct { F0 int32; F1 bool }
type GenS4 struct { F0 uint32; F1 int64; F2 uint }
func genFn0(p0 bool, p1 int32, p2 uint8) uint64 {
    fmt.Println(p0)
    if (p1 < (p1 | (p1 / ((p1) | 1)))) {
        fmt.Println(p2, p0)
        v4 := GenS1{F0: (p2 - 127)}
        _ = v4
    }
    _ = p0
    _ = p1
    _ = p2
    return 982
}
func genFn1(p0 int, p1 uint32) uint16 {
    var v5 uint64 = 743
    var v6 uint = 1
    _ = p0
    _ = p1
    _ = v5
    _ = v6
    return 1
}

func main() {
    var v1 float64 = 1e10
    var v2 uint = 2
    v3 := []string{}
    fmt.Println(v1, v2)
    fmt.Println(v2)
    v2 = (v2 * (v2 ^ 18446744073709551615))
    if (0 == 42) {
        v1 += 1.5
        var v7 int16 = -32768
        v8 := GenS2{F0: -9223372036854775808, F1: 4294967295, F2: 0}
        _ = v7
        _ = v8
    }
    v9 := GenS3{F0: 908, F1: false}
    v1 = v1
    for fi10 := 0; fi10 < 16; fi10++ {
        v2 += 337
        v1 = 1e10
        v1 = float64(1e10)
        _ = fi10
    }
    var v11 float32 = ((0.1 + -2.25) / 1.0)
    v12 := GenS4{F0: 1, F1: 9223372036854775807, F2: v2}
    v3 = append(v3, ("a" + ("hello" + "a")))
    var v13 float32 = v11
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v13:", v13)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    _ = v1
    _ = v11
    _ = v12
    _ = v13
    _ = v2
    _ = v3
    _ = v9
}
