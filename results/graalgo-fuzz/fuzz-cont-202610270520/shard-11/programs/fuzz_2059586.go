package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 int64 }
type GenS2 struct { F0 uint; F1 int64; F2 uint16 }
type GenS3 struct { F0 int }
type GenS4 struct { F0 bool }

func main() {
    var v1 uint8 = 72
    var v2 float32 = (1.5 / 1e10)
    v3 := []string{}
    var v4 uint16 = 65535
    for fi5 := 0; fi5 < 7; fi5++ {
        fmt.Println(v1, fi5, fi5)
        v6 := GenS1{F0: (false && true), F1: 297}
        switch (fi5 >> ((4294967295) & 63)) {
        case 32768:
            var v7 int32 = -2147483648
            _ = v7
        }
        _ = fi5
        _ = v6
    }
    var v8 uint = 65535
    for fi9 := 0; fi9 < 6; fi9++ {
        switch (((("foo bar" + "foo bar") + "a") + "") + "foo bar") {
        case "":
            if (("a" > "\n\t") && false) {
                v10 := GenS1{F0: false, F1: -1}
                _ = v10
            }
        case "foo bar":
            v8 += (^(v8))
        default:
            if (((v8 * v8) != (v8 / ((v8) | 1))) && (v2 == 1e10)) {
                if len(v3) >= 2 {
                    v3 = v3[1:len(v3)]
                }
                switch 0 {
                case 42:
                    fmt.Println(v2, v2)
                case -1:
                    if false {
                        v11 := GenS1{F0: ((-1 > -9223372036854775808) || (false && true)), F1: 42}
                        fmt.Println(v8, fi9, fi9)
                        v12 := GenS2{F0: v8, F1: 42, F2: 494}
                        _ = v11
                        _ = v12
                    }
                case 2147483647:
                    var v13 uint64 = 42
                    _ = v13
                }
                v3 = append(v3, ("hello" + ""))
            } else {
                if (1 == 1) {
                    v3 = append(v3, "\n\t")
                    for fi14 := 0; fi14 < 13; fi14++ {
                        switch "\n\t" {
                        case "\n\t":
                            v15 := GenS1{F0: true, F1: 1}
                            _ = v15
                        }
                        var v16 string = "a"
                        _ = fi14
                        _ = v16
                    }
                    v1 = v1
                }
            }
        }
        _ = fi9
    }
    fmt.Println(v4, v8, v8)
    if (448 < -1) {
        if false {
            if (int64(v1) > 310) {
                v1 += (v1 * v1)
                if (-676 > -1) {
                    var v17 uint8 = uint8(v4)
                    _ = v17
                }
                fmt.Println(v4, v1, v2)
            }
            v18 := GenS2{F0: (v8 & (^(v8))), F1: 211, F2: v4}
            _ = v18
        }
        var v19 int = 229
        switch v19 {
        case -9223372036854775808:
            var v20 int32 = -1
            _ = v20
        default:
            v21 := GenS3{F0: v19}
            _ = v21
        }
        _ = v19
    }
    v22 := GenS2{F0: v8, F1: -159, F2: (v4 & (^(v4)))}
    v1 = 0
    v1 = 1
    var v23 float64 = (1.5 - (0.1 - (-2.25 / (0.0 + float64(v2)))))
    var v24 uint8 = v1
    _ = len(v3)
    for fi25 := 0; fi25 < 5; fi25++ {
        if (((v1 >> (((^(v8))) & 63)) >= (v1 | v1)) && ((!(false || false)) && (true || (2147483647 == -343)))) {
            if len(v3) >= 2 {
                v3 = v3[1:len(v3)]
            }
            for fi26 := 0; fi26 < 3; fi26++ {
                if len(v3) > 0 {
                    v3[0] = ("\n\t" + "foo bar")
                }
                v27 := GenS2{F0: 4294967295, F1: -9223372036854775808, F2: (^(v4))}
                if (!(409 < 1000)) {
                    v28 := GenS4{F0: false}
                    _ = v28
                } else {
                    v8 -= v8
                    v24 += (^(v1))
                    v24 = v1
                    if (!false) {
                        var v29 uint64 = 18446744073709551615
                        if len(v3) >= 2 {
                            v3 = v3[1:len(v3)]
                        }
                        var v30 uint = v8
                        _ = v29
                        _ = v30
                    } else {
                        v4 += (v4 >> (((v8 * (v8 * 1))) & 63))
                        v31 := GenS2{F0: v8, F1: -423, F2: v4}
                        _ = v31
                    }
                }
                _ = fi26
                _ = v27
            }
        } else {
            switch "" {
            case "foo bar":
                var v32 uint8 = 124
                _ = v32
            case "a":
                switch "" {
                case "hello":
                    v24 = 2
                }
            default:
                v2 = float32(((1.5 - float64(-2.25)) / v23))
            }
            if (993 < uint64(v24)) {
                var v33 int = int(v4)
                v33 = (^(v33))
                _ = v33
            }
        }
        _ = fi25
    }
    v23 = ((((0.1 * v23) * float64(1.5)) - 3.14159) / (float64(1.5) - v23))
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v23:", v23)
    fmt.Println("v24:", v24)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v2
    _ = v22
    _ = v23
    _ = v24
    _ = v3
    _ = v4
    _ = v8
}
