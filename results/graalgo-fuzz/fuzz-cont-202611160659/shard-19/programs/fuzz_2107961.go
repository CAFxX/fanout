package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 bool; F2 uint64 }
type GenS2 struct { F0 uint32; F1 int; F2 int32 }
type GenS3 struct { F0 uint16 }
type GenS4 struct { F0 uint }
type GenS5 struct { F0 int8; F1 int64; F2 int32 }
type GenS6 struct { F0 bool }
type GenS7 struct { F0 uint }

func main() {
    v1 := make(map[string]int)
    v2 := make(map[int]string)
    var v3 string = "\n\t"
    v4 := []string{((v3 + v3) + "\n\t"), (v3 + (v3 + "foo bar")), ("\n\t" + v3)}
    var v5 uint32 = 46
    var v6 float32 = 3.14159
    switch 430 {
    case -844:
        var v7 uint = 4294967295
        _ = v7
    case 4294967295:
        if (42 < 1) {
            for fi8 := 0; fi8 < 9; fi8++ {
                for fi9 := 0; fi9 < 6; fi9++ {
                    switch (("\n\t" + (v3 + "\n\t")) + (v3 + ("hello" + v3))) {
                    case "foo bar":
                        delete(v2, fi9)
                    case "hello":
                        fmt.Println(fi8, v6, fi9)
                    default:
                        fmt.Println(v6, v3, fi8)
                    }
                    v10 := GenS1{F0: 1, F1: (65535 == 255), F2: 125}
                    if len(v4) >= 2 {
                        v4 = v4[1:len(v4)]
                    }
                    _ = fi9
                    _ = v10
                }
                if (!(!(!(v5 == (v5 << ((65535) & 63)))))) {
                    var v11 uint32 = 2147483647
                    switch (((v3 + v3) + "a") + "\n\t") {
                    case "":
                        var v12 int = (fi8 | -1)
                        _ = v12
                    case "\n\t":
                        v5 = 1
                    case "a":
                        switch v3 {
                        case "":
                            _ = len(v4)
                        case "a":
                            v3 = v3
                        default:
                            if true {
                                var v13 int64 = -967
                                _ = v13
                            }
                        }
                    default:
                        if len(v4) > 0 {
                            v4[0] = v3
                        }
                    }
                    _ = v11
                } else {
                    v14 := GenS1{F0: 2, F1: (((false || false) || (9223372036854775807 == -939)) || true), F2: uint64(v5)}
                    switch 1 {
                    case -1:
                        v6 = v6
                    case 491:
                        v5 -= 42
                    }
                    fmt.Println(v3)
                    _ = v14
                }
                _ = fi8
            }
            switch int(v5) {
            case -1:
                fmt.Println(v3, v6, v6)
            case 4294967295:
                delete(v1, (("" + (v3 + ((v3 + v3) + ((v3 + "\n\t") + v3)))) + ("\n\t" + "foo bar")))
            }
            if true {
                for fi15 := 0; fi15 < 16; fi15++ {
                    var v16 int = 0
                    for fi17 := 0; fi17 < 15; fi17++ {
                        v3 = v3
                        v3 = v3
                        fmt.Println(fi17, fi17)
                        v18 := GenS1{F0: 579, F1: true, F2: 42}
                        _ = fi17
                        _ = v18
                    }
                    _ = fi15
                    _ = v16
                }
                var v19 int64 = -160
                v5 = (v5 | v5)
                _ = v19
            } else {
                var v20 float32 = (v6 / ((0.1 + ((-2.25 - (-2.25 - v6)) - v6)) * v6))
                _ = len(v4)
                v21 := GenS1{F0: 1, F1: true, F2: 18446744073709551615}
                _ = v20
                _ = v21
            }
        } else {
            v3 = v3
            for fi22 := 0; fi22 < 3; fi22++ {
                fmt.Println(fi22)
                v6 = (((v6 * 0.0) * v6) + v6)
                _ = fi22
            }
        }
    case -9223372036854775808:
        v3 = v3
    default:
        v3 = (("foo bar" + ("a" + v3)) + v3)
    }
    for fi23 := 0; fi23 < 18; fi23++ {
        if len(v4) > 0 {
            v4[0] = ("a" + v3)
        }
        _ = fi23
    }
    v24 := GenS1{F0: 933, F1: false, F2: 18446744073709551615}
    fmt.Println(v3, v3)
    if true {
        if true {
            if true {
                v25 := GenS2{F0: 652, F1: 32767, F2: 2147483647}
                _ = v25
            }
            v5 += (v5 & (v5 >> ((2) & 63)))
        } else {
            v5 -= (v5 >> ((424) & 63))
            for fi26 := 0; fi26 < 18; fi26++ {
                delete(v2, fi26)
                var v27 int = fi26
                _ = fi26
                _ = v27
            }
        }
    }
    var v28 string = v3
    fmt.Println(v5, v3)
    v29 := GenS2{F0: v5, F1: 935, F2: 42}
    var v30 int = 796
    v1[(v3 + (v28 + ("\n\t" + "")))] = (v30 & (v30 << ((1) & 63)))
    if false {
        if (-1 == -32768) {
            v31 := GenS2{F0: 0, F1: v30, F2: 0}
            _ = v31
        }
        if true {
            v4 = append(v4, ("foo bar" + ((v28 + v3) + v28)))
        }
    }
    fmt.Println(v30, v30)
    var v32 int16 = 0
    if (!((true || true) || false)) {
        for fi33 := 0; fi33 < 4; fi33++ {
            switch v30 {
            case 1:
                var v34 int32 = 42
                _ = v34
            case 2147483647:
                var v35 float32 = v6
                _ = v35
            case -547:
                for fi36 := 0; fi36 < 2; fi36++ {
                    v37 := GenS1{F0: 901, F1: false, F2: uint64(v5)}
                    _ = fi36
                    _ = v37
                }
            default:
                var v38 uint32 = 42
                _ = v38
            }
            var v39 int8 = 65
            switch v28 {
            case "\n\t":
                v28 = ((v28 + ((v3 + ("hello" + v3)) + v3)) + "foo bar")
            case "foo bar":
                v2[(fi33 | fi33)] = v28
            default:
                if true {
                    if (4294967295 > 968) {
                        fmt.Println(v28, fi33)
                        var v40 int8 = (v39 ^ v39)
                        for fi41 := 0; fi41 < 7; fi41++ {
                            if false {
                                v40 = v39
                                var v42 float32 = (3.14159 / 0.1)
                                fmt.Println(v39, v39, v40)
                                _ = v42
                            }
                            fmt.Println(v39)
                            _ = fi41
                        }
                        _ = v40
                    } else {
                        switch "hello" {
                        case "":
                            var v43 uint16 = 1
                            _ = v43
                        case "\n\t":
                            v44 := GenS3{F0: 65535}
                            _ = v44
                        default:
                            v45 := GenS1{F0: 0, F1: false, F2: 1}
                            _ = v45
                        }
                    }
                }
            }
            _ = fi33
            _ = v39
        }
        var v46 float64 = (0.1 / 1.0)
        if (2147483647 < 1) {
            if (!true) {
                var v47 int64 = 0
                for fi48 := 0; fi48 < 20; fi48++ {
                    if (!true) {
                        for fi49 := 0; fi49 < 4; fi49++ {
                            fmt.Println(v30)
                            _ = fi49
                        }
                        _ = len(v2)
                        switch "a" {
                        case "\n\t":
                            if (914 < (v47 % (((v47 | (v47 ^ 9223372036854775807))) | 1))) {
                                v5 = 65535
                            } else {
                                switch (v30 + v30) {
                                case -9223372036854775808:
                                    v46 += v46
                                case -637:
                                    if ((-97 >= 100) && true) {
                                        v50 := GenS1{F0: 834, F1: false, F2: 288}
                                        for fi51 := 0; fi51 < 13; fi51++ {
                                            fmt.Println(v30, v6)
                                            _ = fi51
                                        }
                                        _ = v50
                                    }
                                default:
                                    var v52 float64 = (0.0 * (3.14159 + ((float64(v6) / v46) - v46)))
                                    _ = v52
                                }
                                v28 = ""
                            }
                        case "hello":
                            v4 = append(v4, v28)
                        case "foo bar":
                            for fi53 := 0; fi53 < 6; fi53++ {
                                switch fi48 {
                                case 2147483647:
                                    switch "a" {
                                    case "\n\t":
                                        if ((v32 / (((v32 & (v32 / ((-32768) | 1)))) | 1)) > v32) {
                                            if false {
                                                if false {
                                                    switch (fi48 >> ((2) & 63)) {
                                                    case 32767:
                                                        var v54 int32 = -838
                                                        _ = v54
                                                    }
                                                    _ = len(v4)
                                                }
                                                if true {
                                                    var v55 uint16 = 55
                                                    v56 := GenS4{F0: 56}
                                                    _ = v55
                                                    _ = v56
                                                }
                                            } else {
                                                v4 = append(v4, v3)
                                            }
                                            if len(v4) >= 2 {
                                                v4 = v4[1:len(v4)]
                                            }
                                            var v57 float64 = 1e10
                                            _ = v57
                                        } else {
                                            switch v30 {
                                            case 127:
                                                _ = len(v1)
                                            case 2147483647:
                                                switch v30 {
                                                case 32768:
                                                    v58 := GenS5{F0: 1, F1: (v47 + 9223372036854775807), F2: -267}
                                                    _ = v58
                                                case 42:
                                                    fmt.Println(v32)
                                                    v47 = v47
                                                case 376:
                                                    fmt.Println(v3)
                                                default:
                                                    v59 := GenS5{F0: 100, F1: v47, F2: int32(fi48)}
                                                    _ = v59
                                                }
                                            }
                                            v6 -= v6
                                            var v60 uint32 = v5
                                            if true {
                                                var v61 float64 = ((v46 * -2.25) + v46)
                                                v47 -= v47
                                                _ = v61
                                            }
                                            var v62 uint64 = 284
                                            _ = v60
                                            _ = v62
                                        }
                                    case "":
                                        var v63 uint64 = 2
                                        _ = v63
                                    }
                                case 698:
                                    v1["hello"] = fi53
                                }
                                var v64 uint32 = (v5 | v5)
                                if (false && false) {
                                    v65 := GenS6{F0: false}
                                    v5 -= (v64 << ((255) & 63))
                                    _ = v65
                                }
                                _ = fi53
                                _ = v64
                            }
                        default:
                            v30 = (fi48 - (fi48 << ((0) & 63)))
                        }
                    }
                    for fi66 := 0; fi66 < 4; fi66++ {
                        v67 := GenS7{F0: 18446744073709551615}
                        _ = fi66
                        _ = v67
                    }
                    fmt.Println(v3, v46)
                    _ = fi48
                }
                _ = v47
            }
        } else {
            var v68 int = v30
            v28 = (("\n\t" + "a") + "\n\t")
            _ = v68
        }
        _ = v46
    } else {
        var v69 uint16 = 1
        v5 += v5
        _ = v69
    }
    var v70 int64 = -9223372036854775808
    v3 = v3
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v28:", v28)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("v32:", v32)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v70:", v70)
    _ = v1
    _ = v2
    _ = v24
    _ = v28
    _ = v29
    _ = v3
    _ = v30
    _ = v32
    _ = v4
    _ = v5
    _ = v6
    _ = v70
}
