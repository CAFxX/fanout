package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 int64; F2 int }
type GenS2 struct { F0 int8; F1 int32; F2 uint16 }
type GenS3 struct { F0 uint64; F1 int32 }
type GenS4 struct { F0 uint8; F1 int32; F2 uint }

func main() {
    var v1 float32 = 0.1
    var v2 int = 1
    v3 := []uint32{4294967295, 2147483647}
    var v4 float32 = float32(1.5)
    var v5 int = 122
    var v6 uint32 = 2147483647
    var v7 uint32 = v6
    for fi8 := 0; fi8 < 17; fi8++ {
        var v9 uint64 = 1
        v10 := GenS1{F0: 127, F1: 9223372036854775807, F2: fi8}
        v11 := GenS2{F0: 127, F1: 0, F2: 712}
        _ = fi8
        _ = v10
        _ = v11
        _ = v9
    }
    switch "hello" {
    case "foo bar":
        fmt.Println(v4, v5, v5)
    case "":
        v12 := GenS3{F0: 0, F1: int32(v6)}
        _ = v12
    case "a":
        v5 = (v2 >> ((13) & 63))
    }
    v13 := GenS4{F0: 7, F1: -478, F2: 65535}
    for fi14 := 0; fi14 < 14; fi14++ {
        var v15 string = "\n\t"
        var v16 int64 = 42
        _ = fi14
        _ = v15
        _ = v16
    }
    var v17 int64 = -9223372036854775808
    v18 := GenS1{F0: 127, F1: v17, F2: v5}
    v19 := GenS2{F0: 0, F1: 42, F2: 42}
    switch 128 {
    case 4294967295:
        v17 = (v17 & 0)
    }
    v1 -= ((v4 + (-2.25 - v4)) / v1)
    v7 -= (v6 ^ 487)
    fmt.Println("v1:", v1)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v13
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
