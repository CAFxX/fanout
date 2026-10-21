package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 string; F2 uint16 }
type GenS2 struct { F0 bool; F1 int32; F2 uint64 }
type GenS3 struct { F0 string; F1 uint8 }
type GenS4 struct { F0 uint8 }
type GenS5 struct { F0 uint8 }
type GenS6 struct { F0 int64; F1 uint64; F2 int32 }
type GenS7 struct { F0 int32; F1 uint32 }
type GenS8 struct { F0 uint; F1 bool; F2 int }
type GenS9 struct { F0 uint }
type GenS10 struct { F0 int16; F1 uint32 }
type GenS11 struct { F0 uint8; F1 uint32 }
type GenS12 struct { F0 int; F1 int8 }
type GenS13 struct { F0 string; F1 uint }
type GenS14 struct { F0 uint32 }
type GenS15 struct { F0 bool; F1 int8; F2 string }
type GenS16 struct { F0 uint; F1 int32 }
type GenS17 struct { F0 int64; F1 int16; F2 uint }
type GenS18 struct { F0 int64 }
type GenS19 struct { F0 int8; F1 uint; F2 uint64 }
type GenS20 struct { F0 int8 }
type GenS21 struct { F0 bool; F1 int8 }
type GenS22 struct { F0 int8; F1 int; F2 uint32 }
type GenS23 struct { F0 int8; F1 uint8 }
func genFn0(p0 uint, p1 int32) bool {
    var v5 float64 = (0.0 / -2.25)
    for fi6 := 0; fi6 < 13; fi6++ {
        v7 := GenS1{F0: 42, F1: "", F2: 0}
        _ = fi6
        _ = v7
    }
    fmt.Println(p0, v5)
    _ = p0
    _ = p1
    _ = v5
    return (true || false)
}
func genFn1(p0 uint, p1 int32) int32 {
    v8 := GenS2{F0: (144 <= 42), F1: p1, F2: 42}
    var v9 uint = 466
    switch (("" + ("hello" + "hello")) + ("a" + "a")) {
    case "":
        switch -499 {
        case 4294967295:
            v10 := GenS2{F0: (284 < 32767), F1: p1, F2: 166}
            _ = v10
        default:
            switch "\n\t" {
            case "\n\t":
                var v11 int16 = 1
                _ = v11
            default:
                v12 := GenS3{F0: "", F1: 1}
                _ = v12
            }
        }
    case "foo bar":
        switch -61 {
        case -2147483648:
            var v13 uint = p0
            _ = v13
        case 470:
            switch 0 {
            case 0:
                var v14 bool = false
                _ = v14
            case 4294967295:
                if ((!true) && (-329 < -9223372036854775808)) {
                    fmt.Println(p1, p1)
                    v9 = p0
                } else {
                    if (false || true) {
                        var v15 uint32 = 1
                        _ = v15
                    }
                    var v16 bool = (!(uint16(v9) > 892))
                    _ = v16
                }
            case -1:
                fmt.Println(p0, p1)
            default:
                for fi17 := 0; fi17 < 4; fi17++ {
                    v18 := GenS2{F0: (false && false), F1: p1, F2: 42}
                    v9 = p0
                    _ = fi17
                    _ = v18
                }
            }
        }
    }
    var v19 float32 = (1.5 - 0.0)
    fmt.Println(p0, p1)
    _ = p0
    _ = p1
    _ = v19
    _ = v8
    _ = v9
    return int32(v9)
}

