package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 int64 }

func main() {
    v1 := make(map[int]int)
    var v2 uint16 = 2
    v3 := []uint32{}
    var v4 uint = 4294967295
    v5 := make(map[uint32]uint32)
    delete(v1, -318)
    v6 := GenS1{F0: 4294967295, F1: 1}
    v7 := GenS1{F0: 2, F1: 312}
    v2 = (v2 - (v2 + 65535))
    switch ("" + "foo bar") {
    case "foo bar":
        var v8 int16 = -1
        _ = v8
    }
    for fi9 := 0; fi9 < 16; fi9++ {
        var v10 uint = v4
        _ = fi9
        _ = v10
    }
    v2 = 61
    var v11 float64 = 1e10
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v11:", v11)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v11
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
