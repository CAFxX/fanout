package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
type GenS2 struct { F0 string; F1 uint32; F2 bool }
type GenS3 struct { F0 uint16; F1 int8; F2 uint8 }
type GenS4 struct { F0 uint32; F1 int16; F2 uint64 }
type GenS5 struct { F0 uint }
type GenS6 struct { F0 int; F1 uint32 }
type GenS7 struct { F0 bool; F1 uint8 }
type GenS8 struct { F0 uint32 }
type GenS9 struct { F0 string }
type GenS10 struct { F0 uint8; F1 uint; F2 uint }
type GenS11 struct { F0 string; F1 int16 }
type GenS12 struct { F0 uint64; F1 int32; F2 uint }
type GenS13 struct { F0 bool; F1 int }
type GenS14 struct { F0 bool; F1 int; F2 int16 }
type GenS15 struct { F0 uint16 }
func genFn0(p0 uint8) int32 {
    p0 += (p0 >> ((4294967295) & 63))
    v5 := GenS1{F0: -1}
    var v6 int = 42
    _ = p0
    _ = v5
    _ = v6
    return 28
}
func genFn1(p0 uint32, p1 int32, p2 uint16) int8 {
    switch 56 {
    case -2147483648:
        v7 := GenS1{F0: 42}
        _ = v7
    case 32767:
        v8 := []int{-497, 692, 127, 203}
        _ = v8
    case 32768:
        for fi9 := 0; fi9 < 15; fi9++ {
            p0 += 0
            _ = fi9
        }
    default:
        v10 := GenS2{F0: "hello", F1: (p0 & (p0 / (((p0 + p0)) | 1))), F2: false}
        _ = v10
    }
    if ((0.0 / ((float64((1e10 * 3.14159)) - -2.25) / -2.25)) >= (3.14159 * 1e10)) {
        p2 += p2
        p0 += 2147483647
        fmt.Println(p0)
    } else {
        v11 := []int{-682, 1, 128}
        switch 2147483647 {
        case 407:
            if ((0 < 1000) || true) {
                for fi12 := 0; fi12 < 6; fi12++ {
                    if (true && true) {
                        switch ("hello" + "a") {
                        case "foo bar":
                            fmt.Println(p1)
                        case "hello":
                            fmt.Println(p1, p2, p1)
                        default:
                            fmt.Println(p2, fi12, p1)
                        }
                        switch "foo bar" {
                        case "foo bar":
                            v13 := GenS2{F0: (("\n\t" + "foo bar") + "foo bar"), F1: 1, F2: true}
                            _ = v13
                        default:
                            v11 = append(v11, (fi12 * (-(fi12))))
                        }
                    }
                    _ = fi12
                }
            }
        }
        if (!true) {
            var v14 bool = (!("" <= "foo bar"))
            p1 -= (p1 - genFn0(42))
            if len(v11) > 0 {
                v11[0] = -680
            }
            _ = v14
        } else {
            for fi15 := 0; fi15 < 15; fi15++ {
                for fi16 := 0; fi16 < 5; fi16++ {
                    v17 := GenS3{F0: p2, F1: 0, F2: 255}
                    switch "foo bar" {
                    case "":
                        switch "\n\t" {
                        case "":
                            v18 := GenS2{F0: "\n\t", F1: 65535, F2: (1 < -955)}
                            _ = v18
                        }
                    case "foo bar":
                        var v19 uint8 = 0
                        _ = v19
                    case "hello":
                        switch (fi16 / (((fi16 & 128)) | 1)) {
                        case 127:
                            v20 := GenS2{F0: "\n\t", F1: 4294967295, F2: (1 < -567)}
                            _ = v20
                        case 128:
                            fmt.Println(p1, p0)
                        case 902:
                            v21 := GenS3{F0: 778, F1: -1, F2: 255}
                            _ = v21
                        }
                    }
                    _ = fi16
                    _ = v17
                }
                var v22 float32 = -2.25
                _ = fi15
                _ = v22
            }
            var v23 int = -9223372036854775808
            _ = v23
        }
        _ = v11
    }
    var v24 int8 = -1
    fmt.Println(p2)
    _ = p0
    _ = p1
    _ = p2
    _ = v24
    return v24
}

