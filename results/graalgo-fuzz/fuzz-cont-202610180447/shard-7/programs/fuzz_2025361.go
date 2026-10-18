package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 string }
type GenS2 struct { F0 bool }

func main() {
    var v1 int16 = 32767
    v2 := []int{-271}
    v3 := []string{}
    v4 := []uint32{672, 21}
    var v5 float64 = -2.25
    _ = len(v3)
    if (593 != 2147483647) {
        _ = len(v3)
        if (v1 < 0) {
            var v6 uint = 4294967295
            if ((false && true) && ((100 > 104) || true)) {
                if (true || (!(0 == -2147483648))) {
                    v1 -= v1
                    v7 := GenS1{F0: (v1 ^ (-(v1))), F1: ("foo bar" + "a")}
                    _ = v7
                }
                switch 32768 {
                case 0:
                    var v8 uint64 = 873
                    _ = v8
                default:
                    fmt.Println(v1, v6)
                }
                if false {
                    switch "\n\t" {
                    case "":
                        var v9 float32 = (float32(v5) / (0.1 + 1e10))
                        _ = v9
                    }
                    var v10 uint = v6
                    if (32767 <= 1) {
                        v5 += (1.5 * (v5 * (((v5 / v5) / ((0.0 * 3.14159) - (v5 - (-2.25 / v5)))) * (v5 * ((0.0 / v5) / 1e10)))))
                        v11 := GenS2{F0: (!(-64 < 42))}
                        v10 = 18446744073709551615
                        _ = v11
                    }
                    _ = v10
                }
            } else {
                if len(v3) >= 2 {
                    v3 = v3[1:len(v3)]
                }
                var v12 int8 = 0
                _ = v12
            }
            _ = v6
        }
        v5 -= v5
    }
    var v13 int8 = 42
    for fi14 := 0; fi14 < 15; fi14++ {
        fmt.Println(v1, v5)
        for fi15 := 0; fi15 < 20; fi15++ {
            v13 += 0
            _ = fi15
        }
        _ = fi14
    }
    if true {
        _ = len(v2)
        if true {
            if (!(false && ((v13 | v13) < 127))) {
                v2 = append(v2, -288)
                fmt.Println(v5, v13)
                fmt.Println(v1, v1, v5)
            }
        }
    } else {
        var v16 int64 = 1
        v17 := GenS1{F0: v1, F1: "foo bar"}
        var v18 uint32 = 2147483648
        _ = v16
        _ = v17
        _ = v18
    }
    var v19 bool = (v13 > v13)
    for fi20 := 0; fi20 < 8; fi20++ {
        switch 2147483647 {
        case 9223372036854775807:
            var v21 uint = 231
            _ = v21
        default:
            v5 -= v5
        }
        var v22 float32 = float32(3.14159)
        _ = len(v4)
        _ = fi20
        _ = v22
    }
    switch "a" {
    case "\n\t":
        switch "" {
        case "\n\t":
            for fi23 := 0; fi23 < 5; fi23++ {
                if len(v4) > 0 {
                    v4[0] = 42
                }
                if len(v2) >= 2 {
                    v2 = v2[1:len(v2)]
                }
                v1 += (v1 % ((v1) | 1))
                _ = fi23
            }
        default:
            fmt.Println(v13)
        }
    case "foo bar":
        var v24 int = -1
        _ = v24
    case "a":
        var v25 uint16 = 181
        _ = v25
    }
    v13 = (v13 - v13)
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    switch 32767 {
    case 128:
        if (0.1 >= -2.25) {
            for fi26 := 0; fi26 < 13; fi26++ {
                v19 = false
                v27 := GenS2{F0: (v13 <= v13)}
                v13 -= 81
                _ = fi26
                _ = v27
            }
        }
    case 127:
        var v28 float32 = 1e10
        _ = v28
    case 0:
        if false {
            var v29 uint16 = 42
            v30 := GenS1{F0: 32767, F1: ""}
            _ = v29
            _ = v30
        } else {
            v1 -= v1
            v13 -= 97
            var v31 uint = 1
            _ = v31
        }
    }
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    _ = len(v3)
    fmt.Println(v19, v19)
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("v19:", v19)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v13
    _ = v19
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
