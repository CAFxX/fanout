package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 int16 }

func main() {
    var v1 uint8 = 2
    v2 := make(map[uint32]uint32)
    var v3 int = 1
    v4 := make(map[int]bool)
    var v5 float64 = (0.0 - 0.1)
    var v6 uint32 = 2
    var v7 int64 = 204
    fmt.Println(v1, v6, v5)
    switch ("" + "a") {
    case "hello":
        v5 = (0.1 + 3.14159)
    case "foo bar":
        v8 := []int{v3, int(v1)}
        _ = v8
    }
    v5 -= float64((0.0 * 0.0))
    var v9 int16 = -1
    var v10 uint8 = (^(v1))
    switch "" {
    case "":
        v10 = 42
    case "hello":
        fmt.Println(v3, v3, v10)
    }
    switch "a" {
    case "hello":
        for fi11 := 0; fi11 < 2; fi11++ {
            var v12 int8 = 127
            _ = fi11
            _ = v12
        }
    default:
        switch "a" {
        case "a":
            v7 = 0
        case "foo bar":
            v10 = (v1 >> ((42) & 63))
        }
    }
    v13 := []int{(^(v3)), 635, (v3 % ((v3) | 1)), v3}
    v14 := GenS1{F0: 2147483647, F1: -701}
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("slicelen v13:", len(v13), v13)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v13
    _ = v14
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v9
}
