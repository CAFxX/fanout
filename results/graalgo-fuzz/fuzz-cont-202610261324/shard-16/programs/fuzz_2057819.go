package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 uint32; F2 int64 }
func genFn0() int {
    v4 := []int{32767, 931, 32768}
    var v5 uint = 18446744073709551615
    _ = v4
    _ = v5
    return 1
}
func genFn1(p0 int8, p1 uint32) int64 {
    v6 := GenS1{F0: 0, F1: p1, F2: 1}
    if true {
        var v7 int64 = 42
        v8 := []int{genFn0(), -2147483648, -680}
        _ = v7
        _ = v8
    } else {
        p0 += p0
        p0 += p0
        p1 = p1
    }
    switch 0 {
    case -527:
        p0 -= -47
    case 128:
        var v9 float64 = 0.1
        _ = v9
    }
    _ = p0
    _ = p1
    _ = v6
    return 155
}

func main() {
    var v1 uint8 = 47
    var v2 uint = 65535
    v3 := make(map[string]bool)
    for fi10 := 0; fi10 < 7; fi10++ {
        var v11 uint32 = 4294967295
        fmt.Println(v2, v11, v11)
        v1 = (v1 << ((2) & 63))
        _ = fi10
        _ = v11
    }
    var v12 int = 0
    var v13 float32 = 1.5
    v14 := []int{}
    v15 := GenS1{F0: v12, F1: 723, F2: 42}
    v1 = 42
    fmt.Println(v12, v13)
    v1 -= 127
    fmt.Println(v12, v13, v1)
    if ((-27 > 19) && false) {
        _ = len(v3)
        var v16 uint16 = 463
        var v17 int = -1
        _ = v16
        _ = v17
    }
    var v18 int8 = 42
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("slicelen v14:", len(v14), v14)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    _ = v1
    _ = v12
    _ = v13
    _ = v14
    _ = v15
    _ = v18
    _ = v2
    _ = v3
}
