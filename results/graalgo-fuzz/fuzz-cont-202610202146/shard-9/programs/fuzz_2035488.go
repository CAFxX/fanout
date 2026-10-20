package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 int16 }
type GenS2 struct { F0 int; F1 int }
type GenS3 struct { F0 int32; F1 uint8 }
type GenS4 struct { F0 uint64; F1 uint64 }
type GenS5 struct { F0 int64; F1 int32; F2 uint16 }
type GenS6 struct { F0 uint32; F1 uint64 }
type GenS7 struct { F0 uint32; F1 int8; F2 string }
type GenS8 struct { F0 uint16; F1 int64; F2 int16 }
type GenS9 struct { F0 int; F1 uint }
type GenS10 struct { F0 int32; F1 uint }
type GenS11 struct { F0 uint32; F1 bool }
type GenS12 struct { F0 int8; F1 uint64 }
type GenS13 struct { F0 string; F1 int8 }
type GenS14 struct { F0 uint }
type GenS15 struct { F0 uint16; F1 int32 }
type GenS16 struct { F0 string; F1 int }
type GenS17 struct { F0 uint32; F1 uint; F2 int64 }
type GenS18 struct { F0 uint16 }
type GenS19 struct { F0 uint64 }
type GenS20 struct { F0 uint32 }
type GenS21 struct { F0 int8; F1 int; F2 int32 }
type GenS22 struct { F0 int; F1 int32; F2 uint32 }
type GenS23 struct { F0 uint8; F1 int }
func genFn0(p0 int16) uint {
    for fi6 := 0; fi6 < 15; fi6++ {
        fmt.Println(fi6)
        for fi7 := 0; fi7 < 15; fi7++ {
            v8 := []int{(fi7 | (fi6 << ((1) & 63))), -291, fi6, 9223372036854775807}
            var v9 uint32 = 65535
            _ = fi7
            _ = v8
            _ = v9
        }
        if ((!((!false) && (true || false))) && (false && ((float32(-2.25) == 3.14159) && false))) {
            for fi10 := 0; fi10 < 3; fi10++ {
                fmt.Println(fi6, fi6, fi10)
                if (!(true && true)) {
                    if (true || true) {
                        var v11 int64 = 0
                        v12 := GenS1{F0: v11, F1: (p0 - (p0 + 138))}
                        if true {
                            p0 -= p0
                            if true {
                                var v13 int64 = int64(p0)
                                var v14 uint64 = 2
                                fmt.Println(fi10)
                                _ = v13
                                _ = v14
                            }
                        }
                        var v15 uint8 = 42
                        _ = v11
                        _ = v12
                        _ = v15
                    }
                    v16 := GenS1{F0: 9223372036854775807, F1: (p0 << ((42) & 63))}
                    _ = v16
                }
                _ = fi10
            }
            var v17 float64 = -2.25
            v18 := GenS1{F0: -1, F1: p0}
            _ = v17
            _ = v18
        } else {
            var v19 int16 = (^(p0))
            _ = v19
        }
        _ = fi6
    }
    for fi20 := 0; fi20 < 2; fi20++ {
        if false {
            if false {
                switch fi20 {
                case 42:
                    if ((("hello" != "a") || false) && (-788 > -1)) {
                        v21 := []int{fi20}
                        for fi22 := 0; fi22 < 5; fi22++ {
                            _ = len(v21)
                            v23 := GenS1{F0: -1, F1: 236}
                            _ = fi22
                            _ = v23
                        }
                        v21 = append(v21, (fi20 | fi20))
                        _ = v21
                    } else {
                        if true {
                            v24 := []int{-9223372036854775808, (fi20 << ((65535) & 63))}
                            _ = len(v24)
                            _ = v24
                        } else {
                            p0 -= (p0 + p0)
                            var v25 int8 = 0
                            _ = v25
                        }
                        for fi26 := 0; fi26 < 12; fi26++ {
                            p0 = (^(p0))
                            fmt.Println(p0, p0)
                            _ = fi26
                        }
                        v27 := []int{fi20, (fi20 / (((fi20 * fi20)) | 1)), 32767}
                        _ = v27
                    }
                case 32767:
                    v28 := GenS1{F0: 1, F1: 644}
                    _ = v28
                case -770:
                    p0 -= (p0 >> ((65535) & 63))
                }
                var v29 int16 = (p0 * p0)
                _ = v29
            } else {
                var v30 int16 = -32768
                p0 = 32767
                _ = v30
            }
            if true {
                p0 += int16(fi20)
                if false {
                    v31 := []int{}
                    v32 := GenS1{F0: -378, F1: (p0 % ((1) | 1))}
                    if false {
                        if len(v31) >= 2 {
                            v31 = v31[1:len(v31)]
                        }
                        switch ("a" + ("a" + "hello")) {
                        case "\n\t":
                            for fi33 := 0; fi33 < 8; fi33++ {
                                switch ((((("foo bar" + "\n\t") + "a") + ("a" + ("a" + "hello"))) + "a") + "") {
                                case "a":
                                    v34 := GenS1{F0: 204, F1: (p0 << ((22) & 63))}
                                    _ = v34
                                default:
                                    var v35 int16 = (p0 ^ p0)
                                    _ = v35
                                }
                                if false {
                                    var v36 float64 = (float64(-2.25) + 3.14159)
                                    _ = v36
                                } else {
                                    p0 += -654
                                }
                                if ((("a" + "hello") + "a") <= "") {
                                    if false {
                                        p0 += (-(p0))
                                    }
                                } else {
                                    var v37 int16 = 1000
                                    _ = v37
                                }
                                _ = fi33
                            }
                        case "":
                            p0 = p0
                        case "a":
                            switch ("a" + ("foo bar" + "hello")) {
                            case "a":
                                p0 = (-(p0))
                            case "\n\t":
                                p0 -= p0
                            }
                        }
                    } else {
                        var v38 int8 = 0
                        if (-192 > -1) {
                            if (851 == 42) {
                                v38 = (v38 & v38)
                                v38 = 57
                                if len(v31) >= 2 {
                                    v31 = v31[1:len(v31)]
                                }
                            } else {
                                p0 -= (p0 % ((p0) | 1))
                                switch 777 {
                                case -709:
                                    switch (fi20 - fi20) {
                                    case 127:
                                        if (1 >= 0) {
                                            var v39 int16 = 977
                                            if (!true) {
                                                _ = len(v31)
                                                fmt.Println(v39, v39)
                                                if (true && (((v39 << ((65535) & 63)) != int16(v38)) && (!false))) {
                                                    v39 -= (p0 >> ((255) & 63))
                                                }
                                            }
                                            fmt.Println(fi20, p0, v38)
                                            _ = v39
                                        } else {
                                            var v40 int8 = v38
                                            v41 := GenS2{F0: fi20, F1: 4294967295}
                                            _ = v40
                                            _ = v41
                                        }
                                    default:
                                        fmt.Println(p0, fi20)
                                    }
                                case 685:
                                    v42 := GenS1{F0: 9223372036854775807, F1: p0}
                                    _ = v42
                                case -677:
                                    for fi43 := 0; fi43 < 20; fi43++ {
                                        switch (^(fi20)) {
                                        case 128:
                                            v44 := GenS3{F0: 2147483647, F1: 189}
                                            _ = v44
                                        }
                                        v45 := GenS4{F0: 4294967295, F1: 2}
                                        _ = fi43
                                        _ = v45
                                    }
                                default:
                                    if false {
                                        var v46 int8 = (v38 / (((v38 & v38)) | 1))
                                        for fi47 := 0; fi47 < 17; fi47++ {
                                            var v48 uint64 = 93
                                            _ = fi47
                                            _ = v48
                                        }
                                        _ = v46
                                    }
                                }
                                var v49 uint64 = 4294967295
                                _ = v49
                            }
                            var v50 uint16 = 255
                            _ = len(v31)
                            _ = v50
                        } else {
                            switch "hello" {
                            case "hello":
                                var v51 uint = 0
                                _ = v51
                            case "":
                                switch (("\n\t" + "a") + "foo bar") {
                                case "":
                                    var v52 int16 = p0
                                    _ = v52
                                case "\n\t":
                                    fmt.Println(fi20, v38)
                                case "a":
                                    if (((("a" + "hello") + ("foo bar" + ("" + "\n\t"))) != ("" + (("" + "a") + ("foo bar" + "")))) || (255 <= 42)) {
                                        fmt.Println(fi20, fi20)
                                    }
                                }
                            case "a":
                                v38 -= (v38 * 42)
                            }
                        }
                        _ = v38
                    }
                    _ = v31
                    _ = v32
                }
                for fi53 := 0; fi53 < 3; fi53++ {
                    if false {
                        if (763 < 2147483648) {
                            var v54 uint32 = 228
                            v54 = 2
                            switch 4294967295 {
                            case -1:
                                switch "" {
                                case "hello":
                                    p0 = p0
                                case "foo bar":
                                    if (-1 <= p0) {
                                        v55 := []int{127}
                                        p0 -= -10
                                        if false {
                                            var v56 float32 = 0.0
                                            v57 := GenS2{F0: fi20, F1: (fi53 - fi53)}
                                            _ = v56
                                            _ = v57
                                        } else {
                                            fmt.Println(v54)
                                            if true {
                                                if len(v55) > 0 {
                                                    v55[0] = (fi53 - (fi53 / (((fi53 << ((1) & 63))) | 1)))
                                                }
                                                var v58 uint64 = 18446744073709551615
                                                _ = v58
                                            } else {
                                                if ("" >= "") {
                                                    if len(v55) > 0 {
                                                        v55[0] = fi20
                                                    }
                                                } else {
                                                    if len(v55) >= 2 {
                                                        v55 = v55[1:len(v55)]
                                                    }
                                                    if ((!((v54 <= v54) || (!((1.5 / 0.1) >= 1e10)))) || (!true)) {
                                                        p0 = (p0 + -684)
                                                    } else {
                                                        if (((false && (("hello" + "hello") != ("\n\t" + ("foo bar" + "hello")))) && (42 != 0)) || false) {
                                                            p0 += p0
                                                            v59 := GenS5{F0: 42, F1: -491, F2: 873}
                                                            if ((42 == 18446744073709551615) || (-369 != 87)) {
                                                                v60 := GenS5{F0: 9223372036854775807, F1: -53, F2: 278}
                                                                v61 := GenS5{F0: 9223372036854775807, F1: 42, F2: 65535}
                                                                switch (fi53 >> ((uint(v54)) & 63)) {
                                                                case 32768:
                                                                    v62 := GenS6{F0: v54, F1: 260}
                                                                    v54 = (v54 | (v54 - v54))
                                                                    var v63 int64 = -1
                                                                    v64 := GenS7{F0: v54, F1: -128, F2: ("foo bar" + ("a" + "\n\t"))}
                                                                    _ = len(v55)
                                                                    _ = v62
                                                                    _ = v63
                                                                    _ = v64
                                                                case -17:
                                                                    switch fi53 {
                                                                    case 32767:
                                                                        v65 := GenS3{F0: 0, F1: 42}
                                                                        _ = v65
                                                                    }
                                                                case 602:
                                                                    v54 = v54
                                                                default:
                                                                    v66 := GenS4{F0: 4294967295, F1: 0}
                                                                    _ = v66
                                                                }
                                                                _ = v60
                                                                _ = v61
                                                            } else {
                                                                v54 = v54
                                                                if (-79 != 0) {
                                                                    fmt.Println(p0, v54)
                                                                    v67 := GenS8{F0: 42, F1: 42, F2: (^(p0))}
                                                                    p0 += p0
                                                                    _ = v67
                                                                } else {
                                                                    p0 += p0
                                                                    v55 = append(v55, (fi20 >> ((659) & 63)))
                                                                }
                                                                var v68 uint16 = 42
                                                                var v69 float32 = (0.1 * (1e10 * 0.0))
                                                                _ = v68
                                                                _ = v69
                                                            }
                                                            _ = v59
                                                        }
                                                        v70 := GenS4{F0: 0, F1: 4294967295}
                                                        _ = v70
                                                    }
                                                }
                                                v71 := GenS9{F0: (^(fi53)), F1: 94}
                                                v72 := GenS1{F0: -94, F1: 0}
                                                _ = len(v55)
                                                v73 := GenS10{F0: -2147483648, F1: 18446744073709551615}
                                                _ = v71
                                                _ = v72
                                                _ = v73
                                            }
                                        }
                                        _ = v55
                                    } else {
                                        if false {
                                            v54 -= 65535
                                            p0 -= (p0 & 953)
                                            if (18446744073709551615 > 997) {
                                                var v74 int16 = -551
                                                var v75 uint32 = (v54 ^ v54)
                                                _ = v74
                                                _ = v75
                                            }
                                        }
                                    }
                                    v54 -= (^(v54))
                                    if false {
                                        v76 := GenS7{F0: 2147483647, F1: -128, F2: (("hello" + "a") + ("a" + "a"))}
                                        _ = v76
                                    } else {
                                        if (((89 != 1) || (-854 > -515)) && (18446744073709551615 != 2)) {
                                            fmt.Println(fi20, v54, p0)
                                            var v77 float32 = 0.0
                                            var v78 string = ("\n\t" + "")
                                            _ = v77
                                            _ = v78
                                        } else {
                                            fmt.Println(p0, v54)
                                        }
                                        v79 := []int{4294967295, (-(fi53)), fi53}
                                        if len(v79) >= 2 {
                                            v79 = v79[1:len(v79)]
                                        }
                                        v80 := GenS11{F0: 779, F1: false}
                                        v54 -= (v54 % (((^(v54))) | 1))
                                        _ = v79
                                        _ = v80
                                    }
                                }
                            case -2147483648:
                                if (!(603 != 759)) {
                                    switch (-(fi20)) {
                                    case 128:
                                        var v81 string = ((("" + ("\n\t" + "a")) + "\n\t") + ("\n\t" + ("foo bar" + "\n\t")))
                                        _ = v81
                                    default:
                                        var v82 int8 = 1
                                        _ = v82
                                    }
                                }
                            default:
                                fmt.Println(v54, fi53, fi20)
                                if false {
                                    v83 := GenS1{F0: -1, F1: p0}
                                    _ = v83
                                } else {
                                    if true {
                                        v84 := GenS1{F0: -1, F1: (-(p0))}
                                        v54 += v54
                                        var v85 float64 = (3.14159 * ((((1.5 - 3.14159) + -2.25) - float64(1.5)) - (0.1 / 3.14159)))
                                        _ = v84
                                        _ = v85
                                    }
                                    v54 = v54
                                }
                                v54 += (v54 * (v54 >> ((2) & 63)))
                            }
                            _ = v54
                        } else {
                            p0 -= p0
                            v86 := []int{128}
                            v87 := GenS10{F0: 0, F1: 65535}
                            _ = v86
                            _ = v87
                        }
                        if ((((0.1 * -2.25) - (float64(1.5) / 1e10)) > (0.1 * (0.1 + 1.5))) && (((3.14159 / -2.25) - -2.25) != 3.14159)) {
                            if (false && false) {
                                if (42 > 4294967295) {
                                    var v88 float64 = 0.0
                                    var v89 float32 = (0.0 + float32(v88))
                                    _ = v88
                                    _ = v89
                                } else {
                                    v90 := GenS12{F0: 100, F1: 505}
                                    _ = v90
                                }
                                switch "" {
                                case "a":
                                    if ("" == "hello") {
                                        v91 := GenS3{F0: 1, F1: 2}
                                        if (293 <= 0) {
                                            fmt.Println(fi53, fi20, p0)
                                            var v92 uint8 = 1
                                            _ = v92
                                        } else {
                                            v93 := GenS12{F0: -128, F1: 0}
                                            _ = v93
                                        }
                                        fmt.Println(p0, fi20)
                                        _ = v91
                                    }
                                case "foo bar":
                                    var v94 uint64 = 1
                                    _ = v94
                                case "hello":
                                    v95 := []int{(fi53 * fi53)}
                                    _ = v95
                                }
                                var v96 bool = false
                                v97 := []int{fi53, (fi20 % (((fi20 - (fi20 - (fi53 * 0)))) | 1))}
                                _ = v96
                                _ = v97
                            }
                            switch 42 {
                            case 467:
                                v98 := GenS13{F0: "hello", F1: 0}
                                _ = v98
                            default:
                                v99 := GenS12{F0: 42, F1: 135}
                                _ = v99
                            }
                            switch (^(fi53)) {
                            case 128:
                                if true {
                                    if (42 >= 867) {
                                        if false {
                                            p0 = (p0 << ((718) & 63))
                                            var v100 int32 = 1
                                            _ = v100
                                        } else {
                                            if false {
                                                v101 := GenS13{F0: ("hello" + ("hello" + "")), F1: 100}
                                                var v102 int8 = 1
                                                v103 := GenS14{F0: 246}
                                                _ = v101
                                                _ = v102
                                                _ = v103
                                            }
                                            var v104 int16 = 42
                                            _ = v104
                                        }
                                        v105 := GenS15{F0: 651, F1: -1}
                                        if true {
                                            if ((!(!false)) && false) {
                                                var v106 uint16 = 65535
                                                _ = v106
                                            }
                                            switch "" {
                                            case "a":
                                                if (0 <= -9223372036854775808) {
                                                    if false {
                                                        v107 := []int{}
                                                        v108 := GenS16{F0: "", F1: fi20}
                                                        p0 += 1000
                                                        _ = v107
                                                        _ = v108
                                                    }
                                                    var v109 string = "hello"
                                                    _ = v109
                                                }
                                            }
                                            switch "a" {
                                            case "\n\t":
                                                v110 := GenS17{F0: 42, F1: 616, F2: int64(p0)}
                                                _ = v110
                                            case "hello":
                                                p0 = (^(p0))
                                            case "":
                                                v111 := GenS18{F0: 42}
                                                _ = v111
                                            default:
                                                if (!false) {
                                                    switch fi53 {
                                                    case 32767:
                                                        var v112 int8 = 30
                                                        _ = v112
                                                    default:
                                                        p0 = p0
                                                    }
                                                    switch (fi20 - fi20) {
                                                    case -9223372036854775808:
                                                        if true {
                                                            v113 := GenS19{F0: 4294967295}
                                                            v114 := GenS6{F0: 42, F1: 0}
                                                            _ = v113
                                                            _ = v114
                                                        } else {
                                                            var v115 int = (fi53 << ((1) & 63))
                                                            _ = v115
                                                        }
                                                    case -1:
                                                        p0 -= (^(p0))
                                                    default:
                                                        switch "foo bar" {
                                                        case "":
                                                            v116 := GenS18{F0: 0}
                                                            v117 := GenS2{F0: (fi53 >> ((888) & 63)), F1: fi53}
                                                            p0 = -1
                                                            _ = v116
                                                            _ = v117
                                                        case "foo bar":
                                                            fmt.Println(p0, p0, fi53)
                                                        default:
                                                            if false {
                                                                p0 -= p0
                                                                v118 := []int{(fi20 | (^(fi53))), fi20, fi20, fi20}
                                                                if len(v118) >= 2 {
                                                                    v118 = v118[1:len(v118)]
                                                                }
                                                                _ = v118
                                                            }
                                                        }
                                                    }
                                                    v119 := []int{fi53}
                                                    _ = v119
                                                } else {
                                                    p0 = 0
                                                    fmt.Println(p0, fi53, fi20)
                                                    v120 := []int{459, 827}
                                                    var v121 uint16 = 236
                                                    _ = v120
                                                    _ = v121
                                                }
                                            }
                                        } else {
                                            var v122 string = "foo bar"
                                            p0 -= -1
                                            switch ((v122 + v122) + v122) {
                                            case "a":
                                                v123 := []int{}
                                                _ = v123
                                            }
                                            _ = v122
                                        }
                                        _ = v105
                                    }
                                }
                            }
                        }
                    } else {
                        p0 -= (-(p0))
                        v124 := []int{fi53, fi53, (fi20 % ((-9223372036854775808) | 1))}
                        _ = v124
                    }
                    v125 := GenS5{F0: -1, F1: 2147483647, F2: 65535}
                    _ = fi53
                    _ = v125
                }
            } else {
                v126 := GenS20{F0: 1}
                v127 := GenS17{F0: 700, F1: 65535, F2: 245}
                v128 := []int{(fi20 << ((0) & 63)), 1, fi20, (fi20 % (((fi20 >> ((2) & 63))) | 1))}
                _ = v126
                _ = v127
                _ = v128
            }
            for fi129 := 0; fi129 < 1; fi129++ {
                v130 := []int{}
                if (0.0 > 0.1) {
                    if true {
                        p0 = p0
                    }
                    p0 = p0
                }
                _ = fi129
                _ = v130
            }
        }
        fmt.Println(p0, fi20)
        _ = fi20
    }
    var v131 uint = 4294967295
    if false {
        p0 = (p0 * (p0 ^ -1))
        var v132 bool = false
        if (!(false || (308 < 862))) {
            v133 := []int{42}
            v134 := GenS21{F0: 1, F1: 1, F2: 1}
            v132 = ((693 > 42) || (42 > 2147483648))
            _ = v133
            _ = v134
        }
        _ = v132
    } else {
        for fi135 := 0; fi135 < 13; fi135++ {
            var v136 bool = false
            _ = fi135
            _ = v136
        }
        p0 += p0
        var v137 int32 = 1
        _ = v137
    }
    v138 := []int{-1}
    _ = p0
    _ = v131
    _ = v138
    return (v131 << ((v131) & 63))
}

