package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 uint32 }

func main() {
    v1 := []uint32{}
    v2 := make(map[uint32]uint32)
    var v3 string = "foo bar"
    fmt.Println(v3)
    v4 := GenS1{F0: true, F1: 1}
    v5 := GenS1{F0: (0 != 4294967295), F1: 440}
    v6 := GenS1{F0: true, F1: 4294967295}
    v1 = append(v1, 2)
    v7 := GenS1{F0: true, F1: 4294967295}
    if len(v1) > 0 {
        v1[0] = 2147483647
    }
    switch v3 {
    case "":
        fmt.Println(v3)
    case "a":
        fmt.Println(v3)
    case "foo bar":
        v3 = (v3 + v3)
    }
    var v8 uint32 = 2147483647
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
