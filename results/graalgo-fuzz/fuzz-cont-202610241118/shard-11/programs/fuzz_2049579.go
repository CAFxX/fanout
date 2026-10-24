package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 int8 }

func main() {
    v1 := make(map[int]uint32)
    var v2 uint8 = 0
    var v3 uint16 = 264
    fmt.Println(v3, v3)
    switch -1 {
    case 9223372036854775807:
        v3 -= (v3 - 1)
    case 32767:
        for fi4 := 0; fi4 < 1; fi4++ {
            v5 := GenS1{F0: false, F1: 0}
            if false {
                var v6 string = "\n\t"
                v7 := GenS1{F0: ((((0.0 / 1.0) + 3.14159) < 3.14159) || true), F1: -96}
                var v8 int64 = -1
                _ = v6
                _ = v7
                _ = v8
            } else {
                v1[fi4] = 260
                fmt.Println(fi4)
                if false {
                    v9 := v1[fi4]
                    _ = v9
                } else {
                    var v10 uint = 255
                    _ = v10
                }
            }
            var v11 int = (fi4 + (^(fi4)))
            _ = fi4
            _ = v11
            _ = v5
        }
    default:
        fmt.Println(v3, v2)
    }
    v2 = (v2 % (((v2 >> ((255) & 63))) | 1))
    fmt.Println(v2, v3)
    v2 -= (v2 + (^(v2)))
    fmt.Println(v3, v3)
    var v12 float64 = ((1e10 + 1.5) * 0.0)
    for fi13 := 0; fi13 < 7; fi13++ {
        var v14 int32 = 0
        _ = fi13
        _ = v14
    }
    v3 -= 42
    var v15 int16 = -1
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v12:", v12)
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v12
    _ = v15
    _ = v2
    _ = v3
}
