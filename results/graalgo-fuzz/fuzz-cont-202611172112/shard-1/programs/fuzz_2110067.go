package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint32 }
type GenS2 struct { F0 uint32; F1 bool }
type GenS3 struct { F0 int16; F1 int; F2 bool }
func genFn0(p0 int8) uint32 {
    for fi5 := 0; fi5 < 13; fi5++ {
        v6 := []int{(-(fi5)), fi5, fi5}
        for fi7 := 0; fi7 < 14; fi7++ {
            var v8 int16 = 42
            p0 -= (p0 % ((p0) | 1))
            var v9 int8 = (p0 & 42)
            _ = fi7
            _ = v8
            _ = v9
        }
        _ = fi5
        _ = v6
    }
    switch -1 {
    case 166:
        var v10 uint = 18446744073709551615
        _ = v10
    default:
        v11 := []int{-2147483648, 32767}
        _ = v11
    }
    fmt.Println(p0)
    var v12 uint = 2
    _ = p0
    _ = v12
    return 0
}

func main() {
    var v1 uint16 = 65535
    var v2 uint8 = 255
    var v3 uint16 = v1
    var v4 float32 = 0.0
    v13 := []int{9223372036854775807, -9223372036854775808, 889}
    var v14 uint32 = 0
    if false {
        if false {
            v15 := GenS1{F0: (^(v1)), F1: v14}
            for fi16 := 0; fi16 < 9; fi16++ {
                switch fi16 {
                case 9223372036854775807:
                    v3 -= v1
                default:
                    switch "hello" {
                    case "a":
                        for fi17 := 0; fi17 < 9; fi17++ {
                            var v18 int = fi17
                            _ = fi17
                            _ = v18
                        }
                    case "foo bar":
                        v19 := GenS2{F0: (v14 << ((255) & 63)), F1: (v3 != v3)}
                        _ = v19
                    }
                }
                for fi20 := 0; fi20 < 12; fi20++ {
                    switch ("hello" + ("a" + "")) {
                    case "a":
                        v3 += v3
                    default:
                        v13 = append(v13, -117)
                    }
                    var v21 int64 = 882
                    v22 := GenS2{F0: (v14 ^ v14), F1: (!true)}
                    _ = fi20
                    _ = v21
                    _ = v22
                }
                v23 := GenS1{F0: 42, F1: genFn0(106)}
                _ = fi16
                _ = v23
            }
            switch "" {
            case "a":
                var v24 int64 = 271
                _ = v24
            case "foo bar":
                _ = len(v13)
            }
            _ = v15
        }
        if (2147483647 >= -2147483648) {
            switch ("foo bar" + ("foo bar" + "a")) {
            case "a":
                v1 = 65535
            }
        }
    } else {
        v1 += v1
        if true {
            var v25 float32 = v4
            var v26 bool = (!false)
            _ = v25
            _ = v26
        } else {
            var v27 uint16 = (v3 - v3)
            _ = v27
        }
        v28 := GenS1{F0: (v3 ^ (v1 % ((v3) | 1))), F1: v14}
        _ = v28
    }
    for fi29 := 0; fi29 < 8; fi29++ {
        if len(v13) > 0 {
            v13[0] = (fi29 & fi29)
        }
        v2 += 128
        for fi30 := 0; fi30 < 15; fi30++ {
            fmt.Println(fi30)
            v31 := GenS1{F0: 0, F1: v14}
            _ = fi30
            _ = v31
        }
        _ = fi29
    }
    if (false && true) {
        v4 = ((1e10 - (0.0 - v4)) + float32(float64(v4)))
        if (uint8(v3) < v2) {
            v14 = genFn0(-128)
        }
    }
    v4 = (float32(1.5) + (1e10 - v4))
    v32 := GenS3{F0: 42, F1: 42, F2: false}
    v14 += 442
    v33 := GenS2{F0: 230, F1: false}
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v13:", len(v13), v13)
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v13
    _ = v14
    _ = v2
    _ = v3
    _ = v32
    _ = v33
    _ = v4
}
