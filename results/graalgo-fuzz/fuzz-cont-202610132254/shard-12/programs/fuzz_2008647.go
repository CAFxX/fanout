package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 uint32; F2 bool }
type GenS2 struct { F0 uint64; F1 bool; F2 uint8 }
type GenS3 struct { F0 string; F1 int32; F2 uint }
type GenS4 struct { F0 int; F1 int8 }
type GenS5 struct { F0 uint }

func main() {
    v1 := []string{"foo bar", ((("foo bar" + ("foo bar" + "hello")) + "") + ("" + ((("\n\t" + "") + "hello") + "\n\t"))), "a"}
    v2 := make(map[uint32]bool)
    var v3 bool = (!true)
    var v4 uint = 42
    v5 := []uint32{639}
    if true {
        if true {
            v1 = append(v1, "hello")
            switch "hello" {
            case "\n\t":
                v4 = (^(v4))
            case "a":
                v3 = (52 <= 614)
            case "":
                if len(v1) >= 2 {
                    v1 = v1[1:len(v1)]
                }
            default:
                switch -2147483648 {
                case 9223372036854775807:
                    var v6 int = 491
                    _ = v6
                case -9223372036854775808:
                    if v3 {
                        v7 := GenS1{F0: 18446744073709551615, F1: 845, F2: (127 != 1)}
                        for fi8 := 0; fi8 < 18; fi8++ {
                            v9 := GenS1{F0: 42, F1: 4294967295, F2: v3}
                            if (false && (v3 || v3)) {
                                var v10 float32 = (((0.1 * (((1.5 + 0.0) - 3.14159) * -2.25)) - ((float32(-2.25) / -2.25) * (0.1 / float32(0.1)))) * (3.14159 + 1.5))
                                _ = v10
                            }
                            _ = fi8
                            _ = v9
                        }
                        v4 -= v4
                        _ = v7
                    } else {
                        v11 := v2[1]
                        v12 := GenS1{F0: 18446744073709551615, F1: 4294967295, F2: false}
                        v11 = true
                        _ = v11
                        _ = v12
                    }
                case 4294967295:
                    var v13 float32 = 3.14159
                    _ = v13
                default:
                    for fi14 := 0; fi14 < 5; fi14++ {
                        fmt.Println(v3, v4)
                        v15 := GenS2{F0: 210, F1: v3, F2: 128}
                        _ = fi14
                        _ = v15
                    }
                }
            }
        }
    }
    switch "\n\t" {
    case "":
        v16 := GenS3{F0: "a", F1: 445, F2: (v4 / ((v4) | 1))}
        _ = v16
    case "\n\t":
        v17 := GenS4{F0: -2147483648, F1: -128}
        _ = v17
    case "a":
        if false {
            for fi18 := 0; fi18 < 15; fi18++ {
                _ = len(v5)
                _ = fi18
            }
        } else {
            if v3 {
                var v19 float64 = (1.5 + 1e10)
                for fi20 := 0; fi20 < 10; fi20++ {
                    var v21 int = (fi20 + fi20)
                    _ = fi20
                    _ = v21
                }
                _ = v19
            } else {
                v3 = (95 > -128)
                var v22 int16 = 1000
                var v23 float64 = (float64(-2.25) - ((-2.25 * 1.5) - 1e10))
                _ = v22
                _ = v23
            }
            fmt.Println(v4)
            fmt.Println(v4)
        }
    }
    var v24 int32 = -2147483648
    var v25 bool = v3
    v25 = (-874 > -513)
    v3 = true
    switch "" {
    case "\n\t":
        switch "a" {
        case "\n\t":
            if (!v3) {
                switch 1 {
                case 618:
                    for fi26 := 0; fi26 < 18; fi26++ {
                        var v27 float32 = (3.14159 - 0.1)
                        _ = fi26
                        _ = v27
                    }
                }
                var v28 int16 = 177
                _ = v28
            } else {
                v24 = v24
                var v29 int8 = 127
                _ = v29
            }
        case "a":
            if len(v5) >= 2 {
                v5 = v5[1:len(v5)]
            }
        default:
            v30 := GenS5{F0: v4}
            _ = v30
        }
    case "":
        v4 += 869
    case "hello":
        fmt.Println(v4, v25, v4)
    default:
        v25 = v3
    }
    v4 = (^(v4))
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v24:", v24)
    fmt.Println("v25:", v25)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v2
    _ = v24
    _ = v25
    _ = v3
    _ = v4
    _ = v5
}
