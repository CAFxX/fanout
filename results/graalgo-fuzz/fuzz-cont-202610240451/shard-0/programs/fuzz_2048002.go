package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
type GenS2 struct { F0 uint }
type GenS3 struct { F0 uint16; F1 bool; F2 uint64 }

func main() {
    var v1 int32 = 2147483647
    v2 := make(map[uint32]string)
    var v3 int8 = -79
    var v4 float32 = 0.1
    v5 := []uint32{}
    v6 := make(map[string]uint32)
    v1 -= 42
    var v7 int16 = -96
    switch "hello" {
    case "hello":
        switch "\n\t" {
        case "\n\t":
            var v8 int32 = (v1 & -1)
            _ = v8
        case "a":
            if (!false) {
                v7 -= v7
            }
        default:
            v9 := GenS1{F0: 2147483647}
            _ = v9
        }
    }
    var v10 bool = true
    var v11 int32 = 42
    if (502 <= 0) {
        fmt.Println(v3)
    }
    v12 := GenS2{F0: 0}
    if (!false) {
        for fi13 := 0; fi13 < 11; fi13++ {
            var v14 uint64 = 4294967295
            var v15 float32 = v4
            _ = len(v6)
            _ = fi13
            _ = v14
            _ = v15
        }
        v1 -= int32(v3)
        v1 = v11
    }
    v16 := GenS3{F0: 144, F1: (2 == 1), F2: 2}
    v1 = v11
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("maplen v6:", len(v6))
    fmt.Println("v7:", v7)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v16
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
