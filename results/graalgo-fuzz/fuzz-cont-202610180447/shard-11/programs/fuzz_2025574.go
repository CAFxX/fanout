package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
type GenS2 struct { F0 uint16; F1 uint8; F2 uint16 }
type GenS3 struct { F0 int32 }
type GenS4 struct { F0 bool; F1 int16; F2 string }
type GenS5 struct { F0 int }
type GenS6 struct { F0 uint16; F1 uint8; F2 int16 }
type GenS7 struct { F0 uint }
type GenS8 struct { F0 uint32; F1 int64 }
type GenS9 struct { F0 uint8; F1 uint64 }
type GenS10 struct { F0 uint32; F1 int16 }
type GenS11 struct { F0 uint32 }
type GenS12 struct { F0 uint16; F1 int16; F2 uint32 }
type GenS13 struct { F0 uint8 }
type GenS14 struct { F0 int16; F1 uint16; F2 uint16 }
func genFn0(p0 bool) int8 {
    var v5 uint = 255
    for fi6 := 0; fi6 < 16; fi6++ {
        p0 = (fi6 <= fi6)
        v7 := []int{(-(fi6)), (^(fi6))}
        _ = fi6
        _ = v7
    }
    for fi8 := 0; fi8 < 11; fi8++ {
        if (p0 || (true || (uint32(v5) == 2147483648))) {
            p0 = p0
        }
        switch ("foo bar" + "hello") {
        case "\n\t":
            v9 := GenS1{F0: int8(v5)}
            _ = v9
        }
        switch "foo bar" {
        case "a":
            p0 = (1 != 32)
        case "\n\t":
            fmt.Println(p0)
        case "hello":
            p0 = (9223372036854775807 == -9223372036854775808)
        default:
            var v10 int64 = 9223372036854775807
            _ = v10
        }
        _ = fi8
    }
    v11 := GenS2{F0: 42, F1: 195, F2: 1}
    switch "hello" {
    case "a":
        p0 = p0
    }
    _ = p0
    _ = v11
    _ = v5
    return 0
}

