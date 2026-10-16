package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 int64 }

func main() {
    v1 := []int{1, -1, -493, 428}
    var v2 int = 128
    v3 := make(map[uint32]uint32)
    v4 := make(map[int]bool)
    fmt.Println(v2)
    if len(v1) > 0 {
        v1[0] = (-(v2))
    }
    var v5 int16 = 1000
    fmt.Println(v2, v2)
    var v6 uint = 2
    switch v2 {
    case 32767:
        v5 += v5
    case 42:
        fmt.Println(v2)
    default:
        for fi7 := 0; fi7 < 20; fi7++ {
            v8 := GenS1{F0: 0, F1: -9223372036854775808}
            _ = fi7
            _ = v8
        }
    }
    v9 := GenS1{F0: 31, F1: 129}
    v5 = (v5 << (((v6 << ((v6) & 63))) & 63))
    v1 = append(v1, (v2 >> ((v6) & 63)))
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v9
}
