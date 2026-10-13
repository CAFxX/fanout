package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 int }

func main() {
    var v1 int32 = -837
    var v2 int64 = 1
    var v3 bool = false
    v4 := []int{-9223372036854775808, 32767, 32767}
    var v5 bool = (100 == int8(v2))
    var v6 int = 32768
    for fi7 := 0; fi7 < 20; fi7++ {
        v3 = (1 == 80)
        _ = fi7
    }
    v6 -= v6
    if (1.5 >= (-2.25 * float32((0.1 * 0.0)))) {
        v6 = (v6 >> ((2) & 63))
        for fi8 := 0; fi8 < 15; fi8++ {
            v4 = append(v4, fi8)
            v3 = v3
            _ = fi8
        }
        for fi9 := 0; fi9 < 7; fi9++ {
            var v10 uint8 = 2
            if v3 {
                switch fi9 {
                case 0:
                    v11 := GenS1{F0: (v2 - -1), F1: (^(fi9))}
                    _ = v11
                case -1:
                    if len(v4) > 0 {
                        v4[0] = 4294967295
                    }
                default:
                    if (657 == 42) {
                        var v12 string = "hello"
                        var v13 int = fi9
                        _ = v12
                        _ = v13
                    }
                }
            }
            for fi14 := 0; fi14 < 18; fi14++ {
                switch "hello" {
                case "a":
                    var v15 uint = 819
                    _ = v15
                case "\n\t":
                    v6 = (-(fi9))
                case "":
                    fmt.Println(fi14, v2, v5)
                default:
                    switch "a" {
                    case "":
                        var v16 float64 = float64(((1e10 + (3.14159 + 0.0)) / 0.1))
                        _ = v16
                    case "a":
                        fmt.Println(v10, fi14, fi9)
                    case "foo bar":
                        fmt.Println(v1, v10, v5)
                    default:
                        if (((1e10 <= (1.5 - float32(-2.25))) || (!(!(42 < 2)))) || ((1 > 2) || ("a" > "hello"))) {
                            v10 -= (v10 ^ v10)
                            if (v5 || (!false)) {
                                var v17 string = ("a" + (("a" + "hello") + "a"))
                                fmt.Println(v3, fi14, v5)
                                _ = v17
                            } else {
                                var v18 uint8 = (v10 | (v10 | v10))
                                if len(v4) >= 2 {
                                    v4 = v4[1:len(v4)]
                                }
                                _ = v18
                            }
                            if (42 != 41) {
                                var v19 uint32 = 65535
                                if len(v4) >= 2 {
                                    v4 = v4[1:len(v4)]
                                }
                                v6 -= (fi14 ^ 454)
                                _ = v19
                            } else {
                                if (v1 > v1) {
                                    var v20 int32 = v1
                                    v5 = v5
                                    _ = v20
                                }
                                v21 := GenS1{F0: v2, F1: v6}
                                v4 = append(v4, v6)
                                fmt.Println(v10, v1, fi9)
                                _ = v21
                            }
                        }
                    }
                }
                v3 = v5
                v22 := GenS1{F0: -1, F1: (v6 ^ 455)}
                _ = fi14
                _ = v22
            }
            _ = fi9
            _ = v10
        }
    } else {
        if (false || (992 > 1)) {
            v1 += (v1 / ((-1) | 1))
            var v23 int16 = -1
            if len(v4) >= 2 {
                v4 = v4[1:len(v4)]
            }
            _ = v23
        } else {
            var v24 string = "\n\t"
            var v25 uint8 = 42
            _ = v24
            _ = v25
        }
        v6 = v6
    }
    v3 = v5
    if len(v4) >= 2 {
        v4 = v4[1:len(v4)]
    }
    v26 := GenS1{F0: (v2 << ((42) & 63)), F1: v6}
    for fi27 := 0; fi27 < 4; fi27++ {
        switch 722 {
        case 32767:
            fmt.Println(fi27)
        case 4294967295:
            v2 = (v2 % (((v2 >> ((18446744073709551615) & 63))) | 1))
        case -9223372036854775808:
            v4 = append(v4, (v6 % ((fi27) | 1)))
        }
        fmt.Println(fi27, v6, v1)
        _ = fi27
    }
    v28 := GenS1{F0: (v2 - (v2 / (((v2 >> ((255) & 63))) | 1))), F1: v6}
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v26
    _ = v28
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
