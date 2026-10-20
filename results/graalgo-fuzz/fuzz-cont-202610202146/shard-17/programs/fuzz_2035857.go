package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 int; F2 uint8 }
type GenS2 struct { F0 int; F1 int8; F2 int16 }
type GenS3 struct { F0 string; F1 uint64 }
type GenS4 struct { F0 int64; F1 uint16; F2 int8 }
type GenS5 struct { F0 int; F1 int8; F2 uint16 }

func main() {
    var v1 int64 = 42
    var v2 float32 = 0.0
    var v3 int16 = -902
    v4 := []string{"hello", ("a" + ""), ("foo bar" + ("a" + "")), ""}
    var v5 bool = (271 < 4294967295)
    var v6 int = -2147483648
    if (151 <= 2) {
        for fi7 := 0; fi7 < 9; fi7++ {
            v1 += (v1 | -563)
            _ = len(v4)
            for fi8 := 0; fi8 < 7; fi8++ {
                if (v2 < (v2 - v2)) {
                    if true {
                        var v9 uint8 = 255
                        _ = len(v4)
                        _ = v9
                    } else {
                        fmt.Println(fi7, fi8)
                    }
                    if ((3.14159 / 3.14159) < -2.25) {
                        fmt.Println(fi7, v5)
                        v2 -= v2
                        v6 = 32767
                    }
                    if v5 {
                        switch ("hello" + "foo bar") {
                        case "hello":
                            fmt.Println(v5, v1, v1)
                        case "a":
                            v6 -= (v6 * 308)
                        case "":
                            if false {
                                var v10 bool = v5
                                _ = v10
                            }
                        default:
                            if v5 {
                                switch ("foo bar" + ("hello" + "a")) {
                                case "a":
                                    v11 := GenS1{F0: 0, F1: fi7, F2: 0}
                                    _ = v11
                                case "\n\t":
                                    var v12 uint = 255
                                    if ((^(v3)) != (v3 - v3)) {
                                        switch ("hello" + (("foo bar" + "") + "hello")) {
                                        case "":
                                            v13 := GenS1{F0: 580, F1: -939, F2: 128}
                                            _ = v13
                                        case "\n\t":
                                            if (0 >= 208) {
                                                switch 27 {
                                                case 1:
                                                    if (232 >= 150) {
                                                        var v14 string = ("a" + ("\n\t" + "foo bar"))
                                                        _ = v14
                                                    }
                                                case 9223372036854775807:
                                                    if len(v4) >= 2 {
                                                        v4 = v4[1:len(v4)]
                                                    }
                                                case -299:
                                                    _ = len(v4)
                                                }
                                            }
                                        default:
                                            v6 = (v6 ^ 128)
                                        }
                                        var v15 int64 = 0
                                        _ = v15
                                    } else {
                                        if len(v4) >= 2 {
                                            v4 = v4[1:len(v4)]
                                        }
                                    }
                                    _ = v12
                                default:
                                    v16 := GenS1{F0: 4294967295, F1: fi7, F2: 42}
                                    _ = v16
                                }
                                v1 = -186
                                switch v6 {
                                case 816:
                                    v4 = append(v4, "\n\t")
                                default:
                                    var v17 int = (v6 % ((fi7) | 1))
                                    _ = v17
                                }
                            }
                            v18 := GenS2{F0: fi7, F1: 42, F2: v3}
                            v19 := GenS3{F0: "", F1: 18446744073709551615}
                            _ = v18
                            _ = v19
                        }
                        var v20 int64 = (v1 ^ 0)
                        fmt.Println(fi7)
                        _ = v20
                    }
                }
                switch v6 {
                case 42:
                    v1 -= v1
                    if len(v4) > 0 {
                        v4[0] = ("\n\t" + ("" + "\n\t"))
                    }
                    v1 = -9223372036854775808
                case 870:
                    var v21 int8 = 42
                    _ = v21
                default:
                    var v22 uint8 = 127
                    _ = v22
                }
                _ = fi8
            }
            _ = fi7
        }
    } else {
        for fi23 := 0; fi23 < 12; fi23++ {
            for fi24 := 0; fi24 < 17; fi24++ {
                v4 = append(v4, "foo bar")
                _ = fi24
            }
            _ = fi23
        }
    }
    v25 := GenS3{F0: ("a" + ("hello" + "a")), F1: 159}
    v26 := GenS4{F0: v1, F1: 2, F2: 0}
    switch -9223372036854775808 {
    case -422:
        v6 -= (v6 + (v6 | int(v1)))
    case 42:
        v27 := GenS5{F0: -9223372036854775808, F1: 0, F2: 65535}
        _ = v27
    default:
        v6 += v6
    }
    var v28 int16 = int16(v1)
    var v29 int16 = 488
    fmt.Println(v6, v3, v1)
    v30 := GenS3{F0: ("" + "foo bar"), F1: 42}
    var v31 float32 = 0.0
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v28:", v28)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("v31:", v31)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v25
    _ = v26
    _ = v28
    _ = v29
    _ = v3
    _ = v30
    _ = v31
    _ = v4
    _ = v5
    _ = v6
}
