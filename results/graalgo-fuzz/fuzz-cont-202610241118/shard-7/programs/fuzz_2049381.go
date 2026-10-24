package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 uint64 }

func main() {
    v1 := make(map[uint32]int)
    var v2 int32 = 641
    var v3 uint8 = 1
    v2 = 174
    var v4 uint = 1
    var v5 int = 24
    v4 = 852
    var v6 float32 = (0.0 + 1e10)
    fmt.Println(v5, v6, v4)
    if (true || (1e10 != 3.14159)) {
        v5 = -1
        v7 := GenS1{F0: v4, F1: 2}
        _ = v7
    } else {
        var v8 int64 = 0
        var v9 uint16 = uint16(v4)
        if (!false) {
            var v10 uint32 = uint32(v4)
            fmt.Println(v4)
            _ = v10
        }
        _ = v8
        _ = v9
    }
    if false {
        var v11 int64 = -9223372036854775808
        v4 = (v4 << ((v4) & 63))
        _ = v11
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
