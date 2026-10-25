package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 int8 }

func main() {
    var v1 int32 = -1
    var v2 uint64 = 578
    var v3 uint = 138
    var v4 int8 = -75
    v3 -= v3
    v5 := GenS1{F0: uint8(v2)}
    v1 += 2147483647
    v6 := []int{}
    fmt.Println(v2)
    fmt.Println(v3, v3, v3)
    var v7 int16 = 32767
    v8 := GenS1{F0: 182}
    for fi9 := 0; fi9 < 16; fi9++ {
        for fi10 := 0; fi10 < 15; fi10++ {
            fmt.Println(v1, v3)
            var v11 uint16 = 1
            _ = fi10
            _ = v11
        }
        v3 -= v3
        v12 := GenS2{F0: (v4 / ((v4) | 1))}
        _ = fi9
        _ = v12
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v6:", len(v6), v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
