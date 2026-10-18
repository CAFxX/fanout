package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 int8 }

func main() {
    v1 := make(map[int]string)
    var v2 int = 9223372036854775807
    v3 := []string{"hello", "foo bar", ""}
    v4 := make(map[int]uint32)
    var v5 int8 = 79
    var v6 uint = 163
    v2 = v2
    v2 -= v2
    var v7 int64 = -1
    var v8 float64 = (0.1 / 1.0)
    v6 = 780
    var v9 uint8 = 162
    v10 := GenS1{F0: (^(v6)), F1: 56}
    var v11 int = (v2 >> ((v6) & 63))
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v11:", v11)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
