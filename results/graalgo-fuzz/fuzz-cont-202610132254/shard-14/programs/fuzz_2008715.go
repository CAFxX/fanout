package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
type GenS2 struct { F0 int64; F1 uint32 }
type GenS3 struct { F0 uint8 }
type GenS4 struct { F0 uint64; F1 int16; F2 int64 }
type GenS5 struct { F0 uint8 }
type GenS6 struct { F0 bool }
type GenS7 struct { F0 int16; F1 bool }
type GenS8 struct { F0 int64 }
type GenS9 struct { F0 uint16; F1 bool; F2 uint16 }
type GenS10 struct { F0 bool }
type GenS11 struct { F0 bool; F1 uint }
type GenS12 struct { F0 uint64; F1 uint32 }
type GenS13 struct { F0 int64; F1 bool }
type GenS14 struct { F0 int8 }
type GenS15 struct { F0 uint64; F1 int }
type GenS16 struct { F0 uint8; F1 int32; F2 uint }
type GenS17 struct { F0 bool }
func genFn0(p0 uint64) int8 {
    v5 := []int{}
    var v6 bool = (1 == p0)
    switch ("a" + "hello") {
    case "":
        v6 = (65535 >= 2)
    case "\n\t":
        for fi7 := 0; fi7 < 10; fi7++ {
            var v8 int16 = -32768
            _ = fi7
            _ = v8
        }
    }
    fmt.Println(v6, p0)
    v6 = v6
    _ = p0
    _ = v5
    _ = v6
    return -128
}
func genFn1(p0 uint64, p1 bool) string {
    p1 = p1
    v9 := []int{}
    var v10 uint8 = 75
    p0 = p0
    _ = p0
    _ = p1
    _ = v10
    _ = v9
    return "a"
}

