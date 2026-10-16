package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
type GenS2 struct { F0 string; F1 bool; F2 uint8 }
type GenS3 struct { F0 int8; F1 int16; F2 string }

func main() {
    v1 := make(map[int]string)
    var v2 int64 = 9223372036854775807
    var v3 bool = ((v2 >> ((18446744073709551615) & 63)) > 633)
    v4 := []uint32{0}
    var v5 uint64 = 277
    v6 := GenS1{F0: -15}
    for fi7 := 0; fi7 < 4; fi7++ {
        if true {
            for fi8 := 0; fi8 < 5; fi8++ {
                v9 := GenS1{F0: 0}
                var v10 uint16 = 0
                if (0 < 42) {
                    _ = len(v1)
                    fmt.Println(v10, v10)
                    v3 = (!true)
                }
                _ = fi8
                _ = v10
                _ = v9
            }
            v5 = (v5 >> ((261) & 63))
            delete(v1, (fi7 - (fi7 / (((fi7 << ((984) & 63))) | 1))))
        }
        _ = fi7
    }
    if len(v4) >= 2 {
        v4 = v4[1:len(v4)]
    }
    v5 = v5
    switch "foo bar" {
    case "foo bar":
        fmt.Println(v2)
    case "\n\t":
        var v11 string = "foo bar"
        _ = v11
    case "hello":
        var v12 int = -891
        _ = v12
    default:
        v2 = 9223372036854775807
    }
    v3 = (322 != -9223372036854775808)
    var v13 int16 = -32768
    fmt.Println(v2, v13)
    v5 = v5
    if len(v4) > 0 {
        v4[0] = 42
    }
    var v14 int64 = v2
    fmt.Println(v3)
    for fi15 := 0; fi15 < 12; fi15++ {
        if ((3.14159 / (float64((((0.0 / 3.14159) + -2.25) * 1e10)) + -2.25)) == -2.25) {
            switch "\n\t" {
            case "":
                if len(v4) >= 2 {
                    v4 = v4[1:len(v4)]
                }
            case "foo bar":
                var v16 int = fi15
                _ = v16
            default:
                if v3 {
                    v13 = v13
                    v17 := GenS2{F0: ("foo bar" + "\n\t"), F1: ((42 == 1) || true), F2: 127}
                    var v18 uint32 = 809
                    _ = v17
                    _ = v18
                }
            }
            fmt.Println(fi15, v13)
        }
        v1[fi15] = ""
        v19 := GenS3{F0: 1, F1: -32768, F2: ("a" + ("" + (("" + (("" + ("\n\t" + "foo bar")) + ("foo bar" + ("" + "foo bar")))) + "")))}
        _ = fi15
        _ = v19
    }
    fmt.Println(v3, v3)
    v14 = v2
    v2 = (v14 + v14)
    v3 = (false && (2147483647 <= 9223372036854775807))
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v13
    _ = v14
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
