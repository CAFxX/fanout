package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 int32 }
type GenS2 struct { F0 int8; F1 uint }
type GenS3 struct { F0 string }
type GenS4 struct { F0 int64; F1 uint32 }
type GenS5 struct { F0 int8 }
type GenS6 struct { F0 string }
type GenS7 struct { F0 int; F1 int64; F2 int32 }
type GenS8 struct { F0 int16 }
type GenS9 struct { F0 string; F1 uint }
type GenS10 struct { F0 int8; F1 uint8 }
type GenS11 struct { F0 uint64 }
type GenS12 struct { F0 uint }
type GenS13 struct { F0 uint16 }
type GenS14 struct { F0 uint16; F1 bool; F2 int8 }
type GenS15 struct { F0 int16; F1 int16; F2 int8 }
type GenS16 struct { F0 int }
type GenS17 struct { F0 uint }
type GenS18 struct { F0 int32 }
type GenS19 struct { F0 int16 }
type GenS20 struct { F0 int }
type GenS21 struct { F0 uint64 }
type GenS22 struct { F0 uint32; F1 uint8; F2 uint64 }
type GenS23 struct { F0 int8; F1 uint }
type GenS24 struct { F0 int8; F1 uint16 }
type GenS25 struct { F0 uint; F1 uint64 }
type GenS26 struct { F0 int32; F1 int32 }
type GenS27 struct { F0 string }
type GenS28 struct { F0 uint8 }
type GenS29 struct { F0 int8; F1 uint64; F2 string }
func genFn0(p0 int32, p1 int) string {
    v5 := GenS1{F0: 2, F1: p0}
    p1 = (p1 | (p1 >> ((4294967295) & 63)))
    _ = p0
    _ = p1
    _ = v5
    return ""
}
func genFn1(p0 uint, p1 int64) uint16 {
    for fi6 := 0; fi6 < 6; fi6++ {
        var v7 int = fi6
        var v8 int32 = -31
        switch 128 {
        case -9223372036854775808:
            for fi9 := 0; fi9 < 18; fi9++ {
                switch (fi9 - -159) {
                case 128:
                    var v10 bool = (2147483647 < (^(v8)))
                    _ = v10
                case -1:
                    var v11 int = (-(v7))
                    _ = v11
                case 1:
                    var v12 string = ""
                    _ = v12
                default:
                    switch "a" {
                    case "a":
                        var v13 bool = (p1 < 469)
                        _ = v13
                    case "":
                        v14 := []int{-1}
                        _ = v14
                    }
                }
                v8 += v8
                var v15 uint32 = 185
                v16 := []int{(fi9 >> ((4294967295) & 63)), fi6, fi9, (fi6 & 4294967295)}
                _ = fi9
                _ = v15
                _ = v16
            }
        case -337:
            v17 := GenS2{F0: -28, F1: (p0 - (^(p0)))}
            _ = v17
        }
        _ = fi6
        _ = v7
        _ = v8
    }
    p1 += (p1 + (p1 ^ -180))
    _ = p0
    _ = p1
    return 1
}

