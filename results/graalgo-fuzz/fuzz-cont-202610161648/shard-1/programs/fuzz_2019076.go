package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 int16; F2 string }

func main() {
    v1 := []int{32768, -476, 371}
    v2 := make(map[int]uint32)
    v3 := make(map[string]uint32)
    v4 := make(map[uint32]bool)
    var v5 uint16 = 75
    v5 = 32767
    var v6 uint64 = 1
    fmt.Println(v6)
    var v7 int16 = 1000
    for fi8 := 0; fi8 < 9; fi8++ {
        v6 += v6
        _ = fi8
    }
    v9 := GenS1{F0: (v5 & v5), F1: (v7 - v7), F2: "hello"}
    v5 = 65535
    var v10 int32 = -861
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v10:", v10)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v10
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v9
}
