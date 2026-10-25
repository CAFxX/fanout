package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 uint16 }
type GenS2 struct { F0 int }
type GenS3 struct { F0 int8; F1 uint32; F2 uint8 }
type GenS4 struct { F0 uint32; F1 uint8; F2 int64 }
type GenS5 struct { F0 uint; F1 int64 }
type GenS6 struct { F0 int; F1 string }
type GenS7 struct { F0 uint8; F1 bool }
type GenS8 struct { F0 int16; F1 bool; F2 bool }
type GenS9 struct { F0 int16; F1 int8 }
type GenS10 struct { F0 int32; F1 uint }
type GenS11 struct { F0 uint; F1 int }
type GenS12 struct { F0 uint64; F1 uint32; F2 uint16 }
type GenS13 struct { F0 bool; F1 uint; F2 uint16 }
type GenS14 struct { F0 int64; F1 int64 }
type GenS15 struct { F0 string }
type GenS16 struct { F0 int8 }
type GenS17 struct { F0 bool }
type GenS18 struct { F0 int }
type GenS19 struct { F0 uint16; F1 int32 }
func genFn0() int64 {
    var v7 uint16 = 0
    if (42 < 0) {
        v7 += (v7 % (((v7 | v7)) | 1))
        v7 += 255
    } else {
        var v8 uint16 = (v7 << ((2) & 63))
        _ = v8
    }
    for fi9 := 0; fi9 < 13; fi9++ {
        switch fi9 {
        case 0:
            v10 := GenS1{F0: 1000, F1: v7}
            _ = v10
        case -1:
            v7 -= v7
        case 9223372036854775807:
            if (32767 == -1) {
                v7 -= 1
                v7 = v7
            } else {
                for fi11 := 0; fi11 < 14; fi11++ {
                    v12 := GenS2{F0: (-(fi9))}
                    v13 := []int{fi9, (fi11 + (fi11 / (((fi9 | (^(fi9)))) | 1))), fi9}
                    var v14 uint8 = 1
                    _ = fi11
                    _ = v12
                    _ = v13
                    _ = v14
                }
                for fi15 := 0; fi15 < 15; fi15++ {
                    v16 := []int{(fi9 >> ((0) & 63)), fi9, fi15, fi9}
                    _ = fi15
                    _ = v16
                }
                if true {
                    v17 := GenS2{F0: 128}
                    _ = v17
                } else {
                    if true {
                        for fi18 := 0; fi18 < 5; fi18++ {
                            v7 += v7
                            if (uint64(v7) <= 18446744073709551615) {
                                v19 := []int{fi9, (^(fi9)), (fi18 - -2147483648)}
                                switch "" {
                                case "foo bar":
                                    var v20 uint = 42
                                    var v21 int64 = 118
                                    v21 += (v21 >> ((v20) & 63))
                                    v21 += v21
                                    v7 -= (v7 - (v7 % ((226) | 1)))
                                    _ = v20
                                    _ = v21
                                }
                                _ = v19
                            } else {
                                var v22 float32 = float32(1.5)
                                var v23 float32 = (3.14159 / v22)
                                _ = v22
                                _ = v23
                            }
                            v7 = (v7 - (^(v7)))
                            v7 = 429
                            v24 := []int{(fi9 | (fi9 * fi18)), fi18, 4294967295, fi9}
                            v25 := GenS3{F0: 20, F1: 1, F2: 217}
                            _ = fi18
                            _ = v24
                            _ = v25
                        }
                        for fi26 := 0; fi26 < 20; fi26++ {
                            fmt.Println(v7, fi9, fi9)
                            v27 := []int{fi26, (-(fi9)), fi9}
                            v7 = 42
                            _ = fi26
                            _ = v27
                        }
                        var v28 bool = (42 <= 2147483647)
                        _ = v28
                    } else {
                        switch -718 {
                        case -500:
                            var v29 uint = 957
                            _ = v29
                        case -17:
                            v30 := []int{(fi9 & (fi9 ^ 162)), fi9, -9223372036854775808}
                            _ = v30
                        }
                        var v31 int16 = -32768
                        _ = v31
                    }
                }
            }
        default:
            if true {
                switch fi9 {
                case 32767:
                    v32 := GenS1{F0: 599, F1: (v7 + (v7 ^ 310))}
                    _ = v32
                case -2147483648:
                    v33 := GenS3{F0: 100, F1: 65535, F2: 127}
                    _ = v33
                }
                v34 := []int{fi9, 0}
                _ = v34
            } else {
                var v35 uint8 = 128
                for fi36 := 0; fi36 < 12; fi36++ {
                    v35 = 255
                    v37 := GenS4{F0: 433, F1: (^(v35)), F2: int64(v7)}
                    v35 += v35
                    _ = fi36
                    _ = v37
                }
                v38 := []int{(fi9 * -431)}
                _ = v35
                _ = v38
            }
        }
        _ = fi9
    }
    _ = v7
    return -1
}