func main() {
    var v1 int16 = -32768
    var v2 uint8 = 195
    v3 := make(map[int]string)
    var v4 uint64 = 2
    var v5 int8 = 52
    v139 := GenS13{F0: "foo bar", F1: (v5 & (-(v5)))}
    v140 := v3[128]
    v4 = 435
    var v141 uint8 = (v2 >> ((1) & 63))
    v142 := GenS22{F0: 9, F1: -491, F2: uint32(v141)}
    v3[128] = v140
    delete(v3, -1)
    if true {
        var v143 int = 9223372036854775807
        _ = v143
    } else {
        if false {
            var v144 float64 = (((-2.25 / 1e10) / (1e10 + (0.0 - float64(1.5)))) / 1.5)
            v144 += v144
            var v145 int16 = 798
            _ = v144
            _ = v145
        } else {
            if (true && false) {
                v141 = 2
                var v146 int = 127
                v140 = v140
                _ = v146
            }
            var v147 float64 = 1e10
            _ = v147
        }
        fmt.Println(v141, v2)
    }
    v4 = (v4 >> ((65535) & 63))
    switch v140 {
    case "foo bar":
        for fi148 := 0; fi148 < 20; fi148++ {
            var v149 float64 = 3.14159
            v141 += 128
            _ = fi148
            _ = v149
        }
    case "\n\t":
        v150 := []int{}
        _ = v150
    default:
        switch v140 {
        case "hello":
            v5 = v5
        default:
            var v151 float32 = 0.1
            _ = v151
        }
    }
    v152 := GenS16{F0: "\n\t", F1: 32768}
    var v153 uint64 = (v4 | v4)
    if (v4 == v4) {
        fmt.Println(v4, v1, v140)
        _ = len(v3)
        var v154 float64 = float64(float32(((3.14159 + -2.25) - 1e10)))
        _ = v154
    } else {
        switch (("a" + v140) + ((v140 + "hello") + (v140 + ((v140 + (("a" + "foo bar") + v140)) + v140)))) {
        case "":
            switch (v140 + ("a" + (v140 + "\n\t"))) {
            case "\n\t":
                v155 := GenS23{F0: (v2 - v141), F1: -9223372036854775808}
                _ = v155
            case "":
                if true {
                    var v156 uint64 = v4
                    _ = v156
                } else {
                    var v157 float64 = (-2.25 * 1.5)
                    if true {
                        v158 := []int{0, 613}
                        var v159 int64 = -1
                        _ = len(v158)
                        _ = v158
                        _ = v159
                    }
                    _ = v157
                }
            default:
                var v160 uint = genFn0(0)
                _ = v160
            }
        case "\n\t":
            fmt.Println(v141, v1, v5)
        default:
            if (uint16(v2) > 0) {
                var v161 float32 = 1.5
                v1 = (v1 << ((4294967295) & 63))
                v162 := GenS3{F0: 481, F1: (v141 << ((42) & 63))}
                _ = v161
                _ = v162
            } else {
                fmt.Println(v140)
                var v163 uint = 18446744073709551615
                _ = v163
            }
        }
    }
    v1 = v1
    fmt.Println("v1:", v1)
    fmt.Println("v140:", v140)
    fmt.Println("v141:", v141)
    fmt.Println("v153:", v153)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v139
    _ = v140
    _ = v141
    _ = v142
    _ = v152
    _ = v153
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
