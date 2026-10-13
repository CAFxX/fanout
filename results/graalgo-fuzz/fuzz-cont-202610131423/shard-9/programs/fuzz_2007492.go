package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 string; F2 uint }
type GenS2 struct { F0 uint32; F1 string; F2 bool }
type GenS3 struct { F0 int8 }
type GenS4 struct { F0 uint16; F1 int16; F2 uint }
type GenS5 struct { F0 uint32 }
type GenS6 struct { F0 uint }
type GenS7 struct { F0 bool }
type GenS8 struct { F0 uint16; F1 int32; F2 bool }
type GenS9 struct { F0 uint32; F1 string; F2 int }
type GenS10 struct { F0 int; F1 uint64 }
type GenS11 struct { F0 int8 }
type GenS12 struct { F0 uint; F1 uint64; F2 int8 }
type GenS13 struct { F0 int64; F1 uint16 }

func main() {
    var v1 int64 = -1
    var v2 uint8 = 185
    v3 := []string{"\n\t", "foo bar"}
    v4 := make(map[string]uint32)
    if (-681 == -32768) {
        v4["\n\t"] = 0
        v2 = (v2 ^ 42)
    }
    fmt.Println(v1, v2)
    var v5 uint32 = 65535
    if false {
        var v6 uint64 = 4294967295
        switch 127 {
        case -458:
            var v7 uint32 = v5
            _ = v7
        default:
            for fi8 := 0; fi8 < 13; fi8++ {
                fmt.Println(fi8)
                _ = fi8
            }
        }
        if true {
            v1 = (v1 * 1)
            v2 = (v2 << ((440) & 63))
            for fi9 := 0; fi9 < 3; fi9++ {
                if ((fi9 >> ((2) & 63)) == (fi9 & (fi9 >> ((999) & 63)))) {
                    v10 := GenS1{F0: -128, F1: "hello", F2: 636}
                    v11 := GenS2{F0: (v5 | (^(v5))), F1: ((("\n\t" + ("foo bar" + ((("hello" + (("" + "\n\t") + (("hello" + ("" + "")) + (("hello" + "foo bar") + "foo bar")))) + "foo bar") + ("\n\t" + "hello")))) + ("" + "a")) + "foo bar"), F2: false}
                    _ = v10
                    _ = v11
                }
                _ = fi9
            }
        }
        _ = v6
    } else {
        v5 = v5
        if true {
            v12 := GenS2{F0: (v5 | v5), F1: "\n\t", F2: true}
            v5 = v5
            var v13 uint32 = v5
            _ = v12
            _ = v13
        } else {
            v14 := GenS2{F0: 2147483647, F1: "\n\t", F2: (v1 >= (v1 * (v1 + (v1 / (((v1 << ((65535) & 63))) | 1)))))}
            _ = v14
        }
    }
    v5 = (v5 + (v5 + (v5 << ((2) & 63))))
    if (!(42 > 1)) {
        fmt.Println(v1, v1)
    } else {
        fmt.Println(v5, v2, v1)
        for fi15 := 0; fi15 < 19; fi15++ {
            if true {
                v2 = v2
            } else {
                fmt.Println(fi15)
            }
            if len(v3) >= 2 {
                v3 = v3[1:len(v3)]
            }
            for fi16 := 0; fi16 < 19; fi16++ {
                if (18446744073709551615 >= 255) {
                    v17 := GenS1{F0: int8(fi16), F1: "foo bar", F2: uint(v5)}
                    _ = v17
                }
                fmt.Println(v2)
                if false {
                    _ = len(v3)
                    var v18 int = 32767
                    v19 := GenS1{F0: 118, F1: ("\n\t" + ("hello" + "foo bar")), F2: 0}
                    if len(v3) > 0 {
                        v3[0] = "\n\t"
                    }
                    switch "" {
                    case "":
                        switch "\n\t" {
                        case "foo bar":
                            var v20 uint64 = 1
                            _ = v20
                        case "":
                            var v21 uint32 = (v5 * v5)
                            _ = v21
                        case "hello":
                            v22 := GenS2{F0: v5, F1: "a", F2: ((fi15 >> ((337) & 63)) == 1)}
                            _ = v22
                        default:
                            if (4294967295 <= 0) {
                                v18 -= fi16
                                v23 := v4[("a" + "foo bar")]
                                var v24 uint32 = 268
                                fmt.Println(v23, fi16, v23)
                                _ = v23
                                _ = v24
                            }
                        }
                    }
                    _ = v18
                    _ = v19
                } else {
                    v5 -= v5
                }
                _ = fi16
            }
            _ = fi15
        }
    }
    var v25 bool = false
    _ = len(v4)
    v25 = (-937 < 449)
    for fi26 := 0; fi26 < 4; fi26++ {
        if false {
            var v27 int16 = -783
            v28 := GenS2{F0: (v5 * 570), F1: ("" + "\n\t"), F2: (v5 <= 775)}
            _ = v27
            _ = v28
        }
        _ = fi26
    }
    v3 = append(v3, "foo bar")
    fmt.Println(v25)
    for fi29 := 0; fi29 < 2; fi29++ {
        var v30 int8 = 100
        delete(v4, "")
        fmt.Println(v2)
        _ = fi29
        _ = v30
    }
    if (0.1 < 1e10) {
        switch ("a" + "\n\t") {
        case "":
            if len(v3) > 0 {
                v3[0] = ""
            }
        default:
            if v25 {
                fmt.Println(v5, v5)
                _ = len(v3)
                var v31 int32 = -2147483648
                _ = v31
            } else {
                v5 = (v5 + (v5 + v5))
                if v25 {
                    switch 0 {
                    case 317:
                        v32 := GenS1{F0: -1, F1: "foo bar", F2: 890}
                        _ = v32
                    case 42:
                        for fi33 := 0; fi33 < 4; fi33++ {
                            v34 := GenS3{F0: 1}
                            v35 := GenS2{F0: (v5 - (v5 % ((uint32(v2)) | 1))), F1: (("" + "\n\t") + "a"), F2: (false || v25)}
                            _ = fi33
                            _ = v34
                            _ = v35
                        }
                    }
                }
            }
        }
        v2 = (v2 >> ((2) & 63))
        v2 = v2
    } else {
        if (!(32767 == -626)) {
            var v36 uint32 = v5
            _ = v36
        } else {
            if v25 {
                v5 = 42
                switch -322 {
                case -25:
                    for fi37 := 0; fi37 < 6; fi37++ {
                        var v38 int8 = 0
                        _ = fi37
                        _ = v38
                    }
                default:
                    if ((3.14159 + 0.0) > -2.25) {
                        if (9223372036854775807 <= (v1 | v1)) {
                            var v39 int32 = 2147483647
                            for fi40 := 0; fi40 < 14; fi40++ {
                                var v41 int32 = v39
                                for fi42 := 0; fi42 < 4; fi42++ {
                                    if (0 == 169) {
                                        v43 := GenS1{F0: 1, F1: "a", F2: 668}
                                        if true {
                                            if len(v3) > 0 {
                                                v3[0] = "foo bar"
                                            }
                                            if v25 {
                                                switch "foo bar" {
                                                case "a":
                                                    var v44 int = (fi40 / ((fi40) | 1))
                                                    _ = v44
                                                case "hello":
                                                    if len(v3) >= 2 {
                                                        v3 = v3[1:len(v3)]
                                                    }
                                                case "\n\t":
                                                    v45 := GenS1{F0: 0, F1: ("hello" + "foo bar"), F2: 4294967295}
                                                    _ = v45
                                                default:
                                                    if true {
                                                        var v46 int32 = -525
                                                        if (((("" + "\n\t") + "a") + (("hello" + "hello") + "a")) < ("hello" + "foo bar")) {
                                                            if (v25 && (2 != 0)) {
                                                                if false {
                                                                    v47 := GenS4{F0: uint16(v5), F1: 715, F2: 2}
                                                                    _ = v47
                                                                }
                                                                var v48 int = (-(fi42))
                                                                _ = v48
                                                            } else {
                                                                if (float32(((float64(-2.25) + 0.0) * 1.5)) < 0.1) {
                                                                    v25 = false
                                                                    v49 := GenS2{F0: v5, F1: "\n\t", F2: (!(!true))}
                                                                    _ = v49
                                                                } else {
                                                                    v50 := v4[("a" + ("foo bar" + ("a" + "hello")))]
                                                                    var v51 int8 = 42
                                                                    if true {
                                                                        if (811 < 65535) {
                                                                            if (v25 && v25) {
                                                                                v52 := GenS5{F0: (^(v50))}
                                                                                if v25 {
                                                                                    v53 := v4["a"]
                                                                                    _ = v53
                                                                                }
                                                                                v54 := GenS6{F0: 377}
                                                                                _ = v52
                                                                                _ = v54
                                                                            } else {
                                                                                if (0 <= 127) {
                                                                                    if (!((-(v41)) >= v41)) {
                                                                                        if (v1 != v1) {
                                                                                            switch (fi40 << ((65535) & 63)) {
                                                                                            case 9223372036854775807:
                                                                                                if (((1 > 666) && true) || false) {
                                                                                                    var v55 uint16 = 255
                                                                                                    var v56 bool = v25
                                                                                                    switch (fi40 + (^(fi40))) {
                                                                                                    case 0:
                                                                                                        var v57 bool = (42 > (v51 << ((uint(v50)) & 63)))
                                                                                                        _ = v57
                                                                                                    case -2147483648:
                                                                                                        v58 := v4[((("hello" + "") + ("hello" + "\n\t")) + "a")]
                                                                                                        _ = v58
                                                                                                    }
                                                                                                    _ = v55
                                                                                                    _ = v56
                                                                                                }
                                                                                            default:
                                                                                                fmt.Println(fi42, v2)
                                                                                            }
                                                                                            delete(v4, "\n\t")
                                                                                        } else {
                                                                                            if ("a" == ("a" + "a")) {
                                                                                                var v59 uint8 = 127
                                                                                                v60 := GenS7{F0: false}
                                                                                                _ = v59
                                                                                                _ = v60
                                                                                            }
                                                                                            var v61 bool = (!true)
                                                                                            v5 -= (^(v5))
                                                                                            _ = v61
                                                                                        }
                                                                                    }
                                                                                    var v62 uint = 179
                                                                                    var v63 int16 = 0
                                                                                    _ = v62
                                                                                    _ = v63
                                                                                } else {
                                                                                    v5 = v50
                                                                                }
                                                                            }
                                                                            var v64 int32 = (v41 | (v41 >> ((uint(v2)) & 63)))
                                                                            var v65 bool = (695 < 255)
                                                                            _ = v64
                                                                            _ = v65
                                                                        } else {
                                                                            if ((v39 >= (v46 & (v46 | 489))) || false) {
                                                                                v41 += (^(v41))
                                                                                v66 := GenS7{F0: ((397 > 255) || (v25 || v25))}
                                                                                v5 += v50
                                                                                _ = v66
                                                                            } else {
                                                                                if (!("" != "hello")) {
                                                                                    fmt.Println(fi42, v51, v2)
                                                                                    if v25 {
                                                                                        v67 := GenS8{F0: uint16(v2), F1: v41, F2: v25}
                                                                                        v1 -= (v1 % ((int64(v50)) | 1))
                                                                                        fmt.Println(v2, v1, v25)
                                                                                        _ = v67
                                                                                    }
                                                                                }
                                                                                if v25 {
                                                                                    if (!((v2 ^ (v2 + (v2 / ((v2) | 1)))) != 128)) {
                                                                                        v41 -= (^(v46))
                                                                                        switch "hello" {
                                                                                        case "a":
                                                                                            fmt.Println(fi42, v5)
                                                                                        case "":
                                                                                            v5 -= v50
                                                                                        case "foo bar":
                                                                                            fmt.Println(v1, fi40)
                                                                                        default:
                                                                                            v4["foo bar"] = v5
                                                                                            v68 := GenS9{F0: 65535, F1: ("" + ("foo bar" + ("hello" + ((("foo bar" + "\n\t") + "foo bar") + ("" + "\n\t"))))), F2: fi42}
                                                                                            _ = v68
                                                                                        }
                                                                                    }
                                                                                    switch (("a" + (((("\n\t" + "a") + "\n\t") + "\n\t") + "\n\t")) + "foo bar") {
                                                                                    case "foo bar":
                                                                                        v2 = v2
                                                                                    case "\n\t":
                                                                                        if v25 {
                                                                                            fmt.Println(v39)
                                                                                        }
                                                                                    }
                                                                                    v69 := GenS2{F0: (v50 * v50), F1: ("foo bar" + (("\n\t" + "hello") + ("foo bar" + "a"))), F2: (2 > 691)}
                                                                                    _ = v69
                                                                                } else {
                                                                                    switch ("foo bar" + "hello") {
                                                                                    case "foo bar":
                                                                                        if len(v3) > 0 {
                                                                                            v3[0] = "hello"
                                                                                        }
                                                                                    case "a":
                                                                                        if v25 {
                                                                                            v50 -= v5
                                                                                            v70 := GenS10{F0: fi42, F1: 1}
                                                                                            _ = v70
                                                                                        } else {
                                                                                            if v25 {
                                                                                                delete(v4, "a")
                                                                                                fmt.Println(v41, v39, v2)
                                                                                                v4[""] = 284
                                                                                            } else {
                                                                                                var v71 int16 = 117
                                                                                                fmt.Println(fi40)
                                                                                                _ = v71
                                                                                            }
                                                                                        }
                                                                                    case "":
                                                                                        fmt.Println(v5, v25)
                                                                                    default:
                                                                                        v41 -= v39
                                                                                    }
                                                                                    var v72 int = (fi40 * 534)
                                                                                    var v73 uint64 = 4294967295
                                                                                    _ = v72
                                                                                    _ = v73
                                                                                }
                                                                                v2 = v2
                                                                            }
                                                                            fmt.Println(v2)
                                                                            if (1 != 460) {
                                                                                v74 := GenS11{F0: (v51 % ((v51) | 1))}
                                                                                _ = v74
                                                                            } else {
                                                                                if (!true) {
                                                                                    v39 += (^(v39))
                                                                                } else {
                                                                                    if v25 {
                                                                                        var v75 bool = false
                                                                                        _ = v75
                                                                                    }
                                                                                    v50 -= 4294967295
                                                                                    v1 = int64(v51)
                                                                                }
                                                                            }
                                                                            if len(v3) >= 2 {
                                                                                v3 = v3[1:len(v3)]
                                                                            }
                                                                        }
                                                                        _ = len(v4)
                                                                        v46 = v41
                                                                    }
                                                                    _ = v50
                                                                    _ = v51
                                                                }
                                                            }
                                                            var v76 uint16 = 894
                                                            _ = v76
                                                        }
                                                        _ = v46
                                                    } else {
                                                        switch "foo bar" {
                                                        case "foo bar":
                                                            v5 += 229
                                                        }
                                                    }
                                                    v77 := v4["foo bar"]
                                                    v41 += -1
                                                    fmt.Println(fi40)
                                                    _ = v77
                                                }
                                                v78 := GenS10{F0: (fi42 * -159), F1: 42}
                                                _ = v78
                                            }
                                            var v79 int64 = (^(v1))
                                            _ = v79
                                        }
                                        _ = v43
                                    }
                                    v2 = (^(v2))
                                    var v80 uint64 = 992
                                    v81 := GenS12{F0: 2, F1: (v80 / (((v80 << ((1) & 63))) | 1)), F2: -128}
                                    var v82 uint8 = v2
                                    _ = fi42
                                    _ = v80
                                    _ = v81
                                    _ = v82
                                }
                                v25 = ((0 >= 42) || v25)
                                _ = fi40
                                _ = v41
                            }
                            v2 = v2
                            _ = v39
                        } else {
                            v4[""] = v5
                        }
                        delete(v4, ("" + "\n\t"))
                        v83 := GenS1{F0: 42, F1: "", F2: 0}
                        _ = v83
                    } else {
                        var v84 uint16 = 0
                        var v85 int16 = 0
                        v1 = v1
                        _ = v84
                        _ = v85
                    }
                }
                fmt.Println(v5, v2, v1)
            } else {
                var v86 int = 4294967295
                for fi87 := 0; fi87 < 3; fi87++ {
                    switch "hello" {
                    case "\n\t":
                        var v88 uint8 = (v2 * v2)
                        _ = v88
                    case "hello":
                        _ = len(v3)
                    }
                    _ = fi87
                }
                if ((v86 % (((^(v86))) | 1)) < v86) {
                    v25 = false
                    if ("foo bar" > (((("a" + "\n\t") + "foo bar") + "") + ((("\n\t" + "\n\t") + ("\n\t" + "foo bar")) + ("" + ("hello" + "\n\t"))))) {
                        var v89 int64 = v1
                        _ = v89
                    } else {
                        _ = len(v3)
                        if (!(-1 == 1)) {
                            var v90 float64 = 3.14159
                            v91 := GenS9{F0: v5, F1: "foo bar", F2: v86}
                            v92 := GenS13{F0: (v1 << ((uint(v2)) & 63)), F1: 0}
                            _ = v90
                            _ = v91
                            _ = v92
                        }
                    }
                }
                _ = v86
            }
            fmt.Println(v25, v5, v25)
            if (42 > v1) {
                v2 = v2
                v93 := v4["foo bar"]
                var v94 uint32 = (v5 + v93)
                _ = v93
                _ = v94
            }
        }
    }
    fmt.Println(v25, v2)
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v25:", v25)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v25
    _ = v3
    _ = v4
    _ = v5
}
