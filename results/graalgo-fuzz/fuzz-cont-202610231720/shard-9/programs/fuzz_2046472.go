package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 int32 }

func main() {
    var v1 uint32 = 2147483647
    var v2 float32 = float32(1.5)
    var v3 uint = 764
    v4 := make(map[int]int)
    var v5 uint64 = 476
    for fi6 := 0; fi6 < 1; fi6++ {
        v7 := GenS1{F0: (v1 % ((4294967295) | 1)), F1: 951}
        for fi8 := 0; fi8 < 14; fi8++ {
            if true {
                fmt.Println(fi6, v1, v3)
                v9 := []int{}
                _ = v9
            }
            var v10 float32 = v2
            v5 += v5
            _ = fi8
            _ = v10
        }
        _ = fi6
        _ = v7
    }
    fmt.Println(v3, v2)
    v11 := []int{128, 32767, 4294967295}
    if len(v11) >= 2 {
        v11 = v11[1:len(v11)]
    }
    fmt.Println(v5, v2, v2)
    fmt.Println(v5, v3)
    v3 += v3
    for fi12 := 0; fi12 < 18; fi12++ {
        var v13 int32 = -491
        _ = fi12
        _ = v13
    }
    var v14 int32 = 352
    v3 += (v3 + (v3 - v3))
    for fi15 := 0; fi15 < 16; fi15++ {
        switch fi15 {
        case 1:
            fmt.Println(v14)
        default:
            var v16 uint8 = 2
            _ = v16
        }
        _ = fi15
    }
    if true {
        v5 += (v5 - v5)
    }
    var v17 uint16 = 42
    switch "\n\t" {
    case "foo bar":
        var v18 int8 = -128
        _ = v18
    }
    v4[-1] = -9223372036854775808
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v11:", len(v11), v11)
    fmt.Println("v14:", v14)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    _ = v1
    _ = v11
    _ = v14
    _ = v17
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
