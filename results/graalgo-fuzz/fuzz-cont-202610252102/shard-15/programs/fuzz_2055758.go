package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 int64 }
type GenS2 struct { F0 bool; F1 int32; F2 uint }
func genFn0(p0 bool, p1 uint8, p2 int) uint64 {
    for fi4 := 0; fi4 < 3; fi4++ {
        p0 = p0
        p0 = (p0 && p0)
        _ = fi4
    }
    var v5 int8 = 42
    _ = p0
    _ = p1
    _ = p2
    _ = v5
    return 42
}

func main() {
    var v1 float64 = (1e10 + float64(((1e10 * 3.14159) - float32((1e10 - 1e10)))))
    var v2 uint32 = 1
    var v3 float64 = 1e10
    fmt.Println(v1)
    v6 := []int{-446}
    for fi7 := 0; fi7 < 3; fi7++ {
        if len(v6) > 0 {
            v6[0] = (-(fi7))
        }
        _ = fi7
    }
    v1 -= v3
    fmt.Println(v1, v3, v2)
    var v8 float64 = float64(float32((1e10 * -2.25)))
    if false {
        for fi9 := 0; fi9 < 18; fi9++ {
            if (1 <= 644) {
                if true {
                    var v10 int16 = 42
                    _ = len(v6)
                    var v11 uint64 = 509
                    _ = v10
                    _ = v11
                }
                v6 = append(v6, (fi9 | fi9))
            } else {
                if len(v6) >= 2 {
                    v6 = v6[1:len(v6)]
                }
            }
            _ = fi9
        }
    }
    v12 := GenS1{F0: 0, F1: -9223372036854775808}
    v13 := GenS2{F0: ((2 == 569) || false), F1: 42, F2: 830}
    switch "foo bar" {
    case "hello":
        fmt.Println(v3, v2)
    case "":
        if false {
            v6 = append(v6, -2147483648)
            v14 := GenS2{F0: true, F1: 2147483647, F2: 670}
            v6 = append(v6, 32767)
            _ = v14
        }
    default:
        if len(v6) > 0 {
            v6[0] = 9223372036854775807
        }
    }
    var v15 int = 206
    fmt.Println(v15)
    fmt.Println(v1)
    v16 := GenS1{F0: 871, F1: -1}
    fmt.Println(v8, v3, v15)
    fmt.Println("v1:", v1)
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v6:", len(v6), v6)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v12
    _ = v13
    _ = v15
    _ = v16
    _ = v2
    _ = v3
    _ = v6
    _ = v8
}
