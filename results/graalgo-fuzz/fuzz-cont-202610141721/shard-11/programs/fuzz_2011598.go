package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 string }

func main() {
    v1 := []uint32{840, 0}
    var v2 uint = 716
    v3 := make(map[int]int)
    var v4 int32 = -172
    if (!false) {
        fmt.Println(v4)
        if true {
            v2 = 0
        }
        v4 += (v4 >> (((v2 << ((853) & 63))) & 63))
    }
    for fi5 := 0; fi5 < 14; fi5++ {
        for fi6 := 0; fi6 < 12; fi6++ {
            v2 -= 4294967295
            _ = fi6
        }
        _ = fi5
    }
    v2 = (v2 & v2)
    fmt.Println(v4)
    v4 -= v4
    v7 := GenS1{F0: 65535, F1: "a"}
    v8 := GenS1{F0: 65535, F1: "\n\t"}
    _ = len(v3)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v7
    _ = v8
}
