package main

import (
    "fmt"
)

type GenS1 struct { F0 int }

func main() {
    v1 := make(map[uint32]int)
    v2 := []uint32{}
    var v3 uint8 = 1
    switch 32767 {
    case 0:
        _ = len(v1)
    case 32767:
        var v4 int64 = 444
        _ = v4
    }
    fmt.Println(v3)
    for fi5 := 0; fi5 < 13; fi5++ {
        var v6 uint16 = 1
        _ = fi5
        _ = v6
    }
    v3 = v3
    var v7 int64 = 42
    fmt.Println(v7)
    var v8 int32 = 2147483647
    v7 = 0
    v9 := GenS1{F0: 826}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v2
    _ = v3
    _ = v7
    _ = v8
    _ = v9
}
