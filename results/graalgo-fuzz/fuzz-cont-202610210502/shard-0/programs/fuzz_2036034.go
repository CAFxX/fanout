package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 int64; F2 uint }
type GenS2 struct { F0 uint8; F1 uint16; F2 int16 }
func genFn0(p0 int32) int16 {
    var v4 uint64 = 669
    fmt.Println(p0)
    _ = p0
    _ = v4
    return 32767
}

func main() {
    v1 := make(map[uint32]uint32)
    v2 := []uint32{2147483648, 2147483647, 42}
    var v3 uint16 = 255
    var v5 int16 = 32767
    _ = len(v1)
    v5 = (v5 & v5)
    var v6 float32 = ((-2.25 / (-2.25 * -2.25)) - (1e10 + 0.1))
    v7 := GenS1{F0: uint64(v3), F1: 9223372036854775807, F2: 662}
    v3 = 1
    switch "foo bar" {
    case "":
        for fi8 := 0; fi8 < 8; fi8++ {
            v9 := GenS2{F0: 255, F1: v3, F2: v5}
            for fi10 := 0; fi10 < 17; fi10++ {
                _ = len(v1)
                v11 := GenS1{F0: 4294967295, F1: 0, F2: 2}
                _ = fi10
                _ = v11
            }
            v6 -= ((float32(3.14159) * (0.1 - v6)) - (0.0 - v6))
            _ = fi8
            _ = v9
        }
    case "foo bar":
        fmt.Println(v3, v6)
    }
    fmt.Println(v6, v6)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v5
    _ = v6
    _ = v7
}
