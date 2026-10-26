package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 uint8 }
type GenS2 struct { F0 bool }
type GenS3 struct { F0 int32; F1 int8 }
type GenS4 struct { F0 bool; F1 uint16; F2 uint64 }
type GenS5 struct { F0 int32 }
type GenS6 struct { F0 int64 }
type GenS7 struct { F0 uint32 }
type GenS8 struct { F0 int32; F1 int32; F2 bool }
type GenS9 struct { F0 int8; F1 uint16 }
type GenS10 struct { F0 int16; F1 int; F2 string }
type GenS11 struct { F0 int8 }
type GenS12 struct { F0 int32; F1 uint }
type GenS13 struct { F0 int8 }
type GenS14 struct { F0 uint64 }
type GenS15 struct { F0 int16; F1 bool }
type GenS16 struct { F0 uint8; F1 int16 }
type GenS17 struct { F0 int64; F1 string; F2 int8 }
type GenS18 struct { F0 int8 }
type GenS19 struct { F0 uint16; F1 int; F2 int16 }
type GenS20 struct { F0 int32; F1 int32 }
type GenS21 struct { F0 int64 }
func genFn0() uint {
    var v4 int32 = 2147483647
    v5 := GenS1{F0: -128, F1: 2}
    _ = v4
    _ = v5
    return 42
}
func genFn1(p0 uint8) uint8 {
    p0 -= p0
    for fi6 := 0; fi6 < 18; fi6++ {
        var v7 int = fi6
        p0 += p0
        _ = fi6
        _ = v7
    }
    if (!(p0 > (p0 / ((p0) | 1)))) {
        for fi8 := 0; fi8 < 2; fi8++ {
            var v9 int8 = 42
            v9 += v9
            v10 := []int{}
            _ = fi8
            _ = v10
            _ = v9
        }
        if false {
            v11 := GenS1{F0: -68, F1: (^(p0))}
            for fi12 := 0; fi12 < 4; fi12++ {
                for fi13 := 0; fi13 < 8; fi13++ {
                    v14 := []int{712, fi12, (^(fi12)), 0}
                    _ = fi13
                    _ = v14
                }
                p0 += (p0 >> ((4294967295) & 63))
                p0 += (p0 - p0)
                _ = fi12
            }
            _ = v11
        } else {
            if (1.5 <= 1.5) {
                switch ("foo bar" + "a") {
                case "\n\t":
                    p0 = p0
                case "hello":
                    fmt.Println(p0)
                case "":
                    var v15 float64 = (1.5 + 3.14159)
                    _ = v15
                }
                p0 += p0
                p0 -= (^(p0))
            } else {
                for fi16 := 0; fi16 < 5; fi16++ {
                    for fi17 := 0; fi17 < 1; fi17++ {
                        v18 := GenS2{F0: false}
                        _ = fi17
                        _ = v18
                    }
                    var v19 uint16 = 65535
                    if (!false) {
                        p0 = (p0 << ((genFn0()) & 63))
                        v20 := []int{fi16}
                        if true {
                            v21 := GenS3{F0: 1, F1: 32}
                            v19 = (v19 % ((v19) | 1))
                            _ = v21
                        } else {
                            fmt.Println(fi16)
                            for fi22 := 0; fi22 < 16; fi22++ {
                                if false {
                                    switch fi22 {
                                    case 4294967295:
                                        switch "hello" {
                                        case "hello":
                                            if (-1 == 2147483647) {
                                                v20 = append(v20, fi16)
                                                v20 = append(v20, (fi22 - fi16))
                                                v19 = uint16(p0)
                                                p0 = p0
                                            } else {
                                                fmt.Println(p0, v19, v19)
                                            }
                                        case "":
                                            switch (^(fi16)) {
                                            case 0:
                                                fmt.Println(v19, p0, v19)
                                            case -735:
                                                var v23 int16 = -959
                                                _ = v23
                                            case 9223372036854775807:
                                                v20 = append(v20, fi16)
                                            }
                                            var v24 uint8 = p0
                                            var v25 uint16 = (v19 & (v19 * (v19 >> ((4294967295) & 63))))
                                            var v26 string = ("foo bar" + "\n\t")
                                            if (2 == 756) {
                                                if len(v20) > 0 {
                                                    v20[0] = fi16
                                                }
                                                if len(v20) > 0 {
                                                    v20[0] = fi16
                                                }
                                                var v27 float64 = 3.14159
                                                v27 = v27
                                                _ = len(v20)
                                                _ = v27
                                            }
                                            var v28 bool = true
                                            _ = v24
                                            _ = v25
                                            _ = v26
                                            _ = v28
                                        case "\n\t":
                                            fmt.Println(p0, v19)
                                        }
                                    case 2147483647:
                                        p0 += (^(p0))
                                        _ = len(v20)
                                        v19 += v19
                                        p0 -= p0
                                        var v29 float64 = (-2.25 + 1e10)
                                        _ = v29
                                    case 128:
                                        if (false || true) {
                                            var v30 uint32 = uint32(p0)
                                            v31 := GenS3{F0: 1, F1: 127}
                                            _ = v30
                                            _ = v31
                                        }
                                        fmt.Println(fi16)
                                    default:
                                        if len(v20) >= 2 {
                                            v20 = v20[1:len(v20)]
                                        }
                                    }
                                    var v32 string = "\n\t"
                                    v19 += 1
                                    _ = v32
                                }
                                switch "\n\t" {
                                case "a":
                                    v33 := GenS3{F0: 1, F1: 104}
                                    _ = v33
                                case "hello":
                                    v34 := GenS3{F0: 0, F1: -88}
                                    _ = v34
                                default:
                                    var v35 uint64 = 742
                                    _ = v35
                                }
                                _ = fi22
                            }
                            var v36 string = "\n\t"
                            _ = v36
                        }
                        _ = v20
                    }
                    _ = fi16
                    _ = v19
                }
                if (65535 > 1) {
                    var v37 uint = 150
                    var v38 int64 = -719
                    _ = v37
                    _ = v38
                } else {
                    switch 2147483647 {
                    case 330:
                        p0 += (p0 / ((2) | 1))
                    case 127:
                        fmt.Println(p0)
                    case -310:
                        v39 := GenS4{F0: (610 < 2147483647), F1: 0, F2: 664}
                        _ = v39
                    }
                }
            }
            for fi40 := 0; fi40 < 6; fi40++ {
                for fi41 := 0; fi41 < 10; fi41++ {
                    v42 := []int{32768, (fi40 + (fi41 * 1)), fi40, fi40}
                    _ = fi41
                    _ = v42
                }
                _ = fi40
            }
        }
    }
    _ = p0
    return (^(p0))
}

