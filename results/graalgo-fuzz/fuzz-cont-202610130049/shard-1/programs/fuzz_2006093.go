package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }

func main() {
    v1 := []uint32{1, 897, 65535}
    var v2 uint32 = 42
    var v3 uint = 0
    v4 := make(map[uint32]uint32)
    var v5 uint8 = 88
    v6 := make(map[string]int)
    var v7 uint64 = 1
    var v8 bool = (100 != 0)
    var v9 uint32 = 4294967295
    if ((v5 + (v5 ^ v5)) <= v5) {
        v7 -= (^(v7))
        v10 := GenS1{F0: (v8 && (!v8))}
        _ = v10
    }
    v5 = (^(v5))
    var v11 int = 278
    if true {
        v12 := GenS1{F0: (v8 || true)}
        _ = v12
    } else {
        switch (v11 % ((v11) | 1)) {
        case -1:
            var v13 int64 = 412
            _ = v13
        default:
            v2 -= (v9 * (v2 * 623))
        }
        switch "\n\t" {
        case "a":
            var v14 float64 = ((-2.25 + float64((1e10 - (((float32(float64(0.1)) - -2.25) - (float32(0.1) + -2.25)) + 0.1)))) * (3.14159 + -2.25))
            _ = v14
        default:
            delete(v4, v9)
        }
        v7 -= (v7 >> ((4294967295) & 63))
    }
    var v15 bool = false
    v5 = (v5 + v5)
    var v16 int16 = -32768
    v17 := v4[(v2 & 2)]
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v11:", v11)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v11
    _ = v15
    _ = v16
    _ = v17
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
