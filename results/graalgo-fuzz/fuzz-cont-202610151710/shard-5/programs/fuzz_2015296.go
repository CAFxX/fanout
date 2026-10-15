package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 int16; F2 uint }
type GenS2 struct { F0 int8 }

func main() {
    v1 := []int{0, -740, 2147483647}
    var v2 int64 = 9223372036854775807
    var v3 uint = 990
    var v4 bool = false
    var v5 float64 = (-2.25 * 1.5)
    for fi6 := 0; fi6 < 9; fi6++ {
        if v4 {
            if true {
                var v7 int16 = 1000
                v1 = append(v1, 32767)
                for fi8 := 0; fi8 < 6; fi8++ {
                    if ((v7 < 0) || v4) {
                        v3 -= v3
                        var v9 int64 = (^(v2))
                        _ = v9
                    } else {
                        var v10 uint = (^(v3))
                        _ = v10
                    }
                    _ = fi8
                }
                _ = v7
            } else {
                v2 = (v2 / ((-1) | 1))
                v2 = v2
            }
        }
        _ = fi6
    }
    var v11 bool = (1 >= 823)
    var v12 uint16 = 0
    switch 42 {
    case 247:
        var v13 uint64 = 859
        _ = v13
    }
    if len(v1) > 0 {
        v1[0] = -217
    }
    v14 := GenS1{F0: v2, F1: -123, F2: v3}
    var v15 uint64 = 18446744073709551615
    if v11 {
        if len(v1) > 0 {
            v1[0] = 1
        }
        if len(v1) > 0 {
            v1[0] = int(v12)
        }
    }
    fmt.Println(v2, v4)
    var v16 uint64 = uint64(v12)
    switch 1 {
    case 128:
        switch 2147483647 {
        case 128:
            v12 = (^(v12))
        default:
            v15 = 42
        }
    case 4294967295:
        if len(v1) > 0 {
            v1[0] = 1
        }
    default:
        for fi17 := 0; fi17 < 6; fi17++ {
            v18 := GenS2{F0: -1}
            _ = fi17
            _ = v18
        }
    }
    if v11 {
        _ = len(v1)
        switch -1 {
        case -1:
            if len(v1) > 0 {
                v1[0] = -209
            }
        }
    } else {
        fmt.Println(v2, v5, v16)
    }
    v2 += (v2 >> ((4294967295) & 63))
    switch -9223372036854775808 {
    case -703:
        var v19 int64 = v2
        _ = v19
    default:
        fmt.Println(v11, v15)
    }
    var v20 bool = (83 >= 1)
    fmt.Println(v4, v4)
    v2 = 610
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v11
    _ = v12
    _ = v14
    _ = v15
    _ = v16
    _ = v2
    _ = v20
    _ = v3
    _ = v4
    _ = v5
}
