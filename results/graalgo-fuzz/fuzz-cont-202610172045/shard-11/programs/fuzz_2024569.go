package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 int }
type GenS2 struct { F0 int8; F1 uint16 }
type GenS3 struct { F0 uint8; F1 int16 }
type GenS4 struct { F0 int16; F1 int }

func main() {
    var v1 int16 = 42
    v2 := make(map[uint32]int)
    v3 := []string{"foo bar", "hello", "a"}
    v4 := make(map[uint32]uint32)
    for fi5 := 0; fi5 < 14; fi5++ {
        v1 -= (v1 & (v1 | (v1 + (v1 % ((0) | 1)))))
        var v6 float32 = (-2.25 * (1.5 / (float32(0.1) - float32(float64((0.1 + 1.5))))))
        _ = fi5
        _ = v6
    }
    for fi7 := 0; fi7 < 15; fi7++ {
        if (71 != 127) {
            v2[862] = fi7
            v8 := GenS1{F0: 380, F1: fi7}
            v9 := GenS2{F0: -1, F1: 64}
            _ = v8
            _ = v9
        }
        v10 := GenS2{F0: -128, F1: 255}
        _ = fi7
        _ = v10
    }
    delete(v4, 2147483647)
    v11 := GenS3{F0: 1, F1: v1}
    v12 := GenS4{F0: v1, F1: 508}
    v1 = 1
    v1 -= v1
    v3 = append(v3, "\n\t")
    v1 = 1000
    v1 -= 407
    var v13 uint8 = 127
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v11
    _ = v12
    _ = v13
    _ = v2
    _ = v3
    _ = v4
}
