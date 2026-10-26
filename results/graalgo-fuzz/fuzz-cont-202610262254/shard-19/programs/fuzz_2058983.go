package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 int32 }
type GenS2 struct { F0 uint64; F1 int8 }
type GenS3 struct { F0 int; F1 uint64 }

func main() {
    v1 := make(map[string]string)
    v2 := []uint32{228, 4294967295}
    var v3 bool = (1 <= 0)
    v4 := make(map[int]bool)
    v5 := make(map[string]int)
    var v6 float32 = -2.25
    v6 = 0.0
    switch -468 {
    case 4294967295:
        switch 1 {
        case -1:
            var v7 float64 = float64((1.5 - float32(1.5)))
            _ = v7
        default:
            v3 = ((true || true) && v3)
        }
    case -483:
        for fi8 := 0; fi8 < 7; fi8++ {
            v9 := GenS1{F0: -1, F1: 42}
            v6 = float32(0.1)
            fmt.Println(fi8, v6)
            _ = fi8
            _ = v9
        }
    }
    v6 += float32((0.1 - 0.1))
    if true {
        fmt.Println(v3, v6)
        v10 := GenS2{F0: 42, F1: -19}
        if true {
            var v11 float64 = 1.5
            _ = v11
        }
        _ = v10
    }
    v12 := GenS3{F0: -949, F1: 18446744073709551615}
    v6 = (1e10 - float32(3.14159))
    var v13 int16 = 42
    _ = len(v2)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v13:", v13)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v12
    _ = v13
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
