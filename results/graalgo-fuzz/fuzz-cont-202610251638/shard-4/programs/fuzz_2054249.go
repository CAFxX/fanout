package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 int }
type GenS2 struct { F0 uint }
type GenS3 struct { F0 int8 }
type GenS4 struct { F0 uint8; F1 string; F2 int64 }
type GenS5 struct { F0 uint8 }
type GenS6 struct { F0 int32; F1 uint8; F2 bool }
type GenS7 struct { F0 uint32 }
type GenS8 struct { F0 int16; F1 int64; F2 string }
type GenS9 struct { F0 uint8; F1 uint32; F2 int }
type GenS10 struct { F0 uint32; F1 int64 }
type GenS11 struct { F0 int; F1 uint32; F2 string }
type GenS12 struct { F0 string; F1 uint32; F2 uint8 }
func genFn0(p0 int16, p1 string) int {
    switch -9223372036854775808 {
    case 0:
        for fi4 := 0; fi4 < 11; fi4++ {
            v5 := GenS1{F0: 600, F1: fi4}
            var v6 uint = 255
            _ = fi4
            _ = v5
            _ = v6
        }
    case 1:
        if true {
            p0 += (p0 | (p0 | (^(p0))))
            var v7 int = 0
            _ = v7
        }
    default:
        if ((0 >= 128) && (2147483647 != 0)) {
            v8 := []int{-854}
            _ = v8
        } else {
            for fi9 := 0; fi9 < 5; fi9++ {
                p0 -= 0
                _ = fi9
            }
        }
    }
    for fi10 := 0; fi10 < 1; fi10++ {
        if (1 < 843) {
            p1 = (p1 + "foo bar")
        } else {
            var v11 int8 = -43
            for fi12 := 0; fi12 < 1; fi12++ {
                fmt.Println(p0, fi12)
                switch (p1 + "\n\t") {
                case "\n\t":
                    if true {
                        v11 -= (v11 * 100)
                        v11 += v11
                    }
                case "":
                    p1 = ("foo bar" + p1)
                case "a":
                    var v13 float64 = -2.25
                    if (4294967295 < 42) {
                        var v14 bool = true
                        var v15 int64 = 42
                        _ = v14
                        _ = v15
                    } else {
                        if true {
                            var v16 int64 = 42
                            if (true || false) {
                                if true {
                                    v17 := GenS1{F0: 2, F1: fi12}
                                    v16 += v16
                                    _ = v17
                                } else {
                                    v16 = (v16 * v16)
                                    v18 := []int{fi12, fi10, (fi12 >> ((18446744073709551615) & 63))}
                                    _ = v18
                                }
                                v13 -= float64(0.0)
                                v19 := GenS1{F0: 2, F1: (fi10 | 409)}
                                _ = v19
                            }
                            _ = v16
                        }
                        p1 = p1
                    }
                    var v20 int64 = -9223372036854775808
                    _ = v13
                    _ = v20
                default:
                    v21 := GenS2{F0: 0}
                    _ = v21
                }
                _ = fi12
            }
            for fi22 := 0; fi22 < 19; fi22++ {
                fmt.Println(v11, fi22, p0)
                var v23 int64 = 42
                var v24 uint64 = 4294967295
                v25 := []int{fi22}
                _ = fi22
                _ = v23
                _ = v24
                _ = v25
            }
            _ = v11
        }
        fmt.Println(p0)
        _ = fi10
    }
    if false {
        var v26 uint16 = 901
        _ = v26
    } else {
        for fi27 := 0; fi27 < 16; fi27++ {
            if true {
                p0 += p0
            } else {
                for fi28 := 0; fi28 < 2; fi28++ {
                    var v29 uint32 = 42
                    _ = fi28
                    _ = v29
                }
            }
            if false {
                p0 = 32767
                if (!((-2.25 - 1.5) < 0.0)) {
                    var v30 int8 = -1
                    v31 := []int{(fi27 % ((fi27) | 1)), fi27, (fi27 << ((42) & 63)), fi27}
                    _ = v30
                    _ = v31
                } else {
                    var v32 int8 = 0
                    _ = v32
                }
                switch -2147483648 {
                case -2147483648:
                    if true {
                        switch p1 {
                        case "hello":
                            if (p0 >= (^(p0))) {
                                fmt.Println(fi27, p1)
                                p0 -= -1
                                fmt.Println(p1, p0, fi27)
                            }
                        case "a":
                            v33 := GenS1{F0: 238, F1: (fi27 >> ((42) & 63))}
                            _ = v33
                        }
                        for fi34 := 0; fi34 < 20; fi34++ {
                            fmt.Println(p0, p0)
                            _ = fi34
                        }
                        if (42 <= 118) {
                            var v35 int32 = 0
                            _ = v35
                        }
                    } else {
                        p1 = "a"
                        if false {
                            v36 := GenS2{F0: 18446744073709551615}
                            switch p1 {
                            case "":
                                var v37 bool = (127 <= -78)
                                _ = v37
                            case "hello":
                                p0 = p0
                            case "a":
                                fmt.Println(p1)
                            }
                            _ = v36
                        }
                        fmt.Println(fi27, fi27, p0)
                    }
                case 32768:
                    var v38 int8 = 42
                    _ = v38
                case 4294967295:
                    var v39 uint = 2
                    _ = v39
                default:
                    v40 := GenS3{F0: 100}
                    _ = v40
                }
            } else {
                fmt.Println(p0, fi27, fi27)
            }
            _ = fi27
        }
    }
    var v41 uint64 = 583
    _ = p0
    _ = p1
    _ = v41
    return 0
}
func genFn1(p0 int, p1 int8) bool {
    fmt.Println(p1, p0)
    switch "" {
    case "foo bar":
        for fi42 := 0; fi42 < 20; fi42++ {
            for fi43 := 0; fi43 < 17; fi43++ {
                p1 = p1
                if false {
                    if (!(((p1 >> ((282) & 63)) < (p1 * -128)) && false)) {
                        if true {
                            var v44 float64 = 0.0
                            _ = v44
                        }
                        v45 := GenS4{F0: 0, F1: "foo bar", F2: 0}
                        _ = v45
                    } else {
                        p0 = genFn0(-32768, ((((("a" + "hello") + ("a" + "foo bar")) + "hello") + "\n\t") + "\n\t"))
                        v46 := []int{639}
                        p0 = (^(fi42))
                        _ = v46
                    }
                } else {
                    var v47 int32 = -365
                    if true {
                        p0 = 42
                    } else {
                        switch "hello" {
                        case "":
                            p1 = 0
                        case "foo bar":
                            if (9223372036854775807 > -1) {
                                var v48 int32 = (v47 >> ((0) & 63))
                                v49 := GenS1{F0: 35, F1: fi43}
                                _ = v48
                                _ = v49
                            }
                            if (true && false) {
                                var v50 int64 = -1
                                if (42 > p1) {
                                    p0 -= (fi43 + (fi42 + (fi43 << ((879) & 63))))
                                    v51 := GenS1{F0: 4294967295, F1: 32767}
                                    _ = v51
                                } else {
                                    v52 := GenS2{F0: 18446744073709551615}
                                    var v53 bool = (true || true)
                                    v50 -= v50
                                    _ = v52
                                    _ = v53
                                }
                                v54 := GenS1{F0: 0, F1: fi43}
                                v55 := GenS5{F0: 42}
                                _ = v50
                                _ = v54
                                _ = v55
                            } else {
                                var v56 uint8 = 2
                                v57 := []int{}
                                if (508 <= 4294967295) {
                                    switch fi43 {
                                    case 42:
                                        fmt.Println(fi43)
                                    case 128:
                                        if false {
                                            _ = len(v57)
                                            var v58 uint32 = 1
                                            _ = v58
                                        }
                                    case -9223372036854775808:
                                        switch (fi43 * p0) {
                                        case 32767:
                                            var v59 uint8 = (v56 + v56)
                                            _ = v59
                                        default:
                                            switch "\n\t" {
                                            case "hello":
                                                if false {
                                                    if (!((false || (genFn0(1000, "\n\t") <= p0)) || true)) {
                                                        v56 = (v56 - v56)
                                                        var v60 string = "foo bar"
                                                        if false {
                                                            v60 = ("\n\t" + v60)
                                                            var v61 uint32 = 595
                                                            fmt.Println(v56)
                                                            _ = v61
                                                        }
                                                        _ = v60
                                                    }
                                                }
                                            case "a":
                                                if (!false) {
                                                    v62 := GenS4{F0: v56, F1: "foo bar", F2: -9223372036854775808}
                                                    _ = v62
                                                }
                                            case "\n\t":
                                                v57 = append(v57, fi42)
                                            }
                                        }
                                    }
                                } else {
                                    if true {
                                        if true {
                                            fmt.Println(p1)
                                            if len(v57) >= 2 {
                                                v57 = v57[1:len(v57)]
                                            }
                                            p1 += (^(p1))
                                            p1 += (^(p1))
                                        }
                                        if (!((("" + ("hello" + ("\n\t" + ("foo bar" + "\n\t")))) + "a") == ("foo bar" + ("foo bar" + "")))) {
                                            var v63 int64 = 0
                                            v57 = append(v57, 4294967295)
                                            var v64 float64 = 1e10
                                            _ = v63
                                            _ = v64
                                        } else {
                                            if len(v57) > 0 {
                                                v57[0] = (fi42 % (((^(p0))) | 1))
                                            }
                                            if len(v57) > 0 {
                                                v57[0] = (fi43 & 32768)
                                            }
                                            fmt.Println(fi43)
                                        }
                                        fmt.Println(v56, p1)
                                    } else {
                                        var v65 string = ("a" + "foo bar")
                                        _ = v65
                                    }
                                    var v66 uint64 = 18446744073709551615
                                    _ = v66
                                }
                                _ = v56
                                _ = v57
                            }
                            switch "foo bar" {
                            case "":
                                var v67 uint16 = 2
                                _ = v67
                            case "foo bar":
                                if true {
                                    var v68 uint32 = 2147483647
                                    var v69 int16 = 0
                                    v70 := []int{2147483647, genFn0(v69, ("a" + "hello"))}
                                    if len(v70) > 0 {
                                        v70[0] = genFn0(v69, "hello")
                                    }
                                    _ = v68
                                    _ = v69
                                    _ = v70
                                }
                            default:
                                var v71 float32 = -2.25
                                _ = v71
                            }
                        case "\n\t":
                            switch ("\n\t" + "") {
                            case "\n\t":
                                var v72 int16 = 32767
                                _ = v72
                            }
                        }
                    }
                    fmt.Println(fi42, p0)
                    _ = v47
                }
                var v73 bool = ((^(p1)) < (p1 / ((-128) | 1)))
                _ = fi43
                _ = v73
            }
            p0 += p0
            _ = fi42
        }
    case "":
        var v74 float32 = ((float32(1e10) - 1e10) / float32((1e10 - 0.1)))
        _ = v74
    case "hello":
        v75 := []int{(p0 ^ (p0 >> ((18446744073709551615) & 63))), (p0 ^ 128)}
        _ = v75
    default:
        switch ((("\n\t" + ("a" + "hello")) + "foo bar") + ("foo bar" + "a")) {
        case "":
            for fi76 := 0; fi76 < 18; fi76++ {
                v77 := []int{(p0 - genFn0(1000, "foo bar")), p0, (p0 & (^(fi76))), p0}
                _ = fi76
                _ = v77
            }
        case "\n\t":
            var v78 uint32 = 0
            _ = v78
        case "foo bar":
            var v79 uint64 = 4294967295
            _ = v79
        default:
            v80 := []int{-9223372036854775808}
            _ = v80
        }
    }
    if false {
        for fi81 := 0; fi81 < 10; fi81++ {
            v82 := GenS6{F0: -1, F1: 42, F2: true}
            _ = fi81
            _ = v82
        }
    } else {
        for fi83 := 0; fi83 < 15; fi83++ {
            v84 := GenS4{F0: 128, F1: "a", F2: -9223372036854775808}
            _ = fi83
            _ = v84
        }
    }
    _ = p0
    _ = p1
    return true
}

