package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 uint64; F2 uint16 }

func main() {
    var v1 int = 4294967295
    var v2 int16 = -404
    var v3 int = v1
    v4 := []int{(v3 - (v3 - (v1 << ((958) & 63)))), (^(v1))}
    var v5 int16 = (-(v2))
    v6 := GenS1{F0: 65535, F1: 242, F2: 28}
    v3 = -910
    var v7 uint = 0
    if (!false) {
        for fi8 := 0; fi8 < 3; fi8++ {
            _ = len(v4)
            v9 := GenS1{F0: 4294967295, F1: 867, F2: 255}
            _ = fi8
            _ = v9
        }
        v10 := GenS1{F0: v7, F1: 2, F2: 375}
        var v11 int32 = -2147483648
        _ = v10
        _ = v11
    }
    var v12 uint64 = 18446744073709551615
    v7 = (v7 / ((471) | 1))
    fmt.Println(v5, v1, v1)
    v1 = (^(v3))
    switch "" {
    case "a":
        fmt.Println(v3, v5)
    }
    if len(v4) >= 2 {
        v4 = v4[1:len(v4)]
    }
    var v13 uint = 0
    var v14 int16 = int16(v13)
    var v15 bool = false
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v12
    _ = v13
    _ = v14
    _ = v15
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
