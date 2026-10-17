package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 int64; F2 uint32 }
type GenS2 struct { F0 bool; F1 uint; F2 uint }
func genFn0(p0 int64) int {
    var v6 uint32 = 2147483647
    switch -867 {
    case -505:
        fmt.Println(v6, p0)
    default:
        var v7 int8 = 76
        _ = v7
    }
    v8 := GenS1{F0: ("a" + "hello"), F1: p0, F2: (v6 >> ((985) & 63))}
    var v9 int8 = -1
    _ = p0
    _ = v6
    _ = v8
    _ = v9
    return -302
}
func genFn1(p0 uint16) uint32 {
    for fi10 := 0; fi10 < 19; fi10++ {
        fmt.Println(p0, fi10)
        p0 = 65535
        _ = fi10
    }
    if (!(193 <= 1)) {
        v11 := GenS2{F0: true, F1: 0, F2: 65535}
        _ = v11
    }
    switch 0 {
    case 745:
        fmt.Println(p0)
    case -675:
        v12 := GenS1{F0: ("" + (("hello" + "hello") + ("hello" + ""))), F1: -1, F2: 1}
        _ = v12
    default:
        if ((!(32768 <= 32768)) || false) {
            v13 := []int{1, genFn0(-9223372036854775808)}
            p0 = (p0 | (p0 % ((p0) | 1)))
            _ = v13
        } else {
            switch genFn0(1) {
            case -2147483648:
                v14 := GenS2{F0: (!false), F1: 596, F2: 1}
                _ = v14
            case 830:
                v15 := []int{-2147483648, 563, 2147483647, 124}
                _ = v15
            default:
                if (!(false && false)) {
                    var v16 uint16 = (^(p0))
                    fmt.Println(p0, v16)
                    var v17 float64 = -2.25
                    _ = v16
                    _ = v17
                }
            }
            p0 = p0
        }
    }
    fmt.Println(p0)
    _ = p0
    return 550
}

func main() {
    var v1 int16 = 1000
    var v2 int = 1
    v3 := make(map[string]uint32)
    v4 := make(map[uint32]string)
    var v5 uint8 = 42
    for fi18 := 0; fi18 < 17; fi18++ {
        for fi19 := 0; fi19 < 18; fi19++ {
            if ((4294967295 <= 609) && true) {
                var v20 int16 = 993
                _ = v20
            } else {
                v5 -= v5
                var v21 string = "\n\t"
                var v22 int16 = v1
                _ = v21
                _ = v22
            }
            _ = fi19
        }
        _ = fi18
    }
    v23 := []int{(v2 << ((65535) & 63)), (v2 - (v2 + (-(v2))))}
    var v24 int32 = -810
    switch "" {
    case "a":
        delete(v4, 8)
    case "\n\t":
        v24 -= v24
    case "":
        var v25 string = "hello"
        _ = v25
    default:
        v26 := GenS1{F0: ("hello" + ("\n\t" + ((("" + "\n\t") + "") + (("foo bar" + "foo bar") + "")))), F1: 9223372036854775807, F2: 65535}
        _ = v26
    }
    fmt.Println(v1, v5, v24)
    switch "" {
    case "\n\t":
        v2 = v2
    default:
        var v27 int8 = 100
        _ = v27
    }
    v5 -= v5
    v28 := GenS1{F0: "", F1: 42, F2: 4294967295}
    fmt.Println(v24)
    v1 = v1
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v23:", len(v23), v23)
    fmt.Println("v24:", v24)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v23
    _ = v24
    _ = v28
    _ = v3
    _ = v4
    _ = v5
}
