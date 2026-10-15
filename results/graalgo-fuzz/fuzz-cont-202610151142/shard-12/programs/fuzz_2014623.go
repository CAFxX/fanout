package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 bool }
type GenS2 struct { F0 string; F1 string; F2 int }

func main() {
    var v1 uint8 = 109
    var v2 int16 = 274
    var v3 float32 = (0.0 / 1.5)
    var v4 int8 = -1
    v5 := make(map[int]int)
    var v6 int32 = 0
    for fi7 := 0; fi7 < 20; fi7++ {
        v8 := GenS1{F0: fi7, F1: (0 != 2)}
        fmt.Println(v6)
        _ = fi7
        _ = v8
    }
    for fi9 := 0; fi9 < 16; fi9++ {
        var v10 uint64 = 2
        v10 = (v10 ^ (^(v10)))
        v2 = (v2 * (v2 | v2))
        _ = fi9
        _ = v10
    }
    var v11 int32 = v6
    if (1 != 4294967295) {
        var v12 int64 = 9223372036854775807
        v13 := GenS1{F0: 1, F1: true}
        _ = v12
        _ = v13
    }
    _ = len(v5)
    v5[-1] = -1
    fmt.Println(v6, v3, v6)
    v14 := GenS2{F0: ("foo bar" + "hello"), F1: "\n\t", F2: int(v6)}
    fmt.Println(v1, v2)
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v11
    _ = v14
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
