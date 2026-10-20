package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 uint }
type GenS2 struct { F0 uint16; F1 int64; F2 int8 }
type GenS3 struct { F0 bool; F1 int8 }
type GenS4 struct { F0 string }
type GenS5 struct { F0 int32 }
type GenS6 struct { F0 string; F1 uint; F2 bool }
type GenS7 struct { F0 uint64; F1 uint8; F2 int64 }
type GenS8 struct { F0 bool; F1 string; F2 int32 }
type GenS9 struct { F0 int64; F1 uint }
type GenS10 struct { F0 string }
type GenS11 struct { F0 int16; F1 int32; F2 int64 }
type GenS12 struct { F0 uint8; F1 int; F2 string }
type GenS13 struct { F0 uint16; F1 uint }
type GenS14 struct { F0 uint16; F1 int32; F2 uint64 }
type GenS15 struct { F0 uint8; F1 uint64; F2 int16 }
type GenS16 struct { F0 int }
func genFn0(p0 int8, p1 int32, p2 uint64) uint32 {
    if true {
        p1 -= (p1 | (-(p1)))
    }
    p2 += p2
    _ = p0
    _ = p1
    _ = p2
    return 42
}
func genFn1(p0 uint16, p1 int, p2 string) uint64 {
    switch 0 {
    case 42:
        v4 := GenS1{F0: 0, F1: 245}
        _ = v4
    case 280:
        for fi5 := 0; fi5 < 16; fi5++ {
            p0 += p0
            v6 := []int{(p1 ^ fi5), -9223372036854775808, (p1 << ((650) & 63))}
            var v7 int32 = 1
            _ = fi5
            _ = v6
            _ = v7
        }
    default:
        v8 := []int{p1, p1}
        _ = v8
    }
    for fi9 := 0; fi9 < 2; fi9++ {
        switch (-(p1)) {
        case 128:
            p0 = p0
        }
        switch p2 {
        case "a":
            v10 := []int{fi9}
            _ = v10
        default:
            fmt.Println(p0, p2, p0)
        }
        var v11 int16 = 1
        _ = fi9
        _ = v11
    }
    if false {
        var v12 uint8 = 212
        var v13 int = -472
        _ = v12
        _ = v13
    } else {
        p0 -= (^(p0))
    }
    var v14 string = "foo bar"
    switch (p1 ^ 128) {
    case 32767:
        p0 = 32767
    case 66:
        var v15 bool = false
        _ = v15
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v14
    return 2
}

func main() {
    var v1 float64 = (0.1 - 0.0)
    v2 := make(map[string]uint32)
    v3 := make(map[uint32]bool)
    for fi16 := 0; fi16 < 19; fi16++ {
        v1 += 3.14159
        v1 -= v1
        _ = fi16
    }
    if true {
        if false {
            v1 -= 3.14159
            if false {
                if true {
                    v17 := GenS1{F0: -62, F1: 2}
                    var v18 int16 = -1
                    switch -663 {
                    case 32767:
                        v18 += (-(v18))
                    default:
                        if (!(!true)) {
                            for fi19 := 0; fi19 < 17; fi19++ {
                                var v20 int16 = v18
                                _ = fi19
                                _ = v20
                            }
                            v18 += v18
                            if (-80 < 42) {
                                for fi21 := 0; fi21 < 10; fi21++ {
                                    delete(v2, ("foo bar" + (("a" + ("foo bar" + "\n\t")) + "\n\t")))
                                    if true {
                                        var v22 uint8 = 128
                                        var v23 float64 = v1
                                        _ = v22
                                        _ = v23
                                    }
                                    _ = fi21
                                }
                                var v24 float32 = 1e10
                                _ = v24
                            }
                        } else {
                            var v25 int8 = 1
                            for fi26 := 0; fi26 < 1; fi26++ {
                                v25 += (v25 / ((-41) | 1))
                                v27 := GenS2{F0: 242, F1: -803, F2: 77}
                                v18 = -1
                                _ = fi26
                                _ = v27
                            }
                            _ = v25
                        }
                    }
                    _ = v17
                    _ = v18
                }
                if true {
                    v1 = (float64(1.5) - (v1 + v1))
                    var v28 uint8 = 12
                    var v29 string = "hello"
                    _ = v28
                    _ = v29
                } else {
                    v30 := []int{-513, -2147483648, 2147483647, -405}
                    if true {
                        v31 := GenS3{F0: false, F1: -121}
                        for fi32 := 0; fi32 < 6; fi32++ {
                            for fi33 := 0; fi33 < 9; fi33++ {
                                switch "a" {
                                case "hello":
                                    fmt.Println(v1)
                                case "foo bar":
                                    var v34 uint64 = 0
                                    _ = v34
                                default:
                                    v1 -= ((v1 + v1) * v1)
                                }
                                v1 = v1
                                _ = fi33
                            }
                            v35 := GenS2{F0: 781, F1: 0, F2: -1}
                            if false {
                                v36 := GenS4{F0: "foo bar"}
                                var v37 uint32 = 724
                                _ = v36
                                _ = v37
                            } else {
                                var v38 uint64 = 4294967295
                                var v39 uint8 = 159
                                fmt.Println(fi32, v38)
                                _ = v38
                                _ = v39
                            }
                            _ = fi32
                            _ = v35
                        }
                        var v40 float64 = 1.5
                        _ = v31
                        _ = v40
                    } else {
                        v1 = v1
                        if false {
                            switch ("\n\t" + (("\n\t" + "foo bar") + ("\n\t" + ("foo bar" + (("hello" + "a") + ("foo bar" + "\n\t")))))) {
                            case "foo bar":
                                var v41 float32 = -2.25
                                _ = v41
                            case "a":
                                for fi42 := 0; fi42 < 16; fi42++ {
                                    if (true && false) {
                                        switch "a" {
                                        case "hello":
                                            v1 = (v1 + ((float64(float32((float64(-2.25) * v1))) + (3.14159 - 0.0)) * v1))
                                        }
                                        fmt.Println(v1, v1)
                                        v1 -= (v1 - float64(0.1))
                                    } else {
                                        switch (fi42 & 329) {
                                        case 9223372036854775807:
                                            fmt.Println(fi42)
                                        case 905:
                                            if true {
                                                if len(v30) > 0 {
                                                    v30[0] = fi42
                                                }
                                                var v43 int16 = -32768
                                                var v44 int8 = 0
                                                _ = v43
                                                _ = v44
                                            } else {
                                                fmt.Println(v1)
                                                switch "foo bar" {
                                                case "foo bar":
                                                    for fi45 := 0; fi45 < 17; fi45++ {
                                                        var v46 uint64 = 410
                                                        fmt.Println(v1, fi45)
                                                        _ = fi45
                                                        _ = v46
                                                    }
                                                case "\n\t":
                                                    var v47 uint64 = 191
                                                    _ = v47
                                                }
                                            }
                                        }
                                        for fi48 := 0; fi48 < 12; fi48++ {
                                            var v49 uint8 = 255
                                            switch ("a" + ("hello" + ("" + ("\n\t" + "a")))) {
                                            case "a":
                                                switch fi42 {
                                                case -453:
                                                    var v50 uint32 = 0
                                                    v51 := GenS5{F0: -404}
                                                    var v52 bool = false
                                                    _ = v50
                                                    _ = v51
                                                    _ = v52
                                                }
                                            case "\n\t":
                                                v49 = v49
                                            case "foo bar":
                                                v2["foo bar"] = 0
                                            }
                                            v1 = ((v1 - 0.1) * (1.5 - v1))
                                            v1 = float64(0.0)
                                            var v53 int = fi48
                                            _ = fi48
                                            _ = v49
                                            _ = v53
                                        }
                                    }
                                    switch (fi42 + fi42) {
                                    case 1:
                                        if false {
                                            v54 := GenS1{F0: -2147483648, F1: 613}
                                            var v55 float64 = float64((-2.25 + 1.5))
                                            _ = v54
                                            _ = v55
                                        }
                                    }
                                    _ = fi42
                                }
                            }
                            switch 32768 {
                            case -9223372036854775808:
                                v1 -= 1e10
                            }
                            delete(v2, "")
                        }
                    }
                    _ = v30
                }
            } else {
                for fi56 := 0; fi56 < 7; fi56++ {
                    var v57 uint32 = 4294967295
                    fmt.Println(v57, v1)
                    _ = fi56
                    _ = v57
                }
                v58 := GenS2{F0: 0, F1: 9223372036854775807, F2: 109}
                v59 := GenS2{F0: 255, F1: -493, F2: -90}
                _ = v58
                _ = v59
            }
        } else {
            v60 := []int{128}
            switch -9223372036854775808 {
            case 32767:
                v1 = float64((1.5 / ((1e10 * float32(float64(float32(float64(float32((3.14159 + v1))))))) * -2.25)))
            case 4294967295:
                if (18446744073709551615 == 1) {
                    if true {
                        var v61 uint8 = 112
                        _ = v61
                    } else {
                        if (886 <= 0) {
                            switch 128 {
                            case -622:
                                var v62 uint64 = 18446744073709551615
                                _ = v62
                            case -2147483648:
                                v1 -= v1
                            }
                            if true {
                                var v63 int8 = 1
                                _ = v63
                            } else {
                                for fi64 := 0; fi64 < 3; fi64++ {
                                    if false {
                                        v1 += (v1 * float64((0.1 * ((3.14159 * (1e10 + (1e10 - (((0.0 * 0.1) * float32(v1)) / (-2.25 * -2.25))))) + (-2.25 / 0.1)))))
                                        fmt.Println(v1, v1)
                                    }
                                    _ = fi64
                                }
                                v65 := GenS6{F0: "foo bar", F1: 1, F2: false}
                                _ = len(v2)
                                _ = v65
                            }
                        }
                        var v66 int8 = -128
                        _ = v66
                    }
                    v1 -= v1
                    var v67 int = 9223372036854775807
                    _ = v67
                }
            }
            var v68 uint = 238
            _ = v60
            _ = v68
        }
        for fi69 := 0; fi69 < 1; fi69++ {
            var v70 uint = 1
            _ = fi69
            _ = v70
        }
        v71 := GenS3{F0: true, F1: 47}
        _ = v71
    }
    if (42 == 437) {
        v72 := GenS2{F0: 1, F1: 0, F2: -96}
        if (0 >= 32767) {
            var v73 int8 = 1
            _ = v73
        } else {
            v1 = v1
            _ = len(v3)
            v1 = (((0.1 - v1) * ((v1 * (v1 - 1.5)) * v1)) + (v1 - 1.5))
        }
        fmt.Println(v1)
        _ = v72
    }
    v74 := GenS7{F0: 18446744073709551615, F1: 42, F2: -1}
    var v75 int8 = 42
    v76 := GenS7{F0: 4294967295, F1: 127, F2: -9223372036854775808}
    if false {
        v75 = 1
        v77 := GenS8{F0: (2147483647 <= -2147483648), F1: "foo bar", F2: -1}
        for fi78 := 0; fi78 < 11; fi78++ {
            fmt.Println(fi78)
            for fi79 := 0; fi79 < 17; fi79++ {
                if false {
                    var v80 string = ("\n\t" + ("hello" + ""))
                    var v81 int16 = -32768
                    if false {
                        var v82 bool = true
                        switch 42 {
                        case -2147483648:
                            var v83 float64 = 3.14159
                            v84 := GenS5{F0: -191}
                            if v82 {
                                v85 := []int{fi79, (fi78 | 128), (fi79 & (fi79 ^ fi79))}
                                var v86 float64 = v83
                                var v87 uint8 = 92
                                _ = v85
                                _ = v86
                                _ = v87
                            }
                            _ = v83
                            _ = v84
                        case 4294967295:
                            var v88 float32 = 1.5
                            _ = v88
                        }
                        var v89 int8 = (v75 % ((v75) | 1))
                        delete(v3, 2147483647)
                        _ = v82
                        _ = v89
                    }
                    var v90 float64 = v1
                    v81 += (v81 % ((v81) | 1))
                    _ = v80
                    _ = v81
                    _ = v90
                } else {
                    switch fi78 {
                    case 1:
                        var v91 int64 = 9223372036854775807
                        _ = v91
                    default:
                        var v92 float64 = (v1 - v1)
                        _ = v92
                    }
                    var v93 int64 = -9223372036854775808
                    v75 -= (v75 ^ v75)
                    if (!(true && true)) {
                        v94 := GenS9{F0: (v93 / ((v93) | 1)), F1: 4294967295}
                        v1 = 1e10
                        _ = v94
                    }
                    _ = v93
                }
                var v95 int8 = v75
                _ = fi79
                _ = v95
            }
            _ = fi78
        }
        _ = v77
    } else {
        v96 := []int{42}
        _ = v96
    }
    switch (("a" + "\n\t") + "\n\t") {
    case "\n\t":
        fmt.Println(v75, v75)
    case "foo bar":
        for fi97 := 0; fi97 < 11; fi97++ {
            if false {
                for fi98 := 0; fi98 < 7; fi98++ {
                    v1 = -2.25
                    _ = fi98
                }
                fmt.Println(v1, v1, v1)
                v99 := []int{878, fi97, 2147483647}
                _ = v99
            } else {
                switch fi97 {
                case 9223372036854775807:
                    var v100 uint32 = 42
                    _ = v100
                case 128:
                    if (("\n\t" + "a") <= "\n\t") {
                        for fi101 := 0; fi101 < 8; fi101++ {
                            var v102 uint32 = 98
                            var v103 int32 = -650
                            var v104 float32 = (float32(v1) + (float32(v1) + -2.25))
                            switch "\n\t" {
                            case "hello":
                                if (-946 != 1000) {
                                    var v105 uint64 = 2
                                    _ = v105
                                } else {
                                    if (0 <= 354) {
                                        if true {
                                            v103 = -1
                                            var v106 int64 = 811
                                            _ = v106
                                        }
                                        v107 := v3[v102]
                                        switch "\n\t" {
                                        case "a":
                                            v75 -= 0
                                        case "foo bar":
                                            fmt.Println(v104, v1, v75)
                                        case "\n\t":
                                            v108 := GenS8{F0: v107, F1: ("a" + "a"), F2: -2147483648}
                                            _ = v108
                                        default:
                                            fmt.Println(v102)
                                        }
                                        v1 = (v1 - (-2.25 - (((v1 * float64(v104)) + v1) + v1)))
                                        _ = v107
                                    }
                                    var v109 uint8 = 127
                                    switch "" {
                                    case "":
                                        v104 += 1.5
                                    case "a":
                                        v110 := GenS10{F0: "hello"}
                                        _ = v110
                                    case "\n\t":
                                        if (868 > fi101) {
                                            v111 := GenS10{F0: "\n\t"}
                                            v112 := GenS1{F0: 153, F1: 65535}
                                            _ = v111
                                            _ = v112
                                        }
                                    }
                                    _ = v109
                                }
                            }
                            fmt.Println(v103, fi97)
                            _ = fi101
                            _ = v102
                            _ = v103
                            _ = v104
                        }
                        v113 := GenS11{F0: 998, F1: 0, F2: 1}
                        v114 := []int{(-(fi97))}
                        _ = v113
                        _ = v114
                    } else {
                        fmt.Println(v75, v75)
                        if false {
                            var v115 int = fi97
                            _ = v115
                        } else {
                            for fi116 := 0; fi116 < 13; fi116++ {
                                v117 := GenS12{F0: 128, F1: -1, F2: ""}
                                var v118 string = ("hello" + ("\n\t" + "\n\t"))
                                fmt.Println(v75)
                                var v119 int8 = 20
                                _ = fi116
                                _ = v117
                                _ = v118
                                _ = v119
                            }
                            for fi120 := 0; fi120 < 20; fi120++ {
                                switch "hello" {
                                case "":
                                    v121 := []int{(^(fi97)), (fi97 / ((1) | 1)), (fi120 % (((fi97 >> ((575) & 63))) | 1)), fi120}
                                    _ = v121
                                default:
                                    if true {
                                        fmt.Println(v75)
                                        var v122 int = 0
                                        _ = v122
                                    } else {
                                        fmt.Println(fi97, v1)
                                    }
                                }
                                var v123 uint64 = 2
                                var v124 int = fi97
                                switch "foo bar" {
                                case "a":
                                    fmt.Println(fi120)
                                case "\n\t":
                                    v3[938] = (!false)
                                }
                                if false {
                                    var v125 uint32 = 1
                                    _ = v125
                                } else {
                                    if false {
                                        v1 -= 0.1
                                    } else {
                                        var v126 int32 = 955
                                        _ = v126
                                    }
                                    switch "" {
                                    case "foo bar":
                                        v75 -= v75
                                    case "\n\t":
                                        var v127 int = (^(v124))
                                        _ = v127
                                    case "a":
                                        fmt.Println(v75, v1)
                                        var v128 float32 = 1e10
                                        _ = v128
                                    }
                                }
                                _ = fi120
                                _ = v123
                                _ = v124
                            }
                            if true {
                                if false {
                                    v129 := GenS13{F0: 2, F1: 1}
                                    for fi130 := 0; fi130 < 7; fi130++ {
                                        var v131 int32 = 1
                                        v132 := GenS10{F0: "a"}
                                        var v133 uint = 0
                                        _ = fi130
                                        _ = v131
                                        _ = v132
                                        _ = v133
                                    }
                                    _ = v129
                                }
                            }
                        }
                        var v134 uint32 = 334
                        _ = v134
                    }
                case 2147483647:
                    for fi135 := 0; fi135 < 16; fi135++ {
                        v1 -= ((((v1 + (float64((0.0 * 1e10)) + 1e10)) - v1) * v1) + v1)
                        var v136 bool = (0 > 2)
                        _ = fi135
                        _ = v136
                    }
                default:
                    if true {
                        switch "foo bar" {
                        case "hello":
                            switch (fi97 + int(v75)) {
                            case 32767:
                                v137 := GenS13{F0: 42, F1: 504}
                                _ = v137
                            case 42:
                                var v138 float64 = (v1 / 1.0)
                                _ = v138
                            case 9223372036854775807:
                                var v139 bool = (!((v75 % (((v75 * v75)) | 1)) >= (v75 & v75)))
                                _ = v139
                            default:
                                for fi140 := 0; fi140 < 3; fi140++ {
                                    v141 := []int{fi140, (fi97 ^ (fi140 ^ 2147483647)), (fi97 + (fi140 + (fi140 << ((727) & 63)))), fi97}
                                    v142 := GenS9{F0: 0, F1: 65535}
                                    _ = fi140
                                    _ = v141
                                    _ = v142
                                }
                            }
                        case "a":
                            for fi143 := 0; fi143 < 2; fi143++ {
                                if ((false && (0 != -32768)) && false) {
                                    var v144 int = (^(fi97))
                                    v144 = -9223372036854775808
                                    if true {
                                        var v145 uint32 = 65535
                                        v146 := GenS14{F0: 151, F1: -2147483648, F2: 2}
                                        _ = v145
                                        _ = v146
                                    }
                                    var v147 bool = false
                                    v148 := v3[2147483648]
                                    var v149 int32 = -1
                                    _ = v144
                                    _ = v147
                                    _ = v148
                                    _ = v149
                                } else {
                                    if true {
                                        v1 -= float64(0.1)
                                        v150 := GenS15{F0: 2, F1: 244, F2: 42}
                                        v75 = v75
                                        _ = v150
                                    } else {
                                        if (-725 == 1) {
                                            v1 = 0.1
                                            var v151 int32 = 2147483647
                                            var v152 int = (fi143 ^ -433)
                                            v153 := []int{(fi97 ^ fi143), fi143}
                                            _ = v151
                                            _ = v152
                                            _ = v153
                                        } else {
                                            v1 += ((-2.25 / 3.14159) * v1)
                                            v75 = v75
                                        }
                                        switch (fi143 + (fi97 ^ 4294967295)) {
                                        case 1:
                                            var v154 int = fi143
                                            v154 = (v154 | (-(v154)))
                                            _ = v154
                                        case 600:
                                            v1 += (v1 - (v1 / (((v1 * (v1 / float64((1.5 / 0.1)))) * v1) - v1)))
                                        case 893:
                                            if false {
                                                if true {
                                                    v1 = ((((v1 * (((1e10 / (v1 / 1e10)) / (v1 / (0.0 * v1))) * 0.1)) * v1) + (0.1 - (v1 - ((v1 / v1) - (v1 + (v1 - v1)))))) * 0.0)
                                                    v155 := GenS14{F0: 32767, F1: 1, F2: 42}
                                                    v75 = (^(v75))
                                                    _ = v155
                                                }
                                                if ((128 != 46) || true) {
                                                    fmt.Println(fi143)
                                                    v3[370] = true
                                                } else {
                                                    v75 -= v75
                                                }
                                                v156 := []int{fi97}
                                                _ = v156
                                            } else {
                                                var v157 bool = true
                                                var v158 int64 = 1
                                                v1 -= v1
                                                _ = v157
                                                _ = v158
                                            }
                                        }
                                    }
                                }
                                _ = fi143
                            }
                        case "":
                            delete(v2, ((("foo bar" + (("" + "hello") + "hello")) + "hello") + "hello"))
                        default:
                            for fi159 := 0; fi159 < 12; fi159++ {
                                var v160 float32 = float32((((float64(((-2.25 * 3.14159) + -2.25)) + ((float64(float32(v1)) / v1) * 0.0)) - v1) + ((v1 / (v1 / float64(3.14159))) / v1)))
                                fmt.Println(v1)
                                _ = fi159
                                _ = v160
                            }
                        }
                        v75 += v75
                    } else {
                        for fi161 := 0; fi161 < 7; fi161++ {
                            v162 := GenS16{F0: 131}
                            v163 := GenS12{F0: 2, F1: (^(fi161)), F2: ("foo bar" + ("" + ((("hello" + (("a" + ("a" + "\n\t")) + "a")) + "\n\t") + "hello")))}
                            _ = len(v3)
                            _ = fi161
                            _ = v162
                            _ = v163
                        }
                        var v164 int64 = -411
                        _ = v164
                    }
                }
            }
            v165 := GenS3{F0: true, F1: (v75 - (v75 | -1))}
            _ = fi97
            _ = v165
        }
    case "hello":
        for fi166 := 0; fi166 < 6; fi166++ {
            v1 = v1
            fmt.Println(fi166, fi166, v1)
            v75 = (^(v75))
            _ = fi166
        }
    }
    var v167 string = "hello"
    var v168 string = "hello"
    switch 32767 {
    case 0:
        fmt.Println(v167, v1, v167)
    }
    fmt.Println("v1:", v1)
    fmt.Println("v167:", v167)
    fmt.Println("v168:", v168)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v75:", v75)
    _ = v1
    _ = v167
    _ = v168
    _ = v2
    _ = v3
    _ = v74
    _ = v75
    _ = v76
}