func main() {
    v1 := []string{"hello", ("" + "hello")}
    var v2 uint64 = 42
    v3 := []int{-446, 9223372036854775807}
    var v4 uint16 = 255
    if (true && ("foo bar" < ("hello" + "foo bar"))) {
        var v18 uint32 = 450
        if len(v1) >= 2 {
            v1 = v1[1:len(v1)]
        }
        _ = v18
    }
    var v19 float64 = ((-2.25 / ((1.5 + 1e10) - 1e10)) * -2.25)
    var v20 uint64 = v2
    fmt.Println(v4, v4, v19)
    for fi21 := 0; fi21 < 18; fi21++ {
        fmt.Println(v2, v4, fi21)
        _ = len(v1)
        var v22 float32 = 3.14159
        _ = fi21
        _ = v22
    }
    v19 = v19
    for fi23 := 0; fi23 < 7; fi23++ {
        switch "foo bar" {
        case "\n\t":
            if (v4 == v4) {
                for fi24 := 0; fi24 < 16; fi24++ {
                    if (1e10 != (v19 * 1.5)) {
                        if true {
                            if true {
                                var v25 bool = (576 > 1)
                                if (fi23 != 2147483647) {
                                    if (42 == 378) {
                                        _ = len(v3)
                                        var v26 int8 = 100
                                        if (((("" + "foo bar") == "\n\t") && (v2 <= (^(v20)))) && v25) {
                                            var v27 int16 = -32768
                                            var v28 uint8 = 128
                                            _ = v27
                                            _ = v28
                                        } else {
                                            v29 := GenS2{F0: (v26 ^ 38), F1: 18446744073709551615}
                                            _ = v29
                                        }
                                        _ = v26
                                    } else {
                                        fmt.Println(v19, v2)
                                        fmt.Println(v20, v4, v19)
                                        if len(v1) > 0 {
                                            v1[0] = ""
                                        }
                                    }
                                    var v30 bool = (v2 == 456)
                                    if false {
                                        var v31 string = "hello"
                                        switch fi23 {
                                        case 127:
                                            var v32 string = ("" + ((((v31 + "foo bar") + "a") + (v31 + "")) + v31))
                                            _ = v32
                                        case -9223372036854775808:
                                            if false {
                                                if v30 {
                                                    switch ("" + v31) {
                                                    case "a":
                                                        var v33 float32 = 0.0
                                                        _ = v33
                                                    case "":
                                                        _ = len(v3)
                                                    default:
                                                        switch (fi23 / ((fi23) | 1)) {
                                                        case 32768:
                                                            v34 := GenS2{F0: -128, F1: 65535}
                                                            fmt.Println(v25)
                                                            v1 = append(v1, "\n\t")
                                                            _ = v34
                                                        case 1:
                                                            v31 = (v31 + (v31 + (v31 + "")))
                                                        case 32767:
                                                            v35 := GenS2{F0: -128, F1: 140}
                                                            _ = v35
                                                        }
                                                    }
                                                    fmt.Println(v25, v25)
                                                    var v36 float32 = 1.5
                                                    _ = v36
                                                }
                                            }
                                        case 32768:
                                            switch (fi23 * -687) {
                                            case 9223372036854775807:
                                                var v37 int8 = 72
                                                _ = v37
                                            }
                                            v25 = (((0.1 / (v19 + 0.0)) * (0.1 + -2.25)) <= v19)
                                            var v38 uint16 = (v4 ^ v4)
                                            _ = v38
                                        default:
                                            v20 = (v20 | (v2 << ((18446744073709551615) & 63)))
                                        }
                                        var v39 float64 = v19
                                        _ = v31
                                        _ = v39
                                    } else {
                                        var v40 uint64 = (^(v2))
                                        if false {
                                            fmt.Println(v2)
                                            v20 = v20
                                            var v41 int8 = -1
                                            _ = v41
                                        }
                                        _ = len(v1)
                                        _ = v40
                                    }
                                    _ = v30
                                } else {
                                    if v25 {
                                        v42 := GenS2{F0: 1, F1: 0}
                                        _ = v42
                                    } else {
                                        v43 := GenS1{F0: v2, F1: 2147483647}
                                        switch ("foo bar" + "hello") {
                                        case "foo bar":
                                            v44 := GenS2{F0: int8(fi23), F1: 18446744073709551615}
                                            _ = v44
                                        default:
                                            if (!v25) {
                                                _ = len(v1)
                                            } else {
                                                var v45 uint16 = genFn1(42, 9223372036854775807)
                                                var v46 int8 = -128
                                                _ = v45
                                                _ = v46
                                            }
                                        }
                                        v19 -= (((float64((1.5 - ((1.5 / 0.1) - -2.25))) + float64(1e10)) * (-2.25 * (v19 + 0.1))) + float64(float32(3.14159)))
                                        var v47 uint8 = 1
                                        if ((float64((3.14159 / (float32(v19) - float32(v19)))) + -2.25) == v19) {
                                            switch "hello" {
                                            case "foo bar":
                                                var v48 int8 = -14
                                                _ = v48
                                            default:
                                                v20 -= uint64(v4)
                                            }
                                        }
                                        _ = v43
                                        _ = v47
                                    }
                                    if len(v1) >= 2 {
                                        v1 = v1[1:len(v1)]
                                    }
                                }
                                switch "foo bar" {
                                case "a":
                                    v20 = v20
                                case "foo bar":
                                    var v49 uint = 4294967295
                                    var v50 float32 = 1e10
                                    _ = v49
                                    _ = v50
                                case "hello":
                                    if v25 {
                                        v4 = v4
                                        switch ("foo bar" + "") {
                                        case "hello":
                                            var v51 int64 = -1
                                            _ = v51
                                        case "a":
                                            v25 = (424 > 32767)
                                        case "":
                                            if v25 {
                                                v52 := GenS3{F0: ("\n\t" + "foo bar")}
                                                _ = v52
                                            }
                                        default:
                                            switch (fi23 >> ((385) & 63)) {
                                            case 1:
                                                var v53 uint64 = v20
                                                _ = v53
                                            default:
                                                if (!true) {
                                                    var v54 float32 = 3.14159
                                                    _ = v54
                                                } else {
                                                    var v55 uint32 = 0
                                                    if true {
                                                        var v56 float64 = v19
                                                        if (v25 && ("a" > ("foo bar" + "a"))) {
                                                            fmt.Println(v4, v4)
                                                        }
                                                        _ = len(v1)
                                                        if len(v3) > 0 {
                                                            v3[0] = fi24
                                                        }
                                                        _ = v56
                                                    } else {
                                                        switch fi23 {
                                                        case -1:
                                                            v57 := GenS4{F0: -892, F1: 945}
                                                            _ = v57
                                                        default:
                                                            v58 := GenS3{F0: "a"}
                                                            _ = v58
                                                        }
                                                        switch fi23 {
                                                        case 127:
                                                            if ((v55 - (v55 * (v55 >> ((565) & 63)))) > (^(v55))) {
                                                                var v59 string = ("" + "hello")
                                                                var v60 uint64 = (v2 << ((433) & 63))
                                                                _ = v59
                                                                _ = v60
                                                            } else {
                                                                if (v25 && (v2 != v2)) {
                                                                    if ((v2 ^ v20) >= 0) {
                                                                        var v61 uint16 = uint16(v55)
                                                                        _ = v61
                                                                    } else {
                                                                        if (!v25) {
                                                                            if v25 {
                                                                                v2 = 4294967295
                                                                                if (v25 && false) {
                                                                                    v62 := GenS1{F0: (v2 * (^(v20))), F1: int32(v20)}
                                                                                    _ = v62
                                                                                }
                                                                                v55 = uint32(v4)
                                                                                switch "\n\t" {
                                                                                case "":
                                                                                    v63 := GenS3{F0: "foo bar"}
                                                                                    _ = v63
                                                                                default:
                                                                                    fmt.Println(fi23)
                                                                                }
                                                                            } else {
                                                                                switch fi24 {
                                                                                case 0:
                                                                                    var v64 string = "foo bar"
                                                                                    _ = v64
                                                                                }
                                                                                fmt.Println(v2, v2, v55)
                                                                                if len(v3) > 0 {
                                                                                    v3[0] = (fi23 << ((1) & 63))
                                                                                }
                                                                            }
                                                                            v55 = v55
                                                                            var v65 int8 = 127
                                                                            _ = v65
                                                                        }
                                                                    }
                                                                    if v25 {
                                                                        v19 = v19
                                                                        v66 := GenS5{F0: 0}
                                                                        var v67 bool = true
                                                                        fmt.Println(v20, v25, v19)
                                                                        if len(v3) >= 2 {
                                                                            v3 = v3[1:len(v3)]
                                                                        }
                                                                        _ = v66
                                                                        _ = v67
                                                                    }
                                                                }
                                                                v4 -= (v4 | v4)
                                                                _ = len(v3)
                                                            }
                                                        case -1:
                                                            v55 -= (v55 & (v55 - (^(v55))))
                                                        case 9223372036854775807:
                                                            if len(v1) > 0 {
                                                                v1[0] = ""
                                                            }
                                                        }
                                                        if false {
                                                            if false {
                                                                var v68 uint8 = 181
                                                                fmt.Println(v55, v68, v2)
                                                                if ((-(fi23)) <= (fi24 * (fi23 / ((fi24) | 1)))) {
                                                                    fmt.Println(v20, v4)
                                                                    var v69 float32 = 3.14159
                                                                    var v70 float32 = v69
                                                                    _ = v69
                                                                    _ = v70
                                                                }
                                                                v71 := GenS6{F0: ((("hello" + (((("" + "hello") + "hello") + ("foo bar" + "foo bar")) + "")) + "\n\t") + "foo bar")}
                                                                _ = v68
                                                                _ = v71
                                                            } else {
                                                                switch (fi24 - fi23) {
                                                                case 1:
                                                                    var v72 string = ("a" + "")
                                                                    _ = v72
                                                                case 128:
                                                                    var v73 string = (("foo bar" + "foo bar") + "hello")
                                                                    _ = v73
                                                                case 385:
                                                                    v25 = true
                                                                default:
                                                                    switch (fi24 ^ -2147483648) {
                                                                    case -2147483648:
                                                                        var v74 uint = 1
                                                                        _ = v74
                                                                    case -80:
                                                                        v55 = (^(v55))
                                                                    case 1:
                                                                        if v25 {
                                                                            switch 9223372036854775807 {
                                                                            case -2147483648:
                                                                                fmt.Println(v19, v2, fi23)
                                                                            default:
                                                                                if (false || (v20 == (v20 / (((v2 + (^(v2)))) | 1)))) {
                                                                                    v55 = 408
                                                                                    _ = len(v1)
                                                                                    v20 = uint64(v55)
                                                                                }
                                                                                v2 -= (^(v2))
                                                                            }
                                                                            var v75 int8 = 1
                                                                            v25 = v25
                                                                            _ = v75
                                                                        }
                                                                    }
                                                                }
                                                            }
                                                            v76 := GenS4{F0: -9223372036854775808, F1: 177}
                                                            _ = v76
                                                        }
                                                        v2 = (v20 + v20)
                                                    }
                                                    v4 = v4
                                                    _ = v55
                                                }
                                            }
                                        }
                                        v77 := GenS6{F0: "hello"}
                                        _ = v77
                                    }
                                }
                                _ = v25
                            }
                            v78 := GenS7{F0: (fi23 ^ fi24), F1: 9223372036854775807, F2: 960}
                            var v79 bool = ("hello" > ((("hello" + ("hello" + ((("" + (("\n\t" + "foo bar") + ((((((("foo bar" + "\n\t") + "a") + ((("foo bar" + "foo bar") + ("\n\t" + (("hello" + "hello") + ""))) + "foo bar")) + "\n\t") + ("a" + "hello")) + "\n\t") + ("hello" + ("\n\t" + (((("\n\t" + "hello") + "\n\t") + "foo bar") + ("a" + ("\n\t" + "hello")))))))) + ("a" + "hello")) + "\n\t"))) + ("a" + ("" + "hello"))) + ""))
                            _ = v78
                            _ = v79
                        } else {
                            fmt.Println(fi23)
                            switch (("hello" + (((((("hello" + (("a" + "a") + "a")) + "a") + "hello") + "") + "\n\t") + "foo bar")) + ("hello" + "\n\t")) {
                            case "a":
                                if false {
                                    v80 := GenS8{F0: 1}
                                    var v81 string = "foo bar"
                                    v2 += v2
                                    _ = v80
                                    _ = v81
                                } else {
                                    var v82 uint8 = 186
                                    var v83 int64 = -1
                                    if false {
                                        v84 := GenS9{F0: "", F1: 18446744073709551615}
                                        _ = v84
                                    } else {
                                        var v85 int = (fi24 * 800)
                                        fmt.Println(fi24, fi23, v19)
                                        _ = v85
                                    }
                                    _ = v82
                                    _ = v83
                                }
                            }
                        }
                        var v86 string = ("foo bar" + "")
                        if (genFn1(255, -9223372036854775808) < (v4 + (v4 % ((v4) | 1)))) {
                            if len(v1) >= 2 {
                                v1 = v1[1:len(v1)]
                            }
                            v20 -= 4294967295
                            var v87 bool = (127 != -25)
                            _ = v87
                        } else {
                            if false {
                                v20 = (v20 & (v2 - v2))
                                v4 += (v4 / ((v4) | 1))
                                if ((v4 >> ((4294967295) & 63)) >= 255) {
                                    if false {
                                        switch (v86 + v86) {
                                        case "\n\t":
                                            switch (fi23 % (((fi24 >> ((42) & 63))) | 1)) {
                                            case 32768:
                                                v2 -= (v20 ^ v2)
                                            case 282:
                                                if (!(42 < 2147483647)) {
                                                    v19 -= v19
                                                }
                                            case -339:
                                                var v88 int32 = 42
                                                _ = v88
                                            }
                                        case "hello":
                                            var v89 int = 319
                                            _ = v89
                                        }
                                    } else {
                                        var v90 uint = 65535
                                        switch ((v86 + v86) + "\n\t") {
                                        case "foo bar":
                                            if (!false) {
                                                v90 = v90
                                                v91 := GenS10{F0: 100, F1: 2}
                                                switch int(v4) {
                                                case 2147483647:
                                                    v20 = (v20 % ((434) | 1))
                                                case 0:
                                                    fmt.Println(v20)
                                                default:
                                                    v19 -= ((((v19 + 1.5) / ((v19 + v19) + 1.5)) + (v19 / v19)) - float64(3.14159))
                                                    v92 := GenS6{F0: v86}
                                                    var v93 float64 = (v19 * (v19 / 1.0))
                                                    _ = v92
                                                    _ = v93
                                                }
                                                _ = v91
                                            } else {
                                                if len(v3) > 0 {
                                                    v3[0] = fi24
                                                }
                                            }
                                            var v94 uint = (v90 * uint(v2))
                                            v4 = v4
                                            _ = v94
                                        }
                                        _ = v90
                                    }
                                    v95 := GenS11{F0: v2}
                                    v1 = append(v1, "hello")
                                    _ = v95
                                } else {
                                    var v96 float64 = (0.1 + float64(1.5))
                                    _ = v96
                                }
                            }
                            var v97 uint32 = 999
                            _ = v97
                        }
                        _ = v86
                    }
                    switch -775 {
                    case 4294967295:
                        if false {
                            if true {
                                if len(v1) > 0 {
                                    v1[0] = "hello"
                                }
                                var v98 string = ("a" + ("" + "\n\t"))
                                switch -1 {
                                case 42:
                                    v20 = (v20 | v2)
                                case 4294967295:
                                    switch 0 {
                                    case 111:
                                        fmt.Println(v19, v2)
                                    case 2147483647:
                                        fmt.Println(v19)
                                    case 32767:
                                        if false {
                                            var v99 uint64 = (v20 - 30)
                                            fmt.Println(v98)
                                            _ = v99
                                        }
                                    default:
                                        v20 = (^(v20))
                                    }
                                case 32768:
                                    switch 2147483647 {
                                    case 494:
                                        v100 := GenS12{F0: 4294967295}
                                        _ = v100
                                    case 32767:
                                        var v101 float64 = v19
                                        _ = v101
                                    case 4294967295:
                                        var v102 bool = (true && ((true || false) || true))
                                        _ = v102
                                    }
                                }
                                switch 42 {
                                case -2147483648:
                                    fmt.Println(v20, v98)
                                case -1:
                                    var v103 uint32 = 65535
                                    _ = v103
                                default:
                                    switch 618 {
                                    case -296:
                                        switch 1 {
                                        case 2147483647:
                                            v4 = (v4 * v4)
                                            var v104 float64 = v19
                                            _ = v104
                                        }
                                    }
                                }
                                if false {
                                    var v105 uint32 = 65535
                                    v106 := GenS13{F0: v4}
                                    _ = v105
                                    _ = v106
                                } else {
                                    v2 += (v2 | (v20 << ((65535) & 63)))
                                }
                                _ = v98
                            }
                            if false {
                                v19 = 0.1
                                var v107 int64 = 42
                                if (710 <= fi23) {
                                    if false {
                                        var v108 bool = (!((v4 >> ((255) & 63)) > (v4 / ((v4) | 1))))
                                        v2 = (v20 / ((v20) | 1))
                                        v19 -= 1e10
                                        _ = v108
                                    }
                                    v2 += v20
                                }
                                _ = v107
                            } else {
                                v19 -= float64(3.14159)
                                var v109 string = "\n\t"
                                var v110 uint = 877
                                var v111 uint64 = (v2 >> ((uint(v2)) & 63))
                                _ = v109
                                _ = v110
                                _ = v111
                            }
                            switch "" {
                            case "hello":
                                var v112 int = (fi24 << ((1) & 63))
                                var v113 int = 778
                                _ = v112
                                _ = v113
                            case "\n\t":
                                fmt.Println(v4, fi23)
                            case "a":
                                v1 = append(v1, (("foo bar" + "\n\t") + "a"))
                            default:
                                if (0 <= 9223372036854775807) {
                                    switch ("" + ("\n\t" + (("" + "foo bar") + (("" + "\n\t") + "hello")))) {
                                    case "foo bar":
                                        v114 := GenS14{F0: genFn1(74, -1), F1: true, F2: 0}
                                        _ = v114
                                    case "a":
                                        var v115 int8 = 120
                                        _ = v115
                                    }
                                }
                            }
                        }
                    case -2147483648:
                        switch -891 {
                        case 126:
                            if true {
                                v20 -= 42
                                v20 = (v2 | v2)
                            } else {
                                if true {
                                    fmt.Println(v4, fi23)
                                    var v116 uint64 = v20
                                    _ = v116
                                } else {
                                    v19 = (v19 * 0.1)
                                    v117 := GenS15{F0: -182, F1: -309, F2: 42}
                                    if false {
                                        if true {
                                            var v118 uint8 = 255
                                            v119 := GenS16{F0: (fi24 / (((fi23 % (((fi24 ^ fi23)) | 1))) | 1))}
                                            v19 -= ((float64(-2.25) - v19) * -2.25)
                                            if false {
                                                v120 := GenS17{F0: 4294967295}
                                                var v121 int = fi23
                                                _ = v120
                                                _ = v121
                                            } else {
                                                switch int(v20) {
                                                case 1:
                                                    _ = len(v3)
                                                case 0:
                                                    fmt.Println(v4)
                                                }
                                            }
                                            var v122 uint64 = (^(v20))
                                            _ = v118
                                            _ = v119
                                            _ = v122
                                        } else {
                                            var v123 int = 0
                                            _ = v123
                                        }
                                        v4 = 283
                                        fmt.Println(v20, v20)
                                    } else {
                                        v124 := GenS3{F0: "foo bar"}
                                        fmt.Println(v2, v2, fi24)
                                        v19 -= v19
                                        if len(v1) > 0 {
                                            v1[0] = "hello"
                                        }
                                        _ = v124
                                    }
                                    var v125 uint32 = 2147483648
                                    _ = v117
                                    _ = v125
                                }
                            }
                        case 764:
                            if false {
                                v2 = (v20 << ((2) & 63))
                                _ = len(v1)
                            }
                        case 32768:
                            if (-32768 != 971) {
                                switch "a" {
                                case "foo bar":
                                    v2 -= 935
                                }
                            }
                        default:
                            var v126 string = "a"
                            _ = v126
                        }
                    default:
                        if false {
                            v127 := GenS18{F0: -884}
                            switch "foo bar" {
                            case "foo bar":
                                if false {
                                    var v128 string = ""
                                    _ = v128
                                } else {
                                    var v129 int32 = 2147483647
                                    switch "a" {
                                    case "":
                                        switch 42 {
                                        case -444:
                                            if true {
                                                var v130 float64 = v19
                                                switch (fi23 >> ((438) & 63)) {
                                                case -1:
                                                    var v131 uint16 = 2
                                                    _ = v131
                                                case 32767:
                                                    var v132 float64 = v19
                                                    _ = v132
                                                }
                                                _ = v130
                                            }
                                        }
                                    }
                                    switch "a" {
                                    case "foo bar":
                                        if len(v1) > 0 {
                                            v1[0] = (("hello" + "hello") + "\n\t")
                                        }
                                    }
                                    _ = v129
                                }
                            case "":
                                var v133 uint64 = (v20 + (v2 | v20))
                                _ = v133
                            }
                            v134 := GenS8{F0: -91}
                            _ = v127
                            _ = v134
                        }
                    }
                    fmt.Println(v20, v19, v4)
                    _ = fi24
                }
                v2 += (^(v20))
                v3 = append(v3, (fi23 & fi23))
            } else {
                fmt.Println(fi23)
                if true {
                    for fi135 := 0; fi135 < 15; fi135++ {
                        switch (fi135 ^ (^(fi23))) {
                        case 42:
                            v136 := GenS17{F0: 65535}
                            _ = v136
                        case 9223372036854775807:
                            if len(v1) > 0 {
                                v1[0] = (("hello" + (("" + "\n\t") + "")) + ((("hello" + "foo bar") + "a") + ("hello" + "")))
                            }
                        case 2147483647:
                            v19 = (1e10 + ((v19 + v19) / v19))
                        default:
                            v137 := GenS13{F0: v4}
                            _ = v137
                        }
                        _ = fi135
                    }
                    var v138 int64 = 717
                    if (42 >= 42) {
                        var v139 float64 = v19
                        _ = v139
                    }
                    _ = v138
                } else {
                    switch "a" {
                    case "":
                        if (!(-1 > 1)) {
                            for fi140 := 0; fi140 < 6; fi140++ {
                                v141 := GenS19{F0: 301}
                                v20 += (v2 * 2)
                                switch (fi23 + fi140) {
                                case 30:
                                    if true {
                                        var v142 int32 = 50
                                        v20 -= v20
                                        _ = v142
                                    } else {
                                        if ((!true) || true) {
                                            v143 := GenS16{F0: fi140}
                                            var v144 string = "a"
                                            _ = v143
                                            _ = v144
                                        }
                                        var v145 uint8 = 1
                                        if ((v4 == v4) || ((1 != 111) || true)) {
                                            v4 += genFn1(93, 63)
                                        } else {
                                            v145 -= uint8(v20)
                                            v19 = float64((1.5 + 1.5))
                                        }
                                        if (false || true) {
                                            var v146 uint16 = (v4 ^ 42)
                                            v147 := GenS8{F0: 1}
                                            v148 := GenS9{F0: "hello", F1: 0}
                                            switch "a" {
                                            case "a":
                                                if true {
                                                    fmt.Println(v145)
                                                    switch (fi23 + fi140) {
                                                    case 128:
                                                        if false {
                                                            var v149 float64 = ((v19 - -2.25) + v19)
                                                            _ = v149
                                                        }
                                                        fmt.Println(fi23, v4)
                                                    case -2147483648:
                                                        v150 := GenS15{F0: -757, F1: -203, F2: -128}
                                                        _ = v150
                                                    case -575:
                                                        fmt.Println(fi23, v19, v20)
                                                    default:
                                                        switch (^(fi23)) {
                                                        case 127:
                                                            v20 = (v2 & v20)
                                                        case 0:
                                                            var v151 uint64 = v20
                                                            _ = v151
                                                        default:
                                                            _ = len(v3)
                                                        }
                                                    }
                                                }
                                            case "foo bar":
                                                v19 = 0.0
                                            case "\n\t":
                                                if len(v1) > 0 {
                                                    v1[0] = "\n\t"
                                                }
                                            }
                                            v2 = (v20 / ((42) | 1))
                                            _ = v146
                                            _ = v147
                                            _ = v148
                                        }
                                        _ = v145
                                    }
                                case -705:
                                    var v152 uint16 = v4
                                    var v153 bool = (!(448 < -2147483648))
                                    switch fi140 {
                                    case 506:
                                        switch "hello" {
                                        case "\n\t":
                                            var v154 uint64 = v2
                                            _ = v154
                                        case "foo bar":
                                            if v153 {
                                                switch (fi140 | (fi23 % ((2147483647) | 1))) {
                                                case 4294967295:
                                                    var v155 int64 = -483
                                                    _ = v155
                                                case 360:
                                                    v156 := GenS4{F0: 42, F1: 4294967295}
                                                    var v157 uint16 = (v4 - genFn1(4294967295, 92))
                                                    if len(v3) >= 2 {
                                                        v3 = v3[1:len(v3)]
                                                    }
                                                    _ = v156
                                                    _ = v157
                                                case 9223372036854775807:
                                                    var v158 int = fi23
                                                    _ = v158
                                                }
                                            } else {
                                                v152 = 1
                                                switch "hello" {
                                                case "\n\t":
                                                    v2 = 18446744073709551615
                                                case "foo bar":
                                                    v1 = append(v1, "hello")
                                                case "hello":
                                                    if len(v3) >= 2 {
                                                        v3 = v3[1:len(v3)]
                                                    }
                                                default:
                                                    var v159 uint = 65535
                                                    _ = v159
                                                }
                                                if len(v3) >= 2 {
                                                    v3 = v3[1:len(v3)]
                                                }
                                            }
                                        case "hello":
                                            if false {
                                                v19 -= (float64(-2.25) + ((v19 - (1e10 - v19)) * v19))
                                                var v160 bool = (564 < -166)
                                                v161 := GenS20{F0: (^(fi23))}
                                                _ = v160
                                                _ = v161
                                            }
                                        }
                                    case 392:
                                        v152 += (v4 | v4)
                                        var v162 uint32 = 599
                                        v163 := GenS5{F0: -128}
                                        var v164 uint32 = 42
                                        if (!v153) {
                                            if v153 {
                                                v165 := GenS21{F0: 0}
                                                _ = v165
                                            } else {
                                                fmt.Println(v153)
                                                if false {
                                                    v166 := GenS12{F0: 657}
                                                    fmt.Println(v2, v164, v19)
                                                    v164 -= (v162 >> ((uint(v4)) & 63))
                                                    if len(v1) >= 2 {
                                                        v1 = v1[1:len(v1)]
                                                    }
                                                    _ = v166
                                                }
                                            }
                                        } else {
                                            if (!(396 < 350)) {
                                                fmt.Println(fi23)
                                            } else {
                                                var v167 float64 = v19
                                                v152 = (^(v4))
                                                _ = v167
                                            }
                                        }
                                        var v168 int8 = 127
                                        switch 32767 {
                                        case 9223372036854775807:
                                            if v153 {
                                                var v169 float32 = 3.14159
                                                _ = v169
                                            } else {
                                                if len(v3) > 0 {
                                                    v3[0] = fi140
                                                }
                                            }
                                        case -943:
                                            fmt.Println(fi23, v2, v162)
                                        default:
                                            v170 := GenS22{F0: (v162 - (v164 / (((v164 - v162)) | 1))), F1: 2, F2: v2}
                                            v171 := GenS14{F0: v4, F1: v153, F2: v168}
                                            var v172 int = (fi140 % ((fi140) | 1))
                                            _ = v170
                                            _ = v171
                                            _ = v172
                                        }
                                        if v153 {
                                            var v173 int = (fi140 / ((fi140) | 1))
                                            var v174 int16 = int16(v2)
                                            _ = v173
                                            _ = v174
                                        } else {
                                            if true {
                                                if v153 {
                                                    v175 := GenS21{F0: (v2 * (v2 % ((uint64(v162)) | 1)))}
                                                    _ = v175
                                                }
                                                v4 += 42
                                                v176 := GenS9{F0: "foo bar", F1: 190}
                                                _ = v176
                                            }
                                        }
                                        if (2147483647 <= -2147483648) {
                                            v168 += v168
                                            if ((fi23 / (((fi23 | (^(fi23)))) | 1)) < int(v152)) {
                                                var v177 float32 = 0.0
                                                v178 := GenS6{F0: ("foo bar" + (("hello" + "\n\t") + ""))}
                                                if v153 {
                                                    v164 = 2147483648
                                                    var v179 uint = 740
                                                    v180 := GenS23{F0: (v168 | v168), F1: v179}
                                                    _ = v179
                                                    _ = v180
                                                }
                                                _ = v177
                                                _ = v178
                                            } else {
                                                v19 -= (v19 + 3.14159)
                                                switch (fi140 ^ fi23) {
                                                case 1:
                                                    fmt.Println(v153)
                                                case -216:
                                                    switch "\n\t" {
                                                    case "foo bar":
                                                        v181 := GenS16{F0: -2147483648}
                                                        _ = v181
                                                    case "":
                                                        if (42 == v168) {
                                                            _ = len(v1)
                                                            v152 = (v152 % ((304) | 1))
                                                        } else {
                                                            v4 = v4
                                                            if (true && false) {
                                                                switch "\n\t" {
                                                                case "foo bar":
                                                                    v182 := GenS9{F0: "", F1: 565}
                                                                    _ = v182
                                                                }
                                                                if v153 {
                                                                    var v183 string = "hello"
                                                                    v20 += v2
                                                                    if len(v1) >= 2 {
                                                                        v1 = v1[1:len(v1)]
                                                                    }
                                                                    _ = v183
                                                                }
                                                            }
                                                            v153 = false
                                                        }
                                                    default:
                                                        v153 = false
                                                    }
                                                }
                                            }
                                        } else {
                                            _ = len(v3)
                                        }
                                        v152 -= (v4 + (^(v152)))
                                        _ = v162
                                        _ = v163
                                        _ = v164
                                        _ = v168
                                    default:
                                        var v184 uint8 = 136
                                        fmt.Println(v20, v4, fi140)
                                        _ = v184
                                    }
                                    _ = v152
                                    _ = v153
                                case 2147483647:
                                    fmt.Println(v19)
                                    v2 = (v20 * (v20 << ((42) & 63)))
                                default:
                                    switch fi140 {
                                    case 42:
                                        v19 = ((v19 + (-2.25 + v19)) + 1e10)
                                    case -1:
                                        v185 := GenS5{F0: 100}
                                        _ = v185
                                    }
                                }
                                v186 := GenS23{F0: -9, F1: uint(v20)}
                                if true {
                                    switch (fi23 & 42) {
                                    case 128:
                                        v19 = v19
                                    case 2147483647:
                                        v2 -= v20
                                    case 4294967295:
                                        switch (fi23 ^ (fi23 << ((42) & 63))) {
                                        case 2147483647:
                                            fmt.Println(fi23, fi23)
                                            switch (("hello" + "\n\t") + "\n\t") {
                                            case "hello":
                                                if (!true) {
                                                    v4 -= (v4 + 517)
                                                    if false {
                                                        switch (fi23 << ((255) & 63)) {
                                                        case 128:
                                                            var v187 int8 = 0
                                                            _ = v187
                                                        case -1:
                                                            if true {
                                                                var v188 int64 = 1
                                                                switch ("hello" + "\n\t") {
                                                                case "\n\t":
                                                                    var v189 int = fi140
                                                                    _ = v189
                                                                default:
                                                                    var v190 uint32 = 65535
                                                                    _ = v190
                                                                }
                                                                v2 = 556
                                                                var v191 uint16 = (v4 % (((^(v4))) | 1))
                                                                if len(v1) >= 2 {
                                                                    v1 = v1[1:len(v1)]
                                                                }
                                                                if true {
                                                                    v19 += ((v19 + v19) / 1.5)
                                                                } else {
                                                                    switch fi140 {
                                                                    case 431:
                                                                        var v192 uint = 255
                                                                        _ = v192
                                                                    case 2147483647:
                                                                        v193 := GenS21{F0: (^(v2))}
                                                                        _ = v193
                                                                    case 9223372036854775807:
                                                                        v194 := GenS24{F0: 45, F1: 980}
                                                                        _ = v194
                                                                    }
                                                                    if len(v1) > 0 {
                                                                        v1[0] = "a"
                                                                    }
                                                                    if false {
                                                                        var v195 bool = (((((float32(1e10) / ((0.1 * float32(1e10)) - (1e10 / (3.14159 - 1.5)))) * 3.14159) + float32(float64(1e10))) >= float32(0.0)) || true)
                                                                        _ = v195
                                                                    }
                                                                    var v196 int32 = -2147483648
                                                                    _ = v196
                                                                }
                                                                _ = v188
                                                                _ = v191
                                                            }
                                                        default:
                                                            v2 += (v20 - (v2 - v2))
                                                        }
                                                    } else {
                                                        switch ("" + ("\n\t" + (("foo bar" + "foo bar") + "\n\t"))) {
                                                        case "hello":
                                                            switch "foo bar" {
                                                            case "a":
                                                                v197 := GenS22{F0: 0, F1: 127, F2: v2}
                                                                _ = v197
                                                            case "":
                                                                var v198 int = (fi140 / ((fi23) | 1))
                                                                _ = v198
                                                            default:
                                                                if len(v1) >= 2 {
                                                                    v1 = v1[1:len(v1)]
                                                                }
                                                            }
                                                        case "":
                                                            var v199 uint32 = 389
                                                            _ = v199
                                                        case "foo bar":
                                                            v200 := GenS25{F0: 639, F1: (v2 | 4294967295)}
                                                            _ = v200
                                                        }
                                                    }
                                                    v20 += v2
                                                } else {
                                                    v2 += v2
                                                }
                                            default:
                                                v4 = (v4 & (v4 - 588))
                                            }
                                        case 4294967295:
                                            switch ("" + "") {
                                            case "hello":
                                                if false {
                                                    var v201 uint8 = 42
                                                    v4 = 406
                                                    var v202 int64 = 42
                                                    fmt.Println(v20, v19, v19)
                                                    _ = v201
                                                    _ = v202
                                                } else {
                                                    var v203 string = "\n\t"
                                                    _ = v203
                                                }
                                            case "foo bar":
                                                v19 = ((v19 + ((-2.25 + (3.14159 - ((v19 / 1.5) / v19))) / 1.5)) - (v19 * -2.25))
                                                v204 := GenS26{F0: 2147483647, F1: 42}
                                                _ = v204
                                            }
                                        default:
                                            if (3.14159 != 0.0) {
                                                if len(v1) >= 2 {
                                                    v1 = v1[1:len(v1)]
                                                }
                                                v2 -= 99
                                            }
                                        }
                                    }
                                    var v205 float32 = 1e10
                                    v206 := GenS8{F0: 42}
                                    _ = v205
                                    _ = v206
                                }
                                _ = fi140
                                _ = v141
                                _ = v186
                            }
                            v2 = v2
                        } else {
                            v20 += 2
                        }
                    }
                }
            }
        default:
            v1 = append(v1, ("\n\t" + ""))
        }
        if ((v19 * 3.14159) >= (0.0 - (1.5 - 0.0))) {
            for fi207 := 0; fi207 < 9; fi207++ {
                v19 = 0.0
                _ = fi207
            }
            if (4294967295 == 4294967295) {
                var v208 float32 = 0.0
                _ = v208
            } else {
                fmt.Println(fi23, v19)
                v209 := GenS27{F0: "foo bar"}
                v210 := GenS22{F0: 881, F1: 203, F2: v2}
                _ = v209
                _ = v210
            }
            if len(v3) >= 2 {
                v3 = v3[1:len(v3)]
            }
        } else {
            for fi211 := 0; fi211 < 12; fi211++ {
                if ((fi23 * (fi211 - fi211)) > -1) {
                    v212 := GenS4{F0: 42, F1: 42}
                    _ = v212
                } else {
                    var v213 int16 = -1
                    v2 -= v2
                    var v214 uint32 = 42
                    _ = v213
                    _ = v214
                }
                fmt.Println(fi211, v2)
                _ = fi211
            }
        }
        fmt.Println(v20, v2)
        _ = fi23
    }
    if (true && ((false && true) && true)) {
        var v215 uint64 = 526
        v4 -= 672
        _ = v215
    }
    var v216 int8 = 127
    switch -2147483648 {
    case 393:
        v217 := GenS28{F0: 2}
        _ = v217
    default:
        for fi218 := 0; fi218 < 18; fi218++ {
            var v219 int8 = (^(v216))
            _ = fi218
            _ = v219
        }
    }
    var v220 int8 = 42
    fmt.Println(v20)
    v216 = (v220 / ((v216) | 1))
    var v221 int = 0
    v222 := GenS29{F0: (v216 << ((322) & 63)), F1: 42, F2: "hello"}
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v216:", v216)
    fmt.Println("v220:", v220)
    fmt.Println("v221:", v221)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v19
    _ = v2
    _ = v20
    _ = v216
    _ = v220
    _ = v221
    _ = v222
    _ = v3
    _ = v4
}
