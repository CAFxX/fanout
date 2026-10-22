package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 uint }
type GenS2 struct { F0 uint; F1 bool; F2 string }
type GenS3 struct { F0 uint; F1 uint32; F2 uint16 }
type GenS4 struct { F0 string }

func main() {
    var v1 int8 = 83
    var v2 uint8 = 248
    v3 := []string{""}
    v4 := GenS1{F0: -9223372036854775808, F1: 2}
    _ = len(v3)
    if true {
        v2 = 128
    }
    v2 = (^(v2))
    var v5 int64 = -712
    switch 42 {
    case 42:
        var v6 uint8 = (v2 << ((0) & 63))
        _ = v6
    case 1:
        var v7 int = -931
        _ = v7
    case -1:
        v8 := GenS2{F0: 18446744073709551615, F1: true, F2: "\n\t"}
        _ = v8
    }
    if (2 > 32767) {
        var v9 int = 0
        v3 = append(v3, "a")
        var v10 uint = 898
        _ = v10
        _ = v9
    } else {
        for fi11 := 0; fi11 < 6; fi11++ {
            for fi12 := 0; fi12 < 3; fi12++ {
                fmt.Println(fi12, fi12)
                v13 := GenS3{F0: 4294967295, F1: 65535, F2: 520}
                v5 -= 0
                if len(v3) > 0 {
                    v3[0] = ("foo bar" + "\n\t")
                }
                _ = fi12
                _ = v13
            }
            _ = fi11
        }
        fmt.Println(v5)
    }
    fmt.Println(v2, v2, v1)
    v5 = v5
    v2 += 42
    v5 = v5
    v5 += (v5 << ((840) & 63))
    switch "a" {
    case "foo bar":
        for fi14 := 0; fi14 < 9; fi14++ {
            v2 = (v2 - v2)
            _ = fi14
        }
    case "hello":
        if true {
            for fi15 := 0; fi15 < 9; fi15++ {
                var v16 uint64 = 0
                if (false && (-222 >= 2147483647)) {
                    v17 := GenS1{F0: (v5 << ((42) & 63)), F1: 0}
                    if (!false) {
                        if len(v3) >= 2 {
                            v3 = v3[1:len(v3)]
                        }
                    } else {
                        if (!(4294967295 < 4294967295)) {
                            if (2147483647 > 195) {
                                if len(v3) > 0 {
                                    v3[0] = ("" + ("\n\t" + (("a" + "hello") + ("a" + "foo bar"))))
                                }
                            }
                            v16 -= v16
                            fmt.Println(fi15, v16, v1)
                        }
                    }
                    _ = v17
                }
                switch "\n\t" {
                case "":
                    v18 := GenS4{F0: "hello"}
                    _ = v18
                case "\n\t":
                    var v19 uint16 = 65535
                    _ = v19
                case "foo bar":
                    v1 = 97
                default:
                    var v20 bool = (-20 < 42)
                    _ = v20
                }
                _ = fi15
                _ = v16
            }
        }
    default:
        if (false && ((127 < 0) && false)) {
            v5 -= -817
        }
    }
    _ = len(v3)
    var v21 uint16 = 2
    for fi22 := 0; fi22 < 5; fi22++ {
        switch ("" + ("\n\t" + "a")) {
        case "":
            if false {
                if true {
                    fmt.Println(v5)
                    for fi23 := 0; fi23 < 18; fi23++ {
                        var v24 float64 = 0.1
                        switch fi22 {
                        case 2147483647:
                            switch "\n\t" {
                            case "a":
                                v5 += (v5 >> ((1) & 63))
                            case "hello":
                                v25 := GenS1{F0: v5, F1: 255}
                                _ = v25
                            case "foo bar":
                                v21 = (v21 % ((v21) | 1))
                            default:
                                v24 -= (float64(1.5) - (float64(0.0) / (((v24 / float64(0.1)) * 1e10) * float64((0.0 + 1.5)))))
                            }
                            var v26 uint = 2
                            _ = v26
                        case 225:
                            if (v2 < (v2 - (v2 - 0))) {
                                var v27 int8 = (v1 | (v1 ^ -1))
                                _ = v27
                            } else {
                                fmt.Println(v21, v21, fi23)
                            }
                        default:
                            if len(v3) >= 2 {
                                v3 = v3[1:len(v3)]
                            }
                        }
                        switch "foo bar" {
                        case "foo bar":
                            var v28 float64 = (-2.25 + (((3.14159 + (v24 + ((0.1 - v24) / 1.5))) - float64(0.1)) - v24))
                            _ = v28
                        default:
                            _ = len(v3)
                        }
                        _ = fi23
                        _ = v24
                    }
                }
                v5 = (^(v5))
            }
        default:
            fmt.Println(v5, v1)
        }
        _ = fi22
    }
    var v29 float32 = 1e10
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("v29:", v29)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v21
    _ = v29
    _ = v3
    _ = v4
    _ = v5
}
