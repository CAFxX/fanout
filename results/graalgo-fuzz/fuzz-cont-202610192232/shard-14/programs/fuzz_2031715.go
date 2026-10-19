package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }

func main() {
    var v1 uint16 = 59
    v2 := make(map[uint32]bool)
    v3 := []uint32{67, 65535, uint32(v1)}
    var v4 uint16 = 775
    for fi5 := 0; fi5 < 16; fi5++ {
        var v6 int64 = 1
        _ = fi5
        _ = v6
    }
    var v7 uint64 = 0
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    v8 := GenS1{F0: false}
    _ = len(v3)
    v9 := v2[2147483647]
    v10 := GenS1{F0: v9}
    v9 = v9
    fmt.Println(v7)
    _ = len(v3)
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v7:", v7)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v2
    _ = v3
    _ = v4
    _ = v7
    _ = v8
    _ = v9
}
