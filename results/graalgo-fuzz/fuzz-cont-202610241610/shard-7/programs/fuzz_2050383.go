package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 int; F2 uint8 }
type GenS2 struct { F0 string; F1 int64; F2 int16 }
type GenS3 struct { F0 string; F1 uint8 }

func main() {
    var v1 string = ("" + "foo bar")
    v2 := make(map[uint32]bool)
    var v3 int = 220
    var v4 uint64 = 2
    var v5 uint8 = 2
    var v6 string = v1
    var v7 int16 = 42
    v4 -= v4
    delete(v2, 1)
    if false {
        var v8 int8 = 1
        _ = v8
    }
    v9 := GenS1{F0: 489, F1: -1, F2: (v5 + (^(v5)))}
    v10 := GenS2{F0: v1, F1: 9223372036854775807, F2: (v7 >> ((909) & 63))}
    v11 := GenS3{F0: (((v1 + v1) + v6) + v6), F1: v5}
    v4 = v4
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v10
    _ = v11
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v9
}
