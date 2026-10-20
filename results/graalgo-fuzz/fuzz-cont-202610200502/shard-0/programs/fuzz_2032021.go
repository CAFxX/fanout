package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 uint }
func genFn0(p0 string, p1 uint64) int32 {
    p1 += p1
    var v6 uint8 = 128
    p1 = (p1 ^ (p1 * (p1 % ((107) | 1))))
    _ = p0
    _ = p1
    _ = v6
    return -938
}

func main() {
    v1 := make(map[uint32]bool)
    var v2 float32 = (0.0 * 0.1)
    var v3 float32 = 0.1
    var v4 uint32 = 42
    v5 := make(map[uint32]string)
    var v7 uint64 = 4294967295
    v4 = v4
    if false {
        var v8 int = 0
        var v9 int32 = 0
        _ = v8
        _ = v9
    }
    v2 += -2.25
    fmt.Println(v4)
    for fi10 := 0; fi10 < 9; fi10++ {
        v11 := GenS1{F0: true, F1: 0}
        _ = fi10
        _ = v11
    }
    v12 := GenS1{F0: (!false), F1: 4294967295}
    fmt.Println(v4, v2)
    var v13 int = 149
    v4 = (^(v4))
    v14 := []int{(v13 - v13), v13, int(v7)}
    if (false && (255 == 42)) {
        v4 = (v4 ^ v4)
    }
    v3 = v3
    v4 += v4
    var v15 uint16 = 42
    var v16 uint64 = (^(v7))
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v13:", v13)
    fmt.Println("slicelen v14:", len(v14), v14)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v7:", v7)
    _ = v1
    _ = v12
    _ = v13
    _ = v14
    _ = v15
    _ = v16
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v7
}
