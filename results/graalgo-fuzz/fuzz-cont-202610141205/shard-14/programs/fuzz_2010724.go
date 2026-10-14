package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 uint16 }
type GenS2 struct { F0 int16; F1 int; F2 uint8 }
type GenS3 struct { F0 int8 }
type GenS4 struct { F0 int8; F1 bool }
type GenS5 struct { F0 uint64 }
type GenS6 struct { F0 int8; F1 uint8; F2 int32 }
type GenS7 struct { F0 uint16; F1 bool }
type GenS8 struct { F0 int }
type GenS9 struct { F0 int16; F1 uint; F2 int64 }
type GenS10 struct { F0 uint }
type GenS11 struct { F0 int16; F1 string; F2 string }
type GenS12 struct { F0 int32; F1 uint8; F2 uint32 }
type GenS13 struct { F0 uint8; F1 int32; F2 int16 }
type GenS14 struct { F0 int; F1 uint32; F2 string }
type GenS15 struct { F0 bool }
func genFn0(p0 int64, p1 int16, p2 uint32) string {
    for fi7 := 0; fi7 < 10; fi7++ {
        var v8 uint16 = 1
        var v9 uint8 = 42
        v8 = 0
        _ = fi7
        _ = v8
        _ = v9
    }
    fmt.Println(p1)
    if false {
        p0 = -9223372036854775808
        var v10 float64 = 1.5
        _ = v10
    }
    for fi11 := 0; fi11 < 4; fi11++ {
        for fi12 := 0; fi12 < 20; fi12++ {
            v13 := []int{fi11}
            v13 = append(v13, fi12)
            if true {
                var v14 uint16 = 2
                if len(v13) >= 2 {
                    v13 = v13[1:len(v13)]
                }
                _ = v14
            }
            p0 += 16
            _ = fi12
            _ = v13
        }
        _ = fi11
    }
    _ = p0
    _ = p1
    _ = p2
    return "\n\t"
}

