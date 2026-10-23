package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }
type GenS2 struct { F0 int16; F1 int8 }
func genFn0(p0 uint32, p1 uint16, p2 uint32) uint8 {
    if true {
        if (!true) {
            v5 := []int{2147483647, -427, 128}
            fmt.Println(p1)
            _ = v5
        }
        if (float64((1.5 / (0.0 - 1e10))) != float64((1.5 - (1.5 * ((float32(0.0) - 0.0) * (1.5 * (((3.14159 + -2.25) - (0.0 * 1e10)) + 1.5))))))) {
            if (!true) {
                v6 := []int{9223372036854775807, 2147483647}
                var v7 uint16 = p1
                _ = v6
                _ = v7
            } else {
                v8 := []int{-261, 127, 32767}
                if false {
                    var v9 uint64 = 2
                    _ = v9
                }
                _ = v8
            }
        }
    }
    if false {
        for fi10 := 0; fi10 < 16; fi10++ {
            if false {
                p1 -= (p1 % ((p1) | 1))
            } else {
                fmt.Println(p0)
                var v11 uint32 = (^(p2))
                _ = v11
            }
            v12 := GenS1{F0: (!true)}
            _ = fi10
            _ = v12
        }
        var v13 uint16 = 2
        _ = v13
    }
    v14 := []int{-52, 32767, -426, -9223372036854775808}
    _ = p0
    _ = p1
    _ = p2
    _ = v14
    return 128
}

func main() {
    var v1 float32 = float32((-2.25 / (((0.1 + (3.14159 / 1.0)) / float64(0.1)) - -2.25)))
    v2 := make(map[int]int)
    var v3 bool = true
    v4 := []uint32{2}
    delete(v2, 32767)
    var v15 int8 = 127
    var v16 int64 = 0
    v16 -= -9223372036854775808
    var v17 uint8 = 150
    fmt.Println(v15, v17, v3)
    for fi18 := 0; fi18 < 11; fi18++ {
        v16 = (v16 / ((803) | 1))
        v3 = (283 != 716)
        _ = fi18
    }
    v19 := GenS2{F0: -131, F1: (v15 >> ((408) & 63))}
    v20 := GenS2{F0: -256, F1: v15}
    var v21 int16 = 0
    fmt.Println("v1:", v1)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v21:", v21)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    _ = v1
    _ = v15
    _ = v16
    _ = v17
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v3
    _ = v4
}
