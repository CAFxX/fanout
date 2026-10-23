package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 int8 }

func main() {
    var v1 float32 = 1e10
    var v2 int = 9223372036854775807
    v3 := make(map[uint32]int)
    v4 := make(map[int]string)
    v5 := []uint32{}
    _ = len(v5)
    v1 = 0.0
    var v6 int16 = -586
    v7 := GenS1{F0: -289, F1: 42}
    var v8 int64 = -1
    v2 -= (v2 / ((v2) | 1))
    _ = len(v5)
    v8 -= 798
    fmt.Println(v6, v8, v2)
    if (0.1 >= v1) {
        switch "hello" {
        case "\n\t":
            var v9 float64 = (1.5 / (-2.25 * 1e10))
            _ = v9
        case "foo bar":
            var v10 uint16 = 309
            _ = v10
        }
    }
    var v11 uint = 716
    var v12 uint64 = 0
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v11
    _ = v12
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
