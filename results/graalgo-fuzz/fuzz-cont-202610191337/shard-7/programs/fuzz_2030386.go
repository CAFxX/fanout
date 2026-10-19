package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 int8 }
func genFn0(p0 int32) int32 {
    switch -9223372036854775808 {
    case 1:
        v4 := []int{9223372036854775807, 32767}
        _ = v4
    case 128:
        p0 -= (p0 % (((p0 / (((p0 | (p0 + 411))) | 1))) | 1))
    case -50:
        p0 -= p0
    }
    v5 := []int{-607}
    _ = p0
    _ = v5
    return p0
}
func genFn1(p0 int64, p1 string, p2 uint64) int8 {
    v6 := []int{1}
    p0 = p0
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    return 100
}

func main() {
    v1 := make(map[uint32]bool)
    var v2 float32 = 0.1
    var v3 uint64 = 0
    v2 += 0.1
    var v7 uint8 = uint8(v3)
    for fi8 := 0; fi8 < 7; fi8++ {
        v2 -= v2
        _ = fi8
    }
    v2 -= 0.0
    var v9 int = 32767
    _ = len(v1)
    for fi10 := 0; fi10 < 11; fi10++ {
        v7 = v7
        delete(v1, 2147483648)
        var v11 int16 = 1
        _ = fi10
        _ = v11
    }
    v12 := GenS1{F0: -2147483648, F1: 0}
    var v13 int32 = 0
    v2 = v2
    var v14 int32 = 1
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v7:", v7)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v12
    _ = v13
    _ = v14
    _ = v2
    _ = v3
    _ = v7
    _ = v9
}
