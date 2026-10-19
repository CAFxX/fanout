package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 bool }
type GenS2 struct { F0 int32; F1 uint32; F2 int }
func genFn0(p0 uint16) uint64 {
    v4 := []int{0, 42, 128}
    if len(v4) > 0 {
        v4[0] = -63
    }
    var v5 int8 = -128
    fmt.Println(v5)
    _ = p0
    _ = v4
    _ = v5
    return 1
}

func main() {
    v1 := []uint32{0}
    var v2 bool = (!((true && (2147483647 > 0)) || (0 >= -1)))
    var v3 int32 = 0
    _ = len(v1)
    var v6 uint8 = 2
    switch "" {
    case "hello":
        fmt.Println(v6)
    case "a":
        switch "\n\t" {
        case "":
            if (11 == 1) {
                v3 -= (v3 | v3)
            } else {
                if len(v1) > 0 {
                    v1[0] = 501
                }
                if len(v1) >= 2 {
                    v1 = v1[1:len(v1)]
                }
            }
        case "a":
            fmt.Println(v3, v6, v3)
        case "\n\t":
            if len(v1) >= 2 {
                v1 = v1[1:len(v1)]
            }
        default:
            if v2 {
                if (!(!((-861 != -2147483648) && (!true)))) {
                    var v7 uint8 = v6
                    _ = v7
                } else {
                    if v2 {
                        v2 = (733 != -1)
                    }
                }
            } else {
                var v8 uint32 = 65535
                v2 = false
                v2 = (v8 != (v8 / (((v8 - 0)) | 1)))
                _ = v8
            }
        }
    }
    var v9 uint64 = 42
    v6 += (v6 / ((128) | 1))
    v1 = append(v1, 0)
    v2 = v2
    var v10 uint32 = 42
    v11 := GenS1{F0: 127, F1: v2}
    if false {
        for fi12 := 0; fi12 < 20; fi12++ {
            v3 -= 175
            for fi13 := 0; fi13 < 13; fi13++ {
                v1 = append(v1, v10)
                v3 -= v3
                var v14 float64 = ((0.1 / 3.14159) + 1.5)
                _ = fi13
                _ = v14
            }
            _ = fi12
        }
        fmt.Println(v9, v9)
        fmt.Println(v6, v6)
    } else {
        fmt.Println(v2)
    }
    var v15 float64 = (0.1 * 1.5)
    for fi16 := 0; fi16 < 17; fi16++ {
        var v17 uint64 = 4294967295
        _ = fi16
        _ = v17
    }
    switch "" {
    case "":
        v18 := GenS2{F0: (v3 & v3), F1: 2, F2: 1}
        _ = v18
    case "\n\t":
        v19 := GenS2{F0: (v3 << ((4294967295) & 63)), F1: 65535, F2: -1}
        _ = v19
    }
    fmt.Println(v15, v2)
    v10 += 2
    fmt.Println(v3, v15, v10)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v10:", v10)
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v6:", v6)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v15
    _ = v2
    _ = v3
    _ = v6
    _ = v9
}
