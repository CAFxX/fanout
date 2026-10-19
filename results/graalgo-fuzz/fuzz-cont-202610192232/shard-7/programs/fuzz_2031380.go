package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 int64 }
type GenS2 struct { F0 int }
type GenS3 struct { F0 uint8 }
func genFn0(p0 bool) uint64 {
    if (p0 || p0) {
        var v4 int = -2147483648
        _ = v4
    }
    var v5 uint16 = 42
    if false {
        v6 := []int{339}
        p0 = (float32(0.0) != ((((0.0 - (-2.25 * float32(0.1))) + 1e10) / 1.5) / (((-2.25 / 1.5) + ((((0.1 + 1e10) - 0.1) / (0.1 * -2.25)) / ((1e10 + (3.14159 / 0.1)) / (-2.25 / 1e10)))) * -2.25)))
        _ = v6
    }
    v7 := []int{4294967295}
    if len(v7) > 0 {
        v7[0] = 42
    }
    _ = p0
    _ = v5
    _ = v7
    return 2
}

func main() {
    var v1 uint16 = 42
    var v2 int16 = -32768
    var v3 int64 = -1
    var v8 uint16 = v1
    v9 := GenS1{F0: 0, F1: v3}
    for fi10 := 0; fi10 < 7; fi10++ {
        switch 205 {
        case -9223372036854775808:
            v11 := []int{fi10, 2147483647}
            _ = v11
        case -508:
            v2 -= (v2 * (^(v2)))
        }
        switch "\n\t" {
        case "\n\t":
            v12 := GenS2{F0: (fi10 + fi10)}
            _ = v12
        default:
            if false {
                v13 := GenS3{F0: 35}
                if false {
                    fmt.Println(v3, fi10, v3)
                    var v14 float64 = 1e10
                    _ = v14
                }
                _ = v13
            }
        }
        _ = fi10
    }
    v1 = v8
    v15 := []int{-652}
    if len(v15) > 0 {
        v15[0] = -1
    }
    fmt.Println(v2)
    v2 -= int16(v3)
    var v16 float32 = -2.25
    v3 = (v3 + (^(v3)))
    var v17 int = 32767
    for fi18 := 0; fi18 < 19; fi18++ {
        _ = len(v15)
        v19 := GenS1{F0: 0, F1: 1}
        _ = fi18
        _ = v19
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v15:", len(v15), v15)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v15
    _ = v16
    _ = v17
    _ = v2
    _ = v3
    _ = v8
    _ = v9
}
