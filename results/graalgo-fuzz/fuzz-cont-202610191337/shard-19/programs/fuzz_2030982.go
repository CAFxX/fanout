package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 int32 }
type GenS2 struct { F0 int16; F1 int8 }

func main() {
    var v1 uint16 = 0
    v2 := make(map[string]uint32)
    v3 := []uint32{1}
    v4 := make(map[string]string)
    v5 := make(map[uint32]bool)
    var v6 int16 = 32767
    v7 := GenS1{F0: "a", F1: 42}
    v8 := GenS2{F0: 2, F1: -128}
    _ = len(v3)
    fmt.Println(v1)
    v6 = v6
    v3 = append(v3, 0)
    v6 = (v6 << ((186) & 63))
    v3 = append(v3, 0)
    var v9 uint32 = 26
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
