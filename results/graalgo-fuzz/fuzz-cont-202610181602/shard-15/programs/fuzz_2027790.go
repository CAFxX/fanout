package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
type GenS2 struct { F0 uint8; F1 string; F2 uint64 }
type GenS3 struct { F0 uint32 }
type GenS4 struct { F0 uint; F1 int32; F2 int16 }
type GenS5 struct { F0 bool; F1 int32; F2 int8 }
type GenS6 struct { F0 uint64; F1 string; F2 uint8 }
type GenS7 struct { F0 int64 }
type GenS8 struct { F0 uint8; F1 int32; F2 int32 }
type GenS9 struct { F0 uint; F1 uint8 }
type GenS10 struct { F0 uint8; F1 uint64; F2 int64 }

func main() {
    v1 := make(map[string]string)
    var v2 uint16 = 0
    var v3 int8 = 0
    var v4 int32 = -2147483648
    var v5 string = "\n\t"
    var v6 int16 = 42
    for fi7 := 0; fi7 < 15; fi7++ {
        v6 = v6
        v8 := []int{}
        if (true || (!((v3 >> ((1) & 63)) == (v3 - (-(v3)))))) {
            if ((v6 << ((0) & 63)) < (v6 - (v6 ^ v6))) {
                if (42 >= 2) {
                    v9 := GenS1{F0: 42}
                    var v10 float32 = (1e10 / 1.0)
                    _ = v10
                    _ = v9
                } else {
                    if (!false) {
                        for fi11 := 0; fi11 < 4; fi11++ {
                            if len(v8) >= 2 {
                                v8 = v8[1:len(v8)]
                            }
                            v12 := GenS1{F0: 127}
                            var v13 string = ("foo bar" + ("hello" + (v5 + v5)))
                            fmt.Println(fi11, fi11, v5)
                            _ = fi11
                            _ = v12
                            _ = v13
                        }
                    } else {
                        switch 127 {
                        case -428:
                            var v14 float32 = (0.1 + (-2.25 + 1.5))
                            _ = v14
                        case 128:
                            var v15 int16 = (v6 << ((212) & 63))
                            _ = v15
                        case -491:
                            if (false || false) {
                                for fi16 := 0; fi16 < 5; fi16++ {
                                    v17 := GenS1{F0: 95}
                                    _ = fi16
                                    _ = v17
                                }
                                v18 := GenS2{F0: 2, F1: "\n\t", F2: 42}
                                _ = v18
                            }
                        }
                        switch "a" {
                        case "hello":
                            var v19 int8 = 42
                            _ = v19
                        case "\n\t":
                            fmt.Println(v2)
                        }
                        if true {
                            var v20 float32 = 1e10
                            if false {
                                if true {
                                    switch (fi7 << ((42) & 63)) {
                                    case -667:
                                        if len(v8) > 0 {
                                            v8[0] = fi7
                                        }
                                    }
                                    if true {
                                        v5 = v5
                                        var v21 int32 = v4
                                        _ = v21
                                    } else {
                                        if true {
                                            switch fi7 {
                                            case 175:
                                                var v22 int = fi7
                                                _ = v22
                                            case -2147483648:
                                                v5 = (v5 + v5)
                                            default:
                                                v23 := GenS1{F0: v3}
                                                _ = v23
                                            }
                                        }
                                        switch (fi7 >> ((2) & 63)) {
                                        case 2147483647:
                                            v8 = append(v8, fi7)
                                        }
                                    }
                                }
                                v3 -= 100
                            }
                            v24 := GenS2{F0: 176, F1: v5, F2: 42}
                            _ = v20
                            _ = v24
                        } else {
                            v25 := GenS2{F0: 255, F1: ((("a" + v5) + "\n\t") + ""), F2: 1}
                            if true {
                                v26 := GenS2{F0: 8, F1: (("a" + ("\n\t" + "foo bar")) + ((v5 + v5) + v5)), F2: 1}
                                v4 = (v4 % ((516) | 1))
                                _ = v26
                            } else {
                                v27 := GenS3{F0: 2147483648}
                                fmt.Println(v6, v6, v2)
                                switch fi7 {
                                case 4294967295:
                                    for fi28 := 0; fi28 < 8; fi28++ {
                                        var v29 uint32 = 2147483648
                                        var v30 int = (fi7 - (fi28 >> ((1) & 63)))
                                        var v31 bool = (false && true)
                                        _ = fi28
                                        _ = v29
                                        _ = v30
                                        _ = v31
                                    }
                                case 42:
                                    switch (("\n\t" + (((v5 + "foo bar") + (v5 + v5)) + ("foo bar" + ("" + v5)))) + v5) {
                                    case "hello":
                                        v6 = v6
                                    case "":
                                        for fi32 := 0; fi32 < 7; fi32++ {
                                            v6 += v6
                                            fmt.Println(fi32, v3)
                                            v33 := GenS3{F0: 54}
                                            _ = fi32
                                            _ = v33
                                        }
                                    }
                                case 128:
                                    _ = len(v1)
                                default:
                                    for fi34 := 0; fi34 < 2; fi34++ {
                                        v35 := GenS3{F0: 2}
                                        v36 := v1[v5]
                                        _ = fi34
                                        _ = v35
                                        _ = v36
                                    }
                                }
                                _ = v27
                            }
                            _ = v25
                        }
                    }
                    switch v5 {
                    case "\n\t":
                        var v37 int64 = 677
                        _ = v37
                    default:
                        fmt.Println(v5)
                    }
                    _ = len(v8)
                }
                _ = len(v8)
                for fi38 := 0; fi38 < 19; fi38++ {
                    var v39 bool = false
                    _ = fi38
                    _ = v39
                }
            }
            fmt.Println(v2)
        } else {
            var v40 uint8 = 128
            var v41 uint32 = 2147483648
            v42 := GenS2{F0: v40, F1: v5, F2: 835}
            _ = v40
            _ = v41
            _ = v42
        }
        _ = fi7
        _ = v8
    }
    var v43 int64 = 42
    v44 := v1[v5]
    v4 = v4
    for fi45 := 0; fi45 < 8; fi45++ {
        var v46 int64 = (v43 % ((-1) | 1))
        v47 := GenS2{F0: 16, F1: ((("foo bar" + "hello") + (v44 + v5)) + v44), F2: 42}
        _ = fi45
        _ = v46
        _ = v47
    }
    if false {
        v4 = (v4 + v4)
        v6 = (^(v6))
    }
    switch ("hello" + "") {
    case "foo bar":
        if true {
            v43 = v43
            switch ("" + (v5 + (v44 + v5))) {
            case "a":
                if false {
                    switch "a" {
                    case "foo bar":
                        v6 -= (v6 + int16(v4))
                    case "\n\t":
                        if true {
                            var v48 float32 = 0.0
                            _ = v48
                        }
                    }
                    for fi49 := 0; fi49 < 13; fi49++ {
                        v50 := GenS4{F0: 4294967295, F1: -1, F2: v6}
                        var v51 uint16 = (v2 & (v2 * v2))
                        _ = fi49
                        _ = v50
                        _ = v51
                    }
                    var v52 uint8 = 255
                    _ = v52
                }
            case "":
                var v53 int16 = v6
                _ = v53
            case "foo bar":
                for fi54 := 0; fi54 < 8; fi54++ {
                    v4 += v4
                    v5 = "a"
                    _ = fi54
                }
            default:
                v43 = -1
            }
        }
    }
    var v55 uint16 = (v2 >> ((18446744073709551615) & 63))
    if (128 >= 2) {
        for fi56 := 0; fi56 < 6; fi56++ {
            if true {
                v44 = (v5 + (v44 + "a"))
                v55 = (v2 << ((18446744073709551615) & 63))
                for fi57 := 0; fi57 < 16; fi57++ {
                    fmt.Println(v43)
                    if (!(!false)) {
                        _ = len(v1)
                        var v58 uint16 = v55
                        switch "a" {
                        case "foo bar":
                            var v59 int = (-(fi56))
                            _ = v59
                        }
                        _ = v58
                    } else {
                        v60 := GenS1{F0: (v3 << ((2) & 63))}
                        v61 := GenS1{F0: v3}
                        v62 := []int{(fi56 * fi57), (fi56 + fi57), (fi57 & fi56), (fi56 + fi56)}
                        _ = v60
                        _ = v61
                        _ = v62
                    }
                    _ = fi57
                }
            } else {
                var v63 int32 = (^(v4))
                var v64 string = v5
                var v65 uint = 725
                _ = v63
                _ = v64
                _ = v65
            }
            switch v44 {
            case "":
                for fi66 := 0; fi66 < 2; fi66++ {
                    if true {
                        switch ((v44 + v44) + "hello") {
                        case "foo bar":
                            v67 := GenS1{F0: 22}
                            _ = v67
                        default:
                            v1[v5] = v44
                        }
                        var v68 int64 = v43
                        var v69 int64 = (v68 * (v68 | v68))
                        _ = v68
                        _ = v69
                    }
                    v70 := GenS3{F0: 2}
                    if ((-(v4)) != (^(v4))) {
                        if (538 >= 698) {
                            switch (v5 + "hello") {
                            case "\n\t":
                                fmt.Println(v6, v44)
                            }
                        }
                        var v71 uint16 = (v2 >> ((1) & 63))
                        var v72 int32 = v4
                        v2 = v55
                        _ = v71
                        _ = v72
                    } else {
                        var v73 uint32 = 65535
                        _ = v73
                    }
                    _ = fi66
                    _ = v70
                }
            case "hello":
                v74 := GenS5{F0: ((int64(v4) >= (^(v43))) && true), F1: v4, F2: (v3 / ((v3) | 1))}
                _ = v74
            default:
                if (!(fi56 == (fi56 & 32768))) {
                    v75 := GenS6{F0: 18446744073709551615, F1: (((((v5 + v44) + v44) + (v5 + (((v44 + ((v44 + "\n\t") + v5)) + v5) + (v44 + "")))) + ((v44 + "foo bar") + ((v5 + (((v5 + (v5 + v5)) + v44) + "a")) + (("hello" + (v5 + ((v5 + "foo bar") + v44))) + v44)))) + "foo bar"), F2: 216}
                    v5 = (("foo bar" + v5) + v44)
                    var v76 float32 = -2.25
                    _ = v75
                    _ = v76
                } else {
                    for fi77 := 0; fi77 < 9; fi77++ {
                        v78 := GenS7{F0: 1}
                        _ = fi77
                        _ = v78
                    }
                }
            }
            v79 := GenS5{F0: (false || (v3 > (v3 - int8(v6)))), F1: v4, F2: v3}
            _ = fi56
            _ = v79
        }
        v44 = (v5 + ((v44 + v5) + "\n\t"))
    }
    v1[("a" + v5)] = "a"
    for fi80 := 0; fi80 < 7; fi80++ {
        v3 = v3
        switch v44 {
        case "a":
            switch (fi80 >> ((397) & 63)) {
            case -2147483648:
                var v81 bool = false
                _ = v81
            case -105:
                var v82 bool = true
                _ = v82
            default:
                switch fi80 {
                case 42:
                    switch v5 {
                    case "hello":
                        v83 := []int{}
                        _ = v83
                    case "a":
                        if false {
                            if (!true) {
                                v84 := []int{706}
                                if (((!(1.5 != 3.14159)) && true) || true) {
                                    switch v5 {
                                    case "":
                                        for fi85 := 0; fi85 < 3; fi85++ {
                                            var v86 int8 = int8(v4)
                                            _ = fi85
                                            _ = v86
                                        }
                                    default:
                                        var v87 uint8 = 59
                                        _ = v87
                                    }
                                    for fi88 := 0; fi88 < 5; fi88++ {
                                        var v89 bool = (4294967295 < 4294967295)
                                        v3 = v3
                                        var v90 uint = 776
                                        _ = fi88
                                        _ = v89
                                        _ = v90
                                    }
                                }
                                v44 = v44
                                _ = v84
                            } else {
                                v91 := v1[(((((v44 + "a") + v5) + (v5 + v44)) + v5) + (("" + v44) + ("hello" + "hello")))]
                                var v92 int32 = v4
                                var v93 uint = 1
                                _ = v91
                                _ = v92
                                _ = v93
                            }
                            switch fi80 {
                            case 9223372036854775807:
                                v94 := v1[((v44 + v5) + "\n\t")]
                                _ = v94
                            case 1:
                                v55 = v2
                            }
                        } else {
                            fmt.Println(v5, v4)
                            fmt.Println(v4, v2, fi80)
                        }
                    }
                case -431:
                    var v95 int = (fi80 - (fi80 + fi80))
                    _ = v95
                case 4294967295:
                    v3 += (^(v3))
                default:
                    _ = len(v1)
                }
            }
        case "hello":
            if true {
                var v96 int64 = (v43 >> ((uint(v55)) & 63))
                fmt.Println(fi80)
                v97 := []int{(fi80 + fi80), (fi80 % (((fi80 - fi80)) | 1)), fi80, (^(fi80))}
                _ = v96
                _ = v97
            } else {
                var v98 float64 = 3.14159
                if false {
                    if ((v43 % ((v43) | 1)) >= (v43 * (^(v43)))) {
                        var v99 int64 = (v43 ^ int64(v4))
                        _ = v99
                    }
                } else {
                    var v100 uint32 = 0
                    v6 -= (v6 & (v6 / (((v6 << ((2) & 63))) | 1)))
                    v55 -= v55
                    _ = v100
                }
                _ = v98
            }
        default:
            var v101 string = v44
            _ = v101
        }
        if (v4 == v4) {
            switch 9223372036854775807 {
            case -2147483648:
                v44 = v44
            case -890:
                var v102 uint16 = v2
                _ = v102
            default:
                v103 := GenS3{F0: 2147483648}
                _ = v103
            }
        }
        _ = fi80
    }
    if true {
        if false {
            v44 = "hello"
            v3 = v3
            v104 := GenS5{F0: false, F1: (v4 * v4), F2: (v3 ^ (v3 % ((v3) | 1)))}
            _ = v104
        }
        v105 := GenS8{F0: 255, F1: -1, F2: int32(v43)}
        v106 := GenS1{F0: (v3 | 1)}
        _ = v105
        _ = v106
    }
    for fi107 := 0; fi107 < 4; fi107++ {
        if true {
            if false {
                switch "hello" {
                case "\n\t":
                    v43 += 0
                case "":
                    if true {
                        v3 += 100
                        for fi108 := 0; fi108 < 1; fi108++ {
                            v55 -= 665
                            var v109 uint8 = 147
                            v4 -= 0
                            var v110 bool = (82 <= v109)
                            switch (-(fi108)) {
                            case -2147483648:
                                switch (fi107 % (((^(fi107))) | 1)) {
                                case 2147483647:
                                    if (float32((-2.25 + (-2.25 * ((1e10 / 1.0) * 1e10)))) >= 1.5) {
                                        if (v6 < (v6 | (v6 % (((v6 << ((0) & 63))) | 1)))) {
                                            v4 += (v4 + v4)
                                            var v111 string = ((("" + (v5 + "foo bar")) + (v44 + v44)) + ("hello" + "a"))
                                            _ = v111
                                        } else {
                                            switch (fi107 * fi107) {
                                            case -343:
                                                v4 = v4
                                                v109 = v109
                                                v109 = 136
                                            }
                                            v112 := GenS9{F0: 97, F1: (v109 >> ((426) & 63))}
                                            var v113 uint64 = 0
                                            var v114 int64 = v43
                                            _ = v112
                                            _ = v113
                                            _ = v114
                                        }
                                    } else {
                                        v115 := v1["a"]
                                        var v116 int64 = 1
                                        _ = len(v1)
                                        if (0.1 != -2.25) {
                                            v109 = (v109 / ((1) | 1))
                                            if (!false) {
                                                if (v4 != (^(v4))) {
                                                    var v117 float64 = (float64((float32(0.0) + ((0.1 * -2.25) / 1e10))) - (1e10 + 3.14159))
                                                    if ((!v110) || v110) {
                                                        v118 := GenS5{F0: true, F1: v4, F2: (v3 ^ (v3 & v3))}
                                                        if v110 {
                                                            v119 := GenS6{F0: 4294967295, F1: v115, F2: (v109 & v109)}
                                                            fmt.Println(v43, v3)
                                                            _ = v119
                                                        } else {
                                                            v43 = v116
                                                        }
                                                        switch (fi107 << ((uint(v55)) & 63)) {
                                                        case 199:
                                                            var v120 float32 = (float32((3.14159 / -2.25)) / float32(v117))
                                                            _ = v120
                                                        case 128:
                                                            fmt.Println(fi108, v2, v116)
                                                        case 127:
                                                            v110 = ((1e10 > (v117 - v117)) && (42 != 4294967295))
                                                        }
                                                        _ = v118
                                                    } else {
                                                        fmt.Println(v55, fi107)
                                                        var v121 uint16 = (v55 & v55)
                                                        _ = v121
                                                    }
                                                    switch -9223372036854775808 {
                                                    case -778:
                                                        if false {
                                                            var v122 string = v44
                                                            var v123 int32 = (v4 + v4)
                                                            _ = v122
                                                            _ = v123
                                                        }
                                                    case 42:
                                                        v116 += (v116 & (-(v116)))
                                                    case -633:
                                                        if (true || v110) {
                                                            switch (v44 + v115) {
                                                            case "hello":
                                                                v44 = (("" + (((v5 + (((("\n\t" + v44) + ("hello" + v115)) + v44) + v5)) + (v44 + (v44 + ""))) + v5)) + v115)
                                                            case "":
                                                                v124 := GenS2{F0: (v109 - v109), F1: ("\n\t" + v5), F2: 0}
                                                                _ = v124
                                                            case "a":
                                                                v43 = 214
                                                            default:
                                                                v4 += (v4 + 2147483647)
                                                            }
                                                            var v125 int = int(v43)
                                                            delete(v1, v5)
                                                            v109 = v109
                                                            _ = v125
                                                        } else {
                                                            var v126 uint64 = 18446744073709551615
                                                            var v127 string = v115
                                                            fmt.Println(fi108)
                                                            _ = v126
                                                            _ = v127
                                                        }
                                                    }
                                                    _ = v117
                                                } else {
                                                    v128 := GenS1{F0: v3}
                                                    fmt.Println(v55, fi107)
                                                    _ = v128
                                                }
                                                v109 = 127
                                                fmt.Println(v44, v110, v6)
                                            } else {
                                                if (3.14159 >= 1.5) {
                                                    switch v5 {
                                                    case "\n\t":
                                                        v129 := v1[v5]
                                                        _ = v129
                                                    default:
                                                        v5 = ""
                                                    }
                                                } else {
                                                    v44 = v5
                                                }
                                                v55 -= 1
                                            }
                                        } else {
                                            fmt.Println(v115, v5, v4)
                                        }
                                        v130 := GenS8{F0: v109, F1: v4, F2: 42}
                                        _ = v115
                                        _ = v116
                                        _ = v130
                                    }
                                case 32768:
                                    if (false && v110) {
                                        v131 := []int{(fi107 & 2147483647), int(v6), (fi108 % (((fi108 >> ((2) & 63))) | 1)), (fi107 << ((65535) & 63))}
                                        v1[((v44 + v44) + v5)] = "hello"
                                        switch fi108 {
                                        case 2147483647:
                                            var v132 uint64 = 2
                                            _ = v132
                                        }
                                        _ = v131
                                    }
                                case -1:
                                    var v133 float32 = ((0.1 * 0.1) * 3.14159)
                                    _ = v133
                                }
                            case 863:
                                v134 := GenS5{F0: (v110 && (!v110)), F1: (v4 & v4), F2: (v3 / ((v3) | 1))}
                                _ = v134
                            default:
                                v44 = v44
                                v2 += v55
                                if v110 {
                                    v135 := GenS8{F0: (v109 * v109), F1: 270, F2: v4}
                                    v136 := GenS6{F0: 839, F1: ((v5 + (v5 + "")) + v44), F2: (v109 + 221)}
                                    v137 := []int{(fi108 & fi108), fi108}
                                    _ = v135
                                    _ = v136
                                    _ = v137
                                }
                            }
                            if true {
                                v6 = v6
                            }
                            _ = fi108
                            _ = v109
                            _ = v110
                        }
                        var v138 bool = true
                        _ = v138
                    } else {
                        var v139 float32 = (1e10 - -2.25)
                        fmt.Println(v2, v139)
                        v6 = v6
                        _ = v139
                    }
                default:
                    v140 := []int{int(v2)}
                    _ = v140
                }
            } else {
                v141 := []int{}
                delete(v1, v5)
                v2 = v55
                _ = v141
            }
            if true {
                for fi142 := 0; fi142 < 17; fi142++ {
                    v55 += (v2 - 261)
                    v143 := GenS4{F0: 255, F1: int32(v2), F2: -231}
                    switch (fi142 * (-(fi107))) {
                    case 787:
                        switch (fi107 + 1) {
                        case -722:
                            fmt.Println(v5, v44)
                        case -1:
                            var v144 int16 = v6
                            _ = v144
                        }
                    }
                    v44 = v44
                    _ = fi142
                    _ = v143
                }
                switch fi107 {
                case 4294967295:
                    for fi145 := 0; fi145 < 10; fi145++ {
                        if (float64(1e10) < 1.5) {
                            fmt.Println(v5)
                            v146 := []int{1, fi145, fi107}
                            delete(v1, ("hello" + v5))
                            _ = v146
                        } else {
                            v2 = (^(v2))
                        }
                        _ = fi145
                    }
                case -2147483648:
                    var v147 int64 = 42
                    _ = v147
                case 2147483647:
                    v148 := GenS1{F0: (v3 << ((2) & 63))}
                    _ = v148
                default:
                    for fi149 := 0; fi149 < 16; fi149++ {
                        v150 := []int{fi149, (-(fi149)), 128}
                        v151 := GenS2{F0: 42, F1: v44, F2: 18446744073709551615}
                        _ = fi149
                        _ = v150
                        _ = v151
                    }
                }
                v152 := GenS10{F0: 42, F1: 42, F2: (v43 | v43)}
                _ = v152
            }
        } else {
            v153 := []int{(^(fi107)), fi107, fi107, fi107}
            v3 = v3
            _ = v153
        }
        v4 += -1
        var v154 float32 = float32(3.14159)
        _ = fi107
        _ = v154
    }
    if (true && false) {
        var v155 uint = 476
        _ = v155
    }
    for fi156 := 0; fi156 < 16; fi156++ {
        var v157 uint32 = 322
        _ = fi156
        _ = v157
    }
    _ = len(v1)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v43:", v43)
    fmt.Println("v44:", v44)
    fmt.Println("v5:", v5)
    fmt.Println("v55:", v55)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v43
    _ = v44
    _ = v5
    _ = v55
    _ = v6
}
