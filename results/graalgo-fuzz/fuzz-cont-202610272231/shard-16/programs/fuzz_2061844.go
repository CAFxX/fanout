package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 int16 }
func genFn0(p0 uint8, p1 string) int64 {
    v6 := []int{-739}
    p0 += p0
    p1 = p1
    _ = p0
    _ = p1
    _ = v6
    return 42
}

func main() {
    v1 := []uint32{}
    var v2 uint64 = 197
    v3 := make(map[int]bool)
    v4 := make(map[string]uint32)
    var v5 int = 1000
    var v7 int32 = 2147483647
    v2 = v2
    v8 := GenS1{F0: v7, F1: -32768}
    if false {
        var v9 uint16 = uint16(v2)
        var v10 int32 = (v7 % ((v7) | 1))
        var v11 int64 = -903
        _ = v10
        _ = v11
        _ = v9
    } else {
        v12 := GenS1{F0: (v7 - v7), F1: 203}
        for fi13 := 0; fi13 < 8; fi13++ {
            var v14 int64 = -1
            _ = fi13
            _ = v14
        }
        _ = v12
    }
    var v15 int8 = 1
    v2 -= 4294967295
    delete(v3, -2147483648)
    fmt.Println(v5)
    v2 = 1
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v15
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v7
    _ = v8
}
