package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
type GenS2 struct { F0 string; F1 uint8; F2 uint }
type GenS3 struct { F0 int16 }

func main() {
    var v1 uint64 = 18446744073709551615
    var v2 int32 = 42
    v3 := []uint32{2, 0}
    v4 := make(map[string]bool)
    v5 := []int{201}
    fmt.Println(v1)
    v6 := GenS1{F0: uint(v1)}
    v1 = 0
    v2 = v2
    fmt.Println(v2, v2)
    var v7 float32 = 0.0
    var v8 int16 = -32768
    for fi9 := 0; fi9 < 4; fi9++ {
        var v10 uint8 = 128
        var v11 uint16 = 2
        _ = fi9
        _ = v10
        _ = v11
    }
    v12 := GenS2{F0: ("a" + "hello"), F1: 128, F2: 0}
    var v13 int = -762
    if len(v3) > 0 {
        v3[0] = 1
    }
    v13 = (v13 >> ((1) & 63))
    if (!true) {
        var v14 int64 = -9223372036854775808
        fmt.Println(v1, v7)
        v15 := GenS3{F0: v8}
        _ = v14
        _ = v15
    }
    for fi16 := 0; fi16 < 6; fi16++ {
        v7 = (-2.25 - float32(float64(v7)))
        _ = len(v3)
        _ = fi16
    }
    for fi17 := 0; fi17 < 2; fi17++ {
        v1 = v1
        _ = fi17
    }
    switch "foo bar" {
    case "\n\t":
        var v18 float64 = 0.1
        _ = v18
    }
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v12
    _ = v13
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
