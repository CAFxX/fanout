package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 int64 }
type GenS2 struct { F0 int; F1 int64; F2 bool }
type GenS3 struct { F0 int; F1 bool; F2 uint16 }

func main() {
    var v1 int8 = 4
    v2 := make(map[int]string)
    var v3 uint64 = 512
    v4 := GenS1{F0: 1, F1: -1}
    v5 := GenS2{F0: -2147483648, F1: 553, F2: false}
    v6 := GenS3{F0: 128, F1: true, F2: 114}
    fmt.Println(v3, v3)
    switch "hello" {
    case "a":
        v1 -= v1
    case "hello":
        fmt.Println(v3, v3)
    }
    _ = len(v2)
    for fi7 := 0; fi7 < 2; fi7++ {
        var v8 int16 = 42
        _ = fi7
        _ = v8
    }
    v9 := GenS3{F0: -2147483648, F1: false, F2: 584}
    v3 = (v3 / ((v3) | 1))
    v3 = 0
    _ = len(v2)
    fmt.Println(v1)
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v9
}
