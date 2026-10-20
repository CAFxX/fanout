package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
type GenS2 struct { F0 int64; F1 uint16 }
type GenS3 struct { F0 uint }
func genFn0() uint {
    var v5 int = -2147483648
    var v6 uint = 255
    var v7 uint64 = 2
    fmt.Println(v7, v6, v6)
    _ = v5
    _ = v6
    _ = v7
    return (^(v6))
}
func genFn1(p0 uint8) uint16 {
    for fi8 := 0; fi8 < 15; fi8++ {
        v9 := GenS1{F0: -128}
        _ = fi8
        _ = v9
    }
    var v10 uint8 = 160
    _ = p0
    _ = v10
    return 42
}

func main() {
    var v1 uint = 4294967295
    v2 := make(map[string]uint32)
    var v3 float32 = ((float32((1e10 * -2.25)) - 1.5) + ((float32(-2.25) / (0.0 + (3.14159 / 1e10))) * 1.5))
    v4 := make(map[uint32]bool)
    v11 := GenS2{F0: 1, F1: 42}
    var v12 int16 = -1
    v13 := []int{-2147483648, -121}
    var v14 int8 = -1
    fmt.Println(v3, v1, v14)
    if true {
        v15 := GenS1{F0: 66}
        for fi16 := 0; fi16 < 1; fi16++ {
            _ = len(v13)
            v17 := GenS1{F0: v14}
            for fi18 := 0; fi18 < 14; fi18++ {
                v1 = (v1 << ((genFn0()) & 63))
                if true {
                    v19 := GenS3{F0: (v1 - (v1 << (((v1 >> ((v1) & 63))) & 63)))}
                    fmt.Println(v1, fi18)
                    _ = v19
                }
                _ = fi18
            }
            _ = fi16
            _ = v17
        }
        _ = v15
    } else {
        v3 += 1.5
        fmt.Println(v1, v12)
    }
    fmt.Println(v1, v3)
    if (2 <= 457) {
        var v20 uint64 = 1
        _ = v20
    }
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("slicelen v13:", len(v13), v13)
    fmt.Println("v14:", v14)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v11
    _ = v12
    _ = v13
    _ = v14
    _ = v2
    _ = v3
    _ = v4
}