func main() {
    var v1 int = 32768
    v2 := []int{(v1 * 42), v1}
    var v3 uint64 = 0
    var v4 float64 = (1.5 - -2.25)
    if (v4 == ((float64(-2.25) / v4) * 3.14159)) {
        if false {
            var v25 int64 = -9223372036854775808
            v4 = v4
            _ = v25
        } else {
            if false {
                if len(v2) > 0 {
                    v2[0] = v1
                }
                switch "\n\t" {
                case "hello":
                    v26 := GenS4{F0: 65535, F1: -978, F2: v3}
                    _ = v26
                }
            } else {
                v27 := GenS5{F0: 4294967295}
                _ = v27
            }
        }
        if false {
            _ = len(v2)
            var v28 int16 = 0
            _ = v28
        } else {
            if len(v2) >= 2 {
                v2 = v2[1:len(v2)]
            }
        }
        fmt.Println(v3, v1)
    } else {
        var v29 uint8 = 217
        _ = v29
    }
    var v30 string = "\n\t"
    v1 += (v1 & v1)
    for fi31 := 0; fi31 < 5; fi31++ {
        if true {
            for fi32 := 0; fi32 < 15; fi32++ {
                v3 = v3
                v2 = append(v2, (v1 >> ((566) & 63)))
                _ = fi32
            }
        } else {
            var v33 int = (fi31 / ((v1) | 1))
            _ = v33
        }
        _ = fi31
    }
    v3 = v3
    for fi34 := 0; fi34 < 5; fi34++ {
        fmt.Println(fi34, v30, v30)
        fmt.Println(v30, v30)
        if true {
            var v35 int8 = 42
            fmt.Println(fi34)
            for fi36 := 0; fi36 < 3; fi36++ {
                switch v1 {
                case -9223372036854775808:
                    v37 := GenS2{F0: "", F1: 179, F2: true}
                    _ = v37
                case -791:
                    if (true && false) {
                        if false {
                            if ((v1 - 127) <= fi36) {
                                fmt.Println(v1)
                                v38 := GenS3{F0: 32767, F1: v35, F2: 255}
                                var v39 int16 = 1000
                                _ = v38
                                _ = v39
                            }
                            v40 := GenS6{F0: 2147483647, F1: 4294967295}
                            v3 = 2
                            _ = v40
                        } else {
                            if (!(2 > v3)) {
                                v41 := GenS1{F0: 9223372036854775807}
                                var v42 float32 = 0.1
                                var v43 int32 = 0
                                _ = v41
                                _ = v42
                                _ = v43
                            }
                            switch fi36 {
                            case 32768:
                                var v44 bool = true
                                _ = v44
                            default:
                                _ = len(v2)
                            }
                            if false {
                                v45 := GenS7{F0: false, F1: 127}
                                _ = v45
                            } else {
                                var v46 uint64 = (v3 ^ v3)
                                _ = v46
                            }
                        }
                        if true {
                            var v47 uint16 = 55
                            if false {
                                v1 += fi36
                                var v48 uint16 = v47
                                _ = v48
                            } else {
                                _ = len(v2)
                                switch "" {
                                case "\n\t":
                                    if (int16(v47) >= 1000) {
                                        v49 := GenS8{F0: 185}
                                        v1 += (fi34 / ((fi34) | 1))
                                        _ = v49
                                    }
                                case "a":
                                    var v50 float32 = 1.5
                                    _ = v50
                                }
                                switch (fi34 & fi34) {
                                case 19:
                                    v51 := GenS9{F0: v30}
                                    _ = v51
                                case 32767:
                                    v1 -= -946
                                default:
                                    var v52 uint64 = v3
                                    _ = v52
                                }
                                var v53 string = v30
                                _ = v53
                            }
                            if ((0.0 / (3.14159 - 1.5)) != (0.1 - 1e10)) {
                                fmt.Println(fi36)
                                var v54 uint = 42
                                _ = v54
                            }
                            _ = v47
                        }
                    } else {
                        if true {
                            switch v30 {
                            case "\n\t":
                                switch "\n\t" {
                                case "\n\t":
                                    v35 = v35
                                case "hello":
                                    v1 = fi36
                                case "foo bar":
                                    if (v35 > v35) {
                                        var v55 uint8 = 255
                                        _ = v55
                                    }
                                }
                            case "":
                                v3 = (^(v3))
                            case "hello":
                                switch (-(fi36)) {
                                case 1:
                                    fmt.Println(v30, v1)
                                case -441:
                                    var v56 int8 = 42
                                    _ = v56
                                case -2147483648:
                                    var v57 uint16 = 65535
                                    _ = v57
                                default:
                                    fmt.Println(v3, v35, v3)
                                }
                            }
                            if (871 < 477) {
                                v1 = fi36
                                var v58 uint32 = 2147483648
                                _ = v58
                            } else {
                                if len(v2) > 0 {
                                    v2[0] = (fi34 * fi34)
                                }
                                switch v30 {
                                case "foo bar":
                                    fmt.Println(v4, v30)
                                default:
                                    var v59 int32 = -1
                                    _ = v59
                                }
                            }
                            var v60 uint64 = v3
                            _ = v60
                        } else {
                            var v61 int64 = -9223372036854775808
                            var v62 string = v30
                            if len(v2) > 0 {
                                v2[0] = (v1 % (((fi36 % ((v1) | 1))) | 1))
                            }
                            var v63 int16 = -132
                            _ = v61
                            _ = v62
                            _ = v63
                        }
                        v30 = "foo bar"
                        var v64 float32 = 0.1
                        if true {
                            v1 += fi34
                            switch 819 {
                            case 128:
                                fmt.Println(v64, fi34, v3)
                                _ = len(v2)
                                fmt.Println(v64, v1, v30)
                                if (((true && true) || false) || (!false)) {
                                    switch (fi36 % (((fi34 << ((2) & 63))) | 1)) {
                                    case -9223372036854775808:
                                        var v65 float64 = 1e10
                                        _ = v65
                                    case 853:
                                        if len(v2) >= 2 {
                                            v2 = v2[1:len(v2)]
                                        }
                                    default:
                                        switch v30 {
                                        case "a":
                                            fmt.Println(v1)
                                        case "hello":
                                            _ = len(v2)
                                        }
                                    }
                                } else {
                                    fmt.Println(v35, v1, v30)
                                }
                                fmt.Println(v64, fi34)
                            case -419:
                                if true {
                                    if len(v2) > 0 {
                                        v2[0] = (^(fi34))
                                    }
                                } else {
                                    fmt.Println(v30, fi34, v35)
                                    switch 42 {
                                    case 9223372036854775807:
                                        if true {
                                            if false {
                                                switch ((v30 + ("a" + ((v30 + (v30 + "a")) + v30))) + "") {
                                                case "foo bar":
                                                    fmt.Println(fi34, v4)
                                                case "a":
                                                    fmt.Println(v30, v30, v4)
                                                case "":
                                                    if false {
                                                        v66 := GenS5{F0: 65535}
                                                        _ = v66
                                                    }
                                                }
                                                fmt.Println(v1)
                                            } else {
                                                switch 9223372036854775807 {
                                                case 32767:
                                                    var v67 uint64 = v3
                                                    _ = v67
                                                case 42:
                                                    v2 = append(v2, fi34)
                                                default:
                                                    if false {
                                                        v68 := GenS7{F0: true, F1: 2}
                                                        var v69 int32 = genFn0(128)
                                                        v70 := GenS10{F0: 2, F1: 905, F2: uint(v3)}
                                                        _ = v68
                                                        _ = v69
                                                        _ = v70
                                                    } else {
                                                        var v71 int = 60
                                                        if false {
                                                            if false {
                                                                var v72 int = 4294967295
                                                                _ = v72
                                                            }
                                                        } else {
                                                            v1 = 42
                                                            v64 = v64
                                                        }
                                                        v1 = v71
                                                        _ = v71
                                                    }
                                                }
                                                if false {
                                                    switch v30 {
                                                    case "":
                                                        var v73 int32 = -1
                                                        var v74 int = 4294967295
                                                        _ = v73
                                                        _ = v74
                                                    }
                                                    v30 = "hello"
                                                    fmt.Println(v4)
                                                    v2 = append(v2, (-(v1)))
                                                } else {
                                                    fmt.Println(fi34, fi36, v3)
                                                    if (true || true) {
                                                        var v75 int64 = 430
                                                        v35 += 1
                                                        if len(v2) > 0 {
                                                            v2[0] = -91
                                                        }
                                                        _ = v75
                                                    }
                                                    v2 = append(v2, v1)
                                                    if true {
                                                        var v76 int64 = -345
                                                        fmt.Println(fi36, v35)
                                                        fmt.Println(v64, v4, v30)
                                                        fmt.Println(v35, v30)
                                                        v64 += (1.5 / float32((v4 + 0.0)))
                                                        v30 = v30
                                                        if (49 != 14) {
                                                            if (!false) {
                                                                fmt.Println(v76, v76, v35)
                                                            } else {
                                                                var v77 bool = (("a" + (v30 + ("hello" + "hello"))) != ("\n\t" + "foo bar"))
                                                                v64 -= (float32(3.14159) / ((v64 / 1e10) * (float32((0.1 / 0.1)) * ((v64 / float32((float64((float32(v4) / 3.14159)) - v4))) + (v64 * 0.1)))))
                                                                _ = v77
                                                            }
                                                            v3 = v3
                                                        }
                                                        _ = v76
                                                    } else {
                                                        if ((0 > 2147483647) || (79 != -32768)) {
                                                            var v78 int32 = genFn0(42)
                                                            v79 := GenS11{F0: v30, F1: 1}
                                                            _ = v78
                                                            _ = v79
                                                        } else {
                                                            if true {
                                                                if (v4 < (3.14159 / 1.5)) {
                                                                    v80 := GenS5{F0: 255}
                                                                    _ = v80
                                                                }
                                                                v4 = (0.1 + v4)
                                                            }
                                                        }
                                                    }
                                                }
                                                var v81 int = (fi34 / ((v1) | 1))
                                                v1 -= (-(v81))
                                                v35 = 127
                                                _ = v81
                                            }
                                            if (!false) {
                                                v2 = append(v2, v1)
                                            } else {
                                                var v82 float32 = (v64 + v64)
                                                switch (v1 - fi34) {
                                                case 107:
                                                    v30 = ""
                                                case 42:
                                                    var v83 uint64 = v3
                                                    _ = v83
                                                case 951:
                                                    if false {
                                                        v64 = v64
                                                        if false {
                                                            v84 := GenS10{F0: 0, F1: 42, F2: 42}
                                                            _ = v84
                                                        } else {
                                                            if false {
                                                                fmt.Println(v35)
                                                                v85 := GenS7{F0: false, F1: 6}
                                                                if len(v2) >= 2 {
                                                                    v2 = v2[1:len(v2)]
                                                                }
                                                                _ = v85
                                                            } else {
                                                                if true {
                                                                    fmt.Println(v64, v3, v30)
                                                                    switch (v30 + v30) {
                                                                    case "hello":
                                                                        fmt.Println(v1, v30)
                                                                    }
                                                                }
                                                                fmt.Println(v3)
                                                            }
                                                            var v86 string = (v30 + "\n\t")
                                                            fmt.Println(v35)
                                                            _ = v86
                                                        }
                                                    } else {
                                                        var v87 uint32 = 964
                                                        v88 := GenS4{F0: (v87 << ((2) & 63)), F1: -1, F2: (v3 << ((418) & 63))}
                                                        switch fi36 {
                                                        case 127:
                                                            if len(v2) > 0 {
                                                                v2[0] = v1
                                                            }
                                                        case 32767:
                                                            switch (v30 + (v30 + v30)) {
                                                            case "":
                                                                var v89 uint16 = 70
                                                                _ = v89
                                                            case "\n\t":
                                                                v2 = append(v2, 52)
                                                            }
                                                        default:
                                                            v3 -= (v3 | v3)
                                                        }
                                                        v3 += (v3 >> ((841) & 63))
                                                        switch (fi36 / ((v1) | 1)) {
                                                        case 32768:
                                                            v90 := GenS12{F0: v3, F1: 102, F2: 591}
                                                            _ = v90
                                                        case -1:
                                                            v35 = (v35 + (v35 - (v35 - (v35 / ((127) | 1)))))
                                                        case 42:
                                                            v91 := GenS13{F0: (((("hello" + v30) + ("hello" + v30)) + "a") != (v30 + v30)), F1: v1}
                                                            _ = v91
                                                        default:
                                                            var v92 float64 = v4
                                                            _ = v92
                                                        }
                                                        _ = v87
                                                        _ = v88
                                                    }
                                                }
                                                _ = v82
                                            }
                                        } else {
                                            v93 := GenS14{F0: ((((-919 >= 1) && false) || true) && (128 == 42)), F1: -46, F2: -32768}
                                            if false {
                                                if len(v2) > 0 {
                                                    v2[0] = (fi36 ^ (fi36 ^ 127))
                                                }
                                            }
                                            _ = v93
                                        }
                                    case 128:
                                        if false {
                                            v94 := GenS15{F0: 0}
                                            _ = len(v2)
                                            v64 -= 1.5
                                            _ = v94
                                        }
                                    case -948:
                                        v4 -= ((3.14159 + (3.14159 - 0.1)) / v4)
                                    default:
                                        var v95 uint = 673
                                        _ = v95
                                    }
                                    switch ((v30 + v30) + (v30 + "a")) {
                                    case "\n\t":
                                        fmt.Println(v30, fi34, fi36)
                                    case "a":
                                        var v96 int16 = 42
                                        _ = v96
                                    case "hello":
                                        v2 = append(v2, (v1 >> ((uint(v3)) & 63)))
                                    }
                                }
                            case -2147483648:
                                fmt.Println(v64, v35)
                            default:
                                var v97 uint16 = 0
                                _ = v97
                            }
                        } else {
                            var v98 int = 784
                            _ = v98
                        }
                        _ = v64
                    }
                default:
                    switch (fi34 << ((309) & 63)) {
                    case -775:
                        v1 -= (fi34 ^ (fi34 ^ (fi34 + 9223372036854775807)))
                    default:
                        v4 = ((v4 / v4) - ((v4 - -2.25) / (3.14159 * v4)))
                    }
                }
                if false {
                    v4 -= (float64(1e10) - v4)
                } else {
                    var v99 int = v1
                    _ = v99
                }
                v100 := GenS1{F0: -1}
                _ = fi36
                _ = v100
            }
            _ = v35
        } else {
            v3 += (v3 + (^(v3)))
            v30 = v30
        }
        _ = fi34
    }
    v101 := GenS1{F0: -1}
    if ((true && false) || (4294967295 < 859)) {
        v4 -= v4
        fmt.Println(v3)
    }
    for fi102 := 0; fi102 < 20; fi102++ {
        v103 := GenS2{F0: v30, F1: 65535, F2: (255 == 167)}
        var v104 int32 = -1
        _ = fi102
        _ = v103
        _ = v104
    }
    if false {
        if len(v2) >= 2 {
            v2 = v2[1:len(v2)]
        }
        v4 = (v4 / 1.5)
    }
    if (2 != (v3 | (^(v3)))) {
        var v105 float32 = -2.25
        for fi106 := 0; fi106 < 3; fi106++ {
            v1 += (fi106 >> ((109) & 63))
            var v107 string = v30
            _ = fi106
            _ = v107
        }
        _ = v105
    } else {
        switch v30 {
        case "a":
            fmt.Println(v4, v1, v30)
        case "hello":
            for fi108 := 0; fi108 < 16; fi108++ {
                v4 += 0.1
                v2 = append(v2, fi108)
                fmt.Println(v30)
                _ = fi108
            }
        }
        var v109 int16 = 42
        v2 = append(v2, (v1 << ((18446744073709551615) & 63)))
        _ = v109
    }
    var v110 int32 = 985
    v3 += v3
    for fi111 := 0; fi111 < 1; fi111++ {
        var v112 int = 2147483647
        var v113 uint = 3
        _ = fi111
        _ = v112
        _ = v113
    }
    if len(v2) > 0 {
        v2[0] = (v1 & (v1 << ((65535) & 63)))
    }
    for fi114 := 0; fi114 < 2; fi114++ {
        fmt.Println(v3)
        _ = fi114
    }
    if (!(269 == 0)) {
        fmt.Println(v4, v4)
        v4 += 1e10
    } else {
        v115 := GenS10{F0: 15, F1: 18446744073709551615, F2: 0}
        var v116 uint8 = uint8(v3)
        _ = v115
        _ = v116
    }
    if len(v2) > 0 {
        v2[0] = (v1 % ((32768) | 1))
    }
    fmt.Println("v1:", v1)
    fmt.Println("v110:", v110)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v101
    _ = v110
    _ = v2
    _ = v3
    _ = v30
    _ = v4
}
