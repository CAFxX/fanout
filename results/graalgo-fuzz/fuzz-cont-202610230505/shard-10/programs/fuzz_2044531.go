package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 uint32; F2 int32 }
type GenS2 struct { F0 int16; F1 uint8 }
func genFn0(p0 int) string {
    var v6 bool = (977 > 1000)
    switch p0 {
    case 0:
        if (2147483647 < 1) {
            var v7 uint8 = 128
            _ = v7
        }
    default:
        var v8 float32 = 0.0
        _ = v8
    }
    var v9 int64 = 52
    _ = p0
    _ = v6
    _ = v9
    return "foo bar"
}

func main() {
    var v1 int8 = -128
    var v2 uint64 = 2
    var v3 uint16 = uint16(v2)
    var v4 int32 = 2147483647
    var v5 bool = (65535 > 18446744073709551615)
    v10 := GenS1{F0: 1, F1: 472, F2: v4}
    v11 := []int{}
    v12 := GenS1{F0: -32768, F1: 0, F2: (-(v4))}
    if len(v11) > 0 {
        v11[0] = -2147483648
    }
    if true {
        var v13 float64 = float64(3.14159)
        _ = v13
    } else {
        fmt.Println(v3, v3)
    }
    v5 = (!(!v5))
    v14 := GenS2{F0: -32768, F1: 0}
    if (v2 == v2) {
        var v15 uint32 = 1
        if ((v3 ^ (v3 >> ((1) & 63))) == (v3 + (^(v3)))) {
            if v5 {
                for fi16 := 0; fi16 < 5; fi16++ {
                    var v17 int64 = -92
                    v3 += (v3 - (v3 & v3))
                    _ = fi16
                    _ = v17
                }
            } else {
                v15 = (v15 + (^(v15)))
            }
        }
        _ = v15
    }
    var v18 uint = 255
    fmt.Println(v5, v1, v3)
    if v5 {
        v19 := GenS2{F0: 1, F1: 182}
        for fi20 := 0; fi20 < 3; fi20++ {
            v18 = v18
            for fi21 := 0; fi21 < 14; fi21++ {
                if (v3 != 1) {
                    if (!(float64(3.14159) == (0.1 * 3.14159))) {
                        var v22 bool = ((v2 >= v2) || v5)
                        v4 = v4
                        _ = v22
                    }
                } else {
                    var v23 int16 = 0
                    v18 = v18
                    _ = v23
                }
                var v24 uint8 = 154
                _ = fi21
                _ = v24
            }
            var v25 int = (fi20 & -9223372036854775808)
            _ = fi20
            _ = v25
        }
        _ = v19
    }
    var v26 int = -9223372036854775808
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v11:", len(v11), v11)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("v26:", v26)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v14
    _ = v18
    _ = v2
    _ = v26
    _ = v3
    _ = v4
    _ = v5
}
