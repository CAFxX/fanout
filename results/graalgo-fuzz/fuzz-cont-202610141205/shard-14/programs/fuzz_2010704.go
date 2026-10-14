package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 int16; F2 int64 }
type GenS2 struct { F0 uint64; F1 uint16; F2 uint16 }
func genFn0(p0 uint) string {
    fmt.Println(p0)
    if false {
        v5 := []int{2147483647, -148, 1}
        _ = v5
    } else {
        var v6 int32 = -1
        _ = v6
    }
    var v7 uint = 42
    v8 := []int{-263, 1}
    _ = p0
    _ = v7
    _ = v8
    return ("" + ("\n\t" + ""))
}

func main() {
    var v1 int64 = 42
    var v2 int16 = -590
    var v3 bool = false
    v4 := []string{}
    var v9 float64 = 3.14159
    for fi10 := 0; fi10 < 3; fi10++ {
        switch 128 {
        case 127:
            v1 -= (v1 * (^(v1)))
        case 0:
            var v11 int8 = 42
            _ = v11
        case 4294967295:
            fmt.Println(fi10, v3)
        default:
            v2 -= v2
        }
        fmt.Println(fi10, v2)
        var v12 uint16 = 491
        _ = fi10
        _ = v12
    }
    for fi13 := 0; fi13 < 12; fi13++ {
        var v14 int8 = -102
        for fi15 := 0; fi15 < 10; fi15++ {
            if (((v14 >= v14) || false) && (((v14 & (v14 >> ((361) & 63))) <= v14) || true)) {
                var v16 float32 = 3.14159
                _ = v16
            } else {
                if (42 < -883) {
                    _ = len(v4)
                    _ = len(v4)
                    fmt.Println(v14, v2, v2)
                } else {
                    v14 += (^(v14))
                    _ = len(v4)
                }
                _ = len(v4)
                var v17 uint16 = 65535
                _ = v17
            }
            _ = fi15
        }
        _ = fi13
        _ = v14
    }
    fmt.Println(v3, v2)
    for fi18 := 0; fi18 < 12; fi18++ {
        if v3 {
            var v19 float64 = (((v9 * v9) / (v9 - (1e10 / (((v9 / (((float64(3.14159) / v9) / v9) - v9)) - v9) / v9)))) - v9)
            v1 = (v1 | v1)
            _ = v19
        }
        if len(v4) >= 2 {
            v4 = v4[1:len(v4)]
        }
        _ = fi18
    }
    fmt.Println(v9, v1)
    for fi20 := 0; fi20 < 9; fi20++ {
        v4 = append(v4, "hello")
        v1 += 9223372036854775807
        for fi21 := 0; fi21 < 15; fi21++ {
            if false {
                var v22 float32 = 3.14159
                switch -143 {
                case 32768:
                    var v23 float64 = v9
                    _ = v23
                default:
                    var v24 uint32 = 2147483648
                    _ = v24
                }
                var v25 uint = 137
                _ = v22
                _ = v25
            } else {
                v4 = append(v4, "\n\t")
                switch "hello" {
                case "hello":
                    v9 += (v9 * -2.25)
                case "foo bar":
                    switch fi21 {
                    case -9223372036854775808:
                        var v26 int64 = v1
                        _ = v26
                    case 32768:
                        fmt.Println(v3)
                    default:
                        v3 = v3
                    }
                case "a":
                    fmt.Println(v9, v2)
                }
                var v27 uint64 = 42
                _ = v27
            }
            _ = fi21
        }
        _ = fi20
    }
    v9 = (float64(0.0) * 0.0)
    v4 = append(v4, ("\n\t" + "a"))
    for fi28 := 0; fi28 < 4; fi28++ {
        for fi29 := 0; fi29 < 18; fi29++ {
            v30 := GenS1{F0: -1, F1: (v2 ^ v2), F2: 42}
            _ = fi29
            _ = v30
        }
        _ = fi28
    }
    if (1e10 <= v9) {
        var v31 float64 = v9
        fmt.Println(v2, v9)
        v4 = append(v4, ("" + "hello"))
        _ = v31
    } else {
        var v32 int8 = 42
        var v33 int64 = int64(v32)
        _ = v32
        _ = v33
    }
    if len(v4) >= 2 {
        v4 = v4[1:len(v4)]
    }
    switch "foo bar" {
    case "hello":
        v34 := GenS2{F0: 368, F1: 1, F2: 2}
        _ = v34
    }
    for fi35 := 0; fi35 < 4; fi35++ {
        v4 = append(v4, ("\n\t" + "a"))
        var v36 bool = ((0.0 * v9) > 1e10)
        _ = fi35
        _ = v36
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v9
}
