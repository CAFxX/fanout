package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 bool }
type GenS2 struct { F0 uint32 }
type GenS3 struct { F0 uint64; F1 bool; F2 int32 }

func main() {
    var v1 int16 = 42
    var v2 int64 = 42
    var v3 float32 = 0.0
    for fi4 := 0; fi4 < 20; fi4++ {
        var v5 int64 = (v2 / ((1) | 1))
        if true {
            v6 := GenS1{F0: (fi4 % ((fi4) | 1)), F1: false}
            v7 := []int{}
            v5 -= (v2 - (v5 | 42))
            _ = v6
            _ = v7
        }
        _ = fi4
        _ = v5
    }
    v3 -= float32(3.14159)
    v8 := GenS1{F0: -513, F1: (!(474 >= 0))}
    v9 := []int{168}
    _ = len(v9)
    switch 32768 {
    case 66:
        fmt.Println(v2)
    }
    v10 := GenS2{F0: 42}
    v11 := GenS3{F0: 1, F1: true, F2: 0}
    v12 := GenS3{F0: 2, F1: ((v2 <= v2) || true), F2: 2147483647}
    v2 = v2
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v9:", len(v9), v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v2
    _ = v3
    _ = v8
    _ = v9
}
