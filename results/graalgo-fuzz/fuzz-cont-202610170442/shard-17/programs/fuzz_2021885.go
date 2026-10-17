package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 uint8; F2 uint32 }

func main() {
    var v1 float32 = -2.25
    v2 := []uint32{2147483647, 0, 615}
    var v3 int = 231
    var v4 float32 = 3.14159
    var v5 int32 = -2147483648
    v6 := make(map[string]int)
    v1 = v1
    var v7 int8 = 42
    var v8 uint16 = 42
    v7 = (v7 + v7)
    v9 := GenS1{F0: -1, F1: 39, F2: 42}
    if false {
        fmt.Println(v4, v7, v1)
        for fi10 := 0; fi10 < 19; fi10++ {
            var v11 uint8 = 0
            _ = fi10
            _ = v11
        }
    } else {
        for fi12 := 0; fi12 < 7; fi12++ {
            if true {
                fmt.Println(v5, v1, v7)
            }
            var v13 uint16 = (v8 | (v8 & (^(v8))))
            v3 -= v3
            _ = fi12
            _ = v13
        }
    }
    v4 = ((v4 - float32(0.1)) + 0.1)
    for fi14 := 0; fi14 < 4; fi14++ {
        v8 += 32767
        for fi15 := 0; fi15 < 16; fi15++ {
            var v16 uint8 = 128
            _ = len(v6)
            fmt.Println(v3, fi14)
            _ = fi15
            _ = v16
        }
        _ = fi14
    }
    if true {
        if true {
            v7 = (v7 / (((-(v7))) | 1))
        }
    }
    var v17 int8 = v7
    var v18 int32 = 0
    fmt.Println("v1:", v1)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v17
    _ = v18
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
