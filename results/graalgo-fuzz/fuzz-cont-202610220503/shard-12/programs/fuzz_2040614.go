package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 int; F2 uint64 }

func main() {
    var v1 float64 = 1e10
    var v2 uint16 = 255
    v3 := []uint32{}
    if true {
        fmt.Println(v1)
        v1 = v1
        v1 -= (v1 + ((0.1 - ((float64(float32(((v1 / (v1 - float64(1.5))) - 1e10))) / v1) * v1)) / v1))
    } else {
        for fi4 := 0; fi4 < 11; fi4++ {
            var v5 float64 = float64(1.5)
            v2 = v2
            _ = fi4
            _ = v5
        }
    }
    v1 += v1
    var v6 float32 = 1e10
    if false {
        v3 = append(v3, 0)
        if false {
            var v7 int8 = -46
            _ = v7
        }
    }
    var v8 int16 = 0
    for fi9 := 0; fi9 < 11; fi9++ {
        if len(v3) > 0 {
            v3[0] = 2147483647
        }
        v10 := GenS1{F0: 42, F1: -1, F2: 4294967295}
        _ = fi9
        _ = v10
    }
    var v11 uint8 = 141
    v6 += v6
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v6:", v6)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v11
    _ = v2
    _ = v3
    _ = v6
    _ = v8
}
