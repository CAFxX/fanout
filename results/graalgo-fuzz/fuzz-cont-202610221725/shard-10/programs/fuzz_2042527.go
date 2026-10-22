package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 uint64; F2 int16 }

func main() {
    var v1 uint8 = 70
    v2 := []uint32{}
    var v3 int8 = -1
    var v4 float32 = (1.5 / 1.0)
    v5 := make(map[int]uint32)
    var v6 float32 = 0.1
    switch ("hello" + (("foo bar" + "hello") + "a")) {
    case "":
        v4 += float32(0.1)
    case "a":
        var v7 int8 = (v3 << ((2) & 63))
        _ = v7
    case "foo bar":
        v1 += v1
    }
    v8 := GenS1{F0: 1, F1: 2, F2: -593}
    v2 = append(v2, 2)
    v2 = append(v2, 4294967295)
    v1 = v1
    switch "a" {
    case "a":
        v3 += v3
    }
    var v9 uint8 = 127
    var v10 uint = 2
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v8
    _ = v9
}