func main() {
    v1 := make(map[uint32]int)
    var v2 bool = (2147483647 != -690)
    v3 := make(map[int]bool)
    var v85 uint8 = 234
    v2 = true
    v1[23] = 2147483647
    v85 -= (v85 & v85)
    var v86 bool = (!(v2 || false))
    var v87 int16 = 655
    if v2 {
        v88 := GenS1{F0: 18446744073709551615, F1: 4294967295}
        var v89 string = "\n\t"
        _ = v88
        _ = v89
    } else {
        if true {
            for fi90 := 0; fi90 < 17; fi90++ {
                for fi91 := 0; fi91 < 11; fi91++ {
                    switch "\n\t" {
                    case "":
                        v92 := []int{(fi91 + fi90), genFn0(-32768, "foo bar")}
                        _ = v92
                    case "\n\t":
                        v87 += (-(v87))
                    }
                    v85 = 156
                    var v93 uint16 = 656
                    _ = fi91
                    _ = v93
                }
                _ = fi90
            }
            var v94 int32 = -993
            _ = v94
        } else {
            var v95 bool = false
            if true {
                if (123 >= (v85 + 163)) {
                    v95 = (-194 == -909)
                    var v96 uint = 42
                    for fi97 := 0; fi97 < 7; fi97++ {
                        v98 := GenS7{F0: 1}
                        _ = fi97
                        _ = v98
                    }
                    _ = v96
                }
                _ = len(v3)
                var v99 int32 = 0
                _ = v99
            } else {
                v100 := GenS3{F0: 1}
                v101 := []int{-360}
                _ = v100
                _ = v101
            }
            _ = len(v1)
            _ = v95
        }
    }
    v102 := GenS4{F0: (v85 ^ 255), F1: ("" + "a"), F2: -276}
    v103 := []int{4294967295}
    v104 := GenS2{F0: 42}
    v86 = (18446744073709551615 != 4294967295)
    fmt.Println(v85, v85)
    v105 := GenS5{F0: 21}
    if (((42 >= 0) && v2) || true) {
        if v2 {
            for fi106 := 0; fi106 < 8; fi106++ {
                if v2 {
                    var v107 bool = v2
                    fmt.Println(v2, v87)
                    v108 := GenS5{F0: 1}
                    _ = v107
                    _ = v108
                } else {
                    if (v85 != v85) {
                        for fi109 := 0; fi109 < 5; fi109++ {
                            v86 = (!(2 == 1))
                            switch fi106 {
                            case -292:
                                switch "foo bar" {
                                case "\n\t":
                                    var v110 bool = (2147483647 < -601)
                                    _ = v110
                                }
                            case 127:
                                delete(v3, fi109)
                            default:
                                _ = len(v103)
                            }
                            fmt.Println(fi106, v2, v87)
                            var v111 uint32 = 915
                            _ = fi109
                            _ = v111
                        }
                        for fi112 := 0; fi112 < 7; fi112++ {
                            v85 += (v85 - 255)
                            v85 = v85
                            _ = fi112
                        }
                        v113 := GenS5{F0: (v85 << ((566) & 63))}
                        _ = v113
                    }
                }
                if v2 {
                    v2 = v86
                }
                var v114 int64 = 245
                _ = fi106
                _ = v114
            }
            _ = len(v1)
        }
        v87 = 428
        if v2 {
            v85 -= v85
            fmt.Println(v85, v86)
        }
    } else {
        if (!v86) {
            var v115 string = "a"
            _ = len(v1)
            _ = v115
        } else {
            for fi116 := 0; fi116 < 1; fi116++ {
                if len(v103) >= 2 {
                    v103 = v103[1:len(v103)]
                }
                _ = fi116
            }
        }
        for fi117 := 0; fi117 < 2; fi117++ {
            if v2 {
                switch "\n\t" {
                case "hello":
                    if (-2.25 <= 0.0) {
                        for fi118 := 0; fi118 < 5; fi118++ {
                            v87 += (v87 >> ((0) & 63))
                            _ = fi118
                        }
                        var v119 string = "foo bar"
                        v87 += v87
                        _ = v119
                    }
                }
            }
            _ = fi117
        }
        fmt.Println(v86, v2)
    }
    if (!(-2.25 > (-2.25 / 1e10))) {
        for fi120 := 0; fi120 < 6; fi120++ {
            if v86 {
                for fi121 := 0; fi121 < 18; fi121++ {
                    var v122 string = "hello"
                    switch v122 {
                    case "\n\t":
                        if (v2 || (1 < 9223372036854775807)) {
                            var v123 uint8 = v85
                            var v124 uint64 = 18446744073709551615
                            var v125 int64 = 10
                            _ = v123
                            _ = v124
                            _ = v125
                        }
                        v126 := GenS8{F0: v87, F1: -1, F2: "a"}
                        _ = v126
                    default:
                        if true {
                            var v127 int16 = -32768
                            _ = v127
                        } else {
                            var v128 int16 = v87
                            v129 := GenS8{F0: (v128 / (((v128 + -32768)) | 1)), F1: 78, F2: v122}
                            _ = v128
                            _ = v129
                        }
                    }
                    if v86 {
                        v87 = (v87 * (v87 % ((v87) | 1)))
                    } else {
                        var v130 bool = v2
                        _ = v130
                    }
                    v131 := GenS6{F0: 1, F1: 0, F2: v86}
                    fmt.Println(v122, fi121)
                    _ = fi121
                    _ = v122
                    _ = v131
                }
            } else {
                v132 := GenS9{F0: 255, F1: 1, F2: fi120}
                var v133 uint8 = 0
                _ = v132
                _ = v133
            }
            v2 = false
            v86 = true
            _ = fi120
        }
        v87 -= (v87 + v87)
    } else {
        if (4294967295 > 2) {
            v85 -= v85
        }
        if (1.5 == ((-2.25 - 0.0) - 1e10)) {
            if (!(65535 < 524)) {
                switch "\n\t" {
                case "":
                    v86 = (v86 || (!v2))
                case "hello":
                    if v2 {
                        v86 = true
                        v134 := v3[-9223372036854775808]
                        var v135 uint16 = 42
                        _ = v134
                        _ = v135
                    } else {
                        v136 := GenS3{F0: -128}
                        v137 := GenS10{F0: 42, F1: -361}
                        v138 := GenS11{F0: -713, F1: 527, F2: ("" + "a")}
                        _ = v136
                        _ = v137
                        _ = v138
                    }
                case "\n\t":
                    for fi139 := 0; fi139 < 13; fi139++ {
                        v3[fi139] = true
                        for fi140 := 0; fi140 < 2; fi140++ {
                            v86 = false
                            var v141 int16 = (v87 | -32768)
                            _ = fi140
                            _ = v141
                        }
                        _ = fi139
                    }
                }
            } else {
                if false {
                    var v142 bool = (4294967295 > 18446744073709551615)
                    for fi143 := 0; fi143 < 16; fi143++ {
                        switch "" {
                        case "foo bar":
                            v144 := GenS4{F0: (v85 / (((v85 << ((0) & 63))) | 1)), F1: "\n\t", F2: 9223372036854775807}
                            _ = v144
                        }
                        for fi145 := 0; fi145 < 12; fi145++ {
                            var v146 bool = false
                            var v147 uint32 = 4294967295
                            switch 32767 {
                            case 439:
                                if (834 == 656) {
                                    if ((!v2) && (1.5 > (-2.25 - 0.1))) {
                                        var v148 int8 = 59
                                        _ = v148
                                    }
                                    var v149 int16 = v87
                                    v150 := GenS12{F0: "\n\t", F1: v147, F2: 255}
                                    var v151 uint = 255
                                    v152 := GenS11{F0: -2147483648, F1: (v147 & v147), F2: ""}
                                    _ = v149
                                    _ = v150
                                    _ = v151
                                    _ = v152
                                } else {
                                    v86 = v146
                                }
                            case 0:
                                v86 = ((((v147 >= (v147 << ((0) & 63))) && (v146 || v142)) || (!true)) || (-1 != 42))
                            default:
                                var v153 uint32 = (v147 | (^(v147)))
                                _ = v153
                            }
                            _ = fi145
                            _ = v146
                            _ = v147
                        }
                        _ = len(v1)
                        _ = fi143
                    }
                    var v154 float64 = (0.0 - (0.1 * 3.14159))
                    _ = v142
                    _ = v154
                } else {
                    fmt.Println(v86, v86, v87)
                    var v155 int8 = 64
                    _ = v155
                }
            }
        } else {
            var v156 bool = false
            for fi157 := 0; fi157 < 2; fi157++ {
                v86 = v156
                if false {
                    var v158 int16 = (v87 | -32768)
                    var v159 int = (fi157 << ((65535) & 63))
                    _ = v158
                    _ = v159
                } else {
                    var v160 float32 = ((1.5 - float32(((((1.5 / 3.14159) * 0.1) * 1.5) * ((((-2.25 + (0.1 * 0.0)) * 0.0) / -2.25) / 1e10)))) / 3.14159)
                    for fi161 := 0; fi161 < 4; fi161++ {
                        v87 += int16(v85)
                        _ = fi161
                    }
                    var v162 uint = 424
                    _ = v160
                    _ = v162
                }
                for fi163 := 0; fi163 < 1; fi163++ {
                    v3[genFn0(1000, "a")] = ((!(v85 != (v85 ^ v85))) && v2)
                    v164 := GenS8{F0: v87, F1: -1, F2: (("\n\t" + "foo bar") + "")}
                    _ = fi163
                    _ = v164
                }
                _ = fi157
            }
            var v165 int32 = 2147483647
            _ = v156
            _ = v165
        }
    }
    v2 = true
    var v166 uint16 = 65535
    fmt.Println(v2)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v103:", len(v103), v103)
    fmt.Println("v166:", v166)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v85:", v85)
    fmt.Println("v86:", v86)
    fmt.Println("v87:", v87)
    _ = v1
    _ = v102
    _ = v103
    _ = v104
    _ = v105
    _ = v166
    _ = v2
    _ = v3
    _ = v85
    _ = v86
    _ = v87
}