func main() {
    var v1 bool = (false || (0.1 <= -2.25))
    var v2 uint64 = 1
    var v3 uint = 42
    v3 -= v3
    switch "a" {
    case "hello":
        for fi43 := 0; fi43 < 19; fi43++ {
            switch fi43 {
            case 4294967295:
                var v44 float64 = 3.14159
                _ = v44
            case -2147483648:
                switch ("\n\t" + ("" + "a")) {
                case "\n\t":
                    v1 = v1
                }
            case 0:
                for fi45 := 0; fi45 < 8; fi45++ {
                    var v46 bool = true
                    v1 = true
                    var v47 int16 = 32767
                    _ = fi45
                    _ = v46
                    _ = v47
                }
            default:
                v48 := GenS5{F0: -473}
                _ = v48
            }
            v1 = v1
            _ = fi43
        }
    default:
        v49 := GenS4{F0: (189 > 998), F1: 630, F2: 4294967295}
        _ = v49
    }
    if false {
        v2 += v2
    } else {
        v2 -= v2
        for fi50 := 0; fi50 < 18; fi50++ {
            v1 = (int16(v3) != -32768)
            for fi51 := 0; fi51 < 13; fi51++ {
                var v52 uint32 = 4294967295
                if v1 {
                    v53 := []int{(fi51 | (-(fi50))), fi51}
                    v2 = v2
                    if len(v53) > 0 {
                        v53[0] = (fi50 ^ (fi51 - (fi51 ^ (fi50 % ((fi50) | 1)))))
                    }
                    if (1 <= 65535) {
                        var v54 uint16 = uint16(v52)
                        if len(v53) > 0 {
                            v53[0] = 127
                        }
                        var v55 string = ""
                        fmt.Println(v54, v55, fi50)
                        v56 := GenS2{F0: v1}
                        var v57 bool = (3.14159 > 3.14159)
                        if v1 {
                            v2 += (v2 | v2)
                        }
                        v1 = (!v1)
                        var v58 uint16 = v54
                        _ = v54
                        _ = v55
                        _ = v56
                        _ = v57
                        _ = v58
                    }
                    switch ("foo bar" + "a") {
                    case "\n\t":
                        switch -434 {
                        case 9223372036854775807:
                            if (v1 && false) {
                                v59 := GenS4{F0: v1, F1: 1, F2: (v2 - v2)}
                                var v60 uint64 = v2
                                _ = v59
                                _ = v60
                            }
                            var v61 uint16 = 2
                            fmt.Println(fi50)
                            var v62 string = ""
                            _ = v61
                            _ = v62
                        case 128:
                            var v63 float32 = float32(1.5)
                            _ = v63
                        default:
                            var v64 uint64 = 374
                            _ = v64
                        }
                    }
                    _ = v53
                } else {
                    v1 = (true || (!(true && (("foo bar" + "hello") <= (("" + "foo bar") + "hello")))))
                    v65 := GenS2{F0: true}
                    v66 := []int{fi51, fi50, (fi51 / ((fi50) | 1)), fi51}
                    if (fi51 >= fi50) {
                        if (-1 <= 9223372036854775807) {
                            v52 = (v52 | (^(v52)))
                            if len(v66) >= 2 {
                                v66 = v66[1:len(v66)]
                            }
                        }
                        switch fi51 {
                        case 0:
                            var v67 uint8 = 125
                            _ = v67
                        case 9223372036854775807:
                            v1 = v1
                        case 32767:
                            fmt.Println(v52)
                        default:
                            switch "foo bar" {
                            case "foo bar":
                                switch 42 {
                                case 2147483647:
                                    v1 = v1
                                case 128:
                                    fmt.Println(fi50, v2, v52)
                                default:
                                    if v1 {
                                        v2 = v2
                                        if (!(((!false) || ((0.0 - 1e10) != 0.1)) || (32768 > (fi50 - fi51)))) {
                                            var v68 int16 = -741
                                            fmt.Println(fi51)
                                            _ = v68
                                        }
                                    }
                                }
                            case "hello":
                                switch "\n\t" {
                                case "":
                                    switch "\n\t" {
                                    case "\n\t":
                                        v52 -= 2147483647
                                    }
                                }
                                fmt.Println(v3, fi51, v2)
                            }
                        }
                        v66 = append(v66, int(v2))
                    } else {
                        var v69 int64 = -326
                        switch 2147483647 {
                        case -139:
                            switch "" {
                            case "hello":
                                if ((v3 & v3) > v3) {
                                    v70 := GenS4{F0: (!(false || false)), F1: 535, F2: (^(v2))}
                                    if (!v1) {
                                        if v1 {
                                            var v71 uint64 = v2
                                            v3 = (v3 * genFn0())
                                            _ = len(v66)
                                            _ = v71
                                        } else {
                                            var v72 uint64 = (v2 % ((v2) | 1))
                                            var v73 bool = false
                                            _ = v72
                                            _ = v73
                                        }
                                        fmt.Println(v1)
                                    } else {
                                        switch fi51 {
                                        case 1:
                                            v74 := GenS5{F0: -2147483648}
                                            _ = v74
                                        case 32767:
                                            if false {
                                                if len(v66) > 0 {
                                                    v66[0] = fi50
                                                }
                                                v66 = append(v66, fi51)
                                            }
                                        default:
                                            v52 += v52
                                        }
                                        if v1 {
                                            v75 := GenS6{F0: v69}
                                            _ = v75
                                        }
                                        var v76 int32 = 1
                                        v69 -= -569
                                        _ = v76
                                    }
                                    _ = v70
                                }
                            case "foo bar":
                                v77 := GenS2{F0: false}
                                _ = v77
                            case "":
                                v78 := GenS6{F0: v69}
                                _ = v78
                            default:
                                v79 := GenS2{F0: (4294967295 < 42)}
                                _ = v79
                            }
                        case 42:
                            v80 := GenS7{F0: v52}
                            _ = v80
                        case 127:
                            switch "a" {
                            case "a":
                                fmt.Println(v52, v2, fi51)
                            case "hello":
                                if false {
                                    switch 32768 {
                                    case -106:
                                        if (42 >= 403) {
                                            v81 := GenS8{F0: 0, F1: -1, F2: (!true)}
                                            v52 = v52
                                            v82 := GenS9{F0: 1, F1: 65535}
                                            _ = v81
                                            _ = v82
                                        }
                                    default:
                                        fmt.Println(v52, v3)
                                    }
                                } else {
                                    var v83 bool = (0 >= 0)
                                    _ = v83
                                }
                            case "\n\t":
                                if len(v66) >= 2 {
                                    v66 = v66[1:len(v66)]
                                }
                                if (-2.25 < 1e10) {
                                    switch ("hello" + "a") {
                                    case "hello":
                                        v52 -= (v52 | 65535)
                                        var v84 uint16 = uint16(v3)
                                        v3 -= (^(v3))
                                        _ = v84
                                    case "\n\t":
                                        v3 = v3
                                    }
                                } else {
                                    var v85 string = ("hello" + "hello")
                                    _ = v85
                                }
                                v86 := GenS10{F0: 14, F1: (fi50 * 42), F2: ((("a" + "hello") + ("a" + "foo bar")) + "\n\t")}
                                if (v1 && v1) {
                                    if false {
                                        if v1 {
                                            v69 += 9223372036854775807
                                        } else {
                                            if false {
                                                v87 := GenS11{F0: 0}
                                                v52 = 2
                                                var v88 uint64 = (v2 + (^(v2)))
                                                _ = v87
                                                _ = v88
                                            }
                                            if ("" < ("a" + "hello")) {
                                                switch fi50 {
                                                case 2147483647:
                                                    var v89 int64 = v69
                                                    _ = v89
                                                default:
                                                    switch "hello" {
                                                    case "":
                                                        var v90 uint8 = 180
                                                        _ = v90
                                                    }
                                                }
                                                var v91 int64 = (v69 >> ((genFn0()) & 63))
                                                _ = v91
                                            } else {
                                                fmt.Println(v69, v1)
                                                v66 = append(v66, (fi50 * (fi51 | 128)))
                                            }
                                        }
                                    } else {
                                        v1 = ((^(v3)) >= v3)
                                        if true {
                                            v2 -= v2
                                            if (1 > v2) {
                                                fmt.Println(v69, v52, v3)
                                                fmt.Println(v52)
                                                v92 := GenS12{F0: -167, F1: genFn0()}
                                                _ = v92
                                            } else {
                                                var v93 int = -561
                                                _ = v93
                                            }
                                        }
                                        v3 = 65535
                                    }
                                    v94 := GenS13{F0: 42}
                                    _ = v94
                                } else {
                                    v95 := GenS1{F0: -128, F1: 2}
                                    switch "a" {
                                    case "hello":
                                        if false {
                                            var v96 uint32 = 2
                                            switch -465 {
                                            case 9223372036854775807:
                                                v69 += (v69 % ((224) | 1))
                                            case -2147483648:
                                                v3 += 0
                                            case -1:
                                                var v97 int = (fi50 / ((4294967295) | 1))
                                                _ = v97
                                            }
                                            v98 := GenS14{F0: (v2 % (((v2 >> ((genFn0()) & 63))) | 1))}
                                            _ = v96
                                            _ = v98
                                        }
                                    default:
                                        if (!(!false)) {
                                            var v99 bool = v1
                                            _ = v99
                                        } else {
                                            v2 = v2
                                            switch (fi50 / (((fi51 ^ fi51)) | 1)) {
                                            case 9223372036854775807:
                                                if len(v66) > 0 {
                                                    v66[0] = fi50
                                                }
                                            case 32767:
                                                var v100 float64 = 0.1
                                                _ = v100
                                            case -1:
                                                v69 = (v69 / ((v69) | 1))
                                            default:
                                                v101 := GenS4{F0: v1, F1: 1, F2: v2}
                                                _ = len(v66)
                                                _ = v101
                                            }
                                            var v102 int = fi51
                                            _ = v102
                                        }
                                        var v103 int8 = 127
                                        var v104 uint32 = uint32(v3)
                                        v105 := GenS7{F0: (v104 * v104)}
                                        if (!(!true)) {
                                            var v106 int64 = (v69 / (((v69 * (v69 % ((v69) | 1)))) | 1))
                                            switch "foo bar" {
                                            case "\n\t":
                                                _ = len(v66)
                                            case "a":
                                                v104 = (v104 << (((v3 & 1)) & 63))
                                            }
                                            _ = v106
                                        } else {
                                            if false {
                                                var v107 string = "a"
                                                var v108 uint32 = (v52 >> ((1) & 63))
                                                var v109 int8 = (v103 * (v103 + 1))
                                                _ = v107
                                                _ = v108
                                                _ = v109
                                            } else {
                                                var v110 uint16 = 42
                                                v66 = append(v66, fi51)
                                                v104 = v52
                                                if len(v66) > 0 {
                                                    v66[0] = (fi51 * -9223372036854775808)
                                                }
                                                if false {
                                                    if len(v66) > 0 {
                                                        v66[0] = fi51
                                                    }
                                                } else {
                                                    switch ("hello" + "") {
                                                    case "\n\t":
                                                        fmt.Println(fi51, fi51, v110)
                                                    case "hello":
                                                        v3 = genFn0()
                                                    }
                                                }
                                                v104 = 42
                                                switch "a" {
                                                case "\n\t":
                                                    var v111 uint16 = v110
                                                    _ = v111
                                                default:
                                                    v52 = (^(v52))
                                                }
                                                _ = len(v66)
                                                v2 = (^(v2))
                                                var v112 int32 = 0
                                                if len(v66) >= 2 {
                                                    v66 = v66[1:len(v66)]
                                                }
                                                _ = v110
                                                _ = v112
                                            }
                                            fmt.Println(fi50)
                                        }
                                        switch int(v103) {
                                        case 127:
                                            var v113 int8 = (-(v103))
                                            _ = v113
                                        case 128:
                                            if (-2147483648 <= 1) {
                                                fmt.Println(v2, v2, v69)
                                                fmt.Println(v69, fi50, v3)
                                                var v114 uint = 65535
                                                _ = v114
                                            } else {
                                                v104 = (v52 | 2147483647)
                                                var v115 uint = (v3 & (^(v3)))
                                                _ = v115
                                            }
                                        case 2147483647:
                                            _ = len(v66)
                                        }
                                        switch "\n\t" {
                                        case "":
                                            if (169 != 1) {
                                                v116 := GenS15{F0: 527, F1: v1}
                                                v66 = append(v66, fi50)
                                                var v117 float32 = 1.5
                                                v52 += (v52 / (((^(v52))) | 1))
                                                _ = v116
                                                _ = v117
                                            } else {
                                                v118 := GenS9{F0: (v103 - int8(v104)), F1: 222}
                                                switch (("" + "hello") + ("hello" + "hello")) {
                                                case "":
                                                    var v119 bool = (-1 > 1000)
                                                    _ = v119
                                                case "foo bar":
                                                    var v120 uint32 = v104
                                                    _ = v120
                                                case "a":
                                                    if len(v66) >= 2 {
                                                        v66 = v66[1:len(v66)]
                                                    }
                                                    if v1 {
                                                        if len(v66) >= 2 {
                                                            v66 = v66[1:len(v66)]
                                                        }
                                                        var v121 int16 = 32767
                                                        _ = v121
                                                    }
                                                    fmt.Println(v1, v52)
                                                }
                                                if true {
                                                    switch "foo bar" {
                                                    case "a":
                                                        v1 = (!v1)
                                                    case "hello":
                                                        v52 = (v104 * (v52 & (v104 - (^(v52)))))
                                                    case "foo bar":
                                                        v122 := GenS16{F0: 133, F1: 683}
                                                        _ = v122
                                                    }
                                                    var v123 int16 = 42
                                                    _ = v123
                                                }
                                                var v124 float32 = 3.14159
                                                _ = v118
                                                _ = v124
                                            }
                                        case "a":
                                            if (v3 >= (v3 << (((v3 ^ uint(v2))) & 63))) {
                                                fmt.Println(v103, fi50)
                                                v125 := GenS17{F0: (^(v69)), F1: ("hello" + ""), F2: v103}
                                                v52 = 42
                                                _ = v125
                                            }
                                        case "\n\t":
                                            v69 = (v69 % (((v69 - (v69 & v69))) | 1))
                                        default:
                                            v52 += 0
                                        }
                                        _ = v103
                                        _ = v104
                                        _ = v105
                                    }
                                    _ = v95
                                }
                                v3 += (v3 - (v3 % (((v3 - 408)) | 1)))
                                var v126 uint32 = 668
                                _ = v126
                                _ = v86
                            }
                        default:
                            switch 0 {
                            case 127:
                                v1 = (!false)
                            case 2147483647:
                                v127 := GenS16{F0: 127, F1: -1}
                                _ = v127
                            default:
                                var v128 float32 = (3.14159 / 1e10)
                                _ = v128
                            }
                        }
                        var v129 uint = 32
                        v3 += (^(v129))
                        _ = v129
                        _ = v69
                    }
                    var v130 bool = (772 <= -2147483648)
                    _ = v130
                    _ = v65
                    _ = v66
                }
                var v131 float64 = 0.1
                _ = fi51
                _ = v131
                _ = v52
            }
            for fi132 := 0; fi132 < 7; fi132++ {
                var v133 int8 = 0
                if ((fi50 >> ((v3) & 63)) > int(v2)) {
                    v133 += (v133 % (((v133 & (v133 % ((v133) | 1)))) | 1))
                } else {
                    var v134 uint32 = 2147483647
                    _ = v134
                }
                v3 = 18446744073709551615
                _ = fi132
                _ = v133
            }
            _ = fi50
        }
        fmt.Println(v3)
    }
    v2 += v2
    v3 -= (v3 & genFn0())
    if ("\n\t" <= ("foo bar" + ("foo bar" + ("foo bar" + "a")))) {
        v1 = true
        v1 = v1
        if (42 <= 2147483647) {
            for fi135 := 0; fi135 < 9; fi135++ {
                fmt.Println(v3, v1)
                var v136 string = "a"
                _ = fi135
                _ = v136
            }
        } else {
            v137 := GenS10{F0: 1, F1: -2147483648, F2: "a"}
            switch 2147483647 {
            case -1:
                var v138 int16 = 42
                _ = v138
            case 32767:
                v2 = (^(v2))
            case 550:
                switch 42 {
                case 127:
                    v139 := GenS18{F0: -1}
                    _ = v139
                case 32768:
                    v1 = false
                default:
                    switch 127 {
                    case -2147483648:
                        if (v1 || true) {
                            if v1 {
                                v3 = 42
                            }
                            fmt.Println(v2)
                        }
                    case 474:
                        var v140 int = 0
                        _ = v140
                    case -9223372036854775808:
                        for fi141 := 0; fi141 < 4; fi141++ {
                            v1 = ((true || (2147483647 >= -2147483648)) && v1)
                            v1 = (42 <= 2)
                            _ = fi141
                        }
                    }
                }
            }
            var v142 bool = v1
            _ = v137
            _ = v142
        }
    }
    v143 := GenS16{F0: 255, F1: 334}
    for fi144 := 0; fi144 < 5; fi144++ {
        if (!(float64(((1.5 + (1e10 - 1.5)) / 1e10)) > 0.1)) {
            fmt.Println(v3, v2)
        }
        _ = fi144
    }
    for fi145 := 0; fi145 < 19; fi145++ {
        if true {
            v2 += v2
        } else {
            fmt.Println(v3, fi145, v2)
            for fi146 := 0; fi146 < 10; fi146++ {
                fmt.Println(v1)
                var v147 int64 = -9223372036854775808
                _ = fi146
                _ = v147
            }
            v148 := GenS19{F0: 32767, F1: fi145, F2: -484}
            _ = v148
        }
        v149 := GenS20{F0: -1, F1: -2147483648}
        _ = fi145
        _ = v149
    }
    var v150 int = -1
    v151 := GenS21{F0: 9223372036854775807}
    v150 = v150
    fmt.Println(v1, v1, v2)
    var v152 uint32 = uint32(v2)
    var v153 uint16 = 802
    switch 4294967295 {
    case 892:
        switch int(v153) {
        case 42:
            v2 -= (v2 | v2)
        default:
            v154 := GenS16{F0: 42, F1: 32767}
            _ = v154
        }
    case 4294967295:
        v1 = false
    }
    fmt.Println("v1:", v1)
    fmt.Println("v150:", v150)
    fmt.Println("v152:", v152)
    fmt.Println("v153:", v153)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v143
    _ = v150
    _ = v151
    _ = v152
    _ = v153
    _ = v2
    _ = v3
}
