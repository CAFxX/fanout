package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 uint8; F2 uint }

func main() {
    v1 := make(map[uint32]string)
    var v2 int32 = 1
    v3 := []string{""}
    v4 := []uint32{}
    v5 := GenS1{F0: -1, F1: 128, F2: 1}
    var v6 uint64 = 1
    delete(v1, 627)
    var v7 uint8 = 179
    fmt.Println(v2)
    v2 = (v2 >> ((65535) & 63))
    var v8 float64 = 0.1
    for fi9 := 0; fi9 < 16; fi9++ {
        fmt.Println(fi9, v6, fi9)
        v6 = 4294967295
        _ = fi9
    }
    _ = len(v4)
    var v10 int64 = -9223372036854775808
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v10:", v10)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v10
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
