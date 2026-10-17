package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 string }

func main() {
    var v1 uint8 = 211
    var v2 uint = 1
    v3 := []int{-9223372036854775808, 1}
    var v4 float64 = 1e10
    v4 = 0.1
    var v5 string = "foo bar"
    if true {
        var v6 float64 = v4
        v7 := GenS1{F0: int32(v1), F1: v5}
        switch -1 {
        case 32768:
            var v8 float64 = ((1.5 / (v6 - 0.0)) / (0.1 * ((v4 + (float64(3.14159) + v6)) - v6)))
            _ = v8
        default:
            for fi9 := 0; fi9 < 9; fi9++ {
                var v10 uint64 = 42
                switch fi9 {
                case 42:
                    for fi11 := 0; fi11 < 8; fi11++ {
                        fmt.Println(v6)
                        v1 -= 127
                        var v12 bool = (v2 > 65535)
                        _ = fi11
                        _ = v12
                    }
                }
                _ = fi9
                _ = v10
            }
        }
        _ = v6
        _ = v7
    } else {
        v2 -= (v2 | (v2 ^ (^(v2))))
    }
    fmt.Println(v5, v1)
    for fi13 := 0; fi13 < 4; fi13++ {
        if false {
            _ = len(v3)
            var v14 uint64 = 0
            _ = v14
        } else {
            v15 := GenS1{F0: 931, F1: v5}
            for fi16 := 0; fi16 < 15; fi16++ {
                if ((2 <= 503) && (fi16 <= (fi13 / ((fi16) | 1)))) {
                    v1 = (v1 ^ v1)
                    var v17 float32 = 3.14159
                    _ = v17
                } else {
                    var v18 int = fi13
                    var v19 int = v18
                    _ = v18
                    _ = v19
                }
                fmt.Println(fi16, v2)
                _ = fi16
            }
            var v20 string = ""
            _ = v15
            _ = v20
        }
        var v21 int = fi13
        fmt.Println(fi13)
        _ = fi13
        _ = v21
    }
    var v22 int8 = -128
    if true {
        for fi23 := 0; fi23 < 4; fi23++ {
            for fi24 := 0; fi24 < 9; fi24++ {
                v22 += (v22 ^ 100)
                _ = fi24
            }
            _ = fi23
        }
        if len(v3) > 0 {
            v3[0] = 731
        }
    }
    if (("a" + "hello") <= "a") {
        var v25 float32 = float32(v4)
        for fi26 := 0; fi26 < 17; fi26++ {
            var v27 string = "foo bar"
            var v28 uint8 = v1
            switch 4294967295 {
            case 23:
                v25 = (0.0 * 1.5)
            }
            _ = fi26
            _ = v27
            _ = v28
        }
        _ = v25
    } else {
        v4 = float64(((((-2.25 / float32(1.5)) * -2.25) + -2.25) - 1e10))
        v3 = append(v3, 32767)
        _ = len(v3)
    }
    v29 := GenS1{F0: -2147483648, F1: v5}
    var v30 int8 = v22
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v22:", v22)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v30:", v30)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v22
    _ = v29
    _ = v3
    _ = v30
    _ = v4
    _ = v5
}
