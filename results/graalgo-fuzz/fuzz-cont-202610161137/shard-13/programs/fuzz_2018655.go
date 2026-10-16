package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
type GenS2 struct { F0 int32; F1 uint32; F2 uint8 }
type GenS3 struct { F0 int }

func main() {
    var v1 uint16 = 2
    var v2 int = 127
    v3 := []uint32{376, 42}
    v2 += v2
    var v4 uint = 2
    var v5 bool = ("hello" != "\n\t")
    switch v2 {
    case 32768:
        var v6 uint16 = 771
        _ = v6
    case 0:
        v1 = v1
    default:
        fmt.Println(v5)
    }
    switch (^(v2)) {
    case 9223372036854775807:
        v7 := GenS1{F0: 0}
        _ = v7
    case 42:
        v8 := GenS2{F0: 0, F1: 2, F2: 1}
        _ = v8
    case 1:
        var v9 uint = 0
        _ = v9
    }
    var v10 bool = v5
    if (!(42 == 255)) {
        v11 := GenS3{F0: (v2 ^ (v2 << (((v4 ^ 255)) & 63)))}
        _ = v11
    }
    var v12 float32 = 1.5
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v12:", v12)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v10
    _ = v12
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
