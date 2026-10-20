package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 bool; F2 uint8 }
type GenS2 struct { F0 uint; F1 string; F2 uint16 }
func genFn0() uint32 {
    if (!false) {
        v5 := GenS1{F0: 620, F1: true, F2: 2}
        _ = v5
    } else {
        v6 := GenS1{F0: 0, F1: (true || (18446744073709551615 >= 18446744073709551615)), F2: 42}
        _ = v6
    }
    var v7 int64 = -9223372036854775808
    _ = v7
    return 42
}

func main() {
    var v1 int16 = -7
    v2 := []uint32{0, 2147483647, 42, 65535}
    var v3 float32 = 0.0
    v4 := make(map[uint32]bool)
    v1 -= v1
    _ = len(v4)
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    for fi8 := 0; fi8 < 7; fi8++ {
        for fi9 := 0; fi9 < 17; fi9++ {
            fmt.Println(fi8, v1, fi8)
            v10 := GenS1{F0: v1, F1: false, F2: 42}
            v1 = (v1 & (v1 << ((659) & 63)))
            v11 := GenS1{F0: (-(v1)), F1: false, F2: 44}
            _ = fi9
            _ = v10
            _ = v11
        }
        _ = fi8
    }
    v12 := GenS2{F0: 4294967295, F1: (("hello" + "foo bar") + "hello"), F2: 419}
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    var v13 int32 = 2147483647
    fmt.Println(v3, v1)
    var v14 uint32 = 2147483648
    var v15 int64 = 9223372036854775807
    switch "hello" {
    case "":
        if (0 > v13) {
            fmt.Println(v14)
            switch 0 {
            case 0:
                v16 := GenS2{F0: 65535, F1: "foo bar", F2: 2}
                _ = v16
            case 743:
                fmt.Println(v15)
            default:
                var v17 float32 = (1e10 - (v3 + (v3 / v3)))
                _ = v17
            }
        }
    case "\n\t":
        if false {
            v18 := GenS1{F0: (v1 | v1), F1: true, F2: 0}
            _ = v18
        } else {
            v3 -= v3
        }
    default:
        fmt.Println(v3)
    }
    switch ("\n\t" + "a") {
    case "hello":
        v15 = (^(v15))
    case "\n\t":
        v15 -= (v15 % ((v15) | 1))
    default:
        v19 := GenS2{F0: 42, F1: ("foo bar" + "\n\t"), F2: 905}
        _ = v19
    }
    v1 += 1
    v13 = v13
    var v20 float32 = v3
    if ((-(v1)) > v1) {
        v21 := GenS1{F0: (v1 / (((v1 | (v1 + (v1 - 0)))) | 1)), F1: ("hello" > "a"), F2: 128}
        _ = v21
    } else {
        var v22 float64 = 0.1
        v23 := GenS1{F0: -708, F1: (-9223372036854775808 < 2147483647), F2: uint8(v14)}
        if true {
            if true {
                v13 = v13
            }
            for fi24 := 0; fi24 < 3; fi24++ {
                var v25 uint16 = 271
                var v26 float32 = (v3 / (((v3 * 1e10) * ((v20 / v3) + 1e10)) * (v20 / 1.5)))
                if true {
                    for fi27 := 0; fi27 < 18; fi27++ {
                        v26 -= 3.14159
                        _ = fi27
                    }
                    if len(v2) >= 2 {
                        v2 = v2[1:len(v2)]
                    }
                    v28 := v4[675]
                    _ = v28
                }
                _ = fi24
                _ = v25
                _ = v26
            }
        } else {
            v29 := v4[2]
            _ = v29
        }
        _ = v22
        _ = v23
    }
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v20:", v20)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v12
    _ = v13
    _ = v14
    _ = v15
    _ = v2
    _ = v20
    _ = v3
    _ = v4
}
