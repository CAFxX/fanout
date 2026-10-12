package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 int16; F2 uint8 }

func main() {
    var v1 int8 = 1
    v2 := make(map[int]bool)
    var v3 float32 = (1.5 * (3.14159 / 1.0))
    fmt.Println(v1, v3)
    var v4 int16 = 350
    v5 := GenS1{F0: 464, F1: v4, F2: 35}
    for fi6 := 0; fi6 < 17; fi6++ {
        switch "" {
        case "a":
            _ = len(v2)
        case "hello":
            var v7 uint64 = 18446744073709551615
            _ = v7
        case "\n\t":
            v4 += (v4 >> ((352) & 63))
        default:
            var v8 uint = 42
            _ = v8
        }
        _ = fi6
    }
    v3 = 1e10
    var v9 uint8 = 1
    var v10 uint16 = 141
    var v11 string = ""
    delete(v2, 127)
    fmt.Println(v9, v9, v3)
    var v12 uint64 = 1
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v9
}
