package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }
type GenS2 struct { F0 int8; F1 uint }

func main() {
    v1 := []int{}
    var v2 bool = (290 >= 42)
    v3 := make(map[int]uint32)
    v4 := make(map[int]uint32)
    v5 := []string{"", (("" + "hello") + "hello"), ""}
    var v6 uint = 255
    fmt.Println(v6, v6)
    var v7 int16 = 0
    switch "hello" {
    case "a":
        v6 -= (v6 % (((v6 | (v6 / (((v6 << (((^(v6))) & 63))) | 1)))) | 1))
    case "\n\t":
        v8 := GenS1{F0: v2}
        _ = v8
    }
    fmt.Println(v2, v7, v2)
    switch 301 {
    case 128:
        for fi9 := 0; fi9 < 6; fi9++ {
            v10 := GenS2{F0: -1, F1: (v6 ^ 65535)}
            _ = fi9
            _ = v10
        }
    }
    v11 := GenS1{F0: v2}
    fmt.Println(v6, v7, v6)
    for fi12 := 0; fi12 < 6; fi12++ {
        fmt.Println(fi12, v7, v7)
        _ = fi12
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v11
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
