package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 int32; F2 bool }
func genFn0(p0 uint, p1 uint16, p2 uint32) uint {
    v4 := GenS1{F0: 849, F1: 42, F2: true}
    p2 += 2147483647
    p2 -= (^(p2))
    p2 -= (p2 | (^(p2)))
    _ = p0
    _ = p1
    _ = p2
    _ = v4
    return 42
}
func genFn1(p0 uint16) int16 {
    var v5 int16 = 32767
    var v6 string = ("\n\t" + "hello")
    _ = p0
    _ = v5
    _ = v6
    return 32767
}

func main() {
    var v1 int = 0
    v2 := make(map[int]uint32)
    v3 := make(map[int]int)
    v1 -= (v1 + v1)
    v7 := v2[v1]
    var v8 int64 = -1
    for fi9 := 0; fi9 < 18; fi9++ {
        v10 := []int{v1}
        v1 = fi9
        _ = fi9
        _ = v10
    }
    if false {
        switch "a" {
        case "hello":
            var v11 int8 = 67
            _ = v11
        case "\n\t":
            var v12 uint64 = 1
            _ = v12
        }
        var v13 uint64 = 4294967295
        switch "a" {
        case "hello":
            fmt.Println(v1)
        case "\n\t":
            v7 -= (^(v7))
        case "foo bar":
            var v14 uint32 = 761
            _ = v14
        }
        _ = v13
    } else {
        var v15 uint = genFn0(4294967295, 1, (v7 / ((v7) | 1)))
        var v16 bool = true
        var v17 float64 = 0.1
        _ = v15
        _ = v16
        _ = v17
    }
    var v18 uint8 = 139
    v19 := []int{(v1 / (((v1 & v1)) | 1)), v1}
    v1 += v1
    fmt.Println(v1)
    v18 = 1
    v8 += (v8 << ((228) & 63))
    v2[(-(v1))] = (v7 - v7)
    var v20 uint16 = 42
    for fi21 := 0; fi21 < 7; fi21++ {
        var v22 uint64 = 48
        if len(v19) > 0 {
            v19[0] = (v1 * fi21)
        }
        _ = fi21
        _ = v22
    }
    v23 := GenS1{F0: 1, F1: -1, F2: (42 > 32767)}
    fmt.Println("v1:", v1)
    fmt.Println("v18:", v18)
    fmt.Println("slicelen v19:", len(v19), v19)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v20:", v20)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v23
    _ = v3
    _ = v7
    _ = v8
}
