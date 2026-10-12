package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 uint64 }
type GenS2 struct { F0 int32 }

func main() {
    var v1 uint = 18446744073709551615
    v2 := make(map[uint32]uint32)
    v3 := make(map[uint32]bool)
    v4 := GenS1{F0: 128, F1: 18446744073709551615}
    v5 := GenS1{F0: 0, F1: 0}
    switch 0 {
    case -9223372036854775808:
        if (v1 < v1) {
            v6 := GenS1{F0: 221, F1: 451}
            v7 := []int{-516, 41}
            fmt.Println(v1)
            _ = v6
            _ = v7
        }
    default:
        if (!false) {
            v1 += 42
        } else {
            for fi8 := 0; fi8 < 7; fi8++ {
                for fi9 := 0; fi9 < 2; fi9++ {
                    v10 := []int{fi8, (fi8 << ((v1) & 63)), fi9, fi9}
                    if len(v10) >= 2 {
                        v10 = v10[1:len(v10)]
                    }
                    v11 := GenS1{F0: 1, F1: 42}
                    _ = fi9
                    _ = v10
                    _ = v11
                }
                v1 += 255
                v12 := GenS1{F0: 1, F1: 4294967295}
                _ = fi8
                _ = v12
            }
        }
    }
    v1 -= v1
    var v13 float32 = float32(3.14159)
    v14 := GenS1{F0: 0, F1: 1}
    fmt.Println(v1, v13)
    v15 := GenS1{F0: 255, F1: 382}
    var v16 int32 = -1
    var v17 uint8 = 1
    fmt.Println(v13)
    v18 := GenS2{F0: (-(v16))}
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    _ = v1
    _ = v13
    _ = v14
    _ = v15
    _ = v16
    _ = v17
    _ = v18
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
