package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 int64; F2 uint32 }
type GenS2 struct { F0 int8 }

func main() {
    var v1 float32 = float32(-2.25)
    var v2 float64 = 1.5
    v3 := []string{("\n\t" + "\n\t"), (("\n\t" + ("" + "hello")) + (("foo bar" + ("a" + ("\n\t" + ""))) + "hello"))}
    v4 := make(map[uint32]string)
    var v5 int64 = 42
    var v6 uint16 = 32767
    switch 32768 {
    case 32768:
        var v7 float64 = (1e10 + v2)
        _ = v7
    case -1:
        switch ("hello" + "foo bar") {
        case "\n\t":
            fmt.Println(v1, v6, v5)
        case "a":
            var v8 uint8 = 2
            _ = v8
        default:
            var v9 int32 = 42
            _ = v9
        }
    case 0:
        v5 = (v5 * (v5 >> ((65535) & 63)))
    default:
        switch -930 {
        case 899:
            switch "foo bar" {
            case "foo bar":
                var v10 bool = true
                _ = v10
            case "a":
                delete(v4, 53)
            case "hello":
                for fi11 := 0; fi11 < 18; fi11++ {
                    if false {
                        var v12 int64 = v5
                        _ = v12
                    }
                    _ = fi11
                }
            default:
                fmt.Println(v6, v2, v6)
            }
        case 1:
            for fi13 := 0; fi13 < 13; fi13++ {
                var v14 bool = (!false)
                v6 = 1
                v1 = -2.25
                _ = fi13
                _ = v14
            }
        }
    }
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    if (-1 <= -1) {
        if (4294967295 < 255) {
            if false {
                v6 = (v6 * v6)
            } else {
                for fi15 := 0; fi15 < 3; fi15++ {
                    fmt.Println(v5)
                    for fi16 := 0; fi16 < 6; fi16++ {
                        v17 := GenS1{F0: false, F1: (^(v5)), F2: 1}
                        v5 = (v5 - (-(v5)))
                        _ = fi16
                        _ = v17
                    }
                    v5 -= (v5 ^ 1)
                    _ = fi15
                }
                var v18 uint16 = v6
                fmt.Println(v6, v18, v18)
                _ = v18
            }
            if len(v3) >= 2 {
                v3 = v3[1:len(v3)]
            }
        }
        v6 -= v6
    } else {
        v2 -= ((((0.1 + (v2 - ((float64(v1) - v2) / float64(v1)))) / float64(1.5)) / (1.5 / (v2 - 0.1))) * (((((v2 + float64((v1 * (v1 * (0.1 - v1))))) + 1e10) - v2) / 1.0) / (float64(float32(v2)) - ((1e10 / v2) * (v2 / v2)))))
        var v19 int64 = 9223372036854775807
        _ = v19
    }
    if (2 > 4294967295) {
        _ = len(v3)
        v2 -= float64(v1)
    }
    v1 = (v1 + v1)
    var v20 uint64 = 4294967295
    for fi21 := 0; fi21 < 14; fi21++ {
        for fi22 := 0; fi22 < 10; fi22++ {
            if false {
                var v23 uint16 = (v6 & (v6 << ((2) & 63)))
                var v24 uint64 = v20
                _ = len(v3)
                _ = v23
                _ = v24
            }
            var v25 int32 = -2147483648
            _ = fi22
            _ = v25
        }
        v20 += 18446744073709551615
        _ = fi21
    }
    for fi26 := 0; fi26 < 14; fi26++ {
        var v27 int = 2147483647
        if false {
            var v28 int16 = 744
            v1 = float32(1.5)
            v29 := GenS1{F0: (v28 == v28), F1: v5, F2: 2147483647}
            _ = v28
            _ = v29
        } else {
            if (false && false) {
                _ = len(v4)
                v30 := GenS2{F0: 100}
                _ = v30
            } else {
                for fi31 := 0; fi31 < 19; fi31++ {
                    _ = len(v3)
                    _ = len(v4)
                    _ = fi31
                }
                v6 = (^(v6))
                if false {
                    for fi32 := 0; fi32 < 14; fi32++ {
                        switch int(v20) {
                        case 424:
                            var v33 float32 = v1
                            delete(v4, 65535)
                            _ = len(v3)
                            _ = v33
                        default:
                            v5 -= (-(v5))
                        }
                        _ = fi32
                    }
                    var v34 bool = (-401 != -180)
                    fmt.Println(v6, v6, v34)
                    _ = v34
                } else {
                    v5 = (v5 ^ (v5 % (((v5 >> ((108) & 63))) | 1)))
                    var v35 int = (v27 % (((v27 / ((fi26) | 1))) | 1))
                    v6 = (v6 / (((^(v6))) | 1))
                    _ = v35
                }
            }
            v36 := GenS2{F0: -128}
            for fi37 := 0; fi37 < 14; fi37++ {
                fmt.Println(v27, v1)
                _ = fi37
            }
            _ = v36
        }
        _ = fi26
        _ = v27
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v20
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