func main() {
    var v1 uint32 = 1
    var v2 uint64 = 4294967295
    var v3 int64 = 1
    var v4 int32 = -2147483648
    var v5 int64 = (v3 << ((833) & 63))
    v6 := make(map[uint32]string)
    if true {
        for fi15 := 0; fi15 < 7; fi15++ {
            v4 = (-(v4))
            if false {
                var v16 string = "foo bar"
                switch fi15 {
                case -680:
                    v17 := GenS1{F0: "", F1: 42}
                    _ = v17
                case 725:
                    fmt.Println(v5)
                case 0:
                    for fi18 := 0; fi18 < 7; fi18++ {
                        v19 := GenS2{F0: -949, F1: 32768, F2: 2}
                        var v20 int64 = v3
                        switch ((v16 + v16) + ("hello" + v16)) {
                        case "a":
                            var v21 uint8 = 42
                            v2 = (^(v2))
                            if false {
                                var v22 int = fi18
                                _ = v22
                            }
                            _ = v21
                        default:
                            var v23 string = "foo bar"
                            _ = v23
                        }
                        var v24 uint64 = (v2 | 18446744073709551615)
                        switch v16 {
                        case "\n\t":
                            fmt.Println(v2, fi15, v4)
                        case "hello":
                            var v25 bool = false
                            _ = v25
                        case "":
                            v6[1] = (v16 + v16)
                        default:
                            var v26 uint16 = 32767
                            _ = v26
                        }
                        switch (fi15 / ((fi18) | 1)) {
                        case -2147483648:
                            v2 = (v2 << ((2) & 63))
                        case 32767:
                            var v27 string = v16
                            _ = v27
                        }
                        _ = fi18
                        _ = v19
                        _ = v20
                        _ = v24
                    }
                }
                switch -770 {
                case 0:
                    var v28 int32 = v4
                    _ = v28
                }
                _ = v16
            }
            _ = fi15
        }
    } else {
        v29 := GenS2{F0: -32768, F1: 42, F2: 2}
        _ = v29
    }
    var v30 float32 = ((float32((float64(1e10) * 1e10)) - 1.5) + (1e10 / 1.0))
    v31 := GenS2{F0: -464, F1: -2147483648, F2: 2}
    if (-122 >= -40) {
        for fi32 := 0; fi32 < 5; fi32++ {
            for fi33 := 0; fi33 < 2; fi33++ {
                var v34 uint16 = 65535
                _ = fi33
                _ = v34
            }
            v35 := GenS1{F0: (("" + "foo bar") + (("hello" + "a") + "\n\t")), F1: 771}
            switch fi32 {
            case 42:
                v3 += (v5 - (v5 << ((65535) & 63)))
            case 127:
                fmt.Println(fi32)
            }
            _ = fi32
            _ = v35
        }
        v36 := GenS3{F0: 1}
        _ = v36
    } else {
        for fi37 := 0; fi37 < 1; fi37++ {
            v5 += v3
            v38 := GenS4{F0: 1, F1: false}
            switch "a" {
            case "\n\t":
                var v39 uint = 827
                _ = v39
            case "foo bar":
                var v40 string = ("foo bar" + ("foo bar" + "\n\t"))
                _ = v40
            default:
                v4 += (v4 * v4)
            }
            _ = fi37
            _ = v38
        }
        var v41 float32 = v30
        _ = v41
    }
    var v42 uint64 = 2
    if true {
        var v43 int8 = -97
        if (32767 >= 342) {
            for fi44 := 0; fi44 < 10; fi44++ {
                var v45 string = (("foo bar" + "") + ("\n\t" + (("\n\t" + ("foo bar" + "")) + "hello")))
                _ = fi44
                _ = v45
            }
            var v46 bool = (!false)
            switch 128 {
            case 464:
                for fi47 := 0; fi47 < 14; fi47++ {
                    var v48 string = ("a" + "foo bar")
                    for fi49 := 0; fi49 < 8; fi49++ {
                        v6[v1] = v48
                        _ = fi49
                    }
                    _ = fi47
                    _ = v48
                }
            case 32767:
                var v50 int32 = v4
                _ = v50
            case -786:
                if ((v30 + float32((0.0 * 3.14159))) <= (((float32(0.1) + 0.0) + (v30 * v30)) + (v30 - (-2.25 / ((-2.25 + v30) * v30))))) {
                    switch "\n\t" {
                    case "":
                        for fi51 := 0; fi51 < 16; fi51++ {
                            var v52 float32 = (v30 * 0.0)
                            v42 -= v2
                            delete(v6, v1)
                            _ = fi51
                            _ = v52
                        }
                    case "foo bar":
                        switch "\n\t" {
                        case "a":
                            for fi53 := 0; fi53 < 16; fi53++ {
                                var v54 int8 = (v43 - 126)
                                v55 := GenS4{F0: (v54 % ((42) | 1)), F1: false}
                                _ = fi53
                                _ = v54
                                _ = v55
                            }
                        case "":
                            if true {
                                v3 = v3
                                v42 = (v42 / (((v42 / ((v2) | 1))) | 1))
                                var v56 int8 = (v43 + v43)
                                _ = v56
                            } else {
                                v1 += 2147483647
                            }
                        case "\n\t":
                            v57 := v6[(v1 >> ((18446744073709551615) & 63))]
                            _ = v57
                        default:
                            var v58 int = 9223372036854775807
                            _ = v58
                        }
                    case "a":
                        for fi59 := 0; fi59 < 17; fi59++ {
                            for fi60 := 0; fi60 < 5; fi60++ {
                                v61 := v6[v1]
                                var v62 float64 = (-2.25 - 0.1)
                                _ = fi60
                                _ = v61
                                _ = v62
                            }
                            _ = fi59
                        }
                    }
                    fmt.Println(v30)
                } else {
                    v63 := GenS2{F0: 42, F1: 1, F2: 82}
                    v1 -= (^(v1))
                    _ = v63
                }
            default:
                var v64 uint8 = 255
                _ = v64
            }
            _ = v46
        } else {
            var v65 int8 = (v43 >> ((70) & 63))
            var v66 uint = 1
            _ = v65
            _ = v66
        }
        var v67 int64 = (v5 ^ -9223372036854775808)
        _ = v43
        _ = v67
    }
    var v68 uint = 18446744073709551615
    _ = len(v6)
    v68 = v68
    v6[v1] = ("a" + "\n\t")
    for fi69 := 0; fi69 < 3; fi69++ {
        fmt.Println(v4)
        for fi70 := 0; fi70 < 1; fi70++ {
            v71 := GenS4{F0: -126, F1: ((true || false) && false)}
            var v72 uint64 = 18446744073709551615
            var v73 float64 = 0.1
            var v74 int8 = 0
            switch (^(fi69)) {
            case -9223372036854775808:
                if true {
                    var v75 bool = true
                    _ = v75
                }
            case 32768:
                v76 := GenS4{F0: (v74 % (((v74 + 0)) | 1)), F1: (!true)}
                _ = v76
            case -1:
                v77 := GenS3{F0: (-(v74))}
                _ = v77
            default:
                var v78 uint16 = uint16(v68)
                _ = v78
            }
            v79 := GenS3{F0: 42}
            _ = fi70
            _ = v71
            _ = v72
            _ = v73
            _ = v74
            _ = v79
        }
        if (false && (v68 > uint(v1))) {
            switch fi69 {
            case 9223372036854775807:
                fmt.Println(v42, fi69)
            case 0:
                v80 := GenS5{F0: (v42 + (v2 >> ((255) & 63)))}
                _ = v80
            case 2147483647:
                v81 := GenS6{F0: -128, F1: 85, F2: int32(v5)}
                _ = v81
            default:
                v82 := GenS5{F0: (v2 ^ (^(v2)))}
                _ = v82
            }
        } else {
            fmt.Println(v5, v2)
            if true {
                if true {
                    v30 = v30
                    v42 = v2
                    v83 := GenS7{F0: 883, F1: (!(true && (fi69 < (fi69 & 128))))}
                    _ = v83
                } else {
                    switch fi69 {
                    case 127:
                        v84 := GenS8{F0: (fi69 % (((fi69 - -2147483648)) | 1))}
                        _ = v84
                    }
                }
                switch "\n\t" {
                case "":
                    var v85 uint64 = v42
                    _ = v85
                case "a":
                    fmt.Println(v4, v30)
                case "\n\t":
                    var v86 float32 = (v30 - (float32(0.1) - 3.14159))
                    _ = v86
                default:
                    v87 := GenS9{F0: -1, F1: v68, F2: (v5 | (-(v5)))}
                    _ = v87
                }
                for fi88 := 0; fi88 < 20; fi88++ {
                    if (!true) {
                        v89 := GenS10{F0: v68}
                        _ = v89
                    }
                    v4 -= (v4 - v4)
                    if ((^(fi69)) < 128) {
                        if (9223372036854775807 <= (fi88 >> (((v68 | 0)) & 63))) {
                            v90 := []int{int(v42), 2147483647, (fi88 & fi69)}
                            switch -1 {
                            case -439:
                                fmt.Println(v42, v5)
                            default:
                                var v91 int = fi88
                                _ = v91
                            }
                            _ = v90
                        }
                        switch int(v4) {
                        case 128:
                            fmt.Println(v2)
                        case -1:
                            var v92 float32 = v30
                            _ = v92
                        case 669:
                            var v93 float64 = 3.14159
                            _ = v93
                        }
                    }
                    var v94 float64 = float64(float32(3.14159))
                    _ = fi88
                    _ = v94
                }
            } else {
                var v95 int64 = -1
                _ = v95
            }
            switch 1 {
            case 887:
                for fi96 := 0; fi96 < 3; fi96++ {
                    v68 += v68
                    var v97 uint16 = 255
                    if (!((-2.25 + 1e10) > -2.25)) {
                        if false {
                            v98 := GenS8{F0: 42}
                            v99 := []int{(fi96 << (((v68 ^ v68)) & 63)), fi96, (^(fi96)), 182}
                            _ = v98
                            _ = v99
                        } else {
                            v68 += (^(v68))
                            switch ("" + "hello") {
                            case "hello":
                                switch "" {
                                case "hello":
                                    v3 = v3
                                }
                            case "foo bar":
                                var v100 int8 = 127
                                _ = v100
                            case "":
                                v30 += ((((v30 * -2.25) * float32(((1e10 - 0.0) / 1.5))) / v30) / v30)
                            }
                            if (2 < (v68 >> (((v68 >> ((v68) & 63))) & 63))) {
                                fmt.Println(v1, fi69, fi96)
                                var v101 bool = true
                                v102 := GenS2{F0: 61, F1: (fi96 + (fi69 + -364)), F2: 128}
                                _ = v101
                                _ = v102
                            } else {
                                switch fi69 {
                                case -309:
                                    var v103 uint = uint(v1)
                                    _ = v103
                                case 4294967295:
                                    if (!false) {
                                        switch 2147483647 {
                                        case 652:
                                            switch fi69 {
                                            case 2147483647:
                                                var v104 uint32 = (v1 >> ((v68) & 63))
                                                _ = v104
                                            case 4294967295:
                                                fmt.Println(fi69, v97, v4)
                                            }
                                            v5 = (-(v3))
                                        case -827:
                                            v105 := GenS11{F0: -60, F1: ("hello" + "\n\t"), F2: ((("" + "hello") + "hello") + "foo bar")}
                                            _ = v105
                                        default:
                                            if (0.0 > -2.25) {
                                                v97 -= (v97 * (^(v97)))
                                                var v106 int16 = 0
                                                _ = v106
                                            } else {
                                                switch fi69 {
                                                case -1:
                                                    var v107 uint16 = 65535
                                                    var v108 float32 = float32(1e10)
                                                    _ = v107
                                                    _ = v108
                                                case 127:
                                                    v109 := GenS3{F0: 64}
                                                    _ = v109
                                                case 761:
                                                    switch (fi96 ^ (fi69 & 1)) {
                                                    case -577:
                                                        v110 := GenS2{F0: -1, F1: fi69, F2: 128}
                                                        _ = v110
                                                    default:
                                                        v97 += (v97 & v97)
                                                        fmt.Println(v3, fi96)
                                                    }
                                                    v111 := GenS9{F0: 1000, F1: (v68 & (v68 - v68)), F2: v5}
                                                    if true {
                                                        v3 = int64(v4)
                                                    } else {
                                                        v68 = v68
                                                        fmt.Println(v30, v30)
                                                    }
                                                    _ = v111
                                                }
                                            }
                                        }
                                    } else {
                                        v112 := GenS12{F0: v4, F1: 255, F2: (v1 >> ((v68) & 63))}
                                        v113 := GenS9{F0: 1, F1: v68, F2: (v3 & (v5 * v5))}
                                        v6[(v1 + v1)] = "a"
                                        _ = v112
                                        _ = v113
                                    }
                                case 128:
                                    v114 := GenS3{F0: -80}
                                    _ = v114
                                }
                            }
                        }
                    } else {
                        v115 := []int{(fi69 | fi96), (-(fi69)), int(v68)}
                        fmt.Println(v3, v1, v3)
                        _ = v115
                    }
                    v6[v1] = "foo bar"
                    v3 = -1
                    _ = fi96
                    _ = v97
                }
            case 32768:
                if (!false) {
                    v116 := []int{}
                    _ = v116
                } else {
                    for fi117 := 0; fi117 < 20; fi117++ {
                        var v118 float32 = (v30 * ((((v30 * v30) + (float32((1e10 * 0.1)) * 1.5)) - v30) + -2.25))
                        var v119 uint8 = 151
                        _ = fi117
                        _ = v118
                        _ = v119
                    }
                }
            default:
                v120 := GenS4{F0: 1, F1: (1 >= 127)}
                _ = v120
            }
        }
        _ = fi69
    }
    for fi121 := 0; fi121 < 3; fi121++ {
        var v122 string = "foo bar"
        var v123 int8 = 42
        _ = fi121
        _ = v122
        _ = v123
    }
    for fi124 := 0; fi124 < 18; fi124++ {
        v125 := v6[42]
        _ = fi124
        _ = v125
    }
    v126 := GenS8{F0: 2147483647}
    for fi127 := 0; fi127 < 11; fi127++ {
        v128 := GenS9{F0: 1000, F1: (v68 + (v68 << ((42) & 63))), F2: v3}
        _ = fi127
        _ = v128
    }
    var v129 bool = ((v2 & 0) <= 18446744073709551615)
    for fi130 := 0; fi130 < 19; fi130++ {
        for fi131 := 0; fi131 < 13; fi131++ {
            switch "a" {
            case "hello":
                switch (fi130 / ((1) | 1)) {
                case 4294967295:
                    var v132 uint8 = 127
                    if v129 {
                        v133 := v6[2]
                        delete(v6, (v1 + v1))
                        var v134 int8 = -1
                        _ = v133
                        _ = v134
                    }
                    _ = v132
                case 124:
                    var v135 int = (fi131 | -9223372036854775808)
                    _ = v135
                default:
                    if (v129 && v129) {
                        v136 := []int{}
                        var v137 string = ("" + ("a" + "foo bar"))
                        v3 -= v5
                        if (!(v68 >= v68)) {
                            v6[(v1 - 2147483648)] = "\n\t"
                            if false {
                                if v129 {
                                    if (906 >= 1) {
                                        v138 := v6[v1]
                                        if v129 {
                                            v139 := GenS7{F0: 473, F1: (v68 == (^(v68)))}
                                            var v140 uint8 = 0
                                            _ = v139
                                            _ = v140
                                        }
                                        fmt.Println(fi131, v129, v5)
                                        v30 = 1e10
                                        v129 = v129
                                        _ = v138
                                    }
                                } else {
                                    fmt.Println(v4, fi131)
                                    fmt.Println(fi130, v2)
                                }
                            } else {
                                _ = len(v6)
                                var v141 int32 = -61
                                v4 += v4
                                _ = v141
                            }
                        } else {
                            var v142 string = ("" + ("hello" + (v137 + (v137 + ("a" + v137)))))
                            _ = len(v136)
                            _ = v142
                        }
                        v137 = ("\n\t" + ((v137 + "foo bar") + (v137 + (v137 + v137))))
                        var v143 string = "hello"
                        v4 += -1
                        if (v4 >= (-(v4))) {
                            v144 := v6[2147483648]
                            if (18446744073709551615 == (v2 ^ 4294967295)) {
                                v145 := GenS13{F0: 1, F1: 42, F2: -676}
                                fmt.Println(v2, v137)
                                _ = v145
                            } else {
                                var v146 uint32 = v1
                                _ = v146
                            }
                            v147 := GenS1{F0: v143, F1: 611}
                            _ = v144
                            _ = v147
                        } else {
                            delete(v6, v1)
                            v148 := GenS7{F0: 505, F1: false}
                            _ = v148
                        }
                        _ = v136
                        _ = v137
                        _ = v143
                    } else {
                        switch ((("a" + "foo bar") + "\n\t") + ("a" + ("\n\t" + ""))) {
                        case "\n\t":
                            fmt.Println(v129, v68)
                            if false {
                                v1 += v1
                            } else {
                                var v149 uint16 = 593
                                if (v1 <= v1) {
                                    fmt.Println(v42)
                                } else {
                                    if v129 {
                                        v42 = 2
                                        v6[2] = (("hello" + "") + "a")
                                        _ = len(v6)
                                    }
                                    if (v4 < (v4 | v4)) {
                                        var v150 uint = (v68 * (v68 % ((uint(v2)) | 1)))
                                        var v151 uint = (v68 / ((65535) | 1))
                                        var v152 uint32 = 228
                                        v151 = 42
                                        _ = v150
                                        _ = v151
                                        _ = v152
                                    } else {
                                        var v153 int32 = v4
                                        switch int(v2) {
                                        case -427:
                                            v6[v1] = ""
                                        case 128:
                                            fmt.Println(v1)
                                        case 127:
                                            if v129 {
                                                v149 = (v149 >> ((v68) & 63))
                                            }
                                        default:
                                            if (1e10 <= 0.0) {
                                                switch fi130 {
                                                case 32768:
                                                    if (v30 > v30) {
                                                        if (((v30 + 1e10) == v30) || v129) {
                                                            v154 := v6[v1]
                                                            v155 := []int{}
                                                            _ = v154
                                                            _ = v155
                                                        }
                                                        v30 = float32(1e10)
                                                    } else {
                                                        v2 = 18446744073709551615
                                                        var v156 uint8 = 168
                                                        _ = v156
                                                    }
                                                case 315:
                                                    if (!true) {
                                                        if v129 {
                                                            v157 := []int{(^(fi130)), 4294967295, fi131, fi131}
                                                            _ = v157
                                                        }
                                                        if v129 {
                                                            fmt.Println(fi131, v5)
                                                        } else {
                                                            var v158 int16 = -32768
                                                            if (fi130 > (^(fi130))) {
                                                                v153 = v153
                                                            }
                                                            switch ("" + "foo bar") {
                                                            case "hello":
                                                                v153 += v153
                                                            case "":
                                                                v159 := GenS8{F0: fi131}
                                                                _ = v159
                                                            default:
                                                                var v160 uint32 = 42
                                                                _ = v160
                                                            }
                                                            _ = v158
                                                        }
                                                    } else {
                                                        v161 := GenS14{F0: (-(fi130)), F1: 42, F2: "a"}
                                                        var v162 int16 = 1000
                                                        var v163 uint = (v68 & (v68 * uint(v1)))
                                                        _ = v161
                                                        _ = v162
                                                        _ = v163
                                                    }
                                                case 9223372036854775807:
                                                    if (-1 < 0) {
                                                        v68 = 18446744073709551615
                                                    }
                                                }
                                                var v164 float32 = 1.5
                                                _ = v164
                                            } else {
                                                fmt.Println(v149, v149, v5)
                                                v129 = (74 == int8(v1))
                                                v165 := []int{fi131, int(v1), int(v2)}
                                                _ = v165
                                            }
                                        }
                                        var v166 int = 32768
                                        var v167 bool = v129
                                        v30 = v30
                                        v168 := GenS6{F0: 0, F1: 128, F2: v4}
                                        _ = len(v6)
                                        v169 := GenS15{F0: v167}
                                        var v170 int64 = v3
                                        _ = v153
                                        _ = v166
                                        _ = v167
                                        _ = v168
                                        _ = v169
                                        _ = v170
                                    }
                                    _ = len(v6)
                                }
                                _ = v149
                            }
                            v171 := GenS10{F0: (v68 ^ 65535)}
                            _ = v171
                        case "":
                            v172 := []int{}
                            _ = v172
                        }
                        if ((1 == 0) || v129) {
                            v3 -= 42
                        } else {
                            if (v1 >= (v1 ^ 42)) {
                                if (13 >= int8(fi130)) {
                                    v173 := GenS1{F0: ((("\n\t" + ("hello" + "a")) + "foo bar") + "hello"), F1: 22}
                                    _ = v173
                                }
                                fmt.Println(fi131, v30)
                            } else {
                                v2 = (v42 % ((v42) | 1))
                            }
                            v4 = (^(v4))
                        }
                        var v174 string = ("\n\t" + "")
                        v6[(v1 & (v1 % ((v1) | 1)))] = (v174 + v174)
                        _ = v174
                    }
                }
            case "a":
                fmt.Println(v4)
            case "\n\t":
                fmt.Println(fi130)
            }
            v3 += v5
            v5 += (v5 % ((-9223372036854775808) | 1))
            _ = fi131
        }
        _ = fi130
    }
    v6[v1] = "\n\t"
    fmt.Println("v1:", v1)
    fmt.Println("v129:", v129)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("v4:", v4)
    fmt.Println("v42:", v42)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    fmt.Println("v68:", v68)
    _ = v1
    _ = v126
    _ = v129
    _ = v2
    _ = v3
    _ = v30
    _ = v31
    _ = v4
    _ = v42
    _ = v5
    _ = v6
    _ = v68
}
