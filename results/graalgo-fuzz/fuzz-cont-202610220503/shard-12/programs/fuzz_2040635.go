package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
type GenS2 struct { F0 uint64; F1 uint32 }

func main() {
    v1 := make(map[string]string)
    v2 := make(map[int]bool)
    v3 := make(map[string]uint32)
    var v4 int32 = -787
    for fi5 := 0; fi5 < 5; fi5++ {
        if (!false) {
            fmt.Println(fi5, fi5)
        }
        v3[("a" + "\n\t")] = 4294967295
        _ = fi5
    }
    var v6 int = 194
    var v7 uint8 = 76
    fmt.Println(v6)
    for fi8 := 0; fi8 < 13; fi8++ {
        var v9 int64 = 33
        for fi10 := 0; fi10 < 11; fi10++ {
            if ((true || true) || false) {
                v11 := GenS1{F0: 192}
                v12 := GenS2{F0: 1, F1: 805}
                _ = v11
                _ = v12
            }
            _ = fi10
        }
        _ = fi8
        _ = v9
    }
    var v13 int64 = 678
    _ = len(v3)
    v14 := []int{-2147483648, v6, int(v13)}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v13:", v13)
    fmt.Println("slicelen v14:", len(v14), v14)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v13
    _ = v14
    _ = v2
    _ = v3
    _ = v4
    _ = v6
    _ = v7
}
