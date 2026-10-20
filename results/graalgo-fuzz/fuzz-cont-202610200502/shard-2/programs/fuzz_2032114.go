package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 int; F2 int32 }
type GenS2 struct { F0 uint16; F1 uint64 }

func main() {
    v1 := make(map[int]string)
    var v2 uint = 190
    v3 := make(map[string]bool)
    var v4 int8 = 127
    var v5 uint32 = 357
    v5 = (v5 ^ (v5 ^ v5))
    v2 -= v2
    delete(v1, 42)
    for fi6 := 0; fi6 < 17; fi6++ {
        v4 = (v4 / (((v4 >> (((v2 - 18446744073709551615)) & 63))) | 1))
        v7 := GenS1{F0: 127, F1: 1, F2: 801}
        fmt.Println(v5, v2)
        _ = fi6
        _ = v7
    }
    switch 127 {
    case 32768:
        fmt.Println(v2, v4)
    case 340:
        v5 += (v5 % (((v5 & v5)) | 1))
    case 127:
        var v8 uint32 = 912
        _ = v8
    default:
        if false {
            fmt.Println(v5, v2, v5)
            v5 += 4294967295
        } else {
            v2 = 255
            v9 := GenS1{F0: (v4 | v4), F1: -9223372036854775808, F2: 42}
            _ = v9
        }
    }
    for fi10 := 0; fi10 < 20; fi10++ {
        var v11 int = fi10
        var v12 int8 = v4
        v13 := GenS1{F0: (v4 << ((255) & 63)), F1: (v11 >> (((v2 >> (((v2 & v2)) & 63))) & 63)), F2: 0}
        _ = fi10
        _ = v11
        _ = v12
        _ = v13
    }
    v14 := GenS1{F0: (v4 & 127), F1: 127, F2: 1}
    switch 0 {
    case 128:
        switch -9223372036854775808 {
        case 198:
            var v15 float64 = (-2.25 + -2.25)
            _ = v15
        case 1:
            v5 = (v5 % (((^(v5))) | 1))
        default:
            var v16 float32 = float32(3.14159)
            _ = v16
        }
    case 2147483647:
        var v17 uint64 = 340
        _ = v17
    }
    v18 := []int{}
    _ = len(v18)
    v2 += 65535
    v19 := GenS2{F0: 2, F1: 18446744073709551615}
    var v20 uint = v2
    v21 := v1[-9223372036854775808]
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v18:", len(v18), v18)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v21:", v21)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v14
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v3
    _ = v4
    _ = v5
}
