package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 uint64 }

func main() {
    var v1 uint64 = 4294967295
    var v2 uint16 = 42
    var v3 int32 = 2147483647
    for fi4 := 0; fi4 < 16; fi4++ {
        var v5 float32 = (1e10 / 1e10)
        _ = fi4
        _ = v5
    }
    v3 += -1
    fmt.Println(v3)
    v1 += (^(v1))
    v6 := []int{}
    v3 = (v3 >> ((4294967295) & 63))
    var v7 int32 = v3
    v7 += v7
    if len(v6) >= 2 {
        v6 = v6[1:len(v6)]
    }
    v8 := GenS1{F0: true, F1: 1}
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v6:", len(v6), v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v2
    _ = v3
    _ = v6
    _ = v7
    _ = v8
}
