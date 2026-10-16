package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 uint64 }
type GenS2 struct { F0 int16; F1 uint }

func main() {
    v1 := make(map[string]uint32)
    v2 := []uint32{0, 2, 0}
    v3 := make(map[string]uint32)
    var v4 uint = 4294967295
    var v5 int64 = -805
    v6 := make(map[uint32]bool)
    for fi7 := 0; fi7 < 7; fi7++ {
        var v8 int = (fi7 | fi7)
        _ = fi7
        _ = v8
    }
    v5 -= (-(v5))
    v5 += (v5 * 0)
    if (0 < -1) {
        switch "a" {
        case "":
            fmt.Println(v5)
        case "\n\t":
            var v9 string = ("hello" + ((("\n\t" + "a") + "a") + ""))
            _ = v9
        case "a":
            var v10 uint16 = 1
            _ = v10
        default:
            v4 -= (v4 * 4294967295)
        }
        v11 := v1["\n\t"]
        var v12 int32 = 2147483647
        _ = v11
        _ = v12
    }
    for fi13 := 0; fi13 < 1; fi13++ {
        var v14 int32 = 205
        _ = fi13
        _ = v14
    }
    var v15 uint8 = 255
    fmt.Println(v15)
    fmt.Println(v15)
    v5 += (^(v5))
    v4 += (^(v4))
    v15 -= 116
    for fi16 := 0; fi16 < 7; fi16++ {
        fmt.Println(v5, v15)
        for fi17 := 0; fi17 < 9; fi17++ {
            fmt.Println(v15)
            _ = fi17
        }
        fmt.Println(v15, fi16, v4)
        _ = fi16
    }
    fmt.Println(v15, v15)
    switch ("" + (("\n\t" + "a") + "\n\t")) {
    case "hello":
        fmt.Println(v5, v15)
    case "foo bar":
        if (-2.25 > ((float64((1.5 / -2.25)) * (0.1 + 0.1)) / float64((0.1 / 1.0)))) {
            switch 2147483647 {
            case 2147483647:
                var v18 uint16 = 65535
                _ = v18
            default:
                for fi19 := 0; fi19 < 8; fi19++ {
                    for fi20 := 0; fi20 < 10; fi20++ {
                        v15 = v15
                        _ = fi20
                    }
                    _ = fi19
                }
            }
        } else {
            v4 -= (v4 | v4)
            switch "" {
            case "\n\t":
                v21 := GenS1{F0: 2147483648, F1: 497}
                _ = v21
            case "hello":
                for fi22 := 0; fi22 < 6; fi22++ {
                    v23 := GenS2{F0: 32767, F1: v4}
                    v4 += (v4 / ((255) | 1))
                    _ = fi22
                    _ = v23
                }
            case "a":
                var v24 uint64 = uint64(v4)
                _ = v24
            }
        }
    case "":
        v3["foo bar"] = 187
    default:
        var v25 bool = true
        _ = v25
    }
    v26 := GenS1{F0: 894, F1: 18446744073709551615}
    for fi27 := 0; fi27 < 9; fi27++ {
        var v28 int64 = (v5 & (v5 + 9223372036854775807))
        fmt.Println(v15)
        var v29 int32 = 2147483647
        _ = fi27
        _ = v28
        _ = v29
    }
    var v30 uint32 = 908
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v15:", v15)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v30:", v30)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v15
    _ = v2
    _ = v26
    _ = v3
    _ = v30
    _ = v4
    _ = v5
    _ = v6
}
