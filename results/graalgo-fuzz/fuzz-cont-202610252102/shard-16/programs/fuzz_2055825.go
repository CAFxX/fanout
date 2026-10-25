package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 uint8 }
type GenS2 struct { F0 uint16 }
type GenS3 struct { F0 uint64; F1 uint; F2 bool }
type GenS4 struct { F0 int64; F1 uint64; F2 uint8 }
type GenS5 struct { F0 string; F1 int8; F2 uint }
type GenS6 struct { F0 string; F1 int; F2 int }
type GenS7 struct { F0 int16; F1 int64; F2 int32 }

func main() {
    var v1 uint32 = 0
    v2 := make(map[string]string)
    v3 := make(map[int]uint32)
    v4 := []int{1, -1}
    if (0 >= 0) {
        if (!(42 > 354)) {
            if false {
                fmt.Println(v1)
                v3[-290] = v1
                switch 491 {
                case 2147483647:
                    fmt.Println(v1)
                case -110:
                    if (("foo bar" + ("a" + ("a" + ""))) >= "\n\t") {
                        var v5 uint32 = (v1 >> ((255) & 63))
                        _ = v5
                    } else {
                        if len(v4) >= 2 {
                            v4 = v4[1:len(v4)]
                        }
                        if true {
                            v1 += 887
                        } else {
                            v6 := GenS1{F0: 0, F1: uint8(v1)}
                            if false {
                                v1 += 1
                                delete(v3, 127)
                                var v7 int64 = 274
                                _ = v7
                            }
                            _ = v6
                        }
                    }
                case 0:
                    for fi8 := 0; fi8 < 9; fi8++ {
                        delete(v3, fi8)
                        _ = fi8
                    }
                }
            }
            switch -1 {
            case 9223372036854775807:
                switch ("a" + "hello") {
                case "foo bar":
                    if (255 >= 65535) {
                        _ = len(v4)
                        v1 = 4294967295
                    }
                case "\n\t":
                    v9 := GenS1{F0: 42, F1: 226}
                    _ = v9
                }
            }
        } else {
            var v10 string = "hello"
            _ = v10
        }
        v11 := GenS1{F0: -728, F1: 2}
        _ = v11
    } else {
        if true {
            if (!(207 < 4294967295)) {
                if (!(0 == -128)) {
                    if true {
                        v12 := GenS1{F0: 1, F1: 1}
                        _ = v12
                    } else {
                        for fi13 := 0; fi13 < 11; fi13++ {
                            var v14 int64 = 610
                            v2[("a" + "a")] = "a"
                            var v15 int8 = 0
                            _ = fi13
                            _ = v14
                            _ = v15
                        }
                    }
                    v16 := GenS2{F0: 42}
                    _ = v16
                }
            }
            v4 = append(v4, 1)
        } else {
            switch -407 {
            case 260:
                v17 := GenS2{F0: 2}
                _ = v17
            case 128:
                if ((1 >= 2) || (65535 == 32767)) {
                    if len(v4) > 0 {
                        v4[0] = int(v1)
                    }
                    for fi18 := 0; fi18 < 8; fi18++ {
                        v1 = (v1 ^ v1)
                        var v19 int64 = 700
                        _ = fi18
                        _ = v19
                    }
                }
            case -1:
                v1 = v1
            default:
                for fi20 := 0; fi20 < 17; fi20++ {
                    var v21 float32 = (-2.25 - (float32(float64(0.1)) / -2.25))
                    fmt.Println(v1)
                    fmt.Println(v21, v1)
                    _ = fi20
                    _ = v21
                }
            }
            fmt.Println(v1)
            if (-2147483648 > 2147483647) {
                for fi22 := 0; fi22 < 15; fi22++ {
                    for fi23 := 0; fi23 < 6; fi23++ {
                        v24 := v3[fi22]
                        if (float32((((1.5 + float64(1e10)) - (1e10 * (((-2.25 + float64(float32(-2.25))) * 0.0) - 1.5))) * -2.25)) < (float32(1e10) + 1e10)) {
                            if (97 < 100) {
                                if true {
                                    v24 = v24
                                    _ = len(v4)
                                    var v25 float32 = (0.0 / 1.0)
                                    fmt.Println(fi23, v1, v1)
                                    v1 += v1
                                    if (v24 == v1) {
                                        v26 := GenS3{F0: 2, F1: 65535, F2: (42 == 32767)}
                                        v25 = v25
                                        v27 := GenS3{F0: 42, F1: 1, F2: true}
                                        _ = v26
                                        _ = v27
                                    }
                                    _ = v25
                                } else {
                                    v1 = (v1 % ((904) | 1))
                                    var v28 uint32 = (v24 & (v1 << ((65535) & 63)))
                                    v1 = (^(v24))
                                    v29 := GenS1{F0: 9223372036854775807, F1: 255}
                                    _ = v28
                                    _ = v29
                                }
                                v30 := GenS2{F0: 2}
                                v1 -= (v24 & (v1 + (v24 + 1)))
                                _ = v30
                            }
                            fmt.Println(v24)
                        }
                        if true {
                            var v31 float32 = 0.1
                            _ = v31
                        } else {
                            v24 = 4294967295
                            v32 := GenS4{F0: -1, F1: 18446744073709551615, F2: 0}
                            fmt.Println(fi22, v1, v1)
                            _ = v32
                        }
                        switch 42 {
                        case 700:
                            v24 -= 810
                        }
                        if len(v4) > 0 {
                            v4[0] = (fi22 & (^(fi23)))
                        }
                        _ = fi23
                        _ = v24
                    }
                    for fi33 := 0; fi33 < 17; fi33++ {
                        if ((true && (false && true)) || (!true)) {
                            if (-2147483648 > 2147483647) {
                                v1 = v1
                            }
                        }
                        v34 := GenS5{F0: ("hello" + (("hello" + "") + ("a" + ""))), F1: -1, F2: 1}
                        var v35 int8 = -128
                        _ = fi33
                        _ = v34
                        _ = v35
                    }
                    _ = fi22
                }
                var v36 int8 = -128
                v4 = append(v4, 32767)
                _ = v36
            } else {
                switch 32767 {
                case 2147483647:
                    v37 := GenS4{F0: 42, F1: 612, F2: 127}
                    _ = v37
                case 1:
                    var v38 uint8 = 255
                    _ = v38
                case 32767:
                    for fi39 := 0; fi39 < 8; fi39++ {
                        var v40 int32 = -342
                        _ = len(v4)
                        _ = fi39
                        _ = v40
                    }
                }
                for fi41 := 0; fi41 < 17; fi41++ {
                    if (-9223372036854775808 == 9223372036854775807) {
                        if (!true) {
                            v42 := v3[fi41]
                            _ = v42
                        }
                        _ = len(v3)
                        var v43 int32 = 2147483647
                        _ = v43
                    } else {
                        var v44 uint = 1
                        _ = v44
                    }
                    var v45 bool = (false || (false || false))
                    _ = fi41
                    _ = v45
                }
            }
        }
        for fi46 := 0; fi46 < 3; fi46++ {
            v47 := GenS6{F0: ("" + ""), F1: (fi46 ^ -9223372036854775808), F2: fi46}
            for fi48 := 0; fi48 < 18; fi48++ {
                v49 := GenS5{F0: "foo bar", F1: 0, F2: 0}
                var v50 string = "a"
                _ = len(v3)
                if true {
                    var v51 int16 = 1000
                    v52 := GenS2{F0: 42}
                    v50 = "a"
                    _ = v51
                    _ = v52
                }
                _ = fi48
                _ = v49
                _ = v50
            }
            fmt.Println(v1, fi46)
            _ = fi46
            _ = v47
        }
        _ = len(v2)
    }
    v1 -= v1
    v2["foo bar"] = "a"
    if (!(9223372036854775807 <= 42)) {
        if (0 >= 2147483647) {
            switch 1 {
            case 221:
                var v53 int64 = 0
                _ = v53
            case 2147483647:
                fmt.Println(v1)
            case 1:
                var v54 int16 = -50
                _ = v54
            default:
                switch 944 {
                case 9223372036854775807:
                    switch ((("\n\t" + "\n\t") + "") + "hello") {
                    case "hello":
                        var v55 uint32 = (v1 >> ((65535) & 63))
                        _ = v55
                    case "foo bar":
                        v1 += (^(v1))
                    default:
                        var v56 int16 = -740
                        _ = v56
                    }
                default:
                    for fi57 := 0; fi57 < 12; fi57++ {
                        if len(v4) > 0 {
                            v4[0] = (fi57 % ((fi57) | 1))
                        }
                        var v58 uint8 = uint8(v1)
                        _ = fi57
                        _ = v58
                    }
                }
            }
            switch -1 {
            case 372:
                if true {
                    v1 = v1
                }
            default:
                for fi59 := 0; fi59 < 18; fi59++ {
                    var v60 int16 = 0
                    _ = fi59
                    _ = v60
                }
            }
        } else {
            var v61 uint8 = 128
            _ = v61
        }
        v1 += v1
    } else {
        switch 32767 {
        case 32767:
            v62 := GenS7{F0: 209, F1: 173, F2: 643}
            _ = v62
        default:
            if false {
                switch "" {
                case "hello":
                    if true {
                        var v63 uint8 = 42
                        _ = v63
                    }
                }
            } else {
                fmt.Println(v1)
                _ = len(v4)
            }
        }
    }
    var v64 bool = true
    if v64 {
        v65 := GenS2{F0: 4}
        for fi66 := 0; fi66 < 14; fi66++ {
            if v64 {
                if (((!(1e10 == ((1.5 * ((float64(1.5) + 0.0) * float64(1e10))) / -2.25))) || v64) && v64) {
                    for fi67 := 0; fi67 < 7; fi67++ {
                        var v68 int32 = -629
                        switch (fi66 | fi66) {
                        case -158:
                            delete(v3, (fi67 & -9223372036854775808))
                        case 4294967295:
                            v69 := v3[fi66]
                            _ = v69
                        case -1:
                            var v70 uint16 = 65535
                            _ = v70
                        default:
                            var v71 int64 = -287
                            _ = v71
                        }
                        _ = fi67
                        _ = v68
                    }
                    v72 := GenS3{F0: 1, F1: 104, F2: (770 > 0)}
                    for fi73 := 0; fi73 < 11; fi73++ {
                        v74 := GenS1{F0: 1, F1: 42}
                        if len(v4) >= 2 {
                            v4 = v4[1:len(v4)]
                        }
                        _ = fi73
                        _ = v74
                    }
                    _ = v72
                }
                v64 = ((42 <= 42) && (("" + "a") == ("a" + ("a" + (("hello" + "hello") + ("a" + ("hello" + "")))))))
            }
            for fi75 := 0; fi75 < 11; fi75++ {
                if (273 <= 42) {
                    var v76 uint8 = uint8(v1)
                    _ = v76
                }
                v3[128] = v1
                _ = fi75
            }
            _ = fi66
        }
        switch "\n\t" {
        case "foo bar":
            v64 = (9223372036854775807 <= 9223372036854775807)
        case "":
            fmt.Println(v1, v64)
        case "hello":
            if v64 {
                v1 = 65535
            } else {
                v77 := GenS2{F0: 42}
                for fi78 := 0; fi78 < 7; fi78++ {
                    v64 = (2 != 127)
                    _ = fi78
                }
                v64 = v64
                _ = v77
            }
        }
        _ = v65
    } else {
        switch 42 {
        case -9223372036854775808:
            var v79 int8 = -112
            _ = v79
        case -194:
            v64 = true
        case 448:
            for fi80 := 0; fi80 < 3; fi80++ {
                fmt.Println(v1)
                _ = fi80
            }
        }
        var v81 uint64 = 515
        _ = v81
    }
    v82 := GenS5{F0: "foo bar", F1: 0, F2: 2}
    if len(v4) >= 2 {
        v4 = v4[1:len(v4)]
    }
    switch 863 {
    case 502:
        var v83 int = 0
        _ = v83
    case 2147483647:
        if ((v64 || v64) && ((false && (1 > 42)) || v64)) {
            fmt.Println(v64)
            delete(v2, (("\n\t" + "\n\t") + ""))
        } else {
            var v84 bool = false
            if (v84 && v84) {
                fmt.Println(v84)
                fmt.Println(v1, v1, v64)
            }
            var v85 uint32 = 501
            _ = v84
            _ = v85
        }
    case 127:
        var v86 uint64 = 1
        _ = v86
    }
    var v87 int8 = 42
    var v88 uint = 4294967295
    var v89 uint16 = 0
    fmt.Println(v64)
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v64:", v64)
    fmt.Println("v87:", v87)
    fmt.Println("v88:", v88)
    fmt.Println("v89:", v89)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v64
    _ = v82
    _ = v87
    _ = v88
    _ = v89
}
