package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint8; F2 int64 }

func main() {
    v1 := []uint32{0, 65535}
    v2 := make(map[string]uint32)
    var v3 float32 = (((-2.25 * (((1.5 - 1.5) / float32(3.14159)) / 3.14159)) / 1.0) * 3.14159)
    var v4 uint = 458
    v5 := make(map[string]bool)
    var v6 int16 = -138
    if len(v1) > 0 {
        v1[0] = 4294967295
    }
    var v7 uint16 = 0
    fmt.Println(v4)
    for fi8 := 0; fi8 < 11; fi8++ {
        if len(v1) >= 2 {
            v1 = v1[1:len(v1)]
        }
        _ = fi8
    }
    switch "\n\t" {
    case "":
        if false {
            v4 = (v4 << (((v4 - v4)) & 63))
        }
    case "hello":
        var v9 bool = false
        _ = v9
    default:
        v10 := GenS1{F0: 32767, F1: 2, F2: 1}
        _ = v10
    }
    v11 := v2["a"]
    if true {
        v7 -= v7
        for fi12 := 0; fi12 < 14; fi12++ {
            _ = len(v1)
            _ = fi12
        }
    }
    v6 += (v6 * v6)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v11:", v11)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v11
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
