package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
type GenS2 struct { F0 uint32 }

func main() {
    var v1 uint16 = 0
    v2 := []uint32{2147483648, 0, 970}
    var v3 uint64 = 1
    var v4 bool = true
    for fi5 := 0; fi5 < 15; fi5++ {
        v4 = v4
        _ = fi5
    }
    v3 = (v3 << ((279) & 63))
    for fi6 := 0; fi6 < 11; fi6++ {
        var v7 float32 = float32(1e10)
        _ = fi6
        _ = v7
    }
    v8 := GenS1{F0: 0}
    var v9 uint = 4294967295
    v10 := GenS2{F0: 347}
    var v11 uint = (v9 / (((v9 >> ((v9) & 63))) | 1))
    v2 = append(v2, 65535)
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v2
    _ = v3
    _ = v4
    _ = v8
    _ = v9
}
