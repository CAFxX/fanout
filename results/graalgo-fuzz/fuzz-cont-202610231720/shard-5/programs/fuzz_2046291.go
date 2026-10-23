package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 uint; F2 bool }
type GenS2 struct { F0 string; F1 int8; F2 int }
type GenS3 struct { F0 int16 }
type GenS4 struct { F0 int; F1 int; F2 int8 }
type GenS5 struct { F0 uint16; F1 uint64; F2 uint32 }
type GenS6 struct { F0 uint32; F1 uint8 }
type GenS7 struct { F0 string; F1 uint32; F2 uint8 }
type GenS8 struct { F0 int8 }
type GenS9 struct { F0 uint8; F1 string; F2 uint8 }
type GenS10 struct { F0 string }
type GenS11 struct { F0 int16 }
type GenS12 struct { F0 uint; F1 uint }
type GenS13 struct { F0 uint16 }
type GenS14 struct { F0 uint32 }
type GenS15 struct { F0 int; F1 bool; F2 int16 }
type GenS16 struct { F0 int32; F1 int8 }
type GenS17 struct { F0 int16 }
type GenS18 struct { F0 int8; F1 int32; F2 uint32 }
type GenS19 struct { F0 int16 }
type GenS20 struct { F0 int8; F1 uint16 }
type GenS21 struct { F0 bool }
type GenS22 struct { F0 uint; F1 uint64; F2 uint16 }
type GenS23 struct { F0 int32 }
type GenS24 struct { F0 int64; F1 uint16; F2 uint }
type GenS25 struct { F0 int32 }
type GenS26 struct { F0 uint32; F1 string; F2 int16 }
type GenS27 struct { F0 uint64; F1 uint32 }
func genFn0(p0 uint64, p1 uint32) uint8 {
    switch ("hello" + "\n\t") {
    case "\n\t":
        var v4 int8 = 1
        _ = v4
    case "hello":
        var v5 uint = 743
        _ = v5
    default:
        p1 += p1
    }
    p0 = (p0 + (p0 / (((p0 | 4294967295)) | 1)))
    if (-2147483648 >= -672) {
        v6 := GenS1{F0: "\n\t", F1: 18446744073709551615, F2: true}
        v7 := GenS1{F0: "foo bar", F1: 154, F2: true}
        _ = v6
        _ = v7
    }
    var v8 int8 = 127
    var v9 int16 = 1
    _ = p0
    _ = p1
    _ = v8
    _ = v9
    return 2
}
func genFn1(p0 string, p1 bool, p2 string) uint {
    v10 := []int{4294967295, 677, 128, 1}
    var v11 int32 = -338
    if true {
        switch p0 {
        case "\n\t":
            if len(v10) >= 2 {
                v10 = v10[1:len(v10)]
            }
        default:
            v12 := GenS2{F0: p2, F1: 42, F2: -489}
            _ = v12
        }
        if (-9223372036854775808 != 42) {
            var v13 uint16 = 86
            v11 = (v11 - v11)
            switch 332 {
            case 32768:
                for fi14 := 0; fi14 < 7; fi14++ {
                    if p1 {
                        switch 64 {
                        case -876:
                            if (!("a" < (p2 + (p2 + ("" + p0))))) {
                                fmt.Println(fi14)
                            } else {
                                v13 -= v13
                                switch p2 {
                                case "\n\t":
                                    if true {
                                        for fi15 := 0; fi15 < 16; fi15++ {
                                            v11 -= v11
                                            _ = fi15
                                        }
                                    }
                                case "":
                                    v16 := GenS3{F0: 32767}
                                    _ = v16
                                default:
                                    for fi17 := 0; fi17 < 13; fi17++ {
                                        v18 := GenS4{F0: (fi17 % (((fi17 / ((987) | 1))) | 1)), F1: -608, F2: 100}
                                        p0 = p2
                                        _ = fi17
                                        _ = v18
                                    }
                                }
                                _ = len(v10)
                            }
                        }
                        fmt.Println(fi14, v11, p0)
                    } else {
                        var v19 float64 = 3.14159
                        if true {
                            if (((-(v11)) >= (v11 | (v11 ^ 1))) || false) {
                                _ = len(v10)
                                switch (^(fi14)) {
                                case 946:
                                    var v20 uint = 223
                                    _ = v20
                                }
                            }
                            fmt.Println(fi14)
                        } else {
                            v19 += v19
                            if len(v10) >= 2 {
                                v10 = v10[1:len(v10)]
                            }
                            switch p0 {
                            case "\n\t":
                                for fi21 := 0; fi21 < 2; fi21++ {
                                    v13 = v13
                                    v22 := GenS1{F0: p2, F1: 65535, F2: ((fi21 < fi21) || p1)}
                                    if true {
                                        v23 := GenS5{F0: v13, F1: 0, F2: 42}
                                        var v24 string = (p0 + p0)
                                        _ = v23
                                        _ = v24
                                    } else {
                                        v25 := GenS6{F0: 42, F1: 127}
                                        if len(v10) >= 2 {
                                            v10 = v10[1:len(v10)]
                                        }
                                        _ = v25
                                    }
                                    _ = fi21
                                    _ = v22
                                }
                            default:
                                if true {
                                    if true {
                                        fmt.Println(v11, p1)
                                        for fi26 := 0; fi26 < 4; fi26++ {
                                            if p1 {
                                                var v27 float32 = float32(v19)
                                                _ = v27
                                            } else {
                                                _ = len(v10)
                                                v28 := GenS7{F0: p2, F1: 2, F2: 104}
                                                p2 = p0
                                                if len(v10) > 0 {
                                                    v10[0] = fi26
                                                }
                                                _ = v28
                                            }
                                            _ = fi26
                                        }
                                    } else {
                                        p1 = p1
                                    }
                                    if len(v10) >= 2 {
                                        v10 = v10[1:len(v10)]
                                    }
                                }
                            }
                        }
                        _ = v19
                    }
                    if (!p1) {
                        for fi29 := 0; fi29 < 10; fi29++ {
                            v30 := GenS8{F0: 127}
                            var v31 int8 = 96
                            _ = fi29
                            _ = v30
                            _ = v31
                        }
                        fmt.Println(p2, p1)
                        for fi32 := 0; fi32 < 19; fi32++ {
                            _ = len(v10)
                            p1 = p1
                            _ = fi32
                        }
                    }
                    _ = fi14
                }
            case 4294967295:
                p0 = "a"
            default:
                if p1 {
                    if true {
                        var v33 uint32 = 42
                        fmt.Println(v13, p1, p0)
                        fmt.Println(v33, v13)
                        _ = v33
                    }
                    for fi34 := 0; fi34 < 14; fi34++ {
                        for fi35 := 0; fi35 < 19; fi35++ {
                            v11 = v11
                            if p1 {
                                fmt.Println(fi35)
                                var v36 float64 = 3.14159
                                _ = v36
                            }
                            v11 = 829
                            if len(v10) > 0 {
                                v10[0] = 0
                            }
                            _ = fi35
                        }
                        var v37 float64 = (3.14159 - 3.14159)
                        _ = fi34
                        _ = v37
                    }
                }
            }
            _ = v13
        } else {
            for fi38 := 0; fi38 < 11; fi38++ {
                var v39 int8 = 100
                _ = fi38
                _ = v39
            }
            v10 = append(v10, 127)
            switch -1 {
            case 0:
                p1 = p1
            case 32767:
                p1 = p1
            }
        }
        fmt.Println(p2, p2, p1)
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v10
    _ = v11
    return 18446744073709551615
}

func main() {
    v1 := make(map[uint32]int)
    var v2 uint16 = 736
    var v3 uint8 = 255
    for fi40 := 0; fi40 < 4; fi40++ {
        v41 := GenS9{F0: 19, F1: "\n\t", F2: genFn0(2, 1)}
        _ = fi40
        _ = v41
    }
    var v42 int8 = 42
    v43 := v1[65535]
    v44 := GenS10{F0: "hello"}
    for fi45 := 0; fi45 < 18; fi45++ {
        var v46 uint = 4294967295
        for fi47 := 0; fi47 < 8; fi47++ {
            v48 := GenS11{F0: 0}
            v49 := GenS12{F0: v46, F1: (v46 - (^(v46)))}
            if true {
                if (!false) {
                    v1[2] = int(v46)
                    var v50 int64 = 511
                    var v51 float64 = -2.25
                    v51 -= float64((3.14159 + ((1e10 / (float32(v51) * float32((v51 + (v51 - 0.0))))) * 0.1)))
                    var v52 int64 = 0
                    v53 := GenS13{F0: (v2 >> ((v46) & 63))}
                    if (((!(v52 != (v50 + v52))) || true) || true) {
                        if true {
                            var v54 uint = genFn1(("hello" + "\n\t"), (((!(1000 < -32768)) && (127 != v3)) || (v3 >= v3)), (("a" + "hello") + ""))
                            if true {
                                fmt.Println(v54, v52)
                                v55 := GenS8{F0: (v42 * v42)}
                                var v56 int16 = 1000
                                v57 := GenS5{F0: v2, F1: uint64(v3), F2: 2}
                                _ = v55
                                _ = v56
                                _ = v57
                            } else {
                                v3 = v3
                                if (42 > int32(fi47)) {
                                    switch ("foo bar" + ("\n\t" + ("a" + "\n\t"))) {
                                    case "foo bar":
                                        v2 = 65535
                                    case "":
                                        if (859 == -32768) {
                                            fmt.Println(v54, fi47, v43)
                                            v3 = genFn0(1, 337)
                                            var v58 uint8 = v3
                                            _ = v58
                                        } else {
                                            v2 += v2
                                            if (2147483647 != 0) {
                                                if (((!true) && (true || (true || true))) && (!(!false))) {
                                                    var v59 bool = true
                                                    var v60 uint32 = 2147483648
                                                    v61 := GenS4{F0: fi47, F1: (fi47 / (((v43 << ((genFn1("foo bar", true, "hello")) & 63))) | 1)), F2: (v42 * (^(v42)))}
                                                    _ = v59
                                                    _ = v60
                                                    _ = v61
                                                }
                                                switch v43 {
                                                case 32767:
                                                    v62 := GenS9{F0: v3, F1: ("hello" + ""), F2: genFn0(1, 4294967295)}
                                                    _ = v62
                                                case 4294967295:
                                                    switch ("\n\t" + ("hello" + "hello")) {
                                                    case "foo bar":
                                                        v63 := []int{(v43 ^ fi45), fi47, (v43 & (v43 / (((^(fi45))) | 1))), (v43 >> ((v54) & 63))}
                                                        _ = v63
                                                    case "hello":
                                                        var v64 int16 = 32767
                                                        _ = v64
                                                    default:
                                                        var v65 int32 = 2147483647
                                                        _ = v65
                                                    }
                                                case -9223372036854775808:
                                                    v54 += (v54 - (v46 | genFn1("\n\t", (249 == 0), ("foo bar" + "hello"))))
                                                }
                                                if (!true) {
                                                    var v66 uint16 = uint16(v54)
                                                    fmt.Println(v42)
                                                    _ = v66
                                                }
                                            }
                                            v67 := []int{1, 32768, (fi45 % ((-80) | 1))}
                                            _ = v67
                                        }
                                    case "a":
                                        if false {
                                            var v68 uint16 = 42
                                            fmt.Println(fi45)
                                            _ = v68
                                        } else {
                                            if false {
                                                v69 := GenS4{F0: 32767, F1: (^(fi47)), F2: 42}
                                                _ = v69
                                            } else {
                                                fmt.Println(v52)
                                                v43 = 32768
                                                v51 = (0.1 - 0.0)
                                            }
                                            fmt.Println(fi47)
                                            fmt.Println(fi45, v51, fi47)
                                            fmt.Println(fi45)
                                            v70 := GenS14{F0: 1}
                                            var v71 bool = (!(!false))
                                            _ = v70
                                            _ = v71
                                        }
                                    }
                                    switch fi45 {
                                    case 9223372036854775807:
                                        var v72 uint32 = 0
                                        _ = v72
                                    case -866:
                                        v73 := GenS15{F0: -1, F1: true, F2: 1}
                                        _ = v73
                                    default:
                                        v54 = (v46 / ((genFn1("hello", true, "")) | 1))
                                    }
                                } else {
                                    switch (-(fi47)) {
                                    case 2147483647:
                                        var v74 string = (("" + "a") + "\n\t")
                                        _ = v74
                                    case 1:
                                        if (!(!((false || (0 != 1)) && false))) {
                                            v75 := []int{0}
                                            var v76 float64 = v51
                                            _ = v75
                                            _ = v76
                                        } else {
                                            v42 = (^(v42))
                                        }
                                    case -2147483648:
                                        fmt.Println(v42, v46, v54)
                                    }
                                }
                                v3 -= 30
                            }
                            _ = v54
                        }
                    } else {
                        if (("\n\t" + "") > "\n\t") {
                            v52 += (v52 & v52)
                        } else {
                            v77 := []int{v43}
                            fmt.Println(v46)
                            _ = v77
                        }
                        v78 := []int{128, (fi47 + fi45), 9223372036854775807}
                        v1[1] = (fi47 | (fi45 & (fi45 - (fi47 * v43))))
                        _ = v78
                    }
                    var v79 uint8 = 90
                    fmt.Println(v79, v46, v79)
                    if ((v42 < v42) || false) {
                        if (!((0 < (v52 + v50)) && false)) {
                            v51 = float64(-2.25)
                            v1[885] = (v43 - fi45)
                        }
                        v80 := GenS16{F0: -424, F1: (v42 | (v42 >> (((v46 - (v46 - (v46 * 4294967295)))) & 63)))}
                        var v81 float64 = v51
                        if ("" != (((("a" + "foo bar") + "") + "hello") + "a")) {
                            v82 := []int{}
                            if (true && false) {
                                var v83 int = (v43 - (v43 % ((v43) | 1)))
                                if true {
                                    v1[2147483647] = fi45
                                    var v84 int16 = -418
                                    if false {
                                        v42 -= v42
                                        if (v81 > 0.1) {
                                            fmt.Println(v2, v81, v81)
                                            v85 := GenS17{F0: (v84 * (v84 & (v84 % ((989) | 1))))}
                                            v52 += (-(v50))
                                            _ = v85
                                        }
                                    } else {
                                        v84 = (-(v84))
                                        var v86 uint64 = 929
                                        _ = v86
                                    }
                                    _ = v84
                                }
                                if ((genFn0(4294967295, 2147483648) == (v79 >> ((v46) & 63))) || (-1 == v42)) {
                                    var v87 int64 = (v50 - 9223372036854775807)
                                    if (v2 <= (v2 << (((v46 * (v46 + 0))) & 63))) {
                                        var v88 uint8 = genFn0(788, 42)
                                        fmt.Println(v42, fi45, v83)
                                        _ = v88
                                    }
                                    v89 := GenS18{F0: (v42 * (v42 & (v42 & v42))), F1: 447, F2: 65535}
                                    _ = v87
                                    _ = v89
                                } else {
                                    if false {
                                        switch v43 {
                                        case -9223372036854775808:
                                            fmt.Println(v52, v42)
                                        case 0:
                                            fmt.Println(v46, v43)
                                        default:
                                            fmt.Println(fi45)
                                        }
                                        v1[4294967295] = v83
                                        var v90 int = fi45
                                        _ = v90
                                    } else {
                                        v91 := GenS19{F0: 42}
                                        _ = v91
                                    }
                                }
                                _ = v83
                            } else {
                                v92 := v1[740]
                                var v93 uint8 = (v79 % ((v3) | 1))
                                _ = v92
                                _ = v93
                            }
                            if ((3.14159 == 0.1) && true) {
                                v3 = v79
                                if (!false) {
                                    switch ("" + ("hello" + "\n\t")) {
                                    case "":
                                        v94 := GenS20{F0: -128, F1: 42}
                                        _ = v94
                                    case "a":
                                        v46 += uint(v79)
                                    case "hello":
                                        v52 = v52
                                    }
                                } else {
                                    v3 = (v79 ^ v3)
                                    _ = len(v1)
                                    v42 += v42
                                }
                            }
                            _ = v82
                        } else {
                            var v95 uint32 = 4294967295
                            v3 -= (v79 - uint8(v46))
                            var v96 int32 = 471
                            _ = v95
                            _ = v96
                        }
                        _ = v80
                        _ = v81
                    }
                    _ = v50
                    _ = v51
                    _ = v52
                    _ = v53
                    _ = v79
                } else {
                    v97 := GenS12{F0: (^(v46)), F1: genFn1("a", true, "hello")}
                    _ = v97
                }
                if false {
                    if (0.0 >= (1e10 / -2.25)) {
                        v98 := GenS6{F0: 116, F1: 0}
                        if false {
                            switch int(v42) {
                            case 20:
                                switch "\n\t" {
                                case "a":
                                    v99 := []int{v43, (fi47 | fi47), v43, fi47}
                                    _ = v99
                                }
                                v100 := GenS15{F0: fi45, F1: false, F2: 1000}
                                fmt.Println(v43)
                                _ = v100
                            }
                        }
                        v3 = 164
                        var v101 uint16 = (v2 - (v2 - (v2 * v2)))
                        v46 = (^(v46))
                        _ = v101
                        _ = v98
                    }
                    v102 := GenS3{F0: 32767}
                    if false {
                        if false {
                            v43 += 555
                        } else {
                            var v103 uint = 255
                            _ = v103
                        }
                    }
                    _ = v102
                }
                fmt.Println(fi45)
            }
            _ = fi47
            _ = v48
            _ = v49
        }
        v43 += v43
        _ = fi45
        _ = v46
    }
    v104 := GenS21{F0: (true || (771 > 18446744073709551615))}
    if (-9223372036854775808 != 361) {
        for fi105 := 0; fi105 < 17; fi105++ {
            var v106 float32 = 0.1
            _ = fi105
            _ = v106
        }
        v107 := GenS22{F0: 18446744073709551615, F1: 4294967295, F2: 42}
        _ = v107
    } else {
        if true {
            if false {
                v42 += 127
            }
            if (!true) {
                if ((-(v43)) < int(v2)) {
                    var v108 float32 = 0.1
                    _ = v108
                } else {
                    v43 += v43
                    var v109 float32 = (3.14159 + (1e10 - float32((float64(((float32(1.5) - 0.1) - 0.1)) / -2.25))))
                    _ = v109
                }
                var v110 int8 = (-(v42))
                _ = v110
            } else {
                for fi111 := 0; fi111 < 4; fi111++ {
                    var v112 float32 = 1.5
                    v113 := []int{}
                    v1[1] = (fi111 | 4294967295)
                    _ = fi111
                    _ = v112
                    _ = v113
                }
                v2 = (v2 - 264)
            }
            switch (^(v43)) {
            case 32768:
                v1[321] = v43
            case 42:
                for fi114 := 0; fi114 < 20; fi114++ {
                    for fi115 := 0; fi115 < 13; fi115++ {
                        if true {
                            v116 := GenS23{F0: -333}
                            v117 := GenS9{F0: v3, F1: "", F2: 128}
                            var v118 string = "hello"
                            _ = v116
                            _ = v117
                            _ = v118
                        } else {
                            if (("a" != ("" + ("hello" + "hello"))) && ("\n\t" != "")) {
                                _ = len(v1)
                                if ((v3 / ((uint8(v2)) | 1)) == (v3 & (^(v3)))) {
                                    if (float64(1.5) == -2.25) {
                                        fmt.Println(v42)
                                        v3 = uint8(v2)
                                    }
                                    v119 := GenS18{F0: 42, F1: -1, F2: 2147483648}
                                    var v120 int64 = 0
                                    _ = v119
                                    _ = v120
                                } else {
                                    if (float32(0.0) == float32(((3.14159 - (((0.0 / 3.14159) + 1e10) - 0.1)) / float64((3.14159 - 0.0))))) {
                                        v3 -= v3
                                        v121 := GenS24{F0: -155, F1: (v2 % (((^(v2))) | 1)), F2: 18446744073709551615}
                                        _ = v121
                                    }
                                    switch "foo bar" {
                                    case "\n\t":
                                        fmt.Println(v42, v2, v2)
                                        fmt.Println(fi114)
                                        v122 := GenS25{F0: 2147483647}
                                        _ = v122
                                    case "a":
                                        switch "a" {
                                        case "a":
                                            var v123 bool = true
                                            var v124 uint16 = 0
                                            v42 += 1
                                            switch "foo bar" {
                                            case "":
                                                switch int(v3) {
                                                case 127:
                                                    if (!(32767 != 0)) {
                                                        v124 += 1
                                                        v125 := GenS2{F0: "hello", F1: v42, F2: (v43 << ((628) & 63))}
                                                        fmt.Println(v43, v43, v42)
                                                        var v126 uint = 624
                                                        var v127 uint32 = uint32(v126)
                                                        _ = v125
                                                        _ = v126
                                                        _ = v127
                                                    }
                                                    switch (fi115 % ((fi114) | 1)) {
                                                    case 9223372036854775807:
                                                        v123 = (!true)
                                                    }
                                                    v43 = v43
                                                    v128 := []int{32768, fi114}
                                                    switch (("hello" + ("\n\t" + "foo bar")) + "foo bar") {
                                                    case "":
                                                        if false {
                                                            _ = len(v1)
                                                            switch "\n\t" {
                                                            case "foo bar":
                                                                v129 := v1[4294967295]
                                                                _ = v129
                                                            default:
                                                                if ((v2 * (^(v2))) != v2) {
                                                                    v2 = (v2 >> ((0) & 63))
                                                                } else {
                                                                    fmt.Println(v43, v2, fi114)
                                                                }
                                                            }
                                                            if (v123 || (-869 != 32767)) {
                                                                var v130 uint64 = 18446744073709551615
                                                                if v123 {
                                                                    v131 := GenS26{F0: 2, F1: ("" + ("\n\t" + "a")), F2: 32767}
                                                                    v130 = (v130 - v130)
                                                                    _ = v131
                                                                }
                                                                _ = v130
                                                            }
                                                            var v132 uint16 = 1
                                                            _ = v132
                                                        } else {
                                                            v133 := GenS27{F0: 286, F1: 1}
                                                            _ = v133
                                                        }
                                                    case "hello":
                                                        var v134 int32 = -452
                                                        _ = v134
                                                    case "foo bar":
                                                        v135 := GenS13{F0: v2}
                                                        _ = v135
                                                    }
                                                    if v123 {
                                                        if false {
                                                            var v136 float64 = 3.14159
                                                            v1[4294967295] = fi115
                                                            _ = v136
                                                        } else {
                                                            fmt.Println(v2, v2)
                                                            var v137 uint32 = 414
                                                            if len(v128) >= 2 {
                                                                v128 = v128[1:len(v128)]
                                                            }
                                                            _ = v137
                                                        }
                                                        fmt.Println(fi114)
                                                        var v138 bool = (fi114 == int(v42))
                                                        _ = v138
                                                    }
                                                    _ = v128
                                                }
                                            case "hello":
                                                fmt.Println(fi114)
                                                var v139 int32 = -2147483648
                                                _ = v139
                                            }
                                            _ = v123
                                            _ = v124
                                        }
                                    }
                                }
                            } else {
                                _ = len(v1)
                                v43 += -802
                            }
                        }
                        var v140 float64 = 1e10
                        _ = fi115
                        _ = v140
                    }
                    _ = fi114
                }
            }
        } else {
            v42 = v42
            v141 := GenS19{F0: 389}
            _ = v141
        }
        switch (v43 | 116) {
        case 1:
            if (1 == (v3 << ((244) & 63))) {
                v142 := []int{v43, v43, v43}
                switch ("a" + "") {
                case "foo bar":
                    for fi143 := 0; fi143 < 14; fi143++ {
                        v144 := v1[1]
                        switch 128 {
                        case 9223372036854775807:
                            var v145 int16 = 0
                            _ = v145
                        default:
                            var v146 bool = false
                            _ = v146
                        }
                        var v147 int16 = -581
                        _ = fi143
                        _ = v144
                        _ = v147
                    }
                }
                _ = v142
            } else {
                if (2 < 18446744073709551615) {
                    fmt.Println(v2, v2)
                }
                for fi148 := 0; fi148 < 4; fi148++ {
                    var v149 float32 = (1e10 / 3.14159)
                    var v150 uint32 = 1
                    v42 -= (v42 + 67)
                    _ = fi148
                    _ = v149
                    _ = v150
                }
            }
        default:
            var v151 uint8 = v3
            _ = v151
        }
    }
    fmt.Println(v2, v43, v42)
    v2 = v2
    v43 = int(v2)
    if (65535 != 65535) {
        var v152 int = (v43 / ((v43) | 1))
        v152 = v152
        _ = v152
    } else {
        v2 = 2
        var v153 uint32 = 56
        v154 := GenS15{F0: (v43 / ((v43) | 1)), F1: false, F2: 1}
        _ = v153
        _ = v154
    }
    for fi155 := 0; fi155 < 9; fi155++ {
        var v156 uint = 855
        if (935 >= (v156 * (v156 * (v156 + (v156 | (v156 - (v156 / ((503) | 1)))))))) {
            fmt.Println(v156, v43, v43)
            var v157 bool = (42 > 42)
            v2 += v2
            _ = v157
        }
        _ = fi155
        _ = v156
    }
    switch "hello" {
    case "hello":
        for fi158 := 0; fi158 < 13; fi158++ {
            if false {
                fmt.Println(v43, v42)
            }
            var v159 uint8 = v3
            var v160 float32 = 1e10
            _ = fi158
            _ = v159
            _ = v160
        }
    }
    v43 = (v43 / ((0) | 1))
    v161 := []int{(v43 | -794), 401, int(v42)}
    for fi162 := 0; fi162 < 18; fi162++ {
        var v163 uint32 = 581
        _ = fi162
        _ = v163
    }
    v3 -= (v3 << ((2) & 63))
    var v164 int32 = -2147483648
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v161:", len(v161), v161)
    fmt.Println("v164:", v164)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v42:", v42)
    fmt.Println("v43:", v43)
    _ = v1
    _ = v104
    _ = v161
    _ = v164
    _ = v2
    _ = v3
    _ = v42
    _ = v43
    _ = v44
}