func main() {
    v1 := make(map[string]int)
    var v2 uint8 = 42
    var v3 int32 = 586
    var v4 int16 = 1
    for fi12 := 0; fi12 < 8; fi12++ {
        v4 -= (v4 ^ v4)
        if (145 == 4294967295) {
            fmt.Println(v2, v3, v3)
            v13 := GenS3{F0: 706}
            _ = v13
        } else {
            if (!("" <= "\n\t")) {
                switch (("foo bar" + ((("" + "a") + "foo bar") + ("foo bar" + "hello"))) + ("\n\t" + ("" + "foo bar"))) {
                case "a":
                    var v14 bool = (18446744073709551615 < 756)
                    _ = v14
                }
            }
            v15 := GenS1{F0: genFn0((1 == 2147483647))}
            _ = v15
        }
        _ = fi12
    }
    switch -2147483648 {
    case -9223372036854775808:
        for fi16 := 0; fi16 < 18; fi16++ {
            var v17 int16 = (v4 >> ((uint(v2)) & 63))
            if (2 > 46) {
                fmt.Println(v2, v3)
            }
            _ = fi16
            _ = v17
        }
    case 285:
        var v18 int16 = (v4 * (v4 * 1))
        _ = v18
    default:
        v3 = int32(v2)
    }
    v1[""] = 32767
    if true {
        v4 = v4
        v19 := GenS4{F0: ("a" == ((("hello" + "foo bar") + "foo bar") + ("hello" + "hello"))), F1: v4, F2: "a"}
        _ = v19
    }
    v20 := []int{-1, 127, 9223372036854775807, 9223372036854775807}
    var v21 string = "foo bar"
    v2 += v2
    for fi22 := 0; fi22 < 14; fi22++ {
        var v23 uint16 = 426
        var v24 uint32 = 347
        _ = fi22
        _ = v23
        _ = v24
    }
    _ = len(v20)
    if len(v20) >= 2 {
        v20 = v20[1:len(v20)]
    }
    if ((!true) || (!false)) {
        switch 127 {
        case -2147483648:
            if false {
                var v25 int = -9223372036854775808
                v26 := GenS2{F0: 42, F1: 2, F2: 229}
                _ = v25
                _ = v26
            }
        case 4294967295:
            if len(v20) >= 2 {
                v20 = v20[1:len(v20)]
            }
        default:
            for fi27 := 0; fi27 < 3; fi27++ {
                fmt.Println(fi27, v3, v2)
                v28 := GenS3{F0: 0}
                _ = fi27
                _ = v28
            }
        }
    }
    var v29 int16 = int16(v2)
    switch 128 {
    case -85:
        switch "hello" {
        case "\n\t":
            v3 -= (v3 << ((0) & 63))
        case "":
            v29 = v29
        }
    default:
        v20 = append(v20, 9223372036854775807)
    }
    v1[(v21 + (v21 + v21))] = 42
    switch 4294967295 {
    case 42:
        switch "a" {
        case "a":
            var v30 int64 = -129
            _ = v30
        case "foo bar":
            if (2 <= 4294967295) {
                fmt.Println(v2, v21, v4)
                v31 := GenS2{F0: 65535, F1: (v2 / (((v2 >> ((0) & 63))) | 1)), F2: 6}
                _ = v31
            } else {
                v21 = v21
                switch 128 {
                case -2147483648:
                    if ((("a" + v21) + v21) <= v21) {
                        var v32 uint16 = 65535
                        if true {
                            switch v21 {
                            case "\n\t":
                                v33 := GenS5{F0: 1}
                                _ = v33
                            case "hello":
                                switch "hello" {
                                case "\n\t":
                                    switch 1 {
                                    case 9223372036854775807:
                                        var v34 uint16 = v32
                                        _ = v34
                                    case 42:
                                        for fi35 := 0; fi35 < 8; fi35++ {
                                            v36 := GenS6{F0: 42, F1: 44, F2: (^(v4))}
                                            _ = fi35
                                            _ = v36
                                        }
                                    case 128:
                                        v37 := v1[(("a" + "a") + ("\n\t" + v21))]
                                        _ = v37
                                    default:
                                        if (!true) {
                                            var v38 float32 = (0.1 + 0.0)
                                            if (1e10 != float64((1.5 * (v38 * v38)))) {
                                                switch 127 {
                                                case 830:
                                                    for fi39 := 0; fi39 < 7; fi39++ {
                                                        v4 -= v29
                                                        v40 := GenS6{F0: 1, F1: (v2 | 127), F2: int16(v2)}
                                                        for fi41 := 0; fi41 < 8; fi41++ {
                                                            var v42 int16 = v4
                                                            _ = len(v1)
                                                            _ = len(v1)
                                                            _ = fi41
                                                            _ = v42
                                                        }
                                                        _ = fi39
                                                        _ = v40
                                                    }
                                                case -2147483648:
                                                    v43 := GenS2{F0: (v32 ^ (^(v32))), F1: v2, F2: (v32 | v32)}
                                                    _ = v43
                                                case -30:
                                                    for fi44 := 0; fi44 < 13; fi44++ {
                                                        switch "a" {
                                                        case "\n\t":
                                                            if ((v29 >> ((uint(v2)) & 63)) < v29) {
                                                                v2 -= (v2 & (v2 ^ v2))
                                                                if len(v20) > 0 {
                                                                    v20[0] = (fi44 ^ (fi44 - (fi44 + (fi44 + -1))))
                                                                }
                                                                fmt.Println(v21, v2)
                                                            } else {
                                                                v32 -= 42
                                                                var v45 int32 = (v3 * v3)
                                                                var v46 string = "foo bar"
                                                                _ = v45
                                                                _ = v46
                                                            }
                                                        case "":
                                                            delete(v1, v21)
                                                        case "foo bar":
                                                            var v47 bool = false
                                                            _ = v47
                                                        default:
                                                            v1[(v21 + v21)] = fi44
                                                        }
                                                        switch (v21 + v21) {
                                                        case "":
                                                            v48 := v1[((("\n\t" + v21) + "") + v21)]
                                                            _ = v48
                                                        default:
                                                            switch "a" {
                                                            case "\n\t":
                                                                v4 -= (v4 | v4)
                                                            default:
                                                                v49 := v1[v21]
                                                                _ = v49
                                                            }
                                                        }
                                                        _ = fi44
                                                    }
                                                default:
                                                    if false {
                                                        v38 -= 0.1
                                                        v3 += (v3 >> ((285) & 63))
                                                    } else {
                                                        var v50 bool = (1 >= 428)
                                                        v3 += v3
                                                        v51 := v1[(v21 + "a")]
                                                        _ = v50
                                                        _ = v51
                                                    }
                                                }
                                                if len(v20) > 0 {
                                                    v20[0] = 2147483647
                                                }
                                            } else {
                                                switch v21 {
                                                case "hello":
                                                    for fi52 := 0; fi52 < 19; fi52++ {
                                                        var v53 int8 = 127
                                                        if (v32 == 2) {
                                                            for fi54 := 0; fi54 < 19; fi54++ {
                                                                if true {
                                                                    fmt.Println(fi52)
                                                                    v21 = "a"
                                                                    var v55 bool = (333 > 325)
                                                                    _ = v55
                                                                } else {
                                                                    v2 = v2
                                                                    if (false || (!(true || true))) {
                                                                        v2 += v2
                                                                        if (true || ((3.14159 / (v38 - float32(3.14159))) >= v38)) {
                                                                            var v56 string = v21
                                                                            _ = len(v20)
                                                                            if true {
                                                                                var v57 string = "\n\t"
                                                                                fmt.Println(fi52, v53, v57)
                                                                                _ = v57
                                                                            }
                                                                            _ = v56
                                                                        }
                                                                        fmt.Println(v2, v4, v32)
                                                                        switch (^(fi52)) {
                                                                        case -9223372036854775808:
                                                                            v58 := GenS3{F0: v3}
                                                                            v59 := GenS3{F0: v3}
                                                                            fmt.Println(v2, v4, fi54)
                                                                            _ = v58
                                                                            _ = v59
                                                                        case 127:
                                                                            v53 = (v53 ^ (v53 | v53))
                                                                        case 1:
                                                                            v2 -= (v2 >> ((877) & 63))
                                                                        default:
                                                                            v1[("a" + ("foo bar" + (v21 + ((v21 + "a") + "hello"))))] = 230
                                                                        }
                                                                        v4 += v29
                                                                    }
                                                                    var v60 int16 = (v4 % (((^(v4))) | 1))
                                                                    _ = v60
                                                                }
                                                                _ = fi54
                                                            }
                                                            for fi61 := 0; fi61 < 19; fi61++ {
                                                                if true {
                                                                    var v62 int = fi52
                                                                    switch v21 {
                                                                    case "a":
                                                                        v63 := GenS7{F0: 939}
                                                                        _ = v63
                                                                    }
                                                                    _ = v62
                                                                } else {
                                                                    v3 = v3
                                                                }
                                                                v3 -= (v3 & v3)
                                                                switch 42 {
                                                                case 0:
                                                                    v64 := GenS6{F0: (v32 | v32), F1: (^(v2)), F2: (v4 >> ((133) & 63))}
                                                                    _ = v64
                                                                default:
                                                                    if true {
                                                                        var v65 uint16 = (v32 / ((v32) | 1))
                                                                        _ = v65
                                                                    } else {
                                                                        if (true && false) {
                                                                            var v66 int = (fi52 * fi52)
                                                                            if true {
                                                                                fmt.Println(v66, v53, v32)
                                                                                _ = len(v1)
                                                                            }
                                                                            v2 = (v2 * (v2 >> ((18446744073709551615) & 63)))
                                                                            _ = v66
                                                                        }
                                                                        var v67 uint64 = 0
                                                                        v20 = append(v20, 32768)
                                                                        _ = v67
                                                                    }
                                                                }
                                                                _ = fi61
                                                            }
                                                            for fi68 := 0; fi68 < 3; fi68++ {
                                                                var v69 int16 = v29
                                                                if (!(855 <= 2)) {
                                                                    var v70 int32 = -2147483648
                                                                    _ = v70
                                                                }
                                                                switch (fi52 + (-(fi52))) {
                                                                case -561:
                                                                    v71 := GenS8{F0: 504, F1: 1}
                                                                    _ = v71
                                                                default:
                                                                    var v72 float32 = ((1e10 + 0.0) / 1e10)
                                                                    _ = v72
                                                                }
                                                                _ = fi68
                                                                _ = v69
                                                            }
                                                        } else {
                                                            for fi73 := 0; fi73 < 17; fi73++ {
                                                                v74 := GenS9{F0: uint8(v32), F1: 0}
                                                                fmt.Println(v2)
                                                                v3 -= (v3 / ((0) | 1))
                                                                var v75 uint = 1
                                                                v32 -= (v32 * 255)
                                                                fmt.Println(v29)
                                                                _ = fi73
                                                                _ = v74
                                                                _ = v75
                                                            }
                                                            v76 := GenS2{F0: 65535, F1: (^(v2)), F2: 2}
                                                            var v77 string = "\n\t"
                                                            _ = v76
                                                            _ = v77
                                                        }
                                                        _ = fi52
                                                        _ = v53
                                                    }
                                                case "foo bar":
                                                    var v78 int16 = 17
                                                    _ = v78
                                                case "":
                                                    v4 = 1000
                                                default:
                                                    v3 = (v3 << ((255) & 63))
                                                }
                                                v20 = append(v20, 4294967295)
                                                fmt.Println(v2)
                                            }
                                            v1[("foo bar" + v21)] = 127
                                            _ = v38
                                        }
                                    }
                                case "":
                                    v21 = "foo bar"
                                default:
                                    var v79 bool = false
                                    _ = v79
                                }
                            case "":
                                v80 := GenS4{F0: true, F1: (v4 >> ((0) & 63)), F2: v21}
                                _ = v80
                            default:
                                v81 := GenS9{F0: 1, F1: 37}
                                _ = v81
                            }
                        } else {
                            v4 = (^(v29))
                        }
                        var v82 int32 = v3
                        _ = v32
                        _ = v82
                    } else {
                        if ((!(!true)) || true) {
                            delete(v1, (((("" + v21) + ("hello" + v21)) + (v21 + ("a" + "\n\t"))) + ("a" + "hello")))
                        } else {
                            v83 := GenS10{F0: 0, F1: v4}
                            _ = v83
                        }
                        v3 += (-(v3))
                    }
                case 32767:
                    v84 := GenS11{F0: 2}
                    _ = v84
                }
            }
        case "\n\t":
            if true {
                for fi85 := 0; fi85 < 17; fi85++ {
                    for fi86 := 0; fi86 < 12; fi86++ {
                        v3 = -92
                        v87 := GenS8{F0: 65535, F1: 0}
                        fmt.Println(v21, v4)
                        _ = fi86
                        _ = v87
                    }
                    switch fi85 {
                    case 4294967295:
                        v2 += 2
                    case -2147483648:
                        _ = len(v1)
                    case 235:
                        if false {
                            for fi88 := 0; fi88 < 5; fi88++ {
                                if true {
                                    v3 += (v3 | (v3 + v3))
                                    if ((v3 / (((^(v3))) | 1)) != -1) {
                                        v1["\n\t"] = 42
                                        _ = len(v20)
                                        if true {
                                            var v89 uint = 255
                                            fmt.Println(v3, v3, fi85)
                                            fmt.Println(v21, v89)
                                            _ = v89
                                        }
                                        v90 := GenS4{F0: true, F1: v4, F2: (v21 + v21)}
                                        v29 += v4
                                        _ = v90
                                    } else {
                                        var v91 uint64 = 42
                                        if true {
                                            v91 -= 58
                                        } else {
                                            if false {
                                                var v92 int = (fi85 / ((int(v3)) | 1))
                                                _ = v92
                                            } else {
                                                v93 := GenS12{F0: 65535, F1: (v4 >> ((65535) & 63)), F2: 0}
                                                _ = v93
                                            }
                                            var v94 int = -2147483648
                                            var v95 int16 = int16(fi85)
                                            v96 := GenS2{F0: 42, F1: (v2 & 21), F2: 42}
                                            _ = v94
                                            _ = v95
                                            _ = v96
                                        }
                                        _ = v91
                                    }
                                } else {
                                    v97 := GenS4{F0: false, F1: (v4 << ((18446744073709551615) & 63)), F2: ""}
                                    v29 = v4
                                    _ = v97
                                }
                                v98 := GenS3{F0: v3}
                                _ = fi88
                                _ = v98
                            }
                            v99 := GenS13{F0: v2}
                            v100 := GenS11{F0: 349}
                            _ = v100
                            _ = v99
                        }
                    }
                    _ = fi85
                }
                v101 := GenS14{F0: (v29 >> ((4294967295) & 63)), F1: 0, F2: 884}
                _ = v101
            }
        default:
            for fi102 := 0; fi102 < 2; fi102++ {
                var v103 uint = 255
                v104 := GenS12{F0: 779, F1: v29, F2: 4294967295}
                _ = fi102
                _ = v103
                _ = v104
            }
        }
    case 672:
        var v105 int64 = 0
        _ = v105
    default:
        for fi106 := 0; fi106 < 12; fi106++ {
            switch "a" {
            case "":
                switch (fi106 << ((986) & 63)) {
                case 127:
                    v107 := v1[""]
                    _ = v107
                case 4294967295:
                    var v108 int16 = v29
                    _ = v108
                }
            case "a":
                var v109 uint8 = 135
                _ = v109
            default:
                fmt.Println(v4, v4)
            }
            for fi110 := 0; fi110 < 16; fi110++ {
                switch (v21 + ("hello" + "hello")) {
                case "\n\t":
                    v3 -= v3
                case "a":
                    if false {
                        var v111 uint = 1
                        var v112 float32 = 0.0
                        _ = v111
                        _ = v112
                    }
                case "hello":
                    var v113 int8 = genFn0(true)
                    _ = v113
                default:
                    if len(v20) >= 2 {
                        v20 = v20[1:len(v20)]
                    }
                }
                v4 += (v4 & v29)
                v114 := GenS9{F0: v2, F1: 18446744073709551615}
                _ = fi110
                _ = v114
            }
            v4 -= v29
            _ = fi106
        }
    }
    v115 := GenS4{F0: false, F1: (v29 / (((v4 ^ v29)) | 1)), F2: ((v21 + ((v21 + v21) + v21)) + v21)}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v20:", len(v20), v20)
    fmt.Println("v21:", v21)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v115
    _ = v2
    _ = v20
    _ = v21
    _ = v29
    _ = v3
    _ = v4
}
