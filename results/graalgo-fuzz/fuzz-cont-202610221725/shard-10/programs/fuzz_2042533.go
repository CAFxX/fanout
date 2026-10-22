package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
type GenS2 struct { F0 uint64; F1 uint64 }

func main() {
    var v1 int64 = -1
    var v2 uint8 = 0
    var v3 int = -9223372036854775808
    var v4 int = (v3 >> ((0) & 63))
    v5 := []uint32{180}
    var v6 int64 = (v1 >> ((42) & 63))
    var v7 uint16 = 42
    var v8 uint64 = 4294967295
    var v9 int8 = 100
    fmt.Println(v7, v1, v1)
    for fi10 := 0; fi10 < 2; fi10++ {
        var v11 uint64 = v8
        _ = fi10
        _ = v11
    }
    for fi12 := 0; fi12 < 6; fi12++ {
        var v13 int = fi12
        _ = fi12
        _ = v13
    }
    v14 := GenS1{F0: 0}
    switch "\n\t" {
    case "\n\t":
        for fi15 := 0; fi15 < 1; fi15++ {
            v16 := GenS2{F0: (v8 / ((v8) | 1)), F1: v8}
            _ = fi15
            _ = v16
        }
    case "foo bar":
        var v17 int32 = -784
        _ = v17
    }
    var v18 uint8 = (v2 ^ (^(v2)))
    var v19 float32 = (0.0 * 3.14159)
    fmt.Println("v1:", v1)
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v14
    _ = v18
    _ = v19
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
