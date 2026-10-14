package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 uint; F2 uint64 }
type GenS2 struct { F0 uint64; F1 string; F2 int8 }
type GenS3 struct { F0 bool; F1 uint64 }

func main() {
    v1 := []uint32{65535}
    var v2 uint = 0
    var v3 float32 = (1e10 * (((((1.5 + (1.5 / 1.5)) * 0.1) / ((3.14159 / float32((0.1 * 3.14159))) + (3.14159 - (0.0 - 1.5)))) / 1e10) / float32(-2.25)))
    v4 := make(map[int]string)
    var v5 int16 = 1
    v6 := make(map[uint32]uint32)
    var v7 float32 = (-2.25 + 0.1)
    v2 = (v2 / (((^(v2))) | 1))
    for fi8 := 0; fi8 < 11; fi8++ {
        v3 = 0.0
        if (fi8 >= fi8) {
            delete(v4, fi8)
            if ((true && false) || false) {
                fmt.Println(v3, fi8, v2)
                v9 := GenS1{F0: 68, F1: v2, F2: 300}
                _ = v9
            }
        }
        _ = fi8
    }
    v3 = v7
    v10 := GenS2{F0: 2, F1: "\n\t", F2: 67}
    fmt.Println(v5)
    v5 = v5
    _ = len(v6)
    var v11 int8 = 1
    switch -79 {
    case 9223372036854775807:
        v12 := GenS3{F0: (!((v7 - -2.25) < ((v3 + 3.14159) - v3))), F1: 0}
        _ = v12
    }
    var v13 uint32 = 4294967295
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v11:", v11)
    fmt.Println("v13:", v13)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    fmt.Println("v7:", v7)
    _ = v1
    _ = v10
    _ = v11
    _ = v13
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
