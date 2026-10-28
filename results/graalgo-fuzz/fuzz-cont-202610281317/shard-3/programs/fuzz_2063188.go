package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 int32; F1 uint64; F2 int }
type GenS3 struct { F0 uint32; F1 uint64 }
func genFn0(p0 uint8) uint64 {
    p0 += (p0 * (p0 / (((p0 + (p0 / (((^(p0))) | 1)))) | 1)))
    fmt.Println(p0)
    if (!false) {
        fmt.Println(p0)
    }
    v5 := []int{-976, -1, -911}
    _ = p0
    _ = v5
    return 42
}
func genFn1(p0 uint8, p1 int16, p2 uint8) uint16 {
    p1 = -32768
    p2 = p2
    p0 = 127
    _ = p0
    _ = p1
    _ = p2
    return 1
}

func main() {
    var v1 int16 = -1
    v2 := make(map[string]uint32)
    var v3 int = -322
    v4 := []int{v3, (v3 & 48), int(v1), v3}
    var v6 int16 = v1
    v7 := GenS1{F0: 1}
    var v8 uint32 = 335
    var v9 uint = 1
    if true {
        var v10 int64 = 1
        switch "foo bar" {
        case "a":
            var v11 uint32 = 28
            _ = v11
        }
        _ = v10
    }
    var v12 int = v3
    v3 = (v12 / ((v12) | 1))
    v13 := GenS2{F0: -222, F1: 155, F2: 32767}
    fmt.Println(v9, v6, v6)
    v14 := GenS3{F0: (v8 ^ v8), F1: genFn0(237)}
    var v15 uint64 = 646
    v16 := GenS1{F0: 42}
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("v15:", v15)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v6:", v6)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v12
    _ = v13
    _ = v14
    _ = v15
    _ = v16
    _ = v2
    _ = v3
    _ = v4
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
