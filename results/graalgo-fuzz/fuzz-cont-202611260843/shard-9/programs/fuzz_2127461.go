package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 uint64; F2 uint64 }

func main() {
    var v1 uint16 = 634
    var v2 uint64 = 57
    var v3 int64 = 42
    var v4 uint16 = 32767
    v1 = 836
    v4 = 42
    for fi5 := 0; fi5 < 17; fi5++ {
        for fi6 := 0; fi6 < 11; fi6++ {
            v1 -= uint16(v2)
            _ = fi6
        }
        v7 := GenS1{F0: 84, F1: (v2 | v2), F2: (v2 | (^(v2)))}
        _ = fi5
        _ = v7
    }
    v1 = v4
    v3 = v3
    v4 = 592
    v1 = (v1 >> ((343) & 63))
    for fi8 := 0; fi8 < 8; fi8++ {
        fmt.Println(v2, v2, v1)
        _ = fi8
    }
    v4 = (v1 & v4)
    v9 := []int{2147483647, 1, 127, 128}
    v1 = (v1 + (v4 << ((0) & 63)))
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v9:", len(v9), v9)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v9
}
