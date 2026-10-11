package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
type GenS2 struct { F0 uint32; F1 uint8 }
func genFn0(p0 int8, p1 int64, p2 int16) int64 {
    p1 -= 9223372036854775807
    var v5 uint16 = 361
    var v6 int32 = 0
    _ = p0
    _ = p1
    _ = p2
    _ = v5
    _ = v6
    return (-(p1))
}

func main() {
    v1 := make(map[uint32]string)
    var v2 uint8 = 255
    var v3 int16 = 97
    var v4 float64 = 0.0
    var v7 uint64 = 42
    v8 := []int{32767, 42}
    v3 = -192
    fmt.Println(v2)
    var v9 int16 = (v3 ^ v3)
    v10 := v1[1]
    fmt.Println(v3, v7)
    if len(v8) >= 2 {
        v8 = v8[1:len(v8)]
    }
    if (42 <= 1) {
        if len(v8) > 0 {
            v8[0] = 0
        }
        v11 := GenS1{F0: 100}
        for fi12 := 0; fi12 < 6; fi12++ {
            v13 := GenS2{F0: 65535, F1: 0}
            _ = fi12
            _ = v13
        }
        _ = v11
    } else {
        v3 += v9
    }
    v2 += (v2 + (^(v2)))
    v14 := GenS2{F0: 0, F1: 128}
    fmt.Println(v2)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v10:", v10)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v7:", v7)
    fmt.Println("slicelen v8:", len(v8), v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v14
    _ = v2
    _ = v3
    _ = v4
    _ = v7
    _ = v8
    _ = v9
}
