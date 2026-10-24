package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
type GenS2 struct { F0 uint8; F1 int64 }
type GenS3 struct { F0 uint8; F1 uint32; F2 string }
type GenS4 struct { F0 uint8; F1 uint8; F2 int }
type GenS5 struct { F0 uint64; F1 int; F2 uint64 }
type GenS6 struct { F0 bool; F1 int16 }
type GenS7 struct { F0 int64; F1 uint32 }
type GenS8 struct { F0 int16; F1 uint8 }
type GenS9 struct { F0 int64; F1 bool }
type GenS10 struct { F0 string; F1 int32; F2 int64 }
type GenS11 struct { F0 int32 }
type GenS12 struct { F0 string; F1 int64 }
type GenS13 struct { F0 uint32 }
type GenS14 struct { F0 int8; F1 int32 }
type GenS15 struct { F0 uint }
type GenS16 struct { F0 uint16 }
type GenS17 struct { F0 uint8; F1 uint32; F2 bool }
type GenS18 struct { F0 int; F1 int8 }
type GenS19 struct { F0 int64; F1 int8 }
type GenS20 struct { F0 bool; F1 uint64 }
type GenS21 struct { F0 int16; F1 uint8 }
type GenS22 struct { F0 bool; F1 int8 }
type GenS23 struct { F0 uint64; F1 int32 }
type GenS24 struct { F0 uint8; F1 string; F2 int64 }
type GenS25 struct { F0 int }
type GenS26 struct { F0 int; F1 int32 }
type GenS27 struct { F0 uint16; F1 int32; F2 uint8 }
type GenS28 struct { F0 int64; F1 int8; F2 uint8 }
type GenS29 struct { F0 int16; F1 uint16; F2 uint16 }
type GenS30 struct { F0 int8; F1 int16; F2 string }
type GenS31 struct { F0 int8; F1 int16 }
func genFn0(p0 int) uint8 {
    switch p0 {
    case 128:
        var v7 uint16 = 1
        _ = v7
    case 42:
        var v8 uint8 = 11
        _ = v8
    case 127:
        var v9 uint8 = 128
        _ = v9
    }
    if (793 >= 18446744073709551615) {
        var v10 int32 = 0
        v11 := []int{}
        if true {
            var v12 int8 = 35
            if false {
                var v13 uint8 = 2
                for fi14 := 0; fi14 < 15; fi14++ {
                    v15 := GenS1{F0: 65535}
                    switch fi14 {
                    case 127:
                        if len(v11) > 0 {
                            v11[0] = (fi14 & (p0 & p0))
                        }
                    }
                    if true {
                        if len(v11) >= 2 {
                            v11 = v11[1:len(v11)]
                        }
                    } else {
                        switch 42 {
                        case 32767:
                            switch "foo bar" {
                            case "hello":
                                if false {
                                    var v16 float64 = 0.0
                                    v17 := GenS2{F0: 234, F1: -636}
                                    _ = v16
                                    _ = v17
                                } else {
                                    _ = len(v11)
                                    for fi18 := 0; fi18 < 8; fi18++ {
                                        p0 = 2147483647
                                        if (!true) {
                                            switch (fi18 / (((p0 | (p0 << ((65535) & 63)))) | 1)) {
                                            case -2147483648:
                                                var v19 int = fi18
                                                _ = v19
                                            case 128:
                                                var v20 uint32 = 65535
                                                _ = v20
                                            }
                                            var v21 uint = 320
                                            v13 = 255
                                            _ = v21
                                        }
                                        _ = fi18
                                    }
                                    for fi22 := 0; fi22 < 18; fi22++ {
                                        var v23 int = (p0 + 0)
                                        if (!true) {
                                            fmt.Println(fi14, fi14, v13)
                                            if false {
                                                v24 := GenS1{F0: 32767}
                                                fmt.Println(v10, v10)
                                                v10 = v10
                                                _ = v24
                                            }
                                            _ = len(v11)
                                            if false {
                                                v13 -= (^(v13))
                                            }
                                        }
                                        if (!((42 != 173) || true)) {
                                            v25 := GenS3{F0: (v13 << ((0) & 63)), F1: 4294967295, F2: ""}
                                            _ = v25
                                        }
                                        var v26 bool = ((("foo bar" + "\n\t") + "hello") <= "foo bar")
                                        _ = fi22
                                        _ = v23
                                        _ = v26
                                    }
                                }
                            }
                        }
                        var v27 float32 = ((0.1 + float32(0.1)) - -2.25)
                        _ = v27
                    }
                    _ = fi14
                    _ = v15
                }
                var v28 float32 = 1.5
                _ = v13
                _ = v28
            }
            _ = v12
        }
        _ = v10
        _ = v11
    } else {
        var v29 string = "foo bar"
        _ = v29
    }
    var v30 uint32 = 4294967295
    for fi31 := 0; fi31 < 9; fi31++ {
        var v32 int64 = 1
        var v33 string = (("foo bar" + ("foo bar" + ("foo bar" + ("" + "\n\t")))) + "hello")
        v34 := GenS1{F0: 822}
        _ = fi31
        _ = v32
        _ = v33
        _ = v34
    }
    _ = p0
    _ = v30
    return 0
}
func genFn1(p0 uint64, p1 uint32, p2 uint16) string {
    v35 := GenS2{F0: 1, F1: 0}
    if true {
        for fi36 := 0; fi36 < 8; fi36++ {
            if (1e10 >= 1e10) {
                fmt.Println(fi36)
                if false {
                    v37 := GenS4{F0: genFn0(1), F1: 42, F2: fi36}
                    if true {
                        for fi38 := 0; fi38 < 18; fi38++ {
                            fmt.Println(p1)
                            var v39 int32 = -1
                            p1 = p1
                            _ = fi38
                            _ = v39
                        }
                        p0 = p0
                        v40 := []int{(-(fi36)), (fi36 << ((42) & 63)), (^(fi36))}
                        _ = v40
                    }
                    _ = v37
                } else {
                    switch ("\n\t" + ((("\n\t" + "hello") + "\n\t") + "")) {
                    case "\n\t":
                        if ((false || (287 > 42)) && false) {
                            for fi41 := 0; fi41 < 3; fi41++ {
                                p0 = p0
                                _ = fi41
                            }
                        } else {
                            switch ("" + "hello") {
                            case "":
                                if true {
                                    p1 += uint32(p0)
                                    var v42 string = "a"
                                    _ = v42
                                } else {
                                    v43 := []int{1, fi36, (-(fi36))}
                                    var v44 uint = 1
                                    for fi45 := 0; fi45 < 4; fi45++ {
                                        v44 = (v44 ^ 255)
                                        _ = fi45
                                    }
                                    _ = v43
                                    _ = v44
                                }
                            case "hello":
                                if true {
                                    v46 := []int{fi36, (fi36 << ((42) & 63))}
                                    var v47 int16 = 42
                                    _ = len(v46)
                                    _ = v46
                                    _ = v47
                                } else {
                                    if false {
                                        if false {
                                            var v48 uint64 = p0
                                            if false {
                                                if false {
                                                    p2 += (p2 | (p2 + uint16(p0)))
                                                }
                                            } else {
                                                p1 += (^(p1))
                                            }
                                            v48 = (v48 + uint64(p1))
                                            _ = v48
                                        }
                                    } else {
                                        switch (fi36 % ((4294967295) | 1)) {
                                        case -9223372036854775808:
                                            fmt.Println(fi36)
                                        case 32768:
                                            var v49 float32 = 0.1
                                            _ = v49
                                        case -91:
                                            switch (fi36 + 4294967295) {
                                            case 2147483647:
                                                var v50 float32 = 3.14159
                                                _ = v50
                                            }
                                        default:
                                            if false {
                                                var v51 int32 = 0
                                                p1 += 497
                                                _ = v51
                                            } else {
                                                v52 := GenS5{F0: p0, F1: fi36, F2: (p0 % ((18446744073709551615) | 1))}
                                                var v53 uint = uint(p1)
                                                _ = v52
                                                _ = v53
                                            }
                                        }
                                        for fi54 := 0; fi54 < 2; fi54++ {
                                            v55 := []int{}
                                            var v56 uint16 = p2
                                            var v57 string = "hello"
                                            if false {
                                                if true {
                                                    v58 := GenS6{F0: true, F1: 1000}
                                                    v55 = append(v55, fi36)
                                                    _ = v58
                                                } else {
                                                    p1 -= p1
                                                }
                                                if true {
                                                    if (75 < 168) {
                                                        var v59 bool = true
                                                        _ = v59
                                                    }
                                                    if true {
                                                        v60 := GenS7{F0: -1, F1: 2147483648}
                                                        _ = v60
                                                    }
                                                } else {
                                                    v56 = p2
                                                    p2 = p2
                                                    switch 0 {
                                                    case 821:
                                                        v61 := GenS6{F0: false, F1: 0}
                                                        p2 = 255
                                                        var v62 uint32 = p1
                                                        _ = v61
                                                        _ = v62
                                                    default:
                                                        if len(v55) >= 2 {
                                                            v55 = v55[1:len(v55)]
                                                        }
                                                    }
                                                }
                                            } else {
                                                v63 := GenS8{F0: 42, F1: 128}
                                                var v64 uint8 = 144
                                                _ = len(v55)
                                                _ = v63
                                                _ = v64
                                            }
                                            _ = fi54
                                            _ = v55
                                            _ = v56
                                            _ = v57
                                        }
                                    }
                                    fmt.Println(p2, p2, p2)
                                }
                            default:
                                var v65 int32 = 1
                                _ = v65
                            }
                            p1 = (p1 + (p1 ^ (p1 | (^(p1)))))
                        }
                    default:
                        switch "\n\t" {
                        case "foo bar":
                            for fi66 := 0; fi66 < 20; fi66++ {
                                v67 := GenS9{F0: -9223372036854775808, F1: false}
                                var v68 int = fi66
                                var v69 int16 = 359
                                var v70 int64 = 9223372036854775807
                                var v71 uint32 = (p1 / ((p1) | 1))
                                var v72 int8 = -87
                                fmt.Println(fi66)
                                v73 := GenS1{F0: p2}
                                var v74 int8 = 42
                                v74 = (-(v72))
                                _ = fi66
                                _ = v67
                                _ = v68
                                _ = v69
                                _ = v70
                                _ = v71
                                _ = v72
                                _ = v73
                                _ = v74
                            }
                        }
                    }
                    p0 = p0
                    v75 := []int{}
                    _ = v75
                }
                p1 = uint32(p0)
            } else {
                for fi76 := 0; fi76 < 2; fi76++ {
                    var v77 int64 = -448
                    v78 := GenS7{F0: v77, F1: (p1 ^ 2147483648)}
                    var v79 uint8 = 2
                    _ = fi76
                    _ = v77
                    _ = v78
                    _ = v79
                }
                v80 := []int{(fi36 | fi36), (fi36 << ((255) & 63))}
                if ((fi36 & fi36) <= fi36) {
                    switch fi36 {
                    case 204:
                        fmt.Println(fi36)
                    default:
                        if (0.0 == float32(float64(3.14159))) {
                            if (-2147483648 >= -235) {
                                switch "hello" {
                                case "":
                                    v81 := GenS4{F0: 1, F1: genFn0(fi36), F2: -512}
                                    _ = v81
                                case "hello":
                                    if true {
                                        p0 -= (p0 + (p0 >> ((302) & 63)))
                                        for fi82 := 0; fi82 < 4; fi82++ {
                                            var v83 float32 = 1e10
                                            v83 = (-2.25 * ((float32(-2.25) * (0.0 - -2.25)) + v83))
                                            _ = fi82
                                            _ = v83
                                        }
                                        if len(v80) >= 2 {
                                            v80 = v80[1:len(v80)]
                                        }
                                    }
                                default:
                                    var v84 int = (fi36 * 128)
                                    _ = v84
                                }
                                p0 -= (p0 & 4294967295)
                            }
                            switch ((("foo bar" + "hello") + "foo bar") + ("" + "\n\t")) {
                            case "foo bar":
                                for fi85 := 0; fi85 < 12; fi85++ {
                                    fmt.Println(fi85)
                                    p1 = (p1 | p1)
                                    _ = fi85
                                }
                            case "\n\t":
                                for fi86 := 0; fi86 < 14; fi86++ {
                                    if true {
                                        var v87 int16 = -32768
                                        var v88 int32 = 579
                                        p0 -= 4294967295
                                        _ = v87
                                        _ = v88
                                    } else {
                                        v89 := GenS9{F0: 9223372036854775807, F1: (("hello" + ("" + "\n\t")) <= ("foo bar" + ("foo bar" + "a")))}
                                        if true {
                                            fmt.Println(p2, p0)
                                            fmt.Println(fi86)
                                            if len(v80) >= 2 {
                                                v80 = v80[1:len(v80)]
                                            }
                                        } else {
                                            fmt.Println(p0)
                                            var v90 float64 = -2.25
                                            var v91 int16 = int16(p2)
                                            _ = v90
                                            _ = v91
                                        }
                                        _ = v89
                                    }
                                    _ = len(v80)
                                    _ = fi86
                                }
                            }
                        } else {
                            v92 := GenS10{F0: ("" + "foo bar"), F1: int32(fi36), F2: -36}
                            var v93 uint32 = uint32(p2)
                            if true {
                                p0 = p0
                                for fi94 := 0; fi94 < 10; fi94++ {
                                    var v95 int = 9223372036854775807
                                    _ = fi94
                                    _ = v95
                                }
                            } else {
                                v96 := GenS11{F0: 42}
                                fmt.Println(p0)
                                if (42 > uint(p2)) {
                                    var v97 float64 = 0.1
                                    _ = v97
                                }
                                _ = v96
                            }
                            _ = v92
                            _ = v93
                        }
                    }
                    p0 = (p0 ^ p0)
                    fmt.Println(p1, p1, p0)
                }
                _ = v80
            }
            p0 = 1
            _ = fi36
        }
        var v98 int32 = int32(p1)
        var v99 int32 = 42
        _ = v98
        _ = v99
    } else {
        v100 := GenS12{F0: ("\n\t" + ""), F1: -9223372036854775808}
        v101 := []int{127, 9223372036854775807}
        var v102 bool = (-406 != -32768)
        _ = v100
        _ = v101
        _ = v102
    }
    var v103 int32 = 42
    for fi104 := 0; fi104 < 19; fi104++ {
        p0 = p0
        _ = fi104
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v103
    _ = v35
    return "a"
}

func main() {
    var v1 float64 = ((-2.25 + 1e10) * 0.0)
    var v2 string = "hello"
    v3 := make(map[int]uint32)
    v4 := make(map[uint32]bool)
    v5 := make(map[string]bool)
    v6 := make(map[uint32]int)
    v105 := GenS1{F0: 580}
    switch 0 {
    case 127:
        if (false && true) {
            fmt.Println(v1, v1)
            var v106 float32 = -2.25
            switch -325 {
            case 1:
                v107 := []int{-2147483648, 32768, 1, 32767}
                _ = v107
            case -2147483648:
                v106 -= (v106 / (v106 + float32(float64((v106 / ((v106 + v106) - v106))))))
            }
            _ = v106
        }
    case 568:
        for fi108 := 0; fi108 < 2; fi108++ {
            v2 = (("a" + "\n\t") + v2)
            var v109 uint32 = 2
            if (false || (1 != 18446744073709551615)) {
                if false {
                    fmt.Println(v2)
                    v110 := GenS12{F0: (v2 + "\n\t"), F1: 881}
                    v1 = 0.1
                    _ = v110
                } else {
                    v111 := GenS5{F0: 663, F1: (fi108 - fi108), F2: 42}
                    if ((2 > 0) && true) {
                        v1 += (float64(0.0) + v1)
                        v109 += 306
                        if (genFn0(-857) != 128) {
                            v112 := GenS6{F0: (1000 >= 1), F1: -32768}
                            _ = v112
                        }
                    }
                    _ = v111
                }
                v113 := []int{-307, (fi108 + (-(fi108))), fi108}
                var v114 float64 = v1
                _ = v113
                _ = v114
            }
            _ = fi108
            _ = v109
        }
    case 454:
        _ = len(v6)
    }
    var v115 string = ("a" + v2)
    if (0.1 < (v1 + ((v1 + float64(0.0)) * -2.25))) {
        for fi116 := 0; fi116 < 11; fi116++ {
            for fi117 := 0; fi117 < 8; fi117++ {
                var v118 uint8 = 71
                var v119 int16 = 32767
                _ = fi117
                _ = v118
                _ = v119
            }
            v1 -= float64((1.5 - (3.14159 / (0.1 / (0.0 - float32(v1))))))
            _ = fi116
        }
        if true {
            for fi120 := 0; fi120 < 8; fi120++ {
                v121 := GenS13{F0: 65535}
                if ((v1 / 1e10) > v1) {
                    for fi122 := 0; fi122 < 18; fi122++ {
                        v115 = (v2 + v115)
                        v123 := GenS14{F0: 57, F1: 1}
                        v124 := GenS15{F0: 18446744073709551615}
                        if (1 >= 486) {
                            v125 := GenS16{F0: 594}
                            _ = v125
                        } else {
                            var v126 uint32 = 872
                            v2 = (v2 + ((((("hello" + v2) + v2) + (v2 + v115)) + v115) + v115))
                            v127 := GenS11{F0: 428}
                            _ = v126
                            _ = v127
                        }
                        _ = fi122
                        _ = v123
                        _ = v124
                    }
                }
                _ = fi120
                _ = v121
            }
            fmt.Println(v2, v115)
        } else {
            var v128 string = "foo bar"
            _ = v128
        }
        v115 = ""
    }
    var v129 float64 = ((v1 * v1) + v1)
    v130 := GenS14{F0: 127, F1: 0}
    v129 -= 0.1
    var v131 uint = 18446744073709551615
    for fi132 := 0; fi132 < 11; fi132++ {
        switch ("\n\t" + v115) {
        case "hello":
            if (0 >= -9223372036854775808) {
                if true {
                    switch 1 {
                    case 9223372036854775807:
                        if false {
                            fmt.Println(v129, v1, v2)
                            fmt.Println(fi132)
                        } else {
                            v133 := []int{fi132, (fi132 | (-(fi132))), (^(fi132)), fi132}
                            fmt.Println(v2)
                            _ = v133
                        }
                    case 32768:
                        var v134 int64 = -9223372036854775808
                        _ = v134
                    case -109:
                        v135 := []int{fi132, (fi132 - (fi132 ^ (fi132 * 753))), fi132}
                        _ = v135
                    }
                    fmt.Println(fi132, v115, fi132)
                }
            }
        case "a":
            for fi136 := 0; fi136 < 19; fi136++ {
                if (true && false) {
                    if true {
                        var v137 uint64 = 0
                        _ = v137
                    } else {
                        switch fi132 {
                        case 42:
                            if (12 < 42) {
                                if (((fi136 | (fi132 + fi132)) != fi132) && ((4294967295 != 4294967295) && false)) {
                                    if true {
                                        switch "" {
                                        case "":
                                            if (!((fi136 % ((fi136) | 1)) > (^(fi132)))) {
                                                v138 := GenS10{F0: (v2 + ((v2 + v2) + "foo bar")), F1: 63, F2: 42}
                                                _ = v138
                                            }
                                        }
                                        v6[2147483648] = (fi132 ^ (fi132 ^ (^(fi136))))
                                        var v139 int32 = -1
                                        _ = v139
                                    }
                                } else {
                                    var v140 uint64 = 2
                                    v141 := []int{-9223372036854775808}
                                    var v142 bool = (!true)
                                    _ = v140
                                    _ = v141
                                    _ = v142
                                }
                            } else {
                                var v143 bool = false
                                v131 = 18446744073709551615
                                if v143 {
                                    v144 := []int{fi136}
                                    _ = v144
                                }
                                _ = v143
                            }
                        }
                        if (true && true) {
                            v145 := []int{}
                            _ = v145
                        }
                        var v146 int64 = -16
                        v146 -= v146
                        if false {
                            _ = len(v5)
                            v129 = v1
                        } else {
                            if false {
                                if (!(true || (false || (true || false)))) {
                                    v147 := GenS17{F0: 42, F1: 2147483647, F2: false}
                                    fmt.Println(v2, v131, fi132)
                                    _ = v147
                                } else {
                                    v148 := GenS14{F0: 0, F1: 938}
                                    switch -385 {
                                    case -9223372036854775808:
                                        if (false && false) {
                                            if (-553 >= -32768) {
                                                var v149 bool = (-750 != 1)
                                                var v150 int8 = -128
                                                v151 := GenS18{F0: fi132, F1: v150}
                                                _ = v149
                                                _ = v150
                                                _ = v151
                                            }
                                            var v152 uint8 = genFn0(fi136)
                                            var v153 int8 = 127
                                            _ = v152
                                            _ = v153
                                        } else {
                                            v115 = v2
                                            fmt.Println(v146)
                                        }
                                    case 9223372036854775807:
                                        if false {
                                            if ("a" == ((v2 + "foo bar") + v115)) {
                                                v115 = "a"
                                                fmt.Println(v129, v1, v1)
                                            }
                                        } else {
                                            delete(v6, 1)
                                        }
                                    case 128:
                                        switch "" {
                                        case "\n\t":
                                            if (1e10 != 3.14159) {
                                                var v154 string = (("" + (v2 + v115)) + "\n\t")
                                                fmt.Println(v129)
                                                v155 := v6[340]
                                                _ = v154
                                                _ = v155
                                            }
                                        case "":
                                            if false {
                                                fmt.Println(v131)
                                                v156 := []int{}
                                                _ = v156
                                            } else {
                                                fmt.Println(v115, fi136)
                                                var v157 string = (v115 + "a")
                                                if true {
                                                    fmt.Println(v2)
                                                    v158 := GenS19{F0: -1, F1: 0}
                                                    _ = v158
                                                } else {
                                                    v129 = (v129 + (1.5 + (3.14159 / v1)))
                                                    v6[65535] = (fi136 * fi132)
                                                    switch "foo bar" {
                                                    case "a":
                                                        v146 = v146
                                                    case "hello":
                                                        var v159 uint = 732
                                                        _ = v159
                                                    }
                                                }
                                                _ = v157
                                            }
                                        }
                                    }
                                    if true {
                                        fmt.Println(v129, v2, v115)
                                        v160 := []int{(fi136 | fi136), fi136, (fi132 * fi132), (fi132 << ((v131) & 63))}
                                        var v161 int32 = -356
                                        _ = v160
                                        _ = v161
                                    } else {
                                        var v162 string = "hello"
                                        _ = v162
                                    }
                                    _ = v148
                                }
                                var v163 uint8 = 42
                                var v164 int32 = -2147483648
                                _ = v163
                                _ = v164
                            }
                            v165 := GenS14{F0: 0, F1: -2147483648}
                            _ = v165
                        }
                        var v166 float64 = float64(0.0)
                        _ = v146
                        _ = v166
                    }
                } else {
                    if true {
                        v167 := GenS18{F0: fi132, F1: -24}
                        switch 9223372036854775807 {
                        case 32768:
                            switch 4294967295 {
                            case -648:
                                v1 -= v1
                            case -278:
                                var v168 int = fi132
                                _ = v168
                            }
                        case 127:
                            v129 = v129
                        }
                        _ = v167
                    } else {
                        if false {
                            if false {
                                v131 += 0
                            } else {
                                if (false || (1000 == -1)) {
                                    v115 = v2
                                    fmt.Println(v115, v129, v131)
                                } else {
                                    v169 := []int{1, fi132}
                                    fmt.Println(v115, v131, v1)
                                    var v170 float32 = -2.25
                                    _ = v169
                                    _ = v170
                                }
                                fmt.Println(v2, v129, v115)
                                v129 += (v1 + 1.5)
                            }
                            if ((((0 == 42) && true) || false) || true) {
                                v1 = v1
                            } else {
                                if (1.5 <= (-2.25 + v1)) {
                                    switch (("" + (((("\n\t" + ("a" + v2)) + v2) + (v2 + v115)) + ("\n\t" + ("foo bar" + "")))) + v2) {
                                    case "hello":
                                        var v171 int8 = 42
                                        _ = v171
                                    case "foo bar":
                                        var v172 int = (fi132 / ((fi132) | 1))
                                        _ = v172
                                    }
                                    fmt.Println(fi132)
                                }
                                v2 = v2
                                var v173 uint16 = 785
                                v174 := []int{(fi136 & (fi132 & fi136)), (fi132 >> ((v131) & 63)), (fi132 >> (((v131 + (v131 * (v131 * v131)))) & 63))}
                                _ = v173
                                _ = v174
                            }
                        } else {
                            v175 := GenS1{F0: 94}
                            _ = v175
                        }
                        v1 = v1
                    }
                    v176 := GenS2{F0: 42, F1: -123}
                    switch (v115 + (v2 + (v2 + "\n\t"))) {
                    case "hello":
                        var v177 float32 = 3.14159
                        _ = v177
                    default:
                        fmt.Println(v115)
                    }
                    _ = v176
                }
                _ = fi136
            }
        }
        v178 := GenS16{F0: 42}
        _ = fi132
        _ = v178
    }
    for fi179 := 0; fi179 < 20; fi179++ {
        var v180 uint16 = 274
        switch v2 {
        case "\n\t":
            for fi181 := 0; fi181 < 5; fi181++ {
                var v182 uint8 = 2
                var v183 uint32 = 265
                _ = fi181
                _ = v182
                _ = v183
            }
        }
        _ = fi179
        _ = v180
    }
    v184 := GenS20{F0: (!(false || (-1 > 127))), F1: 4294967295}
    var v185 float32 = (1.5 * ((1.5 + -2.25) / 0.1))
    var v186 bool = true
    var v187 uint32 = 4294967295
    v188 := []int{-1}
    if v186 {
        if v186 {
            if (0 != 2) {
                v186 = (4294967295 < uint64(v131))
            } else {
                switch 42 {
                case -1:
                    var v189 uint = 417
                    _ = v189
                default:
                    v1 = 0.0
                }
                if v186 {
                    if v186 {
                        for fi190 := 0; fi190 < 12; fi190++ {
                            v191 := GenS21{F0: -532, F1: 255}
                            v5["a"] = v186
                            _ = fi190
                            _ = v191
                        }
                    }
                } else {
                    for fi192 := 0; fi192 < 20; fi192++ {
                        var v193 uint8 = 104
                        for fi194 := 0; fi194 < 8; fi194++ {
                            if (v193 == v193) {
                                _ = len(v188)
                                v115 = v2
                                _ = len(v6)
                            } else {
                                if v186 {
                                    var v195 int = (fi192 & fi192)
                                    _ = v195
                                } else {
                                    fmt.Println(v2, v1, v131)
                                    fmt.Println(fi192)
                                    var v196 int8 = 0
                                    _ = v196
                                }
                                v197 := v4[(v187 & 2)]
                                _ = v197
                            }
                            var v198 int32 = 1
                            v199 := GenS2{F0: v193, F1: -1}
                            _ = fi194
                            _ = v198
                            _ = v199
                        }
                        v200 := GenS22{F0: v186, F1: 0}
                        _ = fi192
                        _ = v193
                        _ = v200
                    }
                }
                fmt.Println(v186, v187, v187)
            }
            switch v2 {
            case "foo bar":
                for fi201 := 0; fi201 < 11; fi201++ {
                    fmt.Println(v186, fi201)
                    for fi202 := 0; fi202 < 15; fi202++ {
                        if (!v186) {
                            var v203 int32 = 0
                            _ = v203
                        }
                        fmt.Println(v131, v2, v115)
                        v3[(^(fi201))] = 42
                        v131 = 65535
                        v131 += (v131 - (v131 / (((^(v131))) | 1)))
                        _ = fi202
                    }
                    _ = fi201
                }
            }
            if (32767 > -155) {
                var v204 uint = (v131 >> ((v131) & 63))
                fmt.Println(v185)
                _ = v204
            }
        } else {
            switch 632 {
            case 2147483647:
                if (4294967295 != 1) {
                    fmt.Println(v115, v187, v186)
                } else {
                    v205 := GenS20{F0: v186, F1: 412}
                    _ = v205
                }
            case 32768:
                if (("a" <= "\n\t") && (!v186)) {
                    switch v115 {
                    case "foo bar":
                        v1 -= float64(3.14159)
                    case "hello":
                        v129 = -2.25
                    case "":
                        fmt.Println(v186)
                    }
                    v206 := GenS13{F0: (v187 % (((v187 % ((v187) | 1))) | 1))}
                    v187 = 65535
                    _ = v206
                }
            case 444:
                if (171 <= 2) {
                    v207 := GenS7{F0: 9223372036854775807, F1: (v187 ^ v187)}
                    if v186 {
                        if v186 {
                            switch 128 {
                            case 32768:
                                for fi208 := 0; fi208 < 2; fi208++ {
                                    v209 := GenS4{F0: 0, F1: 127, F2: 32767}
                                    v129 = v1
                                    var v210 uint16 = 255
                                    _ = fi208
                                    _ = v209
                                    _ = v210
                                }
                            default:
                                var v211 uint = (v131 + v131)
                                _ = v211
                            }
                            fmt.Println(v129, v1, v131)
                        } else {
                            for fi212 := 0; fi212 < 9; fi212++ {
                                switch fi212 {
                                case -449:
                                    for fi213 := 0; fi213 < 5; fi213++ {
                                        if v186 {
                                            v214 := GenS17{F0: 0, F1: 559, F2: ((false && true) && v186)}
                                            switch fi213 {
                                            case 32767:
                                                v185 -= float32(v129)
                                            default:
                                                v187 = v187
                                            }
                                            v185 = 0.0
                                            _ = v214
                                        } else {
                                            switch 42 {
                                            case 467:
                                                if v186 {
                                                    v215 := GenS23{F0: 42, F1: 718}
                                                    _ = v215
                                                } else {
                                                    var v216 int8 = -101
                                                    v115 = v2
                                                    v6[(v187 ^ 408)] = (-(fi213))
                                                    _ = v216
                                                }
                                            case -99:
                                                v115 = "hello"
                                            case 32767:
                                                if (v186 || false) {
                                                    if (v185 < v185) {
                                                        var v217 int16 = 376
                                                        _ = v217
                                                    }
                                                } else {
                                                    v218 := GenS24{F0: genFn0(fi212), F1: v2, F2: -9223372036854775808}
                                                    v219 := v3[fi212]
                                                    if v186 {
                                                        fmt.Println(v187, fi212)
                                                        if (65535 <= v219) {
                                                            fmt.Println(v185, v129)
                                                            switch fi212 {
                                                            case -9223372036854775808:
                                                                fmt.Println(v115)
                                                            case 9223372036854775807:
                                                                if ((v2 + "foo bar") <= v115) {
                                                                    switch 127 {
                                                                    case -9223372036854775808:
                                                                        var v220 int8 = 127
                                                                        _ = v220
                                                                    case 42:
                                                                        v129 = ((3.14159 + v129) + (((3.14159 / (float64((v185 * v185)) - 1.5)) / v1) + ((1.5 - (v129 / (v129 / (((0.1 - (3.14159 + float64(-2.25))) - 1e10) * (float64(3.14159) - 0.1))))) + float64(float32((v129 * 1e10))))))
                                                                    default:
                                                                        if true {
                                                                            var v221 int64 = -1
                                                                            _ = v221
                                                                        }
                                                                    }
                                                                    v222 := v3[127]
                                                                    _ = v222
                                                                } else {
                                                                    var v223 uint32 = 42
                                                                    v224 := GenS4{F0: 2, F1: 127, F2: (fi213 - fi212)}
                                                                    v5[v2] = v186
                                                                    _ = v223
                                                                    _ = v224
                                                                }
                                                            }
                                                            v5["a"] = v186
                                                        }
                                                        if (((v185 / 1.0) - v185) < v185) {
                                                            if (v187 != (v187 ^ (v219 >> (((v131 - v131)) & 63)))) {
                                                                _ = len(v4)
                                                            }
                                                            v131 -= v131
                                                            v225 := v3[fi213]
                                                            _ = v225
                                                        } else {
                                                            v187 = 150
                                                        }
                                                    }
                                                    if ((v131 % ((1) | 1)) > 2) {
                                                        v226 := GenS24{F0: 2, F1: ((v2 + ("hello" + v2)) + (("hello" + "hello") + v115)), F2: -9223372036854775808}
                                                        _ = v226
                                                    }
                                                    var v227 string = v115
                                                    _ = v218
                                                    _ = v219
                                                    _ = v227
                                                }
                                                if (42 != 65535) {
                                                    var v228 string = v115
                                                    _ = v228
                                                } else {
                                                    if v186 {
                                                        var v229 bool = (1 == 255)
                                                        _ = v229
                                                    } else {
                                                        if v186 {
                                                            v230 := GenS25{F0: (fi213 >> ((v131) & 63))}
                                                            v115 = ((v2 + "a") + "foo bar")
                                                            switch 169 {
                                                            case -1:
                                                                v185 = ((float32(v129) - (0.1 * float32(v1))) - ((0.1 * v185) - float32(v129)))
                                                            case 42:
                                                                v129 += v1
                                                            }
                                                            _ = v230
                                                        }
                                                        if v186 {
                                                            v231 := v3[(fi213 >> ((2) & 63))]
                                                            if (100 > 123) {
                                                                _ = len(v4)
                                                                var v232 uint32 = uint32(v131)
                                                                fmt.Println(v115, v115, v1)
                                                                _ = v232
                                                            }
                                                            _ = v231
                                                        } else {
                                                            v233 := GenS3{F0: 141, F1: v187, F2: v115}
                                                            switch ("" + v115) {
                                                            case "foo bar":
                                                                if v186 {
                                                                    if (!v186) {
                                                                        var v234 string = (("\n\t" + ("\n\t" + "")) + ("a" + ""))
                                                                        _ = v234
                                                                    } else {
                                                                        switch (fi212 << ((v131) & 63)) {
                                                                        case 84:
                                                                            switch (v115 + v115) {
                                                                            case "\n\t":
                                                                                v235 := v4[v187]
                                                                                _ = v235
                                                                            default:
                                                                                switch (((((v115 + "a") + v2) + "\n\t") + v2) + v2) {
                                                                                case "hello":
                                                                                    v236 := GenS8{F0: 751, F1: 2}
                                                                                    _ = v236
                                                                                case "\n\t":
                                                                                    var v237 float64 = 0.0
                                                                                    _ = v237
                                                                                default:
                                                                                    switch "\n\t" {
                                                                                    case "a":
                                                                                        v238 := GenS26{F0: 32768, F1: -1}
                                                                                        _ = v238
                                                                                    case "foo bar":
                                                                                        var v239 string = v2
                                                                                        _ = v239
                                                                                    default:
                                                                                        if (v1 != 3.14159) {
                                                                                            var v240 int16 = 196
                                                                                            if v186 {
                                                                                                switch ((((v2 + v115) + v2) + v115) + v115) {
                                                                                                case "foo bar":
                                                                                                    var v241 int = -1
                                                                                                    _ = v241
                                                                                                case "\n\t":
                                                                                                    fmt.Println(v129)
                                                                                                case "hello":
                                                                                                    switch ((v115 + (v115 + ((v115 + v2) + v2))) + "a") {
                                                                                                    case "foo bar":
                                                                                                        v5[(v2 + (v2 + ("" + v2)))] = false
                                                                                                    }
                                                                                                default:
                                                                                                    switch ((v2 + ((("\n\t" + (("\n\t" + v2) + (((v2 + v2) + v115) + (v115 + v2)))) + v115) + (((v115 + v115) + "") + v115))) + v2) {
                                                                                                    case "a":
                                                                                                        v5[(v115 + v115)] = (((v240 > (v240 << (((^(v131))) & 63))) && (203 == 127)) && (v131 == (v131 % ((v131) | 1))))
                                                                                                    default:
                                                                                                        v188 = append(v188, int(v187))
                                                                                                    }
                                                                                                }
                                                                                                switch (fi212 * (^(fi213))) {
                                                                                                case 128:
                                                                                                    v186 = (((v1 * v1) >= 0.0) || ((42 > 0) || ((!v186) || (933 > 816))))
                                                                                                case -886:
                                                                                                    var v242 float32 = v185
                                                                                                    _ = v242
                                                                                                case 1:
                                                                                                    if (-1 > -423) {
                                                                                                        var v243 int32 = 42
                                                                                                        _ = v243
                                                                                                    }
                                                                                                }
                                                                                                fmt.Println(v131, fi212, v129)
                                                                                            } else {
                                                                                                switch v115 {
                                                                                                case "hello":
                                                                                                    v244 := GenS19{F0: 989, F1: -94}
                                                                                                    _ = v244
                                                                                                case "foo bar":
                                                                                                    fmt.Println(v115, v187)
                                                                                                case "a":
                                                                                                    if ((false && true) || (100 == 127)) {
                                                                                                        if v186 {
                                                                                                            if len(v188) > 0 {
                                                                                                                v188[0] = -228
                                                                                                            }
                                                                                                            v115 = v115
                                                                                                        }
                                                                                                        var v245 int32 = -2147483648
                                                                                                        _ = v245
                                                                                                    }
                                                                                                }
                                                                                                if true {
                                                                                                    _ = len(v188)
                                                                                                    if (2 != 93) {
                                                                                                        fmt.Println(v131, v187, v115)
                                                                                                        v185 = (float32(1e10) - -2.25)
                                                                                                        var v246 float32 = v185
                                                                                                        _ = v246
                                                                                                    }
                                                                                                    _ = len(v4)
                                                                                                } else {
                                                                                                    fmt.Println(fi212, fi213, v2)
                                                                                                    delete(v6, v187)
                                                                                                }
                                                                                                v115 = v2
                                                                                            }
                                                                                            _ = len(v188)
                                                                                            _ = v240
                                                                                        } else {
                                                                                            v247 := GenS11{F0: 42}
                                                                                            v5[(v115 + "hello")] = ((v115 + "a") == (v2 + "hello"))
                                                                                            _ = v247
                                                                                        }
                                                                                    }
                                                                                }
                                                                            }
                                                                        default:
                                                                            if (80 < 2147483647) {
                                                                                switch v2 {
                                                                                case "":
                                                                                    if ((1.5 / (3.14159 / v185)) == (0.0 * v185)) {
                                                                                        if true {
                                                                                            var v248 string = ((("a" + v2) + ((v2 + ("\n\t" + ("" + v2))) + ("" + (v115 + (((v2 + "a") + (v115 + v2)) + ""))))) + (((v115 + ("" + "")) + v115) + (v2 + ((("foo bar" + ("hello" + "foo bar")) + "\n\t") + (v2 + (v2 + (v115 + v115)))))))
                                                                                            _ = v248
                                                                                        } else {
                                                                                            fmt.Println(v131)
                                                                                        }
                                                                                        v249 := GenS27{F0: 665, F1: 2147483647, F2: genFn0((fi212 + (^(fi212))))}
                                                                                        _ = v249
                                                                                    } else {
                                                                                        v2 = (v115 + (((v2 + "\n\t") + "foo bar") + "a"))
                                                                                    }
                                                                                    fmt.Println(v1, v186, v2)
                                                                                    fmt.Println(v1, v129, v129)
                                                                                    if (!v186) {
                                                                                        if (((v187 << (((v131 | (v131 / ((v131) | 1)))) & 63)) >= 388) || v186) {
                                                                                            if len(v188) > 0 {
                                                                                                v188[0] = fi213
                                                                                            }
                                                                                            v250 := GenS11{F0: 42}
                                                                                            var v251 float32 = ((-2.25 - (v185 - float32(float64(0.1)))) / 1.5)
                                                                                            _ = v250
                                                                                            _ = v251
                                                                                        } else {
                                                                                            v131 = v131
                                                                                        }
                                                                                        if (!(((v131 - (^(v131))) == v131) && (2 < 927))) {
                                                                                            v2 = ("" + (v2 + v115))
                                                                                        }
                                                                                        fmt.Println(v131, v129, v186)
                                                                                    }
                                                                                case "foo bar":
                                                                                    v252 := GenS1{F0: 1}
                                                                                    _ = v252
                                                                                default:
                                                                                    v186 = v186
                                                                                }
                                                                                switch fi212 {
                                                                                case 128:
                                                                                    fmt.Println(v131)
                                                                                default:
                                                                                    if (v186 && (-1 > 0)) {
                                                                                        if v186 {
                                                                                            switch v2 {
                                                                                            case "":
                                                                                                var v253 uint8 = 9
                                                                                                _ = v253
                                                                                            case "\n\t":
                                                                                                v131 = v131
                                                                                                v254 := GenS21{F0: -432, F1: 1}
                                                                                                _ = v254
                                                                                            default:
                                                                                                var v255 int8 = -51
                                                                                                _ = v255
                                                                                            }
                                                                                            v115 = ((("\n\t" + v2) + "") + ((("a" + "a") + v115) + ("\n\t" + (v2 + v115))))
                                                                                            v256 := GenS4{F0: 127, F1: genFn0(-82), F2: (fi213 * (fi213 / ((2147483647) | 1)))}
                                                                                            v129 -= v129
                                                                                            _ = v256
                                                                                        }
                                                                                        switch fi212 {
                                                                                        case 573:
                                                                                            var v257 int16 = 1000
                                                                                            _ = v257
                                                                                        case 177:
                                                                                            var v258 int16 = 767
                                                                                            _ = v258
                                                                                        }
                                                                                    }
                                                                                    if v186 {
                                                                                        switch v2 {
                                                                                        case "a":
                                                                                            var v259 int16 = 42
                                                                                            _ = v259
                                                                                        }
                                                                                        v260 := GenS28{F0: -968, F1: 1, F2: 1}
                                                                                        _ = len(v188)
                                                                                        v261 := GenS29{F0: 0, F1: 2, F2: 941}
                                                                                        _ = v260
                                                                                        _ = v261
                                                                                    } else {
                                                                                        v4[(v187 | (^(v187)))] = v186
                                                                                        v262 := GenS23{F0: 2, F1: 2147483647}
                                                                                        v186 = (genFn0((fi212 ^ -1)) != 2)
                                                                                        _ = v262
                                                                                    }
                                                                                }
                                                                            }
                                                                        }
                                                                        fmt.Println(v187)
                                                                    }
                                                                }
                                                            default:
                                                                var v263 bool = (111 != 127)
                                                                _ = v263
                                                            }
                                                            v187 += (v187 << ((255) & 63))
                                                            _ = v233
                                                        }
                                                    }
                                                    var v264 int32 = 1
                                                    v2 = ((((v2 + v115) + (v2 + v2)) + (v115 + "")) + "a")
                                                    var v265 float32 = 1.5
                                                    _ = v264
                                                    _ = v265
                                                }
                                            default:
                                                _ = len(v188)
                                            }
                                            var v266 int = fi213
                                            v267 := GenS30{F0: 1, F1: -281, F2: v115}
                                            v268 := GenS21{F0: 316, F1: 127}
                                            v187 = (v187 >> ((255) & 63))
                                            fmt.Println(fi213, fi213)
                                            var v269 uint16 = 2
                                            _ = v266
                                            _ = v267
                                            _ = v268
                                            _ = v269
                                        }
                                        if len(v188) > 0 {
                                            v188[0] = (fi212 << ((v131) & 63))
                                        }
                                        v131 += 1
                                        v187 = (v187 | (v187 + 4294967295))
                                        _ = fi213
                                    }
                                default:
                                    v187 += v187
                                }
                                if false {
                                    fmt.Println(fi212, v131, v1)
                                }
                                for fi270 := 0; fi270 < 12; fi270++ {
                                    fmt.Println(v131)
                                    v271 := GenS3{F0: 42, F1: 42, F2: v115}
                                    _ = fi270
                                    _ = v271
                                }
                                _ = fi212
                            }
                        }
                    }
                    v187 -= v187
                    _ = v207
                }
            default:
                v272 := GenS23{F0: 4294967295, F1: 194}
                _ = v272
            }
            v273 := GenS19{F0: 846, F1: -128}
            v4[2] = true
            _ = v273
        }
    }
    fmt.Println(v1, v129)
    v274 := GenS31{F0: -54, F1: -75}
    fmt.Println("v1:", v1)
    fmt.Println("v115:", v115)
    fmt.Println("v129:", v129)
    fmt.Println("v131:", v131)
    fmt.Println("v185:", v185)
    fmt.Println("v186:", v186)
    fmt.Println("v187:", v187)
    fmt.Println("slicelen v188:", len(v188), v188)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v105
    _ = v115
    _ = v129
    _ = v130
    _ = v131
    _ = v184
    _ = v185
    _ = v186
    _ = v187
    _ = v188
    _ = v2
    _ = v274
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
