package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 uint; F2 uint32 }
type GenS2 struct { F0 string }
type GenS3 struct { F0 uint8 }
type GenS4 struct { F0 int16; F1 uint16; F2 uint16 }
type GenS5 struct { F0 int32; F1 uint16 }
type GenS6 struct { F0 bool }
type GenS7 struct { F0 bool }
type GenS8 struct { F0 int8; F1 uint8 }
type GenS9 struct { F0 int }
type GenS10 struct { F0 uint8 }
type GenS11 struct { F0 uint64; F1 int64 }

func main() {
    v1 := make(map[int]int)
    var v2 uint64 = 18446744073709551615
    var v3 string = "\n\t"
    var v4 bool = true
    var v5 uint = 964
    v6 := make(map[int]uint32)
    for fi7 := 0; fi7 < 8; fi7++ {
        if false {
            for fi8 := 0; fi8 < 2; fi8++ {
                if v4 {
                    v9 := []int{9223372036854775807, 599}
                    v3 = ""
                    var v10 uint64 = (v2 + 42)
                    _ = v10
                    _ = v9
                } else {
                    switch "hello" {
                    case "":
                        v4 = false
                    case "hello":
                        v11 := []int{(fi7 / ((32767) | 1)), 20}
                        _ = v11
                    }
                    v12 := []int{fi7, (fi7 - (fi7 << ((v5) & 63)))}
                    v2 -= v2
                    _ = v12
                }
                var v13 uint8 = 128
                _ = fi8
                _ = v13
            }
            var v14 int8 = 40
            _ = v14
        }
        fmt.Println(v5, v2)
        _ = fi7
    }
    for fi15 := 0; fi15 < 2; fi15++ {
        var v16 int8 = -128
        _ = fi15
        _ = v16
    }
    fmt.Println(v2)
    if (!false) {
        for fi17 := 0; fi17 < 5; fi17++ {
            switch int(v5) {
            case -1:
                switch ("foo bar" + ("a" + v3)) {
                case "\n\t":
                    v4 = ((!(672 == 4294967295)) && (v4 || true))
                case "a":
                    switch (v3 + (v3 + v3)) {
                    case "foo bar":
                        v18 := GenS1{F0: (!(-894 != 270)), F1: (v5 >> ((v5) & 63)), F2: 1}
                        _ = v18
                    case "\n\t":
                        for fi19 := 0; fi19 < 18; fi19++ {
                            if false {
                                if (!false) {
                                    v20 := []int{(fi17 - (fi17 & fi19)), fi17}
                                    _ = v20
                                } else {
                                    v21 := v1[fi17]
                                    v22 := []int{1, v21, fi17}
                                    var v23 bool = v4
                                    _ = v21
                                    _ = v22
                                    _ = v23
                                }
                            }
                            v24 := GenS1{F0: false, F1: 18446744073709551615, F2: 197}
                            var v25 uint64 = (v2 | (v2 % ((v2) | 1)))
                            _ = fi19
                            _ = v24
                            _ = v25
                        }
                    }
                default:
                    v4 = (v5 < (^(v5)))
                }
            case 2147483647:
                v26 := v6[fi17]
                _ = v26
            }
            for fi27 := 0; fi27 < 20; fi27++ {
                v4 = (1 <= 32767)
                v28 := []int{4294967295}
                _ = fi27
                _ = v28
            }
            switch (fi17 / ((-108) | 1)) {
            case 1:
                for fi29 := 0; fi29 < 15; fi29++ {
                    var v30 float32 = ((1e10 / 0.1) + 1e10)
                    _ = fi29
                    _ = v30
                }
            case -385:
                v1[(fi17 + (fi17 % ((fi17) | 1)))] = fi17
            }
            _ = fi17
        }
        if (false && v4) {
            v31 := GenS1{F0: v4, F1: 255, F2: 210}
            _ = v31
        } else {
            switch v3 {
            case "hello":
                for fi32 := 0; fi32 < 4; fi32++ {
                    v33 := []int{(fi32 * (-(fi32))), fi32, fi32, 127}
                    var v34 uint64 = 18446744073709551615
                    _ = fi32
                    _ = v33
                    _ = v34
                }
            }
        }
        if v4 {
            v35 := []int{127, 9223372036854775807}
            if v4 {
                var v36 int32 = 420
                if (0.1 <= (((0.0 - 0.0) * 1.5) * 0.0)) {
                    for fi37 := 0; fi37 < 20; fi37++ {
                        v38 := GenS1{F0: false, F1: v5, F2: 0}
                        if v4 {
                            var v39 uint64 = (v2 | v2)
                            v4 = (!(v4 && true))
                            _ = v39
                        } else {
                            v40 := GenS1{F0: (-351 < 1), F1: v5, F2: 4294967295}
                            _ = v40
                        }
                        switch (("hello" + ("a" + "\n\t")) + "\n\t") {
                        case "a":
                            if true {
                                if v4 {
                                    v5 += (^(v5))
                                    var v41 float64 = ((0.1 * 0.0) * 0.0)
                                    _ = v41
                                } else {
                                    var v42 int32 = 42
                                    v36 = (^(v42))
                                    for fi43 := 0; fi43 < 6; fi43++ {
                                        var v44 bool = (((false && true) || v4) || (255 == 1))
                                        _ = fi43
                                        _ = v44
                                    }
                                    _ = v42
                                }
                            }
                        case "foo bar":
                            if v4 {
                                v3 = (("a" + ((v3 + ("foo bar" + v3)) + v3)) + v3)
                                v3 = ("a" + "a")
                                v45 := GenS1{F0: (v4 || (v3 >= "")), F1: v5, F2: 1}
                                _ = v45
                            }
                        case "":
                            var v46 uint8 = 233
                            _ = v46
                        }
                        _ = fi37
                        _ = v38
                    }
                }
                var v47 uint = (v5 * (v5 | v5))
                _ = v36
                _ = v47
            }
            _ = v35
        } else {
            switch (v3 + ("\n\t" + (v3 + v3))) {
            case "":
                for fi48 := 0; fi48 < 16; fi48++ {
                    if v4 {
                        for fi49 := 0; fi49 < 19; fi49++ {
                            v2 = v2
                            _ = fi49
                        }
                        v4 = true
                        var v50 string = v3
                        _ = v50
                    }
                    for fi51 := 0; fi51 < 13; fi51++ {
                        var v52 uint32 = 640
                        _ = fi51
                        _ = v52
                    }
                    v53 := GenS1{F0: (-1 <= 1), F1: v5, F2: 4294967295}
                    _ = fi48
                    _ = v53
                }
            case "\n\t":
                if true {
                    for fi54 := 0; fi54 < 6; fi54++ {
                        switch fi54 {
                        case 32768:
                            if (!v4) {
                                v3 = "hello"
                                switch v3 {
                                case "\n\t":
                                    fmt.Println(fi54)
                                case "foo bar":
                                    v3 = v3
                                }
                            } else {
                                v55 := []int{fi54, fi54}
                                _ = v55
                            }
                        case 13:
                            for fi56 := 0; fi56 < 10; fi56++ {
                                var v57 uint16 = 255
                                var v58 uint64 = v2
                                _ = fi56
                                _ = v57
                                _ = v58
                            }
                        case -9223372036854775808:
                            if (v4 && (!(1 < -1))) {
                                v5 = v5
                                v2 = v2
                            } else {
                                var v59 uint32 = 1
                                v60 := GenS2{F0: v3}
                                _ = v59
                                _ = v60
                            }
                        }
                        switch ((((v3 + "foo bar") + v3) + (v3 + (v3 + v3))) + ("hello" + v3)) {
                        case "a":
                            var v61 int = -2147483648
                            _ = v61
                        case "\n\t":
                            v3 = v3
                        case "hello":
                            v4 = (!v4)
                        }
                        _ = fi54
                    }
                    var v62 int = 9223372036854775807
                    v63 := GenS2{F0: ""}
                    _ = v62
                    _ = v63
                }
            default:
                v64 := GenS2{F0: "foo bar"}
                _ = v64
            }
        }
    }
    if (((1.5 / 0.1) > -2.25) || ((4294967295 != 42) || (-9223372036854775808 == 16))) {
        v65 := GenS2{F0: v3}
        if ((v2 >> ((v5) & 63)) <= (v2 / ((2) | 1))) {
            v66 := GenS1{F0: (v4 || (-32768 < -1)), F1: 18446744073709551615, F2: 453}
            _ = v66
        }
        _ = v65
    } else {
        var v67 int64 = -1
        _ = v67
    }
    for fi68 := 0; fi68 < 12; fi68++ {
        var v69 bool = false
        if (v5 != v5) {
            switch "\n\t" {
            case "foo bar":
                for fi70 := 0; fi70 < 17; fi70++ {
                    var v71 uint = v5
                    _ = fi70
                    _ = v71
                }
            case "\n\t":
                var v72 uint8 = 42
                _ = v72
            case "hello":
                fmt.Println(v2, v69, v2)
            default:
                v3 = (v3 + "\n\t")
            }
            v4 = v4
        } else {
            fmt.Println(fi68)
            switch "foo bar" {
            case "hello":
                for fi73 := 0; fi73 < 20; fi73++ {
                    v2 += (v2 % ((v2) | 1))
                    v74 := []int{32768, (^(fi68))}
                    _ = fi73
                    _ = v74
                }
            case "":
                if v4 {
                    var v75 float64 = (0.1 / 1e10)
                    v76 := []int{-9223372036854775808, (fi68 * -703), fi68}
                    if false {
                        v77 := GenS2{F0: v3}
                        v78 := GenS1{F0: (0 == 42), F1: v5, F2: 118}
                        var v79 uint8 = 0
                        _ = v77
                        _ = v78
                        _ = v79
                    } else {
                        var v80 int8 = 0
                        if (v5 == 503) {
                            var v81 float32 = 0.1
                            if len(v76) >= 2 {
                                v76 = v76[1:len(v76)]
                            }
                            for fi82 := 0; fi82 < 4; fi82++ {
                                v80 += v80
                                v5 = (v5 << ((v5) & 63))
                                fmt.Println(v3)
                                _ = fi82
                            }
                            _ = v81
                        } else {
                            var v83 int = fi68
                            v4 = true
                            var v84 uint64 = v2
                            _ = v83
                            _ = v84
                        }
                        _ = v80
                    }
                    _ = v75
                    _ = v76
                }
            default:
                for fi85 := 0; fi85 < 20; fi85++ {
                    v86 := GenS2{F0: (v3 + (v3 + v3))}
                    if v69 {
                        v3 = v3
                        v5 = v5
                        if (128 < 1) {
                            delete(v1, fi68)
                            v87 := v6[(fi85 & (fi68 | (fi85 + 127)))]
                            if (!false) {
                                var v88 float64 = -2.25
                                _ = v88
                            } else {
                                fmt.Println(v87, v2)
                                switch v3 {
                                case "a":
                                    if ((^(v2)) > v2) {
                                        var v89 int = fi68
                                        if v4 {
                                            var v90 float64 = 3.14159
                                            if (false && (!(v4 || (uint8(v87) >= 42)))) {
                                                _ = len(v1)
                                                if (v5 != (v5 % (((v5 ^ 302)) | 1))) {
                                                    v91 := GenS2{F0: (v3 + (v3 + "a"))}
                                                    _ = v91
                                                }
                                                v92 := GenS2{F0: v3}
                                                _ = v92
                                            }
                                            v93 := GenS3{F0: 128}
                                            _ = v90
                                            _ = v93
                                        }
                                        _ = v89
                                    }
                                }
                            }
                            _ = v87
                        } else {
                            v94 := []int{-156, (fi68 + fi68), fi68}
                            var v95 uint64 = (v2 * (^(v2)))
                            if v4 {
                                v95 = 4294967295
                                if v4 {
                                    switch v3 {
                                    case "hello":
                                        var v96 uint = v5
                                        _ = v96
                                    case "\n\t":
                                        var v97 float64 = -2.25
                                        _ = v97
                                    default:
                                        switch -266 {
                                        case 4294967295:
                                            if true {
                                                if len(v94) >= 2 {
                                                    v94 = v94[1:len(v94)]
                                                }
                                                fmt.Println(fi68, v69)
                                            } else {
                                                if len(v94) > 0 {
                                                    v94[0] = fi85
                                                }
                                                v94 = append(v94, fi68)
                                                v5 = v5
                                            }
                                            var v98 uint32 = 176
                                            if len(v94) >= 2 {
                                                v94 = v94[1:len(v94)]
                                            }
                                            _ = v98
                                        case 2147483647:
                                            v95 -= (v2 | v2)
                                        default:
                                            fmt.Println(v4, fi85, fi85)
                                        }
                                    }
                                    var v99 bool = (1e10 == (float32(float64((-2.25 - 0.0))) - 1.5))
                                    _ = v99
                                }
                            }
                            _ = v94
                            _ = v95
                        }
                        if v4 {
                            v100 := []int{(fi68 / ((-2147483648) | 1)), (fi85 >> (((v5 | (v5 * (v5 / ((v5) | 1))))) & 63))}
                            var v101 bool = v69
                            _ = v100
                            _ = v101
                        } else {
                            v5 = v5
                            v3 = "\n\t"
                            var v102 uint8 = uint8(v2)
                            v103 := GenS3{F0: 42}
                            _ = v102
                            _ = v103
                        }
                    } else {
                        fmt.Println(fi68, v2, v4)
                        var v104 float64 = 0.1
                        if true {
                            fmt.Println(v2, v4)
                            v105 := v6[fi68]
                            var v106 int64 = 917
                            _ = v105
                            _ = v106
                        } else {
                            if true {
                                var v107 uint64 = v2
                                switch "a" {
                                case "hello":
                                    var v108 int8 = -128
                                    _ = v108
                                case "foo bar":
                                    if (v104 != v104) {
                                        v109 := GenS4{F0: 1000, F1: uint16(v2), F2: 255}
                                        v3 = "\n\t"
                                        if (true || ((30 >= 2) || (!((2 != uint16(v5)) && (42 != 127))))) {
                                            fmt.Println(v3, v4)
                                        }
                                        _ = v109
                                    }
                                }
                                _ = len(v1)
                                v110 := GenS5{F0: 0, F1: 255}
                                var v111 int = fi68
                                if v69 {
                                    var v112 int8 = 41
                                    if v4 {
                                        var v113 float32 = 1e10
                                        if (((^(fi85)) == int(v107)) && ((((v3 + "foo bar") + (("\n\t" + v3) + v3)) + ((v3 + "a") + v3)) <= ("foo bar" + ((v3 + (v3 + "")) + v3)))) {
                                            v104 += v104
                                            switch ("" + v3) {
                                            case "foo bar":
                                                v69 = false
                                            case "a":
                                                var v114 float32 = 1e10
                                                _ = v114
                                            }
                                            v113 -= v113
                                        } else {
                                            v115 := GenS1{F0: (550 <= (^(v111))), F1: (v5 & (v5 % ((v5) | 1))), F2: 0}
                                            var v116 uint = 0
                                            v113 -= (v113 - (float32(v104) / v113))
                                            _ = v115
                                            _ = v116
                                        }
                                        if (v69 && (625 < int64(v112))) {
                                            v117 := GenS6{F0: false}
                                            if ((v112 ^ -69) > (^(v112))) {
                                                var v118 bool = v4
                                                fmt.Println(fi85, v104, v107)
                                                _ = v118
                                            } else {
                                                var v119 float64 = v104
                                                if (!true) {
                                                    _ = len(v1)
                                                }
                                                _ = v119
                                            }
                                            _ = v117
                                        }
                                        v120 := GenS7{F0: true}
                                        _ = v113
                                        _ = v120
                                    } else {
                                        if v69 {
                                            switch (v111 / (((fi85 - v111)) | 1)) {
                                            case 1:
                                                v111 = (^(fi85))
                                            case 42:
                                                if (1 != 32767) {
                                                    var v121 bool = (42 != 18446744073709551615)
                                                    var v122 int16 = 1000
                                                    var v123 uint32 = 4294967295
                                                    v124 := GenS6{F0: (-1 >= (v122 - v122))}
                                                    _ = v121
                                                    _ = v122
                                                    _ = v123
                                                    _ = v124
                                                } else {
                                                    v111 -= fi85
                                                    if true {
                                                        v125 := GenS8{F0: 1, F1: 1}
                                                        _ = v125
                                                    }
                                                }
                                            default:
                                                var v126 float32 = 0.0
                                                if v4 {
                                                    fmt.Println(v5, v5)
                                                } else {
                                                    var v127 int32 = int32(v5)
                                                    if v4 {
                                                        v4 = false
                                                        var v128 float64 = (v104 / (((v104 / (v104 / 1e10)) + 3.14159) + v104))
                                                        if true {
                                                            v126 -= v126
                                                        } else {
                                                            v129 := GenS9{F0: (^(fi85))}
                                                            v107 += (v107 & (v107 >> ((v5) & 63)))
                                                            v107 = 992
                                                            _ = v129
                                                        }
                                                        _ = v128
                                                    }
                                                    _ = v127
                                                }
                                                _ = v126
                                            }
                                            switch ((((("hello" + "\n\t") + v3) + v3) + v3) + (v3 + (v3 + v3))) {
                                            case "a":
                                                var v130 float32 = float32(v104)
                                                _ = v130
                                            case "":
                                                v6[(^(v111))] = 981
                                            }
                                        } else {
                                            v131 := GenS4{F0: 0, F1: 2, F2: uint16(v5)}
                                            var v132 uint = v5
                                            _ = v131
                                            _ = v132
                                        }
                                    }
                                    v69 = (v4 || (-357 != (v111 + (fi85 ^ (fi85 >> ((uint(v2)) & 63))))))
                                    fmt.Println(v107)
                                    _ = v112
                                } else {
                                    var v133 float64 = v104
                                    var v134 int = (v111 / ((fi68) | 1))
                                    fmt.Println(v107)
                                    _ = v133
                                    _ = v134
                                }
                                v2 += v107
                                v107 = (v107 * v107)
                                v135 := GenS2{F0: "hello"}
                                if true {
                                    v104 = v104
                                    if (!v4) {
                                        fmt.Println(v111, v111, v4)
                                    }
                                    v6[v111] = 746
                                }
                                v104 += ((v104 + 1e10) + (float64(1e10) - v104))
                                v69 = (((0.1 / 1.0) * 3.14159) != -2.25)
                                var v136 int = fi85
                                _ = v107
                                _ = v110
                                _ = v111
                                _ = v135
                                _ = v136
                            } else {
                                fmt.Println(v4)
                                var v137 uint64 = 42
                                _ = v137
                            }
                            switch v3 {
                            case "\n\t":
                                v138 := GenS10{F0: 128}
                                _ = v138
                            default:
                                var v139 int8 = 0
                                _ = v139
                            }
                        }
                        _ = v104
                    }
                    switch (-(fi85)) {
                    case 42:
                        var v140 uint8 = 127
                        _ = v140
                    }
                    _ = fi85
                    _ = v86
                }
            }
        }
        _ = fi68
        _ = v69
    }
    var v141 int = 32767
    var v142 uint8 = 10
    var v143 uint = 255
    var v144 int64 = -9223372036854775808
    _ = len(v1)
    for fi145 := 0; fi145 < 15; fi145++ {
        for fi146 := 0; fi146 < 14; fi146++ {
            delete(v6, 687)
            v142 = 128
            v147 := GenS7{F0: true}
            v2 += v2
            _ = fi146
            _ = v147
        }
        _ = fi145
    }
    var v148 int32 = -2147483648
    var v149 float64 = ((1.5 * -2.25) + (((-2.25 - 0.0) - ((float64(0.0) + 0.0) - ((3.14159 * (((0.1 / 1e10) * 1.5) + 1.5)) * (((-2.25 / 1.0) / 0.1) / -2.25)))) * (((3.14159 * -2.25) * ((1e10 - (1e10 - 1e10)) - (1.5 / float64((0.1 + 0.1))))) * 1e10)))
    v2 = v2
    for fi150 := 0; fi150 < 5; fi150++ {
        var v151 int32 = (v148 << ((18446744073709551615) & 63))
        if (-102 >= 1) {
            _ = len(v1)
        } else {
            for fi152 := 0; fi152 < 19; fi152++ {
                var v153 float64 = v149
                _ = fi152
                _ = v153
            }
            if v4 {
                v148 += 975
            }
            v149 -= (1.5 / 1e10)
        }
        switch (fi150 << ((v5) & 63)) {
        case -2147483648:
            for fi154 := 0; fi154 < 17; fi154++ {
                if v4 {
                    v155 := GenS11{F0: v2, F1: 905}
                    _ = v155
                }
                v141 = (fi150 - fi154)
                _ = fi154
            }
        case 1:
            if true {
                v2 -= v2
            }
        }
        _ = fi150
        _ = v151
    }
    switch -372 {
    case 9223372036854775807:
        fmt.Println(v142)
    default:
        v6[(v141 / (((v141 & (-(v141)))) | 1))] = 65535
    }
    v156 := []int{(v141 | (v141 - v141))}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v141:", v141)
    fmt.Println("v142:", v142)
    fmt.Println("v143:", v143)
    fmt.Println("v144:", v144)
    fmt.Println("v148:", v148)
    fmt.Println("v149:", v149)
    fmt.Println("slicelen v156:", len(v156), v156)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v141
    _ = v142
    _ = v143
    _ = v144
    _ = v148
    _ = v149
    _ = v156
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
