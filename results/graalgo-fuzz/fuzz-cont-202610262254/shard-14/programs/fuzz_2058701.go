package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 uint8; F2 uint }
type GenS2 struct { F0 int; F1 uint8; F2 int32 }
type GenS3 struct { F0 int; F1 uint32 }
type GenS4 struct { F0 string }
type GenS5 struct { F0 int8; F1 uint16 }
type GenS6 struct { F0 uint8; F1 int64; F2 uint }
type GenS7 struct { F0 uint; F1 uint8; F2 bool }
type GenS8 struct { F0 uint32; F1 uint16; F2 uint64 }
type GenS9 struct { F0 uint8; F1 int }
type GenS10 struct { F0 int8; F1 int; F2 uint32 }
type GenS11 struct { F0 uint }
type GenS12 struct { F0 int16; F1 int }
type GenS13 struct { F0 int64; F1 uint16; F2 uint64 }
type GenS14 struct { F0 string; F1 int16; F2 int8 }
type GenS15 struct { F0 int32 }
type GenS16 struct { F0 uint32; F1 uint64; F2 uint }
type GenS17 struct { F0 uint16; F1 int8; F2 int }
type GenS18 struct { F0 uint8; F1 int32; F2 int64 }
type GenS19 struct { F0 uint32; F1 int64 }
type GenS20 struct { F0 uint8; F1 int8 }
type GenS21 struct { F0 uint8; F1 uint; F2 uint8 }
type GenS22 struct { F0 int64; F1 int8 }
func genFn0(p0 int, p1 string, p2 uint) bool {
    if (0 == 100) {
        for fi4 := 0; fi4 < 2; fi4++ {
            switch (p0 / (((p0 + 32767)) | 1)) {
            case 1:
                var v5 float64 = (0.1 / ((1e10 + (((float64(-2.25) + (3.14159 - (-2.25 + 0.1))) * (0.0 / (0.1 - 0.0))) * float64((0.0 * 0.0)))) - (3.14159 / 1.0)))
                _ = v5
            case -1:
                v6 := GenS1{F0: false, F1: 108, F2: p2}
                _ = v6
            case 4294967295:
                fmt.Println(p0, fi4)
            default:
                v7 := GenS2{F0: fi4, F1: 1, F2: -2147483648}
                _ = v7
            }
            if true {
                for fi8 := 0; fi8 < 7; fi8++ {
                    switch p1 {
                    case "a":
                        var v9 int8 = 119
                        var v10 int8 = (-(v9))
                        var v11 uint8 = 1
                        _ = v10
                        _ = v11
                        _ = v9
                    default:
                        p0 = p0
                    }
                    switch ("a" + "") {
                    case "foo bar":
                        p2 = (^(p2))
                        v12 := []int{(fi4 - (p0 & -959)), (fi4 + (p0 | fi4)), 4294967295}
                        var v13 int8 = 1
                        if true {
                            if false {
                                p1 = ((p1 + "") + (p1 + ""))
                            }
                            p1 = p1
                        } else {
                            p1 = (p1 + ("a" + (p1 + (p1 + p1))))
                            var v14 uint16 = 42
                            p1 = p1
                            var v15 bool = true
                            v14 += 1
                            _ = v14
                            _ = v15
                        }
                        _ = v12
                        _ = v13
                    case "\n\t":
                        v16 := []int{42}
                        _ = v16
                    case "a":
                        switch p1 {
                        case "a":
                            var v17 string = p1
                            _ = v17
                        case "hello":
                            fmt.Println(p0)
                        case "":
                            v18 := []int{p0, p0, 589, (fi8 >> ((4294967295) & 63))}
                            _ = v18
                        }
                    default:
                        if (true || true) {
                            p0 += (fi4 / ((589) | 1))
                            v19 := GenS2{F0: fi4, F1: 128, F2: 1}
                            var v20 uint64 = 0
                            _ = v19
                            _ = v20
                        } else {
                            switch p1 {
                            case "hello":
                                v21 := GenS1{F0: (((!(false && (!((-2.25 < 3.14159) && (42 <= -430))))) && false) && false), F1: 2, F2: (p2 >> ((2) & 63))}
                                _ = v21
                            case "foo bar":
                                var v22 float64 = 0.1
                                _ = v22
                            default:
                                fmt.Println(fi4, fi4, p2)
                                if (!(-1 <= -718)) {
                                    p2 = p2
                                } else {
                                    var v23 uint32 = 4294967295
                                    v24 := []int{128, fi4}
                                    if false {
                                        p0 += 1
                                        v25 := GenS1{F0: (4294967295 > uint64(v23)), F1: 2, F2: p2}
                                        var v26 uint = 18446744073709551615
                                        if (!(!true)) {
                                            fmt.Println(v23)
                                            fmt.Println(fi8, v26)
                                        }
                                        _ = v25
                                        _ = v26
                                    } else {
                                        switch ("a" + "hello") {
                                        case "a":
                                            var v27 float64 = 1e10
                                            p0 = fi8
                                            _ = v27
                                        case "foo bar":
                                            v24 = append(v24, fi8)
                                        case "":
                                            if (!true) {
                                                if false {
                                                    var v28 string = (p1 + ("a" + (p1 + (p1 + ("hello" + (("a" + p1) + "hello"))))))
                                                    v29 := GenS3{F0: (fi4 >> ((4294967295) & 63)), F1: 2}
                                                    switch (fi8 ^ fi8) {
                                                    case 127:
                                                        var v30 uint16 = 831
                                                        _ = v30
                                                    case 0:
                                                        v24 = append(v24, fi4)
                                                    }
                                                    _ = v28
                                                    _ = v29
                                                } else {
                                                    if (true || ((1e10 - -2.25) <= 0.1)) {
                                                        var v31 int = 1
                                                        var v32 int32 = 42
                                                        p2 = (p2 & (p2 >> ((2) & 63)))
                                                        _ = v31
                                                        _ = v32
                                                    }
                                                }
                                                var v33 float32 = (0.1 / 1e10)
                                                var v34 bool = false
                                                _ = v33
                                                _ = v34
                                            }
                                        default:
                                            v35 := GenS4{F0: p1}
                                            v36 := GenS5{F0: 42, F1: 596}
                                            if true {
                                                var v37 string = p1
                                                v24 = append(v24, 906)
                                                _ = v37
                                            }
                                            var v38 float64 = 3.14159
                                            v38 = -2.25
                                            _ = v35
                                            _ = v36
                                            _ = v38
                                        }
                                        v23 += (v23 >> ((344) & 63))
                                    }
                                    _ = v23
                                    _ = v24
                                }
                                p1 = "foo bar"
                            }
                            if (-1 >= -1) {
                                var v39 int32 = -494
                                _ = v39
                            }
                            v40 := GenS6{F0: 44, F1: -355, F2: 118}
                            _ = v40
                        }
                        switch p1 {
                        case "hello":
                            p2 += p2
                        case "\n\t":
                            v41 := []int{}
                            _ = v41
                        }
                    }
                    if false {
                        switch p1 {
                        case "foo bar":
                            switch p1 {
                            case "\n\t":
                                if (p0 == (^(p0))) {
                                    switch p1 {
                                    case "foo bar":
                                        if (42 > -1) {
                                            p0 -= 128
                                        }
                                    }
                                } else {
                                    v42 := GenS7{F0: p2, F1: 128, F2: false}
                                    v43 := GenS8{F0: 2147483648, F1: 922, F2: 717}
                                    v44 := GenS5{F0: 0, F1: 2}
                                    v45 := []int{(p0 - -1), (p0 << (((p2 + p2)) & 63)), (fi8 - fi8)}
                                    v45 = append(v45, fi4)
                                    _ = v42
                                    _ = v43
                                    _ = v44
                                    _ = v45
                                }
                            }
                        case "hello":
                            p0 -= (fi4 * p0)
                        case "":
                            p1 = ("foo bar" + "a")
                        default:
                            var v46 uint64 = 4294967295
                            _ = v46
                        }
                        p0 = (fi4 - (^(fi4)))
                        var v47 uint = p2
                        _ = v47
                    }
                    fmt.Println(p1, fi8)
                    if false {
                        var v48 float64 = ((3.14159 + 3.14159) + float64(1.5))
                        if true {
                            v49 := []int{(fi8 << ((p2) & 63))}
                            var v50 uint16 = 1
                            _ = v49
                            _ = v50
                        } else {
                            v51 := []int{(p0 | p0), p0, p0}
                            v48 = (v48 - ((((-2.25 - ((1e10 - (v48 / -2.25)) / v48)) * -2.25) * float64(3.14159)) * (v48 * v48)))
                            _ = v51
                        }
                        _ = v48
                    }
                    _ = fi8
                }
                for fi52 := 0; fi52 < 19; fi52++ {
                    fmt.Println(fi4, fi4)
                    switch "a" {
                    case "foo bar":
                        var v53 bool = (1 < 42)
                        _ = v53
                    case "\n\t":
                        v54 := GenS2{F0: (fi4 | fi4), F1: 245, F2: -77}
                        _ = v54
                    case "hello":
                        switch p1 {
                        case "hello":
                            switch p0 {
                            case -2147483648:
                                p2 -= 0
                            }
                        case "foo bar":
                            switch fi4 {
                            case -1:
                                var v55 uint64 = 1
                                _ = v55
                            default:
                                switch fi4 {
                                case 4294967295:
                                    p2 -= p2
                                case 0:
                                    v56 := []int{(fi52 + p0), 2, (fi52 >> (((p2 + (p2 * p2))) & 63)), 127}
                                    _ = v56
                                }
                            }
                        case "":
                            var v57 uint32 = 2147483647
                            _ = v57
                        }
                        if (!true) {
                            p0 -= 32768
                            p0 = p0
                            switch "foo bar" {
                            case "":
                                fmt.Println(fi52, p2)
                            }
                            switch p1 {
                            case "\n\t":
                                v58 := []int{(fi52 + fi4), p0, fi4, fi52}
                                _ = v58
                            case "foo bar":
                                p0 += 128
                                p1 = p1
                            }
                            switch p1 {
                            case "":
                                v59 := GenS6{F0: 236, F1: 197, F2: (p2 << ((p2) & 63))}
                                p1 = ("a" + "")
                                _ = v59
                            case "\n\t":
                                fmt.Println(p0)
                            case "a":
                                p0 = fi4
                            }
                        }
                        switch fi4 {
                        case 127:
                            var v60 int32 = 2147483647
                            _ = v60
                        case 106:
                            if true {
                                var v61 int32 = 630
                                _ = v61
                            }
                        case -2147483648:
                            p1 = ((("" + p1) + "foo bar") + ("\n\t" + p1))
                        default:
                            var v62 float32 = 0.0
                            var v63 string = ((p1 + (("hello" + p1) + (("" + "\n\t") + ""))) + p1)
                            _ = v62
                            _ = v63
                        }
                        var v64 float64 = 3.14159
                        _ = v64
                    default:
                        if false {
                            p0 += fi4
                        }
                    }
                    _ = fi52
                }
                var v65 bool = true
                _ = v65
            } else {
                var v66 uint16 = 384
                p0 = (fi4 * p0)
                _ = v66
            }
            _ = fi4
        }
        var v67 int64 = -249
        _ = v67
    } else {
        if (-128 <= -764) {
            var v68 bool = (false || true)
            v69 := GenS4{F0: "hello"}
            var v70 uint = (p2 ^ p2)
            _ = v68
            _ = v69
            _ = v70
        } else {
            if (0 > p2) {
                var v71 uint64 = 1
                _ = v71
            } else {
                if false {
                    for fi72 := 0; fi72 < 11; fi72++ {
                        var v73 float32 = (((-2.25 * ((0.1 * 1e10) - 0.0)) + -2.25) - -2.25)
                        var v74 bool = true
                        _ = fi72
                        _ = v73
                        _ = v74
                    }
                    p0 += -2147483648
                    if false {
                        if (976 < uint32(p2)) {
                            v75 := []int{127, (p0 - p0)}
                            _ = v75
                        } else {
                            for fi76 := 0; fi76 < 13; fi76++ {
                                v77 := []int{}
                                switch (p0 % ((fi76) | 1)) {
                                case -9223372036854775808:
                                    p2 = (p2 / (((p2 + (^(p2)))) | 1))
                                case 32768:
                                    if len(v77) >= 2 {
                                        v77 = v77[1:len(v77)]
                                    }
                                }
                                _ = fi76
                                _ = v77
                            }
                            if (!false) {
                                v78 := []int{(p0 % ((p0) | 1)), (p0 >> (((p2 + (p2 + (p2 % (((^(p2))) | 1))))) & 63))}
                                if false {
                                    p1 = ("foo bar" + ((p1 + p1) + ("\n\t" + ((((("hello" + p1) + p1) + (p1 + p1)) + p1) + p1))))
                                } else {
                                    _ = len(v78)
                                    var v79 float64 = 3.14159
                                    if true {
                                        if true {
                                            for fi80 := 0; fi80 < 11; fi80++ {
                                                if (!(!false)) {
                                                    fmt.Println(v79)
                                                    v81 := GenS9{F0: 128, F1: 0}
                                                    var v82 bool = true
                                                    _ = v81
                                                    _ = v82
                                                }
                                                if true {
                                                    switch ("\n\t" + "foo bar") {
                                                    case "a":
                                                        if len(v78) >= 2 {
                                                            v78 = v78[1:len(v78)]
                                                        }
                                                    case "foo bar":
                                                        _ = len(v78)
                                                    }
                                                } else {
                                                    for fi83 := 0; fi83 < 15; fi83++ {
                                                        switch 46 {
                                                        case 0:
                                                            v79 = (v79 / v79)
                                                        default:
                                                            p0 = fi80
                                                        }
                                                        v78 = append(v78, -9223372036854775808)
                                                        p1 = p1
                                                        _ = fi83
                                                    }
                                                    v84 := GenS10{F0: 0, F1: p0, F2: 2147483647}
                                                    _ = v84
                                                }
                                                _ = fi80
                                            }
                                        }
                                    } else {
                                        v85 := GenS11{F0: p2}
                                        v86 := GenS7{F0: p2, F1: 255, F2: false}
                                        _ = v85
                                        _ = v86
                                    }
                                    _ = v79
                                }
                                _ = v78
                            } else {
                                if ((p2 >> ((p2) & 63)) > (p2 >> (((p2 / ((p2) | 1))) & 63))) {
                                    if (false || true) {
                                        v87 := []int{2147483647, 32767}
                                        fmt.Println(p0, p2)
                                        _ = v87
                                    } else {
                                        if false {
                                            if true {
                                                v88 := GenS8{F0: 672, F1: 886, F2: 204}
                                                p0 -= 4294967295
                                                p0 += p0
                                                _ = v88
                                            } else {
                                                v89 := GenS12{F0: 1, F1: 0}
                                                v90 := []int{(-(p0)), -891, (-(p0)), p0}
                                                for fi91 := 0; fi91 < 18; fi91++ {
                                                    for fi92 := 0; fi92 < 7; fi92++ {
                                                        var v93 int32 = 0
                                                        _ = fi92
                                                        _ = v93
                                                    }
                                                    var v94 float64 = 0.1
                                                    _ = fi91
                                                    _ = v94
                                                }
                                                _ = v89
                                                _ = v90
                                            }
                                        }
                                    }
                                }
                                p2 += p2
                            }
                        }
                    }
                }
            }
            fmt.Println(p1)
        }
        fmt.Println(p1, p1, p1)
        v95 := GenS13{F0: -9223372036854775808, F1: 65535, F2: 571}
        _ = v95
    }
    for fi96 := 0; fi96 < 20; fi96++ {
        v97 := []int{fi96, -9223372036854775808, fi96}
        if true {
            var v98 uint8 = 9
            _ = len(v97)
            _ = v98
        } else {
            if len(v97) > 0 {
                v97[0] = p0
            }
        }
        _ = fi96
        _ = v97
    }
    for fi99 := 0; fi99 < 16; fi99++ {
        p2 += 521
        p1 = "hello"
        for fi100 := 0; fi100 < 10; fi100++ {
            p0 += (fi99 >> ((p2) & 63))
            switch ("" + ((p1 + p1) + p1)) {
            case "\n\t":
                v101 := []int{(-(fi99)), p0, fi99}
                _ = v101
            case "a":
                fmt.Println(fi99, p2, p2)
            case "foo bar":
                fmt.Println(fi99, p0, fi99)
            default:
                switch p1 {
                case "a":
                    v102 := GenS8{F0: 486, F1: 255, F2: 2}
                    _ = v102
                case "hello":
                    var v103 uint8 = 112
                    _ = v103
                }
            }
            p2 += (p2 - 888)
            _ = fi100
        }
        _ = fi99
    }
    if true {
        if ((2147483647 != 128) && true) {
            var v104 int = 512
            for fi105 := 0; fi105 < 14; fi105++ {
                v106 := GenS14{F0: ("hello" + p1), F1: 0, F2: 100}
                if true {
                    for fi107 := 0; fi107 < 4; fi107++ {
                        switch (fi107 % (((fi105 | 535)) | 1)) {
                        case -9223372036854775808:
                            fmt.Println(p0, p0, p0)
                        case -1:
                            fmt.Println(fi105)
                        default:
                            if (-8 != -43) {
                                fmt.Println(p2, p1, p0)
                            }
                        }
                        if (false || false) {
                            fmt.Println(v104, p1, fi107)
                            switch "a" {
                            case "hello":
                                v108 := GenS15{F0: 1}
                                _ = v108
                            case "\n\t":
                                p2 += (^(p2))
                            }
                            p0 = (fi107 % ((v104) | 1))
                        }
                        p1 = "a"
                        _ = fi107
                    }
                    var v109 int32 = 0
                    if true {
                        if false {
                            fmt.Println(p1, fi105)
                            switch (p1 + p1) {
                            case "":
                                v109 -= v109
                            case "foo bar":
                                if (!true) {
                                    for fi110 := 0; fi110 < 8; fi110++ {
                                        fmt.Println(v104)
                                        if false {
                                            var v111 uint16 = 2
                                            p0 -= v104
                                            _ = v111
                                        }
                                        var v112 uint64 = 222
                                        _ = fi110
                                        _ = v112
                                    }
                                    fmt.Println(p2, p2)
                                } else {
                                    v113 := GenS16{F0: 65535, F1: 313, F2: 803}
                                    _ = v113
                                }
                            }
                            switch (^(v104)) {
                            case 42:
                                if ((v109 << (((p2 * (^(p2)))) & 63)) != (v109 / (((v109 << ((255) & 63))) | 1))) {
                                    for fi114 := 0; fi114 < 6; fi114++ {
                                        v115 := []int{p0, 9223372036854775807, (^(fi105)), (fi114 | fi114)}
                                        _ = fi114
                                        _ = v115
                                    }
                                    v104 -= (fi105 << ((1) & 63))
                                    v116 := []int{2147483647, p0}
                                    _ = v116
                                }
                            case -1:
                                if ((0.0 + 0.1) > 1e10) {
                                    var v117 string = (p1 + p1)
                                    _ = v117
                                } else {
                                    v118 := []int{}
                                    var v119 string = ("hello" + "a")
                                    var v120 bool = true
                                    _ = v118
                                    _ = v119
                                    _ = v120
                                }
                            case 174:
                                v121 := []int{p0, fi105}
                                _ = v121
                            default:
                                v109 = v109
                            }
                        }
                        for fi122 := 0; fi122 < 3; fi122++ {
                            v123 := []int{v104, (p0 & p0), (p0 % ((fi105) | 1))}
                            _ = fi122
                            _ = v123
                        }
                        p1 = ""
                    }
                    _ = v109
                }
                _ = fi105
                _ = v106
            }
            _ = v104
        }
        var v124 int32 = -1
        for fi125 := 0; fi125 < 7; fi125++ {
            v126 := GenS17{F0: 255, F1: 0, F2: (-(p0))}
            if true {
                var v127 uint32 = 2
                for fi128 := 0; fi128 < 4; fi128++ {
                    if (2 < 42) {
                        if true {
                            switch (fi125 | fi128) {
                            case 257:
                                v129 := []int{(p0 + 127), fi125, 127, int(p2)}
                                v130 := GenS18{F0: 126, F1: -769, F2: -968}
                                var v131 uint = p2
                                _ = v129
                                _ = v130
                                _ = v131
                            case 32768:
                                v132 := GenS19{F0: v127, F1: -9223372036854775808}
                                _ = v132
                            case 128:
                                var v133 uint32 = (v127 >> ((4294967295) & 63))
                                _ = v133
                            }
                        } else {
                            var v134 bool = (false || (((v124 != -421) || false) && false))
                            v135 := []int{(fi125 >> (((^(p2))) & 63)), (fi125 * -1), fi125, fi125}
                            var v136 uint64 = 18446744073709551615
                            _ = v134
                            _ = v135
                            _ = v136
                        }
                    } else {
                        p2 = p2
                        v137 := GenS15{F0: v124}
                        v138 := GenS20{F0: 127, F1: -33}
                        _ = v137
                        _ = v138
                    }
                    _ = fi128
                }
                _ = v127
            }
            _ = fi125
            _ = v126
        }
        _ = v124
    }
    for fi139 := 0; fi139 < 17; fi139++ {
        v140 := []int{fi139}
        _ = fi139
        _ = v140
    }
    _ = p0
    _ = p1
    _ = p2
    return true
}
func genFn1() uint8 {
    var v141 uint16 = 1
    if false {
        v142 := []int{-2147483648}
        _ = v142
    } else {
        v143 := GenS4{F0: ("hello" + "\n\t")}
        if (((-2.25 < 0.1) && false) || (!((!true) && (!(!false))))) {
            var v144 uint16 = 65535
            _ = v144
        } else {
            var v145 uint8 = 59
            v146 := GenS21{F0: (v145 % (((v145 & 42)) | 1)), F1: 0, F2: (v145 % ((v145) | 1))}
            _ = v145
            _ = v146
        }
        for fi147 := 0; fi147 < 14; fi147++ {
            v148 := []int{32767, 2147483647}
            if (!true) {
                for fi149 := 0; fi149 < 7; fi149++ {
                    var v150 int8 = -1
                    if len(v148) >= 2 {
                        v148 = v148[1:len(v148)]
                    }
                    _ = fi149
                    _ = v150
                }
            }
            _ = fi147
            _ = v148
        }
        _ = v143
    }
    _ = v141
    return 127
}

