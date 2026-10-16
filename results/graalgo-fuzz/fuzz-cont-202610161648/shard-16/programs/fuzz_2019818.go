package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 uint8 }
type GenS2 struct { F0 int64; F1 int }
func genFn0(p0 int32) int16 {
    v7 := []int{-9223372036854775808, 32767, 32767}
    for fi8 := 0; fi8 < 17; fi8++ {
        p0 -= -2147483648
        _ = fi8
    }
    var v9 uint = 303
    _ = p0
    _ = v7
    _ = v9
    return 32767
}
func genFn1() int16 {
    var v10 int64 = 42
    v11 := GenS1{F0: -9223372036854775808, F1: 1}
    fmt.Println(v10)
    if (642 == 0) {
        switch 0 {
        case 2147483647:
            v10 = v10
        }
        v12 := GenS1{F0: v10, F1: 255}
        _ = v12
    }
    _ = v10
    _ = v11
    return -486
}

func main() {
    v1 := []int{550}
    v2 := make(map[int]string)
    var v3 float64 = 1e10
    v4 := make(map[uint32]string)
    var v5 string = ("" + "foo bar")
    v6 := make(map[string]uint32)
    v5 = v5
    _ = len(v1)
    if false {
        v13 := GenS1{F0: 363, F1: 251}
        v3 = v3
        _ = v13
    } else {
        fmt.Println(v3, v3)
    }
    v5 = v5
    if len(v1) > 0 {
        v1[0] = 390
    }
    v5 = "\n\t"
    var v14 int64 = 42
    var v15 int = 9223372036854775807
    v3 += v3
    var v16 bool = true
    var v17 int64 = (v14 / ((v14) | 1))
    v5 = v5
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    v18 := GenS2{F0: 0, F1: 127}
    switch v15 {
    case 127:
        var v19 float64 = -2.25
        _ = v19
    case 1:
        fmt.Println(v5, v17, v3)
    case 690:
        if (216 == 128) {
            fmt.Println(v14, v5)
        }
    default:
        for fi20 := 0; fi20 < 15; fi20++ {
            if v16 {
                switch v5 {
                case "":
                    var v21 int64 = v14
                    _ = v21
                case "foo bar":
                    var v22 float64 = 3.14159
                    _ = v22
                default:
                    if (-1 >= 1) {
                        fmt.Println(fi20, v17, v17)
                        v3 -= (3.14159 + 0.1)
                        v16 = (842 <= 2)
                    }
                }
                v17 -= v14
                fmt.Println(v16, fi20, v3)
            } else {
                fmt.Println(v16, fi20, v5)
                var v23 int64 = v14
                _ = v23
            }
            _ = fi20
        }
    }
    var v24 int16 = 1
    v16 = ((18446744073709551615 < 18446744073709551615) || (738 >= 2))
    var v25 int64 = 1
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v24:", v24)
    fmt.Println("v25:", v25)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v14
    _ = v15
    _ = v16
    _ = v17
    _ = v18
    _ = v2
    _ = v24
    _ = v25
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
