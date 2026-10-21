package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 bool }
type GenS2 struct { F0 int16; F1 int64 }
type GenS3 struct { F0 string; F1 int32 }
type GenS4 struct { F0 uint32; F1 uint64; F2 int32 }
type GenS5 struct { F0 uint8; F1 bool; F2 int64 }

func main() {
    var v1 uint32 = 65535
    v2 := []string{("\n\t" + "a"), "hello", ("a" + "\n\t"), ("hello" + "")}
    var v3 float64 = (float64(1e10) * 1.5)
    var v4 float64 = -2.25
    if len(v2) > 0 {
        v2[0] = "\n\t"
    }
    v3 = v3
    if (0 != 65535) {
        var v5 uint8 = 127
        _ = v5
    }
    v6 := GenS1{F0: -476, F1: false}
    var v7 int32 = 1
    for fi8 := 0; fi8 < 15; fi8++ {
        for fi9 := 0; fi9 < 13; fi9++ {
            v1 = (v1 & (v1 - 223))
            v10 := GenS2{F0: -458, F1: 42}
            var v11 uint = 473
            _ = fi9
            _ = v10
            _ = v11
        }
        if (-1 <= -625) {
            for fi12 := 0; fi12 < 14; fi12++ {
                v13 := GenS1{F0: -123, F1: (-1 > -1)}
                _ = fi12
                _ = v13
            }
            switch "foo bar" {
            case "\n\t":
                var v14 int16 = 42
                _ = v14
            default:
                v7 = 42
            }
        } else {
            var v15 int32 = 2147483647
            _ = v15
        }
        _ = fi8
    }
    var v16 float32 = 0.0
    if ((false && true) || (!(!(0 < 990)))) {
        v4 -= 3.14159
        v1 = (v1 + (v1 % (((v1 & (v1 - 0))) | 1)))
    } else {
        switch -9223372036854775808 {
        case 32768:
            if true {
                v7 -= (v7 << ((1) & 63))
            }
        case 9223372036854775807:
            for fi17 := 0; fi17 < 12; fi17++ {
                switch -481 {
                case -446:
                    fmt.Println(fi17, v7)
                case 9223372036854775807:
                    if false {
                        var v18 float32 = v16
                        v3 = v3
                        v16 -= -2.25
                        _ = v18
                    } else {
                        switch (fi17 >> ((4294967295) & 63)) {
                        case 4294967295:
                            v3 = (v3 * v3)
                        case -401:
                            v19 := GenS3{F0: ("\n\t" + ""), F1: (v7 % (((v7 << ((447) & 63))) | 1))}
                            _ = v19
                        case 128:
                            v1 = v1
                        }
                        var v20 int64 = 1
                        _ = v20
                    }
                default:
                    if len(v2) >= 2 {
                        v2 = v2[1:len(v2)]
                    }
                }
                if len(v2) > 0 {
                    v2[0] = "foo bar"
                }
                v16 = (v16 * 1e10)
                _ = fi17
            }
        case 210:
            var v21 float32 = float32((v4 + v4))
            _ = v21
        default:
            v22 := GenS1{F0: -1, F1: (!true)}
            _ = v22
        }
        switch "a" {
        case "hello":
            var v23 string = ((("" + (("foo bar" + "a") + "\n\t")) + "foo bar") + "")
            _ = v23
        case "":
            if false {
                if (1000 <= 32767) {
                    v4 -= (v3 + v4)
                }
                var v24 uint32 = v1
                switch -978 {
                case 4294967295:
                    for fi25 := 0; fi25 < 11; fi25++ {
                        _ = len(v2)
                        fmt.Println(v24, v16, v16)
                        var v26 int32 = (v7 ^ (v7 | 279))
                        _ = fi25
                        _ = v26
                    }
                case 1:
                    var v27 uint64 = 42
                    _ = v27
                case 127:
                    var v28 uint32 = v1
                    _ = v28
                }
                _ = v24
            } else {
                if (false || true) {
                    switch "a" {
                    case "hello":
                        switch ("\n\t" + "foo bar") {
                        case "\n\t":
                            switch 9223372036854775807 {
                            case 942:
                                v7 = 2147483647
                            case 937:
                                switch 128 {
                                case 127:
                                    v29 := GenS1{F0: -1, F1: ((42 != 32767) || true)}
                                    _ = v29
                                case 128:
                                    for fi30 := 0; fi30 < 8; fi30++ {
                                        var v31 uint = 534
                                        var v32 int16 = -193
                                        v33 := GenS2{F0: -575, F1: -9223372036854775808}
                                        _ = fi30
                                        _ = v31
                                        _ = v32
                                        _ = v33
                                    }
                                case -1:
                                    var v34 uint16 = 65535
                                    _ = v34
                                }
                            case 9223372036854775807:
                                var v35 uint32 = (v1 & (^(v1)))
                                _ = v35
                            default:
                                v2 = append(v2, "hello")
                            }
                        case "a":
                            var v36 string = "a"
                            _ = v36
                        }
                    case "foo bar":
                        v7 += -722
                    default:
                        if (!false) {
                            v37 := GenS1{F0: 1, F1: (817 <= 421)}
                            if len(v2) > 0 {
                                v2[0] = (("foo bar" + "\n\t") + ("a" + ("" + "hello")))
                            }
                            v38 := GenS1{F0: 42, F1: false}
                            _ = v37
                            _ = v38
                        }
                    }
                    if false {
                        v39 := GenS2{F0: 1, F1: -1}
                        _ = v39
                    } else {
                        v40 := GenS1{F0: 386, F1: (-9223372036854775808 == 518)}
                        v41 := GenS3{F0: ("hello" + "hello"), F1: v7}
                        for fi42 := 0; fi42 < 7; fi42++ {
                            switch fi42 {
                            case 1:
                                var v43 int8 = int8(v1)
                                _ = v43
                            }
                            _ = fi42
                        }
                        _ = v40
                        _ = v41
                    }
                } else {
                    if (583 == 65535) {
                        for fi44 := 0; fi44 < 4; fi44++ {
                            v4 += v3
                            fmt.Println(v3)
                            for fi45 := 0; fi45 < 17; fi45++ {
                                v46 := GenS2{F0: -1, F1: 1}
                                _ = len(v2)
                                _ = fi45
                                _ = v46
                            }
                            _ = fi44
                        }
                    }
                    v2 = append(v2, "foo bar")
                }
                var v47 uint64 = 1
                _ = v47
            }
        case "\n\t":
            switch "foo bar" {
            case "hello":
                var v48 uint32 = 607
                _ = v48
            case "a":
                v49 := GenS2{F0: 160, F1: 0}
                _ = v49
            case "\n\t":
                v3 += (v4 - float64((float32(1.5) * v16)))
            }
        }
        if (false && true) {
            v16 = float32(1.5)
            v50 := GenS4{F0: 240, F1: 42, F2: -2147483648}
            v51 := GenS3{F0: "", F1: -1}
            _ = v50
            _ = v51
        }
    }
    v52 := GenS5{F0: 0, F1: (!true), F2: -9223372036854775808}
    var v53 float64 = 1.5
    if (!(9223372036854775807 != 9223372036854775807)) {
        v4 = float64(v16)
    } else {
        for fi54 := 0; fi54 < 7; fi54++ {
            v2 = append(v2, (("foo bar" + ("foo bar" + "a")) + (("\n\t" + "a") + "hello")))
            _ = fi54
        }
    }
    var v55 int = 128
    if (!false) {
        for fi56 := 0; fi56 < 3; fi56++ {
            if (((((((v16 / (-2.25 * v16)) + v16) * 3.14159) - float32(0.0)) <= v16) || false) && (-1 >= 1)) {
                for fi57 := 0; fi57 < 5; fi57++ {
                    v7 = v7
                    _ = fi57
                }
                var v58 uint = 815
                fmt.Println(v53, fi56)
                _ = v58
            } else {
                var v59 int64 = 9223372036854775807
                _ = v59
            }
            _ = fi56
        }
        var v60 int16 = 989
        _ = v60
    } else {
        var v61 int16 = 1000
        _ = v61
    }
    var v62 uint16 = 42
    if len(v2) > 0 {
        v2[0] = "foo bar"
    }
    fmt.Println("v1:", v1)
    fmt.Println("v16:", v16)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v53:", v53)
    fmt.Println("v55:", v55)
    fmt.Println("v62:", v62)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v16
    _ = v2
    _ = v3
    _ = v4
    _ = v52
    _ = v53
    _ = v55
    _ = v6
    _ = v62
    _ = v7
}
