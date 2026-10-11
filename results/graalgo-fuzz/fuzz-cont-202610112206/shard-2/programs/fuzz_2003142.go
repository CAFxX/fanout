package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 int16 }

func main() {
    v1 := make(map[uint32]bool)
    var v2 int = 902
    var v3 uint32 = 1
    var v4 int64 = -415
    var v5 int16 = -1
    v6 := []int{1, 2147483647, v2, v2}
    var v7 int32 = 872
    fmt.Println(v4, v3)
    switch (v2 ^ v2) {
    case 0:
        fmt.Println(v4, v3)
    default:
        var v8 uint8 = 64
        _ = v8
    }
    var v9 uint64 = 42
    if (0 > v2) {
        v2 = v2
        var v10 int = 4294967295
        _ = v10
    }
    var v11 uint64 = 18446744073709551615
    fmt.Println(v2)
    fmt.Println(v7, v4)
    v4 += (^(v4))
    var v12 uint16 = 0
    for fi13 := 0; fi13 < 12; fi13++ {
        for fi14 := 0; fi14 < 20; fi14++ {
            if len(v6) > 0 {
                v6[0] = fi13
            }
            fmt.Println(v3, v12, v9)
            v4 += v4
            _ = fi14
        }
        _ = fi13
    }
    delete(v1, (v3 / ((2) | 1)))
    v1[4294967295] = false
    v15 := GenS1{F0: ("" + "\n\t"), F1: (v5 / ((0) | 1))}
    var v16 string = "\n\t"
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    fmt.Println("v7:", v7)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v11
    _ = v12
    _ = v15
    _ = v16
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v9
}
