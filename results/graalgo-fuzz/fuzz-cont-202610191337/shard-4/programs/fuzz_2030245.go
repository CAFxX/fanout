package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 uint8 }

func main() {
    var v1 int16 = 1000
    v2 := []uint32{42, 65535, 2147483648}
    var v3 uint8 = 255
    v4 := []uint32{42, 470, 1, uint32(v3)}
    var v5 bool = true
    v1 += (v1 - v1)
    if len(v4) > 0 {
        v4[0] = 921
    }
    v5 = true
    for fi6 := 0; fi6 < 9; fi6++ {
        var v7 uint32 = uint32(v3)
        _ = fi6
        _ = v7
    }
    v3 += 255
    switch -105 {
    case 1:
        for fi8 := 0; fi8 < 11; fi8++ {
            v9 := GenS1{F0: v3, F1: 255}
            _ = fi8
            _ = v9
        }
    case 995:
        for fi10 := 0; fi10 < 15; fi10++ {
            for fi11 := 0; fi11 < 10; fi11++ {
                var v12 float64 = 0.1
                v5 = (int64(v1) >= -460)
                var v13 uint8 = (v3 + v3)
                _ = fi11
                _ = v12
                _ = v13
            }
            _ = fi10
        }
    }
    var v14 int8 = 85
    fmt.Println("v1:", v1)
    fmt.Println("v14:", v14)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v14
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
