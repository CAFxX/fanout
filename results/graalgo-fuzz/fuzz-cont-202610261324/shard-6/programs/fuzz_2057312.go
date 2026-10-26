package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 uint8 }
type GenS2 struct { F0 int32; F1 uint64 }

func main() {
    var v1 float64 = -2.25
    var v2 bool = true
    v3 := make(map[int]uint32)
    v2 = v2
    v4 := []int{-2147483648}
    if true {
        var v5 bool = true
        var v6 float64 = 1.5
        _ = v5
        _ = v6
    } else {
        if len(v4) > 0 {
            v4[0] = 42
        }
    }
    switch "" {
    case "hello":
        var v7 uint = 4294967295
        _ = v7
    case "":
        if false {
            var v8 bool = v2
            _ = v8
        }
    }
    v9 := GenS1{F0: 42, F1: 1}
    if len(v4) > 0 {
        v4[0] = -2147483648
    }
    switch 0 {
    case 42:
        var v10 int8 = -1
        _ = v10
    default:
        var v11 string = "\n\t"
        _ = v11
    }
    for fi12 := 0; fi12 < 4; fi12++ {
        for fi13 := 0; fi13 < 8; fi13++ {
            v3[2147483647] = 1
            if (v2 || (401 != 508)) {
                _ = len(v4)
                switch 1 {
                case 128:
                    v1 += (((v1 + v1) * -2.25) - (v1 + v1))
                case 0:
                    v2 = v2
                case -2147483648:
                    v1 = 1e10
                }
                v2 = true
                v14 := GenS2{F0: 0, F1: 249}
                _ = v14
            }
            fmt.Println(v1, v1)
            _ = fi13
        }
        _ = fi12
    }
    var v15 uint32 = 13
    switch 656 {
    case -213:
        var v16 int8 = 100
        _ = v16
    case 700:
        _ = len(v3)
    default:
        var v17 bool = v2
        _ = v17
    }
    if len(v4) > 0 {
        v4[0] = -9223372036854775808
    }
    v18 := GenS1{F0: -128, F1: 128}
    if v2 {
        v4 = append(v4, 2147483647)
        if len(v4) >= 2 {
            v4 = v4[1:len(v4)]
        }
    }
    v15 = (^(v15))
    fmt.Println("v1:", v1)
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v4:", len(v4), v4)
    _ = v1
    _ = v15
    _ = v18
    _ = v2
    _ = v3
    _ = v4
    _ = v9
}
