package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
type GenS2 struct { F0 bool; F1 uint64 }

func main() {
    v1 := []string{"foo bar"}
    v2 := []uint32{357, 1}
    v3 := make(map[uint32]uint32)
    var v4 float32 = float32(1e10)
    v5 := []int{-164, 127, 596}
    switch "" {
    case "foo bar":
        v4 = v4
    }
    _ = len(v5)
    var v6 int8 = -128
    v6 += v6
    var v7 float64 = 0.1
    fmt.Println(v6)
    for fi8 := 0; fi8 < 10; fi8++ {
        v6 = 127
        _ = fi8
    }
    switch "hello" {
    case "\n\t":
        v7 = v7
    default:
        var v9 bool = false
        _ = v9
    }
    v1 = append(v1, "a")
    var v10 int32 = 648
    if true {
        var v11 int = 2147483647
        var v12 float64 = -2.25
        var v13 int64 = -497
        _ = v11
        _ = v12
        _ = v13
    }
    var v14 int16 = -1
    v15 := GenS1{F0: (-(v10))}
    var v16 uint = 2
    delete(v3, uint32(v16))
    v17 := GenS2{F0: true, F1: 641}
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v10:", v10)
    fmt.Println("v14:", v14)
    fmt.Println("v16:", v16)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v10
    _ = v14
    _ = v15
    _ = v16
    _ = v17
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
