package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 int; F2 uint }
type GenS2 struct { F0 string }
type GenS3 struct { F0 int8; F1 uint64 }
type GenS4 struct { F0 int8 }
type GenS5 struct { F0 int; F1 uint16 }
type GenS6 struct { F0 uint8; F1 uint64 }
type GenS7 struct { F0 uint16 }
type GenS8 struct { F0 int64; F1 uint16 }
type GenS9 struct { F0 int; F1 uint8; F2 uint64 }
type GenS10 struct { F0 uint32 }

func main() {
    v1 := []uint32{0}
    var v2 uint32 = 4294967295
    v3 := make(map[int]bool)
    var v4 int8 = 1
    var v5 int8 = (v4 & v4)
    v6 := GenS1{F0: 42, F1: 321, F2: 291}
    var v7 int8 = 1
    switch -9223372036854775808 {
    case 893:
        v4 = v4
    case 0:
        switch "\n\t" {
        case "foo bar":
            if (v7 == 65) {
                if ("foo bar" >= "hello") {
                    switch -919 {
                    case 473:
                        var v8 uint32 = v2
                        _ = v8
                    case 128:
                        fmt.Println(v4)
                    default:
                        fmt.Println(v5, v7)
                    }
                } else {
                    var v9 int16 = 32767
                    if (42 < 42) {
                        v9 = v9
                    } else {
                        fmt.Println(v4, v5)
                    }
                    var v10 string = "foo bar"
                    _ = v10
                    _ = v9
                }
                var v11 string = "hello"
                _ = v11
            }
        case "hello":
            var v12 int = -319
            _ = v12
        case "a":
            var v13 uint32 = v2
            _ = v13
        default:
            for fi14 := 0; fi14 < 20; fi14++ {
                v15 := GenS1{F0: 9223372036854775807, F1: int(v2), F2: 707}
                v4 += -26
                v16 := GenS2{F0: ("a" + "hello")}
                _ = fi14
                _ = v15
                _ = v16
            }
        }
    case 4294967295:
        if true {
            v4 = v7
        } else {
            var v17 string = (("hello" + (("hello" + "hello") + "\n\t")) + "")
            _ = v17
        }
    default:
        _ = len(v3)
    }
    v7 = v5
    if true {
        v18 := v3[2147483647]
        fmt.Println(v4)
        _ = v18
    }
    for fi19 := 0; fi19 < 18; fi19++ {
        v20 := GenS3{F0: (v5 & v7), F1: 4294967295}
        var v21 int64 = 9223372036854775807
        _ = fi19
        _ = v20
        _ = v21
    }
    if true {
        for fi22 := 0; fi22 < 15; fi22++ {
            if (661 <= 1) {
                v4 = 127
            }
            _ = fi22
        }
        if (348 >= -513) {
            if true {
                switch 32767 {
                case 979:
                    var v23 string = "foo bar"
                    _ = v23
                case -2147483648:
                    v4 = 108
                case 4294967295:
                    v24 := GenS4{F0: (v5 - (v7 - (v5 | v4)))}
                    _ = v24
                default:
                    switch 32767 {
                    case 32767:
                        v2 = (^(v2))
                    case 0:
                        v25 := GenS3{F0: v7, F1: 917}
                        _ = v25
                    }
                }
                v7 = v5
            } else {
                if false {
                    var v26 uint16 = 32767
                    var v27 uint32 = 2147483647
                    _ = v26
                    _ = v27
                } else {
                    for fi28 := 0; fi28 < 6; fi28++ {
                        switch (fi28 ^ 42) {
                        case 0:
                            for fi29 := 0; fi29 < 1; fi29++ {
                                v2 = v2
                                _ = fi29
                            }
                        case 457:
                            var v30 float32 = 0.1
                            _ = v30
                        }
                        _ = fi28
                    }
                }
                if ("\n\t" > "") {
                    for fi31 := 0; fi31 < 5; fi31++ {
                        _ = len(v3)
                        _ = fi31
                    }
                } else {
                    var v32 uint8 = 39
                    v7 = (v5 / ((v4) | 1))
                    var v33 float32 = (((float32(float64(1.5)) - 0.1) - float32((-2.25 - 0.1))) - 0.0)
                    _ = v32
                    _ = v33
                }
                var v34 float32 = 1e10
                _ = v34
            }
            if ((v2 - (v2 << ((0) & 63))) > v2) {
                var v35 int8 = 1
                v7 = -4
                _ = v35
            } else {
                v36 := GenS3{F0: (v4 + (^(v4))), F1: 951}
                var v37 uint16 = 2
                _ = v36
                _ = v37
            }
            for fi38 := 0; fi38 < 1; fi38++ {
                var v39 uint32 = (v2 << ((4294967295) & 63))
                if ((false && (uint16(v2) != 255)) && (v2 <= 796)) {
                    for fi40 := 0; fi40 < 1; fi40++ {
                        v39 = v39
                        v2 = 42
                        v2 = (v2 % (((v2 * v39)) | 1))
                        _ = fi40
                    }
                    var v41 int16 = -124
                    _ = v41
                } else {
                    var v42 bool = (65535 == 255)
                    v7 = (v7 | v4)
                    _ = v42
                }
                var v43 float64 = -2.25
                _ = fi38
                _ = v39
                _ = v43
            }
        } else {
            var v44 uint8 = 171
            if ((!false) && ((((float32(1.5) + (1e10 + (0.0 * (0.0 * 1e10)))) / 3.14159) + -2.25) > 0.1)) {
                v2 -= 2
                for fi45 := 0; fi45 < 12; fi45++ {
                    if len(v1) >= 2 {
                        v1 = v1[1:len(v1)]
                    }
                    fmt.Println(v7, v7, v4)
                    v46 := v3[fi45]
                    _ = fi45
                    _ = v46
                }
            } else {
                for fi47 := 0; fi47 < 3; fi47++ {
                    v4 += v7
                    _ = fi47
                }
            }
            _ = v44
        }
    } else {
        if (false && true) {
            for fi48 := 0; fi48 < 14; fi48++ {
                if false {
                    for fi49 := 0; fi49 < 8; fi49++ {
                        var v50 bool = false
                        if v50 {
                            if (-412 == 42) {
                                v50 = v50
                                fmt.Println(fi48, v4)
                            } else {
                                v7 = (-(v7))
                                var v51 uint32 = 65535
                                v51 += (v51 % ((v51) | 1))
                                var v52 int = (fi48 & fi48)
                                if v50 {
                                    v2 += v2
                                    var v53 uint = 4294967295
                                    _ = v53
                                } else {
                                    var v54 float32 = 1.5
                                    var v55 uint64 = 53
                                    switch ("" + (("hello" + "\n\t") + "\n\t")) {
                                    case "hello":
                                        switch "\n\t" {
                                        case "hello":
                                            switch (fi49 ^ fi48) {
                                            case -2147483648:
                                                v1 = append(v1, v2)
                                            case -1:
                                                v5 += (v7 ^ v5)
                                            case 1:
                                                if (75 > 2) {
                                                    var v56 int8 = v7
                                                    v52 += v52
                                                    v57 := GenS3{F0: 0, F1: v55}
                                                    _ = v56
                                                    _ = v57
                                                } else {
                                                    if (2 >= 0) {
                                                        var v58 int32 = 942
                                                        _ = v58
                                                    } else {
                                                        v55 = v55
                                                    }
                                                    v52 -= (-(fi49))
                                                }
                                            default:
                                                v2 = (v2 >> ((65535) & 63))
                                            }
                                        case "":
                                            v59 := GenS5{F0: (fi49 - (fi48 % (((fi48 - 9223372036854775807)) | 1))), F1: uint16(v55)}
                                            if v50 {
                                                fmt.Println(v4, v5, v54)
                                            } else {
                                                if true {
                                                    switch "\n\t" {
                                                    case "hello":
                                                        var v60 int16 = 32767
                                                        _ = v60
                                                    }
                                                    v7 += (v7 ^ -128)
                                                    switch fi48 {
                                                    case 0:
                                                        if (0 > 32767) {
                                                            if false {
                                                                if true {
                                                                    v55 -= 748
                                                                    switch "hello" {
                                                                    case "hello":
                                                                        var v61 float32 = v54
                                                                        _ = v61
                                                                    case "a":
                                                                        fmt.Println(fi48, v5, v51)
                                                                    case "":
                                                                        v62 := GenS6{F0: 255, F1: (^(v55))}
                                                                        _ = v62
                                                                    default:
                                                                        v54 = (0.0 * ((v54 * (0.1 + 3.14159)) + 3.14159))
                                                                    }
                                                                }
                                                                var v63 int64 = -1
                                                                _ = v63
                                                            } else {
                                                                var v64 float64 = (float64(1e10) / ((3.14159 - 1e10) - 3.14159))
                                                                _ = v64
                                                            }
                                                        }
                                                    case -339:
                                                        if true {
                                                            v50 = v50
                                                            v5 -= v7
                                                            fmt.Println(v51)
                                                        }
                                                    case 2147483647:
                                                        if len(v1) >= 2 {
                                                            v1 = v1[1:len(v1)]
                                                        }
                                                    }
                                                }
                                                v2 = v51
                                                v4 -= 111
                                            }
                                            _ = v59
                                        case "foo bar":
                                            v65 := GenS3{F0: 64, F1: (^(v55))}
                                            _ = v65
                                        }
                                    default:
                                        var v66 uint32 = v51
                                        _ = v66
                                    }
                                    _ = v54
                                    _ = v55
                                }
                                if v50 {
                                    if (!true) {
                                        var v67 bool = v50
                                        v52 = (^(v52))
                                        _ = v67
                                    } else {
                                        v68 := GenS3{F0: int8(v2), F1: 2}
                                        _ = v68
                                    }
                                    switch "hello" {
                                    case "hello":
                                        var v69 uint = 522
                                        _ = v69
                                    case "":
                                        fmt.Println(v50, v51)
                                    case "a":
                                        if len(v1) >= 2 {
                                            v1 = v1[1:len(v1)]
                                        }
                                    }
                                    v70 := GenS3{F0: v4, F1: 42}
                                    _ = v70
                                }
                                _ = v51
                                _ = v52
                            }
                            if (false && (21 > 128)) {
                                if (!v50) {
                                    var v71 int = (fi48 | (fi49 | (-(fi49))))
                                    switch "" {
                                    case "foo bar":
                                        switch "foo bar" {
                                        case "hello":
                                            if false {
                                                fmt.Println(v71)
                                            }
                                        case "":
                                            if (v50 && true) {
                                                if v50 {
                                                    if v50 {
                                                        var v72 int8 = (v5 | 3)
                                                        _ = v72
                                                    }
                                                    fmt.Println(fi48)
                                                }
                                                v50 = false
                                            }
                                        }
                                    default:
                                        switch (("hello" + "") + "foo bar") {
                                        case "hello":
                                            switch fi49 {
                                            case 9223372036854775807:
                                                if (!(-1 != 2147483647)) {
                                                    v2 -= (v2 + 2)
                                                } else {
                                                    v73 := GenS7{F0: 2}
                                                    if v50 {
                                                        var v74 int64 = 0
                                                        _ = v74
                                                    }
                                                    if (v71 > 9223372036854775807) {
                                                        v75 := GenS6{F0: 2, F1: 197}
                                                        v2 -= (v2 % ((634) | 1))
                                                        if (-2.25 > (-2.25 - 1e10)) {
                                                            if false {
                                                                var v76 int16 = -793
                                                                _ = len(v3)
                                                                _ = v76
                                                            }
                                                        } else {
                                                            var v77 int64 = -1
                                                            var v78 int32 = 0
                                                            _ = v77
                                                            _ = v78
                                                        }
                                                        _ = v75
                                                    }
                                                    _ = v73
                                                }
                                                switch (v71 >> ((65535) & 63)) {
                                                case 127:
                                                    v3[fi49] = false
                                                default:
                                                    if ((!(3.14159 >= 0.0)) || (v50 || (42 < 0))) {
                                                        var v79 bool = (!(v50 && v50))
                                                        delete(v3, 9223372036854775807)
                                                        v80 := GenS5{F0: (fi49 / ((127) | 1)), F1: 0}
                                                        _ = v79
                                                        _ = v80
                                                    } else {
                                                        v2 += 2147483648
                                                        switch (fi49 << ((1) & 63)) {
                                                        case 741:
                                                            v3[fi49] = false
                                                        case 42:
                                                            delete(v3, 2147483647)
                                                        case 507:
                                                            var v81 int32 = -726
                                                            _ = v81
                                                        default:
                                                            v5 = (v5 + (v7 & (-(v7))))
                                                        }
                                                        v5 = (v4 ^ -128)
                                                        v7 += -1
                                                    }
                                                    v50 = v50
                                                }
                                                v82 := GenS6{F0: 42, F1: 2}
                                                v5 += v5
                                                _ = v82
                                            case -1:
                                                v83 := GenS2{F0: "hello"}
                                                _ = v83
                                            case 2147483647:
                                                var v84 uint8 = 157
                                                _ = v84
                                            default:
                                                v85 := GenS8{F0: -1, F1: 104}
                                                _ = v85
                                            }
                                        case "\n\t":
                                            if len(v1) > 0 {
                                                v1[0] = v2
                                            }
                                        case "a":
                                            var v86 uint32 = (v2 & 2)
                                            _ = v86
                                        default:
                                            var v87 uint64 = 522
                                            _ = v87
                                        }
                                    }
                                    _ = v71
                                }
                                if false {
                                    v7 -= int8(v2)
                                } else {
                                    var v88 int8 = v4
                                    delete(v3, fi48)
                                    v89 := GenS9{F0: (fi48 * (-(fi48))), F1: 255, F2: 18446744073709551615}
                                    _ = v88
                                    _ = v89
                                }
                            }
                        }
                        v3[(-(fi49))] = v50
                        _ = fi49
                        _ = v50
                    }
                    switch "hello" {
                    case "hello":
                        for fi90 := 0; fi90 < 9; fi90++ {
                            var v91 float32 = (3.14159 - 0.0)
                            _ = fi90
                            _ = v91
                        }
                    }
                    switch "hello" {
                    case "a":
                        v92 := GenS10{F0: v2}
                        _ = v92
                    case "hello":
                        v7 -= v7
                    }
                } else {
                    var v93 string = (("foo bar" + ("foo bar" + ("a" + ""))) + ((((("" + "foo bar") + "hello") + "\n\t") + ("hello" + ("a" + "foo bar"))) + ""))
                    v94 := v3[fi48]
                    _ = v93
                    _ = v94
                }
                _ = fi48
            }
        } else {
            if true {
                for fi95 := 0; fi95 < 18; fi95++ {
                    for fi96 := 0; fi96 < 10; fi96++ {
                        var v97 int8 = v5
                        v4 = v5
                        var v98 int16 = 401
                        _ = fi96
                        _ = v97
                        _ = v98
                    }
                    v5 -= (-(v7))
                    _ = fi95
                }
                var v99 bool = true
                var v100 string = "a"
                _ = v100
                _ = v99
            }
            v2 -= 2147483648
        }
        var v101 int32 = -2147483648
        _ = v101
    }
    switch 127 {
    case 2147483647:
        v7 = (^(v5))
    default:
        v102 := GenS5{F0: 42, F1: 364}
        _ = v102
    }
    for fi103 := 0; fi103 < 13; fi103++ {
        v4 = 0
        _ = fi103
    }
    v104 := GenS8{F0: 1, F1: 1}
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v104
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