func main() {
    var v1 int16 = 590
    v2 := make(map[int]bool)
    var v3 int16 = v1
    v4 := []int{2147483647, 32767}
    if false {
        v20 := GenS1{F0: 32767, F1: (("hello" + ("\n\t" + "\n\t")) + "a"), F2: 255}
        for fi21 := 0; fi21 < 12; fi21++ {
            v1 -= v3
            if true {
                var v22 int = fi21
                if (0.0 >= ((((3.14159 + float32(3.14159)) * 0.0) - (0.1 * (1.5 / 0.1))) - ((((3.14159 / float32(1e10)) - ((((1.5 * 0.0) / 1.0) * ((0.1 / 1.0) + -2.25)) + float32(0.0))) - (1e10 + 1.5)) + (1.5 / (-2.25 + 3.14159))))) {
                    v23 := GenS3{F0: "a", F1: 239}
                    v24 := v2[v22]
                    var v25 int16 = v3
                    _ = v23
                    _ = v24
                    _ = v25
                }
                _ = v22
            } else {
                if false {
                    var v26 int32 = 144
                    v27 := GenS1{F0: 0, F1: "a", F2: 42}
                    _ = v26
                    _ = v27
                }
                if (!(976 != 1)) {
                    if (1.5 == 3.14159) {
                        delete(v2, fi21)
                        for fi28 := 0; fi28 < 9; fi28++ {
                            if ((("" + "a") + "a") <= ("" + ("foo bar" + ((("hello" + "") + "\n\t") + "foo bar")))) {
                                switch (fi28 + fi28) {
                                case 1:
                                    switch -1 {
                                    case -9223372036854775808:
                                        fmt.Println(fi21, v3)
                                    case 32768:
                                        var v29 bool = true
                                        _ = v29
                                    default:
                                        switch ("hello" + "foo bar") {
                                        case "a":
                                            fmt.Println(v3)
                                        }
                                    }
                                default:
                                    switch "a" {
                                    case "":
                                        v3 = v3
                                    case "a":
                                        v1 = (v1 % ((v1) | 1))
                                        var v30 int8 = 127
                                        _ = v30
                                    }
                                }
                            }
                            _ = fi28
                        }
                        v31 := GenS4{F0: 2}
                        _ = v31
                    } else {
                        var v32 uint16 = 42
                        if (3.14159 > float32(3.14159)) {
                            v32 -= v32
                        }
                        _ = v32
                    }
                    v1 += v1
                    v3 = (v1 / ((1000) | 1))
                }
            }
            _ = fi21
        }
        _ = v20
    }
    var v33 int = -9223372036854775808
    var v34 int = v33
    for fi35 := 0; fi35 < 12; fi35++ {
        for fi36 := 0; fi36 < 9; fi36++ {
            v3 = v3
            _ = fi36
        }
        for fi37 := 0; fi37 < 16; fi37++ {
            var v38 int64 = -622
            if (!(!true)) {
                var v39 uint8 = 1
                var v40 int64 = 9223372036854775807
                _ = v39
                _ = v40
            }
            if false {
                fmt.Println(v1, v38)
                if true {
                    var v41 int = fi35
                    switch 127 {
                    case 127:
                        var v42 float64 = 1e10
                        _ = v42
                    case -2147483648:
                        if true {
                            fmt.Println(v33, v34, v3)
                        }
                    default:
                        var v43 uint8 = 42
                        if (!("\n\t" == (("hello" + "\n\t") + "hello"))) {
                            if (-2.25 == (1e10 + ((((-2.25 / 1e10) - 1e10) - (0.1 * (float64((0.1 - 0.0)) / 1e10))) / (((3.14159 / 3.14159) + 1e10) - 0.0)))) {
                                _ = len(v2)
                                if (0.1 == ((1.5 + (1e10 - -2.25)) + (0.0 + -2.25))) {
                                    v3 -= (v3 >> ((402) & 63))
                                    v44 := v2[954]
                                    _ = v44
                                } else {
                                    var v45 float32 = (3.14159 + 0.0)
                                    _ = v45
                                }
                                if (42 <= 255) {
                                    switch ("a" + "a") {
                                    case "\n\t":
                                        delete(v2, v33)
                                    case "hello":
                                        v1 = -1
                                    case "foo bar":
                                        if true {
                                            var v46 int16 = v1
                                            fmt.Println(v41, v38, fi35)
                                            _ = v46
                                        } else {
                                            switch (fi35 - -659) {
                                            case 1:
                                                v38 = 9223372036854775807
                                            case -548:
                                                v1 = (v3 | v3)
                                                if false {
                                                    v47 := GenS2{F0: (786 < 2147483647), F1: 2147483647, F2: 728}
                                                    if true {
                                                        var v48 int16 = (v3 - v3)
                                                        if true {
                                                            fmt.Println(v33, v1)
                                                        } else {
                                                            _ = len(v2)
                                                            var v49 uint32 = 42
                                                            if (("hello" + "foo bar") == "\n\t") {
                                                                v4 = append(v4, 270)
                                                                v2[v34] = true
                                                            } else {
                                                                switch v41 {
                                                                case 127:
                                                                    if false {
                                                                        if (!(true && false)) {
                                                                            if len(v4) >= 2 {
                                                                                v4 = v4[1:len(v4)]
                                                                            }
                                                                            var v50 float32 = 1.5
                                                                            if true {
                                                                                v43 += v43
                                                                                if (v50 <= ((v50 * ((-2.25 + v50) * 0.0)) - v50)) {
                                                                                    var v51 uint32 = (v49 - v49)
                                                                                    _ = v51
                                                                                }
                                                                            }
                                                                            _ = v50
                                                                        } else {
                                                                            if (18446744073709551615 != 255) {
                                                                                var v52 int16 = (v1 + v1)
                                                                                _ = len(v2)
                                                                                v43 += v43
                                                                                _ = v52
                                                                            }
                                                                        }
                                                                        if true {
                                                                            if true {
                                                                                v53 := GenS2{F0: true, F1: 82, F2: 528}
                                                                                _ = v53
                                                                            } else {
                                                                                var v54 int = fi35
                                                                                if true {
                                                                                    var v55 string = (("hello" + "a") + (((("foo bar" + "\n\t") + ((((("\n\t" + ("\n\t" + ("hello" + (("foo bar" + ("\n\t" + "\n\t")) + (("" + "\n\t") + (("foo bar" + "foo bar") + (("\n\t" + ("foo bar" + (("\n\t" + "hello") + "foo bar"))) + "hello"))))))) + "a") + "foo bar") + (("hello" + "\n\t") + ("" + "foo bar"))) + "")) + "\n\t") + ("hello" + "")))
                                                                                    var v56 uint = 65535
                                                                                    _ = v55
                                                                                    _ = v56
                                                                                }
                                                                                var v57 uint = 198
                                                                                _ = v54
                                                                                _ = v57
                                                                            }
                                                                            v2[42] = ((("foo bar" > ((("" + (("a" + ("\n\t" + "foo bar")) + "\n\t")) + "\n\t") + "a")) || (530 < 0)) || (!true))
                                                                        } else {
                                                                            fmt.Println(v1)
                                                                        }
                                                                    } else {
                                                                        switch "foo bar" {
                                                                        case "hello":
                                                                            switch ("hello" + "\n\t") {
                                                                            case "hello":
                                                                                if len(v4) > 0 {
                                                                                    v4[0] = v41
                                                                                }
                                                                            case "foo bar":
                                                                                v49 -= uint32(v43)
                                                                            default:
                                                                                if true {
                                                                                    v4 = append(v4, 2147483647)
                                                                                }
                                                                            }
                                                                        case "a":
                                                                            if true {
                                                                                if true {
                                                                                    v49 = v49
                                                                                    var v58 int8 = 0
                                                                                    _ = v58
                                                                                }
                                                                            } else {
                                                                                switch "hello" {
                                                                                case "\n\t":
                                                                                    v34 = (^(v34))
                                                                                case "hello":
                                                                                    v41 -= fi37
                                                                                }
                                                                                var v59 uint32 = 646
                                                                                _ = v59
                                                                            }
                                                                        case "foo bar":
                                                                            fmt.Println(fi35)
                                                                            var v60 int = (^(v41))
                                                                            var v61 int64 = (v38 - -9223372036854775808)
                                                                            _ = v60
                                                                            _ = v61
                                                                        }
                                                                    }
                                                                case 4294967295:
                                                                    v62 := GenS5{F0: v43}
                                                                    _ = v62
                                                                case 1:
                                                                    var v63 int = v41
                                                                    v48 += 12
                                                                    _ = v63
                                                                }
                                                            }
                                                            if (false && true) {
                                                                var v64 string = "foo bar"
                                                                if (!true) {
                                                                    if false {
                                                                        if len(v4) >= 2 {
                                                                            v4 = v4[1:len(v4)]
                                                                        }
                                                                        var v65 int8 = -46
                                                                        _ = v65
                                                                    } else {
                                                                        switch v64 {
                                                                        case "\n\t":
                                                                            v43 = 128
                                                                            var v66 string = (v64 + v64)
                                                                            _ = v66
                                                                        case "hello":
                                                                            if true {
                                                                                if false {
                                                                                    var v67 uint16 = 1
                                                                                    _ = v67
                                                                                } else {
                                                                                    var v68 uint32 = 42
                                                                                    _ = v68
                                                                                }
                                                                            }
                                                                        case "a":
                                                                            v69 := GenS6{F0: -1, F1: 1, F2: 1}
                                                                            _ = v69
                                                                        }
                                                                        if false {
                                                                            fmt.Println(v38, v1, v34)
                                                                            var v70 string = ""
                                                                            var v71 int32 = 970
                                                                            _ = v70
                                                                            _ = v71
                                                                        }
                                                                        fmt.Println(v64)
                                                                    }
                                                                    var v72 uint64 = 2
                                                                    switch v33 {
                                                                    case 593:
                                                                        var v73 uint64 = 0
                                                                        _ = v73
                                                                    }
                                                                    switch 127 {
                                                                    case 32767:
                                                                        switch (-(v34)) {
                                                                        case 2147483647:
                                                                            _ = len(v4)
                                                                        }
                                                                    case -9223372036854775808:
                                                                        if false {
                                                                            v74 := GenS3{F0: (((v64 + v64) + v64) + "hello"), F1: (v43 + 1)}
                                                                            switch fi35 {
                                                                            case 2147483647:
                                                                                var v75 int16 = (v1 * v3)
                                                                                _ = v75
                                                                            case -649:
                                                                                var v76 bool = true
                                                                                _ = v76
                                                                            }
                                                                            _ = len(v2)
                                                                            var v77 float32 = 1e10
                                                                            _ = v74
                                                                            _ = v77
                                                                        } else {
                                                                            if true {
                                                                                v41 = fi35
                                                                                var v78 uint8 = v43
                                                                                v79 := GenS3{F0: v64, F1: v43}
                                                                                _ = len(v4)
                                                                                _ = v78
                                                                                _ = v79
                                                                            } else {
                                                                                v4 = append(v4, (v33 >> ((4294967295) & 63)))
                                                                                v48 += v3
                                                                            }
                                                                            var v80 uint16 = 0
                                                                            v72 += (v72 % ((4294967295) | 1))
                                                                            _ = v80
                                                                        }
                                                                    }
                                                                    var v81 int64 = -920
                                                                    v82 := GenS7{F0: 1, F1: 944}
                                                                    var v83 int16 = 0
                                                                    v38 -= v38
                                                                    switch v64 {
                                                                    case "foo bar":
                                                                        v1 -= v48
                                                                    case "hello":
                                                                        v83 -= 1
                                                                    default:
                                                                        var v84 int64 = (v81 + v81)
                                                                        _ = v84
                                                                    }
                                                                    _ = v72
                                                                    _ = v81
                                                                    _ = v82
                                                                    _ = v83
                                                                } else {
                                                                    v85 := GenS8{F0: 42, F1: false, F2: (^(v34))}
                                                                    _ = v85
                                                                }
                                                                v86 := GenS4{F0: (^(v43))}
                                                                _ = v64
                                                                _ = v86
                                                            }
                                                            v87 := GenS9{F0: 1}
                                                            _ = v49
                                                            _ = v87
                                                        }
                                                        v2[fi37] = true
                                                        _ = v48
                                                    }
                                                    _ = v47
                                                }
                                                var v88 uint8 = v43
                                                _ = v88
                                            case 0:
                                                if true {
                                                    if (!false) {
                                                        fmt.Println(v43)
                                                        if true {
                                                            var v89 bool = (true && true)
                                                            _ = v89
                                                        }
                                                    } else {
                                                        if true {
                                                            _ = len(v2)
                                                        } else {
                                                            if (((v3 < v3) && false) || (42 != int8(v1))) {
                                                                switch fi35 {
                                                                case 1:
                                                                    if true {
                                                                        v3 = v1
                                                                        var v90 int8 = 0
                                                                        var v91 bool = false
                                                                        _ = v90
                                                                        _ = v91
                                                                    } else {
                                                                        var v92 uint16 = 998
                                                                        _ = v92
                                                                    }
                                                                case -19:
                                                                    if len(v4) >= 2 {
                                                                        v4 = v4[1:len(v4)]
                                                                    }
                                                                    var v93 uint16 = 1
                                                                    v94 := GenS10{F0: (v3 << ((482) & 63)), F1: 2147483647}
                                                                    v1 += (v1 ^ v3)
                                                                    var v95 int64 = 42
                                                                    _ = v93
                                                                    _ = v94
                                                                    _ = v95
                                                                case 4294967295:
                                                                    v41 -= -2147483648
                                                                }
                                                                v41 += (^(fi37))
                                                            } else {
                                                                var v96 uint64 = 731
                                                                var v97 int = fi35
                                                                fmt.Println(v41, v43, v3)
                                                                _ = v96
                                                                _ = v97
                                                            }
                                                            v98 := GenS10{F0: 974, F1: 65535}
                                                            _ = v98
                                                        }
                                                        v34 = v41
                                                        switch (v33 >> ((42) & 63)) {
                                                        case 9223372036854775807:
                                                            v99 := GenS11{F0: (v43 ^ v43), F1: 42}
                                                            _ = v99
                                                        case -746:
                                                            var v100 uint32 = 0
                                                            _ = v100
                                                        case 42:
                                                            v101 := GenS8{F0: 18446744073709551615, F1: true, F2: (fi37 << ((0) & 63))}
                                                            _ = v101
                                                        }
                                                    }
                                                    switch fi35 {
                                                    case -1:
                                                        fmt.Println(v34)
                                                    case 2147483647:
                                                        if (42 <= 175) {
                                                            v3 -= (v3 ^ v1)
                                                        } else {
                                                            v3 -= (v3 << ((101) & 63))
                                                            if false {
                                                                v4 = append(v4, -668)
                                                                if len(v4) > 0 {
                                                                    v4[0] = 0
                                                                }
                                                            } else {
                                                                var v102 int8 = int8(v43)
                                                                if (1 < v102) {
                                                                    _ = len(v2)
                                                                    if true {
                                                                        v102 -= 100
                                                                        if true {
                                                                            v103 := GenS3{F0: "\n\t", F1: v43}
                                                                            _ = v103
                                                                        } else {
                                                                            v2[-797] = true
                                                                            if false {
                                                                                fmt.Println(v102, v41, v43)
                                                                                _ = len(v4)
                                                                            } else {
                                                                                var v104 bool = (true || (420 > 1))
                                                                                _ = v104
                                                                            }
                                                                            var v105 float32 = -2.25
                                                                            _ = v105
                                                                        }
                                                                    } else {
                                                                        if (2 != 645) {
                                                                            switch (v33 - (v34 << ((0) & 63))) {
                                                                            case 4294967295:
                                                                                if true {
                                                                                    var v106 int64 = v38
                                                                                    fmt.Println(fi37, v3)
                                                                                    _ = v106
                                                                                }
                                                                            case -9223372036854775808:
                                                                                if false {
                                                                                    if true {
                                                                                        var v107 string = ""
                                                                                        _ = v107
                                                                                    } else {
                                                                                        var v108 uint8 = (v43 - 45)
                                                                                        v43 -= v43
                                                                                        _ = v108
                                                                                    }
                                                                                } else {
                                                                                    v33 = int(v102)
                                                                                    v38 += v38
                                                                                }
                                                                            case 9223372036854775807:
                                                                                v38 -= v38
                                                                                var v109 int64 = -339
                                                                                v110 := GenS12{F0: 1, F1: v102}
                                                                                _ = v109
                                                                                _ = v110
                                                                            default:
                                                                                fmt.Println(v43, v3)
                                                                            }
                                                                            var v111 bool = true
                                                                            var v112 int64 = v38
                                                                            _ = v111
                                                                            _ = v112
                                                                        }
                                                                    }
                                                                } else {
                                                                    if true {
                                                                        var v113 uint64 = 4294967295
                                                                        v33 = fi37
                                                                        _ = v113
                                                                    }
                                                                }
                                                                _ = v102
                                                            }
                                                        }
                                                    }
                                                    v114 := GenS13{F0: "\n\t", F1: 272}
                                                    _ = v114
                                                } else {
                                                    var v115 uint16 = 2
                                                    switch (v34 << ((4294967295) & 63)) {
                                                    case 32768:
                                                        if (false || true) {
                                                            fmt.Println(v33, v34, v33)
                                                            v116 := GenS14{F0: 2}
                                                            v43 -= 1
                                                            _ = v116
                                                        } else {
                                                            _ = len(v4)
                                                            switch "a" {
                                                            case "\n\t":
                                                                fmt.Println(fi37)
                                                            case "hello":
                                                                fmt.Println(v43)
                                                            default:
                                                                var v117 string = ("a" + "\n\t")
                                                                _ = v117
                                                            }
                                                        }
                                                    case 4294967295:
                                                        v2[(fi35 / ((fi35) | 1))] = false
                                                    case 32767:
                                                        var v118 uint32 = 4294967295
                                                        _ = v118
                                                    }
                                                    if (18446744073709551615 < 255) {
                                                        fmt.Println(v41, v34, v34)
                                                    }
                                                    _ = v115
                                                }
                                            }
                                        }
                                    }
                                    _ = len(v4)
                                }
                            } else {
                                _ = len(v4)
                                v119 := GenS15{F0: true, F1: 80, F2: ""}
                                _ = v119
                            }
                            v4 = append(v4, (-(v33)))
                        }
                        v120 := GenS16{F0: 2, F1: genFn1(601, 741)}
                        fmt.Println(v41, fi37)
                        v121 := GenS1{F0: 46, F1: ((("" + "") + "a") + ("" + "\n\t")), F2: 1}
                        var v122 uint32 = 2147483648
                        _ = v120
                        _ = v121
                        _ = v122
                        _ = v43
                    }
                    var v123 string = "\n\t"
                    _ = v123
                    _ = v41
                } else {
                    fmt.Println(v38, v38, v33)
                }
                v38 = (-(v38))
            } else {
                v3 -= v1
                if true {
                    var v124 uint32 = 1
                    _ = v124
                }
            }
            _ = fi37
            _ = v38
        }
        fmt.Println(v34, v1)
        _ = fi35
    }
    switch "" {
    case "":
        var v125 uint16 = 32767
        _ = v125
    case "hello":
        if ((("foo bar" + "") + ("\n\t" + "hello")) <= ("" + "\n\t")) {
            var v126 string = "hello"
            _ = len(v2)
            _ = v126
        }
    default:
        var v127 uint8 = 1
        _ = v127
    }
    if (!false) {
        v34 -= (v34 >> ((4294967295) & 63))
        fmt.Println(v1, v33)
    } else {
        v2[4294967295] = false
    }
    v128 := GenS1{F0: 1, F1: "", F2: 0}
    fmt.Println(v34, v3)
    if true {
        v129 := GenS16{F0: 846, F1: -582}
        _ = v129
    }
    if true {
        for fi130 := 0; fi130 < 11; fi130++ {
            var v131 uint32 = 65535
            v1 -= (v3 % ((-1) | 1))
            var v132 int = (v34 & v33)
            _ = fi130
            _ = v131
            _ = v132
        }
    } else {
        v33 += (v34 + -2147483648)
        fmt.Println(v34, v3, v3)
        if true {
            var v133 uint64 = 0
            if ((v3 % ((-1) | 1)) < -32768) {
                v33 -= 32767
                var v134 uint16 = 255
                var v135 int32 = 30
                _ = v134
                _ = v135
            } else {
                switch "" {
                case "":
                    var v136 uint64 = v133
                    _ = v136
                }
                if len(v4) > 0 {
                    v4[0] = v34
                }
                for fi137 := 0; fi137 < 11; fi137++ {
                    var v138 uint8 = 88
                    var v139 uint32 = 0
                    var v140 float32 = ((1.5 * 1e10) - 1.5)
                    _ = fi137
                    _ = v138
                    _ = v139
                    _ = v140
                }
            }
            if (630 > 2) {
                for fi141 := 0; fi141 < 13; fi141++ {
                    v142 := GenS17{F0: 530, F1: v3, F2: 4294967295}
                    for fi143 := 0; fi143 < 18; fi143++ {
                        v33 = fi143
                        var v144 int64 = -1
                        _ = fi143
                        _ = v144
                    }
                    var v145 int16 = (v3 ^ v3)
                    _ = fi141
                    _ = v142
                    _ = v145
                }
                if (-9223372036854775808 <= 776) {
                    switch -825 {
                    case 127:
                        fmt.Println(v1, v3, v1)
                    case 32767:
                        v146 := GenS18{F0: 458}
                        _ = v146
                    case 4294967295:
                        for fi147 := 0; fi147 < 12; fi147++ {
                            switch ("a" + ("a" + "")) {
                            case "hello":
                                for fi148 := 0; fi148 < 15; fi148++ {
                                    if (-2.25 > 1e10) {
                                        fmt.Println(v1, fi148)
                                        v133 -= (v133 - (v133 + (v133 & v133)))
                                        var v149 string = "\n\t"
                                        _ = v149
                                    } else {
                                        v34 = v33
                                        v3 -= v3
                                        var v150 uint8 = 41
                                        _ = v150
                                    }
                                    switch "foo bar" {
                                    case "hello":
                                        if len(v4) > 0 {
                                            v4[0] = 127
                                        }
                                    case "":
                                        var v151 string = ("" + "a")
                                        _ = v151
                                    case "a":
                                        fmt.Println(v33, fi148)
                                    default:
                                        if true {
                                            _ = len(v2)
                                            fmt.Println(v1, v33, v133)
                                            _ = len(v2)
                                            v152 := GenS19{F0: -128, F1: 42, F2: v133}
                                            _ = v152
                                        }
                                    }
                                    var v153 float64 = 0.1
                                    _ = fi148
                                    _ = v153
                                }
                            case "foo bar":
                                for fi154 := 0; fi154 < 10; fi154++ {
                                    switch (v33 / (((-(v33))) | 1)) {
                                    case 269:
                                        var v155 int32 = -2147483648
                                        _ = v155
                                    case -564:
                                        v1 = (v1 % (((v1 % ((-1) | 1))) | 1))
                                    case 917:
                                        v156 := GenS16{F0: 120, F1: genFn1(42, 42)}
                                        var v157 int16 = v3
                                        var v158 bool = (1 < (v133 / ((42) | 1)))
                                        _ = v156
                                        _ = v157
                                        _ = v158
                                    }
                                    var v159 int = int(v133)
                                    var v160 uint64 = (^(v133))
                                    _ = fi154
                                    _ = v159
                                    _ = v160
                                }
                            }
                            _ = fi147
                        }
                    default:
                        for fi161 := 0; fi161 < 6; fi161++ {
                            switch (v33 | fi161) {
                            case -9223372036854775808:
                                var v162 bool = (!true)
                                _ = v162
                            case 698:
                                v163 := v2[(fi161 * fi161)]
                                _ = v163
                            default:
                                fmt.Println(v1)
                            }
                            for fi164 := 0; fi164 < 1; fi164++ {
                                v3 += v3
                                if false {
                                    if ((0.0 - 3.14159) == (1e10 - float64(-2.25))) {
                                        fmt.Println(v1, v3)
                                    } else {
                                        v3 += v3
                                        v3 = 42
                                    }
                                    var v165 int32 = 2147483647
                                    _ = len(v2)
                                    _ = v165
                                } else {
                                    if len(v4) >= 2 {
                                        v4 = v4[1:len(v4)]
                                    }
                                    switch "" {
                                    case "hello":
                                        if true {
                                            v166 := GenS10{F0: (v1 - 0), F1: 42}
                                            fmt.Println(v1, v34)
                                            if (0 > 1) {
                                                if false {
                                                    if false {
                                                        var v167 int = fi164
                                                        _ = v167
                                                    } else {
                                                        v3 = -32768
                                                        v168 := GenS20{F0: -92}
                                                        var v169 uint = 2
                                                        _ = v168
                                                        _ = v169
                                                    }
                                                    switch ("hello" + (((((("a" + "hello") + "a") + (("a" + (("a" + "a") + "hello")) + "a")) + "hello") + "foo bar") + "\n\t")) {
                                                    case "hello":
                                                        switch (fi161 ^ (fi161 + 4294967295)) {
                                                        case 9223372036854775807:
                                                            var v170 uint16 = 973
                                                            _ = v170
                                                        case 0:
                                                            switch 150 {
                                                            case -328:
                                                                var v171 int8 = 100
                                                                _ = v171
                                                            }
                                                        case 42:
                                                            var v172 uint64 = 18446744073709551615
                                                            _ = v172
                                                        default:
                                                            var v173 uint = 1
                                                            _ = v173
                                                        }
                                                    }
                                                } else {
                                                    if false {
                                                        if len(v4) >= 2 {
                                                            v4 = v4[1:len(v4)]
                                                        }
                                                        v174 := v2[v33]
                                                        var v175 bool = (v3 <= 32767)
                                                        _ = v174
                                                        _ = v175
                                                    } else {
                                                        var v176 uint64 = v133
                                                        _ = v176
                                                    }
                                                    if false {
                                                        v3 -= (v3 >> ((2) & 63))
                                                        v133 = 857
                                                        v133 = v133
                                                    } else {
                                                        fmt.Println(v133, v133)
                                                        v34 += 9223372036854775807
                                                        var v177 int32 = -1
                                                        v2[v34] = (1 >= 42)
                                                        if true {
                                                            _ = len(v4)
                                                            if len(v4) > 0 {
                                                                v4[0] = fi164
                                                            }
                                                            v3 = (v1 * (v3 % ((int16(v133)) | 1)))
                                                        }
                                                        _ = v177
                                                    }
                                                    v33 = (-(fi164))
                                                }
                                            }
                                            _ = v166
                                        }
                                        v34 -= 930
                                    case "":
                                        var v178 int8 = 127
                                        _ = v178
                                    }
                                    v179 := GenS19{F0: 0, F1: 4294967295, F2: (^(v133))}
                                    _ = v179
                                }
                                v4 = append(v4, 127)
                                if (0.1 > 0.1) {
                                    v2[(fi164 & 908)] = false
                                    v4 = append(v4, (fi161 / ((fi161) | 1)))
                                    v180 := GenS21{F0: ((-2.25 + float64(0.0)) != ((0.1 - 0.0) * (3.14159 / 1.0))), F1: 73}
                                    _ = v180
                                }
                                _ = fi164
                            }
                            for fi181 := 0; fi181 < 4; fi181++ {
                                v1 -= v3
                                if ((!false) && true) {
                                    v182 := v2[v33]
                                    var v183 float32 = (1e10 / float32(((0.1 + 1.5) + 1.5)))
                                    _ = v182
                                    _ = v183
                                } else {
                                    v184 := GenS1{F0: 468, F1: "hello", F2: 2}
                                    var v185 int64 = 9223372036854775807
                                    _ = v184
                                    _ = v185
                                }
                                switch (fi181 / ((127) | 1)) {
                                case -1:
                                    if len(v4) > 0 {
                                        v4[0] = 2147483647
                                    }
                                case 0:
                                    switch fi181 {
                                    case 127:
                                        v186 := GenS22{F0: int8(v3), F1: (^(fi161)), F2: 4294967295}
                                        _ = v186
                                    case -9223372036854775808:
                                        var v187 int64 = 1
                                        _ = v187
                                    case 0:
                                        v4 = append(v4, v33)
                                    default:
                                        v33 = (fi181 >> ((1) & 63))
                                        switch "a" {
                                        case "":
                                            var v188 int32 = 2147483647
                                            _ = v188
                                        case "foo bar":
                                            if false {
                                                delete(v2, 4294967295)
                                            }
                                        default:
                                            var v189 uint32 = 857
                                            _ = v189
                                        }
                                        v190 := GenS3{F0: "a", F1: 255}
                                        _ = v190
                                    }
                                }
                                fmt.Println(v133, fi161, v3)
                                _ = fi181
                            }
                            _ = fi161
                        }
                    }
                } else {
                    v191 := GenS17{F0: -602, F1: (v3 | v3), F2: 213}
                    v192 := v2[2147483647]
                    _ = v191
                    _ = v192
                }
                var v193 uint16 = 255
                _ = v193
            } else {
                var v194 int32 = -254
                _ = v194
            }
            _ = v133
        }
    }
    v195 := GenS23{F0: 0, F1: 6}
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v33:", v33)
    fmt.Println("v34:", v34)
    fmt.Println("slicelen v4:", len(v4), v4)
    _ = v1
    _ = v128
    _ = v195
    _ = v2
    _ = v3
    _ = v33
    _ = v34
    _ = v4
}