func main() {
    var v1 uint8 = 62
    var v2 float32 = float32(3.14159)
    var v3 int32 = -463
    v4 := make(map[int]bool)
    v5 := []string{"hello", "hello", "\n\t", "foo bar"}
    var v6 float64 = (0.0 + ((3.14159 * 0.1) - 0.0))
    if false {
        delete(v4, -709)
        for fi39 := 0; fi39 < 19; fi39++ {
            for fi40 := 0; fi40 < 8; fi40++ {
                switch "a" {
                case "":
                    var v41 int = fi40
                    _ = v41
                case "foo bar":
                    if true {
                        if true {
                            v42 := v4[(fi39 % (((fi40 | fi39)) | 1))]
                            v43 := GenS3{F0: -104, F1: 383, F2: v1}
                            fmt.Println(v3)
                            _ = v42
                            _ = v43
                        }
                        v44 := GenS3{F0: -128, F1: 42, F2: v1}
                        if false {
                            var v45 string = (("foo bar" + "a") + "")
                            var v46 bool = true
                            v46 = v46
                            switch "foo bar" {
                            case "":
                                v47 := GenS5{F0: 419, F1: -9223372036854775808}
                                _ = v47
                            default:
                                v4[(fi40 & 32768)] = (1e10 == v2)
                            }
                            _ = v45
                            _ = v46
                        }
                        _ = v44
                    } else {
                        var v48 int32 = 1
                        v2 += v2
                        _ = v48
                    }
                default:
                    var v49 uint32 = 0
                    _ = v49
                }
                v50 := GenS6{F0: fi39, F1: ""}
                v51 := GenS7{F0: v1, F1: true}
                _ = fi40
                _ = v50
                _ = v51
            }
            for fi52 := 0; fi52 < 2; fi52++ {
                delete(v4, fi52)
                if false {
                    var v53 float32 = v2
                    _ = v53
                }
                _ = fi52
            }
            _ = fi39
        }
    } else {
        fmt.Println(v2, v2, v1)
        var v54 uint16 = 42
        _ = v54
    }
    for fi55 := 0; fi55 < 20; fi55++ {
        v1 = v1
        _ = fi55
    }
    _ = len(v5)
    for fi56 := 0; fi56 < 13; fi56++ {
        delete(v4, (fi56 ^ 1))
        var v57 uint16 = 830
        _ = fi56
        _ = v57
    }
    for fi58 := 0; fi58 < 14; fi58++ {
        if false {
            v59 := GenS5{F0: 258, F1: 9223372036854775807}
            var v60 int = (^(fi58))
            if false {
                if len(v5) > 0 {
                    v5[0] = "hello"
                }
                var v61 float32 = float32(float64(v2))
                for fi62 := 0; fi62 < 8; fi62++ {
                    v4[(-(fi62))] = ((("" + "hello") > ("hello" + "hello")) && false)
                    if (false && (!(v3 > 42))) {
                        v63 := GenS5{F0: 37, F1: genFn0()}
                        var v64 uint16 = 2
                        v1 -= v1
                        _ = v63
                        _ = v64
                    }
                    if ((v61 >= ((float32(0.1) + v61) / (3.14159 - 1.5))) || true) {
                        var v65 float64 = v6
                        _ = v65
                    } else {
                        switch "a" {
                        case "a":
                            _ = len(v5)
                        case "foo bar":
                            switch "a" {
                            case "a":
                                if true {
                                    fmt.Println(v61, v1)
                                    fmt.Println(v1, v60)
                                    var v66 int8 = 26
                                    v67 := v4[9223372036854775807]
                                    _ = len(v5)
                                    _ = v66
                                    _ = v67
                                } else {
                                    v60 += int(v3)
                                }
                            case "hello":
                                if (1 > -94) {
                                    var v68 int64 = 0
                                    _ = len(v4)
                                    _ = v68
                                }
                            case "foo bar":
                                var v69 uint16 = 58
                                _ = v69
                            default:
                                switch 0 {
                                case 467:
                                    if (v6 == v6) {
                                        v2 = v2
                                        v61 -= v2
                                        var v70 float64 = float64(0.1)
                                        if (842 != 2) {
                                            v71 := GenS1{F0: 1, F1: 65535}
                                            v72 := GenS6{F0: fi58, F1: ((("\n\t" + "foo bar") + "a") + ((("hello" + ("\n\t" + "foo bar")) + "") + ("a" + "\n\t")))}
                                            v6 += (float64(float32(float64(v2))) + v70)
                                            _ = v71
                                            _ = v72
                                        }
                                        _ = v70
                                    } else {
                                        fmt.Println(v1, fi58, v61)
                                    }
                                case 2147483647:
                                    v6 = float64(3.14159)
                                default:
                                    var v73 int16 = -1
                                    _ = v73
                                }
                            }
                        }
                        switch ("a" + "hello") {
                        case "foo bar":
                            if false {
                                var v74 int = (fi58 >> ((65535) & 63))
                                switch (-(v60)) {
                                case -2147483648:
                                    _ = len(v5)
                                default:
                                    v60 -= (v74 + fi62)
                                }
                                fmt.Println(v3, v61, v60)
                                v1 = v1
                                _ = v74
                            }
                        case "hello":
                            v75 := GenS3{F0: -109, F1: 4294967295, F2: (v1 << ((4294967295) & 63))}
                            _ = v75
                        case "\n\t":
                            var v76 int64 = -704
                            _ = v76
                        default:
                            var v77 int16 = 1000
                            _ = v77
                        }
                    }
                    _ = fi62
                }
                _ = v61
            }
            _ = v59
            _ = v60
        }
        _ = fi58
    }
    switch "foo bar" {
    case "hello":
        switch -147 {
        case 127:
            v78 := GenS8{F0: -1, F1: (1 != 42), F2: (false && false)}
            _ = v78
        case 9223372036854775807:
            var v79 bool = true
            _ = v79
        default:
            v2 += (v2 * (1e10 + 0.0))
        }
    case "foo bar":
        if len(v5) > 0 {
            v5[0] = "a"
        }
    }
    if (173 != 0) {
        switch -2147483648 {
        case 155:
            var v80 bool = false
            _ = v80
        case -9:
            v3 = 1
        case -1:
            for fi81 := 0; fi81 < 17; fi81++ {
                v82 := GenS2{F0: fi81}
                v83 := GenS9{F0: 42, F1: 42}
                for fi84 := 0; fi84 < 3; fi84++ {
                    if true {
                        switch -9223372036854775808 {
                        case 4294967295:
                            var v85 int = -2147483648
                            _ = v85
                        case 1:
                            v86 := GenS5{F0: 1, F1: -689}
                            _ = v86
                        case 327:
                            var v87 string = "foo bar"
                            _ = v87
                        }
                    } else {
                        fmt.Println(v1, v6, v6)
                        var v88 int64 = 0
                        _ = v88
                    }
                    if (42 != 0) {
                        fmt.Println(v1)
                        delete(v4, fi84)
                    } else {
                        v89 := GenS3{F0: -24, F1: 42, F2: v1}
                        _ = v89
                    }
                    fmt.Println(v1)
                    _ = fi84
                }
                _ = fi81
                _ = v82
                _ = v83
            }
        default:
            if true {
                var v90 float64 = (v6 - v6)
                for fi91 := 0; fi91 < 15; fi91++ {
                    fmt.Println(v6, v3, v6)
                    for fi92 := 0; fi92 < 17; fi92++ {
                        v93 := GenS10{F0: (v3 >> ((1) & 63)), F1: 18446744073709551615}
                        var v94 uint16 = 0
                        _ = fi92
                        _ = v93
                        _ = v94
                    }
                    _ = fi91
                }
                var v95 uint64 = 2
                _ = v90
                _ = v95
            } else {
                for fi96 := 0; fi96 < 19; fi96++ {
                    if true {
                        v97 := GenS6{F0: (fi96 / ((-329) | 1)), F1: ("" + "foo bar")}
                        switch (fi96 * fi96) {
                        case 127:
                            fmt.Println(v2)
                        case 1:
                            for fi98 := 0; fi98 < 11; fi98++ {
                                var v99 string = "a"
                                v100 := GenS8{F0: 1000, F1: (!(false && true)), F2: (v2 < v2)}
                                _ = fi98
                                _ = v100
                                _ = v99
                            }
                        }
                        fmt.Println(v2, fi96)
                        _ = v97
                    } else {
                        if false {
                            if true {
                                fmt.Println(fi96, fi96, v2)
                                for fi101 := 0; fi101 < 5; fi101++ {
                                    v6 += 0.0
                                    v102 := GenS9{F0: 42, F1: 38}
                                    _ = len(v5)
                                    _ = fi101
                                    _ = v102
                                }
                                v1 += v1
                            }
                        } else {
                            var v103 float64 = float64(v2)
                            v104 := GenS11{F0: 18446744073709551615, F1: (-(fi96))}
                            var v105 uint = 480
                            _ = v103
                            _ = v104
                            _ = v105
                        }
                        var v106 string = "\n\t"
                        _ = v106
                    }
                    for fi107 := 0; fi107 < 8; fi107++ {
                        switch ("\n\t" + "\n\t") {
                        case "a":
                            v108 := GenS12{F0: 681, F1: 42, F2: 278}
                            _ = v108
                        case "\n\t":
                            fmt.Println(v2, v2, v2)
                        default:
                            if ((((-2.25 / float32(float64(v2))) - (v2 * float32(v6))) * 3.14159) == v2) {
                                var v109 int16 = -547
                                if ((fi107 ^ fi96) == 802) {
                                    var v110 uint8 = (v1 / ((v1) | 1))
                                    _ = v110
                                }
                                _ = v109
                            } else {
                                if (0 > 42) {
                                    var v111 string = ("hello" + ("foo bar" + "a"))
                                    _ = v111
                                } else {
                                    v112 := GenS5{F0: 18446744073709551615, F1: -9223372036854775808}
                                    v2 += v2
                                    _ = v112
                                }
                                switch fi107 {
                                case 2147483647:
                                    var v113 int = fi107
                                    _ = v113
                                case -623:
                                    var v114 int32 = (v3 ^ (v3 / (((v3 - -617)) | 1)))
                                    v115 := GenS8{F0: -180, F1: ("" >= "foo bar"), F2: (42 <= 1)}
                                    fmt.Println(v114, fi96, fi96)
                                    _ = v114
                                    _ = v115
                                case 142:
                                    switch int(v3) {
                                    case 97:
                                        var v116 int = (fi96 % ((2147483647) | 1))
                                        _ = v116
                                    case 32768:
                                        var v117 int16 = -203
                                        var v118 int32 = 1
                                        _ = v117
                                        _ = v118
                                    case -1:
                                        v119 := GenS10{F0: (v3 + 477), F1: 346}
                                        _ = v119
                                    default:
                                        var v120 int16 = 996
                                        _ = v120
                                    }
                                    if (true || false) {
                                        if (((v1 ^ 2) > (v1 - v1)) || true) {
                                            var v121 uint = 487
                                            v122 := GenS6{F0: fi96, F1: (("foo bar" + "a") + "hello")}
                                            var v123 uint = 706
                                            _ = v121
                                            _ = v122
                                            _ = v123
                                        } else {
                                            if (0.0 != ((float32(3.14159) / 1e10) - (((v2 / 1e10) / 1.0) - 1.5))) {
                                                delete(v4, (fi96 ^ (fi107 + -209)))
                                            } else {
                                                v124 := GenS2{F0: 9223372036854775807}
                                                switch ("a" + "") {
                                                case "hello":
                                                    v1 -= 128
                                                default:
                                                    if len(v5) >= 2 {
                                                        v5 = v5[1:len(v5)]
                                                    }
                                                }
                                                if false {
                                                    if false {
                                                        v3 = (v3 + (v3 << ((0) & 63)))
                                                        if len(v5) > 0 {
                                                            v5[0] = (("" + ("hello" + ("a" + ("\n\t" + "foo bar")))) + "\n\t")
                                                        }
                                                    }
                                                    fmt.Println(v2, fi107, v1)
                                                } else {
                                                    var v125 bool = true
                                                    _ = v125
                                                }
                                                _ = v124
                                            }
                                            switch (fi107 - fi107) {
                                            case 32768:
                                                switch -2147483648 {
                                                case -1:
                                                    v2 = (-2.25 * 0.1)
                                                case -697:
                                                    v126 := GenS13{F0: (-2.25 <= v6), F1: 255, F2: 334}
                                                    _ = v126
                                                case 2147483647:
                                                    var v127 uint = 2
                                                    _ = v127
                                                }
                                            }
                                            if false {
                                                if false {
                                                    if false {
                                                        v6 -= (v6 + float64(v2))
                                                        v2 += v2
                                                        if (!false) {
                                                            v1 -= 255
                                                            v1 = (v1 / (((^(v1))) | 1))
                                                            v6 -= v6
                                                        }
                                                    }
                                                }
                                                switch 1 {
                                                case -2147483648:
                                                    if false {
                                                        if (!(false && ((true && ((((v6 + float64((float32(float64(1e10)) + 3.14159))) + 1e10) > ((v6 * v6) - (v6 - ((v6 / -2.25) / 1e10)))) || (("hello" + "a") <= ""))) && (!false)))) {
                                                            v128 := GenS14{F0: -358, F1: genFn0()}
                                                            var v129 string = (("a" + "") + "a")
                                                            _ = v128
                                                            _ = v129
                                                        }
                                                        v1 += (v1 / ((v1) | 1))
                                                        var v130 bool = false
                                                        _ = v130
                                                    }
                                                case 2147483647:
                                                    _ = len(v4)
                                                }
                                            }
                                        }
                                        if len(v5) > 0 {
                                            v5[0] = ""
                                        }
                                    }
                                    fmt.Println(v1, fi107)
                                }
                                if (false || true) {
                                    switch (fi107 >> ((42) & 63)) {
                                    case -281:
                                        v4[(fi107 << ((485) & 63))] = true
                                    case 1:
                                        v2 = 0.1
                                    }
                                    if len(v5) >= 2 {
                                        v5 = v5[1:len(v5)]
                                    }
                                    var v131 uint64 = 0
                                    var v132 bool = ((uint16(v131) > 65535) || (!(!true)))
                                    v133 := GenS15{F0: "foo bar"}
                                    fmt.Println(v132, v1, v132)
                                    var v134 uint64 = v131
                                    if ((false && v132) || (2147483647 > 2147483647)) {
                                        v132 = true
                                    } else {
                                        v135 := GenS1{F0: 1000, F1: 42}
                                        v131 += 0
                                        _ = v135
                                    }
                                    fmt.Println(v2)
                                    v6 = (v6 + v6)
                                    if len(v5) >= 2 {
                                        v5 = v5[1:len(v5)]
                                    }
                                    _ = v131
                                    _ = v132
                                    _ = v133
                                    _ = v134
                                } else {
                                    var v136 uint8 = v1
                                    var v137 uint16 = 903
                                    var v138 uint16 = 741
                                    _ = v136
                                    _ = v137
                                    _ = v138
                                }
                            }
                        }
                        fmt.Println(v2, v6)
                        v139 := GenS16{F0: -1}
                        _ = fi107
                        _ = v139
                    }
                    v140 := GenS17{F0: (true && false)}
                    _ = fi96
                    _ = v140
                }
            }
        }
        var v141 bool = true
        v142 := GenS18{F0: 9223372036854775807}
        _ = v141
        _ = v142
    }
    for fi143 := 0; fi143 < 1; fi143++ {
        switch 9223372036854775807 {
        case 32768:
            for fi144 := 0; fi144 < 7; fi144++ {
                var v145 string = "foo bar"
                switch (("\n\t" + "hello") + v145) {
                case "hello":
                    var v146 int32 = (v3 - (v3 ^ v3))
                    _ = v146
                case "foo bar":
                    if (v3 == (v3 & (v3 << ((2) & 63)))) {
                        v1 = (^(v1))
                    } else {
                        var v147 int8 = -1
                        if len(v5) > 0 {
                            v5[0] = v145
                        }
                        v148 := GenS19{F0: 2, F1: v3}
                        _ = v147
                        _ = v148
                    }
                case "a":
                    v5 = append(v5, (v145 + "a"))
                }
                _ = fi144
                _ = v145
            }
        case 0:
            v2 = 3.14159
        default:
            for fi149 := 0; fi149 < 12; fi149++ {
                v1 += (^(v1))
                fmt.Println(v6, v1, v1)
                delete(v4, (fi149 % ((9223372036854775807) | 1)))
                _ = fi149
            }
        }
        _ = fi143
    }
    var v150 int32 = v3
    for fi151 := 0; fi151 < 11; fi151++ {
        for fi152 := 0; fi152 < 13; fi152++ {
            var v153 uint64 = 2
            switch "hello" {
            case "foo bar":
                v3 += (v150 | v150)
                v153 += v153
                v2 = v2
            default:
                if (v2 <= v2) {
                    v150 -= (v150 + int32(fi151))
                } else {
                    if true {
                        var v154 float32 = 3.14159
                        if (42 <= 2) {
                            if false {
                                switch "\n\t" {
                                case "\n\t":
                                    v150 = v3
                                }
                            }
                            delete(v4, -982)
                            var v155 int32 = (v3 % (((v3 ^ (v3 & (v150 | (v3 * v3))))) | 1))
                            _ = v155
                        }
                        v150 = (v150 & v150)
                        _ = v154
                    } else {
                        v153 = v153
                        delete(v4, (fi152 - fi152))
                    }
                    switch fi152 {
                    case -418:
                        var v156 uint64 = 18446744073709551615
                        _ = v156
                    }
                    v5 = append(v5, "hello")
                }
            }
            fmt.Println(fi151, fi151, v150)
            _ = fi152
            _ = v153
        }
        _ = fi151
    }
    fmt.Println("v1:", v1)
    fmt.Println("v150:", v150)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v150
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
