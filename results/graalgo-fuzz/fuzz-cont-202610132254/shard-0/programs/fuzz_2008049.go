package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 bool }

func main() {
    var v1 uint8 = 2
    var v2 int64 = 346
    var v3 int8 = 127
    v4 := []string{"a", ""}
    v5 := make(map[uint32]uint32)
    v3 -= (v3 / ((v3) | 1))
    for fi6 := 0; fi6 < 14; fi6++ {
        v7 := GenS1{F0: v1, F1: (fi6 == (fi6 & fi6))}
        switch 2147483647 {
        case 2147483647:
            v5[104] = 2147483648
        case 4294967295:
            v3 -= 42
        case 812:
            var v8 int16 = -562
            _ = v8
        }
        var v9 uint16 = 65535
        _ = fi6
        _ = v7
        _ = v9
    }
    var v10 int8 = 100
    var v11 uint64 = 546
    if (-2.25 != 0.1) {
        var v12 int64 = v2
        _ = v12
    } else {
        v2 += (v2 >> ((42) & 63))
    }
    _ = len(v5)
    var v13 uint8 = v1
    if (v3 >= v3) {
        v3 = 100
        v13 = (v13 << ((297) & 63))
        var v14 uint32 = 42
        _ = v14
    }
    v2 = v2
    v15 := GenS1{F0: v1, F1: (!true)}
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v13:", v13)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v10
    _ = v11
    _ = v13
    _ = v15
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
