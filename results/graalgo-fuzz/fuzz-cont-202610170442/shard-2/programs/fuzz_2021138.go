package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 int8 }
func genFn0(p0 uint64) string {
    p0 = (p0 % (((p0 - (^(p0)))) | 1))
    fmt.Println(p0)
    _ = p0
    return ("\n\t" + "")
}
func genFn1(p0 int16, p1 uint32) uint {
    var v7 float32 = 0.1
    p1 -= (p1 & p1)
    _ = p0
    _ = p1
    _ = v7
    return 858
}

func main() {
    v1 := make(map[uint32]bool)
    var v2 uint64 = 18446744073709551615
    var v3 int8 = 1
    v4 := []string{"", "a", "foo bar"}
    v5 := make(map[string]int)
    v6 := make(map[string]bool)
    v3 = (v3 * v3)
    if (42 < -2147483648) {
        v3 = (v3 << ((0) & 63))
    }
    v2 = (v2 + (^(v2)))
    if len(v4) >= 2 {
        v4 = v4[1:len(v4)]
    }
    switch "foo bar" {
    case "foo bar":
        delete(v6, ("hello" + "foo bar"))
    case "a":
        var v8 uint8 = 15
        _ = v8
    case "\n\t":
        fmt.Println(v3)
    default:
        for fi9 := 0; fi9 < 1; fi9++ {
            var v10 int32 = -2147483648
            _ = fi9
            _ = v10
        }
    }
    v3 -= (v3 ^ v3)
    var v11 uint32 = 2147483648
    if (-1 <= 2147483647) {
        v3 = v3
    }
    if (v2 <= 294) {
        v3 += (-(v3))
    } else {
        v3 = (v3 % ((v3) | 1))
    }
    var v12 uint16 = 115
    var v13 uint16 = (v12 + (v12 ^ (v12 + v12)))
    if (!(1.5 >= (3.14159 * 3.14159))) {
        var v14 int = 32767
        _ = v14
    } else {
        v15 := GenS1{F0: 94, F1: (v3 | (v3 >> ((0) & 63)))}
        switch "a" {
        case "":
            v6[("a" + "a")] = (!true)
        case "a":
            fmt.Println(v3, v13)
        }
        if (false && (!false)) {
            fmt.Println(v2, v12)
            v4 = append(v4, ("\n\t" + "\n\t"))
            v5[("" + (("a" + "foo bar") + ""))] = 9223372036854775807
        } else {
            v16 := GenS1{F0: 207, F1: 7}
            var v17 uint32 = 4294967295
            var v18 uint64 = (v2 + (v2 + v2))
            _ = v16
            _ = v17
            _ = v18
        }
        _ = v15
    }
    v2 = (v2 / ((v2) | 1))
    switch "a" {
    case "":
        var v19 uint16 = v12
        _ = v19
    case "\n\t":
        for fi20 := 0; fi20 < 7; fi20++ {
            switch 32767 {
            case -526:
                if (42 >= 42) {
                    fmt.Println(v11, v2)
                } else {
                    var v21 float64 = float64(1.5)
                    _ = v21
                }
            }
            v22 := GenS1{F0: 2, F1: (v3 % (((v3 % ((v3) | 1))) | 1))}
            _ = fi20
            _ = v22
        }
    default:
        for fi23 := 0; fi23 < 13; fi23++ {
            if len(v4) > 0 {
                v4[0] = ("" + "a")
            }
            _ = fi23
        }
    }
    var v24 int64 = 9223372036854775807
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v2:", v2)
    fmt.Println("v24:", v24)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v11
    _ = v12
    _ = v13
    _ = v2
    _ = v24
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
