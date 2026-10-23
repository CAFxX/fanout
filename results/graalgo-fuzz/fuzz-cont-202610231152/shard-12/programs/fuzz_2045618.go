package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint; F2 uint8 }
type GenS2 struct { F0 bool; F1 uint; F2 int }

func main() {
    var v1 int32 = 1
    var v2 float32 = (-2.25 / 3.14159)
    var v3 string = ("" + "foo bar")
    v4 := make(map[int]uint32)
    var v5 int16 = -1
    v6 := make(map[string]bool)
    var v7 uint64 = 492
    switch "a" {
    case "":
        v2 += v2
    case "a":
        var v8 int = 128
        _ = v8
    case "foo bar":
        var v9 uint32 = 1
        _ = v9
    default:
        v10 := GenS1{F0: 2, F1: 65535, F2: 1}
        _ = v10
    }
    v11 := []int{}
    var v12 uint64 = v7
    if true {
        var v13 uint16 = 786
        _ = v13
    }
    v14 := GenS2{F0: (false || false), F1: 729, F2: 128}
    v5 -= v5
    if false {
        fmt.Println(v12, v3)
    }
    v11 = append(v11, 32768)
    var v15 string = (v3 + v3)
    var v16 uint8 = 0
    v11 = append(v11, 127)
    fmt.Println(v7, v7, v1)
    fmt.Println(v16, v16)
    var v17 uint16 = 2
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v11:", len(v11), v11)
    fmt.Println("v12:", v12)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    fmt.Println("v7:", v7)
    _ = v1
    _ = v11
    _ = v12
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