func main() {
    var v1 int16 = 32767
    var v2 float32 = (((0.1 + 0.0) / 1.0) + (3.14159 + -2.25))
    var v3 uint32 = 110
    var v4 int32 = -2147483648
    v11 := GenS1{F0: 1}
    for fi12 := 0; fi12 < 13; fi12++ {
        var v13 uint32 = v3
        var v14 int32 = -126
        if false {
            v3 += 803
            v1 = v1
            if true {
                switch "\n\t" {
                case "":
                    for fi15 := 0; fi15 < 1; fi15++ {
                        v14 += 2147483647
                        v13 -= (v13 << ((98) & 63))
                        v16 := GenS1{F0: genFn0(1)}
                        _ = fi15
                        _ = v16
                    }
                }
                for fi17 := 0; fi17 < 5; fi17++ {
                    fmt.Println(v4, v2)
                    v4 += 2147483647
                    _ = fi17
                }
            }
        } else {
            var v18 int32 = -437
            switch ("a" + "\n\t") {
            case "foo bar":
                for fi19 := 0; fi19 < 19; fi19++ {
                    v20 := []int{(fi12 * fi19), fi19, -334, (fi12 | fi19)}
                    if true {
                        if (!false) {
                            if (3.14159 <= ((((float64(1e10) + (0.1 - float64((v2 + 1e10)))) + 1e10) / 1.0) - (1.5 / ((0.0 - -2.25) - 1.5)))) {
                                switch ("foo bar" + "a") {
                                case "a":
                                    if false {
                                        var v21 int = fi19
                                        var v22 float64 = (1e10 + 1.5)
                                        var v23 int = (fi19 >> ((18446744073709551615) & 63))
                                        _ = v21
                                        _ = v22
                                        _ = v23
                                    } else {
                                        _ = len(v20)
                                        v24 := GenS2{F0: -1, F1: 2147483648}
                                        v25 := GenS3{F0: 42}
                                        _ = v24
                                        _ = v25
                                    }
                                case "hello":
                                    fmt.Println(fi19)
                                }
                                if ((2 != 134) && (((1.5 + 3.14159) * -2.25) < (float64(v2) / float64((0.0 * v2))))) {
                                    fmt.Println(v18, fi19, v4)
                                    var v26 int = 42
                                    _ = v26
                                } else {
                                    v27 := GenS1{F0: genFn0(977)}
                                    fmt.Println(v3, fi19, v4)
                                    _ = v27
                                }
                                if false {
                                    if true {
                                        fmt.Println(v14, v4, v18)
                                        v4 += 42
                                    } else {
                                        if (((((v2 - (-2.25 * (0.1 - (float32(((3.14159 + -2.25) - (((3.14159 + 0.0) + 1.5) * (-2.25 * 0.1)))) * 0.0)))) - (v2 / (0.1 - (v2 * v2)))) * v2) <= (1e10 - (float32(0.1) * v2))) || ("" == "\n\t")) {
                                            if ((v1 ^ 933) > 42) {
                                                v13 -= (v3 / ((v3) | 1))
                                            } else {
                                                fmt.Println(v14)
                                                v28 := GenS4{F0: 0, F1: 0, F2: 9223372036854775807}
                                                if true {
                                                    fmt.Println(v2, v4)
                                                    fmt.Println(v13)
                                                    var v29 uint32 = 633
                                                    _ = v29
                                                } else {
                                                    if false {
                                                        v18 += v18
                                                        var v30 string = "hello"
                                                        var v31 uint64 = 1
                                                        _ = v30
                                                        _ = v31
                                                    }
                                                    var v32 int32 = (v18 * (v4 ^ v14))
                                                    fmt.Println(v18, fi19, v3)
                                                    _ = v32
                                                }
                                                if len(v20) >= 2 {
                                                    v20 = v20[1:len(v20)]
                                                }
                                                _ = v28
                                            }
                                            var v33 int8 = -1
                                            if ((1.5 == 1e10) && true) {
                                                var v34 uint32 = (^(v3))
                                                if false {
                                                    switch (("a" + (("a" + "a") + "hello")) + "\n\t") {
                                                    case "\n\t":
                                                        var v35 uint64 = 2
                                                        if true {
                                                            switch ((("foo bar" + ("" + "hello")) + "foo bar") + "a") {
                                                            case "foo bar":
                                                                var v36 bool = true
                                                                _ = v36
                                                            case "a":
                                                                var v37 uint64 = v35
                                                                _ = v37
                                                            default:
                                                                switch -757 {
                                                                case -110:
                                                                    var v38 float32 = v2
                                                                    _ = v38
                                                                }
                                                            }
                                                            var v39 uint8 = 200
                                                            _ = v39
                                                        } else {
                                                            _ = len(v20)
                                                            var v40 int8 = (v33 - (v33 << ((1) & 63)))
                                                            _ = v40
                                                        }
                                                        _ = v35
                                                    }
                                                    switch ("" + (("a" + (("hello" + ("hello" + "foo bar")) + "a")) + "\n\t")) {
                                                    case "hello":
                                                        v41 := GenS5{F0: 1}
                                                        _ = v41
                                                    case "a":
                                                        switch fi12 {
                                                        case 42:
                                                            if true {
                                                                if true {
                                                                    if false {
                                                                        v4 = v4
                                                                        if (int(v1) > fi19) {
                                                                            v42 := GenS2{F0: int64(v1), F1: (^(v34))}
                                                                            var v43 int8 = v33
                                                                            v44 := GenS3{F0: 1}
                                                                            _ = v42
                                                                            _ = v43
                                                                            _ = v44
                                                                        }
                                                                        v3 = (v13 * 975)
                                                                    }
                                                                    var v45 uint32 = 377
                                                                    if len(v20) >= 2 {
                                                                        v20 = v20[1:len(v20)]
                                                                    }
                                                                    _ = v45
                                                                }
                                                                if len(v20) >= 2 {
                                                                    v20 = v20[1:len(v20)]
                                                                }
                                                            } else {
                                                                if (-128 >= v33) {
                                                                    v46 := GenS6{F0: false}
                                                                    v47 := GenS4{F0: 1, F1: v1, F2: -9223372036854775808}
                                                                    v48 := GenS7{F0: (v1 / (((v1 | (^(v1)))) | 1)), F1: (1 <= 4294967295)}
                                                                    _ = v46
                                                                    _ = v47
                                                                    _ = v48
                                                                }
                                                                var v49 int = fi19
                                                                switch v49 {
                                                                case 32768:
                                                                    v20 = append(v20, fi19)
                                                                }
                                                                _ = v49
                                                            }
                                                        case 32768:
                                                            switch "a" {
                                                            case "\n\t":
                                                                v4 += (v4 >> ((65535) & 63))
                                                            case "foo bar":
                                                                if len(v20) > 0 {
                                                                    v20[0] = 274
                                                                }
                                                            case "hello":
                                                                v4 += (v14 >> ((58) & 63))
                                                                if len(v20) >= 2 {
                                                                    v20 = v20[1:len(v20)]
                                                                }
                                                                var v50 uint32 = 217
                                                                _ = v50
                                                            default:
                                                                if (3.14159 <= float32(1e10)) {
                                                                    if true {
                                                                        if true {
                                                                            var v51 int = (fi12 % ((-2147483648) | 1))
                                                                            _ = v51
                                                                        } else {
                                                                            var v52 int8 = genFn0(4294967295)
                                                                            fmt.Println(v3)
                                                                            if (!(65535 < 255)) {
                                                                                v53 := GenS7{F0: v1, F1: (true || false)}
                                                                                _ = v53
                                                                            } else {
                                                                                v4 = 79
                                                                            }
                                                                            _ = v52
                                                                        }
                                                                        v20 = append(v20, (-(fi12)))
                                                                    } else {
                                                                        if (!false) {
                                                                            switch "hello" {
                                                                            case "hello":
                                                                                var v54 uint16 = 0
                                                                                _ = v54
                                                                            case "\n\t":
                                                                                if ((-2.25 - -2.25) != 1.5) {
                                                                                    v55 := GenS6{F0: (!false)}
                                                                                    _ = v55
                                                                                } else {
                                                                                    v20 = append(v20, int(v1))
                                                                                }
                                                                            default:
                                                                                v2 -= ((v2 / 3.14159) / v2)
                                                                            }
                                                                        } else {
                                                                            if true {
                                                                                var v56 uint64 = 496
                                                                                switch 1 {
                                                                                case 32768:
                                                                                    _ = len(v20)
                                                                                case 127:
                                                                                    v57 := GenS8{F0: -1}
                                                                                    _ = v57
                                                                                case -2147483648:
                                                                                    v58 := GenS1{F0: int8(v1)}
                                                                                    _ = v58
                                                                                }
                                                                                if true {
                                                                                    var v59 bool = false
                                                                                    v13 = (v13 % ((2) | 1))
                                                                                    _ = v59
                                                                                } else {
                                                                                    v20 = append(v20, 127)
                                                                                    var v60 uint8 = 42
                                                                                    var v61 bool = (1e10 != ((1e10 * ((3.14159 / 1e10) / 3.14159)) + 3.14159))
                                                                                    _ = v60
                                                                                    _ = v61
                                                                                }
                                                                                _ = v56
                                                                            }
                                                                            var v62 int32 = (-(v4))
                                                                            v18 += v4
                                                                            _ = v62
                                                                        }
                                                                        v2 = v2
                                                                    }
                                                                    v63 := GenS1{F0: (v33 / (((v33 ^ 36)) | 1))}
                                                                    _ = v63
                                                                } else {
                                                                    if false {
                                                                        var v64 int8 = (v33 * v33)
                                                                        if true {
                                                                            fmt.Println(v4)
                                                                            if false {
                                                                                if false {
                                                                                    switch 127 {
                                                                                    case 32767:
                                                                                        var v65 int32 = v4
                                                                                        _ = v65
                                                                                    }
                                                                                    if len(v20) >= 2 {
                                                                                        v20 = v20[1:len(v20)]
                                                                                    }
                                                                                }
                                                                                var v66 bool = true
                                                                                fmt.Println(fi12, v18)
                                                                                _ = v66
                                                                            } else {
                                                                                var v67 uint8 = 93
                                                                                switch "\n\t" {
                                                                                case "hello":
                                                                                    if len(v20) >= 2 {
                                                                                        v20 = v20[1:len(v20)]
                                                                                    }
                                                                                case "foo bar":
                                                                                    fmt.Println(v64, fi12)
                                                                                case "\n\t":
                                                                                    v33 = v64
                                                                                }
                                                                                fmt.Println(v13)
                                                                                _ = v67
                                                                            }
                                                                        } else {
                                                                            var v68 string = ""
                                                                            v69 := GenS9{F0: 255, F1: true, F2: 2}
                                                                            fmt.Println(v33)
                                                                            _ = v68
                                                                            _ = v69
                                                                        }
                                                                        _ = v64
                                                                    } else {
                                                                        if false {
                                                                            v2 = (v2 / 1.0)
                                                                        }
                                                                    }
                                                                    var v70 uint = 42
                                                                    if (-32768 == v1) {
                                                                        v71 := GenS3{F0: 255}
                                                                        _ = v71
                                                                    }
                                                                    var v72 string = (("\n\t" + "\n\t") + (("" + "a") + "hello"))
                                                                    _ = v70
                                                                    _ = v72
                                                                }
                                                            }
                                                        case 32767:
                                                            v34 = (v34 % ((v3) | 1))
                                                        default:
                                                            var v73 int = fi12
                                                            _ = v73
                                                        }
                                                    case "foo bar":
                                                        v14 = v4
                                                        if (((!false) || (727 < -1)) || ((v2 / v2) < 1.5)) {
                                                            var v74 uint = 2
                                                            if len(v20) > 0 {
                                                                v20[0] = 1
                                                            }
                                                            if (v14 < (-(v4))) {
                                                                if (true && true) {
                                                                    v13 += v13
                                                                    v18 -= v4
                                                                    switch (("" + "foo bar") + "hello") {
                                                                    case "a":
                                                                        v75 := GenS10{F0: true}
                                                                        _ = v75
                                                                    case "\n\t":
                                                                        var v76 uint64 = 4294967295
                                                                        _ = v76
                                                                    default:
                                                                        var v77 uint32 = v13
                                                                        _ = v77
                                                                    }
                                                                } else {
                                                                    fmt.Println(v18, v4)
                                                                }
                                                                if false {
                                                                    v34 = (v34 / (((v13 >> ((1) & 63))) | 1))
                                                                    var v78 uint32 = 4294967295
                                                                    if len(v20) >= 2 {
                                                                        v20 = v20[1:len(v20)]
                                                                    }
                                                                    _ = v78
                                                                }
                                                                fmt.Println(v18)
                                                            }
                                                            _ = v74
                                                        } else {
                                                            switch "" {
                                                            case "hello":
                                                                v79 := GenS10{F0: true}
                                                                _ = v79
                                                            default:
                                                                v80 := GenS11{F0: false, F1: 255}
                                                                _ = v80
                                                            }
                                                            var v81 int8 = -53
                                                            _ = v81
                                                        }
                                                    }
                                                } else {
                                                    var v82 uint = 0
                                                    _ = v82
                                                }
                                                _ = v34
                                            }
                                            _ = v33
                                        }
                                        v1 = v1
                                        v83 := GenS12{F0: 182, F1: 65535}
                                        if len(v20) > 0 {
                                            v20[0] = (-(fi19))
                                        }
                                        _ = v83
                                    }
                                    if true {
                                        v84 := GenS11{F0: false, F1: 47}
                                        var v85 int8 = genFn0(4294967295)
                                        _ = v84
                                        _ = v85
                                    }
                                } else {
                                    switch ((("a" + "") + "hello") + ("" + (("hello" + "hello") + ""))) {
                                    case "a":
                                        if ((0 != 1) && false) {
                                            v86 := GenS13{F0: -347, F1: (v4 == (v18 >> ((4294967295) & 63)))}
                                            _ = v86
                                        } else {
                                            fmt.Println(v3, v14, v2)
                                            if false {
                                                if (true || true) {
                                                    if (727 < 255) {
                                                        if true {
                                                            v87 := GenS14{F0: 79}
                                                            v4 -= (v18 << ((2) & 63))
                                                            if len(v20) >= 2 {
                                                                v20 = v20[1:len(v20)]
                                                            }
                                                            _ = v87
                                                        }
                                                    } else {
                                                        if false {
                                                            var v88 bool = true
                                                            if (255 >= 128) {
                                                                switch "" {
                                                                case "a":
                                                                    var v89 uint8 = 42
                                                                    _ = v89
                                                                default:
                                                                    var v90 uint8 = 23
                                                                    _ = v90
                                                                }
                                                            }
                                                            v91 := GenS14{F0: -52}
                                                            _ = v88
                                                            _ = v91
                                                        } else {
                                                            v3 -= v13
                                                            v14 += (v18 | (v14 * 2147483647))
                                                            fmt.Println(v13, fi12)
                                                        }
                                                        if len(v20) > 0 {
                                                            v20[0] = (fi19 * (-(fi12)))
                                                        }
                                                    }
                                                }
                                            } else {
                                                var v92 int64 = 9223372036854775807
                                                var v93 int64 = 0
                                                var v94 int8 = -17
                                                _ = v92
                                                _ = v93
                                                _ = v94
                                            }
                                            var v95 int8 = -128
                                            _ = v95
                                        }
                                    case "foo bar":
                                        switch "" {
                                        case "hello":
                                            fmt.Println(v4)
                                        case "foo bar":
                                            v1 += (^(v1))
                                        case "\n\t":
                                            v96 := GenS3{F0: 2}
                                            _ = v96
                                        }
                                    }
                                    fmt.Println(v14)
                                    v18 += v14
                                    v3 -= v13
                                    var v97 uint8 = 127
                                    _ = v97
                                }
                                v1 += int16(v4)
                            } else {
                                switch ("" + ("foo bar" + "foo bar")) {
                                case "a":
                                    var v98 float64 = 0.1
                                    _ = v98
                                default:
                                    if true {
                                        switch "hello" {
                                        case "a":
                                            var v99 bool = false
                                            _ = v99
                                        case "hello":
                                            var v100 int8 = 0
                                            _ = v100
                                        }
                                        v20 = append(v20, (fi12 - (^(fi12))))
                                    }
                                    switch (fi19 ^ -932) {
                                    case 127:
                                        v3 = 862
                                    }
                                    v101 := GenS10{F0: (1 <= 2)}
                                    _ = v101
                                }
                                var v102 int8 = 100
                                fmt.Println(fi12, v3)
                                _ = v102
                            }
                        }
                        var v103 uint16 = 0
                        v14 = v18
                        _ = v103
                    } else {
                        _ = len(v20)
                    }
                    v1 += v1
                    _ = fi19
                    _ = v20
                }
            case "":
                if (!(42 < 2)) {
                    fmt.Println(v4)
                    v104 := GenS13{F0: 0, F1: (65535 != 1)}
                    _ = v104
                }
            case "a":
                var v105 int8 = int8(v1)
                _ = v105
            default:
                fmt.Println(v4, v18)
            }
            for fi106 := 0; fi106 < 9; fi106++ {
                v107 := GenS9{F0: 65535, F1: ("hello" >= "foo bar"), F2: 1}
                v108 := []int{(fi106 + fi106)}
                var v109 uint = 609
                _ = fi106
                _ = v107
                _ = v108
                _ = v109
            }
            _ = v18
        }
        _ = fi12
        _ = v13
        _ = v14
    }
    var v110 bool = true
    if true {
        switch 646 {
        case 9223372036854775807:
            v111 := GenS11{F0: (v110 || ((v110 || v110) || (v3 >= (v3 ^ v3)))), F1: uint(v3)}
            _ = v111
        case 4294967295:
            v4 -= v4
        case 127:
            if false {
                v112 := []int{2147483647, 740}
                _ = v112
            }
        default:
            for fi113 := 0; fi113 < 19; fi113++ {
                v114 := GenS3{F0: 0}
                _ = fi113
                _ = v114
            }
        }
        var v115 int = 42
        _ = v115
    }
    v3 -= v3
    for fi116 := 0; fi116 < 13; fi116++ {
        fmt.Println(v2, v1)
        if false {
            v4 -= -181
            if v110 {
                v3 += 1
                v117 := []int{fi116, fi116}
                _ = v117
            } else {
                var v118 uint = 4294967295
                _ = v118
            }
            fmt.Println(v110, v3)
        }
        v110 = ((18446744073709551615 == 159) && true)
        _ = fi116
    }
    v110 = false
    for fi119 := 0; fi119 < 2; fi119++ {
        if (18446744073709551615 < 1) {
            var v120 int16 = 823
            _ = v120
        }
        v121 := []int{(fi119 - fi119)}
        _ = len(v121)
        _ = fi119
        _ = v121
    }
    if (!false) {
        v122 := GenS13{F0: 671, F1: ((870 <= 65535) && ((v4 + -1) <= (v4 / (((v4 - v4)) | 1))))}
        for fi123 := 0; fi123 < 14; fi123++ {
            switch fi123 {
            case -883:
                for fi124 := 0; fi124 < 1; fi124++ {
                    var v125 uint32 = 958
                    v126 := GenS10{F0: true}
                    if (v1 < (v1 ^ (v1 | (v1 - (v1 | v1))))) {
                        switch 32767 {
                        case 42:
                            v1 = v1
                        default:
                            var v127 int16 = (v1 * 660)
                            _ = v127
                        }
                        var v128 int64 = 1
                        switch "foo bar" {
                        case "foo bar":
                            var v129 int8 = 7
                            _ = v129
                        case "":
                            v110 = (-806 < v4)
                        }
                        fmt.Println(v2)
                        switch (("a" + "foo bar") + "foo bar") {
                        case "a":
                            var v130 uint8 = 127
                            _ = v130
                        case "hello":
                            v131 := GenS15{F0: 1, F1: fi123}
                            _ = v131
                        }
                        _ = v128
                    }
                    _ = fi124
                    _ = v125
                    _ = v126
                }
            case -860:
                var v132 uint32 = (v3 % (((^(v3))) | 1))
                _ = v132
            case 32768:
                v133 := GenS1{F0: 1}
                _ = v133
            default:
                for fi134 := 0; fi134 < 1; fi134++ {
                    v3 -= (v3 ^ 65535)
                    v1 += v1
                    var v135 bool = v110
                    var v136 uint = 431
                    v136 = (v136 ^ v136)
                    _ = fi134
                    _ = v135
                    _ = v136
                }
            }
            v137 := GenS16{F0: 42, F1: (v4 | v4), F2: 65535}
            _ = fi123
            _ = v137
        }
        _ = v122
    } else {
        v138 := GenS17{F0: v110}
        _ = v138
    }
    v139 := GenS6{F0: v110}
    v1 = -1
    fmt.Println("v1:", v1)
    fmt.Println("v110:", v110)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v11
    _ = v110
    _ = v139
    _ = v2
    _ = v3
    _ = v4
}