func main() {
    var v1 uint32 = 415
    v2 := make(map[uint32]uint32)
    var v3 float32 = (0.0 / float32(((1.5 - (3.14159 + (1.5 / 1.5))) - (float64(-2.25) / 1.0))))
    var v151 uint8 = 42
    switch ("hello" + "") {
    case "foo bar":
        if false {
            for fi152 := 0; fi152 < 13; fi152++ {
                fmt.Println(v151)
                fmt.Println(v151)
                _ = fi152
            }
            v153 := []int{}
            _ = v153
        } else {
            if ((^(v1)) > v1) {
                var v154 int = -9223372036854775808
                _ = v154
            } else {
                var v155 int = -32
                _ = v155
            }
        }
    case "hello":
        var v156 int16 = -158
        _ = v156
    case "":
        var v157 string = "\n\t"
        _ = v157
    default:
        v158 := []int{848, -335}
        _ = v158
    }
    v159 := GenS22{F0: -9223372036854775808, F1: 42}
    var v160 uint8 = 42
    for fi161 := 0; fi161 < 15; fi161++ {
        var v162 bool = ((false && (!("a" > "\n\t"))) && true)
        v1 += (^(v1))
        v160 -= genFn1()
        _ = fi161
        _ = v162
    }
    v163 := GenS17{F0: 65535, F1: 0, F2: 42}
    v164 := GenS15{F0: 0}
    var v165 int32 = 107
    v166 := GenS20{F0: v160, F1: 127}
    v1 -= (v1 & 0)
    if (!false) {
        v167 := []int{2147483647, 0, 235, 1}
        _ = v167
    } else {
        v165 -= (v165 >> ((4294967295) & 63))
    }
    v168 := GenS6{F0: 0, F1: 9223372036854775807, F2: 18446744073709551615}
    v1 = (v1 >> ((1) & 63))
    _ = len(v2)
    switch "\n\t" {
    case "\n\t":
        fmt.Println(v3, v160, v160)
    case "a":
        v165 += int32(v1)
    default:
        v169 := []int{int(v160), 42, 4294967295}
        _ = v169
    }
    v170 := v2[(v1 ^ (v1 / (((v1 / ((v1) | 1))) | 1)))]
    var v171 float64 = (1.5 / 1.0)
    fmt.Println("v1:", v1)
    fmt.Println("v151:", v151)
    fmt.Println("v160:", v160)
    fmt.Println("v165:", v165)
    fmt.Println("v170:", v170)
    fmt.Println("v171:", v171)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    _ = v1
    _ = v151
    _ = v159
    _ = v160
    _ = v163
    _ = v164
    _ = v165
    _ = v166
    _ = v168
    _ = v170
    _ = v171
    _ = v2
    _ = v3
}
