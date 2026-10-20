package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 uint16 }

func main() {
    v1 := make(map[string]uint32)
    var v2 float32 = (1.5 / (1e10 + ((3.14159 * (float32(1.5) * (1e10 * 0.0))) - (0.1 - 3.14159))))
    v3 := make(map[int]uint32)
    v4 := make(map[string]bool)
    var v5 uint16 = 255
    var v6 uint8 = 1
    v6 += 2
    switch "a" {
    case "":
        if true {
            _ = len(v4)
            var v7 int64 = -1
            var v8 uint8 = 128
            _ = v7
            _ = v8
        } else {
            v9 := GenS1{F0: (42 >= -1), F1: (v5 ^ v5)}
            _ = v9
        }
    case "a":
        v6 -= (v6 << ((1) & 63))
    case "hello":
        v6 += (v6 % ((v6) | 1))
    }
    switch "\n\t" {
    case "\n\t":
        for fi10 := 0; fi10 < 16; fi10++ {
            var v11 int8 = 5
            _ = fi10
            _ = v11
        }
    }
    var v12 uint32 = 2147483647
    v5 += 42
    v6 -= v6
    v12 -= v12
    v12 = 89
    var v13 uint16 = (v5 ^ (^(v5)))
    var v14 int64 = 9223372036854775807
    for fi15 := 0; fi15 < 14; fi15++ {
        var v16 uint64 = 2
        _ = fi15
        _ = v16
    }
    _ = len(v4)
    v12 -= v12
    v17 := []int{}
    for fi18 := 0; fi18 < 20; fi18++ {
        if (v5 <= v13) {
            v17 = append(v17, fi18)
            fmt.Println(v2)
            v5 -= 1
        } else {
            v17 = append(v17, 9223372036854775807)
        }
        _ = fi18
    }
    fmt.Println(v6, v13)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("slicelen v17:", len(v17), v17)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v12
    _ = v13
    _ = v14
    _ = v17
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
