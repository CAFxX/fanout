package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 uint64; F2 uint8 }
type GenS2 struct { F0 int16; F1 bool }
type GenS3 struct { F0 int8 }
type GenS4 struct { F0 int8 }
type GenS5 struct { F0 int8 }
type GenS6 struct { F0 int16; F1 uint64 }
type GenS7 struct { F0 int; F1 uint }
type GenS8 struct { F0 uint64; F1 uint32; F2 uint }
type GenS9 struct { F0 uint32; F1 uint8; F2 int16 }

func main() {
    v1 := make(map[string]uint32)
    var v2 int32 = 1
    var v3 uint64 = 1
    var v4 float32 = 3.14159
    v5 := []string{(("a" + "a") + "foo bar"), ("" + ("hello" + ""))}
    v4 += (0.0 - v4)
    v4 = v4
    v6 := GenS1{F0: ((true && (true || (233 <= 4294967295))) && ((2147483647 <= v2) && false)), F1: 762, F2: 61}
    v4 = v4
    switch 42 {
    case 1:
        v2 -= 0
    case 2147483647:
        v3 += v3
    case -233:
        v7 := GenS2{F0: 0, F1: false}
        _ = v7
    }
    var v8 int16 = 6
    fmt.Println(v2)
    fmt.Println(v2, v2, v8)
    fmt.Println(v4, v8, v2)
    if ((true && false) && (!(!(!(32768 == -247))))) {
        if (127 != 1) {
            v9 := GenS3{F0: -128}
            fmt.Println(v8, v3, v3)
            _ = v9
        }
        v3 = (v3 ^ 2)
        v8 += (v8 << ((0) & 63))
    }
    if len(v5) > 0 {
        v5[0] = "hello"
    }
    v10 := GenS4{F0: 50}
    switch "a" {
    case "":
        var v11 uint8 = 255
        _ = v11
    }
    if true {
        delete(v1, (("a" + "hello") + ("\n\t" + "")))
        switch ("foo bar" + ((("" + "a") + "a") + ("a" + ""))) {
        case "":
            switch 991 {
            case 42:
                v2 = v2
            case 32767:
                switch 127 {
                case 127:
                    if (!(false && ((0 >= 65535) && (!false)))) {
                        switch 32768 {
                        case 128:
                            if (2147483647 < 2) {
                                switch int(v2) {
                                case -784:
                                    fmt.Println(v4)
                                case -1:
                                    v3 = v3
                                case 1:
                                    fmt.Println(v2)
                                default:
                                    var v12 int = 32767
                                    _ = v12
                                }
                                _ = len(v5)
                                for fi13 := 0; fi13 < 18; fi13++ {
                                    for fi14 := 0; fi14 < 13; fi14++ {
                                        var v15 int8 = 23
                                        _ = fi14
                                        _ = v15
                                    }
                                    if true {
                                        v16 := GenS5{F0: 127}
                                        v17 := GenS6{F0: (v8 >> ((65535) & 63)), F1: 267}
                                        var v18 float64 = 3.14159
                                        _ = v16
                                        _ = v17
                                        _ = v18
                                    }
                                    _ = fi13
                                }
                            }
                        case 272:
                            v2 = v2
                        case 408:
                            var v19 bool = (9223372036854775807 != 32768)
                            _ = v19
                        }
                        fmt.Println(v2)
                        var v20 string = "a"
                        _ = v20
                    }
                case 4294967295:
                    _ = len(v5)
                }
            case -9223372036854775808:
                v2 += v2
            default:
                v21 := GenS7{F0: 2147483647, F1: 18446744073709551615}
                _ = v21
            }
        }
        v22 := GenS8{F0: v3, F1: 4294967295, F2: 961}
        _ = v22
    }
    if false {
        for fi23 := 0; fi23 < 9; fi23++ {
            _ = len(v1)
            fmt.Println(v8)
            v2 += v2
            _ = fi23
        }
        if true {
            for fi24 := 0; fi24 < 6; fi24++ {
                v2 -= -2147483648
                _ = fi24
            }
            if false {
                for fi25 := 0; fi25 < 5; fi25++ {
                    v26 := GenS1{F0: ((v2 << ((4294967295) & 63)) >= (v2 + v2)), F1: 924, F2: 1}
                    if false {
                        var v27 uint16 = 2
                        switch 2147483647 {
                        case 9223372036854775807:
                            v2 = -2147483648
                        default:
                            v8 -= 42
                        }
                        _ = v27
                    } else {
                        v28 := GenS1{F0: false, F1: 1, F2: 255}
                        if (!(-1 == int32(v8))) {
                            switch "a" {
                            case "\n\t":
                                v4 = v4
                            case "foo bar":
                                for fi29 := 0; fi29 < 9; fi29++ {
                                    switch "foo bar" {
                                    case "a":
                                        v2 = v2
                                    case "foo bar":
                                        _ = len(v1)
                                    }
                                    v3 += v3
                                    var v30 uint = 1
                                    _ = fi29
                                    _ = v30
                                }
                            case "hello":
                                var v31 uint32 = 1
                                _ = v31
                            default:
                                if (true || false) {
                                    v3 -= (v3 + (^(v3)))
                                    v32 := GenS6{F0: v8, F1: (^(v3))}
                                    _ = v32
                                } else {
                                    v5 = append(v5, "")
                                }
                            }
                            _ = len(v1)
                            var v33 int32 = 42
                            _ = v33
                        }
                        _ = v28
                    }
                    _ = fi25
                    _ = v26
                }
                var v34 float32 = ((v4 + v4) - ((1e10 / ((-2.25 * 3.14159) / 1.5)) / v4))
                _ = v34
            }
            var v35 uint64 = (v3 * (^(v3)))
            _ = v35
        } else {
            v5 = append(v5, "\n\t")
            switch 71 {
            case -9223372036854775808:
                for fi36 := 0; fi36 < 9; fi36++ {
                    if true {
                        var v37 uint32 = 42
                        _ = v37
                    } else {
                        v38 := GenS9{F0: 42, F1: 128, F2: 0}
                        if false {
                            var v39 uint64 = 0
                            _ = v39
                        }
                        _ = v38
                    }
                    if (-32768 > (v8 | -32768)) {
                        if len(v5) > 0 {
                            v5[0] = "a"
                        }
                        v8 -= (v8 ^ (v8 + v8))
                    } else {
                        switch (("foo bar" + "") + ("hello" + (("" + "") + ("hello" + ("a" + "hello"))))) {
                        case "a":
                            v40 := v1[("\n\t" + ("a" + "hello"))]
                            _ = v40
                        case "":
                            v41 := GenS4{F0: 28}
                            _ = v41
                        case "\n\t":
                            switch (fi36 + fi36) {
                            case 954:
                                _ = len(v5)
                            default:
                                for fi42 := 0; fi42 < 14; fi42++ {
                                    switch 32767 {
                                    case 0:
                                        v43 := GenS7{F0: (fi36 % ((128) | 1)), F1: 4294967295}
                                        _ = v43
                                    case 42:
                                        v8 = -626
                                    default:
                                        var v44 int16 = (v8 ^ v8)
                                        _ = v44
                                    }
                                    _ = fi42
                                }
                            }
                        }
                        v45 := GenS7{F0: (fi36 + (fi36 % (((fi36 % (((fi36 / (((fi36 + fi36)) | 1))) | 1))) | 1))), F1: 512}
                        _ = v45
                    }
                    _ = fi36
                }
            }
            v8 = (v8 >> ((154) & 63))
        }
    }
    v2 = (v2 / (((v2 * 2147483647)) | 1))
    if len(v5) > 0 {
        v5[0] = "a"
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v10
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v8
}
