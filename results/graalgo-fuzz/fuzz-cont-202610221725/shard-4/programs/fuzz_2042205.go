package main

import (
    "fmt"
)


func main() {
    v1 := []uint32{418, 231}
    v2 := make(map[string]int)
    var v3 int16 = 32767
    v4 := []string{}
    v5 := make(map[uint32]int)
    v3 -= 32767
    v3 = (v3 & v3)
    var v6 uint8 = 35
    var v7 uint64 = 0
    if false {
        for fi8 := 0; fi8 < 19; fi8++ {
            switch (fi8 % ((fi8) | 1)) {
            case -1:
                v7 += (v7 | 2)
            default:
                var v9 int32 = 1
                _ = v9
            }
            _ = fi8
        }
    }
    var v10 uint8 = 42
    v6 = v6
    for fi11 := 0; fi11 < 13; fi11++ {
        var v12 uint16 = 32767
        _ = len(v5)
        fmt.Println(v3, v3, v12)
        _ = fi11
        _ = v12
    }
    fmt.Println(v7)
    var v13 bool = (-865 < v3)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v10:", v10)
    fmt.Println("v13:", v13)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v10
    _ = v13
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
