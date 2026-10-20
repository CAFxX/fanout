package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 bool; F2 uint32 }
type GenS2 struct { F0 int32; F1 uint16 }
func genFn0(p0 uint16) string {
    v4 := []int{42, 50, 0, -2147483648}
    v5 := GenS1{F0: -919, F1: false, F2: 4294967295}
    fmt.Println(p0)
    for fi6 := 0; fi6 < 12; fi6++ {
        switch "foo bar" {
        case "":
            var v7 uint16 = p0
            _ = v7
        default:
            p0 += p0
        }
        var v8 uint32 = 4294967295
        if len(v4) >= 2 {
            v4 = v4[1:len(v4)]
        }
        _ = fi6
        _ = v8
    }
    _ = p0
    _ = v4
    _ = v5
    return (("foo bar" + "a") + (("hello" + ("" + ("" + (("" + "foo bar") + (("\n\t" + ("foo bar" + (("a" + "") + "hello"))) + "foo bar"))))) + "a"))
}
func genFn1(p0 uint32) int16 {
    v9 := GenS2{F0: 2147483647, F1: 0}
    if ((((255 != 32767) && false) && false) || true) {
        v10 := GenS1{F0: 1000, F1: true, F2: (p0 - p0)}
        _ = v10
    } else {
        var v11 float64 = 0.0
        v11 = v11
        if (0 == -1) {
            p0 -= (^(p0))
            var v12 int16 = -445
            _ = v12
        } else {
            v13 := GenS2{F0: 384, F1: 255}
            v11 = (0.0 / v11)
            p0 = (p0 % ((p0) | 1))
            _ = v13
        }
        _ = v11
    }
    _ = p0
    _ = v9
    return 1000
}

func main() {
    v1 := make(map[uint32]int)
    var v2 int32 = 0
    v3 := []int{642, 1, 128}
    var v14 int64 = -9223372036854775808
    if false {
        _ = len(v3)
        for fi15 := 0; fi15 < 14; fi15++ {
            if len(v3) >= 2 {
                v3 = v3[1:len(v3)]
            }
            _ = fi15
        }
        for fi16 := 0; fi16 < 15; fi16++ {
            for fi17 := 0; fi17 < 14; fi17++ {
                var v18 int64 = (v14 ^ v14)
                _ = fi17
                _ = v18
            }
            for fi19 := 0; fi19 < 13; fi19++ {
                var v20 float32 = (0.0 * ((float32((-2.25 + 1e10)) * (((-2.25 * -2.25) + 1.5) + -2.25)) * 3.14159))
                fmt.Println(fi16, fi16)
                var v21 bool = false
                _ = fi19
                _ = v20
                _ = v21
            }
            _ = fi16
        }
    }
    v2 += v2
    v22 := GenS1{F0: -796, F1: false, F2: 2147483648}
    v1[2147483648] = 1
    fmt.Println(v2, v2)
    var v23 uint16 = 0
    var v24 int8 = -1
    switch 331 {
    case -9223372036854775808:
        var v25 int = 2147483647
        _ = v25
    case 340:
        v3 = append(v3, 9223372036854775807)
    default:
        var v26 uint = 4294967295
        _ = v26
    }
    fmt.Println(v24, v2)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("v23:", v23)
    fmt.Println("v24:", v24)
    fmt.Println("slicelen v3:", len(v3), v3)
    _ = v1
    _ = v14
    _ = v2
    _ = v22
    _ = v23
    _ = v24
    _ = v3
}
