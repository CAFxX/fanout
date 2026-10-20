package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 string }
type GenS2 struct { F0 int16; F1 uint64 }

func main() {
    var v1 int8 = 0
    v2 := make(map[int]uint32)
    v3 := []uint32{2147483648}
    var v4 bool = false
    _ = len(v2)
    switch -2147483648 {
    case -1:
        v3 = append(v3, 453)
    }
    switch ("hello" + ("\n\t" + "")) {
    case "hello":
        v4 = (190 >= 229)
    default:
        v5 := GenS1{F0: 42, F1: ("" + "")}
        _ = v5
    }
    var v6 uint = 95
    var v7 uint8 = 2
    v1 = -65
    v8 := GenS2{F0: -1, F1: 4294967295}
    v9 := GenS1{F0: 4294967295, F1: "hello"}
    var v10 int64 = 1
    v10 -= -70
    var v11 uint16 = 411
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v10
    _ = v11
    _ = v2
    _ = v3
    _ = v4
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
