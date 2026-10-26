package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 uint; F2 int }

func main() {
    v1 := make(map[uint32]bool)
    v2 := []int{1}
    var v3 string = ("hello" + "\n\t")
    var v4 uint16 = 1
    var v5 int64 = 426
    var v6 bool = ((false && (true || ((!(42 != 63)) || true))) || true)
    v6 = (255 < 0)
    v7 := GenS1{F0: -60, F1: 4294967295, F2: 2147483647}
    v4 += v4
    var v8 float64 = -2.25
    if v6 {
        v6 = (0 < 674)
        switch v3 {
        case "\n\t":
            v9 := GenS1{F0: -128, F1: 1, F2: -989}
            _ = v9
        case "a":
            fmt.Println(v6)
        case "hello":
            var v10 uint32 = 493
            _ = v10
        }
        switch 127 {
        case 325:
            fmt.Println(v6)
        }
    }
    _ = len(v2)
    v5 = (v5 << ((42) & 63))
    v5 = v5
    v11 := GenS1{F0: -1, F1: 1, F2: 441}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v11
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
