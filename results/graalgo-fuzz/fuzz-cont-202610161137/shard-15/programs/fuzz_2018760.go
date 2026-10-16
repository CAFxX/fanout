package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 int64 }
type GenS2 struct { F0 string }

func main() {
    v1 := make(map[uint32]int)
    var v2 int16 = 0
    v3 := []uint32{2}
    v4 := GenS1{F0: 1, F1: 42}
    for fi5 := 0; fi5 < 5; fi5++ {
        v2 = v2
        var v6 uint8 = 253
        for fi7 := 0; fi7 < 2; fi7++ {
            var v8 uint16 = 663
            _ = fi7
            _ = v8
        }
        _ = fi5
        _ = v6
    }
    v9 := GenS1{F0: 253, F1: -890}
    var v10 bool = (true && true)
    var v11 uint32 = 2147483647
    switch 251 {
    case 32767:
        fmt.Println(v2, v10)
    case 9223372036854775807:
        delete(v1, 443)
    }
    var v12 float32 = (3.14159 * (1.5 * (1.5 - (3.14159 + (1e10 / ((0.0 / 0.1) + (0.0 + 3.14159)))))))
    v2 = (v2 >> ((255) & 63))
    v13 := GenS2{F0: "\n\t"}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v13
    _ = v2
    _ = v3
    _ = v4
    _ = v9
}
