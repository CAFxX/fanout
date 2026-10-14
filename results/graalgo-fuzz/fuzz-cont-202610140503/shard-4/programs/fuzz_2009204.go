package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
type GenS2 struct { F0 uint16 }
type GenS3 struct { F0 int; F1 uint }
type GenS4 struct { F0 int64 }
type GenS5 struct { F0 int32 }
type GenS6 struct { F0 int64 }
type GenS7 struct { F0 uint16; F1 uint64; F2 uint64 }
type GenS8 struct { F0 string; F1 int8; F2 int8 }
type GenS9 struct { F0 int16; F1 int8 }
type GenS10 struct { F0 uint8; F1 int }
type GenS11 struct { F0 int8; F1 int32 }
type GenS12 struct { F0 uint; F1 uint32; F2 uint }
type GenS13 struct { F0 string; F1 uint8 }
type GenS14 struct { F0 int64; F1 uint8 }
type GenS15 struct { F0 uint64; F1 int8; F2 int }
func genFn0() int8 {
    v4 := GenS1{F0: ("a" + ("\n\t" + "hello"))}
    var v5 int = 32767
    v5 -= 4294967295
    v5 -= v5
    _ = v4
    _ = v5
    return 100
}

func main() {
    var v1 uint8 = 0
    var v2 int32 = -758
    var v3 uint = 498
    v2 = (-(v2))
    if (!true) {
        v1 += 1
        v6 := GenS2{F0: 294}
        _ = v6
    } else {
        var v7 float32 = 0.0
        v8 := GenS2{F0: 272}
        var v9 uint64 = 1
        _ = v7
        _ = v8
        _ = v9
    }
    for fi10 := 0; fi10 < 17; fi10++ {
        if (!false) {
            var v11 float64 = (1.5 * (-2.25 + 3.14159))
            _ = v11
        }
        v12 := []int{(fi10 % ((fi10) | 1))}
        _ = fi10
        _ = v12
    }
    var v13 int8 = -1
    for fi14 := 0; fi14 < 3; fi14++ {
        var v15 int8 = v13
        _ = fi14
        _ = v15
    }
    v16 := GenS2{F0: 32767}
    fmt.Println(v2)
    fmt.Println(v13)
    if (750 == 154) {
        v17 := []int{42, -9223372036854775808}
        if len(v17) > 0 {
            v17[0] = -563
        }
        _ = v17
    } else {
        if ((false && (!((!true) || (true || true)))) || (false || (v13 >= genFn0()))) {
            v18 := []int{}
            v19 := GenS2{F0: 65535}
            fmt.Println(v1)
            _ = v18
            _ = v19
        } else {
            for fi20 := 0; fi20 < 6; fi20++ {
                for fi21 := 0; fi21 < 3; fi21++ {
                    var v22 bool = true
                    v22 = (65535 <= v3)
                    var v23 uint8 = (^(v1))
                    _ = fi21
                    _ = v22
                    _ = v23
                }
                v24 := GenS2{F0: 2}
                _ = fi20
                _ = v24
            }
        }
        var v25 float32 = float32(-2.25)
        _ = v25
    }
    if (818 >= 0) {
        v3 = 18446744073709551615
        v26 := []int{}
        var v27 string = "a"
        _ = v26
        _ = v27
    }
    switch -713 {
    case -9223372036854775808:
        v28 := []int{32768, 916, -235, 1}
        _ = v28
    case 32768:
        if false {
            for fi29 := 0; fi29 < 16; fi29++ {
                for fi30 := 0; fi30 < 16; fi30++ {
                    fmt.Println(v2, v1, v13)
                    v31 := []int{fi29, fi29}
                    _ = fi30
                    _ = v31
                }
                _ = fi29
            }
        }
    default:
        switch "\n\t" {
        case "foo bar":
            if false {
                v32 := GenS3{F0: -1, F1: v3}
                if (v3 < 243) {
                    v3 += (v3 & v3)
                }
                if false {
                    var v33 uint16 = 65535
                    _ = v33
                }
                _ = v32
            } else {
                var v34 float32 = ((0.1 - (0.1 * 3.14159)) + 1.5)
                if ((true || (("hello" + "\n\t") <= "foo bar")) || true) {
                    v2 = v2
                } else {
                    for fi35 := 0; fi35 < 17; fi35++ {
                        for fi36 := 0; fi36 < 1; fi36++ {
                            var v37 uint32 = 42
                            var v38 int64 = 9223372036854775807
                            v39 := GenS2{F0: 2}
                            _ = fi36
                            _ = v37
                            _ = v38
                            _ = v39
                        }
                        v40 := GenS2{F0: 65535}
                        v2 = 1
                        _ = fi35
                        _ = v40
                    }
                    if false {
                        v41 := GenS2{F0: 884}
                        var v42 int = 127
                        if false {
                            v43 := GenS4{F0: 935}
                            var v44 uint8 = 184
                            v45 := []int{(v42 ^ 2147483647), (v42 >> ((uint(v44)) & 63))}
                            _ = v43
                            _ = v44
                            _ = v45
                        }
                        _ = v41
                        _ = v42
                    }
                }
                _ = v34
            }
        }
    }
    if (false || false) {
        v46 := GenS5{F0: (-(v2))}
        fmt.Println(v2, v3)
        var v47 uint8 = 236
        _ = v46
        _ = v47
    }
    fmt.Println(v1, v1, v3)
    var v48 int8 = 42
    switch (("\n\t" + "\n\t") + "\n\t") {
    case "foo bar":
        if true {
            v3 = v3
            for fi49 := 0; fi49 < 16; fi49++ {
                fmt.Println(v2, v2, v1)
                _ = fi49
            }
            v2 = v2
        }
    case "\n\t":
        v13 = (v48 | (^(v48)))
    case "":
        for fi50 := 0; fi50 < 1; fi50++ {
            var v51 uint32 = 693
            switch fi50 {
            case 1:
                for fi52 := 0; fi52 < 14; fi52++ {
                    v53 := []int{fi50}
                    if (328 > 255) {
                        v53 = append(v53, 1)
                        v54 := GenS3{F0: (fi50 << ((2) & 63)), F1: (v3 << ((v3) & 63))}
                        switch "" {
                        case "":
                            if false {
                                var v55 uint16 = uint16(v3)
                                v56 := GenS3{F0: (fi50 << (((v3 & 42)) & 63)), F1: (v3 / (((v3 >> (((^(v3))) & 63))) | 1))}
                                var v57 uint8 = v1
                                _ = v55
                                _ = v56
                                _ = v57
                            } else {
                                v2 += v2
                                switch (-(fi50)) {
                                case -99:
                                    var v58 uint8 = (v1 | (v1 << ((v3) & 63)))
                                    _ = v58
                                }
                                var v59 int8 = v13
                                if len(v53) > 0 {
                                    v53[0] = (fi52 | (fi52 ^ (fi52 % ((fi52) | 1))))
                                }
                                _ = v59
                            }
                        case "\n\t":
                            var v60 uint8 = (^(v1))
                            v2 = 42
                            if (false && true) {
                                v61 := GenS6{F0: 42}
                                v51 = v51
                                _ = v61
                            }
                            _ = v60
                        default:
                            if len(v53) >= 2 {
                                v53 = v53[1:len(v53)]
                            }
                        }
                        fmt.Println(v51, fi50)
                        var v62 int32 = 958
                        switch int(v51) {
                        case 1:
                            fmt.Println(v62, v13)
                        default:
                            fmt.Println(v2, fi52)
                        }
                        v3 = v3
                        var v63 int8 = genFn0()
                        if (v2 <= v2) {
                            var v64 int16 = 1000
                            v48 += v13
                            switch "hello" {
                            case "a":
                                var v65 uint = (v3 + (v3 / (((v3 ^ v3)) | 1)))
                                _ = v65
                            case "\n\t":
                                var v66 bool = false
                                _ = v66
                            default:
                                var v67 uint8 = v1
                                _ = v67
                            }
                            _ = v64
                        } else {
                            if (!true) {
                                switch (fi50 << ((18446744073709551615) & 63)) {
                                case -1:
                                    if true {
                                        var v68 string = (("\n\t" + "") + (("hello" + "\n\t") + ("" + ("hello" + ("hello" + "")))))
                                        if len(v53) >= 2 {
                                            v53 = v53[1:len(v53)]
                                        }
                                        if false {
                                            if len(v53) > 0 {
                                                v53[0] = fi50
                                            }
                                            switch (v68 + "hello") {
                                            case "":
                                                if (v51 >= 2) {
                                                    if false {
                                                        var v69 float32 = -2.25
                                                        switch (v68 + ((v68 + v68) + ((v68 + ((v68 + v68) + v68)) + (("hello" + v68) + ((("foo bar" + v68) + v68) + ("a" + ("\n\t" + ((v68 + ("hello" + "\n\t")) + v68)))))))) {
                                                        case "\n\t":
                                                            v48 -= (-(v48))
                                                        case "foo bar":
                                                            if len(v53) > 0 {
                                                                v53[0] = fi52
                                                            }
                                                        default:
                                                            var v70 int8 = (v48 | (v13 & (v63 << ((v3) & 63))))
                                                            _ = v70
                                                        }
                                                        v71 := GenS7{F0: 255, F1: 4294967295, F2: 0}
                                                        v1 = 128
                                                        var v72 bool = false
                                                        _ = v69
                                                        _ = v71
                                                        _ = v72
                                                    }
                                                    var v73 uint16 = 677
                                                    fmt.Println(v51, fi50)
                                                    _ = v73
                                                } else {
                                                    if len(v53) > 0 {
                                                        v53[0] = 32768
                                                    }
                                                    var v74 uint = (v3 | (v3 * 42))
                                                    _ = v74
                                                }
                                            case "foo bar":
                                                var v75 string = (((v68 + ("a" + (v68 + ""))) + v68) + "\n\t")
                                                _ = v75
                                            }
                                            var v76 uint = (v3 >> (((^(v3))) & 63))
                                            fmt.Println(fi52, v2, v3)
                                            _ = v76
                                        } else {
                                            v1 += v1
                                        }
                                        _ = v68
                                    } else {
                                        switch (fi52 - 1) {
                                        case 32768:
                                            var v77 uint8 = v1
                                            _ = v77
                                        }
                                        _ = len(v53)
                                    }
                                case 2147483647:
                                    switch -9223372036854775808 {
                                    case 32767:
                                        if true {
                                            var v78 int = fi50
                                            _ = v78
                                        }
                                    default:
                                        var v79 int8 = v63
                                        _ = v79
                                    }
                                case 127:
                                    var v80 int64 = int64(fi50)
                                    _ = v80
                                default:
                                    switch "hello" {
                                    case "":
                                        if false {
                                            var v81 uint = (v3 + 18446744073709551615)
                                            var v82 string = "hello"
                                            _ = v81
                                            _ = v82
                                        } else {
                                            fmt.Println(v51)
                                            var v83 uint32 = 590
                                            switch 464 {
                                            case 474:
                                                _ = len(v53)
                                            case -2147483648:
                                                var v84 uint8 = 1
                                                _ = v84
                                            default:
                                                if (v2 != int32(v1)) {
                                                    v3 += v3
                                                    if true {
                                                        v1 = (v1 * v1)
                                                        v51 = v83
                                                    } else {
                                                        v85 := GenS8{F0: ("a" + "hello"), F1: (v13 & int8(v3)), F2: 100}
                                                        v1 = (v1 - 42)
                                                        switch "hello" {
                                                        case "\n\t":
                                                            v86 := GenS4{F0: -750}
                                                            _ = v86
                                                        default:
                                                            v62 += (v2 + v2)
                                                        }
                                                        if false {
                                                            v2 = 1
                                                            var v87 uint32 = 2147483648
                                                            v1 = v1
                                                            _ = v87
                                                        }
                                                        switch "foo bar" {
                                                        case "\n\t":
                                                            var v88 int = (^(fi50))
                                                            _ = v88
                                                        case "":
                                                            v51 -= v83
                                                        }
                                                        if (!((!false) || (!true))) {
                                                            var v89 int16 = int16(v3)
                                                            var v90 int8 = (v48 + (v13 + (v13 & v48)))
                                                            _ = v89
                                                            _ = v90
                                                        } else {
                                                            if false {
                                                                var v91 uint64 = 4294967295
                                                                if (!(!(!(1e10 > 1e10)))) {
                                                                    if true {
                                                                        _ = len(v53)
                                                                        v13 -= -128
                                                                        if len(v53) > 0 {
                                                                            v53[0] = 1
                                                                        }
                                                                    }
                                                                }
                                                                _ = v91
                                                            }
                                                            var v92 uint16 = 888
                                                            _ = v92
                                                        }
                                                        if false {
                                                            var v93 uint = 4294967295
                                                            if (!true) {
                                                                v63 -= v48
                                                                var v94 int16 = -156
                                                                _ = v94
                                                            }
                                                            _ = v93
                                                        }
                                                        _ = v85
                                                    }
                                                    var v95 float32 = 1e10
                                                    _ = v95
                                                } else {
                                                    if (((!(true || false)) || true) && true) {
                                                        if len(v53) > 0 {
                                                            v53[0] = (-(fi50))
                                                        }
                                                    }
                                                    switch (("a" + "hello") + ("hello" + "hello")) {
                                                    case "\n\t":
                                                        if true {
                                                            _ = len(v53)
                                                        } else {
                                                            v96 := GenS7{F0: 1, F1: 0, F2: 1}
                                                            _ = v96
                                                        }
                                                    case "a":
                                                        var v97 float32 = 0.0
                                                        _ = v97
                                                    }
                                                }
                                            }
                                            _ = v83
                                        }
                                    case "\n\t":
                                        v98 := GenS3{F0: -55, F1: v3}
                                        _ = v98
                                    default:
                                        var v99 float64 = (-2.25 * (((3.14159 * 0.1) + 0.1) - -2.25))
                                        _ = v99
                                    }
                                }
                                v100 := GenS1{F0: ("hello" + ("" + ""))}
                                v13 += v48
                                _ = v100
                            }
                            if false {
                                fmt.Println(fi50)
                                var v101 bool = true
                                var v102 float64 = 1.5
                                _ = v101
                                _ = v102
                            } else {
                                var v103 float32 = float32((1e10 - 0.0))
                                switch ("hello" + (((("foo bar" + "\n\t") + "a") + ("foo bar" + (("\n\t" + "") + ""))) + "a")) {
                                case "a":
                                    switch 0 {
                                    case 4294967295:
                                        if len(v53) >= 2 {
                                            v53 = v53[1:len(v53)]
                                        }
                                    case 32767:
                                        if true {
                                            v3 -= uint(v51)
                                            v104 := GenS9{F0: 32767, F1: (v48 % (((v63 - v13)) | 1))}
                                            _ = v104
                                        }
                                    }
                                case "hello":
                                    if (false && (((false && false) || true) || true)) {
                                        v103 += v103
                                        var v105 uint8 = v1
                                        var v106 int32 = (^(v62))
                                        _ = v105
                                        _ = v106
                                    }
                                }
                                _ = v103
                            }
                            fmt.Println(fi50, v3)
                        }
                        _ = v54
                        _ = v62
                        _ = v63
                    } else {
                        var v107 float32 = (((0.0 * (((-2.25 - 1e10) / 1e10) + 3.14159)) / float32((-2.25 - (float64(3.14159) * 0.1)))) + 0.1)
                        v107 += v107
                        switch fi52 {
                        case 32767:
                            var v108 string = ("hello" + (("foo bar" + "a") + "foo bar"))
                            _ = v108
                        case -458:
                            v109 := GenS10{F0: (v1 * (v1 << ((uint(v51)) & 63))), F1: -2147483648}
                            _ = v109
                        default:
                            if true {
                                var v110 uint16 = 0
                                v111 := GenS11{F0: (^(v48)), F1: (v2 + v2)}
                                fmt.Println(v2)
                                _ = v110
                                _ = v111
                            } else {
                                if len(v53) > 0 {
                                    v53[0] = fi50
                                }
                            }
                        }
                        if (!true) {
                            v53 = append(v53, -2147483648)
                            var v112 int64 = 706
                            if len(v53) > 0 {
                                v53[0] = -2147483648
                            }
                            var v113 float64 = 0.1
                            _ = v112
                            _ = v113
                        }
                        _ = v107
                    }
                    v1 = 0
                    _ = fi52
                    _ = v53
                }
            case 9223372036854775807:
                switch "\n\t" {
                case "hello":
                    v114 := []int{}
                    _ = v114
                case "a":
                    fmt.Println(v2, fi50, v1)
                case "foo bar":
                    v115 := GenS12{F0: (^(v3)), F1: 4294967295, F2: (v3 / (((v3 % ((829) | 1))) | 1))}
                    _ = v115
                default:
                    if true {
                        v116 := []int{fi50}
                        if len(v116) > 0 {
                            v116[0] = (fi50 % (((fi50 % ((fi50) | 1))) | 1))
                        }
                        _ = v116
                    }
                }
            case 32767:
                v13 -= (v48 >> ((v3) & 63))
            }
            v48 = -27
            _ = fi50
            _ = v51
        }
    }
    for fi117 := 0; fi117 < 5; fi117++ {
        if false {
            var v118 int64 = -9223372036854775808
            v119 := []int{2147483647, -2147483648}
            v13 = int8(fi117)
            _ = v118
            _ = v119
        } else {
            v48 = 127
            var v120 uint64 = 201
            _ = v120
        }
        v3 += v3
        if true {
            if (("" + (((("" + "a") + ("hello" + "foo bar")) + ("hello" + "foo bar")) + (("hello" + "") + "a"))) > ("a" + "")) {
                switch (fi117 * (fi117 / ((578) | 1))) {
                case 471:
                    var v121 uint8 = (v1 / (((^(v1))) | 1))
                    _ = v121
                case 4294967295:
                    for fi122 := 0; fi122 < 5; fi122++ {
                        v123 := GenS1{F0: "a"}
                        switch (fi117 & fi117) {
                        case 32768:
                            v124 := GenS13{F0: "foo bar", F1: (v1 ^ 68)}
                            _ = v124
                        }
                        _ = fi122
                        _ = v123
                    }
                case 0:
                    v1 = v1
                default:
                    if false {
                        switch (fi117 ^ 42) {
                        case 32767:
                            for fi125 := 0; fi125 < 1; fi125++ {
                                v3 = v3
                                var v126 bool = (!false)
                                if ((v2 <= 2147483647) || (!(uint32(v1) <= 4294967295))) {
                                    var v127 int8 = genFn0()
                                    _ = v127
                                } else {
                                    v128 := []int{}
                                    _ = v128
                                }
                                _ = fi125
                                _ = v126
                            }
                        default:
                            if false {
                                v129 := []int{fi117, (fi117 + fi117)}
                                for fi130 := 0; fi130 < 13; fi130++ {
                                    if false {
                                        var v131 uint32 = 387
                                        _ = v131
                                    }
                                    if false {
                                        switch "hello" {
                                        case "foo bar":
                                            fmt.Println(v1, v1)
                                        default:
                                            v3 += (v3 / ((4294967295) | 1))
                                        }
                                        v132 := GenS8{F0: "foo bar", F1: (v48 / ((v48) | 1)), F2: (-(v48))}
                                        _ = v132
                                    } else {
                                        v13 -= (-(v13))
                                        var v133 float32 = (((1e10 - (3.14159 + 3.14159)) * (float32(0.0) / 3.14159)) + float32(0.0))
                                        _ = v133
                                    }
                                    var v134 uint8 = 78
                                    _ = fi130
                                    _ = v134
                                }
                                if (v1 > 45) {
                                    _ = len(v129)
                                } else {
                                    var v135 int64 = 9223372036854775807
                                    v136 := GenS14{F0: v135, F1: 42}
                                    v137 := GenS10{F0: 0, F1: (fi117 << (((v3 - v3)) & 63))}
                                    _ = v135
                                    _ = v136
                                    _ = v137
                                }
                                _ = v129
                            } else {
                                var v138 int = fi117
                                for fi139 := 0; fi139 < 18; fi139++ {
                                    v140 := []int{32768, fi139, fi117, (v138 | 32768)}
                                    _ = fi139
                                    _ = v140
                                }
                                _ = v138
                            }
                        }
                    } else {
                        if (v3 > v3) {
                            for fi141 := 0; fi141 < 3; fi141++ {
                                var v142 int32 = -762
                                if ((true && false) || true) {
                                    switch "a" {
                                    case "\n\t":
                                        v142 = (v2 * v142)
                                    case "a":
                                        v143 := []int{(fi117 << ((v3) & 63)), fi117, -9223372036854775808, fi141}
                                        _ = v143
                                    case "":
                                        fmt.Println(v48)
                                    default:
                                        switch ("hello" + ("" + "")) {
                                        case "":
                                            switch fi117 {
                                            case 9223372036854775807:
                                                v1 -= (^(v1))
                                            default:
                                                v3 = 1
                                            }
                                        case "hello":
                                            v144 := GenS14{F0: 646, F1: v1}
                                            _ = v144
                                        case "a":
                                            var v145 float64 = ((0.1 / 1e10) - 3.14159)
                                            _ = v145
                                        }
                                    }
                                    fmt.Println(v142)
                                } else {
                                    v2 -= (v2 ^ -2147483648)
                                    var v146 float64 = ((3.14159 / (float64(1.5) + (1.5 - (-2.25 + float64(0.1))))) * (float64(0.1) / 3.14159))
                                    var v147 bool = (2147483647 <= 805)
                                    fmt.Println(v3, v147, v142)
                                    _ = v146
                                    _ = v147
                                }
                                v142 = 0
                                _ = fi141
                                _ = v142
                            }
                            fmt.Println(v2, v3, v1)
                        } else {
                            v148 := GenS6{F0: 832}
                            var v149 uint32 = 65535
                            _ = v148
                            _ = v149
                        }
                    }
                }
                v150 := []int{}
                var v151 bool = true
                _ = v150
                _ = v151
            } else {
                var v152 string = "\n\t"
                var v153 int32 = v2
                v154 := []int{fi117, (fi117 >> (((v3 % (((v3 & (^(v3)))) | 1))) & 63)), 42, -2147483648}
                _ = v152
                _ = v153
                _ = v154
            }
        } else {
            if false {
                for fi155 := 0; fi155 < 1; fi155++ {
                    v1 = v1
                    var v156 uint = (v3 ^ (v3 | (v3 | v3)))
                    _ = fi155
                    _ = v156
                }
                if (0 == 0) {
                    if false {
                        for fi157 := 0; fi157 < 19; fi157++ {
                            v158 := []int{fi117}
                            v3 -= uint(v1)
                            v159 := GenS6{F0: int64(fi157)}
                            _ = fi157
                            _ = v158
                            _ = v159
                        }
                        v160 := GenS15{F0: 2, F1: (v13 - genFn0()), F2: (fi117 & -1)}
                        _ = v160
                    }
                    v3 = (v3 % (((v3 | (^(v3)))) | 1))
                    fmt.Println(fi117)
                }
            } else {
                v2 -= 2147483647
                var v161 bool = false
                var v162 uint8 = 255
                _ = v161
                _ = v162
            }
            v48 = 0
            v13 = v13
        }
        _ = fi117
    }
    fmt.Println(v3, v3, v3)
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v48:", v48)
    _ = v1
    _ = v13
    _ = v16
    _ = v2
    _ = v3
    _ = v48
}
