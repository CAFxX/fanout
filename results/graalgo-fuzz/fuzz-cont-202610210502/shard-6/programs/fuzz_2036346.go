package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 uint; F2 bool }

func main() {
    var v1 uint32 = 4294967295
    v2 := make(map[uint32]bool)
    var v3 uint8 = 1
    var v4 bool = true
    v3 -= (v3 / ((128) | 1))
    fmt.Println(v1)
    for fi5 := 0; fi5 < 5; fi5++ {
        v6 := []int{-395, (fi5 << ((255) & 63)), fi5}
        _ = fi5
        _ = v6
    }
    v7 := GenS1{F0: v3, F1: 2, F2: (-1 > 0)}
    switch "a" {
    case "\n\t":
        v2[v1] = true
    case "hello":
        var v8 float32 = 0.1
        _ = v8
    }
    v4 = v4
    if v4 {
        var v9 uint = 1
        _ = v9
    } else {
        if false {
            var v10 int8 = 0
            _ = v10
        }
        v1 += (v1 >> ((851) & 63))
        switch 0 {
        case -2147483648:
            v11 := GenS1{F0: (v3 << ((0) & 63)), F1: 940, F2: (int64(v3) == 42)}
            _ = v11
        }
    }
    var v12 uint32 = (v1 & v1)
    v3 = v3
    if (0 == 1) {
        var v13 int64 = -336
        v3 = v3
        var v14 uint8 = (v3 ^ v3)
        _ = v13
        _ = v14
    }
    for fi15 := 0; fi15 < 2; fi15++ {
        fmt.Println(v12, v4)
        _ = fi15
    }
    var v16 bool = (-2147483648 != 42)
    var v17 int32 = int32(v3)
    v18 := GenS1{F0: v3, F1: 42, F2: v16}
    v16 = (42 != v1)
    for fi19 := 0; fi19 < 9; fi19++ {
        for fi20 := 0; fi20 < 11; fi20++ {
            fmt.Println(fi20, v12, fi19)
            _ = fi20
        }
        _ = fi19
    }
    var v21 bool = v16
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v21:", v21)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v12
    _ = v16
    _ = v17
    _ = v18
    _ = v2
    _ = v21
    _ = v3
    _ = v4
    _ = v7
}
