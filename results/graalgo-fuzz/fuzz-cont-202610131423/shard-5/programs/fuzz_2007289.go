package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 int32 }
type GenS2 struct { F0 uint16; F1 int }

func main() {
    var v1 int8 = -1
    var v2 int8 = (v1 | (v1 / ((v1) | 1)))
    var v3 uint8 = 127
    v4 := make(map[int]uint32)
    var v5 bool = (255 != 89)
    for fi6 := 0; fi6 < 12; fi6++ {
        v2 = v1
        v7 := GenS1{F0: 2147483647, F1: 42}
        for fi8 := 0; fi8 < 20; fi8++ {
            var v9 int = (fi6 ^ (fi8 - (fi8 | (fi6 ^ fi8))))
            _ = fi8
            _ = v9
        }
        _ = fi6
        _ = v7
    }
    v10 := []int{-322, 128, 32767, -356}
    v5 = v5
    v11 := GenS2{F0: 65535, F1: 128}
    var v12 int32 = -484
    var v13 float32 = 0.1
    v13 += v13
    v14 := GenS2{F0: 65535, F1: 32768}
    if (false && v5) {
        for fi15 := 0; fi15 < 1; fi15++ {
            v1 -= 1
            fmt.Println(v5, fi15)
            _ = fi15
        }
        var v16 string = ""
        _ = v16
    }
    for fi17 := 0; fi17 < 19; fi17++ {
        var v18 int16 = -32768
        delete(v4, 763)
        _ = fi17
        _ = v18
    }
    v2 = v2
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v10:", len(v10), v10)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v13
    _ = v14
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
